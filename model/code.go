package model

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

// | USER_NAME   | HASH    | PLAIN_CODE | STDIN    | TITLE       | OPTIONS  | CREATED_AT |
// | ----------- | ------- | ---------- | -------- | ----------- | -------- | --- |
// | VARCHAR(32) | CHAR(8) | TEXT       | TEXT     | VARCHAR(64) | TEXT     | TIMESTAMP |
// | PRI_KEY     | PRI_KEY | NON_NULL   | NON_NULL | NON_NULL    | NON_NULL | 自動挿入されてくれ！ |

type Code struct {
	UserName  string `db:"user_name"`
	PlainCode string `db:"plain_code"`
	Stdin     string `db:"stdin"`
	Title     string `db:"title"`
	Options   string `db:"options"`
}

func makeHash() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 8)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func AddCode(ctx context.Context, code *Code) (hash string, err error) {
	for i := 0; i < 10; i++ {
		h := makeHash()
		c := 0
		err := db.SelectContext(
			ctx,
			&c,
			"SELECT COUNT(*) FROM codes WHERE `user_name` = ? AND `hash` = ?",
			code.UserName,
			h,
		)
		if err != nil {
			return "", err
		}
		if c > 0 {
			continue
		}
		_, err = db.ExecContext(
			ctx,
			"INSERT INTO codes "+
				"(`user_name`, `hash`, `plain_code`, `stdin`, `title`, `options`) "+
				"VALUES (?, ?, ?, ?, ?, ?) WHERE NOT EXIST",
			code.UserName,
			h,
			code.PlainCode,
			code.Stdin,
			code.Title,
			code.Options,
		)
		if err != nil {
			return "", err
		}
		return h, nil
	}
	return "", errors.New("timeout while adding the code")
}

func GetCode(ctx context.Context, userName string, hash string) (Code, error) {
	c := Code{}
	err := db.GetContext(
		ctx,
		&c,
		"SELECT `user_name`, `plain_code`, `stdin`, `title`, `options` "+
			"FROM 22hack16 "+
			"WHERE `user_name` = ? AND `hash` = ?",
		userName,
		hash,
	)
	if err != nil {
		return c, err
	}
	return c, nil
}
