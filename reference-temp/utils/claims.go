package utils

import (
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"server/global"
	systemReq "server/model/system/request"
)

func GetClaims(c *fiber.Ctx) (*systemReq.CustomClaims, error) {
	token := c.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.ADTPL_LOG.Error("从fiber的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// 从fiber的Context中获取从jwt解析出来的用户ID

func GetUserID(c *fiber.Ctx) uint {
	claims := c.Locals("claims")
	if claims == nil {
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

// 从fiber的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *fiber.Ctx) uuid.UUID {
	claims := c.Locals("claims")
	if claims == nil {
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

// 从fiber的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *fiber.Ctx) string {
	claims := c.Locals("claims")
	if claims == nil {
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

//  从fiber的Context中获取从jwt解析出来的用户角色属性
func GetUserInfo(c *fiber.Ctx) *systemReq.CustomClaims {
	claims := c.Locals("claims")
	if claims == nil {
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
