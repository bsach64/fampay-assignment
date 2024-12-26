-- +goose Up
CREATE TABLE videos (
	id SERIAL PRIMARY KEY,
	video_id TEXT NOT NULL UNIQUE,
	title TEXT NOT NULL UNIQUE,
	description TEXT,
	published_at TIMESTAMP NOT NULL,
	channel_id TEXT NOT NULL,
	channel_title TEXT NOT NULL,
	thumbnails JSONB NOT NULL,
	fetched_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_videos_published_at_desc ON videos (published_at DESC);

-- +goose Down
DROP TABLE videos;
