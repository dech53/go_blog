package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"server/global"
	"server/utils"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

type Minio struct{}

func (*Minio) UploadImage(file *multipart.FileHeader) (string, string, error) {
	size := float64(file.Size) / (1024 * 1024)
	if size >= float64(global.Config.Upload.Size) {
		return "", "", fmt.Errorf("the image size exceeds the set size, the current size is: %.2f MB, the set size is: %d MB", size, global.Config.Upload.Size)
	}

	ext := filepath.Ext(file.Filename)
	if _, exists := WhiteImageList[ext]; !exists {
		return "", "", errors.New("don't upload files that aren't image types")
	}

	name := strings.TrimSuffix(file.Filename, ext)
	fileKey := utils.MD5V([]byte(name)) + "-" + time.Now().Format("20060102150405") + ext

	fileReader, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer fileReader.Close()

	contentType := getContentType(ext)

	_, err = global.MinioClient.PutObject(
		context.Background(),
		global.Config.Minio.Bucket,
		fileKey,
		fileReader,
		file.Size,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", "", fmt.Errorf("File upload failed: %v", err)
	}

	return global.Config.Minio.ImgPath + fileKey, fileKey, nil
}

func (*Minio) DeleteImage(key string) error {
	err := global.MinioClient.RemoveObject(
		context.Background(),
		global.Config.Minio.Bucket,
		key,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return fmt.Errorf("File delete failed: %v", err)
	}
	return nil
}

func getContentType(ext string) string {
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".svg":
		return "image/svg+xml"
	case ".tiff":
		return "image/tiff"
	case ".ico":
		return "image/x-icon"
	default:
		return "application/octet-stream"
	}
}
