package entity

type Task struct {
	Id   int64
	Name         string
	IdProject    int64
	Description  string
	Time         int64 
	Priority	 string
	Status       string
	TypeTask     string

}
