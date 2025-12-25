package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetStagedDiff() (string, error) {
	// Check if in a git repo
	_, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()
	if err != nil {
		return "", fmt.Errorf("not a git repository")
	}

	// Get staged changes
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git diff: %v", err)
	}

	diff := string(out)
	if strings.TrimSpace(diff) == "" {
		return "", fmt.Errorf("no staged changes found. Use 'git add' to stage files")
	}

	return diff, nil
}
