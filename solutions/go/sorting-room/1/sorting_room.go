package sorting

import (
	"fmt"
	"strconv"
    
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
    return fmt.Sprintf("This is the number %.1f",f)

}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
    return fmt.Sprintf("This is a box containing the number %.1f",float32(nb.Number()))
	
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    number,assertOkay:=fnb.(FancyNumber)
    if(assertOkay){
        value,_ := strconv.Atoi(number.Value())
        return value
    }
    return 0
    
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    number,assertOkay:=fnb.(FancyNumber)
    if(assertOkay){
        value,_:=strconv.ParseFloat(number.Value(), 64)
        return fmt.Sprintf("This is a fancy box containing the number %.1f",value)
    }
    return "This is a fancy box containing the number 0.0"
	
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
    valueInt,assertInt:=i.(int)
    valueFloat64,assertFloat64:=i.(float64)
    valueNumberBox,assertNumberBox:=i.(NumberBox)
    valueFancyNumberBox,assertFancyNumberBox:=i.(FancyNumberBox)
    if(assertInt || assertFloat64){
        if(assertInt){
            return DescribeNumber(float64(valueInt))
    	}
        return DescribeNumber(valueFloat64)
    }else if(assertNumberBox){
         return DescribeNumberBox(valueNumberBox)
    }else if(assertFancyNumberBox){
        return DescribeFancyNumberBox(valueFancyNumberBox)
    }
    return "Return to sender"
	
}
