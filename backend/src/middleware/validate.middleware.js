const { body, query, validationResult } = require("express-validator");

module.exports = {
  validate: (req, res, next) => {
    const errors = validationResult(req).formatWith(({ msg }) => msg);

    if (!errors.isEmpty()) {
      return res.status(400).json({ error: errors.array().join(", ") });
    }

    next();
  },
  validateLogin: () => [
    body("user.email")
      .notEmpty()
      .withMessage("Email is required")
      .bail()
      .isEmail()
      .withMessage("Invalid email"),
    body("user.password")
      .notEmpty()
      .withMessage("Password is required")
      .bail()
      .isString()
      .withMessage("Password should be a string"),
  ],
  validateRegister: () => [
    body("user.email")
      .notEmpty()
      .withMessage("Email is required")
      .bail()
      .isEmail()
      .withMessage("Invalid email"),
    body("user.first_name")
      .notEmpty()
      .withMessage("First name is required")
      .bail()
      .isString()
      .withMessage("First name should be a string"),
    body("user.last_name")
      .notEmpty()
      .withMessage("Last name is required")
      .bail()
      .isString()
      .withMessage("Last name should be a string"),
    body("user.password")
      .notEmpty()
      .withMessage("Password is required")
      .bail()
      .isString()
      .withMessage("Password should be a string"),
  ],
  validateRequest: () => [
    body("request.start_location")
      .notEmpty()
      .withMessage("Start location is required")
      .bail()
      .isString()
      .withMessage("Start location should be a string"),
    body("request.destination")
      .notEmpty()
      .withMessage("Destination is required")
      .bail()
      .isString()
      .withMessage("Destination should be a string"),
    body("request.date_time")
      .notEmpty()
      .withMessage("Date and time is required")
      .bail()
      .isISO8601()
      .withMessage("Invalid date and time format"),
    body("request.purpose")
      .notEmpty()
      .withMessage("Purpose of travel is required")
      .bail()
      .isString()
      .withMessage("Purpose of travel should be a string"),
  ],
  validateRequestsQuery: () => [
    query("limit", "Limit should be a number").optional().isNumeric(),
    query("page", "Page should be a number").optional().isNumeric(),
  ],
  validateUpdateRequestStatus: () => [
    body("status")
      .exists()
      .withMessage("Status is required")
      .bail()
      .isString()
      .withMessage("Status should be a string")
      .bail()
      .isIn(["approved", "rejected"])
      .withMessage("Status can only be approved or rejected"),
  ],
  validateUpdateRequestRemarks: () => [
    body("remarks")
      .exists()
      .withMessage("Remarks is required")
      .bail()
      .isString()
      .withMessage("Remarks should be a string"),
  ],
};
