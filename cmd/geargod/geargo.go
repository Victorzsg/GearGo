package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/Victorzsg/GearGo/cmd/node"
)

var versionFlag bool
var mainCmd = &cobra.Command{
	Use: "command",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			//        version.Print()
		} else {
			cmd.HelpFunc()(cmd, args)
		}
		fmt.Println(args)
	},
}
func main() {
	mainCmd.AddCommand(node.Cmd())
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}

	//初始化队列，连接数据库或者redis或者分配资源

	//绑定事件

	//监听端口

	//释放资源
}
