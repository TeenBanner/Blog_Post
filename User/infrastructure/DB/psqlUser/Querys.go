package psqlUser

const SqlCreateUserQuery = `INSERT INTO users(id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
const SqlGetUser = `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
const SqlGetUserPosts = `SELECT * FROM posts WHERE owner_id = $1`