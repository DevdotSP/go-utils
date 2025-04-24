package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var structure = map[string][]string{
	"controllers":  {"controller.go"},
	"repositories": {"interface.go", "repository.go"},
	"services":     {"service.go"},
	"models":       {"model.go"},
	"routes":       {"routes.go"},
}

// GenerateModule creates a module folder with boilerplate files.
func GenerateModule() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file")
	}

	// Get base directory and module name from environment variables
	baseDir := os.Getenv("BASE_DIR")
	module := os.Getenv("MODULE_NAME")

	// If either baseDir or module name is missing, return an error
	if baseDir == "" || module == "" {
		return fmt.Errorf("BASE_DIR and MODULE_NAME must be set in the environment variables")
	}

	module = strings.ToLower(module)
	// Define the full path for the new module inside services
	basePath := fmt.Sprintf("%s/package/services/%s", baseDir, module)

	// Iterate through the folder structure
	for folder, files := range structure {
		fullPath := fmt.Sprintf("%s/%s", basePath, folder)
		if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create folder %s: %w", fullPath, err)
		}

		// Create and write the files inside each folder
		for _, file := range files {
			filePath := fmt.Sprintf("%s/%s", fullPath, file)
			f, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", filePath, err)
			}
			defer f.Close()
			_, err = f.WriteString(generateBoilerplate(module, file)) // Removed 'folder'
			if err != nil {
				return fmt.Errorf("failed to write boilerplate to file %s: %w", filePath, err)
			}
		}
	}

	fmt.Printf("Module '%s' generated successfully.\n", module)
	return nil
}

// generateBoilerplate generates the boilerplate code for a given file.
func generateBoilerplate(module, file string) string {
	// Capitalize module name for PascalCase format
	modulePascal := cases.Title(language.English).String(strings.ReplaceAll(module, "_", " "))

	switch file {
	case "controller.go":
		return fmt.Sprintf(`package controllers

import "%s/services"

type Controller struct {
	s *services.Service
}

func NewController(s *services.Service) *Controller {
	return &Controller{s: s}
}
`, module)

	case "interface.go":
		return `package repositories

type User interface {
	// define method signatures here
}
`

	case "repository.go":
		return `package repositories

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) User {
	return &Repository{db: db}
}
`

	case "service.go":
		return fmt.Sprintf(`package services

import "%s/repositories"

type Service struct {
	r repositories.User
}

func NewService(r repositories.User) *Service {
	return &Service{r: r}
}
`, module)

	case "model.go":
		return `package models

// Define your model structs here
`

	case "routes.go":
		return fmt.Sprintf(`package routes

import (
	"%s/controllers"
	"%s/repositories"
	"%s/services"

	"github.com/gofiber/fiber/v3"
)

func Register%sRoute(app fiber.Router) {
	r := repositories.NewRepository(nil) // TODO: Replace nil with actual DB
	s := services.NewService(r)
	c := controllers.NewController(s)

	// Define routes
	app.Post("/demo", c.Demo)
}
`, module, module, module, modulePascal)
	}

	return ""
}
