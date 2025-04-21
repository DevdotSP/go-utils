package sharedModels

import (
	"time"
)

// UserImage represents an uploaded image associated with a user.
type UserImage struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `json:"user_id" gorm:"index"` // Index for faster queries
	ImageType string    `json:"image_type" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;type:timestamptz;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;type:timestamptz;default:now()" json:"updated_at"`
}

// TableName explicitly sets the table name for UserImage.
func (UserImage) TableName() string {
	return "v1.user_image" // Consider using plural naming convention
}
