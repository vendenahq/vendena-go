package vendena

// The OrderFormData model.
type OrderFormData struct {
	URL      string            `json:"url"`
	Elements map[string]string `json:"elements"`
}
