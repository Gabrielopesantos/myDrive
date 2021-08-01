package repository

const (
	updateUserLastLoginQuery = `UPDATE users
								SET last_login = now()
								WHERE email = $1
								RETURNING *`

	getUserQuery = `SELECT user_id, first_name, last_name, email, role, avatar, about, is_email_verified, created_at, updated_at
					FROM users 
					WHERE user_id = $1`

	getNumUsersQuery = `SELECT count(user_id) FROM users`

	getAllUsersQuery = `SELECT user_id, first_name, last_name, email, role, about, avatar, created_at, updated_at
						FROM users
						ORDER BY COALESCE(NULLIF($1, ''), first_name)`
)
