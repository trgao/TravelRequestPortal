const RequestsDatabase = require("../database/requests.database");

class AdminService {
  static async getRequests(id, query) {
    const requests = await RequestsDatabase.getRequestsForAdmin(id, query);
    const count = await RequestsDatabase.getRequestsCountForAdmin(id, query);

    const response = {
      requests: requests,
    };
    if (query.limit == 0) {
      response.total_pages = 0;
    } else if (query.limit != null) {
      response.total_pages = Math.max(1, Math.ceil(count / query.limit));
    }

    return response;
  }
}

module.exports = AdminService;
