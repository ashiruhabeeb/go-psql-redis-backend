package db

var (
	// Postgres queries
	PsqlInsertUser = `INSERT INTO users(firstname, lastname, username, email, password, phone, house_number, street_name, local_area, state, country) VALUES ($1, $2, $3, $4, 5$, $6, $7, $8, $8, $9, $10, $11) RETURNING userid`
	PsqlFetchUserById = `SELECT firstname, lastname, username, email,  FROM users WHERE userid = $1`
	PsqlFetchBtEmail = `SELECT * FROM users WHERE email = $1`
	PsqlFetchUsers = `SELECT * FROM users`

	// Redis queries
)