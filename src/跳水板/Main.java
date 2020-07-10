package 跳水板;

/**
 * @author lihaoze01@baidu.com
 * @date 2019/10/10 3:50 下午
 * @description
 **/
public class Main {
    public static void main(String[] args) {
        Solution solution = new Solution();

        for (int i : solution.divingBoard(1, 1, 4)) {
            System.out.print(i + ",");
        }
    }
}
