package tests

import (
	"testing"
	"walletrpc"
	"fmt"
)

func TestRpc(t *testing.T) {
	if rpc, err := walletrpc.NewClient("123", "123sds", "127.0.0.1", 33325); err != nil {
		fmt.Println(err)
		return
	} else {
		//if t, err := rpc.ListTransactions("", 100, 0); err != nil {
		//	fmt.Println(err)
		//	return
		//} else {
		//	for _, v := range t {
		//		fmt.Println(v)
		//	}
		//}

		fmt.Println(rpc.GetBalance())

		fmt.Println(rpc.SendToaddress("KtxcaUixq1akdnu8FwbhodJAc6KbfiWxPW",11))
	}
}
