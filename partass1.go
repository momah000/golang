package main

import("fmt"
       "time"
       "sync"
       "math/rand")

const (
NumRoutines = 3
NumRequests = 1000
)
// global semaphore monitoring the number of routines
var semRout = make(chan int, NumRoutines)

// global semaphore monitoring console
var semDisp = make(chan int, 1)//display channel
// Waitgroups to ensure that main does not exit until all done




//wait group- its a counter
//put wg.add with main to increase counter and put wg.done to decrement by one
//wg.Wait() - so all go routines finish running //Wg.add adds it
//wd,done()
var wgRout sync.WaitGroup

var wgDisp sync.WaitGroup



type Task struct {
	a, b float32
	disp chan float32
}
//A function that sleeps for a random time between 1 and 15 seconds, adds
//the numbers a and b and sends the result on the display channel. 
//solves the task and sends it to the result on the display channel
func solve(t* Task){
	timer := rand.Intn(15)
	time.Sleep(time.Duration(timer) * time.Second)

	res :=t.a + t.b
	t.disp <- res
	
}
//A function that acts as intermediary between ComputeServer and
//solves
//uses solve(function call) on the task
func handleReq(t* Task){
	solve(t)
	<-semRout
	//wgRout.done()
}

//A A function that uses the channel factory pattern
//(lambda) and listens for requests on the created channel for tasks. It calls the handleReq function.

//receives tasks from the main and creates handleReq in seperate go routine(max 3 routines)
func ComputeServer() (chan *Task){

	channel := make(chan *Task)
	
	
//waitgroupout.add(1)
	go func(){
		 // semaphore for when the goroutine is full
		//tke an input incase code doesnt work (task, ok) take channel
		for {
			task, ok := <-channel
			if !ok{
				break
			}
			semRout<-1
			//task = <- channel
			wgRout.Add(1)
			go handleReq(task)
		}
		wgRout.Done()	
	}()
		
	
	
	return channel
}

//A function that uses the channel factory pattern
//(lambda) and listens for requests on the created channel for results to print to the console.
//whenever a result is received, result is displayed
func  DisplayServer() (chan float32) {
    
	channel2 := make(chan float32)
   
    wgDisp.Add(1)
		

		go func(){
	        for {
	        	result , ok:= <- channel2
	        	if !ok{
	        		break
	        	}
	        	//semDisp<-1
                //fmt.Println("--------")
	        	fmt.Println(result)
	        	//<-semDisp
	        	wgRout.Done()
	        }
	        wgDisp.Done()		
		}()
		

	return channel2
}




//main
// user input
// create task
//send task to computer on a chennel
func main() {
	dispChan := DisplayServer()
	reqChan := ComputeServer()
	for {
		var a, b float32
		semDisp<-1
 // make sure to use semDisp
// …     
 		fmt.Print("Enter two numbers: ")
		fmt.Scanf("%f %f \n", &a, &b)
		fmt.Printf("%f %f \n", a, b)
		if a == 0 && b == 0 {
			break
		}
 		// Create task and send to ComputeServer
 	
	    mainTask := Task{}
	 	mainTask.a = a
 		mainTask.b = b
 		mainTask.disp = dispChan
 		reqChan<- &mainTask //&
 		time.Sleep(1e9)//1e9
 		<-semDisp
 	}
 	  <-semDisp
 	  close(reqChan)
 	  wgRout.Wait()
 	  close(dispChan)
 	  wgDisp.Wait()
 	  close(semRout)
 	  close(semDisp)


 }

	
 // Don’t exit until all is done
