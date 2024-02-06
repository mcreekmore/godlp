package embed

import (
	_ "embed"
)

//go:embed yt-dlp-macos
var YtDlpBinary []byte

//go:embed ffmpeg-macos
var FfmpegBinary []byte