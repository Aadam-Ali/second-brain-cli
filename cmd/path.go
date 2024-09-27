package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/aadam-ali/second-brain-cli/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pathCmd)
}

var pathCmd = &cobra.Command{
	Use:   "path [title]",
	Short: "Output absolute path of the inputted note",
	Long:  "Output absolute path of the inputted given the title in snake case format (filename) without '.md'",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.GetConfig()

		path := findFile(config.Root, args[0])

		fmt.Println(path)
	},
}

func findFile(root string, name string) string {
	notepath := ""
	name = name + ".md"

	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if !entry.IsDir() && name == entry.Name() {
			notepath = path
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return notepath
}
