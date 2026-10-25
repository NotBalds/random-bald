package main

import (
	"golang.org/x/sys/unix"
)

func itoa(val int) string {
	if val == 0 {
		return "0"
	}
	var buf [32]byte
	i := len(buf) - 1
	ng := val < 0
	if ng {
		val = -val
	}
	for val > 0 {
		buf[i] = byte(val%10) + '0'
		i--
		val /= 10
	}
	if ng {
		buf[i] = '-'
		i--
	}
	
	return string(buf[i+1:])
}

func printText(ts unix.Timespec, text string) {
	for _, i := range text {
		unix.Write(1, []byte(string(i)))
		unix.Nanosleep(&ts, nil)
	}
	unix.Write(1, []byte("\n"))
}

func main() {
	ts1 := unix.NsecToTimespec(1000000 * 600)
	tsp := unix.NsecToTimespec(1000000 * 60)
	printText(tsp, "Random bald generator!")
	unix.Nanosleep(&ts1, nil)
	unix.Write(1, []byte("\n"))
	unix.Write(1, []byte("\n"))

	ur, _ := unix.Open("/dev/urandom", unix.O_RDONLY, 0)
	var bald [5]int
	var res int = 0
	for i := range bald {
		var buf []byte = make([]byte, 1)
		_, _ = unix.Read(ur, buf);
		bald[i] = int(buf[0]) % 100
		res += bald[i]
	}

	printText(tsp, "You are " + itoa(bald[0]) + "% bald")
	printText(tsp, "Your sister is " + itoa(bald[1]) + "% bald")
	printText(tsp, "Your brother is " + itoa(bald[2]) + "% bald")
	printText(tsp, "Your dog is " + itoa(bald[3]) + "% bald")
	printText(tsp, "Your cat is " + itoa(bald[4]) + "% bald")

	printText(tsp, "My rating: " + itoa(res) + "%!")
	unix.Write(1, []byte("\n"))
	unix.Write(1, []byte("\n"))

	unix.Close(ur)
}
