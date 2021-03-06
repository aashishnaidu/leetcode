<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../diagonal-traverse-ii "Diagonal Traverse II")
　　　　　　　　　　　　　　　　
Next >

## [1425. Constrained Subset Sum (Hard)](https://leetcode.com/problems/constrained-subset-sum "带限制的子序列和")

<p>Given an integer array&nbsp;<code>nums</code>&nbsp;and an integer <code>k</code>, return the maximum sum of a <strong>non-empty</strong> subset of that array such that for every&nbsp;two <strong>consecutive</strong> integers in the subset,&nbsp;<code>nums[i]</code>&nbsp;and&nbsp;<code>nums[j]</code>, where&nbsp;<code>i &lt; j</code>, the condition&nbsp;<code>j - i &lt;= k</code>&nbsp;is satisfied.</p>

<p>A&nbsp;<em>subset</em>&nbsp;of an array is&nbsp;obtained by deleting some number of elements (can be&nbsp;zero) from the array, leaving the remaining elements in their original order.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [10,2,-10,5,20], k = 2
<strong>Output:</strong> 37
<b>Explanation:</b> The subset is [10, 2, 5, 20].
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [-1,-2,-3], k = 1
<strong>Output:</strong> -1
<b>Explanation:</b> The subset must be non-empty, so we choose the largest number.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [10,-2,-10,-5,20], k = 2
<strong>Output:</strong> 23
<b>Explanation:</b> The subset is [10, -2, -5, 20].
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= nums.length &lt;= 10^5</code></li>
	<li><code>-10^4&nbsp;&lt;= nums[i] &lt;= 10^4</code></li>
</ul>

### Related Topics
  [[Dynamic Programming](../../tag/dynamic-programming/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Use dynamic programming.
</details>

<details>
<summary>Hint 2</summary>
Let dp[i] be the solution for the prefix of the array that ends at index i, if the element at index i is in the subset.
</details>

<details>
<summary>Hint 3</summary>
dp[i] = nums[i] + max(0, dp[i-k], dp[i-k+1], ..., dp[i-1])
</details>

<details>
<summary>Hint 4</summary>
Use a heap with the sliding window technique to optimize the dp.
</details>
