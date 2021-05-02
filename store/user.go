package store

// User contain user-related info
type User struct {
	ID       uint64 `json:"-"`
	Name     string `json:"name"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	// Avatar    string    `json:"avatar"  binding:"required"`
	// Admin     bool      `json:"admin"  binding:"required"`
	// Blocked   bool      `json:"blocked"  binding:"required"`
	// Age       uint8     `json:"age"  binding:"required"`
	// CreatedAt time.Time `json:"created_at"  binding:"required"`
	// UpdatedAt time.Time `json:"updated_at"  binding:"required"`
}
