package sbc_golang

import (
	"fmt"
	"testing"
)

type RadioRxHeader struct {
	// 源设备类型
	DeviceType uint8
	// 接收信道
	Channel uint8
	// 接收RSSI
	Rssi uint8
	// 接收时隙
	Slot uint16
	// 正文长度
	PayloadLen uint16
}

func TestCase1(t *testing.T) {
	a := RadioRxHeader{1, 2, 3, 0x1234, 0x5678}
	data, err := StructToBytes(a)
	fmt.Println(err, data)

	var b RadioRxHeader
	err = BytesToStruct(data, &b)
	fmt.Println(err, b)
}
