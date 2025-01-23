package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "Simple archive",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handlerError(err)
	}
}
func handlerError(err error) {
	_, _ = fmt.Print(os.Stderr, err)
	os.Exit(1)
}
