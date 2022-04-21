package util

// Runner runs something.
// It's designed to be created by DI and loaded in command.
type Runner interface {
	Run() (err error)
}

type RunnerFunc func() (err error)

func (f RunnerFunc) Run() (err error) {
	return f()
}
