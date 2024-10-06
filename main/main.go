package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"sort"
	"test01/main/router"
	"time"
)

func timestamp(t int) string {
	fmt.Println(t)
	unix := time.Unix(int64(t), 0)
	return unix.Format("2006-01-02 15:04:05")
}
func stringCat(t1 string, t2 string) string {
	fmt.Println("sc")
	return t1 + t2
}

func main() {
	fmt.Println("123")
	r := gin.Default()
	r.Static("/staticFiles", "./staticFiles")
	//创建存储引擎 密钥
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	//r.Use(middleware.Init, middleware.Init2)
	r.SetFuncMap(template.FuncMap{
		"UnixTime":  timestamp,
		"stringCat": stringCat,
	})
	r.LoadHTMLGlob("template/**/*")
	router.AdminGroup(r)
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Ints(candidates)
	var a []int
	dfs := func(num, now int) {}
	dfs = func(num, now int) {
		if now == 0 {
			ans = append(ans, a)
			return
		}
		if num == len(candidates) || candidates[num] > now {
			return
		}
		a = append(a, candidates[num])
		fmt.Println(a)
		dfs(num, now-candidates[num])
		a = a[:len(a)-1]
		dfs(num+1, now)
	}
	dfs(0, target)

	return ans
}

func minWindow(s string, t string) string {
	m := make(map[byte]int)
	m1 := make(map[byte]int)
	ans := ""
	for i := 0; i < len(t); i++ {
		m[t[i]]++
		m1[t[i]]++
	}
	cnt := 0
	r := 0
	for i := 0; i < len(s) && r < len(s); {
		if m[s[r]] > 0 || m1[s[r]] > 0 {
			if m[s[r]] <= 0 && m1[s[r]] > 0 {
			} else {
				cnt++
			}
			m[s[r]]--
			for cnt == len(t) {
				for m1[s[i]] == 0 {
					i++
				}
				// fmt.Println(s[i : r+1])
				if r-i+1 < len(ans) || ans == "" {
					ans = s[i : r+1]
				}
				m[s[i]]++
				if m[s[i]] > 0 {
					cnt--
				}
				i++
			}
		}
		r++
	}
	return ans
}
func maxScoreSightseeingPair(values []int) int {
	cnt := values[0] - 1
	ans := 0
	for i := 1; i < len(values); i++ {
		tmp := cnt + values[i]
		ans = max(ans, tmp)
		cnt = max(cnt, values[i]) - 1
	}
	return ans
}
func maximumSubsequenceCount(text string, pattern string) int64 {
	ans := int64(0)
	cnt := 1
	cnt2 := 0
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[1] {
			ans += int64(cnt)
			cnt2++
		}
		if text[i] == pattern[0] {
			cnt++
		}
	}
	return ans + int64(max(0, cnt-1-cnt2))
}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			return []int{m[target-nums[i]], i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
func distinctNames(ideas []string) int64 {
	ans := int64(0)
	m := make([]map[string]bool, 26)
	for i := 0; i < 26; i++ {
		m[i] = make(map[string]bool)
	}
	for i := 0; i < len(ideas); i++ {
		m[ideas[i][0]-'a'][ideas[i][1:]] = true
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[:i]); j++ {
			fmt.Println(m[:i])
			cnt := 0
			for s := range m[i] {
				if m[:i][j][s] {
					cnt++
				}
			}
			ans += int64((len(m[i]) - cnt) * (len(m[:i][j]) - cnt))
		}

	}

	return ans * 2
}

func validSubstringCount(s string, t string) int64 {
	m := make(map[byte]int)
	m1 := make(map[byte]int)
	ans := int64(0)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
		m1[t[i]]++
	}
	cnt := 0
	r := 0
	for i := 0; r < len(s); {
		if m[s[r]] > 0 || m1[s[r]] > 0 {
			if m[s[r]] <= 0 && m1[s[r]] > 0 {
			} else {
				cnt++
			}
			m[s[r]]--
			for cnt == len(t) {
				if m[s[i]] == 0 && m1[s[r]] > 0 {
					cnt--
				}
				m[s[i]]++
				i++
			}
			fmt.Println(i)
			ans += int64(i)
		}
		r++
	}
	return ans
}
func takeCharacters(s string, k int) int {
	cnt := make([]int, 3)
	ans := 0
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1
	}
	l, r := 0, 0
	for r < len(s) {
		cnt[s[r]-'a']--
		for l < r && (cnt[s[0]-'a'] < k || cnt[s[1]-'a'] < k || cnt[s[2]-'a'] < k) {
			cnt[s[l]-'a']++
			l++
		}
		if cnt[0] >= k && cnt[1] >= k && cnt[2] >= k {
			ans = min(ans, len(s)-(r-l+1))
		}
		r++
	}
	return ans
}
func permute(nums []int) [][]int {

	var ans [][]int
	dfs := func(now []int) {}
	dfs = func(now []int) {
		if len(now) == len(nums) {
			ans = append(ans, append([]int(nil), now...))
		}
		for i := 0; i < len(nums); i++ {
			if nums[i] != -11 {
				now = append(now, nums[i])
				nums[i] = -11
				dfs(now)
				nums[i] = now[len(now)-1]
				now = now[:len(now)-1]
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		nums[i] = -11
		dfs([]int{num})
		nums[i] = num
	}
	return ans
}
func solve(board [][]byte) {
	dfs := func(x, y int) {}
	dfs = func(x, y int) {
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
			return
		}
		if board[x][y] == 'X' {
			return
		}
		board[x][y] = '1'
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[len(board)-1][j] == 'O' {
			dfs(len(board)-1, j)
		}
	}
	for i := 1; i < len(board)-1; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			dfs(i, len(board[0])-1)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '1' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}
func timeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	for i := 0; i <= k; i++ {
		if tickets[i]-tickets[k] >= 0 {
			ans += tickets[k]
		} else {
			ans += tickets[i]
		}

	}
	for i := k + 1; i < len(tickets); i++ {
		if tickets[i]-tickets[k]+1 >= 0 {
			ans += tickets[k] - 1
		} else {
			ans += tickets[i]
		}
	}
	return ans
}
func kthCharacter(k int) byte {
	var n []byte
	//s := "a"
	n = append(n, 'a')
	for len(n) >= k {
		tmp := len(n)
		for i := 0; i < tmp; i++ {
			if n[i]-'a' == 26 {
				n = append(n, 'a')
			} else {
				n = append(n, n[i]+1)
			}

		}
	}
	return n[k]
}
func isTrue(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	}
	return false
}
func countOfSubstrings(word string, k int) int {
	m := make(map[byte]int)
	cnt := 0
	l, r := 0, 0
	ans := 0
	for r < len(word) {
		if isTrue(word[r]) {
			m[word[r]]++
		} else {
			cnt++
		}
		for cnt > k {
			if isTrue(word[l]) {
				m[word[l]]--
			} else {
				cnt--
			}
			l++
		}
		if cnt == k && m['a'] > 0 && m['e'] > 0 && m['i'] > 0 && m['o'] > 0 && m['u'] > 0 {
			cnt1, cnt2 := 1, 1
			if isTrue(word[l]) && isTrue(word[r]) {
				for i := l + 1; i < r; i++ {
					if isTrue(word[i]) && m[word[i]] > 1 {
						cnt1++
					} else {
						break
					}
				}
				for i := r - 1; i > l; i-- {
					if isTrue(word[i]) {
						cnt2++
					} else {
						break
					}
				}
			}
			ans += cnt1 + cnt2
		}

		r++
	}

	return ans
}
func canReach(s string, minJump int, maxJump int) bool {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		dp[i] = dp[i-1]
		if i >= minJump && s[i] == '0' && (dp[i-minJump+1] > dp[max(i-maxJump, 0)]) {
			dp[i]++
		}
	}
	return dp[len(s)] > dp[len(s)-1]
}
func exist(board [][]byte, word string) bool {
	dfs := func(x, y, n int) bool { return true }
	dfs = func(x, y, n int) bool {
		if n >= len(word) && x < 0 || y < 0 || x >= len(board)-1 || y >= len(board[0])-1 {
			return false
		}
		if board[x][y] == word[n] {
			if n == len(word)-1 {
				return true
			}
			return dfs(x, y+1, n+1) || dfs(x-1, y, n+1) || dfs(x, y-1, n+1) || dfs(x+1, y, n+1)
		}
		return false
	}
	fmt.Println()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func minimumTime(time []int, totalTrips int) int64 {
	l, r := int64(0), int64(time[0])*int64(totalTrips)

	for l < r {
		mid := l + (r-l)/2
		sum := int64(0)
		for i := 0; i < len(time); i++ {
			sum += mid / int64(time[i])
		}
		if sum >= int64(totalTrips) {
			r = mid
			//ans = min(ans, int64(r))
		} else {
			l = mid + 1
		}

	}
	return l
}
func findAnagrams(s string, p string) []int {
	//m := make(map[byte]int)
	cnt := 0
	m := [26]int{0}
	var ans []int
	for i := 0; i < len(p); i++ {
		if m[p[i]-'a'] == 0 {
			cnt++
		}
		m[p[i]-'a']++
	}
	l, r := 0, 0
	n := 0

	for ; r+len(p) < len(s); r++ {
		m[s[r]]--
		if m[s[r]-'a'] == 0 {
			n++
		}
		for r-l+1 > len(p) {
			m[s[l]]++
			if m[s[l]] == 1 {
				n--
			}
			l++
		}
		fmt.Println(l, r, n)
		if n == cnt && r-l+1 == len(p) {
			ans = append(ans, l)
			m[s[l]]++
			if m[s[l]] == 1 {
				n--
			}
			l++
		}
	}
	return ans
}
func orangesRotting(grid [][]int) int {

	dfs := func(x, y, n int) {

	}
	dfs = func(x, y, n int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return
		}
		if grid[x][y] == 1 {
			grid[x][y] = n
			dfs(x+1, y, n-1)
			dfs(x, y+1, n-1)
			dfs(x, y-1, n-1)
			dfs(x-1, y, n-1)
		}
		if grid[x][y] < 0 {
			grid[x][y] = max(n, grid[x][y])
			dfs(x+1, y, grid[x][y]-1)
			dfs(x, y+1, grid[x][y]-1)
			dfs(x, y-1, grid[x][y]-1)
			dfs(x-1, y, grid[x][y]-1)
		}

	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				grid[i][j] = -1
				dfs(i, j, -1)
			}

		}
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = min(ans, grid[i][j])
		}
	}
	if ans == 0 {
		return -1
	}
	return ans * -1
}
func canCompleteCircuit(gas []int, cost []int) int {

	for i := 0; i < len(gas); i++ {
		if gas[i]-cost[i] < 0 {
			continue
		}
		sum := 0
		j := i
		for ; ; j++ {
			j = j % len(gas)
			sum += gas[j] - cost[j]
			// fmt.Println(i, j, sum)
			if sum < 0 {
				break
			}
			if j == i-1 || (i-1 < 0 && j == len(gas)-1) {
				return i
			}

		}
		if i == len(gas)-1 {
			return -1
		}
		i = j - 1
	}
	return -1
}
