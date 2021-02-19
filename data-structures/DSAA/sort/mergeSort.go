package sort

/* 归并排序
也是一个经典的分治过程：
- 分：递归的分，其终止条件为待分割元素数量为 1，如果只有一个元素那么显然有序
- 治：从两次的返回的有序数组结果中合并出新的有序数组
*/

// MergeSort 启动例程
func MergeSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	mSort(a, 0, n-1)
}

// mSort 进行分操作
func mSort(a []int, le, ri int) {
	// 等号必须存在，否则在 le==ri 时 p 取值将超过索引
	if le >= ri {
		return
	}
	// 取中间位置 p 进行**分**割
	p := (le + ri) / 2
	mSort(a, le, p)
	mSort(a, p+1, ri)
	//合并函数，**治**
	merge(a, le, p, ri)
}

// merge 是真个归并的核心例程
func merge(a []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0 // 用于临时数组的索引

	for ; i <= mid && j <= end; k++ {
		if a[i] <= a[j] {
			tmpArr[k] = a[i]
			i++
		} else {
			tmpArr[k] = a[j]
			j++
		}
	}

	// 将某一剩余数组部分拷贝至临时数组
	for ; i <= mid; i++ {
		tmpArr[k] = a[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = a[j]
		k++
	}

	copy(a[start:end+1], tmpArr)
}
