/**
 * @input A : Integer
 * @input B : Integer
 * 
 * @Output Integer
 */
int gcd(int A, int B) {
    if(B==0)
        return A;
    return gcd(B,A%B);
}
