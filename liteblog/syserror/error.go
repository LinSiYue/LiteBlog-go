package syserror

type Error interface {
	Code() int
	Error() string
	ReasonError() error
}

func NewError(msg string, reason error) Error{
	return &UnKnowError{
		msg : msg, 
		reason : reason,
	}
}