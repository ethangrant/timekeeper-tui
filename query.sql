-- name: GetAllTasks :many
SELECT * FROM tasks;

-- name: GetAllTasksByDate :many
SELECT * FROM tasks where date(created_at) = date(?);

-- name: CreateTask :one
INSERT INTO tasks (title, desc, duration) VALUES (?, ?, ?) RETURNING *;

-- name: UpdateTaskDuration :one
UPDATE tasks SET duration = ? where id = ? RETURNING *;