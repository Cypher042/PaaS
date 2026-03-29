package main


import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func main() {
	fmt.Println("Azure Blob Storage - Clean Example")

	accountName := "repostorer"
	containerName := "quick-sample-container"
	blobName := "sample-blob"

	ctx := context.Background()

	// 🔐 Use CLI credentials (works with `az login`)
	cred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		log.Fatal("Failed to create credential:", err)
	}

	// 🌐 Correct service URL
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	client, err := azblob.NewClient(serviceURL, cred, nil)
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}

	// 📦 Create container (idempotent)
	fmt.Println("Ensuring container exists...")
	_, err = client.CreateContainer(ctx, containerName, nil)
	if err != nil {
		var respErr *azcore.ResponseError
		if errors.As(err, &respErr) && respErr.ErrorCode == "ContainerAlreadyExists" {
			fmt.Println("Container already exists, continuing...")
		} else {
			log.Fatal("Failed to create container:", err)
		}
	}

	// 📤 Upload blob
	fmt.Println("Uploading blob...")
	data := []byte("Hello from Go + Azure 🚀")

	_, err = client.UploadBuffer(ctx, containerName, blobName, data, nil)
	if err != nil {
		log.Fatal("Upload failed:", err)
	}

	// 📋 List blobs
	fmt.Println("Listing blobs:")
	pager := client.NewListBlobsFlatPager(containerName, nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatal("List failed:", err)
		}

		for _, blob := range page.Segment.BlobItems {
			fmt.Println(" -", *blob.Name)
		}
	}

	// 📥 Download blob
	fmt.Println("Downloading blob...")
	resp, err := client.DownloadStream(ctx, containerName, blobName, nil)
	if err != nil {
		log.Fatal("Download failed:", err)
	}

	reader := resp.NewRetryReader(ctx, nil)
	defer reader.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(reader)
	if err != nil {
		log.Fatal("Read failed:", err)
	}

	fmt.Println("Blob content:")
	fmt.Println(buf.String())

	// 🧹 Cleanup
	fmt.Println("Cleaning up...")

	_, err = client.DeleteBlob(ctx, containerName, blobName, nil)
	if err != nil {
		log.Fatal("Delete blob failed:", err)
	}

	_, err = client.DeleteContainer(ctx, containerName, nil)
	if err != nil {
		log.Fatal("Delete container failed:", err)
	}

	fmt.Println("Done ✅")
}
