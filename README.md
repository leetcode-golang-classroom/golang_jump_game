# golang_jump_game

You are given an integer array `nums`. You are initially positioned at the array's **first index**, and each element in the array represents your maximum jump length at that position.

Return `true` *if you can reach the last index, or* `false` *otherwise*.

## Examples

**Example 1:**

```
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

```

**Example 2:**

```
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

```

**Constraints:**

- `1 <= nums.length <= 104`
- `0 <= nums[i] <= 105`

## 解析

給定一個非負整數陣列 nums

每個元素 nums[i] 代表在 從 index 出發最多能夠 jump 的步數

初始從 index = 0 開始

要求寫一個演算法判斷給定的 nums 能不能 jump 到最後一個 index, 也就是 index = len(nums) - 1

從思考何時會導致無法跑到最後一個 index

當發現有一個 index 從該位找出的 nums[index] + index 

與之前累計的最大能前進的 index 都無法超過目前的 index 

也就是只能到達目前的 index 為止

所以每次只要根據 maxPos = max(maxPos, i + nums[i]) 

來推算已目前的 index 出發能夠到達的最大位置 maxPos

然後檢驗 maxPos 是否 > index 即可知道是否能夠走完到最後的 index

而這樣只要 O(N) 的時間複雜度

每次只要儲存當下 maxPos 所以空間複雜度是 O(1)

![](https://i.imgur.com/qNefmNc.png)

## 程式碼
```go
package sol

func canJump(nums []int) bool {
	nLen := len(nums)
	maxPos := nums[0]
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for pos := 0; pos < nLen-1; pos++ {
		maxPos = max(maxPos, pos+nums[pos])
		if maxPos <= pos {
			return false
		}
	}
	return true
}
```
## 困難點

1. 要找出無法達到最後一個 index 的條件

## Solve Point

- [x]  初始化 maxPos 為 nums[0] 因為從 index 0 的最大 所能到的 index = 0 + nums[0]
- [x]  每次計算 maxPos = max(maxPos, i + nums[i]) 計算最大能達到的位置是否能夠超過當下 index
- [x]  當所有 index 都小於 maxPos 則代表可能回傳 true 否則回傳 false