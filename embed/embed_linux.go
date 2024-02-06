package embed

import (
	_ "embed"
)

//go:embed yt-dlp_linux
var YtDlpBinary []byte

//go:embed ffmpeg_linux
var FfmpegBinary []byte