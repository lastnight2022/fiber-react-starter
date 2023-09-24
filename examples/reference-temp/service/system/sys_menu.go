package system

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model/common/request"
	"server/model/system"
	"strconv"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) getMenuTreeMap(authorityId string) (err error, treeMap map[string][]system.SysMenu) {
	var allMenus []system.SysMenu
	var btns []system.SysAuthorityBtn
	treeMap = make(map[string][]system.SysMenu)
	err = global.ADTPL_DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	if err != nil {
		return
	}
	err = global.ADTPL_DB.Where("authority_id = ?", authorityId).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]string)
	for _, v := range btns {
		if btnMap[v.SysMenuID] == nil {
			btnMap[v.SysMenuID] = make(map[string]string)
		}
		btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authorityId
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.ID]
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

func (menuService *MenuService) GetMenuTree(authorityId string) (err error, menus []system.SysMenu) {
	err, menuTree := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (menuService *MenuService) GetInfoList() (err error, list interface{}, total int64) {
	var menuList []system.SysBaseMenu
	err, treeMap := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

func (menuService *MenuService) getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[string][]system.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (menuService *MenuService) AddBaseMenu(menu system.SysBaseMenu) error {
	if !errors.Is(global.ADTPL_DB.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.ADTPL_DB.Create(&menu).Error
}

func (menuService *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]system.SysBaseMenu) {
	var allMenus []system.SysBaseMenu
	treeMap = make(map[string][]system.SysBaseMenu)
	err = global.ADTPL_DB.Order("sort").Preload("MenuBtn").Preload("Parameters").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

func (menuService *MenuService) GetBaseMenuTree() (err error, menus []system.SysBaseMenu) {
	err, treeMap := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

func (menuService *MenuService) AddMenuAuthority(menus []system.SysBaseMenu, authorityId string) (err error) {
	var auth system.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []system.SysMenu) {
	err = global.ADTPL_DB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	return err, menus
}
