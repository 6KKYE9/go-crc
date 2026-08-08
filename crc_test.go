package main

import "testing"

func TestCRC32(t *testing.T) {
	// "123456789" 是 CRC 标准测试向量
	got, err := Sum("crc32", []byte("123456789"))
	if err != nil {
		t.Fatal(err)
	}
	if got != 0xCBF43926 {
		t.Errorf("crc32(\"123456789\")=%#x 想要 0xcbf43926", got)
	}
}

func TestCRC16(t *testing.T) {
	got, err := Sum("crc16", []byte("123456789"))
	if err != nil {
		t.Fatal(err)
	}
	if got != 0xBB3D {
		t.Errorf("crc16(\"123456789\")=%#x 想要 0xbb3d", got)
	}
}

func TestCRC8(t *testing.T) {
	got, err := Sum("crc8", []byte("123456789"))
	if err != nil {
		t.Fatal(err)
	}
	// SMBus CRC-8 标准测试向量
	if got != 0xF4 {
		t.Errorf("crc8(\"123456789\")=%#x 想要 0xf4", got)
	}
}

func TestBadName(t *testing.T) {
	if _, err := Sum("crc99", []byte("x")); err == nil {
		t.Error("不支持的算法该报错")
	}
}
