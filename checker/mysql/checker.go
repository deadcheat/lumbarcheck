package mysql

import (
	"database/sql"

	"github.com/deadcheat/lumbarcheck"

	"github.com/deadcheat/lumbarcheck/constants"
	_ "github.com/go-sql-driver/mysql"
)

// Checker implements checker interface
type Checker struct{}

// New create new Checker
func New() lumbarcheck.Checker {
	return &Checker{}
}

// Check implement func
func (c *Checker) Check(dataSource string) (err error) {
	db, _ := sql.Open(constants.MySQL, dataSource)

	defer db.Close()

	if _, err = db.Query("SELECT 1 as value"); err != nil {
		return err
	}

	return nil
}
