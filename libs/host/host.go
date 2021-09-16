package host

import "os"

func Host() string {
	h, _ := os.Hostname()
	return h
}
