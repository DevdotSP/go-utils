package importingcsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/DevdotSP/go-utils/config"
"github.com/DevdotSP/go-utils/shared-models"
)


// loadCSV reads a CSV file and returns the records as [][]string
func loadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	return reader.ReadAll()
}

func importRegions() {
	records, err := loadCSV("csv-file/dataregion.csv")
	if err != nil {
		log.Fatal("Error reading region CSV:", err)
	}

	for _, row := range records[1:] { // Skip header
		region := sharedModels.Region{
			Code: row[1],
			Name: row[2],
		}
		config.DB.Create(&region)
	}
	fmt.Println("✅ Regions imported successfully")
}

func importProvinces() {
	records, err := loadCSV("csv-file/dataprovince.csv")
	if err != nil {
		log.Fatal("Error reading province CSV:", err)
	}

	for _, row := range records[1:] {
		province := sharedModels.Province{
			Code:       row[1],
			RegionCode: row[2],
			Name:       row[3],
		}
		config.DB.Create(&province)
	}
	fmt.Println("✅ Provinces imported successfully")
}

func importMunicipalities() {
	records, err := loadCSV("csv-file/datamunicipality.csv")
	if err != nil {
		log.Fatal("Error reading municipality CSV:", err)
	}

	for _, row := range records[1:] {
		municipality := sharedModels.Municipality{
			Code:         row[1],
			ProvinceCode: row[2],
			Name:         row[3],
		}
		config.DB.Create(&municipality)
	}
	fmt.Println("✅ Municipalities imported successfully")
}

func importBarangays() {
	records, err := loadCSV("csv-file/databarangay.csv")
	if err != nil {
		log.Fatal("Error reading barangay CSV:", err)
	}

	for _, row := range records[1:] {
		barangay := sharedModels.Barangay{
			Code:             row[1],
			MunicipalityCode: row[2],
			Name:             row[3],
		}
		config.DB.Create(&barangay)
	}
	fmt.Println("✅ Barangays imported successfully")
}

func InsertCSV() {
	importRegions()
	importProvinces()
	importMunicipalities()
	importBarangays()
}