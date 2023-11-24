package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
)

var (
	AccessKey       = os.Getenv("ALI_AK")
	AccessKeySecret = os.Getenv("ALI_SK")
	OssEndpoint     = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName      = os.Getenv("ALI_BUCKET_NAME")
)

// 测试oss基本功能
func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()

	if err != nil {
		// HandleError(err)
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessKeySecret)
	if err != nil {
		// HandleError(err)
		panic(err)
	}
	client = c
}
