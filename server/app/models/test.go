package models

type Test struct {
	ID int
	name string
}

func GetTest() (Test) {
	return Test{1, "testJSON"}
}