package core

import (
	"fmt"
	"os"
	"syscall"
)

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
	return nAmount, nil
}

func ValueFromAmount(amount int64) (value float64) {
	return float64(amount) / float64(COIN)
}

func MoneyRange(nValue int64) (bool) {
	return nValue >= 0 && nValue <= MAX_MONEY
}

func Add(val1, val2 float64) (float64, error) {
	v1, err := AmountFromValue(val1)
	if err != nil {
		return 0, err
	}

	v2, err := AmountFromValue(val2)
	if err != nil {
		return 0, err
	}
	v3 := v1 + v2
	v4 := ValueFromAmount(v3)
	return v4, nil
}

func LockFile(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

//进程锁
func ProcLock(filename string) (*os.File) {
	iManPid := fmt.Sprint(os.Getpid())
	pidFile, _ := os.Create(filename)
	pidFile.WriteString(iManPid)
	if err := LockFile(pidFile); err != nil {
		return nil
	}
	return pidFile
}
