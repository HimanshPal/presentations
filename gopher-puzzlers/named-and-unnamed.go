package main

// START OMIT
type stack []uintptr

func callers() stack {
	return make([]uintptr, 20)
}

func main() {
	callers()
}

// END OMIT
