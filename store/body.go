package store

import "time"

// // Body contains userbody-related info
type Body struct {
	DateTime   time.Time `json:"datetime"`
	Height     float32   `json:"height"`
	Weight     float32   `json:"weight"`
	FatPercent float32   `json:"fat_percent"`
	Waist      float32   `json:"waist"`
	Arms       float32   `json:"arms"`
	Foream     float32   `json:"foream"`
	Chest      float32   `json:"chest"`
	Neck       float32   `json:"neck"`
	Thighs     float32   `json:"thighs"`
	Thigh      float32   `json:"thigh"`
	Calf       float32   `json:"calf"`
}
