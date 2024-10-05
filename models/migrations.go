// models/migrations.go
package models

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&Product{}, &User{}, &CartItem{}); err != nil {
		return err
	}
	return nil
}
