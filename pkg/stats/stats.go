package stats

import (
	"github.com/olim007/bank/pkg/bank/types"
)

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

func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, papayment := range payments {
		categories[papayment.Category] += papayment.Amount
	}

	return categories
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	avgcatcount := map[types.Category]int{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		avgcatcount[payment.Category] += 1
	}
	for key, value := range categories {
		categories[key] = value / types.Money(avgcatcount[key])
	}
	return categories
}