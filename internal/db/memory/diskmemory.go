package memory

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"grpc_course/internal/domain/model"
	"os"
	"sync"
)

type DiskImageStore struct {
	mutex       sync.RWMutex
	imageFolder string
	images      map[string]*ImageInfo
}

func NewDiskImageStore(imageFolder string) model.ImageStore {
	return &DiskImageStore{imageFolder: imageFolder, images: make(map[string]*ImageInfo)}
}

type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string
}

func (d *DiskImageStore) Save(laptopId string, imageType string, imageData bytes.Buffer) (string, error) {
	imageId, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	imagePath := fmt.Sprintf("%s/%s%s", d.imageFolder, imageId, imageType)
	file, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	if _, err := imageData.WriteTo(file); err != nil {
		return "", err
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.images[imageId.String()] = &ImageInfo{
		LaptopID: laptopId,
		Type:     imageType,
		Path:     imagePath,
	}
	return imageId.String(), nil
}

func (d *DiskImageStore) Delete(imageID string) error {
	image := d.images[imageID]
	err := os.Remove(image.Path)
	if err != nil {
		return err
	}
	return nil
}
