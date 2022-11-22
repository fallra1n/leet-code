package main

// import "fmt"

type Void struct{}

func containsDuplicate(nums []int) bool {
    var void Void
    Set := make(map[int]Void)

    for _, num := range nums {
        if _, ok := Set[num]; ok {
            return true
        }
        Set[num] = void
    }

    return false
}

// func main() {
// 	a := []int{1, 2, 3, 4};
// 	fmt.Println(containsDuplicate((a)))	
// }
