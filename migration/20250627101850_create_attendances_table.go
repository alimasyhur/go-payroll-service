package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250627101850_create_attendances_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250627101850_create_attendances_table_up,
		Down:      mig_20250627101850_create_attendances_table_down,
	})
}

func mig_20250627101850_create_attendances_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE attendances (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_uuid UUID NOT NULL,
		date DATE NOT NULL,
		clockin TIME,
		clockout TIME,
		ip VARCHAR(64),
		created_by UUID NOT NULL,
		updated_by UUID NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		UNIQUE(user_uuid, date)
	);`).Error

	return err
}

func mig_20250627101850_create_attendances_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS attendances`).Error
	return err
}
