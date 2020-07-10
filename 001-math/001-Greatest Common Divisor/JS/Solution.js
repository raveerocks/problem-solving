module.exports = { 
 //param A : integer
 //param B : integer
 //return an integer
	gcd : function(A, B){
	     return B==0? A:this.gcd(B,A%B);
	}
};
