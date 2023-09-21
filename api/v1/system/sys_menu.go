package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"
)

type AuthorityMenuApi struct{}

func (a *AuthorityMenuApi) GetMenu(c *fiber.Ctx) error {
	if err, menus := menuService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
		global.ADTPL_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return err
	} else {
		if menus == nil {
			menus = []system.SysMenu{}
		}
		response.OkwithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
	return nil
}

func (a *AuthorityMenuApi) GetBaseMenuTree(c *fiber.Ctx) error {
	if err, menus := menuService.GetBaseMenuTree(); err != nil {
		global.ADTPL_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return err
	} else {
		response.OkwithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
	}
	return nil
}

func (a *AuthorityMenuApi) AddMenuAuthority(c *fiber.Ctx) error {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	_ = c.BodyParser(authorityMenu)
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.ADTPL_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return err
	} else {
		response.OkWithMessage("添加成功", c)
	}
	return nil
}

func (a *AuthorityMenuApi) GetMenuAuthority(c *fiber.Ctx) error {
	var param request.GetAuthorityId
	_ = c.BodyParser(&param)
	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, menus := menuService.GetMenuAuthority(&param); err != nil {
		global.ADTPL_LOG.Error("获取成功!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
		return err
	} else {
		response.OkwithDetailed(fiber.Map{"menus": menus}, "获取成功", c)
	}
	return nil
}

func (a *AuthorityMenuApi) AddBaseMenu(c *fiber.Ctx) error {
	var menu system.SysBaseMenu
	_ = c.BodyParser(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := menuService.AddBaseMenu(menu); err != nil {
		global.ADTPL_LOG.Error("添加失败！", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return err
	} else {
		response.OkWithMessage("添加成功", c)
	}
	return nil
}

func (a *AuthorityMenuApi) DeleteBaseMenu(c *fiber.Ctx) error {
	var menu request.GetById
	_ = c.BodyParser(&menu)
	if err := utils.Verify(menu, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := baseMenuService.DeleteBaseMenu(menu.ID); err != nil {
		global.ADTPL_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return err
	} else {
		response.OkWithMessage("删除成功", c)
	}
	return nil
}

func (a *AuthorityMenuApi) UpdateBaseMenu(c *fiber.Ctx) error {
	var menu system.SysBaseMenu
	_ = c.BodyParser(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
		global.ADTPL_LOG.Error("更新失败！", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return err
	} else {
		response.OkWithMessage("更新成功", c)
	}
	return nil
}

func (a *AuthorityMenuApi) GetBaseMenuById(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, menu := baseMenuService.GetBaseMenuById(idInfo.ID); err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return err
	} else {
		response.OkwithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
	return nil
}

func (a *AuthorityMenuApi) GetMenuList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, menuList, total := menuService.GetInfoList(); err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return err
	} else {
		response.OkwithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
	return nil
}
