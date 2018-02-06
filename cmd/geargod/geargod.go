package geargod

import (
	"github.com/spf13/cobra"
	"fmt"
)

const FuncName = "geargod"
var (
	stopPidFile string
)
// Cmd returns the cobra command for geargo
func Cmd() *cobra.Command {
	geargoCmd.AddCommand(startCmd())
	return geargoCmd
}
var geargoCmd = &cobra.Command{
	Use:   FuncName,
	Short: fmt.Sprintf("%s specific commands.", FuncName),
	Long:  fmt.Sprintf("%s specific commands.", FuncName),
}