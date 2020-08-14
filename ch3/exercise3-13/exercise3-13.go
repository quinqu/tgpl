// write const declarations for kB, MB, through YB as compactly as you can
package main 
import (
	"fmt"
)


const (
	KB = 1000 
	GB = KB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000 
)


func main () {
	fmt.Println("powers")
	fmt.Println(KB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(float32(YB))
}