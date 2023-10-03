//指定賦值
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出
func main ( ){ //建立main 函式，程式進入點
	//執行程式時，從這邊開始
	//輸入訊息到終端:fmt.Println(資料,資料,資料,.....)換行
	//輸入訊息到終端:fmt.Print(資料,資料,資料,.....)不換行
	//宣告變數 var 變數名稱 資料型態 = 適當的資料
	var test bool = true 
	var c rune = 'b'
	var s string = "哈嘍朋友"
	var f float64 = 3.55
    x := 42 
    fmt.Println(x) 

    ptr := &x 
    *ptr = 10  
    fmt.Println(x) 
	f = 3.1415
	fmt.Println(f,x,c,s,test)

}