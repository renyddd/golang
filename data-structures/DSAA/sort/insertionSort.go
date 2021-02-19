package sort

// InsertionSort 进行 P = 1 到 P = N-1 躺排序
// 并且每趟保证从 0 到元素 P 为已排序状态
// 既然 0 到 P-1 是有序的，所以只需要将 P 元素插入到合适的位置即可
func InsertionSort(a []int) {
	// 长度为 0，1 则已为有序状态
	for p := 1; p < len(a); p++ {
		for i := p; i > 0; i-- {
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
			}
		}
	}
}
