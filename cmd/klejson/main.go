package main

import (
	"fmt"
	"log"
	"os"

	"github.com/koron-go/kcalign/internal/klejson"
)

var (
	enableDiode = true
	enableLED   = true
)

func main() {
	l, err := klejson.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	keys := l.Keys().SortByCenter()
	const unit = 19.05
	for i, k := range keys {
		x, y := k.CX*unit, k.CY*unit
		fmt.Printf("SW%d\t%f\t%f\n", i+1, x, y)
	}
	for i, k := range keys {
		x, y := k.CX*unit, k.CY*unit
		fmt.Printf("D%d\t%f\t%f\n", i+1, x, y+8.33)
	}

	nled := 0
	for nrow, row := range l.Rows {
		for ncol, key := range row {
			i := nled + ncol
			if nrow%2 == 1 {
				i = nled + len(row) - ncol - 1
			}
			x, y := key.CX*unit, key.CY*unit
			fmt.Printf("LED%d\t%f\t%f\n", i+1, x, y-4.76)
		}
		nled += len(row)
	}
}
