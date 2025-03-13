package usecase

import (
	"database/sql"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
)

type UserUseCase struct {
	DB *sql.DB
}

func (uc *UserUseCase) CreateUser(user entity.UserRequest) error {
	insertQuery := `INSERT INTO users (name, email) VALUES ($1, $2)`

	_, err := uc.DB.Exec(insertQuery, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) GetUser(id string) (*entity.User, error) {
	var (
		selectQuery = `SELECT name, email FROM users WHERE userid = $1`
		user        = entity.User{}
	)

	err := uc.DB.QueryRow(selectQuery, id).Scan(&user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uc *UserUseCase) UpdateUser(user entity.UserRequest) error {
	updateQuery := `UPDATE users SET name = $1, email = $2 WHERE userid = $3`

	_, err := uc.DB.Exec(updateQuery, user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) DeleteUser(id string) error {
	deleteQuery := `UPDATE users SET deleted_at = UTC_TIMESTAMP() WHERE userid = $1`
	_, err := uc.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}
