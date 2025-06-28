package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250628102225_create_employee_salaries_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250628102225_create_employee_salaries_table_up,
		Down:      mig_20250628102225_create_employee_salaries_table_down,
	})
}

func mig_20250628102225_create_employee_salaries_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE employee_salaries (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_uuid UUID NOT NULL,
		amount FLOAT NOT NULL CHECK (amount > 0),
		active BOOLEAN DEFAULT TRUE,
		effective_date DATE NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`).Error

	return err
}

func mig_20250628102225_create_employee_salaries_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS employee_salaries`).Error
	return err
}
