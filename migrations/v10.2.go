package migrations

func allowLongerPosts(db *datastore) error {
	t, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = t.Exec(`ALTER TABLE posts MODIFY content MEDIUMTEXT`)
	if err != nil {
		t.Rollback()
		return err
	}
	_, err = t.Exec(`ALTER TABLE posts ADD COLUMN content_trunc MEDIUMTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin`)
	if err != nil {
		t.Rollback()
		return err
	}

	err = t.Commit()
	if err != nil {
		t.Rollback()
		return err
	}

	return nil
}

