package main

import "fmt"
import "reflect"
import "runtime"

func main(){
  runtime.GOMAXPROCS(2)
  var angka =10
  var kata = "hello"
  go  reflect1(5,angka)
  reflect2(5,kata)

  var input string
  fmt.Scanln(&input )
}

func reflect1(x int,y int){
  var reflectnumber = reflect.ValueOf(y)
  for i:=0;i<x;i++{
    fmt.Println((i+1),reflectnumber.Type())
  }
}

func reflect2(x int,y string){
  var reflectstring = reflect.ValueOf(y)
  for i:=0;i<x;i++{
    fmt.Println((i+1),reflectstring.Type())
  }
}
