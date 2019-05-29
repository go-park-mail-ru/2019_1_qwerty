package helpers

import (
	"2019_1_qwerty/database"
	"2019_1_qwerty/models"
	"log"
)

const sqlInsertUser = `
INSERT INTO users(nickname, "password")
VALUES ($1, $2)
`

// DBUserCreate - Создание пользовтеля
func DBUserCreate(user *models.User) error {
	_, err := database.Database.Exec(sqlInsertUser, user.Nickname, user.Password)
	if err != nil {
		log.Println(err)
		return models.EUserAE
	}
	return nil
}

const sqlUpdateUserByNickname = `
UPDATE users
SET password = COALESCE(NULLIF($2, ''), password)
WHERE nickname = $1
`

// DBUserUpdate - Обновление данных о пользователе
func DBUserUpdate(nickname string, user *models.User) error {
	log.Println("in update:", nickname, user.Password)
	_, _ = database.Database.Exec(sqlUpdateUserByNickname, nickname, user.Password)
	return nil
}

const sqlUpdateUserAvatarByNickname = `
UPDATE users
SET avatar = $2
WHERE nickname = $1
`

// DBUserUpdateAvatar - Обновление данных о пользователе
func DBUserUpdateAvatar(nickname string, avatar string) error {
	_, _ = database.Database.Exec(sqlUpdateUserAvatarByNickname, nickname, avatar)
	return nil
}

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
		log.Println(err)
		return models.EUserNE
	}

	// upass, _ := hash(user.Password)

	if user.Password != dbPassw {
		return models.EWrongPassword
	}
	return nil
}

const sqlSelectUserByNickname = `
SELECT nickname, avatar
FROM users
WHERE nickname = $1
`

// DBUserGet - Поиск пользователя по нику
func DBUserGet(nickname string) (*models.User, error) {
	user := models.User{}
	row := database.Database.QueryRow(sqlSelectUserByNickname, nickname)
	if err := row.Scan(&user.Nickname, &user.Avatar); err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

const sqlSelectScoreByNickname = `
	SELECT score FROM scores
	WHERE player = $1
	`

// DBUserGetScore - Get user score by nickname
func DBUserGetScore(nickname string) (uint64, error) {
	var score uint64
	row := database.Database.QueryRow(sqlSelectScoreByNickname, nickname)
	if err := row.Scan(&score); err != nil {
		log.Println(err)
		return 0, err
	}
	return score, nil
}

// // hash - Функция хеширования пароля
// func hash(input string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
// }
