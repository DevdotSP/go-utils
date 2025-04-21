package sharedModels

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// WebUser Model.
// @swagger:model
type WebUser struct {
	ID                 int               `gorm:"primarykey;autoIncrement" json:"id"`
	Email              string            `json:"email" gorm:"not null;unique" `
	IsVerified         bool              `json:"is_verified" gorm:"default:false"`
	Token              string            `json:"token,omitempty"`
	FullName           string            `json:"full_name"`
	IsLock             string            `json:"is_lock" gorm:"default:0"`
	MobileNo           string            `json:"mobile_no"`
	MustChangePassword string            `json:"must_change_password" gorm:"default:0"`
	UserName           string            `json:"user_name" gorm:"not null;unique" `
	Password           string            `json:"password"`
	PwdExpiredDate     time.Time         `json:"pwd_expired_date,omitempty"`
	Status             string            `json:"status" gorm:"default:1"`
	RoleID             int               `json:"role_id,omitempty" gorm:"null"`
	Role               Role `gorm:"foreignKey:RoleID;references:ID"` // Foreign key reference
	Logged             string            `json:"logged" gorm:"default:0"`
	FirstName          string            `json:"first_name"`
	MiddleName         string            `json:"middle_name"`
	LastName           string            `json:"last_name"`
	Birthday           time.Time         `json:"birthday,omitempty"`
	CreatedBy          string            `json:"created_by"`
	UpdatedBy          string            `json:"updated_by"`

	// Corrected foreign key references
	UserImage    UserImage      `json:"user_images" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Notification []Notification `json:"notifications" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Address      Address        `json:"address" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time `gorm:"autoCreateTime;type:timestamptz"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;type:timestamptz"`
}

// TableName overrides the default table name
func (WebUser) TableName() string {
	return "v1.web_user"
}

// HashPassword hashes a plain text password
func (u *WebUser) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares a hashed password with a plain text input
func (u *WebUser) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type CustomTime struct {
	time.Time
}

// Implement custom JSON unmarshaling
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1] // Remove quotes from JSON string

	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}

	ct.Time = parsedTime
	return nil
}
