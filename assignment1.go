package main

import (
  "fmt"
  "errors"
  "strings"
)


type Trip struct {
  destination string
  weight float32
  deadline int
}


type Truck struct {
  vehicle, name, destination string
  speed float32
  capacity float32
  load float32
}

type Pickup struct {
  vehicle, name, destination string
  speed float32
  capacity float32
  load float32
  isPrivate bool
}


type TrainCar struct {
  vehicle, name, destination string
  speed float32
  capacity float32
  load float32
  railway string
}




func NewTruck() Truck {
  t := Truck{name: "Truck", vehicle: "Truck", speed: 40, capacity: 10}
  return t
}

// Return a new Pickup with default initializers
func NewPickUp() Pickup {
  p := Pickup{name: "Pickup", vehicle: "Pickup", speed: 60, capacity: 2, isPrivate: true}
  return p
}

// Return a new TrainCar with default initializers
func NewTrainCar() TrainCar {
  c := TrainCar{name: "TrainCar", vehicle: "TrainCar", speed: 30, capacity: 30, railway: "CNR"}
  return c
}

// interface Transporter
type Transporter interface{
  addLoad(trip Trip) error
  print()
}

// Create a trip to toronto
func NewTorontoTrip(weight float32, deadline int) *Trip {
  trip := Trip{weight: weight, destination: "Toronto", deadline: deadline}
  return &trip
}

// Create a trip to montreal
func NewMontrealTrip(weight float32, deadline int) *Trip {
  trip := Trip{weight: weight, destination: "Montreal", deadline: deadline}
  return &trip
}

// Implements methods of the interface
func (t *Truck) addLoad(trip Trip) error {
  if t.destination != "" && t.destination != trip.destination {
    return errors.New("Error: Other destination")
  }
  if (t.load + trip.weight) > t.capacity {
    return errors.New("Error: Out of capacity")
  }  
  if trip.destination == "Toronto" && int(400/t.speed) > trip.deadline {
    return errors.New("Error: cannot reach on time")
  }
  if trip.destination == "Montreal" && int(200/t.speed) > trip.deadline {
    return errors.New("Error: cannot reach on time")
  }
  t.destination = trip.destination
  t.load += trip.weight
  return nil
}

func (p *Pickup) addLoad(trip Trip) error {
  if p.destination != "" && p.destination != trip.destination {
    return errors.New("Error: Other destination")
  }
  if (p.load + trip.weight) > p.capacity {
    return errors.New("Error: Out of capacity")
  }  
  if trip.destination == "Toronto" && int(400/p.speed) > trip.deadline {
    return errors.New("Error: cannot reach on time")
  }
  if trip.destination == "Montreal" && int(200/p.speed) > trip.deadline {
    return errors.New("Error: cannot reach on time")
  }
  p.destination = trip.destination
  p.load += trip.weight
  return nil
}

func (tc *TrainCar) addLoad(trip Trip) error {
  if tc.destination != "" && tc.destination != trip.destination {
    return errors.New("Error: Other destination")
  }
  if (tc.load + trip.weight) > tc.capacity {
    return errors.New("Error: Out of capacity")
  } 
  if trip.destination == "Toronto" && int(400/tc.speed) > trip.deadline {
    return errors.New("Error: cannot reach on time")
  }
  if trip.destination == "Montreal" && int(200/tc.speed) > trip.deadline {
    return errors.New("Error: cannot reach on time")
  }
  tc.destination = trip.destination
  tc.load += trip.weight
  return nil
}

func (t *Truck) print() {
  fmt.Printf("Truck %v to %v with %v tons\n", t.name, t.destination, t.load)
}

func (t *Pickup) print() {
  fmt.Printf("Pickup %v to %v with %v tons (Private: %v)\n", t.name, t.destination, t.load, t.isPrivate)
}

func (t *TrainCar) print() {
  fmt.Printf("TrainCar %v to %v with %v tons (%v)\n", t.name, t.destination, t.load, t.railway)
}

func main() {
  // array to store all vehicles
  var vehicles[6] Transporter

  // list to store all trips
  var trips[] Trip


  t1 := NewTruck()
  t1.name = "A"
  vehicles[0] = &t1
  t2 := NewTruck()
  t2.name = "B"
  vehicles[1] = &t2

 
  pickupA := NewPickUp()
  pickupA.name = "A"
  vehicles[2] = &pickupA
  pickupB := NewPickUp()
  pickupB.name = "B"
  vehicles[3] = &pickupB
  pickupC := NewPickUp()
  pickupC.name = "C"
  vehicles[4] = &pickupC


  trainCar := NewTrainCar()
  trainCar.name = "A"
  vehicles[5] = &trainCar


  var dest string
  var weight float32
  var deadline int
  for {
    fmt.Print("Destination: (t)oronto, (m)ontreal, else exit? ")
    n, err := fmt.Scanf("%s\n", &dest)
    if n != 1 || err != nil {
      fmt.Println(n, err)
      return
    }
    dest = strings.ToLower(dest)

    if !strings.HasPrefix("toronto", dest) &&  !strings.HasPrefix("montreal", dest) {
      fmt.Println("Not going to TO or Montreal, bye!")
      break
    }

    fmt.Print("Weight: ")
    n1, err1 := fmt.Scanf("%g\n", &weight)
    if n1 != 1 || err1 != nil {
      fmt.Println(n1, err1)
      return
    }
    fmt.Print("Deadline (in hours): ")
    n2, err2 := fmt.Scanf("%d\n", &deadline)
    if n2 != 1 || err2 != nil {
      fmt.Println(n2, err2)
      return
    }
    
    var trip *Trip
    if  strings.HasPrefix("toronto", dest) {
      trip = NewTorontoTrip(weight, deadline)      
    }
    if strings.HasPrefix("montreal", dest) {
      trip = NewMontrealTrip(weight, deadline)      
    }
    trips = append(trips, *trip)
    for i := 0 ; i < len(vehicles) ; i++ {
      err := vehicles[i].addLoad(*trip)
      if err != nil {
        fmt.Println(err)
      }
      if err == nil {
       break
      }
    }
  }

  fmt.Println("Trips: ", trips)


  fmt.Println("Vehicles:")
  for i := 0 ; i < len(vehicles); i++ {
    vehicles[i].print()
  }
}