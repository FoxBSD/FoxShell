package main

import(
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main(){
    reader := bufio.NewReader(os.Stdin) // Começamos chamando um leitor para ler os.Stdin
    
    for { // Em um loop infinito

        fmt.Print("$ ") // Mostre o prompt atual
        input, _ := reader.ReadString('\n') // Leia o input do usuário, até ler uma quebra delinha

        runInBackground := false // O estado inicial de runInBackground é falsa

        args := strings.Fields(strings.TrimSpace(input)) // Divida o input do usuário em chamadas e argumentos numa divisão por espaços

        if len(args) == 1 && (args[0] == "exit" || args[0] == "quit"){ // Caso o Input tenha 1 valor e seja exit ou quit, saia do shell
            os.Exit(0)
        }
        
        // Caso o ultimo do valor dos argumentos seja &, runInBackground passa a ser verdadeiro

        if len(args) > 1 && args[len(args)-1] == "&"{
            runInBackground = true

            args = args[:(len(args) - 1)]
        }

        switch(args[0]){ // Chamamos um switch para o primeiro valor de args
        case "cd": // Caso seja cd mudamos de diretorio
            if len(args) == 2{
                os.Chdir(args[1])
            }else{
                fmt.Println("Use cd [diretorio]")
            }
        default: // Em default executamos um comando
            cmd := exec.Command(args[0], args[1:]...) 
            cmd.Stdin = os.Stdin
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            
            // Se runInBackground for verdadeiro, rodar em background
            if runInBackground{
                err := cmd.Start()
                if err != nil{
                    fmt.Println(err)
                }else{
                    fmt.Println("Running command in the background: ", strings.Join(args, " "))
                }
            }else{ // Caso contrário rodar em foreground
                err := cmd.Run()
                if err != nil{
                    fmt.Println(err)
                }
            }
        }
    }
}
