package repository

const (
	updateUserQuery = `UPDATE users
					   SET first_name = COALESCE(NULLIF($1, ''), first_name),
						   last_name = COALESCE(NULLIF($2, ''), last_name),
						   email = COALESCE(NULLIF($3, ''), email),
						   role = COALESCE(NULLIF($4, ''), role),
						   about = COALESCE(NULLIF($5, ''), about),
						   is_email_verified = COALESCE(NULLIF($6, FALSE), is_email_verified),
						   avatar = COALESCE(NULLIF($7, ''), avatar),
						   updated_at = now()
   					   WHERE user_id = $8
					   RETURNING *`

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
