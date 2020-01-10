 1. go常用工具：
    - go tool compile –S [main.go] >> [main.S] 生成main.go的go汇编，并保存在main.S文件中  
    - go test -coverprofile=coverage.txt 执行当前目录下的测试文件，保存结果为coverage.txt
    - go tool cover -html=coverage.txt 生成网页形式的代码覆盖率报告  
    - go build -gcflags -m [file name] 查看编译过程中，优化了哪些代码
 1. defer：
    - defer只会压一个函数进栈  
    - 如果有一串连续的调用，最后一个函数之前的内容都会立刻执行  
    - 显式调用os.Exit(0)，defer不会被触发  
 1. go，结构体实现接口：
    - var _ Interface = (*interfaceImpl)(nil)，检查结构体是否实现了接口  
    - 指针类型对象，可以使用定义为值类型接收者的方法；反过来不行
    - 同一个接口的不同函数，可以分别使用值接收者和指针接收者实现
      - 一般用结构体实现接口，在实际使用的时候都是使用的接口类型；因为一个接口如果既有值接收者实现的方法、又有指针接收者实现的方法，在实例化的时候，也只能用指针类型结构体，因为值类型不被视为实现了接口
    - 接口的同一个函数，不能同时拥有值接收者和指针接收者的实现，**因为go不支持方法重载**
      - 关于方法重载，go不支持两个同名函数/方法，但一个函数一个方法拥有相同的名字，是可以的
    - 使用：尽量使用指针类型接收者实现接口的全部方法 
 1. map未初始化：
    - 无法写入（panic：nil）
    - 可以执行读操作，但使用v, ok := map[key]时，ok = false, v为对应类型的0值  
 1. go import _ 只执行init()函数，不导入任何内容，包括但不限于类型定义、变量、常量和函数

---
###### Mario
I've been pretending to work hard, but you're really growing up.
