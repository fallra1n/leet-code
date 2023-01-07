package main

//
//type Pair struct {
//	key int
//	val string
//}
//
//type MaxHeap struct {
//	arr []Pair
//	n   int
//}
//
//func (mh *MaxHeap) SiftDown(id int) {
//	var max int
//	l := 2*id + 1
//	r := 2*id + 2
//	mh.n = len(mh.arr)
//
//	if l < mh.n {
//		if l == mh.n-1 {
//			max = l
//		} else {
//			if mh.arr[l].key > mh.arr[r].key {
//				max = l
//			} else if mh.arr[l].key < mh.arr[r].key {
//				max = r
//			} else {
//				if mh.arr[l].val < mh.arr[r].val {
//					max = l
//				}
//				if mh.arr[l].val > mh.arr[r].val {
//					max = r
//				}
//			}
//		}
//
//		a := mh.arr[id].key < mh.arr[max].key
//		b := (mh.arr[id].key == mh.arr[max].key) && (mh.arr[id].val > mh.arr[max].val)
//
//		if a || b {
//			temp := mh.arr[id]
//			mh.arr[id] = mh.arr[max]
//			mh.arr[max] = temp
//			mh.SiftDown(max)
//		}
//	}
//}
//
//func (mh *MaxHeap) BuildHeap(data []Pair) {
//	mh.arr = data
//	mh.n = len(data)
//
//	for i := mh.n / 2; i >= 0; i-- {
//		mh.SiftDown(i)
//	}
//}
//
//func topKFrequent(words []string, k int) []string {
//	d := make(map[string]int)
//	for i := 0; i < len(words); i++ {
//		d[words[i]]++
//	}
//
//	var pq []Pair
//	for key, val := range d {
//		pq = append(pq, Pair{val, key})
//	}
//
//	var mHeap MaxHeap
//	mHeap.BuildHeap(pq)
//
//	var ans []string
//	for i := 0; i < k; i++ {
//		ans = append(ans, mHeap.arr[0].val)
//		mHeap.arr[0] = mHeap.arr[len(mHeap.arr)-1]
//		mHeap.arr = mHeap.arr[:len(mHeap.arr)-1]
//		mHeap.SiftDown(0)
//	}
//
//	return ans
//}
//
//func main() {
//	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
//}
