package repository

import (
	"context"
	"database/sql"
)

func BeginTransaction(userRepository *UserRepository) error {
	ctx := context.Background()
	transaction, err := userRepository.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	userRepository.transaction = transaction

	return nil
}

func RollbackTransaction(userRepository *UserRepository) error {
	transaction := userRepository.transaction

	userRepository.transaction = nil

	return transaction.Rollback()
}

func CommitTransaction(userRepository *UserRepository) error {
	transaction := userRepository.transaction

	userRepository.transaction = nil
	return transaction.Commit()
}
