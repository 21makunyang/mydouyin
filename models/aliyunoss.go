package models

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadAliyunOss(videoname string, filename string) {
	Endpoint := "https://oss-cn-guangzhou.aliyuncs.com"
	AccessKeyID := access["keyID"]
	AccessKeyIDSecret := access["ketIDSecret"]

	client, err := oss.New(Endpoint, AccessKeyID, AccessKeyIDSecret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bucket, err := client.Bucket("mkydouyin")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	path := "mydouyin/" + videoname
	err = bucket.PutObjectFromFile(path, filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
