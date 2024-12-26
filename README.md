# Fampay Backend Assignment
## Setup
After cloning the repo and install dependencies, you can then:
1. Create a `.env` file with
```
API_KEYS=<KEY_ONE>,<KEY_TWO>,<KEY_THREE>
DATABASE=<postgres_url>
```
2. `go build`

## Testing

The server starts at port `8080`.
You can test the server by sending it a `CURL` request
```
curl "http://localhost:8080/videos?page=1"
```
which will respond with
```
{
  "next_page": 2,
  "videos": [
    {
      "videoId": "xXhgxZcDDDM",
      "channelId": "UCx3mBoUxGZkSfh0CXNxx9fw",
      "title": "#77 India Women vs West Indies Women, 3rd Odi | Live Cricket Match Today | IND Women vs W Gameplay",
      "description": "sportstalks #indvswi #cricketlive #77 India Women vs West Indies Women, 3rd Odi Live 2nd innings | Live Cricket Match ...",
      "thumbnails": {
        "default": {
          "url": "https://i.ytimg.com/vi/xXhgxZcDDDM/default.jpg",
          "width": 120,
          "height": 90
        },
        "medium": {
          "url": "https://i.ytimg.com/vi/xXhgxZcDDDM/mqdefault.jpg",
          "width": 320,
          "height": 180
        },
        "high": {
          "url": "https://i.ytimg.com/vi/xXhgxZcDDDM/hqdefault.jpg",
          "width": 480,
          "height": 360
        }
      },
      "channelTitle": "Sports Talks",
      "publishedAt": "2024-12-26T17:55:34Z"
    },
    {
    	... 4 more entries
    }
  ]
}
```
you can change the value of page to see older values.
The server responds with the latest videos in reverse chronological order.

## Features and Technologies Used

1. `postgres`: Database for storing responses from YouTube API.
2. `goose`: Used for database migrations
3. `sqlc`: Generates go code from sql code

The server re-requests the youtube API with a different API key if quota has been reached on previous one.
It also has a background worker which fetches new videos from YouTube every 10 seconds.

