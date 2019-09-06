
package main

//Import required libraries
import (
    "bufio"
    "container/list"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
	"strings"
)


var deliverzero = deliver{}


type deliver struct {
    n, Cost int
    p, q int
}

//structure for problem
type data struct {
    fname string
    supply, demand []int
    cost_data [][]int
    set [][]deliver
}


func error_handler(error error) {
	
    if error != nil {
		log.Fatal(error)
    }
}


func Min(p, k int) int {
	
    if p < k {
	
        return p
    }

    return k
}

//main function
func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("\nEnter name of input file: ")
	
    scanner.Scan()
	fname := scanner.Text()
	

	fmt.Print("\nEnter name of file containing initial solution: ")

    scanner.Scan()
	initfname := scanner.Text()
	
	
    t := setdata(fname,initfname)
	
    t.initial_solution(initfname)

    t.SteppingStoneMethod(initfname)

	t.display()

	t.write_file(fname,"solution.txt")
}


func countLines(fname string) int {
	//open file
	f, error := os.Open(fname)
	//handle error if any
	error_handler(error)
	
	//store file in buffer
	fscan := bufio.NewScanner(f)
	q := 0
	//scan file line by line
	for fscan.Scan() {
		q++
	}
	
	//return number of lines
	return q
}

//function to count number of words in first line of file
func countWords(fname string) int {
	//open file
	f, error := os.Open(fname)
	//handle error
	error_handler(error)
	
	//store file in buffer
	fscan := bufio.NewScanner(f)
	
	//split file into lines
	fscan.Split(bufio.ScanLines)
	
	//scan line
	fscan.Scan()
	
	//store words in line
	words := strings.Fields(fscan.Text())
	
	//count and return number of words
	return len(words)
}

//function to read file and store the data into structures
func setdata(fname string,initfname string) *data {
	
	//open file
    f, error := os.Open(fname)
	
	//handle error
    error_handler(error)
	
	//close file
    defer f.Close()
	
	//store file into buffer
    fscan := bufio.NewScanner(f)
	
	//split into words
    fscan.Split(bufio.ScanWords)
	
	//count lines and words
    s := countLines(fname) - 2
    d := countWords(fname) - 2
	
	//skip first line
	for p := 0; p < countWords(fname); p++ {
		fscan.Scan()
	}
	
	//define integer arrays for supply, demand and costs
 	supply := make([]int, s)
	demand := make([]int, d)
	cost_data := make([][]int, s)
	
	//create array for each supply
    for p := 0; p < s; p++ {
        cost_data[p] = make([]int, d)
    }
	
	//read and store cost data
	//for each supply
	for p := 0; p < s; p++ {
		fscan.Scan()
		//for each destination
		for k := 0; k < d; k++ {
			fscan.Scan()
			//read and store cost
			cost_data[p][k], error = strconv.Atoi(fscan.Text())
			
			//handle error
			error_handler(error)
		}
		//scan supply
		fscan.Scan()
        supply[p], error = strconv.Atoi(fscan.Text())
        error_handler(error)
	}
	fscan.Scan()
	
	//for each demand
	for p := 0; p < d; p++ {
		fscan.Scan()
		//read and store demand for each destination
		demand[p], error = strconv.Atoi(fscan.Text())
        error_handler(error)
	}
    
	//create 2d array
    set := make([][]deliver, len(supply))
	//create array for each supply 
    for p := 0; p < len(supply); p++ {
        set[p] = make([]deliver, len(demand))
    }
    //data read from file
    return &data{fname, supply, demand, cost_data, set}
}

//function to create initial solution
func (t *data) initial_solution(initfname string) {
	//for each supply
    for p, unused := 0, 0; p < len(t.supply); p++ {
		//for each demand left
        for q := unused; q < len(t.demand); q++ {
			//find minimum
            n := Min(t.supply[p], t.demand[q])
			//supply or demand is greater than zero
            if n > 0 {
				//set data
                t.set[p][q] = deliver{int(n), t.cost_data[p][q], p, q}
                t.supply[p] -= n
                t.demand[q] -= n
				
				//if supply is zero
                if t.supply[p] == 0 {
                    unused = q
					//break
                    break
                }
            }
        }
    }
}

//function to convert data to list
func (t *data) to_list() *list.List {
	//define new line
    l := list.New()
	
	//for each element in set
    for _, m := range t.set {
        for _, s := range m {
			//if s is not zero
            if s != deliverzero {
                l.PushBack(s)
            }
        }
    }
    return l
}

//function to find next non zero element
func (t *data) Next(s deliver, list *list.List) [2]deliver {
    var next [2]deliver
	//for each element
    for i := list.Front(); i != nil; i = i.Next() {
        a := i.Value.(deliver)
		//if element is not zero
        if a != s {
			//search for next zero element
			//if next element is zero
            if a.p == s.p && next[0] == deliverzero {
                next[0] = a
            }else if a.q == s.q && next[1] == deliverzero {
                next[1] = a
            }
			//if zero element not found
            if next[0] != deliverzero && next[1] != deliverzero {
                break
            }
        }
    }
    return next
}

//function to find marginal path
func (t *data) marginalpath(s deliver) []deliver {
	
	//create list of data
    path := t.to_list()
    path.PushFront(s)
 
	//create pointer
    var next *list.Element
	//loop
    for {
        deleted := 0
		//traverse path
        for i := path.Front(); i != nil; i = next {
            next = i.Next()
            next := t.Next(i.Value.(deliver), path)
			
			//if next is zero 
            if next[0] == deliverzero || next[1] == deliverzero {
				//delete path
                path.Remove(i)
                deleted++
            }
        }
		//stop when all are deleted
        if deleted == 0 {
            break
        }
    }
	
	//create array
    new_path := make([]deliver, path.Len())
    prev := s
	
	//for each path
    for p := 0; p < len(new_path); p++ {
        new_path[p] = prev
        prev = t.Next(prev, path)[p%2]
    }
    return new_path
}

//function implementing stepping stone method
func (t *data) SteppingStoneMethod(initfname string) {
    //declare variablesS
	Max := 0
    var shift []deliver = nil
    shifting := deliverzero
	
	//for each supply
    for p := 0; p < len(t.supply); p++ {
		//for each demand
        for q := 0; q < len(t.demand); q++ {
			//if element is Zero
            if t.set[p][q] == deliverzero {
				//test that element to minimize cost
				test := deliver{0, t.cost_data[p][q], p, q}
				
				//declare variables
				path := t.marginalpath(test)
				change := 0
				optimal := int(math.MaxInt32)
				copy := deliverzero
				flag := true
				
				//for each path
				for _, s := range path {
					
					//true
					if flag {
						change += s.Cost
					}else {
						change -= s.Cost
						//if cost is minimized
						if s.n < optimal {
							copy = s
							optimal = s.n
						}
					}
					flag = !flag
				}
				//if change is less than maximum
				if change < Max {
					shift = path
					shifting = copy
					Max = change
				}
			}
		}
    }
	//if shift is not Zero
    if shift != nil {
        q := shifting.n
        flag := true
		//for each shift
        for _, s := range shift {
			//true
            if flag {
                s.n += q
            }else {
                s.n -= q
            }
            if s.n == 0 {
                t.set[s.p][s.q] = deliverzero
            } else {
                t.set[s.p][s.q] = s
            }
            flag = !flag
        }
		//apply Stepping Store Method till no more optimal solution can be found
        t.SteppingStoneMethod(initfname)
    }
}

//function to print optimal solution
func (t *data) display() {
	//variable to calculate total cost
	totalCosts := 0
	
	fmt.Printf("\nOptimal solution:\n\n")
    //for each supply
    for p := 0; p < len(t.supply); p++ {
		//for each demand
        for q := 0; q < len(t.demand); q++ {
            s := t.set[p][q]
			//if not zero
            if s != deliverzero && s.p == p && s.q == q {
				//print element
                fmt.Printf(" %3d ", int(s.n))
                totalCosts += s.n * s.Cost
            }else {
                fmt.Printf("  -  ")
            }
        }
		//print line
        fmt.Println()
    }
	//display cost
    fmt.Printf("\nTotal Cost: %d\n\n", totalCosts)
} 

//function to write data to file
func (t *data) write_file(fname string,outfile string) {
	totalCosts := 0
	
	//open file to read format
	f, error := os.Open(fname)
	//hadle error
    error_handler(error)
	
	//close file
    defer f.Close()
	
	//store file in buffer
    fscan := bufio.NewScanner(f)
	
	//split into words
    fscan.Split(bufio.ScanWords)
	
	//number of words in line
    w := countWords(fname)
	
	//open or create file file to write
	wf, error := os.Create(outfile)
	//handle error
	error_handler(error)
	
	//write header line
	for p := 0; p < w-1; p++ {
		fscan.Scan()
		wf.WriteString(fscan.Text()+" ")
	}
	fscan.Scan()
	wf.WriteString(fscan.Text()+"\n")
	
	//for each supply
	for p := 0; p < len(t.supply); p++ {
		fscan.Scan()
		wf.WriteString(fscan.Text())
		//for each demand
        for q := 0; q < len(t.demand); q++ {
			//get element
            s := t.set[p][q]
			fscan.Scan()
			
			//if not zero
            if s != deliverzero && s.p == p && s.q == q {
				//write to file
                wf.WriteString(" "+strconv.Itoa(s.n))
                totalCosts += s.n * s.Cost
            } else {
                wf.WriteString(" - ")
            }
        }
		//write new line
		fscan.Scan()
        wf.WriteString(" " + fscan.Text() + "\n")
		
    }
	
		//write new line
		fscan.Scan()
        wf.WriteString(fscan.Text())

		//for each demand
        for q := 0; q < len(t.demand); q++ {
			//write demand
			fscan.Scan()
			wf.WriteString(" " + fscan.Text())
		}
	
	//write minimal cost
    wf.WriteString("\nMinimal cost: "+strconv.Itoa(totalCosts))
}
//END
