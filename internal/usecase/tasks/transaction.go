package tasks

import (
	"context"
	"database/sql"
)

func BeginTransaction(repository *TaskRepository) error {
	ctx := context.Background()
	transaction, err := repository.DbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	repository.transaction = transaction

	return nil
}

func RollbackTransaction(repository *TaskRepository) error {
	transaction := repository.transaction

	repository.transaction = nil

	return transaction.Rollback()
}

func CommitTransaction(repository *TaskRepository) error {
	transaction := repository.transaction

	repository.transaction = nil
	return transaction.Commit()
}
