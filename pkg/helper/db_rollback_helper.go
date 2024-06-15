package helper

import (
	"gorm.io/gorm"
)

func RollbackOrCommitDb(tx *gorm.DB) {
	err := recover()
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
}
