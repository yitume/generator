package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yitume/generator/pkg/version"
)

func init() {
	var short bool
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print out build version information",
		Run: func(cmd *cobra.Command, args []string) {
			if short {
				fmt.Println(version.Info)
			} else {
				fmt.Println(version.Info.LongForm())
			}
		},
	}
	cmd.PersistentFlags().BoolVarP(&short, "short", "s", short, "Displays a short form of the version information")

	RootCmd.AddCommand(cmd)
}
