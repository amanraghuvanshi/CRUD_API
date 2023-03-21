package helper

import "database/sql"

func CommitorRollback(tx *sql.Tx) {

	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
