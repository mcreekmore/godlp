package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mcreekmore/godlp/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// soundcloudCmd represents the soundcloud command
var soundcloudCmd = &cobra.Command{
	Use:   "soundcloud",
	Short: "For downloading tracks from soundcloud",
	Long: `Downloads single tracks or collections from soundcloud and moves 
them to your configs 'music_directory' path when done. Pass an optional
--album flag to set the directory name they will be saved in. yt-dlp doesn't
directly support grabbing albumn titles. By default, the artist name is used
in place.`,
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
		tempDir := cwd + "/temp"

		scArgs := []string{
			"-o", "%(title)s.%(ext)s",
			"--embed-metadata",
			"--embed-thumbnail",
			"--metadata-from-title", "%(album)s",
			"--paths", tempDir,
			args[0], // Assuming soundcloudURL is the first argument
		}

		utils.ExecuteYtDlp(scArgs)

		albumName, _ := cmd.Flags().GetString("album")
		if albumName == "" {
			fmt.Println("No album name provided. Grabbing Artist name instead.")
			artistName, err := utils.ExtractArtistNameFromFile(tempDir)
			if err != nil {
				fmt.Printf("Error extracting artist name: %v\n", err)
				return
			}
			albumName = artistName
		} else {
			fmt.Println("Album flag provided. Writing with ffmpeg...")
			utils.ChangeAlbumNameWithFFmpegConcurrent(tempDir, albumName)
		}

		// Generate the final music directory path
		musicDir := viper.GetString("music_directory")
		saveDir := filepath.Join(musicDir, albumName)

		// Move files from tempDir to musicDir
		err = os.MkdirAll(saveDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating music directory: %v\n", err)
			return
		}

		err = utils.MoveFiles(tempDir, saveDir)
		if err != nil {
			fmt.Printf("Error moving files: %v\n", err)
			return
		}

		fmt.Println("Files moved to:", saveDir)

		// Remove the tempDir when done
		err = os.RemoveAll(tempDir)
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
