package system

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"server/global"
	"server/model/common/request"
	"server/model/system"
	"server/utils"
)

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.ADTPL_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.ADTPL_DB.Create(&u).Error
	return err, u
}

func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	if nil == global.ADTPL_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.ADTPL_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		var am system.SysMenu
		ferr := global.ADTPL_DB.First(&am, "name = ? AND authority_id = ?", user.Authority.DefaultRouter, user.AuthorityId).Error
		if errors.Is(ferr, gorm.ErrRecordNotFound) {
			user.Authority.DefaultRouter = "404"
		}
	}
	return err, &user
}

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.ADTPL_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.ADTPL_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return err, userList, total
}

func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.ADTPL_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUseAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.ADTPL_DB.Where("uuid = ?", uuid).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.ADTPL_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []system.SysUseAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUseAuthority{
				id, v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (userService *UserService) DeleteUser(id int) (err error) {
	var user system.SysUser
	err = global.ADTPL_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.ADTPL_DB.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
	return err
}

func (userService *UserService) SetUserInfo(reqUser system.SysUser) (err error, user system.SysUser) {
	err = global.ADTPL_DB.Updates(&reqUser).Error
	return err, reqUser
}

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.ADTPL_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return err, reqUser
	}
	var am system.SysMenu
	ferr := global.ADTPL_DB.First(&am, "name = ? AND authority_id = ?", reqUser.Authority.DefaultRouter, reqUser.AuthorityId).Error
	if errors.Is(ferr, gorm.ErrRecordNotFound) {
		reqUser.Authority.DefaultRouter = "404"
	}
	return err, reqUser
}

func (userService *UserService) FindUserById(id int) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.ADTPL_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.ADTPL_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.ADTPL_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.MD5V([]byte("123456"))).Error
	return err
}
