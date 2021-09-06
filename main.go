package main
 
import (
    "fmt"
    humanize "github.com/dustin/go-humanize"
    "strings"
    "math"
)
var idr, usd, euro, gbp float64
var menu int

//I'm using package humanize to use its Comma function, it's good to separate the thousands.
func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(math.Round(amount),2) //I'm also using math.Round to give better calculation, this round the float to the nearest integer
	stringValue := strings.Replace(humanizeValue, ",", ".", -1) //This to change the strings of , to ., with unlimited search
  return stringValue}

func FormatForeign(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount,2)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
  return stringValue}

func main(){
  fmt.Println("Welcome to Foreign Exchange menu in Zinedine Bank, please select what do you want to exchange :\n1.\tIndonesia Rupiah to United States Dollar\n2.\tIndonesia Rupiah to Euro\n3.\tGB Pounds to Galleon")
  fmt.Println("Input your menu : ")
  fmt.Scan(&menu)
  if menu == 1{
    fmt.Print("Input your IDR currency :\nRp. ")
    fmt.Scanf("%f",&idr)
    usd = idr * 0.000071
    fmt.Print("IDR : Rp. ", 
    FormatRupiah(idr),",00\n")
    fmt.Println("USD : $ ", FormatForeign(usd))
  } else if menu == 2 {
    fmt.Print("Input your IDR Currency :\nRp. ")
    fmt.Scanf("%f",&idr)
    euro = idr * 0.000060
    fmt.Print("IDR : Rp. ", 
    FormatRupiah(idr),",00\n")
    fmt.Println("EUR : € ", FormatForeign(euro))
  } else if menu == 3 {
    fmt.Print("Input your GB Pounds value :\t")
    fmt.Scanf("%f", &gbp)
    knut := gbp*100
    sickle := knut/29
    galleon := math.Round(sickle/17)
    leftknut := knut - (galleon*17*29)
    leftsickle := math.Floor(leftknut/29)
    fmt.Printf("Your knut(s) value :\t%.0f\n", knut)
    fmt.Printf("Your Galleon(s) value :\t%.0f\n", galleon)
    fmt.Printf("Knut(s) left  :\t%.0f\n", leftknut)
    fmt.Printf("Sickle(s) left :\t%.0f\n", leftsickle)
  } else {
    fmt.Println("Wrong input, please return to the menu selection")
  }
}