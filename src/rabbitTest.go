package main

//假设兔子出生二个月后，也就是第三个月初，有繁衍能力，每个月生一对兔子；一对刚出生的兔子，假设兔子都不死，一年能生多少兔子
import (
	"fmt"
)

//12month
var Time int = 12

type rabbing struct {
	old int
}

func main() {
	rabbingList := make([]rabbing, 10000)
	lengthOld := 1
	LengthNew := 1
	rabbingList[0] = rabbing{1}
	for i := 1; i <= Time; i++ {
		for j := 0; j < lengthOld; j++ {
			if rabbingList[j].old > 2 {
				rabbingList[LengthNew] = rabbing{2}
				LengthNew++
			}
			rabbingList[j].old++
		}
		lengthOld = LengthNew
		fmt.Println(LengthNew)
	}

}
