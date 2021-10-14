package cmd

import (
	"github.com/spf13/cobra"
	"luckDraw/models"
)

var fillredisCmd = &cobra.Command{
	Use:"fillredis",
	Short: "填充纸条数据到redis",
	Long: "填充数据到redis",
	Run:fillRedis,
}
func init(){
	rootCmd.AddCommand(fillredisCmd)
}
func fillRedis(cmd *cobra.Command,args []string){
	models.FillUserOnRedis() //填充数据到redis,运行程序一次后可将该代码注释
}