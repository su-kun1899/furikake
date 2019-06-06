package main

import (
	"flag"
	"fmt"
	"github.com/su-kun1899/furikake"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	headerOpt  = flag.String("h", "", "header values separated by comma")
)

func main() {
	flag.Parse()
	header := strings.Split(*headerOpt, ",")
	fName := flag.Args()[0]

	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatal(err)
	}

	csv, err := furikake.ToCsv(header, string(bytes))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, _ = fmt.Println(csv)
}
