/**
 * @input A : Integer
 * @input B : Integer
 * 
 * @Output Integer
 */
func gcd(A int , B int )  (int) {
    if(B==0){
        return A
    }
    return gcd(B,A%B)
}
