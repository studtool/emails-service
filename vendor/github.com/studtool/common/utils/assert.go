package utils

func AssertOk(err error) {
	if err != nil {
		panic(err)
	}
}
