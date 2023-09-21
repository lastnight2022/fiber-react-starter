package system

import (
	"server/global"
	"server/model/common/request"
	"server/model/system"
	systemReq "server/model/system/request"
)

type OperationRecordService struct{}

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.ADTPL_DB.Create(&sysOperationRecord).Error
	return err
}

func (operationRecordService *OperationRecordService) DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.ADTPL_DB.Delete(&[]system.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

func (operationRecordService *OperationRecordService) DeleteSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.ADTPL_DB.Delete(&sysOperationRecord).Error
	return err
}

func (operationRecordService *OperationRecordService) GetSysOperationRecord(id uint) (err error, sysOperationRecord system.SysOperationRecord) {
	err = global.ADTPL_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

func (operationRecordService *OperationRecordService) GetSysOperationRecordInfoList(info systemReq.SysOperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.ADTPL_DB.Model(&system.SysOperationRecord{})
	var sysOperationRecords []system.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
