package document

import (
	"context"
	"fmt"
	constant"gh5-backend/pkg/constants"
	"gh5-backend/pkg/gcs"
	"mime/multipart"
	"os"
	"time"
)

func generateUniqueFileName(name string) string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	uniqueFilename := fmt.Sprintf("%s_%d.png", name, timestamp)
	return uniqueFilename
}

func UploadAndSavePath(ctx context.Context, document *multipart.FileHeader, dir string, name string) (string, error) {
	// Generate a unique filename or use your own logic
	fileName := generateUniqueFileName(name)
	filePath := dir + "/" + fileName

	documentOpen, err := document.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open document: %v", err)
	}
	defer documentOpen.Close()

	bucketName := os.Getenv("GOOGLE_BUCKET_NAME")

	// Initialize the GCS client
	gcsClient, err := gcs.NewGCSClient(ctx, bucketName, os.Getenv(constant.WORK_DIR)+os.Getenv(constant.SERVICE_ACCOUNT_FILENAME))
	if err != nil {
		return "", fmt.Errorf("failed to create GCS client: %v", err)
	}

	// Upload the file to Google Cloud Storage
	fileUrl, err := gcsClient.UploadFile(ctx, documentOpen, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to GCS: %v", err)
	}

	return fileUrl, nil
}
