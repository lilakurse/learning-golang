package examples

import
	"fmt"

func massive(){
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	min:=48
	for _,value:=range x{
		if min>value{
			min=value
		}
	}
	fmt.Print(min)
}

