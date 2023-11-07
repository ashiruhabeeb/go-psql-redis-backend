package db

var (
	// Postgres queries
	PsqlInsertUser = `INSERT INTO users(firstname, lastname, username, email, password, phone, house_number, street_name, local_area, state, country) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING userid`
	PsqlFetchUserById = `SELECT firstname, lastname, username, email,  FROM users WHERE userid = ?`

	// Redis queries
)