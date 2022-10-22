func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums3 := []int{}
    var medRes float64

    nums3 = append(nums1, nums2...)
    sort.Ints(nums3)
    
    medRes = median(nums3)
    return medRes
}

func median(nums3 []int) float64 {
    lenth := len(nums3)
    var med float64
    
    n1 := nums3[len(nums3)/2]
	    
    if lenth % 2 == 1 {
        med = float64(n1)
    } else {
        n2 := nums3[(len(nums3)/2)-1]
        med = (float64(n1) + float64(n2)) / 2
    }
    return med
}