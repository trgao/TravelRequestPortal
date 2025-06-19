const AdminService = require("../services/admin.service");

class AdminController {
  static async admin(req, res) {
    try {
      const response = await AdminService.getRequests(req.user.id, req.query);
      console.log(
        "AdminController.admin: get travel requests of employees for admin",
        req.user.id
      );

      return res.status(200).json(response);
    } catch (error) {
      console.error("AdminController.admin:", error);
      return res.status(500).json({ error: error.message });
    }
  }
}

module.exports = AdminController;
