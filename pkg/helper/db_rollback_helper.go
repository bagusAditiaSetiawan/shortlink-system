package helper

import "gorm.io/gorm"

func RollbackOrCommitDb(tx *gorm.DB) {
	error := recover()
	if error != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
