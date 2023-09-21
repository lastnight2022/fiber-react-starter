package system

type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	CasbinService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
}
