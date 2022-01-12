package main

import "fmt"

func main() {
	fileOne := &file{name: "file_one"}
	fileTwo := &file{name: "file_two"}
	fileThree := &file{name: "file_three"}

	folderOne := &folder{
		filesAndFolders: []node{fileOne},
		name:            "folder_one",
	}

	folderTwo := &folder{
		filesAndFolders: []node{folderOne, fileTwo, fileThree},
		name:            "folder_two",
	}
	fmt.Println("hierarchy for folder_two")
	folderTwo.print("  ")
	fmt.Println("------------------------------------------------------------------------")

	clonnedFolder := folderTwo.clone()
	fmt.Println("hierarchy for clonned folder_two")
	clonnedFolder.print("  ")
	fmt.Println("------------------------------------------------------------------------")

	folderTwo.addToName("_test")
	fmt.Println("hierarchy for renamed folder_two")
	folderTwo.print("  ")
	fmt.Println("------------------------------------------------------------------------")

	folderOne.addToName("_test_2")
	// cloneFolder := folder2.clone()
	fmt.Println("hierarchy for folder_two with renamded folder_one")
	folderTwo.print("  ")
}
