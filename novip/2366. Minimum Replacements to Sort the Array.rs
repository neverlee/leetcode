// 2366. Minimum Replacements to Sort the Array

func minimumReplacement(nums []int) int64 {
    ret := int64(0)
    for i:=len(nums)-2; i>=0; i-- {
        if nums[i] > nums[i+1] {
            part := (nums[i]+nums[i+1]-1) / nums[i+1]
            ret += int64(part-1)
            nums[i] = nums[i]/part
        }
    }
    return ret
}
