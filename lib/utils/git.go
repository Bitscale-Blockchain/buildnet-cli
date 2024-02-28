package utils

import (
	"fmt"
)

func GitCommitChanges(message string, projectDir string) {
	ShellExecute("git add .", projectDir)
	ShellExecute(fmt.Sprintf("git commit -m '%s'", message), projectDir)
}
