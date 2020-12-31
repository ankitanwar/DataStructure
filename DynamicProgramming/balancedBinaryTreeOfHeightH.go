package dynamicprogramming

//BalancedBinaryTreeOfHeightH :Given a height of the binary tree we need to tell the number of balanced binary tree possible of given height
func BalancedBinaryTreeOfHeightH(h int) int {
	//we can have 3 types of height of left and right sub tree
	/*1-> (h-1),(h-1)
	  2-> (h-2),(h-1)
	  3-> (h-1),(h-2)

	  so the total binary tree possible would be (h-1)*(h-2)+(h-2)*(h-1)+(h-1)*(h-1)
	*/

	if h == 0 || h == 1 {
		return 1
	}
	return BalancedBinaryTreeOfHeightH(h-1) * (2*BalancedBinaryTreeOfHeightH(h-2) + BalancedBinaryTreeOfHeightH(h-1))

}
