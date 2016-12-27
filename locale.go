package vendena

// The Locale model.
type Locale struct {
	ID                int64  `json:"id"`
	Title             string `json:"title"`
	Code              string `json:"code"`
	ThousandSeparator string `json:"thousand_separator"`
	DecimalMark       string `json:"decimal_mark"`
}
