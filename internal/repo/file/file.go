package file

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fd239/gopher_keeper/internal/models"
	uuid "github.com/satori/go.uuid"
	"os"
	"sync"
)

type DiskImageStore struct {
	mutex  sync.RWMutex
	folder string
	files  map[string]*models.DataFile
}

func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		folder: imageFolder,
		files:  make(map[string]*models.DataFile),
	}
}

// Save file save
func (store *DiskImageStore) Save(_ context.Context, fileType string, fileData bytes.Buffer) (string, error) {
	fileId := uuid.NewV4()
	filePath := fmt.Sprintf("%s/%s%s", store.folder, fileId, fileType)

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = fileData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.files[fileId.String()] = &models.DataFile{
		FileId: fileId,
		Type:   models.TypeFile,
		Path:   filePath,
	}

	return fileId.String(), nil
}
