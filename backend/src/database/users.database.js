const db = require("./database");

class UsersDatabase {
  static async checkLastLogin(userId) {
    const result = await db.query(
      `SELECT last_login_at FROM users WHERE id = $1`,
      [userId]
    );

    if (result.rows.length == 0) {
      throw new Error("User not found");
    }

    return result.rows[0].last_login_at;
  }

  static async findUserByEmail(email) {
    const result = await db.query(`SELECT * FROM users WHERE email = $1`, [
      email,
    ]);

    if (result.rows.length == 0) {
      throw new Error("User not found");
    }

    return result.rows[0];
  }

  static async findUserById(id) {
    const result = await db.query(`SELECT * FROM users WHERE id = $1`, [id]);

    if (result.rows.length == 0) {
      throw new Error("User not found");
    }

    return result.rows[0];
  }

  static async createAdmin(user) {
    const findResult = await db.query(`SELECT * FROM users WHERE email = $1`, [
      user.email,
    ]);

    if (findResult.rows.length > 0) {
      throw new Error("User already exists");
    }

    const client = await db.client();

    try {
      await client.query("BEGIN");
      const findResult = await db.query(
        `SELECT * FROM companies WHERE name = $1`,
        [user.company]
      );

      if (findResult.rows.length > 0) {
        throw new Error("Company already exists");
      }

      const companyResult = await db.query(
        `INSERT INTO companies VALUES(DEFAULT, $1) RETURNING *`,
        [user.company]
      );

      if (companyResult.rows.length == 0) {
        throw new Error("Unable to create admin account");
      }

      const company = companyResult.rows[0];
      const userResult = await db.query(
        `INSERT INTO users VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7) RETURNING *`,
        [
          user.email,
          user.first_name,
          user.last_name,
          user.password_hash,
          user.is_admin,
          new Date(),
          company.id,
        ]
      );

      if (userResult.rows.length == 0) {
        throw new Error("Unable to create admin account");
      }

      await client.query("COMMIT");
      return userResult.rows[0];
    } catch (e) {
      await client.query("ROLLBACK");
      throw e;
    } finally {
      client.release();
    }
  }

  static async createEmployee(user, adminId) {
    const findResult = await db.query(`SELECT * FROM users WHERE email = $1`, [
      user.email,
    ]);

    if (findResult.rows.length > 0) {
      throw new Error("User already exists");
    }

    const adminResult = await db.query(`SELECT * FROM users WHERE id = $1`, [
      adminId,
    ]);

    if (adminResult.rows.length == 0) {
      throw new Error("Admin does not exist");
    }

    const admin = adminResult.rows[0];

    const userResult = await db.query(
      `INSERT INTO users VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7) RETURNING *`,
      [
        user.email,
        user.first_name,
        user.last_name,
        user.password_hash,
        user.is_admin,
        new Date(0),
        admin.company_id,
      ]
    );

    if (userResult.rows.length == 0) {
      throw new Error("Unable to create employee account");
    }

    return userResult.rows[0];
  }

  static async updateLastLoginAt(userId) {
    const result = await db.query(
      `UPDATE users SET last_login_at = $1
      WHERE id = $2 RETURNING *`,
      [new Date(), userId]
    );

    if (result.rows.length == 0) {
      throw new Error("Unable to update last login at");
    }

    return result.rows[0];
  }
}

module.exports = UsersDatabase;
