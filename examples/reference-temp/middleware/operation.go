package middleware

import (
	"bytes"
	"encoding/json"
	"net/url"
	"server/global"
	"server/model/system"
	"server/service"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

type responseBodyWriter struct {
	Response fiber.Response
	body     *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.Response.BodyWriter().Write(b)
}

func OperationRecord() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body []byte
		var userId int
		if c.Request().Header.IsGet() {
			//var err error
			// TODO : fix it
			body = c.Request().Body()
		} else {
			query := string(c.Request().URI().QueryString())
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}

		claims, _ := utils.GetClaims(c)
		if claims.ID != 0 {
			userId = int(claims.ID)
		} else {
			id, err := strconv.Atoi(c.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}
		record := system.SysOperationRecord{
			Ip:     c.IP(),
			Method: c.Method(),
			Path:   c.Path(),
			Agent:  string(c.Request().Header.UserAgent()),
			Body:   string(body),
			UserID: userId,
		}

		writer := responseBodyWriter{
			Response: *c.Response(),
			body:     &bytes.Buffer{},
		}

		// TODO: fix it
		writer.Write(body)

		now := time.Now()
		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Context().Err().Error()
		record.Status = c.Response().StatusCode()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
			global.ADTPL_LOG.Error("create operation record err:", zap.Error(err))
		}
		return nil
	}
}
