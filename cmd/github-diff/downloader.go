package main

import (
	"fmt"
	"github.com/actions-go/toolkit/core"
	"github.com/actions-go/toolkit/github"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"time"
)

type Downloader struct {
	Arg ActionArg
}

func NewDownloader(arg ActionArg) *Downloader {
	return &Downloader{
		Arg: arg,
	}
}

func getGithubAPIServerURL() string {
	apiServerURL := core.GetInputOrDefault("GITHUB_API_URL", "https://api.github.com")
	return apiServerURL
}

func (d *Downloader) DownloadDiff() ([]byte, error) {
	actionEnv := github.ParseActionEnv()

	diffURL := fmt.Sprintf("%s/repos/%s/%s/pulls/%d",
		getGithubAPIServerURL(),
		actionEnv.Repo.Owner,
		actionEnv.Repo.Repo,
		actionEnv.Issue.Number)

	// Create request
	req, err := http.NewRequest("GET", diffURL, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Set authentication headers
	req.Header.Set("Authorization", "Bearer "+d.Arg.Token)
	req.Header.Set("Accept", "application/vnd.github.diff")
	req.Header.Set("User-Agent", "GitHub-Action-PR-Diff-Downloader")

	// Send request
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to download diff: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("diff download failed with status %d", resp.StatusCode)
	}

	// Read response content
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read diff content: %w", err)
	}

	Get().Debug("Downloading diff from URL", zap.String("diffContent", string(content)))
	return content, nil
}

func (d *Downloader) SaveDiffToFile(content []byte, filename string) error {
	return os.WriteFile(filename, content, 0644)
}
