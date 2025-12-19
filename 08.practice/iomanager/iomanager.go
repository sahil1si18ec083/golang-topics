package iomanager

type IOManager interface {
	WriteResult(data struct{}) error
	ReadFile() ([]string, error)
}
