const { compare, hash } = require("bcrypt");
const UsersDatabase = require("../database/users.database");
const { sign } = require("jsonwebtoken");

class AuthService {
  static generateToken(account) {
    const user = {
      id: account.id,
      email: account.email,
      is_admin: account.is_admin,
      iat: Math.ceil((Date.now() + 1) / 1000),
    };
    return sign(user, process.env.JWT_SECRET, { expiresIn: "1d" });
  }

  static async login(user) {
    const account = await UsersDatabase.findUserByEmail(user.email);
    const matched = await compare(user.password, account.password_hash);

    if (!matched) {
      throw new Error("Password is not correct");
    }

    await UsersDatabase.updateLastLoginAt(account.id);
    const token = this.generateToken(account);

    return {
      token: token,
      isAdmin: account.is_admin,
      userId: account.id,
    };
  }

  static async register(user) {
    const passwordHash = await hash(user.password, 10);
    const userWithPasswordHash = {
      email: user.email,
      first_name: user.first_name,
      last_name: user.last_name,
      password_hash: passwordHash,
      is_admin: true,
      company: user.company,
    };
    const account = await UsersDatabase.createAdmin(userWithPasswordHash);

    return {
      token: await this.generateToken(account),
      isAdmin: account.is_admin,
      userId: account.id,
    };
  }

  static async onboard(user, adminId) {
    const account = await UsersDatabase.findUserById(adminId);
    const passwordHash = await hash(user.password, 10);
    const userWithPasswordHash = {
      email: user.email,
      first_name: user.first_name,
      last_name: user.last_name,
      password_hash: passwordHash,
      is_admin: user.is_admin,
      company_id: account.company_id,
    };
    await UsersDatabase.createEmployee(userWithPasswordHash, adminId);
  }
}

module.exports = AuthService;
