package main

func main() {
	client := &client{}
	usHumidifier := &americanHumidifier{}

	client.plugInAmericanSocket(usHumidifier)

	euHumidifier := &europeanHumidifier{}
	euHumidifierAdapter := &euAdapter{
		humidifier: euHumidifier,
	}

	client.plugInAmericanSocket(euHumidifierAdapter)
}
