package main

import "fmt"

func isPalindrome(num int) string{
   input_num := num
   var rem int
   res := 0
   for num>0 {
      rem = num % 10
      res = (res * 10) + rem
      num = num / 10
   }
   if input_num == res {
      return "Palindrome"
   } else {
      return "Not a Palindrome"
   }
}

func main(){
   fmt.Println(isPalindrome(121))
   fmt.Println(isPalindrome(123))
}
