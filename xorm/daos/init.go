package daos

import (
	"fmt"
	"time"

	"github.com/ma-guo/admin-core/config"

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
func (dao *Dao) GetCache(keyPrefix string, keyArgs ...interface{}) (interface{}, bool) {
	key := keyPrefix
	for _, arg := range keyArgs {
		key += fmt.Sprintf(":%v", arg)
	}
	return localCache.Get(key)
}

func (dao *Dao) SetCache(value interface{}, duration time.Duration, keyPrefix string, keyArgs ...interface{}) {
	key := keyPrefix
	for _, arg := range keyArgs {
		key += fmt.Sprintf(":%v", arg)
	}
	localCache.Set(key, value, duration)
}

var localCache *cache.Cache

func init() {
	localCache = cache.New(5*time.Minute, 30*time.Second)
}
