package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aadam-ali/second-brain-cli/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolP("private", "p", false, "Mark note as private")
}

var newCmd = &cobra.Command{
	Use:   "new [title]",
	Short: "Create a new note",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		config := config.GetConfig()

		title := args[0]
		isPrivate, _ := cmd.Flags().GetBool("private")
		id := time.Now().Format("20060102150405")

		kebabCaseTitle := convertTitleToKebabCase(title)
		filepath := constructNotePath(config.Inbox, id, kebabCaseTitle, isPrivate)
		content := renderNoteContent(id, title)

		createNote(filepath, content)
	},
}

func convertTitleToKebabCase(title string) string {
	lowercaseTitle := strings.ToLower(title)

	alphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	alphanumericTitle := alphanumericRegex.ReplaceAllString(lowercaseTitle, "")

	whitespaceRegex := regexp.MustCompile(`\s+`)
	kebabCaseTitle := whitespaceRegex.ReplaceAllString(alphanumericTitle, "-")

	return kebabCaseTitle
}

func constructNotePath(inbox string, id string, kebabCaseTitle string, isPrivate bool) string {
	extension := "md"
	if isPrivate {
		extension = "private.md"
	}

	return fmt.Sprintf("%s/%s-%s.%s", inbox, id, kebabCaseTitle, extension)
}

func renderNoteContent(id string, title string) string {
	return fmt.Sprintf(`---
id: %s
tags: []
---

# %s

Related:`, id, title)
}

func createNote(path string, content string) {
	f, _ := os.Create(path)
	defer f.Close()
	_, err := f.Write([]byte(content))

	if err != nil {
		panic(err)
	}

	fmt.Println(path)
}
