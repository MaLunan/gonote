## 跨平台编译

SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64

## godep 依赖管理工具

执行以下命令安装godep工具

```go
go get github.com/tools/godep
```

```go
godep save     将依赖项输出并复制到Godeps.json文件中
godep go       使用保存的依赖项运行go工具
godep get      下载并安装具有指定依赖项的包
godep path     打印依赖的GOPATH路径
godep restore  在GOPATH中拉取依赖的版本
godep update   更新选定的包或go版本
godep diff     显示当前和以前保存的依赖项集之间的差异
godep version  查看版本信息

```

## go 基本命令

`go run` 快速执行go文件
`go build`编译程序 生成二进制文件
`go install` 安装可执行文件到bin目录
`go env` 显示go的环境变量
`go fmt` 格式化编码风格

## go module

go module是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，go module将是Go语言默认的依赖管理工具。


## 基础语法

fmt.Println() 打印

fmt.printf(%d,123) 格式化打印
格式符：%o输出八进制，%d输出十进制，%x输出十六进制，%f输出浮点型数据，%c输出单个字符，%s输出字符串，%l输出长整型。但是也有稍微不是很常用的格式符，%p就是其中之一。相信大家在日常中使用得比较少，其实它的输出格式也是六进制，跟%x的区别在于，%p输出的长度是一致的8位16进制符（即32位2进制符）。
%t布尔类型 
cap可以求出slice最大扩张容量，不能超出数组限制
len 长度

## 变量

```go 
var name string
var name string= "ma"
var ( //批量定义
    name string
    age int
)
```
## 常量 

```go
const name string ="ma"
const (
    name string ="ma"
    age int = 18
)
const ( //递增
    num =iota
    num2
    num3
)
```

## 数据类型

```go
//布尔类型 占位符 %t
var a bool =true
var b bool =false

//整数和浮点数 占位符 %d %x十六进制 %f 浮点数
int float 

//字符串  占位符%s 万能占位符 %v
//用"" 或者``
var str string="hello"
str :="hello" //简化声明

var b rune //代表 b是一个utf-8的字符
byte //字节类型 

```

## 字符串操作

```go
//长度 len(str)
//拼接 +,fmt.Sprintf
//分割 string.Split(str,",") //返回一个切片
//包含 string.Contains(str,"str") //返回一个布尔类型
//前缀或后缀判断 string.HasPrefix,strings.HasSuffix
//子串出现的位置 strings.Index(),strings.LastIndex()
//join操作 strings.Join(a[]string,sep string)

```
 

## 包的访问规则

大写意味着这个函数/变量是可导出的
小写是私有的,外部不能访问


## 时间和日期

time包

```go
now:=time.Now()//获取当前时间
Year:=now.Year()//获取当前是那一年
momth:=now.Momth()//获取当前是那一月份
day:=now.Day()//获取当前是那一天

hour:=now.Hour()//获取当前时
minute:=now.Minute()//获取当前分
send:=now.Second()//获取当前秒
now.Unix()//获取当前时间戳
time.tick()//每隔多少时间执行一次的定时器
now.Format("2006/01/02 15:04:05")//格式化日期 必须是2006 go语言诞生日期
```

## for循环


```go
for i:=0;i<=10;i++{
    break //终止循环
    continue //终止本次循环
}
//第一部分和第二部分可以省略
i:=0
for i<10{
    i+=2
}
//可以赋值多个变量
for i:=0,j:=0;i<10&&j<10;i,j=i++,j++{

}
```

## switch语句

```go
switch 条件{
    case 1:
    
    case 2：

    case 3:
    
    case 4:
}

switch a:2; a{
    case 1:
    
    case 2：

    case 3:
        fallthrough //穿透 条件成功继续往下执行
    
    case 4,5,6,7:

    default:
}

```

## 函数

```go
func 函数名 (参数名 参数类型 ) 返回值类型 {

}
func add ( a,b int ) int {//类型一样可以省略

}
func add (a,b int)(int int){//多个返回值

}
func add (a,b int)(sun int sub int){//返回值命名
    sum=a+b  //这里就不用重新命名
    sub=a-b
    return
}
func add (a,b...int)(int int){//多个可变参数
    b[1] 
    b[2]
}
```

## 数组

var 变量名= [长度]类型=(值) 
var 变量名= [...]类型=(值)  //未知长度
var 变量名= [5]类型=(索引:值,索引:值)  //给指定索引赋值

var 变量名= [...]struct=({名称:类型}{值:值})  //前面声明名称类型后面指定值 

``` go
例:var test= [...]struct {
        name string
        age  uint8
    }{
        {"user1", 10}, // 可省略元素类型。
        {"user2", 20}, // 别忘了最后一行的逗号。
    }
```

内置函数 len 和 cap 都返回数组长度 (元素数量)。


### 数组传参

(定义函数内变量名 *[长度]类型)

## 切片 Slice

切片的定义：var 变量名 []类型，

nil 判断切片是否为空

### 通过make来创建切片

```go
    var slice []type = make([]type, len)
    slice  := make([]type, len)
    slice  := make([]type, len, cap)
```

`append` 尾部添加

copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。

## 指针

&（取地址）和*（根据地址取值）

取变量指针的语法如下：

``` go
 ptr := &v    // v的类型为T


```

v:代表被取地址的变量，类型为T
ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。

### new

声明
``` go

func new(Type) *Type

```

1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。


### make

make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型

声明
``` go
func make(t Type, size ...IntegerType) Type
```

### new与make的区别

1.二者都是用来做内存分配的。
2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

## map

map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用

声明

``` go
 map[KeyType]ValueType
```

KeyType:表示键的类型。
ValueType:表示键对应的值的类型。

map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

```go
 make(map[KeyType]ValueType, [cap])
```

也可以在声明时填充元素

```go 
userInfo := map[string]string{
        "username": "pprof.cn",
        "password": "123456",
    }
```


### 判断值是否存在

Go语言中有个判断map中键是否存在的特殊写法，格式如下:

``` go
 value, ok := map[key]

 //例子
 
func main() {
    scoreMap := make(map[string]int)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    // 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
    v, ok := scoreMap["张三"]
    if ok {
        fmt.Println(v)
    } else {
        fmt.Println("查无此人")
    }
}
 
```

### map遍历

```go
scoreMap := make(map[string]int)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    scoreMap["王五"] = 60
for k, v := range scoreMap {
        fmt.Println(k, v)
    }

//或者只遍历key
for k := range scoreMap {
        fmt.Println(k)
    }
```

### map删除

```go
 delete(map, key)
```

map:表示要删除键值对的map
key:表示要删除的键值对的键

## 结构体

### 自定义类型

```go
//将MyInt定义为int类型
type MyInt int
//通过Type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
```

### 类型别名

```go
//类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
type TypeAlias = Type
```

### 结构体

使用type和struct关键字来定义结构体

```go
  type 类型名 struct {
        字段名 字段类型
        字段名 字段类型
        …
    }
```
1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
3.字段类型：表示结构体字段的具体类型。

举个例子，我们定义一个Person（人）结构体，代码如下：

```go
  type person struct {
        name string
        city string
        age  int8
    }

//或者同类型写在一行
  type person1 struct {
        name, city string
        age        int8
    }
```

### 结构体实例化

只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

``` go
  var 结构体实例 结构体类型
```

例子

```go 

type person struct {
    name string
    city string
    age  int8
}

func main() {
    var p1 person
    p1.name = "pprof.cn"
    p1.city = "北京"
    p1.age = 18
    fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
    fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"pprof.cn", city:"北京", age:18}
}
```

### 匿名结构体

```go
 var user struct{Name string; Age int}
```

### 指针结构体

```go
 var p2 = new(person)
```
使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
```go
p3 := &person{}
```


### 结构体初始化

键值对

```go
p6 := &person{
    name: "pprof.cn",
    city: "北京",
    age:  18,
}
```

可以不写键值

```go
p8 := &person{
    "pprof.cn",
    "北京",
    18,
}
```
注意： 1.必须初始化结构体的所有字段。
    2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
    3.该方式不能和键值初始化方式混用

### 创建构造函数

```go
func newPerson(name, city string, age int8) *person {
    return &person{
        name: name,
        city: city,
        age:  age,
    }
}
```

### 方法和接收者

声明
```go
 func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
        函数体
    }
```

指针类型接收者
 1.需要修改接收者中的值
2.接收者是拷贝代价比较大的大对象
3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

```go
func (p *Person) SetAge(newAge int8) {
        p.age = newAge
    }
```

值类型接收者
修改操作只是针对副本，无法修改接收者变量本身。

```go
func (p Person) SetAge2(newAge int8) {
    p.age = newAge
}
```

### 嵌套结构体

```go
//Address 地址结构体
type Address struct {
    Province string
    City     string
}

//User 用户结构体
type User struct {
    Name    string
    Gender  string
    Address Address
}
```