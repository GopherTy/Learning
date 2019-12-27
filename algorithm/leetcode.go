package algorithm

// TwoSum Problem 1
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
func TwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	indices := make([]int, 0)
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			indices = append(indices, m[target-v])
			indices = append(indices, i)
			return indices
		}
		m[v] = i
	}
	return indices
	// The other method
	// if len(nums) == 0 {
	//     return nil
	// }
	// indices := make([]int,0)
	// for i := 0 ;i < len(nums) - 1 ; i ++ {
	//     for j := i+1; j < len(nums) ; j ++ {
	//         if nums[i] + nums[j] == target {
	//             indices = append(indices,i)
	//             indices = append(indices,j)
	//             return indices
	//         }
	//     }
	// }
	// return indices
}

// AddTwoNumbers  Problem 2
// You are given two non-empty linked lists representing two non-negative integers.
//  The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 理解错误导致了一直做不出来，请注意它的链表连接顺序。
	head := new(ListNode)

	var t1, t2, tmp *ListNode
	t1 = l1
	t2 = l2
	tmp = head
	var carry int
	var x, y int
	for t1 != nil || t2 != nil {
		if t1 != nil {
			x = t1.Val
		} else {
			x = 0
		}
		if t2 != nil {
			y = t2.Val
		} else {
			y = 0
		}

		sum := carry + x + y
		carry = sum / 10

		tmp.Next = new(ListNode)
		tmp.Next.Val = sum % 10
		tmp = tmp.Next
		if t1 != nil {
			t1 = t1.Next
		}
		if t2 != nil {
			t2 = t2.Next
		}
	}
	if carry > 0 {
		tmp.Next = new(ListNode)
		tmp.Next.Val = carry
	}
	return head.Next
}

// ListNode list node struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// LengthOfLongestSubstrings Problem 3
// Given a string, find the length of the longest substring without repeating characters.
func LengthOfLongestSubstrings(s string) (l int) {
	// s is null return 0
	if s == "" {
		return
	}

	// reserve the substring
	var str string
	var tmpL int
	str = string(s[0])
	l = len(str)
	k := 1
	for i := 1; i < len(s); i++ {
		tmpL = len(str)

		// traverse substring to find the longest substring
		for j := 0; j < len(str); j++ {
			if s[i] == str[j] {
				str = ""
				i = k
				k++
				break
			}
		}

		// add substring
		str += string(s[i])

		// the longest substring length
		if tmpL < len(str) {
			tmpL = len(str)
		}
		if l < tmpL {
			l = tmpL
		}
	}

	return
}

//FindMedianSortedArrays Problem 4
// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.
func FindMedianSortedArrays(nums1 []int, nums2 []int) (m float64) {
	return
}
