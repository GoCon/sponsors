package sponsors

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Applicant struct {
	Name string
	Plan Plan
	Next bool
}

func ParseCSV(r io.Reader) (map[Plan][]*Applicant, error) {
	cr := csv.NewReader(r)
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}

	plans := Plans()
	records = records[1:] // skip header
	applicants := make(map[Plan][]*Applicant, len(plans))
	names := make(map[string]bool, len(records))
	for _, record := range records {

		// skip duplicated company
		if names[record[0]] {
			fmt.Fprintln(os.Stderr, record[0], "is duplicated")
			continue
		}
		names[record[0]] = true

		a := &Applicant{
			Name: record[0],
			Plan: Plan(record[1]),
		}
		next, err := strconv.ParseBool(record[2])
		if err != nil {
			return nil, err
		}
		a.Next = next

		applicants[a.Plan] = append(applicants[a.Plan], a)
	}

	return applicants, nil
}
