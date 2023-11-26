package aliyun_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zjy-z/cloud-station/store"
	"github.com/zjy-z/cloud-station/store/aliyun"
)

var (
	uploader store.Uploader
)

var (
	AccessKey       = os.Getenv("ALI_AK")
	AccessKeySecret = os.Getenv("ALI_SK")
	OssEndpoint     = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName      = os.Getenv("ALI_BUCKET_NAME")
)

func TestUpload(t *testing.T) {
	// 使用 assert 编写测试用例的断言
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		t.Log("upload ok")
	}
}

func init() {
	ali, err := aliyun.NewAliOssStore(&aliyun.Options{
		Endpoint:     OssEndpoint,
		AccessKey:    AccessKey,
		AccessSecret: AccessKeySecret,
	})
	if err != nil {
		panic(err)
	}

	uploader = ali
}
