package configuration

func Init() {
	// 初始化仓储的接口与实例的对应关系
	RepositoryInit()
	// 初始化 service 的接口与实例的对应关系
	ServiceInit()
}
