package internal

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	Username = os.Getenv("DOCKERHUB_USERNAME")
	Password = os.Getenv("DOCKERHUB_PASSWORD")
)

type DockerMetadata struct {
	ImageName string
	Path      string
	ImageTag  string
}

func ValidateDockerHubCredentials() error {
	if strings.TrimSpace(Username) == "" {
		return fmt.Errorf("DOCKERHUB_USERNAME is required")
	}

	if strings.TrimSpace(Password) == "" {
		return fmt.Errorf("DOCKERHUB_PASSWORD is required")
	}

	return nil
}

func DetectDocker(path string) bool {
	if _, err := os.Stat(path + "/Dockerfile"); err == nil {
		return true
	}
	return false
}

func RunCommand(ctx context.Context, name string, args ...string) error {
	cmd := exec.CommandContext(ctx, name, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"command failed: %s %v\nerr: %v\nstderr: %s",
			name, args, err, stderr.String(),
		)
	}

	return nil
}

func BuildDocker(ctx context.Context, metadata DockerMetadata) error {
	fullImage := fmt.Sprintf("%s/%s:%s", Username, metadata.ImageName, metadata.ImageTag)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	return RunCommand(ctx, "docker", "build", "-t", fullImage, metadata.Path)
}

func DockerLogin(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "docker", "login", "-u", Username, "--password-stdin")
	cmd.Stdin = bytes.NewBufferString(Password)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker login failed: %v\nstderr: %s", err, stderr.String())
	}

	return nil
}

func PushToDockerHub(ctx context.Context, metadata DockerMetadata) error {
	fullImage := fmt.Sprintf("%s/%s:%s", Username, metadata.ImageName, metadata.ImageTag)

	if err := DockerLogin(ctx); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	return RunCommand(ctx, "docker", "push", fullImage)
}
