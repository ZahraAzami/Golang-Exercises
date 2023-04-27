package main

func SpaceAge(data float64, Planet string) float64 {
	var dataEarth float64

	switch {
	case Planet == "Mercury":
		dataEarth = data / 0.2408467
	case Planet == "Venus":
		dataEarth = data / 0.61519726
	case Planet == "Earth":
		dataEarth = data / 31557600
	case Planet == "Mars":
		dataEarth = data / 1.8808158
	case Planet == "Jupiter":
		dataEarth = data / 11.862615
	case Planet == "Saturn":
		dataEarth = data / 29.447498
	case Planet == "Uranus":
		dataEarth = data / 84.016846
	case Planet == "Neptune":
		dataEarth = data / 164.79132
	}
	return dataEarth
}
