int gcd(int A, int B) {
    return B==0? A:gcd(B,A%B);
    
}

int lcm(int A, int B){
        return (A/gcd(A,B))*B;
    
}

int Solution::solve(int A, int B, int C) {
    return A/lcm(B,C);
}
