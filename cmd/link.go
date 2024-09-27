package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/aadam-ali/second-brain-cli/config"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(linkCmd)
}

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Generate a wiki link for the selected note",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		config := config.GetConfig()

		files := listFiles(config.Root)

		idx, _ := fuzzyfinder.Find(files, func(i int) string { return files[i] })

		filenameWithoutExt, _ := strings.CutSuffix(files[idx], ".md")

		fmt.Printf(" [[%s]]", filenameWithoutExt)
	},
}

func listFiles(root string) []string {
	var files []string

	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if !entry.IsDir() && filepath.Ext(path) == ".md" {
			files = append(files, entry.Name())
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}
