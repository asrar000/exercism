package chance
import "math/rand"
// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    return rand.Intn(20) + 1
	panic("Please implement the RollADie function")
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    return rand.Float64()*12.0
	panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    s:=[]string {"ant","beaver","cat","dog","elephant","fox","giraffe","hedgehog"}
    
    rand.Shuffle(len(s), func(i, j int) {
	s[i], s[j] = s[j], s[i]
})
    return s
    
	panic("Please implement the ShuffleAnimals function")
}
