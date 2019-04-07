package helpers

import (
	"2019_1_qwerty/database"
	"2019_1_qwerty/models"
)

const sqlInsertUser = `
INSERT INTO users(nickname, email, hashedPassword, avatar)
VALUES ($1, $2, $3, $4)
`

func DBUserCreate(user *models.User) (*models.Users, error) {
	_, err := database.Database.Exec(sqlInsertUser, user.Nickname, user.Email, user.Password, user.Avatar)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

const sqlSelectUserByNickname = `
SELECT nickname, email, avatar
FROM users
WHERE nickname = $1
`

func DBUserGetInfoByNickname(nickname string) (*models.User, error) {
	user models.User
	row := database.Database.QueryRow(sqlSelectUserByNickname, nickname)
	if err := row.Scan(&user.Nickname, &user.Email, &user.Avatar); err != nil {
		return nil, err
	}
}
