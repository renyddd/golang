package sort

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		// 比较元素从 0 开始，由于是与 j+1 进行比较
		// 故比较至 N-已有序元素-1
		// 已有序元素：被冒泡至末尾的元素
		for j := 0; j < len(a)-i-1; j++ {
			if a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
