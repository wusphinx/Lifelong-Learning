# [go细节与小技巧101](https://gfw.go101.org/details-and-tips/101.html)
## 关于defer的"反直觉"
> 一个延迟函数调用的实参和被其调用的函数值均是在注册此延迟函数时被**估值**的

且看一例
```golang
package main

func main() {
	var f = func(x int) {
		println(x)
	}

	var n = 1
	defer f(n)

	f = func(x int) {
		println(3)
        n = 2
	}
}
```
实际打印的为**1**，but why?
