package  main
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main(){

}
func isSymmetric(root *TreeNode) bool {
    return check(root,root)
}
func check(left, right *TreeNode) bool{
    if left == nil && right == nil{
        return true
    }

    if left == nil || right == nil{
        return false
    }
    return left.Val == right.Val && check(left.Right,right.Left) && check(left.Left,right.Right)
}