type DynamicArray struct {
	arr      []int 
	size     int   
	capacity int   
	}
func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		arr:      make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.size == da.capacity {
		da.Resize()
	}
	da.arr[da.size] = n
	da.size++
}

func (da *DynamicArray) Popback() int {
	da.size--
	return da.arr[da.size]
}

func (da *DynamicArray) Resize() {
	da.capacity = da.capacity * 2
	newArr := make([]int, da.capacity)
	for i := 0; i < da.size; i++ {
		newArr[i] = da.arr[i]
	}
	da.arr = newArr
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}