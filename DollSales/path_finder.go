package dolldelivery

// PathFinder must be implemented
type PathFinder func(startLocation, targetLocation string, streets []Street) (distance int, path []string)

// Street defines a connection between two houses
type Street struct {
	From     string
	To       string
	Distance int
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

// Create a list for next possible nodes.
func NextNodelist(from string, streets []Street) []Street {
	var arr []Street
	for _, v := range streets {
		if v.From == from {
			arr = append(arr, v)
		}
	}
	return arr
}
