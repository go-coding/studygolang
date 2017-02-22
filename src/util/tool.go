// Copyright 2013 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author: polaris	polaris@studygolang.com

package util

import (
	"fmt"
	"global"
	"regexp"
	"strings"

	"github.com/polaris1119/goutils"
)

// 获取头像
func Gravatar(avatar string, emailI interface{}, size uint16, isHttps bool) string {
	cdnDomain := global.App.CDNHttp
	gravatarDomain := "http://gravatar.com"
	if isHttps {
		cdnDomain = global.App.CDNHttps
		gravatarDomain = "https://secure.gravatar.com"
	}
	if avatar != "" {
		return fmt.Sprintf("%savatar/%s?imageView2/2/w/%d", cdnDomain, avatar, size)
	}

	email, ok := emailI.(string)
	if !ok {
		return fmt.Sprintf("%savatar/gopher28.png?imageView2/2/w/%d", cdnDomain, size)
	}
	return fmt.Sprintf("%s/avatar/%s?s=%d", gravatarDomain, goutils.Md5(email), size)
}

// 内嵌 Wide iframe 版
func EmbedWide(content string) string {
	if !strings.Contains(content, "&lt;iframe") {
		return content
	}

	reg := regexp.MustCompile(`&lt;iframe.*src=.*(https://wide\.b3log\.org/playground.*\.go).*/iframe&gt;`)
	return reg.ReplaceAllString(content, `<iframe src="$1?embed=true" width="100%" height="600"></iframe>`)
}
