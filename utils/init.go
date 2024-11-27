package utils

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ma-guo/niuhe"
)

const passwordSalt = "admin"

var httpMethods = map[int]string{
	niuhe.GET:     http.MethodGet,
	niuhe.POST:    http.MethodPost,
	niuhe.PUT:     http.MethodPut,
	niuhe.DELETE:  http.MethodDelete,
	niuhe.PATCH:   http.MethodPatch,
	niuhe.HEAD:    http.MethodHead,
	niuhe.OPTIONS: http.MethodOptions,
}

// func GenPassword(password string) string {
// 	return password
// }

// HashPassword 使用 Bcrypt 算法生成密码哈希值

// 从字符串中分割出int64数组
func SplitInt64(ids string) []int64 {
	rows := strings.Split(ids, ",")
	result := make([]int64, 0)
	for _, row := range rows {
		if row == "" {
			continue
		}
		id, err := strconv.ParseInt(row, 10, 64)
		if err == nil {
			result = append(result, id)
		}
	}
	return result
}

func GetRunDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		niuhe.LogInfo("%v", err)
		return "", err
	}
	dir := filepath.Dir(exe)
	return dir, nil
}

// 获取当前运行程序的目录
func GetTmpFileName(fileName string) (string, error) {
	dir, err := GetRunDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, fileName), nil
}

func SaveFile(file multipart.File, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}

	return nil
}

func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		niuhe.LogInfo("open src %v", err)
		return
	}
	defer in.Close()

	err = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		niuhe.LogInfo("mkdir %v", err)
		return

	}
	out, err := os.Create(dst)
	if err != nil {
		niuhe.LogInfo("create dst %v", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		niuhe.LogInfo("copy %v", err)
	}
	return
}

func GetHttpMethod() map[int]string {
	return httpMethods
}
