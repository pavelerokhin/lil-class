package utl

import "github.com/PuerkitoBio/goquery"

// TODO: GET THOSE CONSTS FROM YAML
const (
	MANDATORY_SCORE = 1

	SELECTOR_BASE     = "div a, div p"
	SELECTOR_SPAN_P   = "div span, div p"
	SELECTOR_SUBTITLE = "h2 span:nth-child(2)"
	SELECTOR_TITLE    = "h1, h2 span:first-child"
)

type Check struct {
	Checker     func(*goquery.Selection) bool
	IsMandatory bool
	IsMultiple  bool
	Selector    string
}

var RegexCheck = map[string]Check{
	"base": {
		IsMandatory: true,
		Selector:    SELECTOR_BASE,
	},
	"date": {
		Checker:  IsDate,
		Selector: SELECTOR_SPAN_P,
	},
	"subtitle": {
		Checker:  IsSubtitle,
		Selector: SELECTOR_SUBTITLE,
	},
	"tag": {
		Checker:    IsTag,
		IsMultiple: true,
		Selector:   SELECTOR_SPAN_P,
	},
	"title": {
		Checker:     IsTitle,
		IsMandatory: true,
		Selector:    SELECTOR_TITLE,
	},
}

func IsDate(element *goquery.Selection) bool {
	return true
}

func IsSubtitle(element *goquery.Selection) bool {
	return true
}

func IsTag(element *goquery.Selection) bool {
	return true
}

func IsTitle(element *goquery.Selection) bool {
	return true
}
