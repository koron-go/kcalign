package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/koron-go/kcalign/internal/klejson"
)

var (
	enableDiode = true
	enableLED   = true
)

type dim struct {
	x float64
	y float64
}

func (d dim) String() string {
	return fmt.Sprintf("%.2f,%.2f", d.x, d.y)
}

func (d *dim) Set(s string) error {
	if !strings.ContainsRune(s, ',') {
		var v float64
		_, err := fmt.Sscanf(s, "%f", &v)
		if err != nil {
			return err
		}
		d.x, d.y = v, v
		return nil
	}
	var x, y float64
	_, err := fmt.Sscanf(s, "%f,%f", &x, &y)
	if err != nil {
		return err
	}
	d.x, d.y = x, y
	return nil
}

func (d dim) Get() interface{} {
	return d
}

func main() {
	var (
		origin = dim{}
		unit   = dim{x: 19.05, y: 19.05}

		enableDiode = false
		offDiode    = dim{y: +8.33}

		enableLED = false
		offLED    = dim{y: -4.76}

		sortBy string
	)
	flag.Var(&origin, "origin", "the origin coordinate")
	flag.Var(&unit, "unit", "unit dimension in millimeter")
	flag.StringVar(&sortBy, "sort", "col,row", "sort priority \"col,row\" or \"row,col\"")
	flag.BoolVar(&enableDiode, "diode", false, "output diodes")
	flag.Var(&offDiode, "diode_offset", "diode offset")
	flag.BoolVar(&enableLED, "led", false, "output LEDs")
	flag.Var(&offLED, "led_offset", "LED offset")
	flag.Parse()

	var w io.Writer = os.Stdout
	l, err := klejson.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	keys := l.Keys()
	switch sortBy {
	case "col,row":
		keys = keys.SortByColRow()
	case "row,col":
		keys = keys.SortByRowCol()
	}

	// align switches
	for i, k := range keys {
		x := k.CX*unit.x + origin.x
		y := k.CY*unit.y + origin.y
		fmt.Fprintf(w, "SW%d\t%f\t%f\n", i+1, x, y)
	}

	// align diodes if required
	if enableDiode {
		for i, k := range keys {
			x := k.CX*unit.x + offDiode.x + origin.x
			y := k.CY*unit.y + offDiode.y + origin.y
			fmt.Fprintf(w, "D%d\t%f\t%f\n", i+1, x, y)
		}
	}

	// align LEDs if required
	if enableLED {
		nled := 0
		for nrow, row := range l.Rows {
			for ncol, key := range row {
				i := nled + ncol
				if nrow%2 == 1 {
					i = nled + len(row) - ncol - 1
				}
				x := key.CX*unit.x + offLED.x + origin.x
				y := key.CY*unit.y + offLED.y + origin.y
				fmt.Fprintf(w, "LED%d\t%f\t%f\n", i+1, x, y)
			}
			nled += len(row)
		}
	}
}
