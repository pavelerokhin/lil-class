package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/pavelerokhin/lil-class-selector-tool/utl"
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
	bet := utl.OnlyWithText(utl.BaseElementsFormDocument(d))
	sort.Sort(utl.ByTextLength(bet))

	// get 80% of elements // TODO: REFACTOR THIS: IT'S NOT COMFORTABLE TO SET minFreq, RATHER SET IT LIKE QUANTILE
	betHead := bet.GetHeadByFreq(0.2)

	csHead := utl.NewClassStats(betHead)
	csHead.PrintInOrderFreq()

	// common parent
	commonParent := csHead.GetCommonParentWithin(3)
	if commonParent != nil {
		fmt.Printf("common parent found: %v\n", commonParent)
	} else {
		fmt.Println("common parent not found")
		// sufficient parents
		csHead.GetSufficientParents()
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
