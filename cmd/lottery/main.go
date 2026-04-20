package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GoCon/sponsors"
)

var (
	flagGoldCount   int
	flagSilverCount int
	flagLunchCount  int
	flagDrinkCount  int
)

func init() {
	flag.IntVar(&flagGoldCount, "g", 6, "counts of gold plan")
	flag.IntVar(&flagSilverCount, "s", 12, "counts of silver plan")
	flag.IntVar(&flagLunchCount, "l", 2, "counts of lunch plan")
	flag.IntVar(&flagDrinkCount, "d", 2, "counts of drink plan")
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
		GoldCount:   flagGoldCount,
		SilverCount: flagSilverCount,
		LunchCount:  flagLunchCount,
		DrinkCount:  flagDrinkCount,
	}

	result := l.Do(applicants)
	result.Show(os.Stdout)

	return nil
}
