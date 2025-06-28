package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250628102454_create_reimbursements_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250628102454_create_reimbursements_table_up,
		Down:      mig_20250628102454_create_reimbursements_table_down,
	})
}

func mig_20250628102454_create_reimbursements_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE reimbursements (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_uuid UUID NOT NULL,
		amount FLOAT NOT NULL CHECK (amount > 0),
		description TEXT,
		date DATE NOT NULL,
		created_by UUID NOT NULL,
		updated_by UUID NOT NULL,
		ip VARCHAR(64),
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`).Error

	return err
}

func mig_20250628102454_create_reimbursements_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS reimbursements`).Error
	return err
}
