
package funfacts

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

var facts = FunFacts{
	Sun: []string{
		"Temperatur i Solens kjerne",
		"Temperatur på ytre lag av Solen",
	},
	Luna: []string{
		"Temperatur på Månens overflate om natten",
		"Temperatur på Månens overflate om dagen",
	},
	Terra: []string{
		"Høyeste temperatur målt på Jordens overflate",
		"Laveste temperatur målt på Jordens overflate",
		"Temperatur i Jordens indre kjerne.",
	},
}

func GetFunFacts(about string) []string {
	switch about {
	case "sun":
		return facts.Sun
	case "luna":
		return facts.Luna
	case "terra":
		return facts.Terra
	default:
		return []string{}
	}
}