package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250628102845_create_payslips_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250628102845_create_payslips_table_up,
		Down:      mig_20250628102845_create_payslips_table_down,
	})
}

func mig_20250628102845_create_payslips_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE payslips (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		payroll_uuid UUID NOT NULL,
		user_uuid UUID NOT NULL,
		base_salary FLOAT NOT NULL,
		work_days INT NOT NULL,
		overtime FLOAT NOT NULL,
		reimburse FLOAT NOT NULL,
		total FLOAT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		UNIQUE (payroll_uuid, user_uuid)
	);`).Error

	return err
}

func mig_20250628102845_create_payslips_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS payslips`).Error
	return err
}
