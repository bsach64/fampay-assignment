-- name: AddVideo :exec
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
ON CONFLICT (video_id) DO NOTHING;

-- name: GetVideos :many
SELECT
	video_id, title, description, published_at, channel_id, channel_title, thumbnails
FROM videos ORDER BY published_at DESC OFFSET $1 LIMIT $2;
