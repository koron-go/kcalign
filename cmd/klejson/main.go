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
	//e := json.NewEncoder(os.Stdout)
	//e.SetIndent("", "  ")
	//e.Encode(keys)
	const unit = 19.05
	for i, k := range keys {
		fmt.Printf("SW%d\t%f\t%f\n", i+1, k.CX*unit, k.CY*unit)
	}
}
