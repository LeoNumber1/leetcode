package main

import "fmt"

func main() {
	//["ParkingSystem","addCar","addCar","addCar","addCar"]
	init := [][]int{{1, 1, 0}, {1}, {2}, {3}, {1}}
	ps := Constructor(init[0][0], init[0][1], init[0][2])
	for i := 1; i < len(init); i++ {
		fmt.Println(ps.AddCar(init[i][0]))
	}
}

//52 ms-100.00%	7.2 MB-100.00%
type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > 0 {
			this.big--
			return true
		} else {
			return false
		}
	case 2:
		if this.medium > 0 {
			this.medium--
			return true
		} else {
			return false
		}
	case 3:
		if this.small > 0 {
			this.small--
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
