package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sk2023-z/cloud-station/store"
)

var (
	// 对象是否实现了接口的约束
	// _ store.Uploader = &AliOssStore{}
	// 不需要类型的值只要类型的判断

	_ store.Uploader = (*AliOssStore)(nil)
)

type AliOssStore struct {
	client   *oss.Client
	listener oss.ProgressListener
}
type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint, access_key, access_secret has one empty")
	}
	return nil
}

func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client:   c,
		listener: NewDefaultProgressListener(),
	}, nil
}

func (s *AliOssStore) Upload(bucketName, objectKey, fileName string) error {

	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	if err := bucket.PutObjectFromFile(objectKey, fileName, oss.Progress(s.listener)); err != nil {
		return err
	}

	downloadURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}

	fmt.Printf("文件下载URL: %s \n", downloadURL)
	fmt.Println("请在1天之内下载.")
	return nil
}
