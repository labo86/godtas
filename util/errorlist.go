package util

type ErrorList struct {
	list []error
}

func (o *ErrorList) append(e error) {
	o.list = append(o.list, e)
}

func (o *ErrorList) FirstError() error {
	if len(o.list) == 0 {
		return nil
	}

	return o.list[0]
}
