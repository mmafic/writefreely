package migrations

func addRedditUrlColumnToPosts(db *datastore) error {
	t, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = t.Exec(`ALTER TABLE posts ADD COLUMN reddit_url ` + db.typeVarChar(255) + ` NULL`)
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

