// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import macaron "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()
	m.Use(Renderer())
	m.Get("/", func(ctx *macaron.Context) {

		ctx.HTML(200, "index")
	})
	m.Get("/auth/vk", func() string {
		return "рано"
	})
	m.Run()
}

// Renderer sets up template renderer
func Renderer() macaron.Handler {
	opts := macaron.RenderOptions{
		Layout: "layout",
	}
	return macaron.Renderer(opts)
}
