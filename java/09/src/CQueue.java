import java.util.Stack;

class CQueue {

    private Stack<Integer> s1;
    private Stack<Integer> s2;

    public CQueue() {
        s1 = new Stack<>();
        s2 = new Stack<>();
    }

    public void appendTail(int value) {
        s1.push(value);
    }

    public int deleteHead() {
        if (s1.empty()) {
            return -1;
        }

        while (!s1.empty()){
            s2.push(s1.pop());
        }

        int head = (int)s2.pop();

        while (!s2.empty()){
            s1.push(s2.pop());
        }
        return head;
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue obj = new CQueue();
 * obj.appendTail(value);
 * int param_2 = obj.deleteHead();
 */