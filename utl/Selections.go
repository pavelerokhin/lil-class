package utl

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type Selections []*goquery.Selection

func GetClassStats(s Selections) ClassStats {
	cls := NewClassStats()
	var total = 0
	for _, element := range s {
		classes, exists := element.Attr("class")
		if !exists {
			continue
		}

		for _, c := range strings.Split(classes, " ") {
			if cls.Has(c) {
				cls[c].Freq++
			} else {
				cls.Add(c)
				cls[c].Selection = element
				cls[c].Freq = 1
				if len(element.Nodes) > 0 && element.Nodes[0].Type == html.ElementNode {
					cls[c].Type = element.Nodes[0].Data
				}
			}
			total++
		}
	}
	cls.SetTotal(total)

	return cls
}

func (s Selections) GetHeadByFreq(minFreq float64) Selections {
	var out Selections
	classStats := GetClassStats(s)
	for _, stats := range classStats {
		if minFreq < stats.Freq/float64(stats.Total) {
			out = append(out, stats.Selection)
		}
	}

	return out
}

func (s Selections) GetTailByFreq(maxFreq float64) Selections {
	var out Selections
	classStats := GetClassStats(s)
	for _, stats := range classStats {
		if maxFreq > stats.Freq/float64(stats.Total) {
			out = append(out, stats.Selection)
		}
	}

	return out
}
