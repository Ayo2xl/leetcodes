func twoSum(nums []int, target int) []int {
   for i, v:= range nums {
        for j, h:= range nums {
            if i == j { 
                continue 
            }else if v + h == target {
                return []int {i,j}
            }
        }
    }
return []int {}
}