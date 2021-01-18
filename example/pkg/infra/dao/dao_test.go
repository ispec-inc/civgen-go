package dao

import (
	"os"
	"testing"

	"github.com/ispec-inc/civgen-go/example/pkg/mysql"
	"github.com/tanimutomo/sqlfile"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	err := mysql.Setup()
	if err != nil {
		panic(err)
	}
	db = mysql.DB

	os.Exit(m.Run())
}

func prepareTestData(filepath string) error {
	s := sqlfile.New()
	if err := s.Files("./testdata/delete.sql", filepath); err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if _, err := s.Exec(sqlDB); err != nil {
		return err
	}
	return nil
}
