# 思路
首先确定解 `运载能力` 的范围。最大值是所有包裹的重量和`sum(weights)`，这样一次可以将所有包裹运输走，最小是所有包裹中最重的一个包裹的重量`max(weights)`，这样运输的次数是最多的。
即解的范围为 `[max(weights), sum(weights)]`
在这个范围内搜索一个左边界，使得 ship(weights, mid) <= days

