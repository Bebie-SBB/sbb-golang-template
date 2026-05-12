package interfaces

type IModule interface {
	Init() error
	SetRoutes() error
}
