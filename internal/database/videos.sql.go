// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: videos.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const addVideo = `-- name: AddVideo :exec
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
`

type AddVideoParams struct {
	VideoID      string
	Title        string
	Description  sql.NullString
	PublishedAt  time.Time
	ChannelID    string
	ChannelTitle string
	Thumbnails   json.RawMessage
}

func (q *Queries) AddVideo(ctx context.Context, arg AddVideoParams) error {
	_, err := q.db.ExecContext(ctx, addVideo,
		arg.VideoID,
		arg.Title,
		arg.Description,
		arg.PublishedAt,
		arg.ChannelID,
		arg.ChannelTitle,
		arg.Thumbnails,
	)
	return err
}

const getVideos = `-- name: GetVideos :many
SELECT
	video_id, title, description, published_at, channel_id, channel_title, thumbnails
FROM videos ORDER BY published_at DESC OFFSET $1 LIMIT $2
`

type GetVideosParams struct {
	Offset int32
	Limit  int32
}

type GetVideosRow struct {
	VideoID      string
	Title        string
	Description  sql.NullString
	PublishedAt  time.Time
	ChannelID    string
	ChannelTitle string
	Thumbnails   json.RawMessage
}

func (q *Queries) GetVideos(ctx context.Context, arg GetVideosParams) ([]GetVideosRow, error) {
	rows, err := q.db.QueryContext(ctx, getVideos, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVideosRow
	for rows.Next() {
		var i GetVideosRow
		if err := rows.Scan(
			&i.VideoID,
			&i.Title,
			&i.Description,
			&i.PublishedAt,
			&i.ChannelID,
			&i.ChannelTitle,
			&i.Thumbnails,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
