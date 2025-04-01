package array

// ArrUint32 type of array uint 32
type ArrUint32 uint32

// InArray checking, return bool and index
func (s ArrUint32) InArray(val uint32, array []uint32) (exists bool, index int) {
	exists = false
	index = -1

	for i, s := range array {
		if s == val {
			exists = true
			index = i
			return
		}
	}

	return
}

// Remove array by value
func (s ArrUint32) Remove(array []uint32, value uint32) []uint32 {
	isExist, index := s.InArray(value, array)
	if isExist {
		array = append(array[:index], array[(index+1):]...)
	}

	return array
}

// RemoveByIndex is remove array by index
func (s ArrUint32) RemoveByIndex(array []uint32, index int) []uint32 {
	return append(array[:index], array[(index+1):]...)
}
