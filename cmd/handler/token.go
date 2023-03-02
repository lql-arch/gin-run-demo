package handler

import (
	"douSheng/cmd/class"
	"douSheng/sql"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
	"time"
)

type toUserUser struct {
	class.User
	times time.Time
}

// Claims Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var tokenList = make(map[string]toUserUser)
var jwtSecret = []byte("controller.Token")

func init() {
	go DeleteTokenList()
}

// GenerateToken 根据用户的用户名和密码产生token
func GenerateToken(username, password string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// Substring 从开头截取token的长为size内容
func Substring(token string, size int) string {
	var str strings.Builder

	for i, x := range token {
		if i == size {
			break
		}
		str.WriteByte(byte(x))
	}

	return str.String()
}

// FindUserToken 查询用户的token,先在tokenList中查询,如果不存在,就在数据库中查询.
// 缺少实时性,如果使用user,则尽可能使用不能被修改的数据
func FindUserToken(token string) (class.User, bool) {
	user, ok := tokenList[token]

	if ok {
		return user.User, true
	}

	user.User, ok = sql.FindUser(token)
	if !ok {
		log.Println("token does not exist.")

		return user.User, false
	}

	user.times = time.Now().Add(7 * 24 * time.Hour)
	tokenList[token] = user

	return user.User, true
}

// DeleteTokenList 删除存在时间大于一周的token记录
func DeleteTokenList() {
	for k, v := range tokenList {
		if v.times.Before(time.Now()) {
			delete(tokenList, k)
		}
	}

	time.AfterFunc(3*24*time.Hour, DeleteTokenList)
}

// ReplaceToken 对token进行简单的替换
func ReplaceToken(token string) string {
	var str strings.Builder

	for i, x := range token {
		t := x - int32(i)
		if t <= 0 {
			t = x + int32(i)
		}
		str.WriteByte(byte(t))
	}

	return str.String()
}
