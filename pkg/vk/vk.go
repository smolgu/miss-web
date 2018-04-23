// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vk

import (
	"strconv"

	"github.com/smolgu/miss-web/pkg/setting"

	"github.com/zhuharev/vk"
	"github.com/zhuharev/vkutil"
)

const (
	callbackURL = "http://smlgu.ru/auth/vk/callback"
	clientID    = "4164850"
)

// AuthURL генерирует ссылку, на которую должен перейти пользователь для авторизации через вконтакте
func AuthURL() string {
	// TODO: hardcode
	//uri, _ := vk.GetAuthURL("https://smlgu.ru/auth/vk/callback", "code", strconv.Itoa(setting.App.Vk.AppID), strconv.Itoa(int(vkutil.UPAll)))
	uri, _ := vk.GetAuthURL(callbackURL, "code", clientID, strconv.Itoa(int(vkutil.UPAll)))
	return uri
}

// GetTokenByCode возвращает токен по code (vk code authorization flow)
func GetTokenByCode(code string) (string, error) {
	at, _, err := vk.GetAccessTokenByCode(code, clientID, setting.App.Vk.AppSecret, callbackURL)
	return at, err
}
