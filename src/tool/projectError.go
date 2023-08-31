package tool

type ProjectError struct {
	Err string
}

func (e *ProjectError) Error() string {
	return e.Err
}
