//指標基本
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出

func main ( ){ //建立main 函式，程式進入點
//1.建立存放資料變數
var x int =5
fmt.Println("原來的資料",x)
//2.取得記憶體位置:&變數名稱
fmt.Println("資料的記憶體位置",&x)
//3.存放到指標變數。注意指標資料型態: *資料型態
var xPtr *int = &x
fmt.Println("資料的記憶體位置",xPtr)
//4.反解指標變數: *指標變數名稱
fmt.Println("反解回原來的資料",*xPtr)

// 另外練習
var word string = "Hello"
fmt.Println(word)
var wordPtr *string = &word
fmt.Println(wordPtr)

// fmt.Println(*wordPtr)

var word2 string= *wordPtr
fmt.Println(word2)
}