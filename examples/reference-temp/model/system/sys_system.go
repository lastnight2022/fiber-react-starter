package system

import "server/config"

type System struct {
	Config config.Server `json:"config"`
}
