package password

import (
	"github.com/ma-guo/niuhe"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	salt string `json:"salt"`
	pwd  string `json:"pwd"`
}

func NewPassword(pwd string) *Password {
	return &Password{
		salt: "admin",
		pwd:  pwd,
	}
}
func (pwd *Password) SetSalt(salt string) {
	pwd.salt = salt
}

func (pwd *Password) SetPwd(password string) {
	pwd.pwd = password
}

func (pwd *Password) Hash() string {
	password := pwd.salt + pwd.pwd
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		niuhe.LogInfo("Hash password error: %v", err)
		return password
	}
	return string(hashedPassword)
}

// ComparePasswords 比较输入的密码与哈希值是否匹配
func (pwd *Password) Compare(hashed string) bool {
	password := pwd.salt + pwd.pwd
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
