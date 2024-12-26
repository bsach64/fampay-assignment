package ytapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Gets yt videos from Youtube API in reverse chronological order i.e newest first
func (c *Client) PublishedVideosByDate(searchQuery string, PublishedAfter time.Time) (DataResponse, error) {
	baseURL := "https://www.googleapis.com/youtube/v3/search?part=snippet&order=date&type=video"
	fullURL := fmt.Sprintf("%v&q=%v&key=%v&publishedAfter=%v", baseURL, searchQuery, c.apiKeys[0], PublishedAfter.Format(time.RFC3339))

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return DataResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return DataResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return DataResponse{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return DataResponse{}, err
	}

	Data := DataResponse{}
	err = json.Unmarshal(dat, &Data)
	if err != nil {
		return DataResponse{}, err
	}

	return Data, err
}
