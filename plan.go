package sponsors

import "fmt"

type Plan string

func Plans() []Plan {
	return []Plan{
		PlanGold,
		PlanSilver,
		PlanLunch,
		PlanDrink,
	}
}

const (
	PlanGold   Plan = "gold"
	PlanSilver Plan = "silver"
	PlanLunch  Plan = "lunch"
	PlanDrink  Plan = "drink"
)

func (p Plan) IsLottery() bool {
	switch p {
	case PlanGold, PlanSilver, PlanLunch, PlanDrink:
		return true
	default:
		return false
	}
}

type ErrUnknownPlan struct{}

func (e ErrUnknownPlan) Error() string {
	return "unknown plan"
}

func (p Plan) Title() (string, error) {
	switch p {
	case PlanGold:
		return `"Go"ld`, nil
	case PlanSilver:
		return "Silver", nil
	case PlanLunch:
		return "Lunch", nil
	case PlanDrink:
		return "Drink", nil
	default:
		return "unknown plan", fmt.Errorf("unknown plan: %s", p)
	}
}

func (p Plan) Next() (Plan, error) {
	switch p {
	case PlanGold:
		return PlanSilver, nil
	case PlanSilver:
		return PlanLunch, nil
	case PlanLunch:
		return PlanDrink, nil
	default:
		return "", fmt.Errorf("unknown plan: %s", p)
	}
}
