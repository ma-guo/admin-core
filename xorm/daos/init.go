package daos

import (
	"fmt"
	"time"

	"github.com/ma-guo/admin-core/config"

	"github.com/ma-guo/niuhe"
	"github.com/ma-guo/niuhe/db"
	cache "github.com/patrickmn/go-cache"
	"xorm.io/xorm"
)

type Dao struct {
	_db *db.DB
}

func NewDao() *Dao {
	return &Dao{
		_db: db.NewDB(config.MainDB),
	}
}

func (dao *Dao) Close() {
	if dao._db != nil {
		dao._db.Close()
	}
}

func (dao *Dao) db() *xorm.Session {
	if dao._db != nil {
		return dao._db.GetDB()
	}
	// 理论上执行不到这里
	return nil
}

func (dao *Dao) Atom(fn func() error) error {
	if dao._db != nil {
		return dao._db.Atom(fn)
	}
	return nil
}

// 获取缓存
// @param prefix 缓存前缀
// @param args 缓存键值
func (dao *Dao) GetCache(prefix string, args ...any) (any, bool) {
	key := prefix
	for _, arg := range args {
		key += fmt.Sprintf(":%v", arg)
	}
	return localCache.Get(key)
}

// 设置缓存
// @param val 缓存值
// @param prefix 缓存前缀
// @param args 缓存键值
func (dao *Dao) SetCache(val any, prefix string, args ...any) {
	key := prefix
	for _, arg := range args {
		key += fmt.Sprintf(":%v", arg)
	}
	localCache.Set(key, val, 0) // 使用默认设置的过期日期
}

// 删除缓存
// @param prefix 缓存前缀
// @param args 缓存键值
func (dao *Dao) DeleteCache(prefix string, args ...any) {
	key := prefix
	for _, arg := range args {
		key += fmt.Sprintf(":%v", arg)
	}
	localCache.Delete(key)
}

// 查找记录
// @param row 要查找结构
func (dao *Dao) GetBy(row any) (bool, error) {
	has, err := dao.db().Get(row)
	if err != nil {
		return false, err
	}
	return has, nil
}

// 批量插入记录
// @param rows 要插入的数据
func (dao *Dao) Insert(rows ...any) (bool, error) {
	if len(rows) == 0 {
		return false, nil
	}
	session := dao.db()
	if err := session.Begin(); err != nil {
		niuhe.LogInfo("begin error %v", err)
		return false, err
	}
	totalAffected := int64(0)
	// 2000条是经验值, 可根据自己需要更改
	batchSize := 2000
	for i := 0; i < len(rows); i += batchSize {
		end := i + batchSize
		if end > len(rows) {
			end = len(rows)
		}
		affected, err := session.Insert(rows[i:end]...)
		if err != nil {
			niuhe.LogInfo("Insert error: %v， %v， %v", err, i, end)
			session.Rollback()
			return false, err
		}
		totalAffected += affected
	}

	if err := session.Commit(); err != nil {
		niuhe.LogInfo("commit error %v", err)
		return false, err
	}
	return totalAffected > 0, nil
}

// 插入记录, 插入的记录 AUTO_INCREMENT 字段会被填充
// @param row 要插入的字段
func (dao *Dao) InsertOne(row any) (bool, error) {
	session := dao.db()
	if err := session.Begin(); err != nil {
		niuhe.LogInfo("begin error %v", err)
		return false, err
	}
	affected, err := session.InsertOne(row)
	if err != nil {
		session.Rollback()
		return false, err
	}
	if err = session.Commit(); err != nil {
		niuhe.LogInfo("commit error %v", err)
		return false, err
	}
	return affected > 0, err
}

// 更新记录, row 结构中需要包含 id 字段
// @param row 结构中需要包含 id 字段
func (dao *Dao) Update(id int64, row any) (bool, error) {
	session := dao.db()
	if err := session.Begin(); err != nil {
		niuhe.LogInfo("begin error %v", err)
		return false, err
	}
	affected, err := session.Where("`id`=?", id).AllCols().Update(row)
	if err != nil {
		session.Rollback()
		return false, err
	}
	if err = session.Commit(); err != nil {
		niuhe.LogInfo("commit error %v", err)
		return false, err
	}
	return affected > 0, err
}

// 删除记录
// @param row 要删除的空表结构, 需为指针
func (dao *Dao) Delete(ids []int64, row any) (bool, error) {
	if len(ids) == 0 {
		return false, nil
	}
	session := dao.db()
	if err := session.Begin(); err != nil {
		niuhe.LogInfo("begin error %v", err)
		return false, err
	}
	affected, err := session.In("`id`", ids).Delete(row)
	if err != nil {
		session.Rollback()
		return false, err
	}
	if err = session.Commit(); err != nil {
		niuhe.LogInfo("commit error %v", err)
		return false, err
	}
	return affected > 0, err
}

// 分页查找
// @param page 页码
// @param size 每页大小
func (dao *Dao) Limit(session *xorm.Session, page, size int) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	session.Limit(size, (page-1)*size)
}

// like 查找
func (dao *Dao) Like(session *xorm.Session, col string, val any) {
	value := fmt.Sprintf("%v", val)
	if len(col) > 0 && value != "" && value != "0" {
		session.Where(fmt.Sprintf("%v LIKE ?", col), "%"+value+"%")
	}
}

var localCache *cache.Cache

func init() {
	localCache = cache.New(5*time.Minute, 30*time.Second)
}
