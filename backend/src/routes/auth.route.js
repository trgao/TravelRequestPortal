const express = require("express");
const AuthController = require("../controllers/auth.controller.js");
const {
  validateLogin,
  validateRegister,
  validate,
} = require("../middleware/validate.middleware.js");
const {
  authenticateToken,
  authoriseAdmin,
} = require("../middleware/auth.middleware.js");

const authRouter = express.Router();

authRouter.post("/login", validateLogin(), validate, AuthController.login);

authRouter.post(
  "/register",
  validateRegister(),
  validate,
  AuthController.register
);

authRouter.post(
  "/onboard",
  authenticateToken,
  authoriseAdmin,
  validateRegister(),
  validate,
  AuthController.onboard
);

authRouter.post("/logout", AuthController.logout);

module.exports = authRouter;
