package examples

import
	"fmt"
func half(x int)(int,bool){
	if x % 2==0{
		fmt.Print("1,true")
		return 1,true
	}
	fmt.Print("0,false")
	return 0,false

}
//поиск максимального значения
func max(args...int)int{
	p_max:=0
	for _,value:= range args{
		if value>p_max{
			p_max=value
		}
	}
	return p_max
}

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 1
		return
	}
}

func fib(n int)int{
	if n==0 {
		return 0;
	}else if n==1{
		return 1;
	}else{
		return fib(n-1)+fib(n-2)
	}
}

func main(){
	half (5)
	fmt.Println(max(5,2,6))
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 0
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 2
	fmt.Print("Фибоначчи:",fib(5))
}