package pdq

import (
	"fmt"
)

func CalcInstability(outgoingDependenciesNum int, incomingDependenciesNum int) (float64, error) {
	if outgoingDependenciesNum < 0 || incomingDependenciesNum < 0 {
		return -1, fmt.Errorf("invalid metric value: outgoingDependenciesNum %v, incomingDependenciesNum %v; values must not be less than 0", outgoingDependenciesNum, incomingDependenciesNum)
	}

	totalDependencies := outgoingDependenciesNum + incomingDependenciesNum
	if totalDependencies == 0 {
		return 0, nil
	}

	instability := float64(outgoingDependenciesNum) / float64(totalDependencies)

	return instability, nil
}