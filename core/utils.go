package core

const (
	CENT      = 1000000
	COIN      = 100 * CENT
	MAX_MONEY = 120000000 * COIN
)

func roundint64(d float64) (int64) {
	if d > 0 {
		return int64(d + 0.5)
	} else {
		return int64(d - 0.5)
	}
}

func roundint(d float64) (int) {
	if d > 0 {
		return int(d + 0.5)
	} else {
		return int(d - 0.5)
	}
}

func abs64(d int64) (int64) {
	if d > 0 {
		return d
	} else {
		return -d
	}
}
func AmountFromValue(value float64) (int64, error) {
	nAmount := roundint64(value * COIN);
	if MoneyRange(nAmount) {
		return nAmount, nil
	} else {
		return 0, nil
	}
}

func ValueFromAmount(amount int64) (value float64) {
	return float64(amount) / float64(COIN)
}

func MoneyRange(nValue int64) (bool) {
	return nValue >= 0 && nValue <= MAX_MONEY
}

/* 计算数值
	v1, err := core.AmountFromValue(10000000.00009999)
	if err != nil {
		return
	}
	v2, err := core.AmountFromValue(10000000.99999999)
	if err != nil {
		return
	}
	v3 := v1 + v2
	v4 := core.ValueFromAmount(v3)
	v5 := strconv.FormatFloat(v4,'f',-1,64)
	fmt.Println(v1, v2, v3, v4,v5)
	return
 */
