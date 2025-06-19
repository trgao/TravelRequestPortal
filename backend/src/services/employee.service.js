const RequestsDatabase = require("../database/requests.database");

class EmployeeService {
  static async getRequests(id, limit, page) {
    const requests = await RequestsDatabase.getRequestsForEmployee(
      id,
      limit,
      page
    );
    const count = await RequestsDatabase.getRequestsCountForEmployee(id);

    console.log(count);
    const response = {
      requests: requests,
    };
    if (limit == 0) {
      response.total_pages = 0;
    } else if (limit != null) {
      response.total_pages = Math.max(1, Math.ceil(count / limit));
    }

    return response;
  }
}

module.exports = EmployeeService;
