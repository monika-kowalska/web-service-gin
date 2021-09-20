package test_data

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/monika-kowalska/web-service-gin/config"
	"github.com/monika-kowalska/web-service-gin/models"
)

// Initializes SQLite database used for testing
func init() {
	config.ConnectDataBase(":memory:")
	config.DB.Exec("PRAGMA foreign_keys = ON")

	// database, err := gorm.Open("sqlite3", ":memory:")

}

func ResetDB() *gorm.DB {
	config.DB.DropTableIfExists(&models.Campaign{}) // Note: Order matters
	config.DB.AutoMigrate(&models.Campaign{})
	// fmt.Println("dupa dupa dupa")
	// fmt.Println(filepath.Abs("/db.sql"))
	if err := runSQLFile(config.DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("error while initializing test database: %s", err))
	}
	return config.DB
}

func getSQLFile() string {
	return "/Users/monikakowalska/projects/my_projects/playground/go/web-service-gin/test_data/db.sql"
}

func GetTestCaseFolder() string {
	return "/Users/monikakowalska/projects/my_projects/playground/go/web-service-gin/test_data/test_case_data"
}

// Executes SQL file specified by file argument
func runSQLFile(db *gorm.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if result := db.Exec(line); result.Error != nil {
			fmt.Println(line)
			return result.Error
		}
	}
	return nil
}
