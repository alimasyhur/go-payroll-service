package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250628004552_create_overtimes_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250628004552_create_overtimes_table_up,
		Down:      mig_20250628004552_create_overtimes_table_down,
	})
}

func mig_20250628004552_create_overtimes_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE overtimes (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_uuid UUID NOT NULL REFERENCES users(uuid) ON DELETE CASCADE,
		date DATE NOT NULL,
		hours FLOAT NOT NULL CHECK (hours > 0 AND hours <= 3),
		ip VARCHAR(64),
		created_by UUID NOT NULL REFERENCES users(uuid),
		updated_by UUID NOT NULL REFERENCES users(uuid),
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`).Error

	return err
}

func mig_20250628004552_create_overtimes_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS overtimes`).Error
	return err
}
