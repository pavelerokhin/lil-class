package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/pavelerokhin/lil-class-selector-tool/utl"
	"golang.org/x/net/html"
)

func main() {
	// recover on panic
	defer func() {
		if err := recover(); err != nil {
			check(err)
		}
	}()

	url := os.Args[1]

	s := utl.NewScrapper()
	d, err := s.GetDocument(url)
	check(err)
	bet := utl.SelectOnlyWithText(utl.SelectBaseElements(d))
	sort.Sort(utl.ByTextLength(bet))

	// get 80% of elements // TODO: REFACTOR THIS: IT'S NOT COMFORTABLE TO SET minFreq, RATHER SET IT LIKE QUANTILE
	betHead := bet.GetHeadByFreq(0.2)

	csHead := utl.GetClassStats(betHead)
	csHead.PrintInOrderFreq()
	commonParent := csHead.GetCommonParentWithin(100)
	if commonParent != nil {
		fmt.Printf("common parent found: %v\n", commonParent)
	} else {
		fmt.Println("common parent not found")
	}
}

func check(err any) {
	if err != nil {
		fmt.Println(err)
		help()
		os.Exit(1)
	}
}

func help() {
	fmt.Println("*****\nUSAGE:\nlilclass <URL>")
}

func printElements(ee utl.Selections) {
	for i, e := range ee {
		// get class name
		class, exists := e.Attr("class")
		if !exists {
			class = "-"
		}
		// get element type
		eType := "unknown"
		if len(e.Nodes) > 0 && e.Nodes[0].Type == html.ElementNode {
			eType = e.Nodes[0].Data
		}
		// text
		text := e.Text()

		fmt.Printf("%d * TYPE: %s * CLASS: %s * LENGTH: %d * TEXT: %v\n", i, eType, class, len(text), text)
	}
}
