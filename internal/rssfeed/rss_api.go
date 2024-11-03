package rssfeed

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

var newFeedClient = NewClient(time.Second * 10)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create request%w", err)
	}

	request.Header.Add("User-agent", "gator")
	response, err := newFeedClient.httpClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("failed to get request %w", err)
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read response body%w", err)
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)

	if err != nil {
		return nil, fmt.Errorf("failed to parse xml %w", err)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}

	return &rssFeed, nil
}
