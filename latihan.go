package main
import "fmt"
func theMaximumAchievableX(num int, t int) int {
    var genap bool
    if num %2 ==0{
        genap = true
    }else{
        genap = false
    }
    hasil := num + t
    if genap == true{
        hasil += 1
    }else{
        hasil+= 2
    }
    return hasil
}
func main(){
	var n, t int
	fmt.Scan(&n, &t)
	hasil:= theMaximumAchievableX(n, t)
	fmt.Print(hasil)
}