package cmd

import (
	"fmt"
	"godlp/embed"
	"godlp/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current working directory: %v\n", err)
			return
		}

		// Generate the temporary directory path
		tempDlDir := cwd + "/temp"

		scArgs := []string{
			"-o", "%(title)s.%(ext)s",
			"--embed-metadata",
			"--embed-thumbnail",
			"--metadata-from-title", "%(album)s",
			// "--parse-metadata", "%(title)s:%(album)s",
			"--paths", tempDlDir,
			args[0], // Assuming soundcloudURL is the first argument
		}

		musicDir := viper.GetString("music_directory")

		embed.ExecuteYtDlp(scArgs)

		albumName, _ := cmd.Flags().GetString("album")
		if albumName == "" {
			fmt.Println("No album name provided. Grabbing Artist name instead.")
			// Extract artist name from the file
			artistName, err := utils.ExtractArtistNameFromTempDir(tempDlDir)
			if err != nil {
				fmt.Printf("Error extracting artist name: %v\n", err)
				return
			}

			// Use the extracted artist name as the album name
			albumName = artistName
		}

		// Generate the final music directory path
		saveDir := filepath.Join(musicDir, albumName)

		// Move files from tempDir to musicDir
		err = os.MkdirAll(saveDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating music directory: %v\n", err)
			return
		}

		err = utils.MoveFiles(tempDlDir, saveDir)
		if err != nil {
			fmt.Printf("Error moving files: %v\n", err)
			return
		}

		fmt.Println("Files moved to:", saveDir)

		// Remove the tempDir when done
		err = os.RemoveAll(tempDlDir)
		if err != nil {
			log.Printf("Error removing temp directory: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(soundcloudCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// soundcloudCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	soundcloudCmd.Flags().StringP("album", "a", "", "if it's an album, use this string to name the save directory")
}
