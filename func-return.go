//函數回傳
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出

func show (msg string){
	if msg =="Hello"{
		return
	}
	fmt.Println(msg)
}

func add (n1 int , n2 int) int {
	var result int = n1 + n2
	// fmt.Println(result)
	return result
}

func test ( ) (int ,string){
	return 5,"test"
}
func main ( ){ //建立main 函式，程式進入點
//基本return運作
// show("Hello")
// show("Hellow")
// show("Hi")

//單一回傳
var x int = add (3,4)
fmt.Println(x*20)

//多個回傳
var x int 
var s string
x,s=test()
fmt.Println(x,s)
}