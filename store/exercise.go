package store

// Exercise contains all exercise
type Exercise struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
