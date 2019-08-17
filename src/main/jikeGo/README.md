## 1.1 基本数据类型

bool

string

int int8 int16 int32 int64

uint uint8 uint16 uint32 uint64 uintptr

byte // alias for unit8

rune // alias for int32, represents a unicode point

float32 float64

complex64 complex128

## 1.2 类型的预定义值
1.math.MaxInt64
2.math.MaxFloat64
3.math.MaxUint32

## 1.3 指针类型

与其它主要的编程语言的差异

1.不支持指针运算
2.string是值类型，其默认的初始化值为空字符串，而不是nil

## 1.4 运算符
没有前置++,--
+,-,*,/,%,++,--

## 1.5 比较运算符

==, !=, >, <, >=, <=

## 1.5.1 比较数组
相同维数且数据值一样

## 1.6 逻辑运算符
&&, ||, !

## 1.7 位运算符
&, |, ^, <<, >>
与其它主要编程语言的差异
&^按位置置零

1 &^ 0 -- 1
1 &^ 1 -- 0
0 &^ 1 -- 0
0 &^ 0 -- 0

# 2.循环结构
Go语言仅支持循环关键字 for

# 2.1 while
while 条件循环
while(n<5)

n := 0
for n < 5 {
   n++
   fmt.Println(n)
}

# 2.2 无限循环
while(true)

n := 0
for {
...
}

# 2.3 if条件

- 1.condition表达式结果必须为布尔值
- 2.支持变量赋值

if condition {

} else {

}

if condition - 1 {

} else if condition - 2 {

} else {

}

# 2.4 switch
1.条件表达式不限制为常量或者整数
2.单个case中，可以出现多个结果选项，使用逗号分隔
3.与C语言等规则相反，Go语言不需要用break来明确推出一个case
4.可以不设定switch之后的条件表达式，在此种情况下，整个switch结构与多个if...else...的逻辑作用等同

