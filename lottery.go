package sponsors

import (
	"fmt"
	"io"
	"math/rand/v2"
	"slices"
	"time"
)

type Lottery struct {
	PlaTinumCount int
	GoldCount     int
	SilverCount   int
}

func (l *Lottery) Limit(p Plan) int {
	switch p {
	case PlanPlaTinum:
		return l.PlaTinumCount
	case PlanGold:
		return l.GoldCount
	case PlanSilver:
		return l.SilverCount
	}
	return -1
}

func (l *Lottery) Do(applicants map[Plan][]*Applicant) LotteryResult {
	plans := Plans()
	doLotteries := make([]func() (Plan, []*Applicant), 0, len(plans))
	for _, plan := range plans {
		doLotteries = append(doLotteries, func() (Plan, []*Applicant) {
			sponsors, nexts := l.doPlan(applicants[plan], l.Limit(plan))
			applicants[plan.Next()] = append(applicants[plan.Next()], nexts...)
			return plan, sponsors
		})
	}

	result := make(LotteryResult)
	for _, f := range doLotteries {
		plan, sponsors := f()
		result[plan] = sponsors
	}

	return result
}

func (l *Lottery) doPlan(as []*Applicant, n int) (sponsors, nexts []*Applicant) {
	rand.Shuffle(len(as), func(i, j int) {
		as[i], as[j] = as[j], as[i]
	})

	n = min(len(as), n)
	if n < 0 {
		n = len(as)
	}

	sponsors = slices.Clone(as[:n])

	if len(as)-n <= 0 {
		return sponsors, nil
	}

	for _, a := range as[n:] {
		if a.Next {
			nexts = append(nexts, a)
		}
	}

	return sponsors, nexts
}

type LotteryResult map[Plan][]*Applicant

func (r LotteryResult) Show(w io.Writer) {
	for _, plan := range Plans() {
		if !plan.IsLottery() {
			continue
		}
		fmt.Fprintf(w, "==== %s sponsor ====\n", plan.Title())
		for _, applicant := range r[plan] {
			r.printApplicant(w, applicant.Name, r.PlanDelay(plan))
		}
		fmt.Fprintln(w)
		time.Sleep(1 * time.Second)
	}
}

func (r LotteryResult) PlanDelay(p Plan) time.Duration {
	switch p {
	case PlanPlaTinum:
		return 1 * time.Second
	case PlanGold:
		return 1 * time.Second
	case PlanSilver:
		return 300 * time.Millisecond
	}
	return 0
}

func (r LotteryResult) printApplicant(w io.Writer, name string, d time.Duration) {
	var n int
	for _, c := range name {
		fmt.Fprintf(w, "%c", c)
		n++
		if n < 6 {
			time.Sleep(d)
		}
	}
	fmt.Fprintln(w)
	time.Sleep(d)
}
