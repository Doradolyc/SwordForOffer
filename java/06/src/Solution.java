import java.awt.*;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {

    public class ListNode{
        int val;
        ListNode next;
        ListNode(int x) {
            val = x;
        }
    }

    // o(n)
    // o(n)
    public int[] reversePrint(ListNode head) {
        ArrayList<Integer> list = new ArrayList<Integer>();

        while (head != null) {
            list.add(head.val);
            head = head.next;
        }

        Collections.reverse(list);

        int[] res = new int[list.size()];

        for(int i = 0; i < list.size(); i++){
            res[i] = (int)list.get(i);
        }

        return res;
    }
}