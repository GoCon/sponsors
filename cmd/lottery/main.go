package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GoCon/sponsors"
)

var (
	flagPlaTinumCount int
	flagGoldCount     int
	flagSilverCount   int
)

func init() {
	flag.IntVar(&flagPlaTinumCount, "p", 2, "counts of platinum plan")
	flag.IntVar(&flagGoldCount, "g", 2, "counts of gold plan")
	flag.IntVar(&flagSilverCount, "s", 18, "counts of silver plan")
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	applicants, err := sponsors.ParseCSV(os.Stdin)
	if err != nil {
		return err
	}

	l := &sponsors.Lottery{
		PlaTinumCount: flagPlaTinumCount,
		GoldCount:     flagGoldCount,
		SilverCount:   flagSilverCount,
	}

	result := l.Do(applicants)
	result.Show(os.Stdout)

	return nil
}
