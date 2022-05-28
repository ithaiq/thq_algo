package algo

import "sort"

//https://leetcode-cn.com/problems/subsets/
//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

//https://leetcode.cn/problems/subsets-ii/
//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//示例 1：
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	sort.Ints(nums) ///
	backtrack2(nums, 0, list, &result)
	return result
}

func backtrack2(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		backtrack2(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

//https://leetcode.cn/problems/permutations/
//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//示例 1：
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := make([]bool, len(nums)) ///
	backtrack3(nums, visited, list, &result)
	return result
}

func backtrack3(nums []int, visited []bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		list = append(list, nums[i])
		visited[i] = true
		backtrack3(nums, visited, list, result)
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}

//https://leetcode.cn/problems/permutations-ii/
//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//示例 1：
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := make([]bool, len(nums))
	sort.Ints(nums) ///
	backtrack4(nums, visited, list, &result)
	return result
}

func backtrack4(nums []int, visited []bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		// 上一个元素和当前相同，并且上一个没有访问过就跳过
		if i != 0 && nums[i] == nums[i-1] && !visited[i-1] { ///
			continue
		}
		list = append(list, nums[i])
		visited[i] = true
		backtrack4(nums, visited, list, result)
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}
