function gcd(A , B){
	    return B==0?A:gcd(B,A%B);
}
	
function lcm(A , B){
         return ( Math.floor(A/gcd(A,B)))*B;
}
	
module.exports = { 
 //param A : integer
 //param B : integer
 //param C : integer
 //return an integer
	solve : function(A, B, C){
	    return  Math.floor(A/lcm(B,C));
	}
};
