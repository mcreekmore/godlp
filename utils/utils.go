package utils

import (
	"fmt"
	"github.com/mcreekmore/godlp/embed"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
func ExecuteYtDlp(args []string) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "embed")
	if err != nil {
		fmt.Printf("Error creating temporary directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Write embedded yt-dlp binary to a temporary file
	ytDlpBinaryPath := tempDir + "/yt-dlp"
	err = os.WriteFile(ytDlpBinaryPath, embed.YtDlpBinary, 0755)
	if err != nil {
		fmt.Printf("Error writing yt-dlp binary to temporary location: %v\n", err)
		return
	}

	// Write embedded ffmpeg binary to a temporary file
	ffmpegBinaryPath := tempDir + "/ffmpeg"
	err = os.WriteFile(ffmpegBinaryPath, embed.FfmpegBinary, 0755)
	if err != nil {
		fmt.Printf("Error writing ffmpeg binary to temporary location: %v\n", err)
		return
	}

	// Append ffmpeg path to args
	args = append(args, "--ffmpeg-location", ffmpegBinaryPath)

	// Execute yt-dlp from the temporary location
	cmd := exec.Command(ytDlpBinaryPath, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))
		fmt.Printf("Error executing yt-dlp command: %v\n", err)
		return
	}

	fmt.Println(string(output))
}

func ExecuteFfmpeg(args []string) error {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "embed")
	if err != nil {
		fmt.Printf("Error creating temporary directory: %v\n", err)
		return err
	}
	defer os.RemoveAll(tempDir)

	// Write embedded ffmpeg binary to a temporary file
	ffmpegBinaryPath := tempDir + "/ffmpeg"
	err = os.WriteFile(ffmpegBinaryPath, embed.FfmpegBinary, 0755)
	if err != nil {
		fmt.Printf("Error writing ffmpeg binary to temporary location: %v\n", err)
		return err
	}

	// Execute ffmpeg from the temporary location
	cmd := exec.Command(ffmpegBinaryPath, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))
		fmt.Printf("Error executing ffmpeg command: %v\n", err)
		return err
	}

	fmt.Println(string(output))
	return nil
}

func MoveFiles(source, dest string) error {
	files, err := filepath.Glob(filepath.Join(source, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println("file", file)
		_, fileName := filepath.Split(file)
		inputFile, err := os.Open(file)
		if err != nil {
			fmt.Printf("Couldn't open source file: %v\n", err)
		}
		outputFile, err := os.Create(dest + "/" + fileName)
		if err != nil {
			inputFile.Close()
			fmt.Printf("Couldn't open dest file: %v\n", err)
			return err
		}
		defer outputFile.Close()
		_, err = io.Copy(outputFile, inputFile)
		inputFile.Close()
		if err != nil {
			fmt.Printf("Writing to output file failed: %v\n", err)
			return err
		}
		// The copy was successful, so now delete the original file
		err = os.Remove(file)
		if err != nil {
			fmt.Printf("Failed removing original file: %v\n", err)
			return err
		}
	}

	return nil
}

func ExtractArtistNameFromFile(tempDir string) (string, error) {
	files, err := filepath.Glob(filepath.Join(tempDir, "*"))
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", fmt.Errorf("no files found in temp directory")
	}

	// Assuming the first file in tempDir contains the downloaded file name
	fileName := filepath.Base(files[0])

	// Split the filename using the "-" delimiter
	parts := strings.Split(fileName, "-")
	if len(parts) < 2 {
		return "", fmt.Errorf("unable to extract artist name from the filename")
	}

	// Extract the artist name and trim whitespaces
	artistName := strings.TrimSpace(parts[0])

	return artistName, nil
}

func ChangeAlbumNameWithFFmpeg(directory, albumName string) error {
	files, err := filepath.Glob(filepath.Join(directory, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		ffmpegArgs := []string{
			"-y",
			"-i", file,
			"-metadata", fmt.Sprintf("album=%s", albumName),
			"-c", "copy",
			file + "_temp.mp3",
		}
		err := ExecuteFfmpeg(ffmpegArgs)
		if err != nil {
			return err
		}

		// Replace the original file with the one containing updated metadata
		err = os.Rename(file+"_temp.mp3", file)
		if err != nil {
			return err
		}
	}

	return nil
}
