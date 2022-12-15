package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
)

//Implement instructions to turn .fsh_history in a circular log.
const maxHistorySize int = 10 // Define max_size of HistoryBuffer

var history [maxHistorySize]string // The History Buffer with fixed size

var historyCounter int = 0 // The counter of total number add to buffer

var historyPointer int = 0 // The pointer to the current positition in buffer


func main(){
    for { // In an infinite loop
        reader := bufio.NewReader(os.Stdin) // We start by calling a reader to read os.Stdin
    
        fmt.Print("$ ") // Show the current prompt
        input, _ := reader.ReadString('\n') // Read the user's input, until a newline is read
        
        // adding command to history buffer
        history[historyPointer] = input
        historyCounter++
        historyPointer = (historyPointer + 1) % maxHistorySize

        if strings.HasSuffix(input, "&"){ // Check if the string contains any & before the end
            input = strings.TrimSuffix(input, "&") // If there is, remove it
        }

        args := strings.Fields(strings.TrimSpace(input)) // Split the user's input into calls and arguments on a space-delimited basis

        if len(args) == 0 { // If nothing is typed, continue
            continue
        }
        
        switch(args[0]){ // We call a switch on the first value of args
        case "exit","quit": // If the Input has 1 value and it's exit or quit, exit the shell
            return

        case "cd": // If it's cd, change directory
            if len(args) == 2{
                err := os.Chdir(args[1])
                if err != nil{
                    log.Printf("Error changing directory: %v\n", err)
                }
            }else{
                log.Println("Use cd [directory]")
            }
        default: // In default, run a command

            runInBackground := false // The initial state of runInBackground is false
            
            // If the last value of the arguments is &, runInBackground becomes true
            if len(args) > 1 && args[len(args)-1] == "&"{
                runInBackground = true
    
                args = args[:(len(args) - 1)]
            }

            cmd := exec.Command(args[0], args[1:]...) 
            cmd.Stdin = os.Stdin
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            
            // If runInBackground is true, run in the background
            if runInBackground{
                err := cmd.Start()
                if err != nil{
                    log.Printf("Error running command in background: %v\n", err)
                }else{
                    log.Printf("Running command in the background: %s\n", strings.Join(args, " "))
                }
            }else{ // Otherwise, run in the foreground
                err := cmd.Run()
                if err != nil{
                    log.Printf("Error running command: %v\n", err)
                }
            }
        }
    }
}
