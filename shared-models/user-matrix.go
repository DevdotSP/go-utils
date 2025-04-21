package sharedModels

import (
	"encoding/json"
	"sort"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Error Model.
// @swagger:model
type ErrorModel struct {
	RetCode int    `json:"ret_code"` // Error Code
	Message string `json:"message"`  // Error Message
}

// Error Model
// @swagger:model
type CreateRoleRequest struct {
	Role       Role  `json:"role"`
	SidebarIDs []int `json:"sidebar_ids"`
}

// ChangePassReq Model.
// @swagger:model
type ChangePassReq struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateUserRequest struct {
	Email     string    `json:"email,omitempty"`
	MobileNo  string    `json:"mobile_no,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	RoleID    int       `json:"role_id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	MiddleName string   `json:"middle_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	UpdatedBy string    `json:"updated_by,omitempty"`
}

type RoleDTO struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// Role Model.
// @swagger:model
type Role struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code   string `gorm:"unique;not null" json:"code"`
	Name   string `gorm:"unique;not null" json:"name"`
	Status string `gorm:"not null" json:"status"`

	// One-to-One Relationship with UserRoleSidebar
	UserRoleSidebar *UserRoleSidebar `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	// Computed field (Not stored in DB)
	SidebarItems []*SidebarItem `gorm:"-" json:"sidebar_items"`
}

// TableName overrides the default table name
func (Role) TableName() string {
	return "v1.role"
}

func (r *Role) LoadSidebarItems() error {
	if r.UserRoleSidebar == nil || len(r.UserRoleSidebar.SidebarItems) == 0 {
		r.SidebarItems = []*SidebarItem{} // ✅ Ensure empty slice instead of nil
		return nil
	}

	var flatSidebarItems []*SidebarItem
	if err := json.Unmarshal(r.UserRoleSidebar.SidebarItems, &flatSidebarItems); err != nil {
		return err
	}

	// ✅ Ensure unique and properly nested sidebar items
	r.SidebarItems = BuildSidebarHierarchy(flatSidebarItems)

	// ✅ Sort sidebar items by ID
	sort.Slice(r.SidebarItems, func(i, j int) bool {
		return r.SidebarItems[i].ID < r.SidebarItems[j].ID
	})

	return nil
}

func BuildSidebarHierarchy(items []*SidebarItem) []*SidebarItem {
	itemMap := make(map[int]*SidebarItem)
	var rootItems []*SidebarItem

	// ✅ Step 1: Initialize item map and ensure children slices
	for _, item := range items {
		if item == nil || item.ID == 0 {
			continue // ✅ Skip invalid items
		}
		item.Children = []*SidebarItem{} // ✅ Ensure empty children slice
		itemMap[item.ID] = item
	}

	// ✅ Step 2: Assign children to their parents
	for _, item := range items {
		if item.ParentID != nil { // If item has a parent
			if parent, exists := itemMap[*item.ParentID]; exists {
				parent.Children = append(parent.Children, item) // ✅ Correct reference
			}
		} else {
			rootItems = append(rootItems, item) // ✅ Root level items
		}
	}

	// ✅ Step 3: Sort root items and children by ID (ascending)
	sort.Slice(rootItems, func(i, j int) bool {
		return rootItems[i].ID < rootItems[j].ID
	})

	for _, item := range itemMap {
		sort.Slice(item.Children, func(i, j int) bool {
			return item.Children[i].ID < item.Children[j].ID
		})
	}

	return rootItems
}

// BeforeSave encodes sidebar items before updating/creating
func (r *Role) BeforeSave(tx *gorm.DB) error {
	if r.UserRoleSidebar != nil && len(r.SidebarItems) > 0 {
		data, err := json.Marshal(r.SidebarItems)
		if err != nil {
			return err
		}
		r.UserRoleSidebar.SidebarItems = datatypes.JSON(data)
	}
	return nil
}

// SidebarItem model
type SidebarItem struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string         `gorm:"not null;unique" json:"title"`
	Icon      string         `gorm:"not null" json:"icon"`
	Route     *string        `gorm:"unique" json:"route,omitempty"`
	ParentID  *int           `json:"parent_id,omitempty"`
	Parent    *SidebarItem   `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent,omitempty"`
	Children  []*SidebarItem `gorm:"foreignKey:ParentID" json:"children,omitempty"` // Use pointer slice ✅
	IsEnabled *bool          `gorm:"default:true" json:"is_enabled"`
}

// TableName overrides the default table name
func (SidebarItem) TableName() string {
	return "v1.sidebar_item"
}

// UserRoleSidebar model
type UserRoleSidebar struct {
	RoleID       int            `gorm:"primaryKey;not null" json:"role_id"`
	SidebarItems datatypes.JSON `gorm:"type:jsonb;not null" json:"sidebar_items"`
	IsEnabled    bool           `gorm:"default:true" json:"is_enabled"`

	Role Role `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;" json:"-"`
}

// TableName overrides the default table name
func (UserRoleSidebar) TableName() string {
	return "v1.user_role_sidebar"
}

// BeforeCreate initializes SidebarItems as an empty array if nil
func (u *UserRoleSidebar) BeforeCreate(tx *gorm.DB) error {
	if u.SidebarItems == nil {
		u.SidebarItems = datatypes.JSON("[]")
	}
	return nil
}
