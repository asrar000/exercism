package thefarm

import (
	"fmt"
	"errors"
)
// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	totalFodder, err := fc.FodderAmount(numberOfCows)
	if err != nil {
		return 0.0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return (totalFodder / float64(numberOfCows)) * factor, nil
}
func ValidateInputAndDivideFood(fc FodderCalculator,numberOfCows int)(float64,error){
    if(numberOfCows<=0){
        return 0.0,errors.New("invalid number of cows")
    }
    return DivideFood(fc,numberOfCows)
}
    
// TODO: define the 'ValidateInputAndDivideFood' function

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct{
     numberOfCows int
     customMessage string
}

func (e *InvalidCowsError) Error() error{
    return errors.New(fmt.Sprintf("%d cows are invalid: %s",e.numberOfCows,e.customMessage))
}
func ValidateNumberOfCows(numberOfCows int) error{
    if(numberOfCows<0){
        s := InvalidCowsError {
    	numberOfCows: numberOfCows,
    	customMessage: "there are no negative cows",
		}
        return s.Error()
    }else if(numberOfCows==0){
        s := InvalidCowsError {
    	numberOfCows: numberOfCows,
    	customMessage: "no cows don't need food",
		}
        return s.Error()
    }
    
    return nil
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
