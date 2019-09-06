package main

import "fmt"
import "math"
import "sync"
import "math/rand"

type Stack struct{
	carryStuff [] Triangle
}

/*unc (s *Stack) add(t *Triangle) {
	a := s.carryStuff
	a = append(a, &t)
}*/
//func

type Point struct {
 x float64
 y float64
}


type Triangle struct {
 A Point
 B Point
 C Point
}

func (t Triangle) Perimeter() float64 {
	axvalue := t.A.x
	ayvalue := t.A.y

	bxvalue := t.B.x
	byvalue := t.B.y

	cxvalue := t.C.x
	cyvalue := t.C.y

	a:= math.Sqrt( ((axvalue-bxvalue)*(axvalue-bxvalue)) + ((ayvalue-byvalue)*(axvalue-byvalue)) ) //between a  and b
	b:= math.Sqrt( ((bxvalue-cxvalue)*(bxvalue-cxvalue)) + ((byvalue-cyvalue)*(byvalue-cyvalue)) )  //b andc
	c:= math.Sqrt( ((cxvalue-axvalue)*(cxvalue-axvalue)) + ((cyvalue-ayvalue)*(cyvalue-ayvalue)) ) // c and a

	return a + b + c

}


func (t Triangle) Area() float64{
	axvalue := t.A.x
	ayvalue := t.A.y

	bxvalue := t.B.x
	byvalue := t.B.y

	cxvalue := t.C.x
	cyvalue := t.C.y

	a := bxvalue - axvalue
	d := cyvalue - ayvalue
	b := cxvalue - axvalue 
	c := byvalue - ayvalue

	ans := 0.5*((a*d)-(b*c))

	return math.Abs(ans)
}

//give a result of type stack
func triangles10000() (result [10000]Triangle) {
	 rand.Seed(2120)
	 for i := 0; i < 10000; i++ {
		 result[i].A= Point{rand.Float64()*100.,rand.Float64()*100.}
		 result[i].B= Point{rand.Float64()*100.,rand.Float64()*100.}
		 result[i].C= Point{rand.Float64()*100.,rand.Float64()*100.}
	 }
	 return
}

var wgGroup sync.WaitGroup

var hiCh = make(chan float64, 1)
var liCh = make(chan float64, 1)


func classifyTriangles(highRatio *Stack, lowRatio *Stack, ratioThreshold float64, triangles []Triangle){
	
   
	for _, t := range triangles{
		ratio := t.Perimeter()/t.Area()
		if ratio > ratioThreshold {
			hiCh <- 1
			highRatio.carryStuff = append(highRatio.carryStuff, t)    
			<-hiCh
		}else{
			liCh<- 1
			lowRatio.carryStuff = append(lowRatio.carryStuff, t) 
			<-liCh
		}
	}
	
	wgGroup.Done()
}




func main() {

	hr := Stack{}
	lr := Stack{}
	
	triArray:=triangles10000() //made for when made we need to put it in a stack 
	
	for i:=0; i<10; i++ {
		
		//fmt.Println("hello")
		triArraySlice := triArray[(i*1000) : ((i+1)*1000)]
		

		wgGroup.Add(1)
		
		go classifyTriangles(&hr, &lr, 1.0, triArraySlice) 
		
	}

	wgGroup.Wait()
	fmt.Println("number of triangles in the stack bigger than the ratio threshold : ", len(hr.carryStuff) )
	fmt.Println("number of triangles in the stack smaller than the ratio threshold : ", len(lr.carryStuff) )
	fmt.Println("item on top on the smaller stack: ", lr.carryStuff[ (len(lr.carryStuff)-1) ] )
	fmt.Println("item on top on the larger stack: ", hr.carryStuff[ (len(hr.carryStuff)-1) ] )
}



