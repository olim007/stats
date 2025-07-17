package stats

import "github.com/olim007/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	sum := 0
	for _, c := range payments {
		sum += int(c.Amount)
	}
	avg := sum / len(payments)
	return types.Money(avg)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	count := 0
	for _, p := range payments {
		if p.Category == category {
			sum += int(p.Amount)
			count += 1
		}
	}
	result := types.Money(sum / count)
	return result
}
