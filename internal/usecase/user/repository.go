package user

import (
	"TimeTracker/internal/entity"
	"database/sql"
	"net/http"
)

type UserRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserRepository(dbHandler *sql.DB) *UserRepository {
	return &UserRepository{
		dbHandler: dbHandler,
	}
}

func (ur UserRepository) AddUser(user *entity.Users) (*entity.Users, *entity.ResponseError) {
	query := `
				WITH ids AS(
    			INSERT INTO users(name, surname, patronymic, address)
    			VALUES($1, $2, $3, $4)
    			RETURNING id
           		)
				INSERT INTO user_documents(user_id, passport_number, passport_serial_number)
				SELECT id, $5, $6 FROM ids;`
	rows, err := ur.dbHandler.Query(query, user.Name, user.Surname, user.Patronymic, user.Address, user.PassportNumber, user.PassportSerial)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var userId string
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &entity.Users{
		ID:             userId,
		Name:           user.Name,
		Surname:        user.Surname,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
		PassportSerial: user.PassportSerial,
		PassportNumber: user.PassportNumber,
	}, nil
}

func (ur UserRepository) UpdateUser(user *entity.Users) *entity.ResponseError {
	query := `
		WITH upd AS(
    	UPDATE users SET name=$1, surname=$2, patronymic=$3, address=$4 WHERE id=$5
		) UPDATE user_documents
    		SET passport_serial_number=$6, passport_number=$7 WHERE user_id=$5;`
	rows, err := ur.dbHandler.Exec(
		query, user.Name, user.Surname, user.Patronymic, user.Address, user.ID, user.PassportNumber, user.PassportSerial)
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if rowsAffected == 0 {
		return &entity.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}
	return nil
}

func (ur UserRepository) DeleteUser(userId string) *entity.ResponseError {
	query := `
		DELETE FROM users
		WHERE id=$1
		RETURNING id`
	rows, err := ur.transaction.Query(
		query, userId)
	defer rows.Close()
	if err != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			return &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (ur UserRepository) GetUser(userId string) (*entity.Users, *entity.ResponseError) {
	query := `
			SELECT u.id, u.name, u.surname, u.patronymic, u.address, ud.passport_serial_number, ud.passport_number
			FROM users AS u
        	JOIN user_documents AS ud ON u.id = ud.user_id
			WHERE u.id = $1;`
	rows, err := ur.dbHandler.Query(query, userId)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var id, name, surname, patronimyc, address, passport_serilal_number, passport_number string
	for rows.Next() {
		err := rows.Scan(&id, &name, &surname, &patronimyc, &address, &passport_serilal_number, &passport_number)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &entity.Users{
		ID:             id,
		Name:           name,
		Surname:        surname,
		Patronymic:     patronimyc,
		Address:        address,
		PassportSerial: passport_serilal_number,
		PassportNumber: passport_number,
	}, nil
}

func (ur UserRepository) GetAllUsers(limit int, offset int) ([]*entity.Users, *entity.ResponseError) {
	query := `
			SELECT u.id, u.name, u.surname, u.patronymic, u.address, ud.passport_serial_number, ud.passport_number
			FROM users AS u
			JOIN user_documents AS ud ON u.id = ud.user_id
			ORDER BY u.id
			LIMIT $1 OFFSET $2;
`
	rows, err := ur.dbHandler.Query(query, limit, offset)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()
	users := make([]*entity.Users, 0)
	var id, name, surname, patronimyc, address, passport_serilal_number, passport_number string
	for rows.Next() {
		err := rows.Scan(&id, &name, &surname, &patronimyc, &address, &passport_serilal_number, &passport_number)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		user := &entity.Users{
			ID:             id,
			Name:           name,
			Surname:        surname,
			Patronymic:     patronimyc,
			Address:        address,
			PassportSerial: passport_serilal_number,
			PassportNumber: passport_number,
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return users, nil
}

// TODO тут жестко, нужна норм фильтрация
func (ur UserRepository) GetUserByParams(limit int, offset int, filterValue string) ([]*entity.Users, *entity.ResponseError) {
	query := `
			SELECT u.id, u.name, u.surname, u.patronymic, u.address, ud.passport_serial_number, ud.passport_number
			FROM users AS u
			JOIN user_documents AS ud ON u.id = ud.user_id
			ORDER BY u.id
			LIMIT $1 OFFSET $2;
`
	rows, err := ur.dbHandler.Query(query, limit, offset)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	users := make([]*entity.Users, 0)
	var id, name, surname, patronimyc, address, passport_serilal_number, passport_number string
	for rows.Next() {
		err := rows.Scan(&id, &name, &surname, &patronimyc, &address, &passport_serilal_number, &passport_number)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		user := &entity.Users{
			ID:             id,
			Name:           name,
			Surname:        surname,
			Patronymic:     patronimyc,
			Address:        address,
			PassportSerial: passport_serilal_number,
			PassportNumber: passport_number,
		}
		if user.Name == filterValue || user.Surname == filterValue ||
			user.Patronymic == filterValue || user.Address == filterValue ||
			user.PassportSerial == filterValue || user.PassportNumber == filterValue {
			users = append(users, user)
		}

	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return users, nil
}

func (ur UserRepository) Info(passportSerial string, passportNumber string) (*entity.Users, *entity.ResponseError) {
	query := `
			SELECT u.id, u.name, u.surname, u.patronymic, u.address, ud.passport_serial_number, ud.passport_number
			FROM users AS u
         	JOIN user_documents AS ud ON u.id = ud.user_id
			WHERE ud.passport_serial_number = $1 AND ud.passport_number = $2;`
	rows, err := ur.dbHandler.Query(query, passportSerial, passportNumber)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var id, name, surname, patronimyc, address, passport_serilal_number, passport_number string
	for rows.Next() {
		err := rows.Scan(&id, &name, &surname, &patronimyc, &address, &passport_serilal_number, &passport_number)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &entity.Users{
		ID:             id,
		Name:           name,
		Surname:        surname,
		Patronymic:     patronimyc,
		Address:        address,
		PassportSerial: passport_serilal_number,
		PassportNumber: passport_number,
	}, nil
}

func (ur UserRepository) AddUserApi(user *entity.Users) (*entity.Users, *entity.ResponseError) {
	query := `
				WITH ids AS(
    			INSERT INTO users(name, surname, patronymic, address)
    			VALUES($1, $2, $3, $4)
    			RETURNING id
           		)
				INSERT INTO user_documents(user_id, passport_number, passport_serial_number)
				SELECT id, $5, $6 FROM ids;`
	rows, err := ur.dbHandler.Query(query, user.Name, user.Surname, user.Patronymic, user.Address, user.PassportNumber, user.PassportSerial)
	if err != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer rows.Close()
	var userId string
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			return nil, &entity.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}
	if rows.Err() != nil {
		return nil, &entity.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &entity.Users{
		ID:             userId,
		Name:           user.Name,
		Surname:        user.Surname,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
		PassportSerial: user.PassportSerial,
		PassportNumber: user.PassportNumber,
	}, nil
}
