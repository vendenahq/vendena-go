package vendena

// The OrderNotificationValidation model.
type OrderNotificationValidation struct {
	Querystring string `json:"querystring"`
}

// The OrderNotificationValidationResult model.
type OrderNotificationValidationResult struct {
	Valid        bool   `json:"valid"`
	ResponseText string `json:"response_text"`
}
