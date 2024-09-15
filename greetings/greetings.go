package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello は指定された名前の人に対する挨拶を返します。
func Hello(name string) (string, error) {
	// 挨拶メッセージを作成します。
	// もし名前が入力にない場合エラーメッセージを返す
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat はランダムに選択された挨拶メッセージを返します。
func randomFormat() string {
	// 挨拶メッセージの配列を定義します。
	formats := []string{
		"こんにちは、%v さん！",
		"いらっしゃいませ、%v さん！",
		"%v さん、お会いできて嬉しいです！",
		"%v さん、ようこそ！",
	}

	// ランダムに選択されたメッセージを返します。
	return formats[rand.Intn(len(formats))]
}
