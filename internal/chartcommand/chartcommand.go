package chartcommand

import (
	"database/sql"

	"github.com/EdgeLordKirito/ChartMaster/internal/chartvalidation"
	"github.com/EdgeLordKirito/ChartMaster/internal/dbutils"
	"github.com/spf13/cobra"
)

func AddCommand(cmd *cobra.Command, args []string) error {
	// Adds an new chart under the specified name with the content of the specified file

	chartName := args[0]
	filePath := args[1]

	data, err := chartvalidation.ReadAndValidate(filePath)
	if err != nil {
		// error return for reading error or validation error
		return err
	}

	// error is of type DatabasePathError
	dbPath, err := dbutils.GetDatabasePath()
	if err != nil {
		return err
	}
	db, err := dbutils.OpenDatabase(dbPath)
	if err != nil {
		return &dbutils.DBConnectionError{Internal: err}
	}
	defer db.Close()

	// Prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO Charts (name, data) VALUES (?, ?)")
	if err != nil {
		return &dbutils.SQLPreparationError{Internal: err}
	}
	defer stmt.Close()

	// Execute the insert statement
	_, err = stmt.Exec(chartName, data)
	if err != nil {
		// Check for unique constraint violation
		if err.Error() == "UNIQUE constraint failed: Charts.name" {
			return dbutils.ErrChartExists
		}
		// wrap unexpected error
		return &dbutils.SQLExecutionError{Internal: err}
	}

	// Chart added successfully
	return nil
}

func UpdateCommand(cmd *cobra.Command, args []string) error {
	// Updates an specified chart with the content of the specified file
	chartName := args[0]
	filePath := args[1]

	data, err := chartvalidation.ReadAndValidate(filePath)
	if err != nil {
		return err
	}

	// error is of type DatabasePathError
	dbPath, err := dbutils.GetDatabasePath()
	if err != nil {
		return err
	}
	db, err := dbutils.OpenDatabase(dbPath)
	if err != nil {
		return &dbutils.DBConnectionError{Internal: err}
	}
	defer db.Close()

	// Prepare the update statement
	stmt, err := db.Prepare("UPDATE Charts SET data = ? WHERE name = ?")
	if err != nil {
		return &dbutils.SQLPreparationError{Internal: err}
	}
	defer stmt.Close()

	// Execute the update statement
	_, err = stmt.Exec(data, chartName)
	if err != nil {
		// Wrap unexpected error
		return &dbutils.SQLExecutionError{Internal: err}
	}

	// Chart updated successfully
	return nil
}

func ReadCommand(cmd *cobra.Command, args []string) error {
	// Outputs the saved csv data for the specified chart
	chartName := args[0]

	// error is of type DatabasePathError
	dbPath, err := dbutils.GetDatabasePath()
	if err != nil {
		return err
	}
	db, err := dbutils.OpenDatabase(dbPath)
	if err != nil {
		return &dbutils.DBConnectionError{Internal: err}
	}
	defer db.Close()

	// Prepare the select statement
	stmt, err := db.Prepare("SELECT data FROM Charts WHERE name = ?")
	if err != nil {
		return &dbutils.SQLPreparationError{Internal: err}
	}
	defer stmt.Close()

	// Execute the select statement
	var data []byte
	err = stmt.QueryRow(chartName).Scan(&data)
	if err != nil {
		// Check for SQL not found error
		if err == sql.ErrNoRows {
			return dbutils.ErrChartNotExists
		}
		// Wrap unexpected error
		return &dbutils.SQLExecutionError{Internal: err}
	}

	// Save the retrieved BLOB data as needed (you can handle it here)
	// Example:
	// someGlobalVariable = data

	return nil
}

func RemoveCommand(cmd *cobra.Command, args []string) error {
	// Removes an specifed chart
	chartName := args[0]

	// error is of type DatabasePathError
	dbPath, err := dbutils.GetDatabasePath()
	if err != nil {
		return err
	}
	db, err := dbutils.OpenDatabase(dbPath)
	if err != nil {
		return &dbutils.DBConnectionError{Internal: err}
	}
	defer db.Close()

	// Prepare the delete statement
	stmt, err := db.Prepare("DELETE FROM Charts WHERE name = ?")
	if err != nil {
		return &dbutils.SQLPreparationError{Internal: err}
	}
	defer stmt.Close()

	// Execute the delete statement
	_, err = stmt.Exec(chartName)
	if err != nil {
		// Wrap unexpected error
		return &dbutils.SQLExecutionError{Internal: err}
	}

	// Chart removed successfully
	return nil
}

func ListCommand(cmd *cobra.Command, args []string) error {
	// lists all saved charts
	// error is of type DatabasePathError
	dbPath, err := dbutils.GetDatabasePath()
	if err != nil {
		return err
	}
	db, err := dbutils.OpenDatabase(dbPath)
	if err != nil {
		return &dbutils.DBConnectionError{Internal: err}
	}
	defer db.Close()

	// Prepare the select statement to retrieve all chart names
	stmt, err := db.Prepare("SELECT name FROM Charts")
	if err != nil {
		return &dbutils.SQLPreparationError{Internal: err}
	}
	defer stmt.Close()

	// Execute the select statement
	rows, err := stmt.Query()
	if err != nil {
		return &dbutils.SQLExecutionError{Internal: err}
	}
	defer rows.Close()

	// Initialize a slice to hold chart names
	var chartNames []string

	// Iterate through the rows and scan the chart names
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return &dbutils.SQLExecutionError{Internal: err}
		}
		chartNames = append(chartNames, name)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return &dbutils.SQLExecutionError{Internal: err}
	}

	// Save or use the chartNames slice as needed
	// Example:
	// someGlobalVariable = chartNames

	return nil
}
