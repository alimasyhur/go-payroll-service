package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250627175541_create_attendance_periods_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250627175541_create_attendance_periods_table_up,
		Down:      mig_20250627175541_create_attendance_periods_table_down,
	})
}

func mig_20250627175541_create_attendance_periods_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE attendance_periods (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		start_date DATE NOT NULL,
		end_date DATE NOT NULL,
		is_closed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`).Error

	return err
}

func mig_20250627175541_create_attendance_periods_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS attendance_periods`).Error
	return err
}
