class Solution {
    // o(n + m)
    // o(1)
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0) {
            return false;
        }

        int rows = matrix.length;
        int columns = matrix[0].length;
        int num = 0;

        int r = 0, c = columns - 1;

        while (r < rows && c >= 0){
            num = matrix[r][c];

            if(num == target) {
                return true;
            }else if (num > target) {
                c--;
            }else {
                r++;
            }
        }

        return false;
    }
}