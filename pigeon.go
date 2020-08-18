package pigeon

type Pigeon interface{
	Notify() error
}

func Deliver(p Pigeon)  error {
	return p.Notify()
}