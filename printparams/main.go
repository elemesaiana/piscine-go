package main

import (
	"github.com/01-edu/z01"
	"os.Args"
  "fmt"
)

func lastword (s string) string {
	l:=0
  for range s{
  x:=""
  c:=0
  str:=[]rune(s)
  for i:=l-1;i>=0;i--{
  if str[i]==' '{
  c=i
  break
  }
  }
  for i:=c+1;i<l;i++{
  x=x+string(str[i])
  }
	return x
}

func main() {
	args:=os.Args[1:]
  	l:=0
    a:=""
  c:=0
  for range args{
  l++
  }
 if l==1{
 for _,let:= range s{
 a=a+string(lastword(let))
 }
 b:=[]rune(a)
 for i:=range b{
 z01.PrintRune(b[i])
 }
 z01.PrintRune('\n')
}else{
z01.PrintRune('\n')
}
}
