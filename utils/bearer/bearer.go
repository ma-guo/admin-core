package bearer

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ma-guo/niuhe"
)

// {"Username":"dequankim","Uid":0,"ExpiresAt":1714562882,"Token":"","aud":"dequankim","exp":1714562882,"nbf":1713957482}
// {"sub":"admin","deptId":1,"dataScope":1,"exp":1713963837,"userId":2,"iat":1713956637,"authorities":["ROLE_ADMIN"],"jti":"ef14c07a6e4e4838bc21b17785e1da4a"}
type Bearer struct {
	Username    string   `json:"-"`
	Uid         int64    `json:"userId"`
	Token       string   `json:"token,omitempty"`
	Role        string   `json:"sub,omitempty"`
	DeptId      int64    `json:"deptId,omitempty"`
	Authorities []string `json:"authorities,omitempty"` // 权限集合
	jwt.StandardClaims
}

func NewBearer(uid int64, username string) *Bearer {
	return &Bearer{
		Username: username,
		Uid:      uid,
	}
}

const secretKey string = "admincore-secret-key"

// 生成 token
func (bea *Bearer) GenToken() error {
	now := time.Now()
	bea.ExpiresAt = now.Add(time.Hour * 24 * 7).Unix()
	bea.StandardClaims = jwt.StandardClaims{
		Audience:  bea.Username,
		ExpiresAt: bea.ExpiresAt,
		NotBefore: now.Unix() - 600,
		IssuedAt:  now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, bea)
	sign, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}
	bea.Token = sign
	return nil
}

// 解析 token
func (bea *Bearer) Parse(tokenStr string) error {
	token, err := jwt.ParseWithClaims(tokenStr, &Bearer{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		// niuhe.LogInfo("%v\n%v", tokenStr, err)
		return err
	}
	tmp, succ := token.Claims.(*Bearer)
	if !succ {
		niuhe.LogInfo("%v", tmp)
		return niuhe.NewCommError(-1, "token claims error")
	}
	if tmp.Valid() != nil {
		return niuhe.NewCommError(-1, "token invalid")
	}
	bea.Username = tmp.Audience
	bea.Uid = tmp.Uid
	bea.ExpiresAt = tmp.ExpiresAt

	return nil
}
