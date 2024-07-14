package user

import (
	"context"
	"database/sql"
)

func BeginTransaction(repository *UserRepository) error {
	ctx := context.Background()
	transaction, err := repository.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	repository.transaction = transaction

	return nil
}

func RollbackTransaction(repository *UserRepository) error {
	transaction := repository.transaction

	repository.transaction = nil

	return transaction.Rollback()
}

func CommitTransaction(repository *UserRepository) error {
	transaction := repository.transaction

	repository.transaction = nil
	return transaction.Commit()
}
