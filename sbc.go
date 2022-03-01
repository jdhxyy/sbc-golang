// Copyright 2022-2022 The SUMEC Authors. All rights reserved.
// 结构体转换模块
// Authors: jdh99 <jdh821@163.com>

package sbc_golang

import (
	"bytes"
	"encoding/binary"
)

// StructToBytes 结构体转字节流
func StructToBytes(s interface{}) ([]uint8, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, s)
	if err != nil {
		return nil, err
	} else {
		return buf.Bytes(), err
	}
}

// BytesToStruct 字节流转结构体 结构体中的元素首字母要求大写
// s是结构体指针,保存转换后的结构体
func BytesToStruct(data []uint8, s interface{}) error {
	err := binary.Read(bytes.NewBuffer(data), binary.LittleEndian, s)
	return err
}
