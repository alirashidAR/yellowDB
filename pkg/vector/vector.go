package vector

type Vector []float64


func (v Vector) Dimension () int {
	return len(v)
}


func (v Vector) Equal (v2 Vector) bool {
	if ( len(v) != len(v2) ) {
		return false
	}

	for i := 0; i < len(v); i++ {
		if (v[i] != v2[i]) {
			return false
		}
	}

	return true
}