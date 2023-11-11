package db

var (
	// Postgres queries
	PsqlInsertUser = `INSERT INTO users(userid, firstname, lastname, username, email, password, phone) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING userid;`
	PsqlFetchUserById = `SELECT userid, firstname, lastname, username, email, phone, created_at, updated_at FROM users WHERE userid = $1;`
	PsqlFetchUserByEmail = `SELECT userid, firstname, lastname, username, email, phone, created_at, updated_at FROM users WHERE email = $1;`
	PsqlFetchUserByUsername = `SELECT userid, firstname, lastname, username, email, phone, created_at, updated_at FROM users WHERE username = $1;`
	PsqlFetchUsers = `SELECT userid, firstname, lastname, username, email, phone, created_at, updated_at FROM users;`
	PsqlUpdateUser = `UPDATE users SET firstname = $2, lastname = $3, username = $4, phone = $5 WHERE userid = $1;`
	PsqlDeleteUser = `DELETE FROM users WHERE userid = $1;`
	// Redis queries
)