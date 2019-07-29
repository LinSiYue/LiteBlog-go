package syserror

type Error404 struct {
	UnKnowError
}

func (ue Error404) Code() int{
	return 1002
}

func (ue Error404) Error() string{
	return "非法访问"
}