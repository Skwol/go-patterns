package main

func main() {
	comp := &computer{}

	onCommand := &onCommand{
		device: comp,
	}

	offCommand := &offCommand{
		device: comp,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
