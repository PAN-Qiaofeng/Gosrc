/*
 * @Description:
 * @version:
 * @Author: pqf
 * @Date: 2021-06-06 14:12:47
 * @LastEditors: pqf
 * @LastEditTime: 2021-06-12 14:10:01
 */
package mysort

// func sortArray(nums []int) []int {
// 	heapSort(nums)
// 	return nums
// }

// func heapSort(nums []int) []int {
// 	end := len(nums) - 1
// 	for i := end / 2; i >= 0; i-- {
// 		sink(nums, i, end)
// 	}
// 	for i := end; i >= 0; i-- {
// 		nums[0], nums[i] = nums[i], nums[0]
// 		end--
// 		sink(nums, 0, end)
// 	}
// 	return nums
// }

// func sink(heap []int, root, end int) {
// 	for {
// 		child := root*2 + 1
// 		if child > end {
// 			return
// 		}
// 		if child < end && heap[child] <= heap[child+1] {
// 			child++
// 		}
// 		if heap[root] > heap[child] {
// 			return
// 		}
// 		heap[root], heap[child] = heap[child], heap[root]
// 		root = child
// 	}
// }

// func quickSort(nums []int, left, right int) {
// 	if left > right {
// 		return
// 	}
// 	i, j, base := left, right, nums[left]
// 	for i < j {
// 		for nums[j] >= base && i < j {
// 			j--
// 		}
// 		for nums[i] <= base && i < j {
// 			i++
// 		}
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// 	nums[i], nums[left] = nums[left], nums[i]
// 	quickSort(nums, left, i-1)
// 	quickSort(nums, i+1, right)
// }

// 1 void QuickSort(int *a, int left,int right)
// 2 {
// 3     if (a == NULL || left < 0 || right <= 0 || left>right)
// 4         return;
// 5     stack<int>temp;
// 6     int i, j;
// 7     //（注意保存顺序）先将初始状态的左右指针压栈
// 8     temp.push(right);//先存右指针
// 9     temp.push(left);//再存左指针
// 10     while (!temp.empty())
// 11     {
// 12         i = temp.top();//先弹出左指针
// 13         temp.pop();
// 14         j = temp.top();//再弹出右指针
// 15         temp.pop();
// 16         if (i < j)
// 17         {
// 18             int k = Pritation(a, i, j);
// 19             if (k > i)
// 20             {
// 21                 temp.push(k - 1);//保存中间变量
// 22                 temp.push(i);  //保存中间变量
// 23             }
// 24             if (j > k)
// 25             {
// 26                 temp.push(j);
// 27                 temp.push(k + 1);
// 28             }
// 29         }
// 30
// 31     }
// 32
// 33 }

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r, i := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	copy(result[i:], left[l:])
	copy(result[i+len(left)-l:], right[r:])
	return result
}
