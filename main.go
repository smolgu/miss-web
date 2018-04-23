// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"
	"os"

	"github.com/smolgu/miss-web/pkg/client"
	"github.com/smolgu/miss-web/pkg/setting"
	"github.com/smolgu/miss-web/pkg/vk"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
	macaron "gopkg.in/macaron.v1"
)

func main() {

	app := &cli.App{
		Action: run,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "dev",
			},
			cli.StringFlag{
				Name:  "config",
				Value: "conf/app.yaml",
			},
		},
	}

	app.Run(os.Args)
}

// Renderer sets up template renderer
func Renderer() macaron.Handler {
	opts := macaron.RenderOptions{
		Layout: "layout",
	}
	return macaron.Renderer(opts)
}

func run(ctx *cli.Context) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	setting.Dev = ctx.Bool("dev")

	conn, err := grpc.Dial(":4455", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	protoClient := client.NewLoveClient(conn)

	m := macaron.Classic()
	m.Use(Renderer())
	m.Get("/", func(ctx *macaron.Context) {

		ctx.HTML(200, "index")
	})
	m.Get("/auth/vk", func(ctx *macaron.Context) {
		ctx.Redirect(vk.AuthURL())
	})
	m.Get("/auth/vk/callback", func(ctx *macaron.Context) {
		var code = ctx.Query("code")

		at, err := vk.GetTokenByCode(code)
		if err != nil {
			log.Printf("err get access token: %v", err)
			ctx.Error(500, err.Error())
			return
		}

		reply, err := protoClient.VkAuth(context.Background(), &client.VkAuthRequest{VkToken: at})
		if err != nil {
			log.Printf("err rpc vk_auth: %v", err)
			ctx.Error(500, err.Error())
			return
		}

		log.Printf("authorize success user=%d", reply.GetUser().GetId())

		ctx.Redirect("/")
	})
	m.Run()
}
