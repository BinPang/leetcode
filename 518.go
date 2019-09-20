package main

import "sort"

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	sort.Ints(coins)
	dp := make(map[int]map[int]int, 0)
	return _change(amount, coins, 0, len(coins)-1, &dp)
}

func _change(amount int, coins []int, start, end int, dp *map[int]map[int]int) int {
	if _, ok := (*dp)[start]; ok {
		if v, ok1 := (*dp)[start][end]; ok1 {
			return v
		}
	}
	r := 0
	for i := end; i >= start; i-- {
		if coins[i] < amount {
			r += _change(amount-coins[i], coins, start, i, dp)
		} else if coins[i] == amount {
			r += 1
		}
	}
	if (*dp)[start] == nil {
		(*dp)[start] = map[int]int{end: r}
	}
	return r

}

/*
You are given coins of different denominations and a total amount of money.
Write a function to compute the number of combinations that make up that amount.
You may assume that you have infinite number of each kind of coin.



Example 1:

Input: amount = 5, coins = [1, 2, 5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
Example 3:

Input: amount = 10, coins = [10]
Output: 1


Note:

You can assume that

0 <= amount <= 5000
1 <= coin <= 5000
the number of coins is less than 500
the answer is guaranteed to fit into signed 32-bit integer
*/
