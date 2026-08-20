package main
import "fmt"

func main() {
	s := "héllo"
fmt.Println(s[0]) // 104 — the byte 'h', not a rune
for i, r := range s {
    fmt.Printf("%d: %c\n", i, r) // i is byte position, r is rune
}
}