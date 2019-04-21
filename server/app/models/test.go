package models

type Test struct {
	ID int `json:"id"`
	name string `json:"name"`
}

type TestCollection struct {
	Tests []Test `json:"items"`
}

func GetTest() (TestCollection) {
	testCollection := TestCollection{
		[]Test {
			{1, "test"},
			{2, "endpoint"},
			{3, "set"},
			{4, "done"},
		},
	}
	return testCollection
}