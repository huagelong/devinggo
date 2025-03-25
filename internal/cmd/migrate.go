// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"devinggo/modules/system/pkg/utils"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"path/filepath"
	"time"
)

var (
	MigrateUp = &gcmd.Command{
		Name:        "migrate:up",
		Brief:       "migrate:up [-n N] Apply all or N up migrations",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)

			if err != nil {
				g.Log().Panic(ctx, err)
			}
			opts := gcmd.GetOpt("n")
			m := migrateDB(ctx)
			if opts == nil {
				if err := m.Up(); err != nil {
					defer m.Close()
					return err
				} else {
					g.Log().Debug(ctx, "migrations database up success")
				}
			} else {
				limit := opts.Int()
				if limit >= 0 {
					if err := m.Steps(limit); err != nil {
						defer m.Close()
						return err
					} else {
						g.Log().Debug(ctx, "migrations database up success")
					}
				}
			}
			defer m.Close()
			return
		},
	}

	MigrateDown = &gcmd.Command{
		Name:        "migrate:down",
		Brief:       "migrate:down [-n N] Apply all or N down migrations",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			opts := gcmd.GetOpt("n")
			m := migrateDB(ctx)
			if opts == nil {
				if err := m.Down(); err != nil {
					defer m.Close()
					return err
				} else {
					g.Log().Debug(ctx, "migrations database down success")
				}
			} else {
				limit := opts.Int()
				if limit >= 0 {
					if err := m.Steps(-limit); err != nil {
						defer m.Close()
						return err
					} else {
						g.Log().Debug(ctx, "migrations database down success")
					}
				}
			}
			defer m.Close()
			return
		},
	}

	MigrateGoto = &gcmd.Command{
		Name:        "migrate:goto",
		Brief:       "migrate:goto -v version Migrate to version v",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			opts := gcmd.GetOpt("v")
			m := migrateDB(ctx)
			if opts == nil {
				defer m.Close()
				return gerror.New("version must input")
			} else {
				v := opts.Uint()
				if err := m.Migrate(v); err != nil {
					defer m.Close()
					return err
				} else {
					g.Log().Debug(ctx, "migrations database down success")
				}
			}
			defer m.Close()
			return
		},
	}

	MigrateForce = &gcmd.Command{
		Name:        "migrate:force",
		Brief:       "migrate:force v version  Set version V but don't run migration (ignores dirty state)",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			opts := gcmd.GetOpt("v")
			g.Log().Debug(ctx, "opts:", gcmd.GetOptAll())
			m := migrateDB(ctx)
			if opts == nil {
				defer m.Close()
				return gerror.New("version must input")
			} else {
				v := opts.Int()
				if err := m.Force(v); err != nil {
					defer m.Close()
					return err
				} else {
					g.Log().Debug(ctx, "migrations database force success")
				}
			}
			defer m.Close()
			return
		},
	}

	MigrateCreate = &gcmd.Command{
		Name:        "migrate:create",
		Brief:       "migrate:create -name name Create a set of timestamped up/down migrations titled NAME",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			opts := gcmd.GetOpt("name")
			if opts == nil {
				return gerror.New("migrate name must input")
			} else {
				name := opts.String()
				var version string
				var err error

				dbType := utils.GetDbType()

				directory := "./resource/migrations"
				if dbType == "postgres" {
					directory = "./resource/migrations_pgsql"
				}

				dir := filepath.Clean(directory)
				ext := ".sql"
				startTime := time.Now()
				timezone, err := time.LoadLocation("UTC")
				if err != nil {
					g.Log().Panic(ctx, err)
				}
				version = startTime.In(timezone).Format("20060102150405")

				versionGlob := filepath.Join(dir, version+"_*"+ext)
				matches, err := filepath.Glob(versionGlob)

				if err != nil {
					g.Log().Panic(ctx, err)
				}

				if len(matches) > 0 {
					g.Log().Panic(ctx, "duplicate migration version:", version)
				}

				if err = os.MkdirAll(dir, os.ModePerm); err != nil {
					return err
				}

				for _, direction := range []string{"up", "down"} {
					basename := fmt.Sprintf("%s_%s.%s%s", version, name, direction, ext)
					filename := filepath.Join(dir, basename)

					if err = createFile(filename); err != nil {
						return err
					}
					absPath, _ := filepath.Abs(filename)
					g.Log().Debug(ctx, absPath)
				}
			}
			return
		},
	}
)

func createFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return err
	}

	return f.Close()
}

func migrateDB(ctx context.Context) (m *migrate.Migrate) {
	db, err := g.DB().Master()
	if err != nil {
		g.Log().Panic(ctx, err)
	}

	dbConfig := g.DB().GetConfig()
	link := dbConfig.Link
	dbType := utils.GetDbType()

	dbName := ""
	if g.IsEmpty(link) {
		dbName = dbConfig.Name
	} else {
		dbName, err = utils.GetConnectDbName(link)
		if err != nil {
			g.Log().Panic(ctx, err)
		}
	}

	g.Log().Debug(ctx, "dbName:", dbName)
	if g.IsEmpty(dbName) {
		g.Log().Fatal(ctx, "migrateDB error: dbName is null")
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		g.Log().Panic(ctx, err)
	}

	var driver database.Driver
	migrationsDir := "file://resource/migrations"
	if dbType == "postgres" {
		driver, err = postgres.WithConnection(ctx, conn, &postgres.Config{
			MigrationsTable: "system_migrations",
			DatabaseName:    dbName,
		})
		migrationsDir = "file://resource/migrations_pgsql"
	} else {
		driver, err = mysql.WithConnection(ctx, conn, &mysql.Config{
			MigrationsTable: "system_migrations",
			DatabaseName:    dbName,
		})
	}

	if err != nil {
		g.Log().Panic(ctx, err)
	}

	m, err = migrate.NewWithDatabaseInstance(
		migrationsDir,
		dbName, driver)
	if err != nil {
		g.Log().Panic(ctx, err)
	}

	return m
}
