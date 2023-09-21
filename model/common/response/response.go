package response

import (
	"github.com/gofiber/fiber/v2"
)

const (
	ERROR   = 500
	SUCCESS = 200
)

func Result(code int, msg string, data interface{}, c *fiber.Ctx) {
	c.Status(200).JSON(fiber.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Ok(c *fiber.Ctx) {
	Result(SUCCESS, "操作成功", map[string]interface{}{}, c)
}

func OkWithMessage(message string, c *fiber.Ctx) {
	Result(SUCCESS, message, map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *fiber.Ctx) {
	Result(SUCCESS, "操作成功", data, c)
}

func OkwithDetailed(data interface{}, message string, c *fiber.Ctx) {
	Result(SUCCESS, message, data, c)
}

func Fail(c *fiber.Ctx) {
	Result(ERROR, "操作失败", map[string]interface{}{}, c)
}

func FailWithMessage(message string, c *fiber.Ctx) {
	Result(ERROR, message, map[string]interface{}{}, c)
}

func FailWithDetailed(data interface{}, message string, c *fiber.Ctx) {
	Result(ERROR, message, data, c)
}
