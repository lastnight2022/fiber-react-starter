package system

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model/common/request"
	"server/model/system"
)

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.ADTPL_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.ADTPL_DB.Create(&api).Error
}

func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	err = global.ADTPL_DB.Delete(&api).Error
	CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err
}

func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.ADTPL_DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

func (apiService *ApiService) GetAllApis() (err error, apis []system.SysApi) {
	err = global.ADTPL_DB.Find(&apis).Error
	return
}

func (apiService *ApiService) GetApiById(id int) (err error, api system.SysApi) {
	err = global.ADTPL_DB.Where("id = ?", id).First(&api).Error
	return
}

func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.ADTPL_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.ADTPL_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.ADTPL_DB.Save(&api).Error
		}
	}
	return err
}

func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	err = global.ADTPL_DB.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
	return err
}

func (apiService *ApiService) DeleteApiByIds(ids []string) (err error) {
	return global.ADTPL_DB.Delete(&system.SysApi{}, "id in ?", ids).Error
}
