package containers

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

// type for containers/heap
type OrderNums[] int 

func (on *OrderNums) Len() int {
	return len(*on)
}

func (on OrderNums) Less(i, j int) bool {
	return  on[i] < on[j] 
}

func (on OrderNums) Swap(i, j int) {
	on[i], on[j] = on[j], on[i]
}


func (on *OrderNums) Push(x interface{})  {
	*on = append(*on, x.(int))
}

func (on *OrderNums) Pop() interface{} {
	old := *on 
	n := len(old) - 1
	x := old[n]
	*on = old[: n]
	return x
}

func main() {

	
	/* container/list */
	listInstance := list.New()
	e0 := listInstance.PushBack(49)
	e1 := listInstance.PushFront(80)
	e2 := listInstance.PushFront(58)
	e3 := listInstance.PushBack(10)

	listInstance.InsertBefore(7, e0)
	listInstance.InsertBefore(253, e1)
	listInstance.InsertBefore(2371, e2)
	listInstance.InsertAfter(423, e3)

	fmt.Println("*********** List ***********")
	fmt.Println("------- Step:1 --------")
	for e := listInstance.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d", e.Value.(int))
	}
	fmt.Printf("\n")
	listInstance.MoveToFront(e3)
	listInstance.MoveToBack(e0)
	listInstance.Remove(e1)

	fmt.Println("-- Step 2:")
	for e := l.Front(); e != nil; e = e.Next() {
	fmt.Printf("%d ", e.Value.(int))
	}
	fmt.Printf("\n")


	// *** container/ring ***
	// create the ring and populate it
	blake := []string{"the", "invisible", "worm"}

	r := ring.New(len(blake))
	for i:=0; i < r.Len(); i++ {
		r.Value = blake[i]
		r = r.Next()
	}

	// move(2 % r.Len())=1 elements forward in the ring
	r = r.Move(2)

	fmt.Printf("\n*** RING ***\n")
	// print all the ring values with ring.Do()
	r.Do(func(x interface {}) {
		fmt.Printf("%s\n", x.(string))
	})

	//*** container/heap
	h := &OrderNums{34, 234, 121, 467, 896, 2112, 467, 10, 42, 3245, 56875}
	heap.Init(h)

	fmt.Printf("\n*** HEAP ***\n")
	fmt.Printf("min: %d\n", (*h)[0])
	fmt.Printf("heap:\n")
	for h.Len() > 0 {
	fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Printf("\n")


}