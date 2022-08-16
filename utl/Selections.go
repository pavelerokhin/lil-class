package utl

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Selections []*goquery.Selection

func (s Selections) GetHeadByFreq(minFreq float64) Selections {
	var out Selections
	classStats := NewClassStats(s)
	for _, stats := range classStats {
		if minFreq < stats.Freq/float64(stats.Total) {
			out = append(out, stats.Selection)
		}
	}

	return out
}

func (s Selections) GetTailByFreq(maxFreq float64) Selections {
	var out Selections
	classStats := NewClassStats(s)
	for _, stats := range classStats {
		if maxFreq > stats.Freq/float64(stats.Total) {
			out = append(out, stats.Selection)
		}
	}

	return out
}

func (s Selections) printElements() {
	for i, e := range s {
		text := e.Text()
		fmt.Printf("%d. type: %s class: %s length: %d * text: %v\n", i, GetType(e), GetClass(e), len(text), text)
	}
}
