package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250628164627_create_audit_logs_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250628164627_create_audit_logs_table_up,
		Down:      mig_20250628164627_create_audit_logs_table_down,
	})
}

func mig_20250628164627_create_audit_logs_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE audit_logs (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_uuid UUID NOT NULL,
		action VARCHAR(50) NOT NULL,
		entity VARCHAR(100) NOT NULL,
		entity_uuid UUID NOT NULL,
		ip VARCHAR(64),
		request_id VARCHAR(64),
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`).Error

	return err
}

func mig_20250628164627_create_audit_logs_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS audit_logs`).Error
	return err
}
