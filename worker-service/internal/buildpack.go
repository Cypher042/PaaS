package internal

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

const DefaultBuilder = "paketobuildpacks/builder:base"

func BuildBuildpack(ctx context.Context, path, image string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"pack", "build", image,
		"--path", path,
		"--builder", DefaultBuilder,
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"buildpack build failed: %v\nstderr: %s",
			err,
			stderr.String(),
		)
	}

	return nil
}

func TagImage(ctx context.Context, localImage, remoteImage string) error {
	cmd := exec.CommandContext(ctx, "docker", "tag", localImage, remoteImage)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker tag failed: %v\n%s", err, stderr.String())
	}
	return nil
}
