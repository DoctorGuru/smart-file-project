package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileRepo interface {
	Move(src string, category string) error
}

type fileRepo struct {
	baseDir string
}

func NewFileRepo(baseDir string) FileRepo {
	return &fileRepo{baseDir: baseDir}
}

func (r *fileRepo) Move(src string, category string) error {
	destDir := filepath.Join(r.baseDir, category)
	os.MkdirAll(destDir, os.ModePerm)

	filename := filepath.Base(src)
	destPath := filepath.Join(destDir, filename)

	i := 1
	ext := filepath.Ext(filename)
	nameOnly := filename[:len(filename)-len(ext)]
	for {
		if _, err := os.Stat(destPath); os.IsNotExist(err) {

			break
		}

		destPath = filepath.Join(destDir, fmt.Sprintf("%s_%d%s", nameOnly, i, ext))
		i++
	}

	maxRetries := 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := os.Rename(src, destPath)
		if err == nil {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
		if attempt == maxRetries {
			return fmt.Errorf("failed to move %s after %d attempts: %w", src, maxRetries, err)
		}
	}

	return nil
}
