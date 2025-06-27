package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20250627100233_create_users_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250627100233_create_users_table_up,
		Down:      mig_20250627100233_create_users_table_down,
	})
}

func mig_20250627100233_create_users_table_up(tx *gorm.DB) error {
	err := tx.Exec(`
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
			CREATE TYPE user_role AS ENUM ('admin', 'employee');
		END IF;
	END $$;

	CREATE TABLE users (
		uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		username VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		role user_role NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`).Error

	return err
}

func mig_20250627100233_create_users_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS users`).Error
	return err
}
