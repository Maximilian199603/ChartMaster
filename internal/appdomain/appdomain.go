package appdomain

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/EdgeLordKirito/ChartMaster/cmd/chartmaster/appinfo"
)

var ErrUnsupportedOS = errors.New("unsupported operating system")

func DomainPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	var domainPath string
	switch runtime.GOOS {
	case "windows":
		domainPath = filepath.Join(homeDir, "AppData", "Local", appinfo.AppName)
	case "darwin": // macOS
		domainPath = filepath.Join(homeDir, "Library", "Application Support", appinfo.AppName)
	case "linux":
		domainPath = filepath.Join(homeDir, ".config", appinfo.AppName)
	case "freebsd", "openbsd", "netbsd":
		domainPath = filepath.Join(homeDir, ".config", appinfo.AppName)
	default:
		return "", ErrUnsupportedOS
	}
	return domainPath, nil
}
