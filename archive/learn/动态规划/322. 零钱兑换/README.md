# 322. 零钱兑换

假定我们有一个函数f(), f(amount) = solution， 输入是amount，返回是 凑成amount的最小数量 那么我们可以推出

对于coins = [1, 2, 5]

f(amount) = min(f(amount - 1) + f(amount - 2) + f(amount - 5)) + 1

自顶向下推出来了，然后是自底向上

从f(0) = 0 开始出发
```go
f(i) = min(f(i - 1) + f(i - 2) + f(i - 5)) + 1
```
