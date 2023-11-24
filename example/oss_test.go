package example_test

import (
	"fmt"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
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
	c, err := oss.New("Endpoint", "AccessKeyId", "AccessKeySecret")
	if err != nil {
		// HandleError(err)
		panic(err)
	}
	client = c
}
