type ParkingSystem struct {
    big, medium, small capacity
}

type capacity struct {
    limit, occupied int
}

func (this *capacity) addCar() bool {
    if this.occupied == this.limit {
        return false
    }
    this.occupied++
    return true
}

func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{
        big: capacity{limit:big},
        medium: capacity{limit:medium},
        small: capacity{limit:small},
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType {
        case 1:
            return this.big.addCar()
        case 2:
            return this.medium.addCar()
        case 3:
            return this.small.addCar()
    }
    return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */