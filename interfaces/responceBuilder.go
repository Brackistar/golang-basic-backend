package interfaces

type ResponseBuilder[T any] interface {
	SetStatusCode(uint)
	SetBody(string)
	AddHeader(string, string)
	Clear()
	Build() *T
}
