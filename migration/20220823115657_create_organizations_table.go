package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20220823115657_create_organizations_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20220823115657_create_organizations_table_up,
		Down:      mig_20220823115657_create_organizations_table_down,
	})
}

func mig_20220823115657_create_organizations_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS organizations (
		id int(11) unsigned NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		location text NOT NULL,
		created_at timestamp NOT NULL DEFAULT current_timestamp(),
		updated_at timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
		PRIMARY KEY (id)
	)`).Error

	return err
}

func mig_20220823115657_create_organizations_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS organizations`).Error
	return err
}
