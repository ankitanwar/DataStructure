package backtracking

import (
	"fmt"
)

//We have given 2 non identical items we need to cal the numbers of ways in which we can arrange these 2 items into 3 boxes

//Every box will have 3 choice either it will not place the item or it will place the first item or it will place the second item

//Number of boxes will always will be larger than the items

//Permutation : To calc the permutation
func Permutation(given string, boxes, k int) {
	items := make([]bool, len(given))
	helperPermu(1, boxes, 0, 2, &items, "", given)

}

func helperPermu(cb, tb, ssf, ts int, items *[]bool, asf, given string) {
	if cb > tb {
		if ssf == ts {
			fmt.Println(asf)
		}
		return
	}
	for i := 0; i < ts; i++ {
		if (*items)[i] == false {
			(*items)[i] = true
			helperPermu(cb+1, tb, ssf+1, ts, items, asf+string(given[i]), given)
			(*items)[i] = false
		}
	}
	helperPermu(cb+1, tb, ssf, ts, items, asf, given)
}
