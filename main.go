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

	cr := csv.NewReader(os.Stdin)
	records, err := cr.ReadAll()
	if err != nil {
		return err
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
			return err
		}

		applicants[a.plan] = append(applicants[a.plan], a)
	}

	time.Sleep(2 * time.Second)

	// Platinum "Go"ld sponsor
	fmt.Println(`==== Platinum "Go"ld sponsor ====`)
	applicants[planGold] = append(applicants[planGold], lottery(applicants[planPlaTinum], 2, 1*time.Second)...)
	fmt.Println()

	// "Go"ld sponsor
	fmt.Println(`==== "Go"ld sponsor ====`)
	applicants[planSilver] = append(applicants[planSilver], lottery(applicants[planGold], 2, 1*time.Second)...)
	fmt.Println()

	// Silver sponsor
	fmt.Println("==== Silver sponsor ====")
	applicants[planBronze] = append(applicants[planBronze], lottery(applicants[planSilver], 10, 100 * time.Millisecond)...)
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

	return nil
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
