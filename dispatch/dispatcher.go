package dispatch

type dispatcher struct {

}

func Dispatcher() *dispatcher {
	return &dispatcher{}
}

func (d *dispatcher) Process(path, method, action string, args *map[string]string) (interface{}, *HttpError) {
	return nil, &HttpError{}
}

type HttpError struct {
	Code int
	message string
}

func (he HttpError) Error() string {
	return he.message
}