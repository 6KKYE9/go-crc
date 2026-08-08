# go-crc

又在终端里为个编码解码去开浏览器搜网页？别了，一行命令的事。

## 用法

```bash
go-crc crc32 "123456789"   # 输出 0xcbf43926
```

库函数 `Sum(name string, data []byte) (uint64, error)` 可直接调用。
