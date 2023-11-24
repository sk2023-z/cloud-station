package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint   = "oss-cn-beijing.aliyuncs.com"
	accessKey  = "xxx"
	secretKey  = "xxx"
	bucketName = "xxx"
	uploadFile = ""
)

func upload(file_path string) error {
	// 实例化客户端
	client, err := oss.New(endpoint, accessKey, secretKey)
	if err != nil {
		// HandleError(err)
		return err
	}
	// 获取bucket对象

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 上传文件

	return bucket.PutObjectFromFile(file_path, file_path)
}

func validate() error {
	if endpoint == "" || accessKey == "" || secretKey == "" {
		return fmt.Errorf("endpoint, access_key, secret_key has one empty")
	}

	if uploadFile == "" {
		return fmt.Errorf("upload file path required")
	}

	return nil
}

func loadParams() {
	flag.StringVar(&uploadFile, "f", "", "上传文件的名称")
	flag.Parse()
}

func main() {
	// 参数加载
	loadParams()

	// 参数验证
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常：%s\n", err)
		os.Exit(1)
	}

	if err := upload(uploadFile); err != nil {
		fmt.Printf("上传文件异常：%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("文件：%s 上传完成\n", uploadFile)
}
