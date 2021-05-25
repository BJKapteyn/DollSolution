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

	//Fan out from the root location, removing nodes as they are checked
	var FindPath PathFinder = func(startLocation, targetLocation string, streets []Street) (distance int, path []string) {
		// Keep original array intact
		neighborhood := streets
		var nodesToCheck = PossiblePathList(startLocation, neighborhood)
		var paths [][]Street
		// Keep track of this to know when to add a new path array or when to add to existing one
		var addedPaths = 0
		var currentSearchLevel = 0

		// Create inital array of paths to be added to later
		for _, s := range nodesToCheck {
			path := []Street{s}
			append(paths, path)
		}
		// Keep track of how many nodes I'm currently looking at
		var totalNodes = len(nodesToCheck)
		var pathCompletionStatus = FoundTargetList(len(nodesToCheck))

		// Loop through the list removing nodes as it checks them
		for len(neighborhood) > 0 {
			var streetToCheck = start;
			var streetToRemove = ""
			var nextStreetToRemove = ""
			var nextPath = ""
			var totalAddedPaths = 0

			//Loop through possible streets from root
			for i := 0; i < len(nodesToCheck); i++ {
				// Skip if path is complete or deadend or if target is found, add it to path
				if pathCompletionStatus {
					continue
				} else if (HasTarget(nodesToCheck)) {
					var foundTarget street = FindStreet(nodesToCheck, target);
					append(paths[i], foundTarget)
					pathCompletionStatus[i] = true
				}

				//loop through inner nodes
				for j := 0; j < totalNodes; j++ {
					//figure out if the path leg needs a new path or if it can be added to existing path
					if j > currentSearchLevel {
						newPath := paths[i]
						append(newPath, nodesToCheck[j])
						//Add completion status for new path
						append(pathCompletionStatus, false)
					} else {
						append(paths[j], nodesToCheck[j])
					}
				}

				if i != totalNodes {
					if(nextStreetToRemove != "") {
						streetToRemove = nextStreetToRemove
					} else {
						streetToRemove = paths[i][len(paths[i]) - 1].From
					}
					// Reset
					var nextIndex = 0
					var nextLegIndexStart []Street = paths[nextIndex]
					var nextStreet Street = paths[nextIndex][nextLegIndexStart]
					nextPath = nextNode.To
					nextStreetToRemove = nextPath;
					neighborhood = RemoveNode(streetToRemove, neighborhood)
					nodesToCheck = PossiblePathList(nextPath, nCopy)
				}
			}
			totalAddedPaths += addedPaths
			totalNodes += addedPaths
			addedPaths = 0
			currentLevel++
		}
		// From here I wanted to remove the paths that didnt end at the target, 
		// tally up the distances for each, and then return the one with the least distance
		almost := [...]String {"Thanks for", " the fun", " weekend."}
		return 0, almost
	}

}

// For creating the inital array of booleans representing finished or deadend paths
func FoundTargetList(numberOfStartingStreets int) []bool {
	var boolArr [numberOfStartingStreets]bool

	for i := 0; i < numberOfStartingStreets; i++ {
		boolArr[i] = false
	}
	return boolArr
}

// Find target in streets stemming from a node
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

func HasTarget(possibleStreets []Street, target string) {
	for _, s := possibleStreets {
		if s.To == target {
			return true
		}
	}
	return false
}
