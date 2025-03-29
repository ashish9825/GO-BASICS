package main
import 	"fmt"


 func main(){
	age :=23
	name := "ashish"
	height := 23.342

	fmt.Println("Age",age,"height",height,"name",name)
	fmt.Printf("age is %.3f \n",height)
	fmt.Printf("name is %s \n",name)

	//to print the type of the variable

	fmt.Printf("type of name is %T \n",name)
 }