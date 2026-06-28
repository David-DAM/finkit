package currency

type Currency struct {
	IsoCode    string `json:"iso_code"`
	IsoNumeric string `json:"iso_numeric"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	StartDate  string `json:"start_date"`
}

type Currencies struct {
	Currencies []Currency
}

type Rate struct {
	Date  string  `json:"date"`
	Base  string  `json:"base"`
	Quote string  `json:"quote"`
	Rate  float64 `json:"rate"`
}
