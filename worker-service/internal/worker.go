package internal

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

type OrchestrateRequest struct {
	RepoURL       string
	ImageName     string
	ImageTag      string
	KeepWorkspace bool
}

type OrchestrateResult struct {
	Image     string
	ClonePath string
}

type Orchestrator struct{}

func CloneRepo(ctx context.Context, repoURL, dest string) error {
	cmd := exec.CommandContext(ctx, "git", "clone", repoURL, dest)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("git error: %v, details: %s", err, stderr.String())
	}
	return nil
}
