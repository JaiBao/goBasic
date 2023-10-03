//函式
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出

func test (msg string ) {//(參數名稱 參數型態)
	fmt.Println(msg)
}

func add (n1 int , n2 int){
var result int =n1+n2
fmt.Println(result)
}

func sum (max int){
var result int = 0
var n int
for n=1;n<=max;n++{
	result+=n
}
fmt.Println(result)
}
func main ( ){ //建立main 函式，程式進入點
//基礎函式語法演練
// test("333" )
// test("444" )
// add(3,2)
// add(5,8)

//計算1+2+3+...+Max 函式包裝
sum(100)
sum(200)
sum(300)
}