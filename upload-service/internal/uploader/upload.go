package uploader

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func Upload(githubURL string, projectID string) error {
	ctx := context.Background()

	parts := strings.Split(strings.TrimSuffix(githubURL, ".git"), "/")
	owner := parts[len(parts)-2]
	repo := parts[len(parts)-1]

	tarURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/tarball", owner, repo)

	log.Println(" Fetching tarball:", tarURL)

	resp, err := http.Get(tarURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to fetch tarball: %s", resp.Status)
	}

	// GitHub sends gzip compressed tar
	gzr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}

	tr := tar.NewReader(gzr)

	// Azure setup
	accountName := "repostorer"
	containerName := projectID

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return err
	}

	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	client, err := azblob.NewClient(serviceURL, cred, nil)
	if err != nil {
		return err
	}

	_, err = client.CreateContainer(ctx, containerName, nil)
	if err != nil {
		var respErr *azcore.ResponseError
		if !(errors.As(err, &respErr) && respErr.ErrorCode == "ContainerAlreadyExists") {
			return err
		}
	}

	fileCount := 0

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if header.Typeflag == tar.TypeDir {
			continue
		}

		// 🔥 GitHub tar has a top-level folder → remove it
		parts := strings.SplitN(header.Name, "/", 2)
		if len(parts) < 2 {
			continue
		}
		blobName := filepath.ToSlash(parts[1])

		log.Println("Uploading:", blobName)

		_, err = client.UploadStream(
			ctx,
			containerName,
			blobName,
			tr,
			nil,
		)

		if err != nil {
			log.Printf("❌ Failed: %s | %v", blobName, err)
			continue
		}

		fileCount++
	}

	log.Printf("🎉 Done. Uploaded %d files\n", fileCount)
	return nil
}