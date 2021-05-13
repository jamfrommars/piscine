package piscine

func StrRev(s string) string {
	reversed := ""
	for _, index := range s {
		reversed = string(index) + reversed
	}
	return reversed
}
