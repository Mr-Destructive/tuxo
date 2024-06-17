package ssg

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func StartServer(port int, baseFolder string) error {
	// Resolve the absolute path of the base folder
	absBaseFolder, err := filepath.Abs(baseFolder)
	if err != nil {
		return err
	}

	// Check if the base folder exists
	_, err = os.Stat(absBaseFolder)
	if err != nil {
		return fmt.Errorf("base folder '%s' does not exist", absBaseFolder)
	}

	// Create a file server handler rooted at the base folder
	fileServer := http.FileServer(http.Dir(absBaseFolder))

	// Register the file server handler and start the HTTP server
	http.Handle("/", http.StripPrefix("/", fileServer))
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server on port %d...\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}
