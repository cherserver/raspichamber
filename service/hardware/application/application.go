package application

type application struct {
}

func New() *application {
	return &application{}
}

func (a *application) Init() error {
	return nil
}

func (a *application) Stop() {

}
