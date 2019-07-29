package syserror

type UnKnowError struct {
	msg string
	reason error
}

func (ue UnKnowError) Code() int{
	return 1000
}

func (ue UnKnowError) Error() string{
	if len(ue.msg) == 0{
		return "未知错误"
	}
	return ue.msg
}

func (ue UnKnowError) ReasonError() error{
	return ue.reason
}