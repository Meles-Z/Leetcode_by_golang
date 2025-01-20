package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	p1, p2, p := m-1, n-1, m+n-1

	for p1 >= 0 && p2 >= 0 {
		// Place the larger element at the end of nums1
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// If there are remaining elements in nums2, copy them
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
func main() {
	n1 := []int{2, 3, 4, 5, 0, 0, 0}
	n2 := []int{6}
	m := 4
	n := 1
	merge(n1, m, n2, n)
}
