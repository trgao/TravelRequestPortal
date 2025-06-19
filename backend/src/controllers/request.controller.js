const RequestService = require("../services/request.service");

class RequestController {
  static async getSummary(req, res) {
    try {
      const response = await RequestService.getSummary(req.user.id);
      console.log(
        "RequestController.getSummary: successfully retrieved summary for",
        req.user.id
      );

      return res.status(200).json(response);
    } catch (error) {
      console.error("RequestController.getSummary:", error);
      return res.status(500).json({ error: error.message });
    }
  }

  static async createRequest(req, res) {
    try {
      await RequestService.createRequest(req.body.request, req.user.id);
      console.log(
        "RequestController.createRequest: successfully created a new travel request"
      );

      return res.status(200).json({
        message: "Successfully submitted travel request",
      });
    } catch (error) {
      console.error("RequestController.createRequest:", error);
      return res.status(500).json({ error: error.message });
    }
  }

  static async updateRequestStatus(req, res) {
    try {
      await RequestService.updateRequestStatus(
        req.params.id,
        req.body.status,
        req.user.id
      );
      console.log(
        "RequestController.updateRequestStatus: updated request status"
      );

      return res.status(200).json({
        message: "Successfully updated travel request status",
      });
    } catch (error) {
      console.error("RequestController.updateRequestStatus:", error);
      return res.status(500).json({ error: error.message });
    }
  }

  static async updateRequestRemarks(req, res) {
    try {
      await RequestService.updateRequestRemarks(
        req.params.id,
        req.body.remarks,
        req.user.id
      );
      console.log(
        "RequestController.updateRequestRemarks: updated request remarks"
      );

      return res.status(200).json({
        message: "Successfully updated travel request remarks",
      });
    } catch (error) {
      console.error("RequestController.updateRequestRemarks:", error);
      return res.status(500).json({ error: error.message });
    }
  }
}

module.exports = RequestController;
