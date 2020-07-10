public class Solution {
    
    public int solve(int A, int B, int C) {
        return A/lcm(B,C);
    }
    
    public int gcd(int A, int B) {
        return B==0? A:gcd(B,A%B);
    }
    
    private int lcm(int A, int B){
        return (A/gcd(A,B))*B;
    }
}
