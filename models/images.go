package models

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

type ImageService interface {
	Create(galleryID uint, r io.Reader, filename string) error
	ByGalleryID(galleryID uint) ([]Image, error)
	Delete(i *Image) error
}

type imageService struct{}

type Image struct {
	GalleryID uint
	Filename  string
}

func NewImageService() ImageService {
	return &imageService{}
}

func (is *imageService) Create(galleryID uint, r io.Reader, filename string) error {
	path, err := is.mkImagePath(galleryID)
	if err != nil {
		return err
	}

	dst, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}
	return nil
}

func (is *imageService) ByGalleryID(galleryID uint) ([]Image, error) {
	path := is.imagePath(galleryID)
	strings, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return nil, err
	}
	ret := make([]Image, len(strings))
	for i, imgStr := range strings {
		ret[i] = Image{
			Filename:  filepath.Base(imgStr),
			GalleryID: galleryID,
		}
	}
	return ret, nil
}

func (is *imageService) Delete(i *Image) error {
	return os.Remove(i.RelativePath())
}

func (is *imageService) imagePath(galleryID uint) string {
	return filepath.Join("images", "galleries", fmt.Sprintf("%v", galleryID))
}

func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := filepath.Join("images", "galleries",
		fmt.Sprintf("%v", galleryID))
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}

func (i *Image) Path() string {
	tmp := url.URL{
		Path: "/" + i.RelativePath(),
	}
	return tmp.String()
}

func (i *Image) RelativePath() string {
	galleryID := fmt.Sprintf("%v", i.GalleryID)
	return filepath.ToSlash(filepath.Join("images", "galleries", galleryID, i.Filename))
}
