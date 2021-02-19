package sort

// my ref: https://blog.csdn.net/weixin_43391310/article/details/106177042

func QuickSort(a []int) {
	n := len(a)
	qSort(a, 0, n-1)
}

func qSort(a []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition(a, start, end)
	qSort(a, start, pivot-1)
	qSort(a, pivot+1, end)
}

func partition(a []int, left, right int) int {
	// 将最右边的几点记为 pivot
	pivot := right

	// 不过届，即避免交换重复元素
	for left < right {
		// 在左边遇见比 pivot 大的
		if a[left] > a[pivot] {
			// 在右边遇见比 pivot 小的
			if a[right] < a[pivot] {
				a[left], a[right] = a[right], a[left]
			} else {
				right--
			}
		} else {
			left++
		}
	}

	/* 完成最后一组需要交换的数据
		如：{4, 3} 此时 left 停留在 4；right, pivot 停留在 3 上
	 */
	if a[left] != a[pivot] {
		a[left], a[pivot] = a[pivot], a[right]
	}

	// pivot 的值始终没有改变过
	// left 索引才是遍历到了真正的数据大小的分割点上
	return left
}
