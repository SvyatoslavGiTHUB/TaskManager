package output

type ResultState uint8

const (
	CORRECT ResultState = 0
)

type Response struct {
	Result    ResultState
	Data      interface{}
	ErrorText string
}

func Correct(data interface{}) Response {
	return Response{Result: CORRECT,
		Data:      data,
		ErrorText: "",
	}
}