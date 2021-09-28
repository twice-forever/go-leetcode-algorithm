package problem

var (
	Nums1 = []int{2}
	Nums2 = []int{}
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	nums3 := make([]int, 0, nums1Len+nums2Len)

	nums1Head := 0
	nums2Head := 0
	for {
		if nums1Len == 0 {
			nums3 = append(nums3, nums2...)
			break
		}
		if nums2Len == 0 {
			nums3 = append(nums3, nums1...)
			break
		}

		if nums1[nums1Head] < nums2[nums2Head] {
			nums3 = append(nums3, nums1[nums1Head])
			nums1Head++
			if nums1Head == nums1Len {
				nums3 = append(nums3, nums2[nums2Head:]...)
				break
			}
		} else {
			nums3 = append(nums3, nums2[nums2Head])
			nums2Head++
			if nums2Head == nums2Len {
				nums3 = append(nums3, nums1[nums1Head:]...)
				break
			}
		}
	}

	var medianSort float64
	nums3Len := len(nums3)
	if nums3Len%2 == 0 {
		num1 := nums3[nums3Len/2-1]
		num2 := nums3[nums3Len/2]
		medianSort = (float64(num1) + float64(num2)) / 2
	} else {
		medianSort = float64(nums3[nums3Len/2])
	}
	return medianSort
}
