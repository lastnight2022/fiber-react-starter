package service

type Service interface {
	Health() bool
	Greeting(name string) string
}

type GreeterService struct {
}

func (GreeterService) Health() bool {
	return true
}

func (GreeterService) Greeting(name string) (greeting string) {
	greeting = "GO-KIT Hello" + name
	return
}
