package system

import (
	"go.uber.org/zap"
	"server/config"
	"server/global"
	"server/model/system"
	"server/utils"
)

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.ADTPL_CONFIG
}

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.ADTPL_VP.Set(k, v)
	}
	err = global.ADTPL_VP.WriteConfig()
	return err
}

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.ADTPL_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.ADTPL_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.ADTPL_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
