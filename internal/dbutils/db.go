package dbutils

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/EdgeLordKirito/ChartMaster/cmd/chartmaster/appinfo"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var ErrChartExists = errors.New("chart already exists in the database")
var ErrChartNotExists = errors.New("chart does not exist in the database")

type SQLPreparationError struct {
	Internal error
}

func (e *SQLPreparationError) Error() string {
	return fmt.Sprintf("SQL preparation: %v", e.Internal)
}

type SQLExecutionError struct {
	Internal error
}

func (e *SQLExecutionError) Error() string {
	return fmt.Sprintf("SQL execution: %v", e.Internal)
}

type DBConnectionError struct {
	Internal error
}

func (e *DBConnectionError) Error() string {
	return fmt.Sprintf("Database connection: %v", e.Internal)
}

// OpenDatabase opens the SQLite database located at the provided path.
func OpenDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDatabasePath() (string, error) {
	var dbPath string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", &DBConnectionError{Internal: err}
	}

	switch runtime.GOOS {
	case "windows":
		dbPath = filepath.Join(homeDir, "AppData", "Local", appinfo.AppName, "database.db")
	case "darwin": // macOS
		dbPath = filepath.Join(homeDir, "Library", "Application Support", appinfo.AppName, "database.db")
	case "linux":
		dbPath = filepath.Join(homeDir, ".config", appinfo.AppName, "database.db")
	case "freebsd", "openbsd", "netbsd":
		dbPath = filepath.Join(homeDir, ".config", appinfo.AppName, "database.db")
	default:
		return "", &DBConnectionError{Internal: fmt.Errorf("unsupported operating system")}
	}

	// Check if the file exists
	if _, err := os.Stat(dbPath); err != nil {
		return "", &DBConnectionError{Internal: err}
	}

	return dbPath, nil
}
