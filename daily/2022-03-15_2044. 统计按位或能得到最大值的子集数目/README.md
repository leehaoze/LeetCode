# 
一开始想了个自上向下递归，但是一直被重复问题困扰，现在细细想一下，
短的子数组的异或结果一定比长的子数组的异或结果小

f([3,2,1,5])

0: 0000
1: 0001
2: 0010
3: 0011
4: 0100
5: 0101
6: 0110

3 
    = 3, maxValue = 3, count = 1
3,2 
    = 3, maxValue = 3, count = 2
    = 3|2 = 3, maxValue = 3, count = 2
3,2,1
    = 3|1 = 3, maxValue = 3, count = 
