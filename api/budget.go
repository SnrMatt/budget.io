package api

//Handle Budget Routes

type Bill struct {
	name        string
	descirption string
	occurance   string
	total       float64
}

type Budget struct {
	income    float64
	occurance string
	bills     []Bill
}

func newBudget(income float64, occurance string) *Budget {
	return &Budget{
		income:    income,
		occurance: occurance,
		bills:     make([]Bill, 0),
	}
}
