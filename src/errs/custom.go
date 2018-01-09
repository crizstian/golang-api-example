package errs

type Custom interface {
	Result() (string, string, error)
}

type CError struct {
	Status string
	Msg    string
	Err    error
}

func (ce *CError) Result() (string, string, error) {
	return ce.Status, ce.Msg, ce.Err
}

// create new custom error
func NewCustom(s string, m string, e error) *CError {
	return &CError{Status: s, Msg: m, Err: e}
}
