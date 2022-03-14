# 518.零钱兑换II
定义f(amount) = count(solutions)

自顶向下
f(amount) = sum(f(amount - range coin))

自底向上
f(0) = 0

```go
for _, coin := range coins {
    f(i) += f(i-coin)
}
```

# 错误case
amount = 5, coins = [1,2,5]

amount = 1 时, solution = 1, count(solution) = 1
amount = 2 时, solution = 1 + 1 | 2, count(solution) = 2
amount = 3 时, solution = 1 + 1 + 1 | 1 + 2 , count(solution) = 2
amount = 4 时, solution = 1 + 1 + 1 + 1 | 1 + 1 + 2 | 2 + 2|
amount = 5 时, solution = 1 + 1 + 1 + 1 + 1 | 1 + 1 + 1 + 2 | 1 + 2 + 2 | 5 |

这个代码会将 1,2 和 2,1 认为是两种组成amount的解.

原先定义的子问题
problem(i) = sum(f(i- range coin)) 这个子问题造成的后果就是上边的将 1,2 和 2,1 认为是两个解，但其实是一个解

正确的子问题应该为
problem(k, i) = problem(k-1, i) + problem(k, i-k)

即