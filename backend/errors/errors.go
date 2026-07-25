package errors

type PublicErrorData struct {
	Code      string
	PublicMsg string
}

type Public interface {
	error
	PublicData() PublicErrorData
}
