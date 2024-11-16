package main

import (
	"fmt"
)

// RemoveElementAt 是一个泛型函数，用于从切片中删除指定索引处的元素。
// 如果删除后切片的长度远小于其容量，则会缩小底层数组的大小以节省内存。
// 缩容的阈值可以根据需要进行调整。
func RemoveElementAt[T any](s []T, index int) ([]T, error) {
	// 检查索引是否超出范围
	if index < 0 || index >= len(s) {
		return nil, fmt.Errorf("索引超出范围")
	}

	// 执行删除操作：通过将索引前后的元素合并到一个新的切片中
	s = append(s[:index], s[index+1:]...)

	// 可选：如果切片长度显著小于其容量，则缩小切片的容量
	// 这里使用一个简单的启发式规则：如果长度小于容量的25%，则缩容
	if len(s)*4 < cap(s) {
		// 创建一个新切片，长度相同，但容量减半
		newSlice := make([]T, len(s), cap(s)/2)
		// 将现有切片的内容复制到新切片中
		copy(newSlice, s)
		// 更新切片引用
		s = newSlice
	}

	// 返回修改后的切片
	return s, nil
}

func main() {
	// 示例用法
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", s)

	// 调用 RemoveElementAt 函数删除索引为 2 的元素
	var err error
	s, err = RemoveElementAt(s, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("删除后切片:", s)
	}
}
