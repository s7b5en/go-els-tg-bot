package fetcher

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/teimurjan/go-els-tg-bot/errs"
	"github.com/teimurjan/go-els-tg-bot/texts"
)

const url = "https://els.kg/en/find_tracking"
const statusSelector = "span.br-pill"

type trackingStatusFetcher struct{}

// NewTrackingStatusFetcher creates a new instance of tracking status fetcher
func NewTrackingStatusFetcher() *trackingStatusFetcher {
	return &trackingStatusFetcher{}
}

// Fetch fetches oreder status by tracking
func (t *trackingStatusFetcher) Fetch(trackingNumber string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return "", err
	}

	q := req.URL.Query()
	q.Add("q", trackingNumber)
	req.URL.RawQuery = q.Encode()
	doc, err := goquery.NewDocument(req.URL.String())
	if err != nil {
		log.Print(err)
		return "", err
	}

	status := strings.TrimSpace(doc.Find(statusSelector).Text())
	if status == "" {
		return "", errs.NewErr(
			errs.NoSuchTrackingErrCode,
			texts.GetTrackingNotExistsMessage(trackingNumber),
		)
	}
	return status, nil
}
