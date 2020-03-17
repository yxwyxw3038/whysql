## random
mysql SQL语句自动生成器

## 用例
```go
package main

import (
	"fmt"
	"github.com/gohouse/random"
)

func main() {
	fmt.Println(random.Rand())
	fmt.Println(random.Random(12))
	fmt.Println(random.Random(12,random.T_ALL))
	fmt.Println(random.RandomBetween(6, 11))
	fmt.Println(random.RandomBetween(6,11, random.T_ALL))
}
```
输出
```shell script
KNgUdYQxeOmZSLDZQAcQOYGGNeeAUa
yhbsh85wzj3o
0sllQyYHxm4p
e0ehi7
oHrf1S8kB
```

## api说明
- func Rand() string  
随机生成6-32位的随机字符串(长度类型皆随机)  

- Random(length int, fill ...RandType) string  
随机生成指定长度的随机字符串(类型可选或随机)  

- RandomBetween(min, max int, fill ...RandType) string  
随机生成指定长度区间的随机字符串(类型可选或随机)


## 字符组成
```go
// RandType ...
type RandType int

const (
	// 大写字母
	T_CAPITAL RandType = iota + 1
	// 小写字母
	T_LOWERCASE
	// 数字
	NUMBERIC
	// 小写字母+数字
	T_LOWERCASE_NUMBERIC
	// 大写字母+数字
	T_CAPITAL_NUMBERIC
	// 大写字母+小写字母
	T_CAPITAL_LOWERCASE
	// 数字+字母
	T_ALL
)
```
字符串
```go
StrNumberic  = `0123456789`
StrLowercase = `abcdefghijklmnopqrstuvwxyz`
StrCapital   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
```
