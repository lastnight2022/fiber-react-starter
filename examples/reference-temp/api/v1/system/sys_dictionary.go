package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/system"
	"server/model/system/request"
	"server/utils"
)

type DictionaryApi struct{}

func (s *DictionaryApi) CreateSysDictionary(c *fiber.Ctx) error {
	var dictionary system.SysDictionary
	_ = c.BodyParser(&dictionary)
	if err := dictionaryService.CreateSysDictionary(dictionary); err != nil {
		global.ADTPL_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return err
	} else {
		response.OkWithMessage("创建成功", c)
	}
	return nil
}

func (s *DictionaryApi) DeleteSysDictionary(c *fiber.Ctx) error {
	var dictionary system.SysDictionary
	_ = c.BodyParser(dictionary)
	if err := dictionaryService.DeleteSysDictionary(dictionary); err != nil {
		global.ADTPL_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return err
	} else {
		response.OkWithMessage("删除成功", c)
	}
	return nil
}

func (s *DictionaryApi) UpdatesysDictionary(c *fiber.Ctx) error {
	var dictionary system.SysDictionary
	_ = c.BodyParser(&dictionary)
	if err := dictionaryService.UpdateSysDictionary(&dictionary); err != nil {
		global.ADTPL_LOG.Error("更新失败！", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return err
	} else {
		response.OkWithMessage("更新成功", c)
	}
	return nil
}

func (s *DictionaryApi) FindSysDictionary(c *fiber.Ctx) error {
	var dictionary system.SysDictionary
	_ = c.BodyParser(&dictionary)
	if err, sysDictionary := dictionaryService.GetSysDictionary(dictionary.Type, dictionary.ID); err != nil {
		global.ADTPL_LOG.Error("查询失败！", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return err
	} else {
		response.OkwithDetailed(fiber.Map{"resysDictionary": sysDictionary}, "查询成功", c)
	}
	return nil
}

func (s *DictionaryApi) GetSysDictionaryList(c *fiber.Ctx) error {
	var pageInfo request.SysDictionarySearch
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, list, total := dictionaryService.GetSysDictionaryInfoList(pageInfo); err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkwithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
	return nil
}
