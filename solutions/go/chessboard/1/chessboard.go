package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool
type Chessboard map[string] []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count:=0
    for _,value:= range cb[file] {
    	if(value){
            count+=1
        }
    }
    return count
	
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if (rank<1 || rank>8){
        return 0
    }
    	
    count:=0
    for _,value:= range cb {
    	if(value[rank-1]){
            count+=1
        }
    }
    return count
	
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count:=0
    for _,x:=range cb{
        for range x{
            count++
        }
    }
    return count
	
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    count:=0
    for _,value:=range cb{
        for _,occupied:=range value{
            if(occupied){
                count++
            }
        }
    }
    return count
	
}
