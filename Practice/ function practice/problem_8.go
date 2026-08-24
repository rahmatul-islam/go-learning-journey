package main
import"fmt"

func getinput()(int ,int,int){

	var a,b,c int 
	fmt.Print("Enter first Number:")
	fmt.Scan(&a)
	fmt.Print("Enter second Number:")
	fmt.Scan(&b)
	fmt.Print("Enter third Number:")
	fmt.Scan(&c)

	return a,b,c
}

func BiggerNumber(a,b,c int)( int ){
	if a>b&&a>c {
	return a
		
	}else if a<b&&b>c {
	 return b
		
	}else{
		return c
		
	}
}

func main(){
	a,b,c :=getinput();
	fmt.Println("Big Number is:",BiggerNumber(a,b,c))

}