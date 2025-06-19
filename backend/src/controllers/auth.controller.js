const AuthService = require("../services/auth.service");

class AuthController {
  static async login(req, res) {
    try {
      const { token, isAdmin, userId } = await AuthService.login(req.body.user);
      console.log(
        "AuthController.login: successfully logged in",
        req.body.user.email
      );

      return res
        .status(200)
        .cookie("token", token, {
          httpOnly: true,
        })
        .json({
          message: "Successfully logged in",
          is_admin: isAdmin,
          user_id: userId,
        });
    } catch (error) {
      console.error("AuthController.login:", error);
      return res.status(500).json({ error: "Incorrect email or password" });
    }
  }

  static async register(req, res) {
    try {
      const { token, isAdmin, userId } = await AuthService.register(
        req.body.user
      );
      console.log(
        "AuthController.register: successfully registered a new admin",
        req.body.user.email
      );

      return res
        .status(200)
        .cookie("token", token, {
          httpOnly: true,
        })
        .json({
          message: "Successfully registered a new account",
          is_admin: isAdmin,
          user_id: userId,
        });
    } catch (error) {
      console.error("AuthController.register:", error);
      return res.status(500).json({ error: error.message });
    }
  }

  static async onboard(req, res) {
    try {
      await AuthService.onboard(req.body.user, req.user.id);
      console.log(
        "AuthController.onboard: successfully registered a new employee",
        req.body.user.email
      );

      return res.status(200).json({
        message: "Successfully registered a new account",
      });
    } catch (error) {
      console.error("AuthController.onboard:", error);
      return res.status(500).json({ error: error.message });
    }
  }

  static async logout(req, res) {
    return res
      .clearCookie("token")
      .status(200)
      .json({ message: "Successfully logged out" });
  }
}

module.exports = AuthController;
