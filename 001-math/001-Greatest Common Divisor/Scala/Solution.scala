class Solution {
    def gcd(A: Int, B: Int): Int  = {
        if(B==0)
             return A;
        return gcd(B,A%B);
    }
}
