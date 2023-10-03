//基礎運算
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出
func main ( ){ //建立main 函式，程式進入點
	//算數運算:+, - , * , /
	var x int 
	x= 3 + 3*3
	fmt.Println(x)
	//指定運算: = , += , -= , *=, /=
  x=5
  x *=2
  fmt.Println(x)
	//單元運算: ++ , --
	x++
	fmt.Println(x)
	//比較運算: > , < , >= , <= , ==
	var result bool=4==3
	fmt.Println(result)
	//邏輯運算: ! , && < ||
	// result = !true
	// result = true && false //and,且,&&:兩個都要對才會是true
	result = true || false //or,或,||:只要一邊true就會是true
	fmt.Println(result)

}