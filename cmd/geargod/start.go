package geargod

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"strconv"
	"github.com/Victorzsg/GearGo/libgeargo-server"
)

var chaincodeDevMode bool
func startCmd() *cobra.Command {
	return nodeStartCmd
}
var nodeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the node.",
	Long:  `Starts a node that interacts with the network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(args)
	},
}
func serve(args []string) error {

	var port int
	for _,param := range args {
		s := strings.Split(param, "=")
		if s[0] ==  "port" {
			port, _ = strconv.Atoi(s[1])
			fmt.Println(port)
		}
	}

	instance := &libgeargo_server.Geargo_daemon{Port: port,}
	instance.Start()

	return nil
}
