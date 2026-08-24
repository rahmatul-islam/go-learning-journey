package main
import "fmt"

func isLeapYear(year int) bool{

	if year%4==0&&year%100!=0||(year%400==0 ){
		return true;
	}else{
		return false;
	}
}
func getinput()int{
	var year int
	fmt.Println("Enter year:")
	fmt.Scan(&year)
	return year
}


func main(){
	year :=getinput()
	fmt.Print(isLeapYear(year))


}