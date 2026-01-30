package partyrobot
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
    return "Welcome to my party, "+name+"!"
	panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    return fmt.Sprintf("Happy birthday %s! You are now %d years old!",name,age)
	panic("Please implement the HappyBirthday function")
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    var welcomeMessage string="Welcome to my party, "+name+"!"
    var seat string="You will be sitting next to "+neighbor+"."
    var d=fmt.Sprintf("%0.1f",distance)
    var tableName=fmt.Sprintf("%d",table)
	if(len(tableName)<3){
        if(len(tableName)==1){
            tableName="00"+tableName
        }else{
            tableName="0"+tableName
        }
    }
    return welcomeMessage+"\n"+"You have been assigned to table "+tableName+". Your table is "+direction+", exactly "+d+" meters from here.\n"+seat
	panic("Please implement the AssignTable function")
}
