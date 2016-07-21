package sqapp


type Action func(interface{}) interface{};
type Listener func(interface{});

type application struct{
	connect IModule
}

type IAppContext interface{
	dispatch(Action);
	subscribe(Listener)
}

type IModule interface {

}

//type application struct{
//
//};

func Application() application{
	return application{};
}

//
//
//type application struct {
//	//middlwares Middleware
//
//	config map[string]Config
//}
//
//type IApplicationConfigurator interface {
//	IRunner
//	//Use(middleware Middleware) IApplicationConfigurator
//}
//
//type IRunner interface {
//	Run(port ...string);
//}
//
////func (app *Application) Use(middleware Middleware) IApplicationConfigurator {
////	return app;
////}
//
//func (app *Application) Run(port ...string) {
//}
//
//func App() IApplicationConfigurator {
//	return &Application{}
//}