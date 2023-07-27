package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Entry struct {
	Name string
	IsDir bool
}

func listDir(path string) ([]Entry, error) {
    entries := make([]Entry, 0)
    files, err := ioutil.ReadDir(path)
    if err != nil {
        return nil, err
    }

    for _, file := range files {
        entry := Entry{
            Name: file.Name(),
            IsDir: file.IsDir(),
        }
        entries = append(entries, entry)
    }

    return entries, nil
}

func printEntries(entries []Entry) {
    for _, entry := range entries {
        if entry.IsDir {
            fmt.Printf("[%s]\n", entry.Name)
        } else {
            fmt.Println(entry.Name)
        }
    }
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readFile(file string) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Contenido de: ", file)
	fmt.Println(string(contents))
}

func currentFolder() string {
	currentF, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	CF := filepath.Base(currentF)
	return CF
	//fmt.Println("Nombre de la carpeta actual: " + CF)
}

func currentFiles() {
	//Una carpeta atras: dir := /Golang
	dir := "."
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	currentFolderName := currentFolder()
	fmt.Printf("Archivos en la carpeta '%s':\n", currentFolderName)

	for _, file := range files {
		fmt.Println(file.Name())
	}

}

func main() {

	//reader := bufio.NewReader((os.Stdin))
	
	/*Read a single current files
	for {
		fmt.Println("¿Qué archivo quiere leer? (Introduzca el nombre)")
		currentFiles()
		fmt.Println("\n")

		file, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		file = strings.TrimSpace(file)

		readFile(file)

		fmt.Println("Deseas leer otro archivo? (si/no)")
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		answer = strings.ToLower(strings.TrimSpace(answer))
		if answer != "si" {
			fmt.Println("Cerrando programa ...")
			break
		}

	}
	*/
	
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Directorio actual: %s\n", currentPath)

	for {
		entries, err := listDir(currentPath)
		if err != nil {
			log.Fatal(err)
		}

		printEntries(entries)

		fmt.Println("\n1. Abrir archivo")
		fmt.Println("2. Cambiar directorio")
		fmt.Println("3. Navegar una carpeta hacia atrás")
		fmt.Println("4. Salir")

		var choice int
		fmt.Print("Ingrese una opción: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var filename string
			fmt.Print("Ingrese el nombre del archivo a abrir: ")
			fmt.Scan(&filename)
			readFile(filepath.Join(currentPath, filename))
			fmt.Print("Presione Enter para continuar...")
			fmt.Scanln()
			clearScreen()
		case 2:
			var foldername string
			fmt.Print("Ingrese el nombre de la carpeta a explorar: ")
			fmt.Scan(&foldername)
			newpath := filepath.Join(currentPath, foldername)
			_, err := os.Stat(newpath)
			if err == nil {
				currentPath = newpath
			} else {
				fmt.Println("Carpeta no encontrada.")
			}
			clearScreen()
		case 3:
			currentPath = filepath.Dir(currentPath)
			clearScreen()
		case 4:
			fmt.Println("Saliendo del explorador.")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}



}