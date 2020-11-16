package tools

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-guns/config"
	"go-guns/global"
	"strconv"
	"time"
)

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type JwtAuth struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// 根据用户的用户名和密码产生token
func GenerateToken(auth JwtAuth) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	timeout := (time.Duration)(config.JwtCfg.Timeout)
	expireTime := nowTime.Add(timeout * time.Second)

	claims := JwtAuth{
		UserId: auth.UserId,
		Role:   auth.Role,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "sting",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 指定加密密钥
	var jwtSecret = []byte(config.JwtCfg.Secret)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*JwtAuth, error) {

	// 指定加密密钥
	var jwtSecret = []byte(config.JwtCfg.Secret)
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtAuth{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*JwtAuth); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}

func GetUserFromClaims(c *gin.Context) JwtAuth {
	if claims, ok := c.Get(global.AUTH_CLAIMS); ok {
		user := (claims).(JwtAuth)
		return user
	}
	return JwtAuth{}
}

func GetUserIdStr(c *gin.Context) string {
	id := GetUserId(c)
	return strconv.Itoa(id)
}

func GetUserId(c *gin.Context) int {
	user := GetJwtAuth(c)
	if user != nil {
		return user.UserId
	}
	return 0
}

func GetJwtAuth(c *gin.Context) *JwtAuth {
	if claims, ok := c.Get(global.AUTH_CLAIMS); ok {
		return (claims).(*JwtAuth)
	}
	return nil
}
