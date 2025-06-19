const EmployeeService = require("../services/employee.service");

class EmployeeController {
  static async employee(req, res) {
    try {
      const response = await EmployeeService.getRequests(
        req.user.id,
        req.query.limit,
        req.query.page
      );
      console.log(
        "EmployeeController.employee: get travel requests for employee",
        req.user.id
      );

      return res.status(200).json(response);
    } catch (error) {
      console.error("EmployeeController.employee:", error);
      return res.status(500).json({ error: error.message });
    }
  }
}

module.exports = EmployeeController;
