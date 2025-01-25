package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "Vlc",
	Short: "Pack vlc",
	Run:   pack,
}

const packedExtension = "vls"

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handlerError(errors.New("pack needs at least 1 argument"))
	}
	filePath := args[0]
	r, err := os.Open(filePath)
	if err != nil {
		handlerError(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handlerError(err)
	}

	packed := ""

	fmt.Println(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handlerError(err)
	}

}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
