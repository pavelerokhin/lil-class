package utl

import (
	"fmt"
	"sort"

	"github.com/PuerkitoBio/goquery"
)

type ClassStats map[string]*Stats

func (cs ClassStats) GetCommonParentWithin(deltaBetweenElements int) *goquery.Selection {
	var commonParent *goquery.Selection
	var parents Selections
	for _, s := range cs {
		parents = append(parents, s.GetParents())
	}

	commonParent = parents[0]
	maxLength := commonParent.Length()
	for _, p := range parents {
		maxLength = MaxInt(maxLength, p.Length())
		commonParent = commonParent.FilterSelection(p)
	}
	if maxLength-commonParent.Length() <= deltaBetweenElements {
		return commonParent
	}
	return nil
}

type Stats struct {
	Freq      float64
	Selection *goquery.Selection
	Type      string
	Total     int
}

func (s *Stats) GetParents() *goquery.Selection {
	return s.Selection.Parents()
}

func (s *Stats) String() string {
	return fmt.Sprintf("type: %s freq: %.0f total: %d", s.Type, s.Freq, s.Total)
}

func NewClassStats() ClassStats {
	return make(ClassStats)
}

func (cs ClassStats) Add(c string) {
	cs[c] = &Stats{}
}

func (cs ClassStats) Has(c string) bool {
	_, exists := cs[c]
	return exists
}

func (cs ClassStats) PrintInOrderAlpha() {
	keys := make([]string, 0, len(cs))
	for k := range cs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("name: %s %s\n", k, cs[k].String())
	}
}

func (cs ClassStats) PrintInOrderFreq() {
	keys := make([]string, 0, len(cs))

	for key := range cs {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return cs[keys[i]].Freq < cs[keys[j]].Freq
	})

	for _, k := range keys {
		fmt.Printf("name: %s %s\n", k, cs[k].String())
	}
}

func (cs ClassStats) SetTotal(total int) {
	for _, stats := range cs {
		stats.Total = total
	}
}

func (cs ClassStats) Total() int {
	for _, v := range cs {
		return v.Total
	}
	return 0
}