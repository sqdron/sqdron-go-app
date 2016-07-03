package sqapp

type Application struct {
	middlware []interface{}
	config map[string]Config
}

type IApplicationConfigurator interface {
	IRunner
	Use(middleware interface{}) IApplicationConfigurator
}

type IRunner interface {
	Run(port ...string);
}

func (app *Application) Use(middleware interface{}) IApplicationConfigurator {
	app.middlware = append(app.middlware, middleware);
	return app;
}

func (app *Application) Run(port ...string) {
}

func App() IApplicationConfigurator {
	return &Application{}
}