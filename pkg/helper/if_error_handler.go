package helper

func IfErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
