package gen

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func getLocalTestData(jFile string) ([]byte, error) {
	jFilePath := filepath.Join(problemSpecificationsDir, jFile)
	if _, err := os.Stat(jFilePath); err != nil {
		return nil, fmt.Errorf("canonical-data.json can't be found: %w", err)
	}

	fmt.Printf("[LOCAL] source: %s\n", jFilePath)

	jTestData, err := os.ReadFile(jFilePath)
	if err != nil {
		return nil, err
	}
	return jTestData, nil
}

func getLatestLocalCommitMessage(jFile string) (Header, error) {
	c := exec.Command("git", "log", "-1", "--oneline", jFile)
	c.Dir = problemSpecificationsDir

	msg, err := c.Output()
	if err != nil {
		return Header{}, fmt.Errorf("failed to determine latest commit message of %s: %w", jFile, err)
	}
	return Header{Commit: string(bytes.TrimSpace(msg)), Origin: defaultOrigin}, nil
}
