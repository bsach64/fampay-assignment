-- name: AddVideo :one
INSERT INTO videos (
	video_id, title, description, published_at, channel_id, channel_title, thumbnails
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7
)
ON CONFLICT (video_id) DO NOTHING
RETURNING *;
