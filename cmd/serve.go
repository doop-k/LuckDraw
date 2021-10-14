package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"luckDraw/models"
	"sync"
)

var wg *sync.WaitGroup
var lock sync.Mutex
var serveCmd = &cobra.Command{
	Use:"serve",
	Short:"启动服务",
	Long:"启动抽取纸条服务",
	Run:runServe,
}
func init(){
	rootCmd.AddCommand(serveCmd)
	initConfig()
}
func runServe(cmd *cobra.Command,args []string){
	models.RunLuckDraw(wg,lock)
}
func initConfig(){
	wg = &sync.WaitGroup{}
	lock = sync.Mutex{}
	fmt.Println("------------------------------")
	fmt.Println("当前用户:",len(models.Users))
	fmt.Println("当前纸条:",len(models.Papers))
	fmt.Println("------------------------------")
}