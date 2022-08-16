package utl

import (
	"github.com/PuerkitoBio/goquery"
)

type ByTextLength []*goquery.Selection

func (a ByTextLength) Len() int           { return len(a) }
func (a ByTextLength) Less(i, j int) bool { return len(a[i].Text()) < len(a[j].Text()) }
func (a ByTextLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Elements(d *goquery.Document, selector string) Selections {
	var sub Selections
	s := d.Find(selector)
	s.Each(func(i int, e *goquery.Selection) {
		sub = append(sub, e)
	})
	return sub
}

func SelectBaseElements(d *goquery.Document) Selections {
	BASE_SELECTOR := "div a, div p"
	return Elements(d, BASE_SELECTOR)
}

func SelectOnlyWithText(s Selections) Selections {
	var sub Selections
	for _, e := range s {
		if len(e.Text()) > 0 {
			sub = append(sub, e)
		}
	}

	return sub
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
