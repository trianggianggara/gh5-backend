package gcs

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCSClient struct {
	client *storage.Client
	bucket string
}

func NewGCSClient(ctx context.Context, bucket, credentialsFile string) (*GCSClient, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}

	return &GCSClient{
		client: client,
		bucket: bucket,
	}, nil
}

func (g *GCSClient) UploadFile(ctx context.Context, src io.Reader, filename string) (string, error) {
	objectName := fmt.Sprintf("uploads/%d-%s", time.Now().Unix(), filename)
	bucket := g.client.Bucket(g.bucket)
	obj := bucket.Object(objectName)

	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, src); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	// Make the object publicly accessible (optional)
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", g.bucket, objectName), nil
}
