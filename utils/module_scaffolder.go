package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

// GenerateModule creates a module folder under package/services/{moduleName}
func GenerateModule(module string) error {
	if module == "" {
		return fmt.Errorf("module name is required")
	}

	module = strings.ToLower(module)

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Build base path: {cwd}/package/services/{module}
	basePath := filepath.Join(cwd, "package", "services", module)

	// Create folders and files
	for folder, files := range structure {
		fullPath := filepath.Join(basePath, folder)
		if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create folder %s: %w", fullPath, err)
		}

		for _, file := range files {
			filePath := filepath.Join(fullPath, file)
			f, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", filePath, err)
			}
			defer f.Close()

			_, err = f.WriteString(generateBoilerplate(module, file))
			if err != nil {
				return fmt.Errorf("failed to write to file %s: %w", filePath, err)
			}
		}
	}

	fmt.Printf("Module '%s' created in 'package/services/%s'\n", module, module)
	return nil
}

func generateBoilerplate(module, file string) string {
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
