package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/koron-go/kcalign"
)

func main() {
	err := fmtjson2go(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}

func fmtjson2go(w io.Writer, r io.Reader) error {
	var f *kcalign.Formatter
	err := json.NewDecoder(r).Decode(&f)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "%#v\n", f)
	return err
}
