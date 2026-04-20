package sponsors_test

import (
	"bytes"
	"flag"
	"io"
	"testing"
	"testing/synctest"

	"github.com/GoCon/sponsors"

	"github.com/tenntenn/golden"
)

var (
	flagUpdateGolden bool
)

func init() {
	flag.BoolVar(&flagUpdateGolden, "update-golden", false, "update golden files")
}

func TestLotteryResult_Show(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		result sponsors.LotteryResult
	}{
		"normal": {result: sponsors.LotteryResult{
			sponsors.PlanGold: []*sponsors.Applicant{
				{Name: "GoldTinum 01", Plan: sponsors.PlanGold, Next: true},
				{Name: "GoldTinum 02", Plan: sponsors.PlanGold, Next: true},
				{Name: "GoldTinum 03", Plan: sponsors.PlanGold, Next: true},
			},
			sponsors.PlanSilver: []*sponsors.Applicant{
				{Name: "SilverTinum 01", Plan: sponsors.PlanSilver, Next: true},
				{Name: "SilverTinum 02", Plan: sponsors.PlanSilver, Next: true},
				{Name: "SilverTinum 03", Plan: sponsors.PlanSilver, Next: true},
			},
			sponsors.PlanLunch: []*sponsors.Applicant{
				{Name: "LunchTinum 01", Plan: sponsors.PlanLunch, Next: true},
				{Name: "LunchTinum 02", Plan: sponsors.PlanLunch, Next: true},
			},
			sponsors.PlanDrink: []*sponsors.Applicant{
				{Name: "DrinkTinum 01", Plan: sponsors.PlanDrink, Next: true},
				{Name: "DrinkTinum 02", Plan: sponsors.PlanDrink, Next: true},
			},
		}},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			synctest.Test(t, func(*testing.T) {
				t.Parallel()

				var got bytes.Buffer
				w := io.MultiWriter(&got, t.Output())
				tt.result.Show(w)

				if diff := golden.Check(t, flagUpdateGolden, "testdata", t.Name(), &got); diff != "" {
					t.Error("unmatched with golden file", diff)
				}
			})
		})
	}
}
