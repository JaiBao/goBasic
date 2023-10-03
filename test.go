//判斷式
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出
func main ( ){ //建立main 函式，程式進入點
//基本
// if true {
// 	fmt.Println("Go")
// }else{
// 	fmt.Println("Not Go")
// }
//情境
var money int
fmt.Println("請問想領多少錢?")
fmt.Scanln(&money)
if money<100{
	fmt.Println("Too Few")
	}else if money<=100000{
	fmt.Println("OK")
}else{
	fmt.Println("Too Much")
}
fmt.Println("執行完畢")
}