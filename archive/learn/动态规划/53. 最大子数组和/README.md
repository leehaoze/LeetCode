# 53. 最大子数组和

f(nums) = 
    nums[0] + f(nums[1:]) if nums[0] > 0
    f(nums[1:]) if nums[1] <= 0


dp定义为以i结尾的最大子数组和为dp[i]
那么怎么从dp[i-1] 推到 dp[i]

对于i，要么加入到i-1这个子数组中，要么自己重新开始一个子数组
dp[i] = max(dp[i-1]+nums[i], nums[i]) 