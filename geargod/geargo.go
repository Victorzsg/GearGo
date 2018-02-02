package main

import (
	"github.com/spf13/cobra"
	"github.com/Victorzsg/GearGo/libgeargo-server"
	"fmt"
	"github.com/spf13/pflag"
)

func newDaemonCommand() *cobra.Command{

	cmd := &cobra.Command{
		Use:           "geargod [OPTIONS]",
		Short:         "A self-sufficient runtime for containers.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          libgeargo_server.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()
			return runDaemon(flags)
		},
	}
	flags := cmd.Flags()
	fmt.Println(flags)
	return cmd
}

func runDaemon(flag *pflag.FlagSet) error{
	fmt.Println(flag)
	return nil
}

func main() {

	//解析入参
	cmd := newDaemonCommand()
	fmt.Println(cmd)

	//初始化队列，连接数据库或者redis或者分配资源

	//绑定事件

	//监听端口

	//释放资源
}
