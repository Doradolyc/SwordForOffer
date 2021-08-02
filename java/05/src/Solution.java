class Solution {
    // o(n)
    // o(1)
//    public String replaceSpace(String s) {
//        return s.replaceAll(" ", "%20");
//    }

    // o(n)
    // o(n)
    public String replaceSpace(String s) {
        for(int i = 0; i < s.length(); i++){
            if(s.charAt(i) == ' '){
                s = s.substring(0, i) + "%20" + s.substring(i + 1, s.length());
                i+= 2;
            }
        }
        return s;
    }

    public static void main(String[] args) {
        String s = "We are happy.";

        Solution so = new Solution();
        System.out.println(so.replaceSpace(s));
    }


}