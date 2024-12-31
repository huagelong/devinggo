// Package driver
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package driver

import (
	"devinggo/modules/system/pkg/utils/config"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"regexp"
	"strings"

	"github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/xwb1989/sqlparser"
)

type MysqlDriver struct {
	*mysql.Driver
}

func init() {
	// It here registers my custom driver in package initialization function "init".
	// You can later use this type in the database configuration.
	if err := gdb.Register("mysql", &MysqlDriver{}); err != nil {
		panic(err)
	}
}

// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
func (d *MysqlDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &MysqlDriver{
		&mysql.Driver{
			Core: core,
		},
	}, nil
}

// DoFilter is a bind_hook function, which filters the sql and its arguments before it's committed to underlying driver.
// The parameter `link` specifies the current database connection operation object. You can modify the sql
// string `sql` and its arguments `args` as you wish before they're committed to driver.
func (d *MysqlDriver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	enable := config.GetConfigBool(ctx, "tenant.enable", false)
	if !enable {
		return sql, args, nil
	}
	tenantId := GetTenant(ctx)
	if g.IsNil(tenantId) {
		return sql, args, nil
	}
	//格式化贴牌驱动
	smt, err := sqlparser.Parse(sql)
	if err == nil {
		switch stmt := smt.(type) {
		case *sqlparser.Select:
			if has, where, err := d.format(ctx, stmt.From, stmt.Where); err != nil {
				return sql, args, err
			} else if has {
				stmt.Where = where
				sql = d.renameSql(smt, args)
			}
		case *sqlparser.Insert:
			colCnt := len(stmt.Columns)
			exist := false
			for j, v := range stmt.Columns {
				if strings.ToLower(v.String()) == GetColumnName() && ctx.Value(GetColumnName()) != nil {
					exist = true
					for i := 0; i < len(args)/colCnt; i++ {
						args[i*colCnt+j] = tenantId
					}
				}
			}
			if !exist {
				stmt.Columns = append(stmt.Columns, sqlparser.NewColIdent(GetColumnName()))
				switch stmt.Rows.(type) {
				case *sqlparser.Select:
					rows := stmt.Rows.(*sqlparser.Select)
					if has, where, err := d.format(ctx, rows.From, rows.Where); err != nil {
						return sql, args, err
					} else if has {
						rows.Where = where
						sql = d.renameSql(smt, args)
					}
					rows.SelectExprs = append(rows.SelectExprs, &sqlparser.AliasedExpr{
						Expr: &sqlparser.ColName{
							Name: sqlparser.NewColIdent(GetColumnName()),
						},
					})
				case sqlparser.Values:
					var values []sqlparser.ValTuple
					buf := sqlparser.NewTrackedBuffer(func(buf *sqlparser.TrackedBuffer, node sqlparser.SQLNode) {
						v1 := node.(sqlparser.ValTuple)
						v1 = append(v1, sqlparser.NewValArg([]byte(fmt.Sprintf(":v%d", colCnt+1))))
						values = append(values, v1)
					})

					stmt.Rows.Format(buf)
					stmt.Rows = sqlparser.Values(values)
					var argList []interface{}
					for i := 0; i < len(args)/colCnt; i++ {
						argList = append(argList, args[i*colCnt:i*colCnt+colCnt]...)
						argList = append(argList, GetColumnName())
					}
					args = argList

				}
				sql = d.renameSql(smt, args)
			}

		case *sqlparser.Update:
			if has, where, err := d.format(ctx, stmt.TableExprs, stmt.Where); err != nil {
				return sql, args, err
			} else if has {
				stmt.Where = where
				sql = d.renameSql(smt, args)
			}

		case *sqlparser.Delete:
			if has, where, err := d.format(ctx, stmt.TableExprs, stmt.Where); err != nil {
				return sql, args, err
			} else if has {
				stmt.Where = where
				sql = d.renameSql(smt, args)
			}
		}
	}
	return sql, args, nil
}

func (d *MysqlDriver) format(ctx context.Context, tes sqlparser.TableExprs, where *sqlparser.Where) (bool, *sqlparser.Where, error) {
	tenantId := GetTenant(ctx)
	if !g.IsNil(tenantId) {
		if len(tes) > 0 {
			var frm *sqlparser.AliasedTableExpr
			switch v := tes[0].(type) {
			case *sqlparser.JoinTableExpr:
				switch v1 := v.LeftExpr.(type) {
				case *sqlparser.JoinTableExpr:
					return d.format(ctx, []sqlparser.TableExpr{v1}, where)
				default:
					frm = v.LeftExpr.(*sqlparser.AliasedTableExpr)
				}
			case *sqlparser.AliasedTableExpr:
				frm = v
			}

			temp, ok := frm.Expr.(sqlparser.TableName)
			if !ok {
				return false, nil, nil
			}

			tName := temp.Name.String()
			if tt := frm.As.String(); tt != "" {
				tName = tt
			}
			g.Log().Info(ctx, "table name:", tName)
			if where == nil {
				where = sqlparser.NewWhere(sqlparser.WhereStr, sqlparser.NewValArg([]byte(fmt.Sprintf("%s.%s=%d", tName, GetColumnName(), tenantId))))
			} else {
				where.Expr = &sqlparser.AndExpr{
					Left:  where.Expr,
					Right: sqlparser.NewValArg([]byte(fmt.Sprintf("%s.%s=%d", tName, GetColumnName(), tenantId))),
				}
			}

			return true, where, nil
		}

	}
	return false, where, nil
}
func (d *MysqlDriver) renameSql(stm sqlparser.Statement, args []interface{}) string {
	v := sqlparser.String(stm)

	reg := regexp.MustCompile(`:v\d*`)
	v = reg.ReplaceAllString(v, "?")
	return v
}
