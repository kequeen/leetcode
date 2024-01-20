package leetcode

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
// 两个正序数组的中位数，其实不就是归并排序取中位数么
// 代码写得比较丑，但也还是AC了
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num1Len := len(nums1)
	num2Len := len(nums2)
	var mid float64
	if num1Len == 0 && num2Len == 0 {
		return 0
	}
	if num1Len == 0 && num2Len != 0 {
		if num2Len%2 == 1 {
			mid = float64(nums2[num2Len/2])
		} else {
			mid = float64(nums2[num2Len/2]+nums2[num2Len/2-1]) / 2
		}
		return mid
	}

	if num1Len != 0 && num2Len == 0 {
		if num1Len%2 == 1 {
			mid = float64(nums1[num1Len/2])
		} else {
			mid = float64(nums1[num1Len/2]+nums1[num1Len/2-1]) / 2
		}
		return mid
	}

	midLen := 0
	if (num1Len+num2Len)%2 == 1 {
		midLen = (num1Len + num2Len) / 2
	} else {
		midLen = (num1Len+num2Len)/2 + 1
	}

	mergeNums := make([]int, num1Len+num2Len)

	i := 0
	j := 0
	place := 0
	for i < num1Len && j < num2Len && place <= midLen {
		if nums1[i] < nums2[j] {
			mergeNums[place] = nums1[i]
			i++
		} else {
			mergeNums[place] = nums2[j]
			j++
		}
		place++
	}
	for i < num1Len && place <= midLen {
		mergeNums[place] = nums1[i]
		i++
		place++
	}

	for j < num2Len && place <= midLen {
		mergeNums[place] = nums2[j]
		j++
		place++
	}

	if (num1Len+num2Len)%2 == 1 {
		//如果是奇数，那就是最后一个值
		mid = float64(mergeNums[midLen])
	} else {
		//如果是偶数，那就是最后两个值
		mid = float64(mergeNums[(num1Len+num2Len)/2]+mergeNums[(num1Len+num2Len)/2-1]) / 2
	}
	return mid
}
