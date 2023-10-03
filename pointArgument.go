//指標參數
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出

// Pass by value
// func add (x  int){
// 	x= x+10
// 	fmt.Println("Add Function", x )
// }

// func main ( ){ //建立main 函式，程式進入點
// var a int =10
// add (a)// Pass by value 獨立
// fmt.Println("Main Function", a)
// }



func add (xPtr  *int){
	*xPtr= *xPtr+10
	fmt.Println("Add Function", *xPtr )
}

func main ( ){ //建立main 函式，程式進入點
var a int =10

// var aPtr *int=&a 
// add (aPtr)// Pass by Pointer 連動

add(&a)// Pass by Pointer 連動
fmt.Println("Main Function", a)


//和使用者要求輸入，運用到指標參數 Pass by pointer
var msg string
fmt.Scanln(&msg)//傳遞字串變數的指標(記憶體位置)
fmt.Println(msg)
}