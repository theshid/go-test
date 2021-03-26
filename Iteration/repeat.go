package Iteration

func Repeat(input string, iteration int)string{
	var repeated string
	for i:=0;i<iteration ;i++  {
       repeated+=input
	}
	return repeated
}

