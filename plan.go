package sponsors

type Plan string

func Plans() []Plan {
	return []Plan{
		PlanPlaTinum,
		PlanGold,
		PlanSilver,
		PlanBronze,
		PlanFree,
	}
}

const (
	PlanPlaTinum Plan = "platinum"
	PlanGold     Plan = "gold"
	PlanSilver   Plan = "silver"
	PlanBronze   Plan = "bronze"
	PlanFree     Plan = "free"
)

func (p Plan) IsLottery() bool {
	switch p {
	case PlanPlaTinum, PlanGold, PlanSilver:
		return true
	default:
		return false
	}
}

func (p Plan) Title() string {
	switch p {
	case PlanPlaTinum:
		return `Platinum "Go"ld`
	case PlanGold:
		return `"Go"ld`
	case PlanSilver:
		return "Silver"
	case PlanBronze:
		return "Bronze"
	case PlanFree:
		return "Free"
	}

	return "unknown plan"
}

func (p Plan) Next() Plan {
	switch p {
	case PlanPlaTinum:
		return PlanGold
	case PlanGold:
		return PlanSilver
	case PlanSilver:
		return PlanBronze
	case PlanBronze:
		return PlanBronze
	default:
		return PlanFree
	}
}
