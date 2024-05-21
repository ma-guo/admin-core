package fileupload

import (
	"os"
	"strings"
	"time"

	"github.com/ma-guo/admin-core/app/common/consts"
	"github.com/ma-guo/admin-core/config"
	"github.com/ma-guo/admin-core/utils"

	"github.com/ma-guo/niuhe"
)

// 本机存储
type Local struct {
	Provider
	dictDao
	ossurl string // 当前存储路径
	prefix string // 文件前缀
}

func NewLocal(dict map[string]string) *Local {
	local := &Local{}
	local.data = dict
	local.ossurl = local.find(consts.FileOssurl)
	local.prefix = local.find(consts.FilePrefix)
	return local
}

// 上传文件到七牛云, name 即为 key
func (local *Local) Upload(localFile, name string) (string, string, error) {
	fileDir := local.ossurl
	if fileDir == "" {

		tmpDir, err := utils.GetRunDir()
		if err != nil {
			return "", "", err
		}
		fileDir = tmpDir
	}
	now := time.Now()
	fileName := strings.Join([]string{local.prefix, now.Format("060101"), now.Format("150405") + "-" + name}, "/")
	fileUrl := config.Config.Host + "/api/v1/files/fetch/?url=" + fileName
	key := fileName
	fileName = strings.Join([]string{fileDir, key}, "/")
	err := utils.CopyFile(localFile, fileName)
	if err != nil {
		niuhe.LogInfo("%v, %v, %v, %v", fileName, key, fileUrl, localFile)
		niuhe.LogInfo("%v", err)
		return "", "", err
	}
	// 用全路径作为 key
	return fileUrl, fileName, nil
}

// 删除本地文件
func (local *Local) Delete(key string) error {
	fileName := key
	if !strings.HasPrefix(key, "/") {
		fileDir := local.ossurl
		if fileDir == "" {
			tmpDir, err := utils.GetRunDir()
			if err != nil {
				return err
			}
			fileDir = tmpDir
		}
		fileName = strings.Join([]string{fileDir, key}, "/")
	}

	err := os.Remove(fileName)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil
}
