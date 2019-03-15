人肉翻译:
> Search uses binary search to find and return the smallest index i
> in [0, n) at which f(i) is true, assuming that on the range [0, n),
> f(i) == true implies f(i+1) == true. 
使用二分搜索去寻找在[0,n)之间最小的索引值i, i满足以下条件: f(i)是true

`f(i) == true` 意味着`f(i+1) == true`

That is, Search requires that
// f is false for some (possibly empty) prefix of the input range [0, n)
// and then true for the (possibly empty) remainder; Search returns
// the first true index. If there is no such index, Search returns n.
// (Note that the "not found" return value is not -1 as in, for instance,
// strings.Index.)
搜索要求函数`f` 会在`[0,n)` 中的左侧都是返回false,而右侧的都返回true.然后`Search`函数会返回第一个为true的索引值`i`
如果没有这个索引值,会直接返回n

// Search calls f(i) only for i in the range [0, n).
`Search` 的回调函数`f(i)` 的i只会在范围[0,n)之中

// A common use of Search is to find the index i for a value x in
// a sorted, indexable data structure such as an array or slice.
// In this case, the argument f, typically a closure, captures the value
// to be searched for, and how the data structure is indexed and
// ordered.
//
// For instance, given a slice data sorted in ascending order,
// the call Search(len(data), func(i int) bool { return data[i] >= 23 })
// returns the smallest index i such that data[i] >= 23. If the caller
// wants to find whether 23 is in the slice, it must test data[i] == 23
// separately.

举个例子,一个升序的`slice`,调用函数
```
Search(   len(data), 
          func(i int) bool 
          { 
            return data[i] >= 23 
          }
      )
```
会返回第一个slice 的元素大于等于23 的值

```

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h    // h 基本可以等于 (i+j)/2 即索引的中值
		// i ≤ h < j
		if !f(h) {   // 如果是false 即 从[h,j)
			i = h + 1 // preserves f(i-1) == false
		} else {    // 否则[i,h) 
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.   // for 的跳出条件i==j
	return i
}
```
