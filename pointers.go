package demo

import (
	"fmt"
)

/*
指针保存了变量的内存地址

类型 *T 是指向类型 T 的值的指针。其零值是 nil 。
var p *int

& 符号会生成一个指向其作用对象的指针。
i := 42
p = &i

* 符号表示指针指向的底层的值。
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i

与 C 不同，Go 没有指针运算。
*/


func Point() {
    i,j := 42,2701

    p := &i
    fmt.Println(*p)
    
    *p = 21  
    fmt.Println(i)  

    p = &j
    *p = *p / 37
    fmt.Println(j)
}