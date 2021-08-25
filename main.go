package main

import "fmt"

func main() {
	var treasure_posision []string

	for y := 5; y >=0; y-- {
		for x := 0; x <= 7; x++ {
			if y==5 || y==0{
				fmt.Printf("%s","#")
			} else if y==3 &&( x == 2 || x==3 || x==4){
				fmt.Printf("%s","#")
			} else if y==2 &&( x == 4 || x==6){
				fmt.Printf("%s","#")
			} else if y==1 &&( x == 2){
				fmt.Printf("%s","#")
			}else if y==1 &&( x == 1){
				fmt.Printf("%s","X")
			}else {
				if x==0 || x==7{
					fmt.Printf("%s","#")
				}else {
					fmt.Printf("%s","$")
					coor:=fmt.Sprintf("X,Y=%d,%d",x,y)
					treasure_posision=append(treasure_posision, coor)
				}
				
			}
			
		}
		fmt.Println()
	}

	fmt.Println("...output a list of probable coordinate points where the treasure might be located....")
	for _,i:= range treasure_posision{
		fmt.Println(i)
	}
}