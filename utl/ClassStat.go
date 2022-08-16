package utl

import (
	"fmt"
	"golang.org/x/net/html"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ClassStats map[string]*Stats

// NewClassStats is a constructor method
func NewClassStats(s Selections) ClassStats {
	cls := make(ClassStats)
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
func (cs ClassStats) GetSufficientParents() Selections {
	var parents, suff Selections
	for _, s := range cs {
		parent := s.GetSufficientParent()
		if parent != nil {
			suff = append(suff, parent)
		}
		parents = append(parents, parent)
	}

	// base cases for the recursion
	if len(parents) == 0 {
		return nil
	}
	if len(suff) == 0 {
		cs2 := NewClassStats(parents)
		return cs2.GetSufficientParents()
	}

	return suff
}

type Stats struct {
	Freq      float64
	Selection *goquery.Selection
	Type      string
	Total     int
}

func (s *Stats) GetParent() *goquery.Selection {
	return s.Selection.Parent()
}

func (s *Stats) GetParents() *goquery.Selection {
	return s.Selection.Parents()
}

func (s *Stats) GetSufficientParent() *goquery.Selection {
	var sufficientParent *goquery.Selection
	parent := s.GetParent()
	// check if parent is  sufficient
	if parent != nil && IsParentSufficient(parent) {
		sufficientParent = parent
	}
	return sufficientParent
}

func (s *Stats) String() string {
	return fmt.Sprintf("type: %s freq: %.0f total: %d", s.Type, s.Freq, s.Total)
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
