package main

import "fmt"

type node interface {
	print(string)
	clone() node
	getName() string
	addToName(string)
}

type file struct {
	name string
}

func (f file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f file) clone() node {
	return &file{name: f.name + "_clone"}
}

func (f file) getName() string {
	return f.name
}

func (f *file) addToName(str string) {
	f.name = fmt.Sprintf("%s%s", f.name, str)
}

type folder struct {
	filesAndFolders []node
	name            string
}

func (f folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.filesAndFolders {
		i.print(indentation + indentation)
	}
}

func (f folder) clone() node {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []node
	for _, i := range f.filesAndFolders {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.filesAndFolders = tempChildren
	return cloneFolder
}

func (f folder) getName() string {
	return f.name
}

func (f *folder) addToName(str string) {
	f.name = fmt.Sprintf("%s%s", f.name, str)
	for i := range f.filesAndFolders {
		f.filesAndFolders[i].addToName(str)
	}
}
