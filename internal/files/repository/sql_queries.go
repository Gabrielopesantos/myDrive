package repository

const (
	createFileInsertion = `INSERT INTO files(file_id, file_owner_id, file, filename, extension, size, description, tags, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now(), now()) 
						RETURNING file_id, file_owner_id, file, filename, extension, size, description, created_at, updated_at`

	getFileByIdQuery = `SELECT file_id, file_owner_id, file, filename, extension, size, description, created_at, updated_at
						FROM files
						WHERE file_id = $1`
)
