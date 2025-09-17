# base-n

一个轻量级的Go库，用于在任意进制的数字系统中进行编码和解码。

## 特性

- 将数字转换为自定义进制表示法，并支持反向转换
- 支持负数处理
- 内置常用进制预设：
  - Base62（字母数字：0-9，a-z，A-Z）
  - Base3（0，1，2）
  - 特殊的Base3（0，o，O）
- 支持创建自定义字母表用于特殊的编码方案

## 安装

```bash
go get github.com/muleiwu/base-n
```

## 用法

### 基本用法

```go
package main

import (
    "fmt"
    BaseN "github.com/muleiwu/base-n"
)

func main() {
    // 创建Base62编码器/解码器
    base62 := BaseN.NewBase62()
    
    // 将整数编码为Base62字符串
    encoded := base62.Encode(12345)
    fmt.Println("编码结果:", encoded)
    
    // 将Base62字符串解码回整数
    decoded, err := base62.Decode(encoded)
    if err != nil {
        panic(err)
    }
    fmt.Println("解码结果:", decoded)
}
```

### 自定义字母表

```go
package main

import (
    "fmt"
    BaseN "github.com/muleiwu/base-n"
)

func main() {
    // 使用自定义字符集创建自定义进制
    customBase := BaseN.NewBaseN([]byte("01"))  // 二进制
    
    // 使用自定义进制进行编码和解码
    encoded := customBase.Encode(10)
    fmt.Println("二进制编码:", encoded)  // 输出: 1010
    
    decoded, _ := customBase.Decode(encoded)
    fmt.Println("解码结果:", decoded)  // 输出: 10
}
```

### 预定义进制

```go
// Base3（0，1，2）
base3 := BaseN.NewBase3Number()

// 特殊Base3（0，o，O）
base3Special := BaseN.NewBase30oO()

// Base62（0-9，a-z，A-Z）
base62 := BaseN.NewBase62()
```

## API参考

### 构造函数

- `NewBaseN(chars []byte) *BaseN`：使用自定义字符创建新的Base-N编码器/解码器
- `NewBase3Number() *BaseN`：创建Base3编码器/解码器（0，1，2）
- `NewBase30oO() *BaseN`：创建特殊的Base3编码器/解码器（0，o，O）
- `NewBase62() *BaseN`：创建Base62编码器/解码器（0-9，a-z，A-Z）

### 方法

- `Encode(n int64) string`：将整数转换为base-n字符串
- `Decode(s string) (int64, error)`：将base-n字符串转换回整数

## 许可证

MIT
