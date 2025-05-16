package migration

import (
	"fmt"
	"log"

	"github.com/DevdotSP/go-utils/config"
	sharedModels "github.com/DevdotSP/go-utils/shared-models"
)

func MigrationTable() {

	err := config.DB.AutoMigrate(
		&sharedModels.WebUser{}, // ✅ User table first
		&sharedModels.Role{}, // ✅ Role (if User depends on Role)
		&sharedModels.UserLoginHistory{},
		&sharedModels.Address{},
		&sharedModels.Region{},
		&sharedModels.Province{},
		&sharedModels.Municipality{},
		&sharedModels.Barangay{},
		&sharedModels.SidebarItem{},
		&sharedModels.UserRoleSidebar{},
		&sharedModels.UserImage{},  // ✅ Move UserImage after User
		&sharedModels.Notification{}, // ✅ Move Notification after User
		&sharedModels.PasswordResetToken{},
		&sharedModels.Advertisement{},
		&sharedModels.UserExportRequest{},
	)
	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	fmt.Println("✅ Database AutoMigration completed successfully!")
}