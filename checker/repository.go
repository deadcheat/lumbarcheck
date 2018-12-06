package checker

import (
	"fmt"

	"github.com/deadcheat/lumbarcheck"
	"github.com/deadcheat/lumbarcheck/checker/mysql"
	"github.com/deadcheat/lumbarcheck/constants"
)

// Create returns Checker interface implementation
func Create(database string) (lumbarcheck.Checker, error) {
	switch database {
	case constants.MySQL:
		return mysql.New(), nil
	}

	return nil, fmt.Errorf("undefined or checker unimplemented for database: %s", database)
}
