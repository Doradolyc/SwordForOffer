import java.util.ArrayList;

class Solution {
    public int numWays(int n) {
        ArrayList<Integer> list = new ArrayList<>();

        // a[0] = 1
        list.add(1);
        //a[1] = 1;
        list.add(1);
        //a[2] = 2;
        list.add(2);

        for(int i = 3; i <= n; i++){
            list.add((list.get(i - 1) + list.get(i - 2)) % 1000000007);
        }

        return list.get(n);
    }
}