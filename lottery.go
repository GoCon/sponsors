package sponsors

import (
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"slices"
	"time"
)

type Lottery struct {
	GoldCount   int
	SilverCount int
	LunchCount  int
	DrinkCount  int
}

func (l *Lottery) Limit(p Plan) int {
	switch p {
	case PlanGold:
		return l.GoldCount
	case PlanSilver:
		return l.SilverCount
	case PlanLunch:
		return l.LunchCount
	case PlanDrink:
		return l.DrinkCount
	}
	return -1
}

func (l *Lottery) Do(applicants map[Plan][]*Applicant) LotteryResult {
	plans := Plans()
	doLotteries := make([]func() (Plan, []*Applicant), 0, len(plans))
	for _, plan := range plans {
		doLotteries = append(doLotteries, func() (Plan, []*Applicant) {
			sponsors := l.doPlan(applicants[plan], l.Limit(plan))
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

func (l *Lottery) doPlan(as []*Applicant, n int) (sponsors []*Applicant) {
	rand.Shuffle(len(as), func(i, j int) {
		as[i], as[j] = as[j], as[i]
	})

	n = min(len(as), n)
	if n < 0 {
		n = len(as)
	}

	sponsors = slices.Clone(as[:n])

	if len(as)-n <= 0 {
		return sponsors
	}

	return sponsors
}

type LotteryResult map[Plan][]*Applicant

func (r LotteryResult) Show(w io.Writer) {
	for _, plan := range Plans() {
		if !plan.IsLottery() {
			continue
		}
		title, err := plan.Title()
		if err, ok := errors.AsType[ErrUnknownPlan](err); ok {
			title = err.Error()
		}
		fmt.Fprintf(w, "==== %s sponsor ====\n", title)
		for _, applicant := range r[plan] {
			r.printApplicant(w, applicant.Name, r.PlanDelay(plan))
		}
		fmt.Fprintln(w)
		time.Sleep(1 * time.Second)
	}
}

func (r LotteryResult) PlanDelay(p Plan) time.Duration {
	switch p {
	case PlanGold:
		return 1 * time.Second
	case PlanSilver, PlanLunch, PlanDrink:
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
