package utils

var currencyMultiplier = map[string]float64{"USD": 1.0, "RUB": 77.85, "TWD": 28.72, "HKD": 7.75,
	"JPY": 105.57, "MYR": 4.15, "THB": 31.29, "IDR": 14688.25, "PHP": 48.54,
}

func NewCurrencyProvider() CurrencyRates {
	return CurrencyRates{}
}

type CurrencyRates struct{}

func (r CurrencyRates) Rate(currency string) float64 {
	if mul, ok := currencyMultiplier[currency]; ok {
		return mul
	}
	return currencyMultiplier["USD"]
}
