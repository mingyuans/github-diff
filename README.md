# github-diff
A GitHub action to download the diff of the GitHub Pull Request. 


# Quick Start

```yaml
      - name: Download PR Diff
        if: ${{ github.event_name == 'pull_request' }}
        id: pr_diff_downloader
        uses: mingyuans/github-diff@main
        with:
          logger-level: debug
          token: ${{ secrets.GITHUB_TOKEN }}
```