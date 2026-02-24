package expenses
import (
    "errors"
    "fmt"
)
// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var filteredRecords []Record
    for _,value:= range in{
        if(predicate(value)){
            filteredRecords=append(filteredRecords,value)
        }
    }
    return filteredRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func( r Record) bool{
        if(p.From<=r.Day && r.Day<=p.To){
            return true 
        }
        return false
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
    return func( r Record) bool{
        if(r.Category==c){
            return true 
        }
        return false
    }
	
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    	filteredRecords:=Filter(in, ByDaysPeriod(p)) 
    	totalExpense:=0.0
	    for _,value:= range filteredRecords{
            totalExpense+=value.Amount
        }
    	return totalExpense
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var hasError=true
    totalExpense:=0.0
    for _,value:= range in{
        if(value.Category==c){
            hasError=false
            break
        }
    }
    if(hasError){
        errorString:=fmt.Sprintf("unknown category %s",c)
        return 0,errors.New(errorString)
    }
    filteredByPeriod:=Filter(in, ByDaysPeriod(p))     
	filteredByPeriodAndCategory:=Filter(filteredByPeriod, ByCategory(c))
                                        
	for _,value:= range filteredByPeriodAndCategory{
            totalExpense+=value.Amount
    }
    return totalExpense,nil
                                        
}
