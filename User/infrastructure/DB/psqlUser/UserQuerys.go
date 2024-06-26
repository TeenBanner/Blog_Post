package psqlUser

const SqlCreateUserQuery = `INSERT INTO users(id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

const SqlGetUser = `SELECT id, name, email, created_at, updated_at FROM users WHERE email = $1`

const SqlUpdateUserName = `UPDATE users SET name = $1, updated_at = $2 WHERE email = $3`
const SqlUpdateUserEmail = `UPDATE users SET email = $1, updated_at =$2 WHERE email = $3`
const SqlUpdateUserPassword = `UPDATE users SET password = $1, updated_at = $2 WHERE email = $3`

const SqlAdminGetAllUsers = `SELECT id, name, email, created_at, updated_at FROM users`

const SqlGetUserByName = `SELECT id, name, email, created_at, updated_at FROM users WHERE name = $1`

const SqlLoginCompareEmails = `SELECT email FROM users WHERE email = $1`
const SqlLoginGetHashdPasswordWithEmail = `SELECT password FROM users WHERE email = $1`

const SqlGetUserName = `SELECT name FROM users WHERE email = $1`
const SqlFindUserEmailByName = `SELECT email FROM users WHERE name = $1`

const SqlDeleteAccount = `DELETE FROM users WHERE email = $1`
