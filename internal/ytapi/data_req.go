package ytapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Gets yt videos from Youtube API in reverse chronological order i.e newest first
func (c *Client) PublishedVideosByDate(searchQuery string, PublishedAfter time.Time) (DataResponse, error) {
	baseURL := "https://www.googleapis.com/youtube/v3/search?part=snippet&order=date&type=video"

	validKeyIdx := -1
	for i, key := range c.apiKeys {
		if key.quotaReached == false {
			validKeyIdx = i
			break
		}
	}

	if validKeyIdx == -1 {
		return DataResponse{}, fmt.Errorf("all api keys have reached their quota")
	}

	fullURL := fmt.Sprintf("%v&q=%v&key=%v&publishedAfter=%v", baseURL, searchQuery, c.apiKeys[validKeyIdx].apiKey, PublishedAfter.Format(time.RFC3339))

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
		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return DataResponse{}, err
		}

		errMsg := YTError{}
		err = json.Unmarshal(dat, &errMsg)
		if err != nil {
			return DataResponse{}, err
		}

		if errMsg.Error.Code == 403 {
			for _, e := range errMsg.Error.Errors {
				if e.Reason == "quotaExceeded" {
					log.Printf("quota exceeded! trying different key")
				}
				c.apiKeys[validKeyIdx].quotaReached = true
				return c.PublishedVideosByDate(searchQuery, PublishedAfter)
			}
		}
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
