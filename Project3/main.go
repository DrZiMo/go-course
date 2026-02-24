package main

import (
	"bufio"
	"fmt"
	"note-app/note"
	"os"
	"strings"
)

func main() {
	for {
		choice := showMenu()

		decision(choice)
	}
}

func showMenu() int {
	fmt.Println("\nNOTES APP")
	fmt.Println("1. show notes")
	fmt.Println("2. new note")

	fmt.Print("\nchoose: ")
	var value int
	fmt.Scanln(&value)
	fmt.Print("\n")
	return value
}

func decision(choice int) {
	switch choice {
	case 1:
		showExistingNotes()
	case 2:
		saveNote()
	}
}

func saveNote() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving note failed!")
		return
	}

	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {
	title := getNoteInput("Note title: ")
	content := getNoteInput("Note content: ")

	return title, content
}

func getNoteInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func extractExistingNotes() ([]string, error) {
	files, err := os.ReadDir("notes")
	var fileNames []string

	if err != nil {
		fmt.Print("Error: ", err)
		return nil, err
	}

	for _, file := range files {
		fileName := file.Name()

		fileName = strings.ReplaceAll(fileName, "_", " ")
		fileName = strings.TrimSuffix(fileName, ".json")
		fileNames = append(fileNames, fileName)
	}

	return fileNames, nil
}

func showExistingNotes() {
	fileNames, err := extractExistingNotes()

	if err != nil {
		fmt.Print("Failed to extract notes!")
		return
	}

	fmt.Println("Existing notes: ")

	if len(fileNames) > 0 {
		for index, file := range fileNames {
			fmt.Printf("%v. %v", index+1, file)
			fmt.Println()
		}
	} else {
		fmt.Println("No existing notes!")
	}
}
