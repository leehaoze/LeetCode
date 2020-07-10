package _309最佳买卖股票时机含冷冻期;

import static java.lang.Integer.max;

class Solution {
    private static final int INIT = 0;
    private static final int BUY = 1;
    private static final int FREEZE = 2;

    public int operate(int[] prices, int day, int status, int money) {
        // 到了最后一天还没有卖出
        if (day > prices.length - 1) {
            return money;
        }

        switch (status) {
            case INIT:
                money = max(
                        // 什么也不做
                        operate(prices, day + 1, status, money),
                        // 买入
                        operate(prices, day + 1, status + 1, money - prices[day])
                );
                break;
            case BUY:
                money = max(
                        operate(prices, day + 1, status, money),
                        operate(prices, day + 1, status + 1, money + prices[day])
                );
                break;
            case FREEZE:
                money = operate(prices, day + 1, INIT, money);
                break;
        }

        return money;
    }


    public int maxProfit(int[] prices) {
        /**
         * 初始状态  可以 买入 无操作 买入后转入 买入状态
         * 买入状态  可以 卖出 无操作 卖出后转入 冷冻状态
         * 冷冻状态  可以 无操作 过一天后转入初始状态
         */
        if (prices == null || prices.length == 0) {
            return 0;
        }

        // 二位数组代表 第i天 持有的钱数 以及 处于的状态
        int[][] dp = new int[prices.length][3];


        // 第一天 持股
        dp[0][0] = 0;
        // 第一天 不持股
        dp[0][1] = -prices[0];


        for (int i = 1; i < prices.length; i++) {
            // 第 i 天 进入不持股状态, 昨天是冷冻期，或者昨天就是不持股
            dp[i][0] = max(dp[i - 1][2], dp[i - 1][0]);
            // 第 i 天 进入持股状态, 昨天是不持股，或者昨天是冷冻期，今天解除了
            dp[i][1] = max(dp[i - 1][0] - prices[i], dp[i - 1][1]);
            // 第 i 天 进入冷冻期状态, 昨天持股，今天卖出
            dp[i][2] = dp[i - 1][1] + prices[i];
        }

        return max(dp[prices.length - 1][0], dp[prices.length - 1][2]);
    }
}
