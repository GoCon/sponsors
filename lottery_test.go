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
				{Name: "GoldTinum 01", Plan: sponsors.PlanGold},
				{Name: "GoldTinum 02", Plan: sponsors.PlanGold},
				{Name: "GoldTinum 03", Plan: sponsors.PlanGold},
			},
			sponsors.PlanSilver: []*sponsors.Applicant{
				{Name: "SilverTinum 01", Plan: sponsors.PlanSilver},
				{Name: "SilverTinum 02", Plan: sponsors.PlanSilver},
				{Name: "SilverTinum 03", Plan: sponsors.PlanSilver},
			},
			sponsors.PlanLunch: []*sponsors.Applicant{
				{Name: "LunchTinum 01", Plan: sponsors.PlanLunch},
				{Name: "LunchTinum 02", Plan: sponsors.PlanLunch},
			},
			sponsors.PlanDrink: []*sponsors.Applicant{
				{Name: "DrinkTinum 01", Plan: sponsors.PlanDrink},
				{Name: "DrinkTinum 02", Plan: sponsors.PlanDrink},
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
