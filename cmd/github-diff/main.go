package main

import (
	"github.com/actions-go/toolkit/core"
	"go.uber.org/zap"
	"os"
)

func main() {
	arg := ParseArg()
	Init(arg.LoggerLevel)
	downloader := NewDownloader(arg)
	downloadedDiff, downErr := downloader.DownloadDiff()
	if downErr != nil {
		Get().Error("Failed to download PR diff", zap.Error(downErr))
		core.SetFailed(downErr.Error())
		os.Exit(2)
	}
	saveErr := downloader.SaveDiffToFile(downloadedDiff, arg.FileName)
	if saveErr != nil {
		Get().Error("Failed to save PR diff to file", zap.Error(saveErr))
		core.SetFailed(saveErr.Error())
		os.Exit(3)
	} else {
		Get().Info("PR diff downloaded and saved successfully", zap.String("file_name", arg.FileName))
		core.SetOutput("file-path", arg.FileName)
	}
}
