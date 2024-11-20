func fibonanci (n int) int {
	if n == 0  {
	return 0
		}else  if  n==1 {
		return 1 
		}else {
		return fibonanci (n-1) + fibonanci(n-2)

	 }

}