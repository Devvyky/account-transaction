-- name: GetOperationType :one
SELECT * FROM operation_types
WHERE operation_type_id = $1 LIMIT 1;

