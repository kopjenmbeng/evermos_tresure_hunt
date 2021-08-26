package main

import "fmt"

func main() {
	var treasure_posision []string
	var obstacle []string

	fmt.Println("...first grid initiate....")
	for y := 5; y >=0; y-- {
		for x := 0; x <= 7; x++ {
			if y==5 || y==0{
				fmt.Printf("%s","#")
				o:=fmt.Sprintf("%d.%d",x,y)
				obstacle=append(obstacle, o)
			} else if y==3 &&( x == 2 || x==3 || x==4){
				fmt.Printf("%s","#")
				o:=fmt.Sprintf("%d.%d",x,y)
				obstacle=append(obstacle, o)
			} else if y==2 &&( x == 4 || x==6){
				fmt.Printf("%s","#")
				o:=fmt.Sprintf("%d.%d",x,y)
				obstacle=append(obstacle, o)
			} else if y==1 &&( x == 2){
				fmt.Printf("%s","#")
				o:=fmt.Sprintf("%d.%d",x,y)
				obstacle=append(obstacle, o)
			}else if y==1 &&( x == 1){
				fmt.Printf("%s","X")
				o:=fmt.Sprintf("%d.%d",x,y)
				obstacle=append(obstacle, o)
			}else {
				if x==0 || x==7{
					fmt.Printf("%s","#")
					o:=fmt.Sprintf("%d.%d",x,y)
					obstacle=append(obstacle, o)
				}else {
					fmt.Printf("%s",".")
				}
				
			}
			
		}
		fmt.Println()
	}
	// get probabilty treasure coordinate
	for StartY := 1; StartY< 5; StartY++ {
		for StartX := 2; StartX<= 7; StartX++ {
			coor:=fmt.Sprintf("%d.%d",StartX,StartY)
			if IsExist(coor,obstacle){
				break
			}else {
				for startMinY := StartY-1; startMinY > 0; startMinY-- {
					coor:=fmt.Sprintf("%d.%d",StartX,startMinY)
					if IsExist(coor,obstacle){
						break
					}else {
						treasure_posision=append(treasure_posision, coor)
					}
				}
				// fmt.Printf("%s",pr)
			}
			// fmt.Println()
		}
		// fmt.Println()
	}

	fmt.Println("...output a list of probable coordinate points where the treasure might be located....")
	for _,i:= range treasure_posision{
		fmt.Println(i)
	}

	fmt.Println("...final grid where the treasure might be located....")
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
					coor:=fmt.Sprintf("%d.%d",x,y)
					if IsExist(coor,treasure_posision){
						fmt.Printf("%s","$")
					}else{
						fmt.Printf("%s",".")
					}
				}
				
			}
			
		}
		fmt.Println()
	}
}

func IsExist(value string,list []string)bool{
	for _,i:= range list{
		if i==value{
			return true
		}
	}
	return false
}