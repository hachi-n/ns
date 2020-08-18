package pigeon

type Pigeon interface{
	Notify() error
}

func Throw(p Pigeon)  error {
	return p.Notify()
}