const express = require("express");
const RequestController = require("../controllers/request.controller.js");
const EmployeeController = require("../controllers/employee.controller.js");
const AdminController = require("../controllers/admin.controller.js");
const {
  validateRequest,
  validateRequestsQuery,
  validate,
  validateUpdateRequestStatus,
  validateUpdateRequestRemarks,
} = require("../middleware/validate.middleware.js");
const {
  authenticateToken,
  authoriseUser,
  authoriseAdmin,
} = require("../middleware/auth.middleware.js");

const resourceRouter = express.Router();

resourceRouter.post(
  "/request",
  authenticateToken,
  validateRequest(),
  validate,
  RequestController.createRequest
);

resourceRouter.get(
  "/employee/:id",
  authenticateToken,
  authoriseUser,
  validateRequestsQuery(),
  validate,
  EmployeeController.employee
);

resourceRouter.get(
  "/admin/:id",
  authenticateToken,
  authoriseAdmin,
  authoriseUser,
  validateRequestsQuery(),
  validate,
  AdminController.admin
);

resourceRouter.patch(
  "/request/:id/status",
  authenticateToken,
  authoriseAdmin,
  validateUpdateRequestStatus(),
  validate,
  RequestController.updateRequestStatus
);

resourceRouter.patch(
  "/request/:id/remarks",
  authenticateToken,
  authoriseAdmin,
  validateUpdateRequestRemarks(),
  validate,
  RequestController.updateRequestRemarks
);

resourceRouter.get(
  "/summary",
  authenticateToken,
  authoriseAdmin,
  RequestController.getSummary
);

module.exports = resourceRouter;
