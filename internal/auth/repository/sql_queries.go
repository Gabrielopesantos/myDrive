package repository

const (
	createUserQuery = `INSERT INTO users(first_name, last_name, email, password, role, about, avatar, last_login, created_at, updated_at)
						VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'user'), $6, $7,  now(), now(), now())
						RETURNING user_id, first_name, last_name, email, role, about, is_email_verified, last_login, created_at, updated_at`

	findByEmailQuery = `SELECT user_id, password, first_name, last_name, email, role, avatar, about, is_email_verified, created_at, updated_at
						FROM users
						WHERE email = $1`
)
