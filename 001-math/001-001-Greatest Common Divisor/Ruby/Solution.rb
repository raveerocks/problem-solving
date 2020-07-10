class Solution
    # @param a : integer
    # @param b : integer
    # @return an integer
    def gcd(a, b)
        if(b==0)
            return a
        end
            return gcd(b,a%b)
    end
end
