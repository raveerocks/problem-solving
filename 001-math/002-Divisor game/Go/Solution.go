/**
 * @input A : Integer
 * @input B : Integer
 * @input C : Integer
 * 
 * @Output Integer
 */
func solve(A int , B int , C int )  (int) {
    return A/lcm(B,C)
}

func gcd(A int, B int) (int){
    if(B==0) {
        return A
    }
    return gcd(B,A%B)
}

 func lcm(A int, B int)(int){
         return (A/gcd(A,B))*B
 }
