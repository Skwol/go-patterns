package main

func main() {
	file1 := &file{name: "file_one"}
	file2 := &file{name: "file_two"}
	file3 := &file{name: "file_three"}

	folder1 := &folder{name: "folder_one"}
	folder1.add(file1)

	folder2 := &folder{name: "folder_two"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("test")
}
