package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/wmsx/group_svc/models"
	group "github.com/wmsx/group_svc/proto/group"
	"github.com/wmsx/group_svc/setting"
	"github.com/wmsx/group_svc/handler"
)

const name = "wm.sx.svc.group"

func main() {
	service := micro.NewService(
		micro.Name(name),
		micro.Version("latest"),
		micro.Flags(
			&cli.StringFlag{
				Name:    "env",
				Usage:   "指定运行环境",
				Value:   "dev",
				EnvVars: []string{"RUN_ENV"},
			},
		),
	)

	var env string
	service.Init(
		micro.Action(func(c *cli.Context) error {
			env = c.String("env")
			return nil
		}),

		micro.BeforeStart(func() (err error) {
			if err = setting.SetUp(name, env); err != nil {
				return err
			}
			if err = models.SetUp(); err != nil {
				return err
			}
			return nil
		}),
		micro.AfterStop(func() error {
			models.CloseDB()
			return nil
		}),
	)

	_ = group.RegisterGroupHandler(service.Server(), new(handler.GroupHandler))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
