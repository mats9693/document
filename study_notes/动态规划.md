# 学习笔记：动态规划  
> 动态规划（dynamic programming）：运筹学的一个分支，求解决策过程最优化的数学方法。  

动态规划通常用来求解含有最优子结构的问题，一般求解过程步骤如下：  
1. 总结出状态转移方程
1. 将状态转移方程翻译成代码
1. 代码优化

状态转移方程基本可以直译成代码，所以动态规划算法的主要难点在于总结状态方程和最后的代码优化。  

顺便说一下，其实下面的两个例子并不是严格的动态规划算法，因为它们没有“最优子结构”，故应称为“使用动态规划思想”；但现在大家都这么用，所以也没有必要抓着定义不放了。

---  
## example 1：Fibonacci（入门）  
先来一个简单的例子，Fibonacci数列的动态规划解法。  
### 解：  
#### 定义：
dp(m)为Fibonacci数列的第m个值  
#### 易知：  
<a href="https://www.codecogs.com/eqnedit.php?latex=dp(m)=\begin{cases}&space;\text{&space;dp(m-1)&space;&plus;&space;dp(m-2)&space;}&space;&&space;m&space;\geq&space;2&space;\\&space;\text{&space;0&space;}&space;&&space;m&space;=&space;0&space;\\&space;\text{&space;1&space;}&space;&&space;m&space;=&space;1&space;\end{cases}" target="_blank"><img src="https://latex.codecogs.com/gif.latex?dp(m)=\begin{cases}&space;\text{&space;dp(m-1)&space;&plus;&space;dp(m-2)&space;}&space;&&space;m&space;\geq&space;2&space;\\&space;\text{&space;0&space;}&space;&&space;m&space;=&space;0&space;\\&space;\text{&space;1&space;}&space;&&space;m&space;=&space;1&space;\end{cases}" title="dp(m)=\begin{cases} \text{ dp(m-1) + dp(m-2) } & m \geq 2 \\ \text{ 0 } & m = 0 \\ \text{ 1 } & m = 1 \end{cases}" /></a>  
#### 代码：
```go  
func fibonacci(n int) (result int) {
    switch n {
    case 0, 1: result = n
    default:
        dp := make([]int, n+1)
        dp[1] = 1
        for i := 2; i < len(dp); i++ {
            dp[i] = dp[i-1] + dp[i-2]
        }
        result = dp[n]
    }

    return
}
```
#### 优化：  
动态规划算法的一个特点，就是用空间换时间；它的优化主要即是存储空间上的优化。  
例如在本例中，dp[3]的值只在计算dp[4]、dp[5]时使用，对其进行优化：
```go  
func fibonacciOptimize(n int) (result int) {
    switch n {
    case 0, 1: result = n
    default:
        dp := make([]int, 2)
        dp[1] = 1
        for i := 2; i < len(dp); i++ {
            dp[i%2] = dp[0] + dp[1]
        }
        result = dp[n]
    }

    return
}
```
这样一来，函数的空间复杂度就从O(n)变成了O(1)。  

---  
## example 2：实现一个简单正则引擎  
[leetcode q10](https://leetcode.com/problems/regular-expression-matching/)  
本例仅实现小写字母、“.”、“\*”的正则匹配。
### 解：  
#### 定义：  
s：字符串 lens：s的长度  
p：正则表达式 lenp：p的长度  
dp[i][j]为s的前i个字符（s[:i]）与p的前j个字符（p[:j]）是否匹配
#### 分析：  
小写字母和“.”的匹配很简单，关键在于“\*”；  
这里我们选择将正则表达式中的“\*”代换成它的前一个字符，与字符串对应位置的字符进行匹配，结果只有两种：  
 - 匹配成功：```p[j-2] == s[i-1] || p[j-2] == '.'```，此时
```go 
dp[i][j] = dp[i-1][j] // 既然s[i-1] == p[j-2]，p[j-1]又是“*”；那么是否添加s[i-1]，匹配结果是一致的。 
```
 - 匹配失败：此时```dp[i][j] = dp[i][j-2]```。  

得到状态转移方程：  
<a href="https://www.codecogs.com/eqnedit.php?latex=dp[i][j]=\begin{cases}&space;\text{&space;dp[i][j]&space;=&space;dp[i-1][j]&space;}&space;&&space;p[j-1]&space;==&space;'*'&space;\quad&space;and&space;\quad&space;(p[j-2]&space;==&space;s[i-1]&space;\quad&space;or&space;\quad&space;p[j-2]&space;==&space;'.')&space;\\&space;\text{&space;dp[i][j]&space;=&space;dp[i][j-2]&space;}&space;&&space;p[j-1]&space;==&space;'*'&space;\quad&space;and&space;\quad&space;!(p[j-2]&space;==&space;s[i-1]&space;\quad&space;or&space;\quad&space;p[j-2]&space;==&space;'.')&space;\\&space;\text{&space;dp[i][j]&space;=&space;dp[i-1][j-1]&space;}&space;&&space;p[j-1]&space;==&space;s[i-1]&space;\quad&space;or&space;\quad&space;p[j-1]&space;==&space;'.'&space;\end{cases}" target="_blank"><img src="https://latex.codecogs.com/gif.latex?dp[i][j]=\begin{cases}&space;\text{&space;dp[i][j]&space;=&space;dp[i-1][j]&space;}&space;&&space;p[j-1]&space;==&space;'*'&space;\quad&space;and&space;\quad&space;(p[j-2]&space;==&space;s[i-1]&space;\quad&space;or&space;\quad&space;p[j-2]&space;==&space;'.')&space;\\&space;\text{&space;dp[i][j]&space;=&space;dp[i][j-2]&space;}&space;&&space;p[j-1]&space;==&space;'*'&space;\quad&space;and&space;\quad&space;!(p[j-2]&space;==&space;s[i-1]&space;\quad&space;or&space;\quad&space;p[j-2]&space;==&space;'.')&space;\\&space;\text{&space;dp[i][j]&space;=&space;dp[i-1][j-1]&space;}&space;&&space;p[j-1]&space;==&space;s[i-1]&space;\quad&space;or&space;\quad&space;p[j-1]&space;==&space;'.'&space;\end{cases}" title="dp[i][j]=\begin{cases} \text{ dp[i][j] = dp[i-1][j] } & p[j-1] == '*' \quad and \quad (p[j-2] == s[i-1] \quad or \quad p[j-2] == '.') \\ \text{ dp[i][j] = dp[i][j-2] } & p[j-1] == '*' \quad and \quad !(p[j-2] == s[i-1] \quad or \quad p[j-2] == '.') \\ \text{ dp[i][j] = dp[i-1][j-1] } & p[j-1] == s[i-1] \quad or \quad p[j-1] == '.' \end{cases}" /></a>
#### 代码：  
```go  
func isMatchDP(s, p string) bool {
    lens, lenp := len(s), len(p)
    dp := make([][]bool, lens+1)
    for i := range dp {
        dp[i] = make([]bool, lenp + 1)
    }

    dp[0][0] = true // s == p == "" ; dp[0][1] = false, default, omit

    for j := 2; j <= lenp; j++ {
        dp[0][j] = dp[0][j-2] && p[j-1] == '*'// s: "a", p: "b*c*d*a"
    }

    for i := 1; i <= lens; i++ { // i: ith
        for j := 1; j <= lenp; j++ { // j: jth
            if p[j-1] == '*' {
                dp[i][j] = dp[i][j-2] || dp[i-1][j] && (p[j-2] == s[i-1] || p[j-2] == '.') // skill only on bool type
            } else if p[j-1] == s[i-1] || p[j-1] == '.' {
                dp[i][j] = dp[i-1][j-1]
            }
        }
    }

    return dp[lens][lenp]
}
```
#### 优化：  
请自行完成该函数的优化。

--- 
## Summary  
动态规划的核心是状态转移方程，总结出状态转移方程，剩下的就是编程的基本功了。  

---
###### Mario
I've been pretending to work hard, but you're really growing up.
