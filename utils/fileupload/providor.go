package fileupload

// 文件存储提供商
type Provider interface {
	// 上传文件到 oss, 返回 url 和 key
	Upload(localFile, name, fileType string) (string, string, error)
	// 删除文件
	Delete(key string) error
}

type dictDao struct {
	data map[string]string
}

func newDict(data map[string]string) *dictDao {
	return &dictDao{
		data: data,
	}
}

func (dao *dictDao) find(name string) string {
	if value, ok := dao.data[name]; ok {
		return value
	}
	return ""
}
