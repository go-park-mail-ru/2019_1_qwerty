package helpers

import (
	"2019_1_qwerty/database"
	"2019_1_qwerty/models"
)

const sqlInsertUser = `
INSERT INTO users(nickname, email, "password")
VALUES ($1, $2, $3)
`

// DBUserCreate - Создание пользовтеля
func DBUserCreate(user *models.User) error {
	_, err := database.Database.Exec(sqlInsertUser, user.Nickname, user.Email, user.Password)
	if err != nil {
		return models.EUserAE
	}
	return nil
}

// // DBUserUpdate - Обновление данных о пользователе
// func DBUserUpdate(user *models.User) error {
// 	return nil
// }

const sqlSelectUserPasswordByNickname = `
SELECT password
FROM users
WHERE nickname = $1
`

// DBUserValidate - Валидация по nickname\password
func DBUserValidate(user *models.User) error {
	var dbPassw string
	row := database.Database.QueryRow(sqlSelectUserPasswordByNickname, user.Nickname)
	if err := row.Scan(&dbPassw); err != nil {
		return models.EUserNE
	}

	// upass, _ := hash(user.Password)

	if user.Password != dbPassw {
		return models.EWrongPassword
	}
	return nil
}

// const sqlSelectUserByNickname = `
// SELECT nickname, email, avatar
// FROM users
// WHERE nickname = $1
// `

// // DBUserGet - Поиск пользователя по нику
// func DBUserGet(nickname string) (*models.User, error) {
// 	var user models.User
// 	row := database.Database.QueryRow(sqlSelectUserByNickname, nickname)
// 	if err := row.Scan(&user.Nickname, &user.Email, &user.Avatar); err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// // hash - Функция хеширования пароля
// func hash(input string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
// }
