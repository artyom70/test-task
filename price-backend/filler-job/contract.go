package fillerjob

type CurrencyPairs struct {
	FromCurrency string `db:"from_currency"`
	ToCurrency   string `db:"to_currency"`
}
