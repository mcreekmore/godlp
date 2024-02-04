package cmd

import (
	"fmt"
	"godlp/embed"

	"github.com/spf13/cobra"
)

// soundcloudCmd represents the soundcloud command
var soundcloudCmd = &cobra.Command{
	Use:   "soundcloud",
	Short: "For downloading tracks from soundcloud",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a SoundCloud URL.")
			return
		}

		// soundcloudURL := args[0]
		dlpArgs := []string{
			"-o",
			"%(title)s.%(ext)s",
			"--embed-metadata",
			"--embed-thumbnail",
			"--metadata-from-title",
			"'%(album)s'",
			args[0], // Assuming soundcloudURL is the first argument
		}

		embed.ExecuteYtDlp(dlpArgs)

		// // Call yt-dlp command with the provided URL
		// dlpCmd := exec.Command("yt-dlp", "-o", "'%(title)s.%(ext)s'", "--embed-metadata", "--embed-thumbnail", "--metadata-from-title", "'%(album)s'", soundcloudURL)
	},
}

func init() {
	rootCmd.AddCommand(soundcloudCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// soundcloudCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// soundcloudCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
