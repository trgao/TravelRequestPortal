const RequestsDatabase = require("../database/requests.database");

class RequestService {
  static async getSummary(userId) {
    const pending = await RequestsDatabase.getPending(userId);
    const approved = await RequestsDatabase.getApproved(userId);
    const rejected = await RequestsDatabase.getRejected(userId);

    return {
      pending: pending,
      approved: approved,
      rejected: rejected,
    };
  }

  static async createRequest(request, userId) {
    const newRequest = await RequestsDatabase.createRequest(request, userId);
    return newRequest;
  }

  static async updateRequestStatus(requestId, status, adminId) {
    await RequestsDatabase.checkAdminPermission(requestId, adminId);

    const updatedRequest = await RequestsDatabase.updateRequestStatus(
      requestId,
      status
    );
    return updatedRequest;
  }

  static async updateRequestRemarks(requestId, remarks, adminId) {
    await RequestsDatabase.checkAdminPermission(requestId, adminId);

    const updatedRequest = await RequestsDatabase.updateRequestRemarks(
      requestId,
      remarks
    );
    return updatedRequest;
  }
}

module.exports = RequestService;
