package models

type Customer struct {
	MonthyIncome      float64 `json:"monthlyIncome"`
	CurrentDebitRatio float64 `json:"currentDebitRation"`
	CreditScore       int     `json:"creditScore"`
	DefaultHistory    History `json:"defaultHistory"`
}

type History struct {
	hasRecenteDefault     bool `json:"hasRecenteDealt"`
	defaultCount          int  `json:"defaultCount"`
	maxDelay              int  `json:"maxDelay"`
	hasActiveRestrictions bool `json:"hasActiveRestrictions"`
}
