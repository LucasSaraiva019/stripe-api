package datastruct

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	Description string `json:"description"`
}
