package initialize

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"server/global"
)

func InitMinio() *minio.Client {
	minioCfg := global.Config.Minio
	minioClient, err := minio.New(minioCfg.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioCfg.AccessKey, minioCfg.SecretKey, ""),
		Secure: minioCfg.UseSSL,
	})
	if err != nil {
		log.Fatalln("MinIO 初始化失败:", err)
	}
	return minioClient
}
