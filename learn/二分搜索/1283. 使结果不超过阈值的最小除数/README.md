解的范围
[1,max(nums)]

# 看了题解

一个是人家计算 nums 中每个数处于一个除数 然后向上取整的方式
我是
```go
		sum += num / t
		if num%t != 0 {
			sum += 1
		}
```

人家是
```go
    sum += (num - 1) / t + 1
```