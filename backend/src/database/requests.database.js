const db = require("./database");

class RequestsDatabase {
  static async checkAdminPermission(requestId, adminId) {
    const result = await db.query(
      `SELECT * FROM requests
      INNER JOIN users on user_id = users.id
      WHERE requests.id = $1 AND
      company_id = (SELECT company_id FROM users WHERE id = $2)`,
      [requestId, adminId]
    );

    if (result.rows.length == 0) {
      throw new Error("You do not have permission to update request");
    }
  }

  static async getPending(userId) {
    const result = await db.query(
      `SELECT COUNT(*) FROM requests
      INNER JOIN users ON user_id = users.id
      WHERE status = 'pending' AND
      company_id = (SELECT company_id FROM users WHERE id = $1)`,
      [userId]
    );
    return result.rows[0].count;
  }

  static async getApproved(userId) {
    const result = await db.query(
      `SELECT COUNT(*) FROM requests
      INNER JOIN users ON user_id = users.id
      WHERE status = 'approved' AND
      company_id = (SELECT company_id FROM users WHERE id = $1)`,
      [userId]
    );
    return result.rows[0].count;
  }

  static async getRejected(userId) {
    const result = await db.query(
      `SELECT COUNT(*) FROM requests
      INNER JOIN users ON user_id = users.id
      WHERE status = 'rejected' AND
      company_id = (SELECT company_id FROM users WHERE id = $1)`,
      [userId]
    );
    return result.rows[0].count;
  }

  static async getRequestsCountForEmployee(id) {
    const request = await db.query(
      `SELECT count(*) FROM requests WHERE user_id = $1`,
      [id]
    );
    return request.rows[0].count;
  }

  static async getRequestsForEmployee(id, limit, page) {
    let queryString = `SELECT * FROM requests WHERE user_id = ${id} ORDER BY id`;
    if (limit != null) {
      queryString += ` LIMIT ${limit}`;
    }
    if (page != null) {
      queryString += ` OFFSET ${(page - 1) * limit}`;
    }

    const request = await db.query(queryString);

    return request.rows;
  }

  static async getRequestsCountForAdmin(id, query) {
    const request = await db.query(
      `SELECT COUNT(*) FROM requests
      INNER JOIN users ON user_id = users.id
      WHERE company_id = (SELECT company_id FROM users WHERE id = $1)
      ${
        query.employee != null
          ? `AND (first_name || ' ' || last_name ILIKE '%${query.employee}%' OR
        first_name || ' ' || last_name ILIKE '%${query.employee}%')`
          : ""
      }`,
      [id]
    );
    return request.rows[0].count;
  }

  static async getRequestsForAdmin(id, query) {
    let queryString = `SELECT requests.id, start_location, destination, date_time,
      purpose, status, admin_remarks, first_name, last_name FROM requests
      INNER JOIN users ON user_id = users.id
      WHERE company_id = (SELECT company_id FROM users WHERE id = ${id})`;

    if (query.employee != null) {
      queryString += ` AND (first_name || ' ' || last_name ILIKE '%${query.employee}%' OR
      first_name || ' ' || last_name ILIKE '%${query.employee}%')`;
    }
    queryString += ` ORDER BY requests.id`;
    if (query.limit != null) {
      queryString += ` LIMIT ${query.limit}`;
    }
    if (query.page != null) {
      queryString += ` OFFSET ${(query.page - 1) * query.limit}`;
    }

    const request = await db.query(queryString);

    return request.rows;
  }

  static async createRequest(request, userId) {
    const result = await db.query(
      `INSERT INTO requests VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7) RETURNING *`,
      [
        request.start_location,
        request.destination,
        request.date_time,
        request.purpose,
        "pending",
        "",
        userId,
      ]
    );

    if (result.rows.length == 0) {
      throw new Error("Unable to create request");
    }

    return result.rows[0];
  }

  static async updateRequestStatus(id, status) {
    const result = await db.query(
      `UPDATE requests SET status = $1 WHERE id = $2 RETURNING *`,
      [status, id]
    );

    if (result.rows.length == 0) {
      throw new Error("Unable to update request status");
    }

    return result.rows[0];
  }

  static async updateRequestRemarks(id, remarks) {
    const result = await db.query(
      `UPDATE requests SET admin_remarks = $1 WHERE id = $2 RETURNING *`,
      [remarks, id]
    );

    if (result.rows.length == 0) {
      throw new Error("Unable to update request remarks");
    }

    return result.rows[0];
  }
}

module.exports = RequestsDatabase;
