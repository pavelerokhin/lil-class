package utl

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type ByTextLength Selections

func (a ByTextLength) Len() int           { return len(a) }
func (a ByTextLength) Less(i, j int) bool { return len(a[i].Text()) < len(a[j].Text()) }
func (a ByTextLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func BaseElementsFormDocument(d *goquery.Document) Selections {
	return Elements(d, SELECTOR_BASE)
}

func GetClass(e *goquery.Selection) string {
	class, exists := e.Attr("class")
	if !exists {
		class = "-"
	}

	return class
}

func GetType(e *goquery.Selection) string {
	eType := "unknown"
	if len(e.Nodes) > 0 && e.Nodes[0].Type == html.ElementNode {
		eType = e.Nodes[0].Data
	}

	return eType
}

func IsParentSufficient(element *goquery.Selection) bool {
	var score int
	for _, check := range RegexCheck {
		if !Has(element, check.Selector) {
			continue
		}
		if check.Checker != nil && check.Checker(element) {
			score++
			continue
		}
		score++
	}

	return score > MANDATORY_SCORE
}

func Elements(d *goquery.Document, selector string) Selections {
	var sub Selections
	s := d.Find(selector)
	s.Each(func(i int, e *goquery.Selection) {
		sub = append(sub, e)
	})

	return sub
}

func Has(element *goquery.Selection, selector string) bool {
	selection := element.ChildrenFiltered(selector)
	return selection != nil
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func OnlyWithText(s Selections) Selections {
	var sub Selections
	for _, e := range s {
		if len(e.Text()) > 0 {
			sub = append(sub, e)
		}
	}
	return sub
}
