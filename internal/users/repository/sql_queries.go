package repository

const (
	createUserQuery = `INSERT INTO users(first_name, last_name, email, password, role, about, avatar, created_at, updated_at)
					VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'user'), $6, $7,  now(), now()) 
					RETURNING user_id, first_name, last_name, email, role, about, is_email_verified, created_at, updated_at`
	//RETURNING *`

	getUserQuery = `SELECT user_id, first_name, last_name, email, role, avatar, about, is_email_verified, created_at, updated_at
					FROM users 
					WHERE user_id = $1`
)
