//陣列基礎
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出


func main ( ){ //建立main 函式，程式進入點
//整數陣列
// var numbers [3] int//三筆資料整數
// numbers[0] = 15 //第一筆資料
// numbers[1] = 30
// numbers[2] = 18

// fmt.Println(numbers)
// fmt.Println(numbers[1]*100)
// //字串陣列
// var names [2] string=[2]string{"john" ,"Marry"}
// fmt.Println(names)

// //取得陣列長度
// fmt.Println(len(numbers))
// fmt.Println(len(names))
//逐一取得陣列中資料

// var grades [5] int=[5]int{60 , 90 ,75 , 80 ,30}
// var sum int=0
// var index int
// for index=0;index<len(grades);index++{
// 	fmt.Println(grades[index])
// 	sum+=grades[index]
// }
// fmt.Println(sum)//總和
// fmt.Println(sum/len(grades))//平均數

// var grades [5] int=[5]int{60 , 90 ,75 , 80 ,30}
// var sum int=grades[0]+grades[1]+grades[2]
// fmt.Println(sum)

//使用者輸入資料
var grades [5] int
var index int
//逐一輸入
fmt.Println("請逐一輸入資料")
for index=0;index<len(grades);index++{
	fmt.Scanln(&grades[index])
}
// 逐一取得陣列中資料
var sum int=0
for index=0;index<len(grades);index++{
	fmt.Println(grades[index])
	sum+=grades[index]
}
fmt.Println(sum)//總和
fmt.Println(sum/len(grades))//平均數
}