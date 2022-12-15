package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
)

func main(){
    for { // Em um loop infinito
        reader := bufio.NewReader(os.Stdin) // Começamos chamando um leitor para ler os.Stdin
    
        fmt.Print("$ ") // Mostre o prompt atual
        input, _ := reader.ReadString('\n') // Leia o input do usuário, até ler uma quebra delinha
        
        if strings.HasSuffix(input, "&"){ // Verifique se a string contém algum & antes do final
            input = strings.TrimSuffix(input, "&") // Se houver remova-o
        }

        args := strings.Fields(strings.TrimSpace(input)) // Divida o input do usuário em chamadas e argumentos numa divisão por espaços

        if len(args) == 0 { // Se nada for digitado continue
            continue
        }
        
        switch(args[0]){ // Chamamos um switch para o primeiro valor de args
        case "exit","quit": // Caso o Input tenha 1 valor e seja exit ou quit, saia do shell
            return

        case "cd": // Caso seja cd mudamos de diretorio
            if len(args) == 2{
                err := os.Chdir(args[1])
                if err != nil{
                    log.Printf("Error changing directory: %v\n", err)
                }
            }else{
                log.Println("Use cd [diretorio]")
            }
        default: // Em default executamos um comando

            runInBackground := false // O estado inicial de runInBackground é falsa
            
            // Caso o ultimo do valor dos argumentos seja &, runInBackground passa a ser verdadeiro
            if len(args) > 1 && args[len(args)-1] == "&"{
                runInBackground = true
    
                args = args[:(len(args) - 1)]
            }

            cmd := exec.Command(args[0], args[1:]...) 
            cmd.Stdin = os.Stdin
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            
            // Se runInBackground for verdadeiro, rodar em background
            if runInBackground{
                err := cmd.Start()
                if err != nil{
                    log.Printf("Error running command in background: %v\n", err)
                }else{
                    log.Printf("Running command in the background: %s\n", strings.Join(args, " "))
                }
            }else{ // Caso contrário rodar em foreground
                err := cmd.Run()
                if err != nil{
                    log.Printf("Error running command: %v\n", err)
                }
            }
        }
    }
}
