package usecase

import (
	"database/sql"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
)

type UserUseCase struct {
	DB *sql.DB
}

func (uc *UserUseCase) CreateUser(user entity.UserRequest) error {
	insertQuery := `INSERT INTO users (username, email) VALUES (?, ?)`

	_, err := uc.DB.Exec(insertQuery, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) GetUser(id int) (*entity.User, error) {
	var (
		selectQuery = `SELECT id, username, email, deletedAt FROM users WHERE id = ?`
		user        = entity.User{}
	)

	err := uc.DB.QueryRow(selectQuery, id).Scan(&user.ID, &user.Name, &user.Email, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uc *UserUseCase) UpdateUser(user entity.UserRequest) (*entity.User, error) {
	updateQuery := `UPDATE users SET username = ?, email = ? WHERE id = ?`

	_, err := uc.DB.Exec(updateQuery, user.Name, user.Email, user.ID)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (uc *UserUseCase) DeleteUser(id int) (*entity.User, error) {
	deleteQuery := `UPDATE users SET deletedAt = UTC_TIMESTAMP() WHERE id = ?`
	_, err := uc.DB.Exec(deleteQuery, id)
	if err != nil {
		return nil, err
	}
	return &entity.User{ID: id}, nil
}
