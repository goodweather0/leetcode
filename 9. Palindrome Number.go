func isPalindrome(x int) bool {
    t:=x
    var n int
    if x<0{
        return false
    }
        for{
            a:=t%10
            b:=t/10
            n=n*10+a
            if b==0{
                break
            }
            
           t=b
        }
    
     return n==x
}
