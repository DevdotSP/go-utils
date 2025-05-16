package importingcsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/DevdotSP/go-utils/config"
	sharedModels "github.com/DevdotSP/go-utils/shared-models"
)

const batchSize = 500

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

	var regions []sharedModels.Region
	for _, row := range records[1:] { // Skip header
		if len(row) < 3 {
			log.Println("Skipping invalid region row:", row)
			continue
		}
		regions = append(regions, sharedModels.Region{
			Code: row[1],
			Name: row[2],
		})
	}

	if err := config.DB.CreateInBatches(&regions, batchSize).Error; err != nil {
		log.Fatal("Error inserting regions:", err)
	}

	fmt.Println("✅ Regions imported successfully")
}

func importProvinces() {
	records, err := loadCSV("csv-file/dataprovince.csv")
	if err != nil {
		log.Fatal("Error reading province CSV:", err)
	}

	var provinces []sharedModels.Province
	for _, row := range records[1:] {
		if len(row) < 4 {
			log.Println("Skipping invalid province row:", row)
			continue
		}
		provinces = append(provinces, sharedModels.Province{
			Code:       row[1],
			RegionCode: row[2],
			Name:       row[3],
		})
	}

	if err := config.DB.CreateInBatches(&provinces, batchSize).Error; err != nil {
		log.Fatal("Error inserting provinces:", err)
	}

	fmt.Println("✅ Provinces imported successfully")
}

func importMunicipalities() {
	records, err := loadCSV("csv-file/datamunicipality.csv")
	if err != nil {
		log.Fatal("Error reading municipality CSV:", err)
	}

	var municipalities []sharedModels.Municipality
	for _, row := range records[1:] {
		if len(row) < 4 {
			log.Println("Skipping invalid municipality row:", row)
			continue
		}
		municipalities = append(municipalities, sharedModels.Municipality{
			Code:         row[1],
			ProvinceCode: row[2],
			Name:         row[3],
		})
	}

	if err := config.DB.CreateInBatches(&municipalities, batchSize).Error; err != nil {
		log.Fatal("Error inserting municipalities:", err)
	}

	fmt.Println("✅ Municipalities imported successfully")
}

func importBarangays() {
	records, err := loadCSV("csv-file/databarangay.csv")
	if err != nil {
		log.Fatal("Error reading barangay CSV:", err)
	}

	var barangays []sharedModels.Barangay
	for _, row := range records[1:] {
		if len(row) < 4 {
			log.Println("Skipping invalid barangay row:", row)
			continue
		}
		barangays = append(barangays, sharedModels.Barangay{
			Code:             row[1],
			MunicipalityCode: row[2],
			Name:             row[3],
		})
	}

	if err := config.DB.CreateInBatches(&barangays, batchSize).Error; err != nil {
		log.Fatal("Error inserting barangays:", err)
	}

	fmt.Println("✅ Barangays imported successfully")
}

func InsertCSV() {
	importRegions()
	importProvinces()
	importMunicipalities()
	importBarangays()
}
