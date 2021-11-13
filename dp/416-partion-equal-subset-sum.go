package dp

import "sort"

func canPartition(nums []int) bool {
	sort.Slice(nums,func(i,j int)bool{
		return nums[i]<nums[j]
	})
	sum := 0
	for _,num:=range nums{
		sum+=num
	}
	if sum % 2==1{
		return false
	}
	dp := make([]int,sum/2+1)
	for i:=0;i<len(nums);i++{
		// 总量weight是j，dp[j]是价值，num[i]是价值，也是weight
		for j:=sum/2; j>=nums[i];j--{
			dp[j] = max(dp[j],dp[j-nums[i]]+nums[i])
		}
	}
	// wight = sum/2
	if dp[sum/2] == sum/2 {
		return true
	}
	return false
}

