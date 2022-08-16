package utl

import (
	"golang.org/x/net/html"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SubSelection []*goquery.Selection

func GetClassStats(s SubSelection) ClassStats {
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

func (s SubSelection) GetHeadByFreq(minFreq float64) SubSelection {
	var out SubSelection
	classStats := GetClassStats(s)
	for _, stats := range classStats {
		if minFreq < stats.Freq/float64(stats.Total) {
			out = append(out, stats.Selection)
		}
	}

	return out
}

func (s SubSelection) GetTailByFreq(maxFreq float64) SubSelection {
	var out SubSelection
	classStats := GetClassStats(s)
	for _, stats := range classStats {
		if maxFreq > stats.Freq/float64(stats.Total) {
			out = append(out, stats.Selection)
		}
	}

	return out
}
