# Go测试框架使用记录（go convey + go mock + go monkey）
> 本文记录了笔者使用这三个测试框架的过程，还有一些实际应用中遇到的问题。

### [Go convey](https://github.com/smartystreets/goconvey)
> 这是一个辅助库，兼容go原生的单元测试；可以管理和执行测试用例，同时提供了丰富的断言函数，还带有一个web界面。

web界面：在对应目录下执行goconvey，然后打开http://127.0.0.1:8080。

### [Go mock](https://github.com/golang/mock)
> 这是一个官方库，可以完成对一个接口的模拟。

生成mock文件：mockgen -source [x.go] -destination [xx.go]，在x.go中包含接口定义，xx.go则是mock的接口实现：
 - 两个文件的目录均为当前目录；
 - 如果想要将xx.go重定向到其他的目录，该路径需要提前创建。  
 
mock次数：如果接口的一个函数被多次调用，则在mock时需要给出‘次数’参数（e.g. ```mockObj.EXPECT().FuncOne().Return(nil).Times(3)```）

### [Go monkey](https://github.com/bouk/monkey)
> 这个库可以用来对函数、方法和过程打桩。

禁止编译期间的inline优化：go会在编译期间优化一些简短的函数、将其插入到调用位置，所以对inline函数的打桩是无效的；可以在执行测试时添加```-gcflags=-l```参数解决。

---
###### Mario
I've been pretending to work hard, but you're really growing up.
