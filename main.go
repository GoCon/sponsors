package main

import (
	"encoding/csv"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
)

func main() {
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
	// Platinum "Go"ld sponsor
	fmt.Println(`==== Platinum "Go"ld sponsor ====`)
	goldNext := lottery(applicants[planPlaTinum], 2, 1*time.Second)
	applicants[planGold] = append(applicants[planGold], goldNext...)
	fmt.Println()

	// "Go"ld sponsor
	fmt.Println(`==== "Go"ld sponsor ====`)
	siliverNext := lottery(applicants[planGold], 2, 1*time.Second)
	applicants[planSilver] = append(applicants[planSilver], siliverNext...)
	fmt.Println()

	// Silver sponsor
	fmt.Println("==== Silver sponsor ====")
	bronzeNext := lottery(applicants[planSilver], 10, 100*time.Millisecond)
	applicants[planBronze] = append(applicants[planBronze], bronzeNext...)
	fmt.Println()

	// Bronze sponsor
	fmt.Println("==== Bronze sponsor ====")
	shuffle(applicants[planBronze])
	for _, a := range applicants[planBronze] {
		fmt.Println(a.company)
	}
	fmt.Println()

	// Free sponsor
	fmt.Println("==== Free sponsor ====")
	shuffle(applicants[planFree])
	for _, a := range applicants[planFree] {
		fmt.Println(a.company)
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
	for i := range n {
		printCompany(as[i].company, d)
	}

	var nexts []applicant
	for _, a := range as[2:] {
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
