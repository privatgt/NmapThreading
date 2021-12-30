package main

import (
   "os"
   "os/exec"
   "fmt"
   "strconv"
   "strings"
   "sync"
)
var wg sync.WaitGroup
func largedivisor(div,thread int) int{
num:=1
   for i := 2; i != div; i++{
   	if (div%i==0) && (i<thread) && (i!=div) {
   	     if num==1{
   	         fmt.Println("lowest thread number is",i)
   	     }
   	     num=i
   	}
   } 
   fmt.Println("will run",num,"threads")
   return num
}
func scanner(port,allport,thread int){
	defer wg.Done()
	if port+(allport/thread) < allport{
	allport=port+(allport/thread)
	}
	 cmd := exec.Command("/bin/bash", "-c","nmap "+os.Args[1]+" -A -sT -p "+strconv.Itoa(port)+"-"+strconv.Itoa(allport))
	 output, _ := cmd.CombinedOutput()
	 if strings.Contains(string(output),"SERVICE"){
	 	fmt.Println(string(output))
	 }
}
func main(){
	if len(os.Args) > 2 {
		thread, _ := strconv.Atoi(os.Args[2])
		if 65535%thread>0{
		thread = largedivisor(65535,thread)
		}
		port:=0
		fmt.Println("Adding", thread,"workers")
		wg.Add(thread)
		for i := 0; i < thread; i++{
			go scanner(port,65535,thread)
			port+=65535/thread
		}
		wg.Wait()
		fmt.Println("All Workers Completed")
	}
}
