package main

var pc [256]byte 

//pc[i] is the population count of i 

func init() {
	for i:= range pc{
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	
}