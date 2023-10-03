//變數
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出
func main ( ){ //建立main 函式，程式進入點
//基本輸出: fmt.Println(資料,資料,資料,...)
// fmt.Println(3,"hello",true)
//基本輸出: fmt.Println(&變數名稱,&變數名稱,&變數名稱,...)
//&變數名稱: 取得變數的指表 (pointer)
// var x int 
// fmt.Print("輸入一個數字")
// fmt.Scanln(&x)
// fmt.Println(x)
// 相乘
var x int 
var y int
//完整
// fmt.Print("輸入第一個數字:")
// fmt.Scanln(&x)
// fmt.Print("輸入第二個數字:")
// fmt.Scanln(&y)

// var result int =x+y
// fmt.Println(result)
//進階
fmt.Print("輸入兩個數字,用空格隔開:")
fmt.Scanln(&x,&y)
var result int =x+y
fmt.Println(result)

}