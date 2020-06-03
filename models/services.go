package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Services struct {
	db      *gorm.DB
	Gallery GalleryService
	User    UserService
	Image   ImageService
}

type ServicesConfig func(*Services) error

// NewServices establishes database connection and creates model services
func NewServices(cfgs ...ServicesConfig) (*Services, error) {
	var s Services

	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
	// db, err := gorm.Open(dialect, connectionInfo)
	// if err != nil {
	// 	return nil, err
	// }
	// db.LogMode(true)

	// return &Services{
	// 	db:      db,
	// 	User:    NewUserService(db),
	// 	Gallery: NewGalleryService(db),
	// 	Image:   NewImageService(),
	// }, nil
}

// Close closes the database connection
func (s *Services) Close() error {
	return s.db.Close()
}

// AutoMigrate automatically migrates all the tables
func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}, &Gallery{}, &pwReset{}).Error
}

// DestructiveReset drops all the tables and rebuilds them
func (s *Services) DestructiveReset() error {
	err := s.db.DropTableIfExists(&User{}, &Gallery{}, &pwReset{}).Error
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}

func WithGorm(dialect, connectionInfo string) ServicesConfig {
	return func(s *Services) error {
		db, err := gorm.Open(dialect, connectionInfo)
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

func WithLogMode(mode bool) ServicesConfig {
	return func(s *Services) error {
		s.db.LogMode(mode)
		return nil
	}
}

func WithUser(pepper, hmacKey string) ServicesConfig {
	return func(s *Services) error {
		s.User = NewUserService(s.db, pepper, hmacKey)
		return nil
	}
}

func WithGallery() ServicesConfig {
	return func(s *Services) error {
		s.Gallery = NewGalleryService(s.db)
		return nil
	}
}

func WithImage() ServicesConfig {
	return func(s *Services) error {
		s.Image = NewImageService()
		return nil
	}
}
