package main

import ("fmt")

func main()  {
	for i := 0; i < 5; i++ {
		fmt.Println(i)	
	}
	fmt.Println("------------")
	for i := 0; i < 5; i++ {
		if i > 3{
			break;
		}
		fmt.Println(i)
	}

	fmt.Println("------------")
	for i := 0; i < 5; i++ {
		if i < 3{
			continue;
		}
		fmt.Println(i)
	}

	fmt.Println("------------")
	a := 0
	for a<3  {
		fmt.Println(a)
		a++
	}

}