package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250628102814_create_payrolls_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250628102814_create_payrolls_table_up,
		Down:      mig_20250628102814_create_payrolls_table_down,
	})
}

func mig_20250628102814_create_payrolls_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE payrolls (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		attendance_period_uuid UUID NOT NULL,
		processed_at TIMESTAMP NOT NULL,
		ip VARCHAR(64),
		created_by UUID NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`).Error

	return err
}

func mig_20250628102814_create_payrolls_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS payrolls`).Error
	return err
}
