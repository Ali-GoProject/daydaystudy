package main
var a = "Get"

func main() {
   n()
   m()
   n()
}

func n() { print(a) }

func m() {
   a = "O"
   print(a)
}
