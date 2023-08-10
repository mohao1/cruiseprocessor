package cruiseprocessor

// Polling Scheduling algorithm(Polling)
func Polling(chanId int, workNum int) (int, error) {
	return chanId % workNum, nil
}
