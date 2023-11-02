type UndergroundSystem struct {
	stations map[string]map[string][]int
	ongoing  map[int]*UndergroundTravelRoute
}

type UndergroundTravelRoute struct {
	timeIn    int
	stationIn string
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		stations: make(map[string]map[string][]int),
		ongoing:  make(map[int]*UndergroundTravelRoute),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	route := &UndergroundTravelRoute{
		timeIn:    t,
		stationIn: stationName,
	}
	this.ongoing[id] = route
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	route := this.ongoing[id]
	stationIn := route.stationIn
	stationOut := stationName
	timeTaken := t - route.timeIn
	if _, ok := this.stations[stationIn]; !ok {
		this.stations[stationIn] = make(map[string][]int)
	}
	if _, ok := this.stations[stationIn][stationOut]; !ok {
		this.stations[stationIn][stationOut] = make([]int, 0)
	}
	this.stations[stationIn][stationOut] = append(this.stations[stationIn][stationOut], timeTaken)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	return average(this.stations[startStation][endStation])
}

func average(slice []int) float64 {
	var sum float64
	for _, v := range slice {
		sum += float64(v)
	}
	return sum / float64(len(slice))
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */