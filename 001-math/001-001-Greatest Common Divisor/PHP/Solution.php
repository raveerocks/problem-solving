<?php
	/**
	* @input $a -> integer
	* @input $b -> integer
	*/
	// return an integer //

	function gcd($a, $b){
		//WRITE YOUR CODE HERE
		if($b==0)
		    return $a;
		return gcd($b,$a%$b);

	}
?>
