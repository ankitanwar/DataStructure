package linkedlist

//NonRepeating : to get the first non Repeating character
func NonRepeating(s string) {
	present := [257]bool{}
	pointers := [257]*Dnode{}

	list := Dlinkedlist{}

	for i := 0; i < len(s); i++ {
		p := string(s[i])
		index := int(s[i]) - int('a')
		if present[index] == false && pointers[index] == nil {
			in := list.Insert(p)
			pointers[index] = in
		} else if present[index] == false && pointers[index] != nil {

			present[index] = true
			list.Ddlete(p)
			pointers[index] = nil
		} else if present[index] == true {
			continue
		}

		list.Print()
	}
}
