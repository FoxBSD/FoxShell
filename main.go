package main

import(
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    for {
    fmt.Print("$ ")
    input, _ := reader.ReadString('\n')

    args := strings.Split(strings.TrimSpace(input)," ")

    if len(args) == 1 && (args[0] == "exit" || args[0] == "quit"){
        os.Exit(0)
    }

    cmd := exec.Command(args[0], args[1:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}
}
