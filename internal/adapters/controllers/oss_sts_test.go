// @Author nono.he 2023/4/12 12:15:00
package controllers

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func TestOssSts(t *testing.T) {

	client, err := oss.New("oss-cn-guangzhou.aliyuncs.com",
		"STS.NUVwyPBDVcDQP2FMoNiG4Q8hD",
		"J156xkBHYJqPVrP7i5jF3Jqzvc1xEXLdJfMnWRek7Dpe",
		oss.SecurityToken(""+
			"CAIS8wF1q6Ft5B2yfSjIr5bjPMPkr5t31Ia6UhT3qW8bZcgYvv3Dpjz2IHhOe3FuBegXs/o3lG1X7v4elqd0UIRyTkzfY8JH9o5Q/VsN7Uc+K5Tng4YfgbiJREKxaXeiruKwDsz9SNTCAITPD3nPii50x5bjaDymRCbLGJaViJlhHL91N0vCGlggPtpNIRZ4o8I3LGbYMe3XUiTnmW3NFkFlyGEe4CFdkf3jnZDFsUSO1AShl7dM/d/LT8L6P5U2DvBWSMyo2eF6TK3F3RNL5gJCnKUM1/Qepm+Z4onMXgcJs0jZa7KI6LN1JQp+fbMqmAWELAAgMpcagAGUne7K3213pDddHOjgYD2gUvzu5BmwvzXVKtjdOBR8Mf3QyUcc6kngHhXZGC+LDh9bOFSP6NTLgJnY7MmwHONpXVwwrk5Vs18wkN1LZlt+lmBDSBRm9XUp6p+Qx7S0CIgNy3XFn33BIkySlhrlZLJQugzkw9xH2dSqXPbHz88fjA=="))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写Bucket名称，例如examplebucket。
	bucketName := "dangguancaihua"
	// 填写Object的完整路径，完整路径中不能包含Bucket名称，例如exampledir/exampleobject.txt。
	objectName := "test-dir/test2.jpg"
	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	filepath := "D:\\tmp\\default.jpg"
	bucket, err := client.Bucket(bucketName)
	// 通过STS授权第三方上传文件。
	err = bucket.PutObjectFromFile(objectName, filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("upload success")
}
