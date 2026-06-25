package main

import (
	"devinggo/internal/cmd"
	_ "devinggo/internal/logic"
	_ "devinggo/internal/packed"
	_ "devinggo/modules/bootstrap/logic"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
