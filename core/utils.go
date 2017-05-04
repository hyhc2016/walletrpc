package core

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"errors"
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

//进程锁
func ProcLock(filename string) {
	iManPid := fmt.Sprint(os.Getpid())
	tmpDir := os.TempDir()
	if err := ProcExsit(tmpDir); err == nil {
		pidFile, _ := os.Create(filename)
		defer pidFile.Close()
		pidFile.WriteString(iManPid)
	} else {
		os.Exit(1)
	}
}

// 判断进程是否启动
func ProcExsit(filename string) (err error) {
	iManPidFile, err := os.Open(filename)
	defer iManPidFile.Close()
	if err == nil {
		filePid, err := ioutil.ReadAll(iManPidFile)
		if err == nil {
			pidStr := fmt.Sprintf("%s", filePid)
			pid, _ := strconv.Atoi(pidStr)
			_, err := os.FindProcess(pid)

			if err == nil {
				return errors.New("[ERROR] 进程已经启动.")
			}
		}
	}
	return nil
}
