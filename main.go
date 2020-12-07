package main

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

func main() {
	ksuid, err := ksuid.NewRandom()
	//KSUIDの生成
	if err != nil {
		fmt.Errorf("failed to generate new KSUID: %s", err.Error())
	}

	//取得
	fmt.Println(ksuid.Get())
	//KSUID内の時刻を取得
	fmt.Println(ksuid.Time().UTC())
}
