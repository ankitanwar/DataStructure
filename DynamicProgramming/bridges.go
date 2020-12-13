package dynamicprogramming

import (
	"fmt"
	"sort"
)

type bridge struct {
	north int
	south int
}

//NonOverllapingBridges : To find out the maxium non overlappingsub bridges
func NonOverllapingBridges(north, south []int) {
	bridges := []bridge{}
	for i := 0; i < len(north); i++ {
		temp := bridge{
			north: north[i],
			south: south[i],
		}
		bridges = append(bridges, temp)
	}
	sort.Slice(bridges, func(i, j int) bool { return bridges[i].north < bridges[j].north })

	lis := []int{}
	lis = append(lis, 1)
	for i := 1; i < len(bridges); i++ {
		var temp int = 1
		for j := 0; j < i; j++ {
			if bridges[j].south <= bridges[i].south {
				temp = maximum(temp, lis[j]+1)
			}
		}
		lis = append(lis, temp)
	}
	fmt.Println(lis)

}
