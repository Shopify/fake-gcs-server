package backend

import "fmt"

// Conditions constrain storage store methods to act on specific generations of objects.
//
// The zero value is an empty set of constraints.
type Conditions struct {
	// GenerationMatch specifies that the object must have the given generation for the operation to occur.
	// If GenerationMatch is zero, it has no effect.
	GenerationMatch int64
}

func validWithConditions(obj Object, conditions Conditions) error {
	if conditions.GenerationMatch != 0 && obj.Generation != conditions.GenerationMatch {
		return fmt.Errorf("Precondition Failed, conditionNotMet")
	}
	return nil
}

func emptyConditions(conditions Conditions) bool {
	return conditions.GenerationMatch == 0
}
