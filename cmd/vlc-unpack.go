package cmd

import (
	"archiver/lib"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcUnPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack vlc using variable-length code",
	Run:   unPack,
}

const unpackedExtension = "txt"

func unPack(_ *cobra.Command, args []string) {
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

	packed := lib.Decode(string(data))

	fmt.Println(string(data))

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handlerError(err)
	}

}

// TODO: refactor this
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
