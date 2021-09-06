package helper

func ErrorHandling(err error) {
	if err != nil {
		panic(err)
	}
}
