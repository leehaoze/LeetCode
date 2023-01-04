# 91. 解码方法

## 自上向下

11106 =>
1 1106 11 106 => 11 存在映射 才可以考虑 111 06 => 超过两位不考虑 1110 6 => 超过两位不考虑

f(11106) = f(1) * f(1106)  子问题必须都有解
f(111) + f(06)
f(1110) + f(6)

可以先生成一个 编码存在映射关系的map用来判断子问题是否有解

先写一版递归的把

## 递归

初始化好map后，首先定位分隔符的范围 对于11106，分隔符将字符串分割为 [0, d] 和 (d, len(str)-1]
d属于[0,len(str)-1), 左闭右开

```go
func numDecodings(s string) int {
	length := len(s)
	if length == 0 {
		return 1
	}

	ret := 0
	if length <= 2 {
		if exists[s] {
			ret += 1
		}
	}

	for i := 0; i <= 1 && i < length-1; i++ {
		ss := s[0 : i+1]
		if exists[ss] {
			b := numDecodings(s[i+1:])
			if b == 0 {
				continue
			}
			ret += b
		}
	}

	return ret
}

func main() {
	// fmt.Println(numDecodings("12"))
	// fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("0"))
}

```

这里写了一会才写出来，其实求解的时候，分为两部分，左半部分不用太大，1位或2位即可，并且必须有解，右半部分当做子问题求解
## 备忘录
既然超时了，那就上备忘录，然后就过了
```go
var backup = make(map[string]int)

func numDecodings(s string) int {
	if v, exists := backup[s]; exists {
		return v
	}
	length := len(s)
	if length == 0 {
		return 1
	}

	ret := 0
	if length <= 2 {
		if exists[s] {
			ret += 1
		}
	}

	for i := 0; i <= 1 && i < length-1; i++ {
		ss := s[0 : i+1]
		if exists[ss] {
			b := numDecodings(s[i+1:])
			if b == 0 {
				continue
			}
			ret += b
		}
	}

	backup[s] = ret
	return ret
}
```

## 自下向上

11106
自顶向下是从左往右求解，自下向上我们就从右向左求解

f(6) = exists(6)
f(06) =  exists(0) + f(6) + exists(06)
f(106) = exists(1) + f(06) + exists(10) + f(6)

