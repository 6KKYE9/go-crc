# go-crc

零依赖的 CRC 校验和计算工具，支持 crc8 / crc16 / crc32 / crc32c。

## 用法

```bash
go-crc crc32 "123456789"   # 输出 0xcbf43926
```

库函数 `Sum(name string, data []byte) (uint64, error)` 可直接调用。
