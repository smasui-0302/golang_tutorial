package greetings

import (
	"errors"
	"fmt"
)

// Hello は指定された名前の人に対する挨拶を返します。
func Hello(name string) (string, error) {
	// 挨拶メッセージを作成します。
	// もし名前が入力にない場合エラーメッセージを返す
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("こんにちは、%vさん。ようこそ！", name)
	return message, nil
}
