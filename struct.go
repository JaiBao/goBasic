//定義、實體化結構
package main //可執行程式必須使用main 封包
import "fmt"//載入內建的 fmt 封包 ，用來做基本輸入輸出

type Person struct {
	name string
	age int

}
func main ( ){ //建立main 函式，程式進入點
	var p1 Person = Person{"阿東" , 18}
	fmt.Println("第一位同學:" , p1.name , p1.age)
	var p2 Person = Person{
		age:46 , name:"Feta",
	}
	p2.name = "餅乾" 
	fmt.Println("第二位同學:" , p2.name , p2.age)
}