package sharedModels

import "time"

// UserLoginHistory Model.
// @swagger:model
type UserLoginHistory struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Action    string    `gorm:"column:action" json:"action"`
	UserName  string    `gorm:"column:user_name" json:"user_name"`
	UpdatedBy string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;type:timestamptz;column:updated_at" json:"updated_at"`
}

// TableName overrides the default table name
func (UserLoginHistory) TableName() string {
	return "v1.user_login_history"
}
