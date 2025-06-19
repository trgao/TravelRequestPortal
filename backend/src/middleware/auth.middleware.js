const { verify } = require("jsonwebtoken");
const UsersDatabase = require("../database/users.database");

module.exports = {
  authenticateToken: (req, res, next) => {
    const authCookie = req.cookies["token"];

    if (authCookie == null) {
      console.error("There is no access token");
      return res.sendStatus(401);
    }

    verify(authCookie, process.env.JWT_SECRET, async (err, payload) => {
      if (err) {
        console.error("Invalid access token");
        return res.sendStatus(403);
      }

      const lastLoginAt = new Date(
        await UsersDatabase.checkLastLogin(payload.id)
      );
      if (lastLoginAt.getTime() > payload.iat * 1000) {
        return res.clearCookie("token").sendStatus(401);
      }

      req.user = payload;
      next();
    });
  },
  authoriseAdmin: (req, res, next) => {
    if (req.user == null || !req.user.is_admin) {
      console.error("User is not an admin");
      return res.sendStatus(403);
    }

    next();
  },
  authoriseUser: (req, res, next) => {
    if (req.user == null || req.user.id != req.params.id) {
      console.error("User does not have permission to view requests");
      return res.sendStatus(403);
    }

    next();
  },
};
