package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/aadam-ali/second-brain-cli/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dailyCmd)
}

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "create a daily note",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()

		filepath := cfg.DailyNotePath

		dailyNoteExists := checkIfDailyNoteExists(filepath)

		if !dailyNoteExists {
			content := renderDailyNoteContent(cfg.Yesterday, cfg.Today, cfg.Tomorrow)
			createNote(filepath, content)

			fmt.Println(filepath)
		} else {
			fmt.Printf("Note already exists: %s\n", filepath)
		}
	},
}

func checkIfDailyNoteExists(filepath string) bool {
	_, err := os.Stat(filepath)

	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}

	log.Fatal(err)

	return false
}

func renderDailyNoteContent(yesterday string, today string, tomorrow string) string {
	return fmt.Sprintf(`# %s

[[%s]] - [[%s]]

## Journal


## Notes Created Today
`, today, yesterday, tomorrow)
}