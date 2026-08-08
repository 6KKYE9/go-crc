package main

import (
	"fmt"
	"hash/crc32"
)

// 各算法参数。refl 表示是否按位反射（常见 CRC 都是反射的）。
// crc8 用非反射的 SMBus 参数，crc16 用反射后的 ARC 参数，
// crc32 / crc32c 直接走标准库（也是标准库，不引入第三方依赖）。
var params = map[string]struct {
	width  int
	poly   uint64
	init   uint64
	xorOut uint64
	refl   bool
}{
	"crc8":   {8, 0x07, 0x00, 0x00, false},
	"crc16":  {16, 0xA001, 0x0000, 0x0000, true},
	"crc32":  {32, 0xEDB88320, 0xFFFFFFFF, 0xFFFFFFFF, true},
	"crc32c": {32, 0x82F63B78, 0xFFFFFFFF, 0xFFFFFFFF, true},
}

// Sum 按指定算法名计算校验值。
// 支持：crc8 / crc16 / crc32 / crc32c。
func Sum(name string, data []byte) (uint64, error) {
	p, ok := params[name]
	if !ok {
		return 0, fmt.Errorf("不支持的算法: %s", name)
	}
	// 这两个标准库已经算好，直接返回
	if name == "crc32" {
		return uint64(crc32.ChecksumIEEE(data)), nil
	}
	if name == "crc32c" {
		return uint64(crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli))), nil
	}

	crc := p.init
	mask := (uint64(1) << p.width) - 1
	top := uint64(1) << (p.width - 1)
	for _, b := range data {
		if p.refl {
			crc ^= uint64(b)
			for i := 0; i < 8; i++ {
				if crc&1 != 0 {
					crc = (crc >> 1) ^ p.poly
				} else {
					crc >>= 1
				}
			}
		} else {
			crc ^= uint64(b) << (p.width - 8)
			for i := 0; i < 8; i++ {
				if crc&top != 0 {
					crc = (crc << 1) ^ p.poly
				} else {
					crc <<= 1
				}
			}
		}
	}
	return (crc & mask) ^ p.xorOut, nil
}
