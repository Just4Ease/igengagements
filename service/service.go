package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"ig.engagements/log"
	"ig.engagements/models"
	"strings"
)

var (
	ErrReadingIGData      = errors.New("error reading provided account's data")
	ErrFailedFetchProfile = errors.New("failed fetch instagram profile")
)

type IgEngagementService struct {
	c      *colly.Collector
	logger log.Logger
}

func NewIGEngagementService(logger log.Logger) *IgEngagementService {
	return &IgEngagementService{c: colly.NewCollector(), logger: logger}
}

func (ig IgEngagementService) GetEngagementRating(username string) error {
	closeSignal := make(chan bool)
	errChan := make(chan error)
	dataset := ""

	link := fmt.Sprintf("https://instagram.com/%s", username)

	ig.c.OnHTML("script", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "edge_owner_to_timeline_media") {
			dataset = e.Text
			closeSignal <- true
		}
	})

	go func(c *colly.Collector, closeSignal chan bool, errChan chan error) {
		// This is so that we can exit the crawling process after we retrieved the 1st page.
		go func() {
			if err := c.Visit(link); err != nil {
				errChan <- err
				return
			}
		}()
		<-closeSignal
		errChan <- nil // Send no errors.
	}(ig.c, closeSignal, errChan)

	if err := <-errChan; err != nil {
		ig.logger.WithError(err).WithField("link", link).Error("failed to crawl/fetch data instagram composed profile url")
		return ErrFailedFetchProfile
	}

	var dmp models.WebDumpedJson
	clean := strings.ReplaceAll(dataset, "window._sharedData = ", "")
	clean = strings.ReplaceAll(clean, ";", "")
	if err := json.Unmarshal([]byte(strings.TrimSpace(clean)), &dmp); err != nil {
		ig.logger.WithError(err).Error("failed to parse IG scraped json web dump")
		return ErrReadingIGData
	}

	followers := 0
	likes := 0
	comments := 0
	rates := 0.0

	for _, profile := range dmp.EntryData.ProfilePage {
		followers += profile.Graphql.User.EdgeFollowedBy.Count

		// loop comments & likes concurrently:
		for _, edge := range profile.Graphql.User.EdgeOwnerToTimelineMedia.Edges {
			likes += edge.Node.EdgeLikedBy.Count
			comments += edge.Node.EdgeMediaToComment.Count
			rate := (float64(edge.Node.EdgeLikedBy.Count) + float64(edge.Node.EdgeMediaToComment.Count)) / float64(followers)
			rates += rate * 100 // "rates" is the sum of all `rate` in %.
		}
	}

	// TODO: Define struct for this final result..., mount frontend, mount http Server
	fmt.Printf("Followers: %d\n", followers)
	fmt.Printf("Last 12 Post Likes: %d\n", likes)
	fmt.Printf("Last 12 Post Comments: %d\n", comments)
	fmt.Printf("Engagement Rate: %f\n", rates/12)

	return nil
}
