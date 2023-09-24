package system

type RouterGroup struct {
	ApiRouter
	AuthorityRouter
	AuthorityBtnRouter
	BaseRouter
	CasbinRouter
	DictionaryRouter
	DictionaryDetailRouter
	JwtRouter
	MenuRouter
	OperationRecordRouter
	SysRouter
	UserRouter
}
