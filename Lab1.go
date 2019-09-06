package main

import "fmt"
import "errors"
import "strings"




/*type errorOutofCapacity struct {
    message string
}
type errorOtherDestination struct {
    message string
}
type noTime struct {
    message string
}

func NewerrorOutofCapacity("Out of Capacity") *errorOutofCapacity{
    return &errorOu

}

var outOfcapacity error =  errors.New("Out of Capacity")
var otherDestination error = errors.New("Other destination")
var noTime error = errors.New("Not enough time")
*/

type Trip struct {
    destination string
    weight float32
    deadline int
}

type Vehicle struct{
    vehicle string
    name string
    destination string    
    speed float32   
    capacity float32
    loadToCarry float32
}

type Truck struct{ 
    Vehicle 
}

type PickUp struct{
    Vehicle
    isPrivate bool
}

type TrainCar struct{
    Vehicle
    railway string
}

func NewTruck(t* Truck)(a* Truck){
    t.name = "truck"
    t.vehicle = "truck"
    t.destination = ""
    t.speed= 40
    t.capacity = 10
    t.loadToCarry = 0
    return
}

func NewPickup(p *PickUp)(b *PickUp){
    p.name = "NewPickup"
    p.vehicle = "pickup"
    p.destination =""
    p.speed= 60
    p.capacity = 2
    p.loadToCarry = 0
    p.isPrivate = true   
    return
}

func NewTrainCar(c *TrainCar)(e *TrainCar){
    c.name = "trainCar"
    c.vehicle = "trainCar"
    c.destination =""
    c.speed= 30
    c.capacity = 30
    c.loadToCarry = 0
    c.railway = "CNR"
    return
}


type Transporter interface {
    addload(Trip) // Trip as an argument returning error if the trnasporter has insufficient capacity to carry the weight
    print()  // prints transporter to console
}


func (t *Truck) addLoad(a *Trip) error {
    //trip = Toronto

        if t.destination!= " " && t.destination != a.destination {
            return errors.New("Error: Other destination")
        }  

        if (t.loadToCarry + a.weight)> t.capacity {
            return errors.New("Error: Out of Capacity")
        }
        
        if a.destination == "Toronto" && (400/t.speed) > float32(a.deadline){
            return errors.New("Error: cannot reach on time")

        }
        if a.destination == "Montreal" && (200/t.speed) > float32(a.deadline){
            return errors.New("Error: cannot reach on time")
        }
        t.destination = a.destination
        t.loadToCarry +=  a.weight 
        t.loadToCarry +=  a.weight 
        return nil    
}

func (t *PickUp) addLoad(a *Trip) error{
       if t.destination!= " " && t.destination != a.destination {
            return errors.New("Error: Other destination")
        }  

        if (t.loadToCarry + a.weight)> t.capacity {
            return errors.New("Error: Out of Capacity")
        }
        
        if a.destination == "Toronto" && (400/t.speed) > float32(a.deadline){
            return errors.New("Error: cannot reach on time")

        }
        if a.destination == "Montreal" && (200/t.speed) > float32(a.deadline){
            return errors.New("Error: cannot reach on time")
        }
        t.destination = a.destination
        t.loadToCarry +=  a.weight  
        return nil    
}

func (t *TrainCar) addLoad(a *Trip) error{ 
        if t.destination!= " " && t.destination != a.destination {
            return errors.New("Error: Other destination")
        } 

        if (t.loadToCarry + a.weight)> t.capacity {
            return errors.New("Error: Out of Capacity")
        }
        
        if a.destination == "Toronto" && (400/t.speed) > float32(a.deadline){
            return errors.New("Error: cannot reach on time")

        }
        if a.destination == "Montreal" && (200/t.speed) > float32(a.deadline){
            return errors.New("Error: cannot reach on time")
        }
        t.destination = a.destination
        t.loadToCarry +=  a.weight 
        return nil    
}



func (t *Truck) print(){
    fmt.Println("Truck to",  t.destination, " with " , t.capacity ,"tons.")
}

func (p *PickUp) print(){
    fmt.Println("PickUp to ", p.destination,  " with ", p.capacity, "tons (Private: ", p.isPrivate)  
}

func (c *TrainCar) print(){
    fmt.Println("TrainCar to",  c.destination, "with", c.capacity,  "tons ", c.railway ) 
}


func (t *Trip) NewTorontoTrip(weight float32, deadline int, destination string ){
    t.weight = weight
    t.deadline = deadline
    t.destination = "Toronto"    

}

func (m *Trip) NewMontrealTrip(weight float32, deadline int, destination string){
    m.weight = weight
    m.deadline = deadline
    m.destination = "Montreal"
    
}



func main() {
    var destination string
      var weight float32
      var deadline int



    truck1:= Truck{}
    truck2:= Truck{}
    pickup1:= PickUp{}
    pickup2:= PickUp{}
    pickup3:= PickUp{}
    trainCar1:= TrainCar{}


    NewTruck(&truck1) 
    NewTruck(&truck2)
    NewPickup(&pickup1)
    NewPickup(&pickup2)  
    NewPickup(&pickup3)
    NewTrainCar(&trainCar1)
    
    /*a:=Trip{}
    b:= Trip{} 
    c:=Trip{}
    d:= Trip{} */
    
   /* a.NewTorontoTrip(8,12)
    b.NewTorontoTrip(8,20)
    c.NewMontrealTrip(8,12)
    d.NewMontrealTrip(5,3)*/

    //Users/Momah/Desktop/ParadigmsLabs

    
    var listofTrips[] Trip 
    var transport[6] Transporter
    //count := 0
    
    transport[0] = &truck1
    transport[1] = &truck2
    transport[2] = &pickup1
    transport[3] = &pickup2
    transport[4] = &pickup3
    transport[5] = &TrainCar
    
     for{
            fmt.Print("Destination: (t)oronto, (m)ontreal, else exit? ")
                n, err := fmt.Scanf("%s\n", &destination)
                if n != 1 || err != nil {
                  fmt.Println(n, err)
                  return
                }
                destination = strings.ToLower(destination)

            if destination == "Quit" || destination == "quit" || destination == "q" || destination == "Q"{
               fmt.Println("Not going to Toronto or Montreal byeee ! ")
               break
            } 

            if !strings.HasPrefix("toronto", destination) &&  !strings.HasPrefix("montreal", destination) {
                fmt.Println("Not going to TO or Montreal, bye!")
                break
            }

            fmt.Print("Weight: ")
            m, error1 := fmt.Scanf("%g\n", &weight)
            if m != 1 || error1 != nil {
                fmt.Println(m, error1)
            return
            }

            fmt.Print("Deadline (in hours): ")
            n2, error2 := fmt.Scanf("%d\n", &deadline)
            if n2 != 1 || error2 != nil {
                fmt.Println(n2, error2)
                 return
            }

           var trip *Trip
           if  strings.HasPrefix("toronto", destination) {
                trip = NewTorontoTrip(weight, deadline)      
            }
            if strings.HasPrefix("montreal", destination) {
                trip = NewMontrealTrip(weight, deadline)      
            }
            listofTrips = append(trips, *trip)
            for i := 0 ; i < len(transport) ; i++ {
                 err := transport[i].addload(trip)
                 if err != nil {
                      fmt.Println(err)
                    }
                if err == nil {
                 break
            }
        }
    }
    fmt.Println("Trips: ", listofTrips)
     for i:=0; i<len(transport); i++{
        transport[i].print()
    } 
}








        


        

            




        


        
    




    
      
    
    

        










