package solution

import (
	"sort"
) //import

func Solution(A []int) int {
	// Sort the array elements in ascending order.
	sort.Ints(A)

	// Traverse ordered array and check that the value at any current index
	// is equivalent to that index plus 1.  If it is not, we have found the
	// missing element.
	for i, val := range A {
		if i+1 != val {
			return i + 1
		} //if
	} //for

	// If no missing element was found then N+1, or the length of the array
	// plus 1 is our missing element.
	return len(A) + 1
} //Solution
