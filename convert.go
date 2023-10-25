package convert

// Dereferences converts pointer to its non-pointer type and copies it on the stack.
func Dereference[T any](t *T) T {
    if t == nil {
        var ret T
        return ret
    }
    return *t
}

// Pointer allocates a copy of an object on the heap and returns a pointer to it.
func Pointer[T any](t T) *T {
    return &t
}

// SlicePointers converts a slice of values to a slice of pointers to those values. It does not copy the values.
func SlicePointers[T any](s []T) []*T {
    ret := make([]*T, len(s))
    for i := range s {
        ret[i] = &s[i]
    }
    return ret
}
