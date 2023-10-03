//迴圈break continue
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出
func main ( ){ //建立main 函式，程式進入點
//break
// var x int = 0
// for x<5{
// 	if x==3{
// 		break
// 	}
// 	fmt.Println(x)
// 	x++
// }


//continue
// var x int
// for x=0;x<5;x++ {
// 	if x%2 == 0{// %取餘數，x是偶數
// 		continue
// 	}
// 	fmt.Println(x)
// }


//實際，持續讓使用者輸入數字，計算總何。直到使用者輸入0為止
var result int = 0
for true{
	//無窮迴圈
	var n int
	fmt.Scanln(&n)
	if n==0{
		break
	}
	result += n
}
fmt.Println("總和",result)

}