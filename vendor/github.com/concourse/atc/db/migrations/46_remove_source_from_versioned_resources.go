package migrations

import "github.com/BurntSushi/migration"

func RemoveSourceFromVersionedResources(tx migration.LimitedTx) error {
	_, err := tx.Exec(`
		ALTER TABLE versioned_resources DROP COLUMN source
	`)
	if err != nil {
		return err
	}

	return nil
}
