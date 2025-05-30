package command

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

func capitalizeFirst(str string) string {
	caser := cases.Title(language.English)
	return caser.String(str)
}

func getRelativePathFromModule(absPath string, moduleName string) (string, error) {
	i := strings.Index(absPath, moduleName)
	if i == -1 {
		return "", fmt.Errorf("module name %s not found in path", moduleName)
	}
	return absPath[i:], nil
}

func generateBoilerplate(module, file string) string {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err).Error()
	}

	relPath, _ := getRelativePathFromModule(cwd, "gofiberv3")
	// relPath is now: gofiberv3/package/services/customer/services

	basePath := filepath.Join(relPath, "package", "services")

	capitalModule := capitalizeFirst(module)

	// Format module names
	modulePascal := cases.Title(language.English).String(strings.ReplaceAll(module, "_", " "))
	moduleCamel := strings.ToLower(module[:1]) + module[1:]

	switch file {
	case "controller.go":
		return fmt.Sprintf(`package controllers

import (
	"github.com/gofiber/fiber/v3"
	"%s/%s/services"
)

type Controller struct {
	s *services.Service
}

func NewController(s *services.Service) *Controller {
	return &Controller{s: s}
}

func (ctl *Controller) Index(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "List all %s"})
}

func (ctl *Controller) Show(c fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "Show %s", "id": id})
}

func (ctl *Controller) Store(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Store new %s"})
}

func (ctl *Controller) Update(c fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "Update %s", "id": id})
}

func (ctl *Controller) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "Delete %s", "id": id})
}
`, basePath, module, module, module, module, module, module)

	case "interface.go":
		return fmt.Sprintf(`package repositories

type %s interface {
	// Define your repository methods here
}
`, modulePascal)

	case "repository.go":
		return fmt.Sprintf(`package repositories

import (
	"gorm.io/gorm"
	"%s/%s/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(entity *models.%s) error {
	return r.db.Create(entity).Error
}

func (r *Repository) FindByID(id uint) (*models.%s, error) {
	var entity models.%s
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *Repository) FindAll() ([]models.%s, error) {
	var entities []models.%s
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *Repository) Update(entity *models.%s) error {
	return r.db.Save(entity).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&models.%s{}, id).Error
}
`, basePath, module, capitalModule, capitalModule, capitalModule, capitalModule, capitalModule, capitalModule, capitalModule)

	case "service.go":
		return fmt.Sprintf(`package services

import "%s/%s/repositories"

type Service struct {
	r repositories.%s
}

func NewService(r repositories.%s) *Service {
	return &Service{r: r}
}
`, basePath, module, modulePascal, modulePascal)

	case "model.go":
		return fmt.Sprintf(`package models

		type %s struct {}  // Define your model structs here


`, capitalModule)

	case "routes.go":
		return fmt.Sprintf(`package routes

import (
	"%s/%s/controllers"
	"%s/%s/repositories"
	"%s/%s/services"

	"github.com/gofiber/fiber/v3"
)

func Register%sRoute(app fiber.Router) {
	r := repositories.NewRepository(nil) // TODO: Replace nil with actual DB
	s := services.NewService(r)
	c := controllers.NewController(s)

	group := app.Group("/%s")
	group.Get("/", c.Index)
	group.Get("/:id", c.Show)
	group.Post("/", c.Store)
	group.Put("/:id", c.Update)
	group.Delete("/:id", c.Delete)
}
`, basePath, module, basePath, module, basePath, module, modulePascal, moduleCamel)
	}

	return ""
}


