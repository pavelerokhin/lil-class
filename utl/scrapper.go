package utl

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Scrapper struct{}

func NewScrapper() *Scrapper {
	return &Scrapper{}
}

func (s *Scrapper) GetDocument(url string) (*goquery.Document, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("error, status code: %v", response.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	return document, nil
}
