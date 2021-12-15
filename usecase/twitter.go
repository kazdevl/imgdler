package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kazdevl/imgdler/entity"
	"github.com/sivchari/gotwtr"
)

type TwitterUsecase struct {
	c *gotwtr.Client
}

func NewTwitterUsecase(token string) *TwitterUsecase {
	return &TwitterUsecase{
		c: gotwtr.New(token),
	}
}

func (t *TwitterUsecase) FetchContent(author, keyword string, max int) ([]entity.Pages, error) {
	query := fmt.Sprintf("from:%s -is:retweet \"%s\"", author, keyword)
	res, err := t.c.SearchRecentTweets(context.Background(), query, &gotwtr.SearchTweetsOption{
		Expansions:  []gotwtr.Expansion{gotwtr.ExpansionAttachmentsMediaKeys},
		MediaFields: []gotwtr.MediaField{gotwtr.MediaFieldMediaKey, gotwtr.MediaFieldURL},
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldAttachments, gotwtr.TweetFieldCreatedAt},
		MaxResults:  max,
		StartTime:   time.Now().In(time.UTC).AddDate(0, 0, -1),
	})
	if err != nil {
		return []entity.Pages{}, err
	}
	pagesList := make([]entity.Pages, len(res.Tweets))
	for index, tweet := range res.Tweets {
		pagesList[index].Datetime, _ = time.Parse(time.RFC3339, tweet.CreatedAt)
		pagesList[index].Links = t.getTweetImageLinks(tweet, res.Includes.Media)
	}

	return pagesList, nil
}

func (t *TwitterUsecase) getTweetImageLinks(tweet *gotwtr.Tweet, media []*gotwtr.Media) []string {
	links := make([]string, 0, 4)
	for _, key := range tweet.Attachments.MediaKeys {
		for _, m := range media {
			if key == m.MediaKey {
				links = append(links, m.URL)
			}
		}
	}
	return links
}
