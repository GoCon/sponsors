package sponsors

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

func (p Plan) Title() string {
	switch p {
	case PlanGold:
		return `"Go"ld`
	case PlanSilver:
		return "Silver"
	case PlanLunch:
		return "Lunch"
	case PlanDrink:
		return "Drink"
	default:
		return "unknown plan"
	}
}

func (p Plan) Next() Plan {
	switch p {
	case PlanGold:
		return PlanSilver
	case PlanSilver:
		return PlanLunch
	case PlanLunch:
		return PlanDrink
	default:
		return ""
	}
}
