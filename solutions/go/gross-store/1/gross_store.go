package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    unitOfMeasurement:=map[string]int{}
    unitOfMeasurement["quarter_of_a_dozen"]=3
    unitOfMeasurement["half_of_a_dozen"]=6
    unitOfMeasurement["dozen"]=12
    unitOfMeasurement["small_gross"]=120
    unitOfMeasurement["gross"]=144
    unitOfMeasurement["great_gross"]=1728

    return unitOfMeasurement
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return map[string]int{}
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    _, unitExist := units[unit]
    if(!unitExist){
        return false
    }
    _,itemExist := bill[item]
    if(itemExist){
        bill[item]+=units[unit]
        return true
    }
    bill[item]=units[unit]
    return true
    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    _, unitExist := units[unit]
    _,itemExist := bill[item]
    if(!unitExist||!itemExist||bill[item]-units[unit]<0){
        return false
    }else if(bill[item]-units[unit]==0){
        delete(bill,item)
    }else{
        bill[item]-=units[unit]
    }
    return true 
	
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    quantity,ifExist:=bill[item]
    return quantity,ifExist
	
}
