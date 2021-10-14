package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:"luckDraw",
	Short: "纸条抽取",
	Long: "redis抽取纸条",
}

func Execute(){
	if err :=rootCmd.Execute();err!=nil{
		fmt.Println("err:",err)
		return
	}
}
