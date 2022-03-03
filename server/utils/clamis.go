package utils

import (
	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

func CasdoorGetClaims(token string) (cusClaims *systemReq.CustomClaims, err error) {
	claims, err := auth.ParseJwtToken(token)
	if err == nil {
		id, err := strconv.Atoi(claims.Id)
		if err == nil {
			cusClaims = &systemReq.CustomClaims{}
			cusClaims.ID = uint(id)
			cusClaims.Username = claims.Name
			cusClaims.NickName = claims.DisplayName
			cusClaims.AuthorityId = claims.Tag
		}
	}
	return
}

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	claims, err := CasdoorGetClaims(token) // parse claim
	if err!=nil{
		j := NewJWT()
		claims, err = j.ParseToken(token)
		if err != nil {
			global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
		}
	}
	return claims, err
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.ID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.AuthorityId
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}
