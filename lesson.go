package main

import "fmt"

/*

Buni darsda topshiriq qilib bergan edingiz
Qo'shimcha qilib shuni qo'shdim hechdan ko'ra . . .

*/

func main() {
  ppp := func (n int) int {
    return n
  }

  Pow := pow(ppp, 4, 2)

  fmt.Println(Pow())
}

func pow(fn func (int) int,num int,  p int) (func () int) {
  theFunc := func() int {
    res := 1
    for p > 0 {
      res *= fn(num)
      p --
    }
    return res
  }
  return theFunc
}
