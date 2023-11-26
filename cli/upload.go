package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zjy-z/cloud-station/store"
	"github.com/zjy-z/cloud-station/store/aliyun"
)

var (
	ossProvider  string
	ossEndpoint  string
	accessKey    string
	accessSecret string
	bucketName   string
	uploadFile   string
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件上传",
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			upload store.Uploader
			err    error
		)
		switch ossProvider {
		case "aliyun":
			upload, err = aliyun.NewAliOssStore(&aliyun.Options{
				Endpoint:     ossEndpoint,
				AccessKey:    accessKey,
				AccessSecret: accessSecret,
			})
		case "tx":
		case "aws":
		default:
			return fmt.Errorf("not support oss storage provider")
		}

		if err != nil {
			return err
		}

		return upload.Upload(bucketName, uploadFile, uploadFile)
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&ossEndpoint, "endpoint", "e", "oss-cn-beijing.aliyuncs.com", "oss storage provider enpoint")
	f.StringVarP(&accessKey, "access_key", "k", "", "oss storage provider ak")
	f.StringVarP(&accessSecret, "access_secret", "s", "", "oss storage provider sk")
	f.StringVarP(&bucketName, "bucket_name", "b", "", "oss storage provider bucket name")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "upload file name")
	RootCmd.AddCommand(UploadCmd)
}
