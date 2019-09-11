package main

/*
step0: TODO time O(kn^2) space O(kn)
//第i天的第k次交易
dp[k][i] = max(
	dp[k][i-1] //在第i天不做交易，例如：1,10,9,8,7,6,5,4,3,2,1。卖出10元后不做交易
	p[i]-p[j] + dp[k-1][j-1] //j=[0,i-1]。在第i天一定做交易，并且是在第j这一天买入
							//注意这里j可以等于0：第0天买入，第i天卖出，dp[k-1][j-1]不合法使用dp技巧加一行解决，少做一次交易
)

step1: TODO time O(kn) space O(kn)
//使用示例，假设k=2
//第一天的交易,i=1
dp[2][1] = max(
	dp[2][0]
	p[1]-p[0]+dp[1][-1] //j=0，第一天的价格-第0天的价格
)
//第二天的交易,i=2
dp[2][2] = max(
	dp[2][1]
	p[2]-p[0]+dp[1][-1] //j=0
	p[2]-p[1]+dp[1][0]  //j=1
)
//第三天的交易,i=3
dp[2][3] = max(
	dp[2][2]
	p[3]-p[0]+dp[1][-1] //j=0
	p[3]-p[1]+dp[1][0]  //j=1
	p[3]-p[2]+dp[1][1]  //j=2
)
//第四天的交易,i=4
dp[2][4] = max(
	dp[2][3]
	p[4]-p[0]+dp[1][-1] //j=0
	p[4]-p[1]+dp[1][0]  //j=1
	p[4]-p[2]+dp[1][1]  //j=2
	p[4]-p[3]+dp[1][2]  //j=3
)

//第五天的交易,i=5
dp[2][5] = max(
	dp[2][4]
	p[5]-p[0]+dp[1][-1] //j=0
	p[5]-p[1]+dp[1][0]  //j=1
	p[5]-p[2]+dp[1][1]  //j=2
	p[5]-p[3]+dp[1][2]  //j=3
	p[5]-p[4]+dp[1][3]  //j=4
)
注意到我们求第x天的交易时，例如第四天、第五天
第四天中是p[4]-xxx0, p[4]-xxx1, p[4]-xxx2, p[4]-xxx3
第五天中是p[5]-xxx0, p[5]-xxx1, p[5]-xxx2, p[5]-xxx3, p[5]-xxx4
这里都有减xxx，我们可以提出来，为啥可以提出来：是因为减去的xxx和p[4], p[5]没有多大关系，我们可以缓存起来

step2: TODO time O(kn) space O(n)
一个测试用例是:交易10亿次，价格为1万次，因此当我们新建这么内存区，空间复杂度太高，我们来降低空间复杂度
为啥有这个想法：我们先来实验下填表的过程，还是以k=2,prices=[1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9]为例
	1	2	4	2	5	7	2	4	9	0	9
---------------------------------------------
00	00	00	00	00	00	00	00	00	00	00	00 //我们这里用2个0来表示边界
00	0	1	3	3	4	6	6	6	8	8	9  //第一次交易
00	0	1	3	3	6	8	8	8	13	13	17 //第二次交易
---------------------------------------------
我们以price=7(下标为6)，看下怎么计算出来的一次交易收益6，二次交易收益为8
dp[1][6] = max(
	dp[1][5]
	p[5] + x //注意这里的x为max(p[0]+dp[0][-1], p[1]+dp[0][0], p[2]+dp[0][1], p[3]+dp[0][2], p[4]+dp[0][3])
)
dp[2][6] = max(
	dp[2][5]
	p[5] + x //注意这里的x为max(p[0]+dp[1][-1], p[1]+dp[1][0], p[2]+dp[1][1], p[3]+dp[1][2], p[4]+dp[1][3])
)
我们来看这个x可以发现第n次交易只与第n-1次交易计算出来的dp值相关：第2次交易与第0次交易dp无关，第三次交易与第0次与第1次交易无关...
因此我们可以用2n的空间来存储dp值，而不用kn的空间了

step3: TODO time O(kn) space O(n)
注意到这里k非常大，例如10亿（共10亿次交易），但是我们的数据只有[]int{1, 2}
其实不管多少次交易，最大的收益只能是1，因此2次dp，如果获取最大的收益已经相同了，那么直接break
*/

func main() {
	println(maxProfit(2, []int{}))
	println(maxProfit(2, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}))
	println(maxProfit(0, []int{1, 3}))
}

func maxProfit(k int, prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, l+1)
	}
	tmp := 0
	for kk := 1; kk <= k; kk++ {
		tmp = dp[kk&1^1][0] - prices[0]
		for ll := 1; ll <= l; ll++ {
			tmp = max(tmp, dp[kk&1^1][ll-1]-prices[ll-1])
			dp[kk%2][ll] = max(dp[kk%2][ll-1], tmp+prices[ll-1])
		}
		if dp[0][l] == dp[1][l] {
			return dp[0][l]
		}
	}
	//println(fmt.Sprintf("%+v", dp))
	return dp[k%2][l]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
