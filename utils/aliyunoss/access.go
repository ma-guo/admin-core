package aliyunoss

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ma-guo/admin-core/app/common/consts"
	"github.com/ma-guo/admin-core/config"
	"github.com/ma-guo/admin-core/xorm/services"
	"github.com/ma-guo/niuhe"
	cache "github.com/patrickmn/go-cache"
)

type Aliyun struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	Buncket         string
	isAliyun        bool
}

var localCache *cache.Cache

func NewAliyun() *Aliyun {
	key := "aliyunClientCache"
	// 先查看缓存
	if tmp, has := localCache.Get(key); has {
		return tmp.(*Aliyun)
	}
	tmp := &Aliyun{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Endpoint:        "",
		Buncket:         "",
		isAliyun:        false,
	}
	svc := services.NewSvc()
	defer svc.Close()
	vendor, has, err := svc.Vendor().GetByKey(consts.FileVendorKey, consts.FileVendorKey)
	// 查看是否有 aliyun 配置
	if !has || err != nil {
		niuhe.LogInfo("GetByKey error: %v", err)
		return tmp
	}
	if vendor.Value != consts.FileVendorEnum.Aliyun.Value {
		return tmp
	}
	values, err := svc.Vendor().GetByVendorToMap(consts.FileVendorEnum.Aliyun.Value)
	if err != nil {
		niuhe.LogInfo("GetByVendorToMap error: %v", err)
		return tmp
	}
	accessKeyID := values[consts.FileAccesKey]
	accessKeySecret := values[consts.FileSecretKey]
	endpoint := values[consts.FileEndpoint]
	bucketName := values[consts.FileBucket]

	tmp = &Aliyun{
		AccessKeyId:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		Endpoint:        endpoint,
		Buncket:         bucketName,
		isAliyun:        true,
	}
	localCache.Set(key, tmp, 24*time.Hour)
	return tmp
}

// 添加 aliyun oss 访问 url
func (ali *Aliyun) SignUrl(fileUrl string, expires time.Duration) string {
	if !ali.isAliyun || !config.Config.OssEncodeUrl {
		return fileUrl
	}
	if !strings.Contains(fileUrl, ali.Endpoint) {
		return fileUrl
	}
	if !strings.Contains(fileUrl, ali.Buncket) {
		return fileUrl
	}
	// 先查看缓存
	if tmpSign, has := localCache.Get(fileUrl); has {
		return tmpSign.(string)
	}
	prefix := fmt.Sprintf("https://%v.%s/", ali.Buncket, ali.Endpoint)
	objectKey := fileUrl[len(prefix):]

	// 创建 OSS 客户端
	client, err := oss.New(ali.Endpoint, ali.AccessKeyId, ali.AccessKeySecret)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return fileUrl
	}

	// 获取 Bucket
	bucket, err := client.Bucket(ali.Buncket)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return fileUrl
	}

	// 生成签名 URL（用于 GET 下载）
	signedURL, err := bucket.SignURL(objectKey, oss.HTTPGet, int64(expires.Seconds()))
	if err != nil {
		niuhe.LogInfo("%v", err)
		return fileUrl
	}
	// signUrl 需要做 decodeURIComponent 操作
	decoded := ali.decodeURIComponent(signedURL)
	// niuhe.LogInfo("SignUrl: %s\n%v\n%v\n%v\n%v", decoded, prefix, objectKey, fileUrl, signedURL)
	localCache.Set(fileUrl, decoded, expires)
	return decoded
}
func (ali *Aliyun) decodeURIComponent(signedURL string) string {
	decoded, err := url.QueryUnescape(signedURL)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return signedURL
	}
	return decoded
}

func init() {
	localCache = cache.New(24*time.Hour, 10*time.Minute)
}
