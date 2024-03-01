package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
)

var (
	flagPlaTinumCount int
	flagGoldCount     int
	flagSilverCount   int
)

func init() {
	flag.IntVar(&flagPlaTinumCount, "p", 2, "counts of platinum plan")
	flag.IntVar(&flagGoldCount, "g", 2, "counts of gold plan")
	flag.IntVar(&flagSilverCount, "s", 10, "counts of silver plan")
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type applicant struct {
	company string
	plan    plan
	next    bool
}

type plan string

const (
	planPlaTinum plan = "platinum"
	planGold     plan = "gold"
	planSilver   plan = "silver"
	planBronze   plan = "bronze"
	planFree     plan = "free"
)

func (p plan) IsLottery() bool {
	switch p {
	case planPlaTinum:
		return true
	case planGold:
		return true
	case planSilver:
		return true
	case planBronze:
		return false
	case planFree:
		return false
	}

	return false
}

func (p plan) Title() string {
	switch p {
	case planPlaTinum:
		return `Platinum "Go"ld`
	case planGold:
		return `"Go"ld`
	case planSilver:
		return "Silver"
	case planBronze:
		return "Bronze"
	case planFree:
		return "Free"
	}

	return "unknown plan"
}

func (p plan) Limit() int {
	switch p {
	case planPlaTinum:
		return flagPlaTinumCount
	case planGold:
		return flagGoldCount
	case planSilver:
		return flagSilverCount
	}
	return -1
}

func (p plan) Delay() time.Duration {
	switch p {
	case planPlaTinum:
		return 1 * time.Second
	case planGold:
		return 1 * time.Second
	case planSilver:
		return 300 * time.Millisecond
	}
	return 0
}

func (p plan) Next() plan {
	switch p {
	case planPlaTinum:
		return planGold
	case planGold:
		return planSilver
	case planSilver:
		return planBronze
	case planBronze:
		return planBronze
	default:
		return planFree
	}
}

func run() error {

	applicants, err := parseCSV()
	if err != nil {
		return err
	}

	lotteryAll(applicants)

	return nil
}

func parseCSV() (map[plan][]applicant, error) {
	cr := csv.NewReader(os.Stdin)
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}

	records = records[1:] // skip header
	applicants := make(map[plan][]applicant, len(records))
	companies := make(map[string]bool, len(records))
	for _, record := range records {

		// skip duplicated company
		if companies[record[0]] {
			fmt.Fprintln(os.Stderr, record[0], "is duplicated")
			continue
		}
		companies[record[0]] = true

		a := applicant{
			company: record[0],
			plan:    plan(record[1]),
		}
		a.next, err = strconv.ParseBool(record[2])
		if err != nil {
			return nil, err
		}

		applicants[a.plan] = append(applicants[a.plan], a)
	}

	return applicants, nil
}

func lotteryAll(applicants map[plan][]applicant) {
	plans := []plan{
		planPlaTinum,
		planGold,
		planSilver,
		planBronze,
		planFree,
	}

	var doLotteries []func()
	for _, plan := range plans {
		doLotteries = append(doLotteries, func() {
			fmt.Printf("==== %s sponsor ====\n", plan.Title())
			nexts := lottery(applicants[plan], plan.Limit(), plan.Delay())
			applicants[plan.Next()] = append(applicants[plan.Next()], nexts...)
			fmt.Println()
		})
	}

	for _, f := range doLotteries {
		f()
		time.Sleep(1 * time.Second)
	}
}

func shuffle(as []applicant) {
	rand.Shuffle(len(as), func(i, j int) {
		as[i], as[j] = as[j], as[i]
	})
}

func lottery(as []applicant, n int, d time.Duration) []applicant {
	shuffle(as)
	n = min(len(as), n)
	if n < 0 {
		n = len(as)
	}

	for i := range n {
		printCompany(as[i].company, d)
	}

	if len(as)-n <= 0 {
		return nil
	}

	var nexts []applicant
	for _, a := range as[n:] {
		if a.next {
			nexts = append(nexts, a)
		}
	}

	return nexts
}

func printCompany(name string, d time.Duration) {
	var n int
	for _, c := range name {
		fmt.Printf("%c", c)
		n++
		if n < 6 {
			time.Sleep(d)
		}
	}
	fmt.Println()
	time.Sleep(d)
}
