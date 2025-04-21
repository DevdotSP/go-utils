package sharedModels

import (
	"time"
)

type Notification struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    *int      `json:"user_id" gorm:"index;default:null"` // Nullable for broadcasts
	UserType  string    `json:"user_type" gorm:"not null"`         // "customer", "agent", "officer", or "system"
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"not null"`
	Token     *string   `json:"token" gorm:"default:null"`         // Nullable for topic notifications
	Topic     *string   `json:"topic" gorm:"default:null"`         // New field for topic-based notifications
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	TargetAll bool      `json:"target_all" gorm:"default:false"`   // If true, broadcast to all
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName explicitly sets the table name
func (Notification) TableName() string {
	return "v1.notification"
}


type SubscriptionRequest struct {
	Token string `json:"token" validate:"required"`
	Topic string `json:"topic" validate:"required"`
}
