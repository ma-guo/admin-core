package fileupload

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ma-guo/admin-core/app/common/consts"

	"github.com/ma-guo/niuhe"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Tencent struct {
	Provider
	secretKey string
	secretId  string
	cosHost   string // https...myqcloud.com
	prefix    string
}

func NewTencent(dict map[string]string) *Tencent {
	dao := newDict(dict)
	return &Tencent{
		secretKey: dao.find(consts.FileSecretKey),
		secretId:  dao.find(consts.FileAccesKey),
		cosHost:   dao.find(consts.FileOssurl),
		prefix:    dao.find(consts.FilePrefix),
	}
}

func (tencent *Tencent) getClient() (*cos.Client, error) {
	// ossUrl := fmt.Sprintf("https://%v.cos.%s.myqcloud.com", tencent.cosHost, tencent.CosZone)
	targetUrl, err := url.Parse(tencent.cosHost)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return nil, err
	}
	baseUrl := &cos.BaseURL{BucketURL: targetUrl}
	client := cos.NewClient(baseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过 secretId, secretKey 初始化认证，这里 secretId 和 secretKey 支持直接传入字符串或从环境变量中读取
			SecretID:  tencent.secretId,
			SecretKey: tencent.secretKey,
		},
	})
	return client, nil
}

func (tencent *Tencent) Upload(localFile, name, fileType string) (string, string, error) {
	client, err := tencent.getClient()
	if err != nil {
		niuhe.LogInfo("%v", err)
		return "", "", err
	}

	// 上传文件
	file, err := os.Open(localFile)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return "", "", err
	}
	defer file.Close()
	key := fmt.Sprintf("%s/%s", tencent.prefix, name)
	_, err = client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return "", "", err
	}

	// 返回文件访问地址
	return fmt.Sprintf("%v/%v", tencent.cosHost, key), key, nil
}

func (tencent *Tencent) Delete(key string) error {
	client, err := tencent.getClient()
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	_, err = client.Object.Delete(context.Background(), key, nil)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil

}
