package iteration

func Repeat(str string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += str
	}
	return repeated
}
