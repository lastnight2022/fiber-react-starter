package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

type LogLayout struct {
	Time      time.Time
	Metadata  map[string]interface{}
	Path      string
	Query     string
	Body      string
	IP        string
	UserAgent string
	Error     string
	Cost      time.Duration
	Source    string
}

type Logger struct {
	Filter        func(c *fiber.Ctx) bool
	FilterKeyword func(layout *LogLayout) bool
	AuthProcess   func(c *fiber.Ctx, layout *LogLayout)
	Print         func(layout LogLayout)
	Source        string
}

func (l Logger) SetLoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		path := c.Path()
		query := c.Request().URI().QueryString()
		var body []byte
		if l.Filter != nil && !l.Filter(c) {
			body = c.Body()
			// TODO: fix it 将原body写回
			c.Request().SetBody(body)
		}
		c.Next()
		cost := time.Since(start)
		layout := LogLayout{
			Time:      time.Now(),
			Path:      path,
			Query:     string(query),
			IP:        c.IP(),
			UserAgent: string(c.Request().Header.UserAgent()),
			Error:     strings.TrimRight(c.Context().Err().Error(), "\n"),
			Cost:      cost,
			Source:    l.Source,
		}
		if l.Filter != nil && !l.Filter(c) {
			layout.Body = string(body)
		}
		l.AuthProcess(c, &layout)
		if l.FilterKeyword != nil {
			l.FilterKeyword(&layout)
		}
		l.Print(layout)

		return nil
	}
}

func DefaultLogger() fiber.Handler {
	return Logger{
		Print: func(layout LogLayout) {
			v, _ := json.Marshal(layout)
			fmt.Println(string(v))
		},
		Source: "ADTPL",
	}.SetLoggerMiddleware()
}
