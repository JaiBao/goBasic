//迴圈
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出
func main ( ){ //建立main 函式，程式進入點
//基本
//簡易
/*var x int = 0
for x<3{
	fmt.Println(x)
	x++
}*/

//複雜
/*var x int
//for 初始化;條件;執行動作
for x=0;x<3;x++{
fmt/Println(x)
}*/


//實際 1+2+3...+50

//簡易
/*var result int = 0
var x int  = 1
for x<=50{
	result =result + x
	x++
}
fmt.Println(result)*/


//複雜

var x int
var result int = 0 
for x =1;x<=50;x++{
result = result + x
}
fmt.Println(result)


}