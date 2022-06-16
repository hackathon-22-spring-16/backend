package model

import (
	"context"
)

// | USER_NAME   | HASH    | PLAIN_CODE | STDIN    | TITLE       | OPTIONS  | CREATED_AT |
// | ----------- | ------- | ---------- | -------- | ----------- | -------- | --- |
// | VARCHAR(32) | CHAR(8) | TEXT       | TEXT     | VARCHAR(64) | TEXT     | TIMESTAMP |
// | PRI_KEY     | PRI_KEY | NON_NULL   | NON_NULL | NON_NULL    | NON_NULL | 自動挿入されてくれ！ |

type Code struct {
	UserName  string
	Hash      string
	PlainCode string
	Stdin     string
	Title     string
	Options   string
}

func AddCode(ctx context.Context, code *Code) error {
	return nil
}

func GetCode(ctx context.Context, userName string, hash string) (*Code, error) {
	return nil, nil
}
