package _121买卖股票的最佳时机;

class Solution {
    public int execute(int[] prices, int day, int status, int money) {
        if (day == prices.length - 1) {
            return 0;
        }

        // 可以买入
        if (status == 0) {
            // 买入
            int a = execute(prices, day + 1, 1, money - prices[day]);
            // 不买入
            int b = execute(prices, day + 1, 0, money);

            return Math.max(a, b);
        } else if (status == 1) {
            // 可以卖出
            int a = execute(prices, day + 1, 1, money);
            int b = money + prices[day];
            return Math.max(a, b);
        }

        return 0;
    }


    public int maxProfit(int[] prices) {
        if (null == prices || prices.length <= 1) {
            return 0;
        }

        int dp[] = new int[prices.length];

        dp[0] = 0;
        int minDay = prices[0];

        for (int i = 1; i < prices.length; i++) {
            dp[i] = Math.max(dp[i - 1], prices[i] - minDay);
            minDay = Math.min(minDay, prices[i]);
        }

        return dp[prices.length - 1];
    }
}
