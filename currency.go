package vendena

// The Currency model.
type Currency struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	Precision int    `json:"precision"`
	Unit      string `json:"unit"`
	Format    string `json:"format"`
}
