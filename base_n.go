package base_n

import (
	"fmt"
)

type BaseN struct {
	chars []byte
}

func NewBaseN(chars []byte) *BaseN {
	return &BaseN{chars}
}

func NewBase3Number() *BaseN {
	return NewBaseN([]byte{'0', '1', '2'})
}

func NewBase30oO() *BaseN {
	return NewBaseN([]byte{'0', 'o', 'O'})
}

func NewBase62() *BaseN {
	return NewBaseN([]byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}

// Encode 数字转任意进制字符串
func (b *BaseN) Encode(n int64) string {
	if n == 0 {
		return string(b.chars[0])
	}

	var result []byte
	negative := n < 0

	// Special case for min int64 to prevent overflow
	// 在 Encode 函数里对负数进行编码时，会先将负数转为正数（通过 n = -n）。
	// 但是对于 int64 的最小值 -9223372036854775808，这个操作会导致溢出，因为 int64 的最大值是 9223372036854775807（比最小值的绝对值小1）。
	// 我在代码中添加了一个特殊处理，对 int64 最小值的情况单独处理：
	// 检测是否是 int64 最小值 (-9223372036854775808)
	// 如果是，则使用 uint64 来处理计算，避免溢出
	if n == -9223372036854775808 {
		// Convert to uint64 for calculation
		unsigned := uint64(9223372036854775808)
		charLen := uint64(len(b.chars))

		for unsigned > 0 {
			remainder := unsigned % charLen
			unsigned = unsigned / charLen
			result = append([]byte{b.chars[remainder]}, result...)
		}

		return "-" + string(result)
	}

	if negative {
		n = -n
	}

	charLen := int64(len(b.chars))

	for n > 0 {
		remainder := n % charLen
		n = n / charLen
		result = append([]byte{b.chars[remainder]}, result...)
	}

	if negative {
		return "-" + string(result)
	}
	return string(result)
}

// Decode 任意进制字符串转数字
func (b *BaseN) Decode(s string) (int64, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string")
	}

	values := make(map[byte]int64, len(b.chars))

	for i, item := range b.chars {
		values[item] = int64(i)
	}

	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	var result int64
	baseSize := int64(len(b.chars))

	for i := 0; i < len(s); i++ {
		char := s[i]
		val, exists := values[char]
		if !exists {
			return 0, fmt.Errorf("invalid character: %c", char)
		}
		result = result*baseSize + val
	}

	if negative {
		result = -result
	}
	return result, nil
}
