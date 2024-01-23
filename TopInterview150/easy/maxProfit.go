package easy

func maxProfit(prices []int) int {
	max := 0
	for i := len(prices) - 1; i >= 1; i-- {
		if prices[i] <= max {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			if prices[i] <= max {
				break
			}
			if prices[i]-prices[j] > max {
				max = prices[i] - prices[j]
			}
		}
	}
	return max
}

// Better solution
func maxProfit2(prices []int) int {
	max := 0
	for i := len(prices) - 1; i >= 1; i-- {
		if prices[i] <= max {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			if prices[i] <= max {
				break
			}
			if prices[i]-prices[j] > max {
				max = prices[i] - prices[j]
			}
		}
	}
	return max
}
