package main

import ("fmt"; "flag")

func main() {
	maxVal := flag.Int("max", 6, "the max value")
	flag.Parse()
	fmt.Println(flag.Args(), *maxVal)
}