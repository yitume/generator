package cmd

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"

	"github.com/yitume/generator/internal/gen"
	"github.com/yitume/generator/pkg/arg"
)

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new project from mysql database",
	Run:   newProject,
}

func init() {
	newCmd.PersistentFlags().StringVarP(&arg.Database, "db", "d", "shop", `指定数据库名`)
	newCmd.PersistentFlags().StringVarP(&arg.MySQL, "mysql", "m",
		"root:root@mysql+tcp(127.0.0.1:3306)/information_schema", `指定存储(MySQL等)地址`)
	newCmd.PersistentFlags().StringVarP(&arg.Out, "out", "o", "./dist", `指定输出目录`)
	newCmd.PersistentFlags().StringVarP(&arg.SshTunnel, "ssh", "s", "", `开启ssh隧道`)
	newCmd.PersistentFlags().StringVarP(&arg.Table, "table", "t", "%%", `指定表名`)
	newCmd.PersistentFlags().StringVarP(&arg.Module, "module", "M", "git.yitum.com/saas/shop-admin", `指定项目module`)
	newCmd.PersistentFlags().StringVarP(&arg.TmplDir, "tmpl-dir", "T", "tmpl", `指定渲染模板目录`)
	RootCmd.AddCommand(newCmd)
}

func newProject(cmd *cobra.Command, args []string) {
	// 根据数据库解析mysql的table schema
	tableSchemas, err := gen.GetTableSchemas(arg.MySQL, arg.Database, arg.Table)
	if err != nil {
		log.Panic("[GetTableSchemas] getSchema fail", err.Error())
		return
	}

	// 将解析出来的schema数组转换为map格式
	schemaTpls := gen.GetSchemaTpls(tableSchemas)
	gen.Render(schemaTpls)
}
