package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    if(balance>=0 && balance<1000){
        return 0.5
    }else if(balance>=1000 && balance<5000){
        return 1.621
    }else if(balance>=5000){
        return 2.475
    }
    return 3.213
	panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    if(balance>=0 && balance<1000){
        return float64(InterestRate(balance)/100)*balance
    }else if(balance>=1000 && balance<5000){
        return float64(InterestRate(balance)/100)*balance
    }else if(balance>=5000){
        return float64(InterestRate(balance)/100)*balance
    }
    return float64(InterestRate(balance)/100)*balance
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return balance+Interest(balance)
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    yearCount:=0
    for balance<targetBalance{
        balance=AnnualBalanceUpdate(balance)
        yearCount++
    }
    return yearCount
	panic("Please implement the YearsBeforeDesiredBalance function")
}
