package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int)int{
    if(time==0){
        return len(layers)*2
    }
    return len(layers)*time
}

func Quantities(layers []string)(noodlesNeeded int,sauceNeeded float64){
    for i:=0;i<len(layers);i++{
        if(layers[i]=="noodles"){
            noodlesNeeded+=50
        }else if(layers[i]=="sauce"){
            sauceNeeded+=0.2
        }
    }
    return
    
}

func AddSecretIngredient(friendIngredient []string,myIngredient []string){
    myIngredient[len(myIngredient)-1]=friendIngredient[len(friendIngredient)-1]
}

func ScaleRecipe(amounts []float64,portions int) []float64{
    var scaledAmounts []float64
    for i:=0;i<len(amounts);i++{
        scaledAmounts=append(scaledAmounts,amounts[i]*(float64(portions)/2.0))
    }
    return scaledAmounts
}
// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
