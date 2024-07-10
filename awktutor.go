package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Display the introduction
	displayLesson("lessons/introduction.txt")

	// Wait for user to press Enter to continue
	fmt.Println("\nPress Enter to start the first lesson...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Display the first lesson
	displayLesson("lessons/lesson1.txt")

	// Get user input and execute the awk command
	getUserInputAndExecute()
}

func displayLesson(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
	}
}

func getUserInputAndExecute() {
	fmt.Println("\nNow it's your turn! Write your awk command below:")
	fmt.Print("$ ")
	inputReader := bufio.NewReader(os.Stdin)
	command, _ := inputReader.ReadString('\n')
	command = strings.TrimSpace(command)

	fmt.Println("\nYour Output:")
	runAwkCommand(command)
}

func runAwkCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
	}
}
