package check

type ResultState uint8

const (
	SUCCESS ResultState = 0
	FAILED  ResultState = 1
)

type Response struct {
	Result    ResultState
	Data      interface{}
	ErrorText string
}

func Success(data interface{}) Response {
	return Response{Result: SUCCESS,
		Data:      data,
		ErrorText: "",
	}
}