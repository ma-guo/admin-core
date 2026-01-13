package fileupload

import (
	"context"
	"fmt"

	"github.com/ma-guo/admin-core/app/common/consts"

	"github.com/ma-guo/niuhe"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"

	"github.com/qiniu/go-sdk/v7/storage"
)

// 七牛存储, 上传文件到七牛云
type Qiniu struct {
	Provider
	zone      *storage.Zone
	accessKey string
	secretKey string
	bucket    string
	ossurl    string
	prefix    string
}

func NewQiniu(dict map[string]string) *Qiniu {
	dao := newDict(dict)
	niu := &Qiniu{
		accessKey: dao.find(consts.FileAccesKey),
		secretKey: dao.find(consts.FileSecretKey),
		bucket:    dao.find(consts.FileBucket),
		ossurl:    dao.find(consts.FileOssurl),
		prefix:    dao.find(consts.FilePrefix),
	}
	niu.setZone(dao.find(consts.FileZone))
	return niu
}

// 上传文件到七牛云, name 即为 key
func (niu *Qiniu) Upload(localFile, name, fileType string) (string, string, error) {
	policy := storage.PutPolicy{
		Scope: niu.bucket,
	}

	mac := qbox.NewMac(niu.accessKey, niu.secretKey)
	token := policy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          niu.zone, // 空间对应的机房
		UseHTTPS:      false,    // 是否使用https域名
		UseCdnDomains: false,    // 上传是否使用CDN上传加速
	}
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	extra := storage.PutExtra{
		Params: map[string]string{
			"x:name": name,
		},
	}
	key := niu.prefix + "/" + name
	err := uploader.PutFile(context.Background(), &ret, token, key, localFile, &extra)
	if err != nil {
		niuhe.LogInfo("%v, %v, %v, %v", localFile, key, token, niu.zone)
		niuhe.LogInfo("%v", err)
		return "", "", err
	}
	fileUrl := fmt.Sprintf("%s/%s", niu.ossurl, ret.Key)
	// niuhe.LogInfo("Upload success: %v, %v", ret.Hash, ret.Key)
	return fileUrl, ret.Key, nil
}

func (niu *Qiniu) Delete(key string) error {
	mac := auth.New(niu.accessKey, niu.secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
		Zone:     niu.zone,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	bucketManager := storage.NewBucketManager(mac, &cfg)

	err := bucketManager.Delete(niu.bucket, key)
	if err != nil {
		niuhe.LogInfo("Delete file error: %v", err)
		return err
	}
	return nil
}

// 根据配置的区域名称获取区域
func (niu *Qiniu) setZone(name string) {
	zones := map[string]storage.Zone{
		consts.QiniuZoneEnum.Huadong.Value:   storage.ZoneHuadong,
		consts.QiniuZoneEnum.Huabei.Value:    storage.ZoneHuabei,
		consts.QiniuZoneEnum.Huanan.Value:    storage.ZoneHuanan,
		consts.QiniuZoneEnum.Beimei.Value:    storage.ZoneBeimei,
		consts.QiniuZoneEnum.Xinjiapo.Value:  storage.ZoneXinjiapo,
		consts.QiniuZoneEnum.ZheJiang2.Value: storage.ZoneHuadongZheJiang2,
	}
	if z, ok := zones[name]; ok {
		niu.zone = &z
		return
	}
	niu.zone = &storage.ZoneHuadong
}
