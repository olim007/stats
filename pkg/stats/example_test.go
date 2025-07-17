package stats

func ExampleAvg() {
    payments := []types.Payment{
    {
	ID:       1,
	Amount:   2,
	Category: "visa",
    },
    {
	ID:       2,
	Amount:   6,
	Category: "milly",
    },
    {
	ID:       1,
	Amount:   3,
	Category: "visa",
    },
    {
	ID:       2,
	Amount:   9,
	Category: "milly",
    },
    }

    var money types.Money = Avg(payments)
    fmt.Println(money)

    // Output:
    // 5
}

func ExampleTotalInCategory() {
    payments := []types.Payment{
    {
	ID:       1,
	Amount:   2,
	Category: "visa",
    },
    {
	ID:       2,
	Amount:   6,
	Category: "milly",
    },
    {
        ID:       1,
	Amount:   3,
	Category: "visa",
    },
    {
	ID:       2,
	Amount:   9,
	Category: "milly",
    },
    }

    var money types.Money = TotalInCategory(payments, "milly")
    fmt.Println(money)

    // Output:
    // 7
}
