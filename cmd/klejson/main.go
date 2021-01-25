package main

import (
	"fmt"
	"log"
	"os"

	"github.com/koron-go/kcalign/internal/klejson"
)

func main() {
	l, err := klejson.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	keys := l.Keys().SortByCenter()
	const unit = 19.05
	for i, k := range keys {
		x , y:= k.CX*unit, k.CY*unit
		fmt.Printf("SW%d\t%f\t%f\n", i+1, x, y)
		fmt.Printf("D%d\t%f\t%f\n", i+1, x, y + 8.33)
		fmt.Printf("LED%d\t%f\t%f\n", i+1, x, y - 4.76)
	}
}
