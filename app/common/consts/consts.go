package consts

const (
	Authorization  string = "Authorization" // 认证信息关键字 - Authorization
	Bearer         string = "Bearer"        // 认证方式 - Bearer
	FullTimeLayout string = "2006-01-02 15:04:05"
	Sheet1         string = "Sheet1"
	// CodeNoCommRsp  int    = -10000 // 方法内部已经处理了返回, 不再统一返回结果
	AuthError int = 1000 // 认证错误 - 当前页面已失效，请重新登录
	PersError int = 1001 // 权限错误
	// 文件存储相关变量
	FileVendorKey  string = "fileVendor"
	FileCurrentKey string = "fileVendorCurrent"
	FileAccesKey   string = "fileVendorAccessKey"
	FileSecretKey  string = "fileVendorSecretKey"
	FileBucket     string = "fileVendorBucket"
	FileEndpoint   string = "fileVendorEndpoint"
	FileOssurl     string = "fileVendorOssurl"
	FileZone       string = "fileVendorZone"
	FilePrefix     string = "fileVendorPrefix"
)
