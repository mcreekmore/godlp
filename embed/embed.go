package embed

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed yt-dlp
var ytDlpBinary []byte

//go:embed ffmpeg
var ffmpegBinary []byte

// ExecuteYtDlp executes yt-dlp from the embedded binary
func ExecuteYtDlp(args []string) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "yt-dlp-embed")
	if err != nil {
		fmt.Printf("Error creating temporary directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Write embedded yt-dlp binary to a temporary file
	ytDlpBinaryPath := tempDir + "/yt-dlp"
	err = os.WriteFile(ytDlpBinaryPath, ytDlpBinary, 0755)
	if err != nil {
		fmt.Printf("Error writing yt-dlp binary to temporary location: %v\n", err)
		return
	}

	// Write embedded ffmpeg binary to a temporary file
	ffmpegBinaryPath := tempDir + "/ffmpeg"
	err = os.WriteFile(ffmpegBinaryPath, ffmpegBinary, 0755)
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
