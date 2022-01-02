package main

import "fmt"

func main() {
	database := map[string]string{
		"brasil":         "Feliz Natal!",
		"alemanha":       "Frohliche Weihnachten!",
		"austria":        "Frohe Weihnacht!",
		"coreia":         "Chuk Sung Tan!",
		"espanha":        "Feliz Navidad!",
		"grecia":         "Kala Christougena!",
		"estados-unidos": "Merry Christmas!",
		"inglaterra":     "Merry Christmas!",
		"australia":      "Merry Christmas!",
		"portugal":       "Feliz Natal!",
		"suecia":         "God Jul!",
		"turquia":        "Mutlu Noeller",
		"argentina":      "Feliz Navidad!",
		"chile":          "Feliz Navidad!",
		"mexico":         "Feliz Navidad!",
		"antardida":      "Merry Christmas!",
		"canada":         "Merry Christmas!",
		"irlanda":        "Nollaig Shona Dhuit!",
		"belgica":        "Zalig Kerstfeest!",
		"italia":         "Buon Natale!",
		"libia":          "Buon Natale!",
		"siria":          "Milad Mubarak!",
		"marrocos":       "Milad Mubarak!",
		"japao":          "Merii Kurisumasu!",
	}

	var country string

	for {
		_, err := fmt.Scanf("%s", &country)
		if err != nil {
			break
		}

		value, ok := database[country]

		if ok {
			fmt.Println(value)
		} else {
			fmt.Println("--- NOT FOUND ---")
		}

	}

}
