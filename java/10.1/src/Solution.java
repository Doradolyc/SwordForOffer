import java.util.ArrayList;

class Solution {
    // o(n)
    // o(n)
    public int fib(int n) {
        ArrayList<Integer> list = new ArrayList<>();

        // a[0] = 0
        list.add(0);

        //a[1] = 1;
        list.add(1);

        for(int i = 2; i <= n; i++){
            list.add((list.get(i - 1) + list.get(i - 2)) % 1000000007);
        }

        return list.get(n);
    }
}