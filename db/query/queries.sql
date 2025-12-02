-- name: CreateURL :one
INSERT INTO urls (short_code, long_url)
VALUES ($1, $2)
RETURNING *;

-- name: GetByShortCode :one
SELECT * FROM urls WHERE short_code = $1 LIMIT 1;

-- name: UpdateClicks :exec
UPDATE urls SET clicks = clicks + 1 WHERE short_code = $1;

-- name: UpdateLongUrl :exec
UPDATE urls SET long_url = $2 WHERE short_code = $1
RETURNING *;

-- name: DeleteByShortCode :exec
DELETE FROM urls WHERE short_code = $1 RETURNING id;
