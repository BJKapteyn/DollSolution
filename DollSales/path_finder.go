package main

// PathFinder must be implemented
type PathFinder func(startLocation, targetLocation string, streets []Street) (distance int, path []string)

// Street defines a connection between two houses
type Street struct {
	From     string
	To       string
	Distance int
}

func main() {
	var FindPath PathFinder = func(startLocation, targetLocation string, streets []Street) (distance int, path []string) {
		// Keep original array intact
		neighborhood := streets
		var nodesToCheck = PossiblePathList(startLocation, neighborhood)
		var paths [][]Street
		// Keep track of this to know when to add a new path array or when to add to existing one
		var addedPaths = 0

		// Create inital array of paths to be added to later
		for _, s := range nodesToCheck {
			path := []Street{s}
			append(paths, path)
		}

	}

}

// For creating the inital array of booleans representing finished or deadend paths
func FoundTargetList(numberOfStartingStreets int) []bool {
	var boolArr [numberOfStartingStreets]bool

	for i := 0; i < numberOfStartingStreets; i++ {
		boolArr[i] := false
	}
	return boolArr
}

// Find target in array or streets stemming from node
func FindStreet(streets []Street, target string) bool {
	for _, v := range streets {
		if v.To == target {
			return true
		}
	}
	return false
}

// Calculate distances of a full path
func CalculateDistance(path []Street) {
	var distance = 0

	for _, v := range path {
		distance += v.Distance
	}
}

// For removing a node from the neighborhood list after checking available neighboring nodes
func RemoveNode(street string, streets []Street) {
	var arr []Street

	for _, v := range streets {
		if v.To != street {
			append(arr, v)
		}
	}
}

// Create a list for next possible nodes
func PossiblePathList(from string, streets []Street) []Street {
	var arr []Street

	for _, v := range streets {
		if v.From == from {
			arr = append(arr, v)
		}
	}
	return arr
}
