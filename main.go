package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	output = flag.String("o", "", "output SVG file (default GENE_SYMBOL.svg)")
	width  = flag.Int("w", 700, "SVG output width")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] GENE_SYMBOL [PROTEIN CHANGES ...]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Where options are:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if *output == "" {
		*output = flag.Arg(0) + ".svg"
	}

	f, err := os.OpenFile(*output, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprintln(os.Stderr, "HGNC Symbol: ", flag.Arg(0))

	acc, err := GetProtID(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Uniprot/SwissProt Accession: ", acc)

	data, err := GetPfamGraphicData(acc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Drawing diagram to", *output)
	DrawSVG(f, *width, flag.Args()[1:], data)
}
