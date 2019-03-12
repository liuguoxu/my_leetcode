//翻转字符串中的单词

package main 

import (
    "fmt"
)

func reverse_words(str string )string {
    ll := len(str)

    if ll <2 {
        return str
    }

    if ll==2{
        return str[1]+str[0]
    }

    for p,q,r:=0,1,2;
}

func main() {
    str := "Hello World"

    fmt.Println(reverse_words(str))
}
