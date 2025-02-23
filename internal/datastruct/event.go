package datastruct

import (
	"strconv"
	"time"
)

type Event struct {
	EventType       EventType
	Date            time.Time
	Symbol          string
	Price           float64
	Quantity        int
	Currency        string
	FeeTax          float64
	Exchange        string
	NKD             float64
	FeeCurrency     string
	DoNotAdjustCash Bool
	Note            string
}

func (e Event) ToCsvRow() []string {
	row := make([]string, 0, 12)

	row[0] = e.EventType.ToString()
	row[1] = e.Date.Format(time.DateOnly)
	row[2] = e.Symbol
	row[3] = strconv.FormatFloat(e.Price, 'f', 3, 64)
	row[4] = strconv.FormatInt(int64(e.Quantity), 10)
	row[5] = e.Currency
	row[6] = strconv.FormatFloat(e.FeeTax, 'f', 3, 64)
	row[7] = e.Exchange
	row[8] = strconv.FormatFloat(e.NKD, 'f', 3, 64)
	row[9] = e.FeeCurrency
	if e.DoNotAdjustCash.Exists {
		row[10] = strconv.FormatBool(e.DoNotAdjustCash.Value)
	} else {
		row[10] = ""
	}
	row[11] = e.Note

	return row
}
