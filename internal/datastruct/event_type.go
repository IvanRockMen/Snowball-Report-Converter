package datastruct

const (
	EmptyEventType           EventType = ""
	BuyEventType             EventType = "Buy"
	SellEventType            EventType = "Sell"
	DividendEventType        EventType = "Dividend"
	StockAsDividendEventType EventType = "Stock_As_Dividend"
	SplitEventType           EventType = "Split"
	SpinoffEventType         EventType = "Spinoff"
	FeeEventType             EventType = "Fee"
	AmortisationEventType    EventType = "Amortisation"
	RepaymentEventType       EventType = "Repayment"
	CashInEventType          EventType = "Cash_In"
	CashOutEventType         EventType = "Cash_Out"
	CashGainEventType        EventType = "Cash_Gain"
	CashExpenseEventType     EventType = "Cash_Expense"
	CashConvertEventType     EventType = "Cash_Convert"
)

type EventType string

func (et EventType) ToString() string {
	return string(et)
}
