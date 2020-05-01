package models

import "github.com/jinzhu/gorm"

type Services struct {
	db      *gorm.DB
	Gallery GalleryService
	User    UserService
}

// NewServices establishes database connection and creates model services
func NewServices(connectionInfo string) (*Services, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Services{
		db:      db,
		User:    NewUserService(db),
		Gallery: &galleryGorm{},
	}, nil
}

// Close closes the database connection
func (s *Services) Close() error {
	return s.db.Close()
}

// AutoMigrate automatically migrates all the tables
func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}, &Gallery{}).Error
}

// DestructiveReset drops all the tables and rebuilds them
func (s *Services) DestructiveReset() error {
	err := s.db.DropTableIfExists(&User{}, &Gallery{}).Error
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}
