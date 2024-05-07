package fileupload

import (
	"fmt"

	"github.com/ma-guo/admin-core/app/common/consts"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ziipin-server/niuhe"
)

type Aliyun struct {
	Provider
	acessKey   string // accessKeyId
	secretKey  string // accessKeySecret
	endpoint   string // endpoint
	bucketName string // bucketName
	ossurl     string // ossurl
	prefix     string
}

func NewAliyun(dict map[string]string) *Aliyun {
	dao := newDict(dict)
	yun := &Aliyun{
		acessKey:   dao.find(consts.FileAccesKey),
		secretKey:  dao.find(consts.FileSecretKey),
		endpoint:   dao.find(consts.FileEndpoint),
		bucketName: dao.find(consts.FileBucket),
		prefix:     dao.find(consts.FilePrefix),
	}
	yun.ossurl = fmt.Sprintf("https://%s.%s", yun.bucketName, yun.endpoint)
	return yun
}

func (aliyun *Aliyun) getBucket() (*oss.Bucket, error) {
	client, err := oss.New(aliyun.endpoint, aliyun.acessKey, aliyun.secretKey)
	if err != nil {
		niuhe.LogInfo("New client error: %v", err)
		return nil, err
	}

	bucket, err := client.Bucket(aliyun.bucketName)
	if err != nil {
		niuhe.LogInfo("Get bucket error: %v", err)
		return nil, err
	}

	return bucket, nil

}

// Upload 上传文件到阿里云
func (aliyun *Aliyun) Upload(localFile, name string) (string, string, error) {
	bucket, err := aliyun.getBucket()
	if err != nil {
		return "", "", err
	}

	key := fmt.Sprintf("%s/%s", aliyun.prefix, name)
	err = bucket.PutObjectFromFile(key, localFile)
	if err != nil {
		niuhe.LogInfo("Put object error: %v", err)
		return "", "", err
	}

	return fmt.Sprintf("%s/%s", aliyun.ossurl, key), key, nil
}

func (aliyun *Aliyun) Delete(key string) error {
	bucket, err := aliyun.getBucket()
	if err != nil {
		niuhe.LogInfo("Get bucket error: %v", err)
		return err
	}

	err = bucket.DeleteObject(key)
	if err != nil {
		niuhe.LogInfo("Delete object error: %v", err)
		return err
	}

	return nil
}
