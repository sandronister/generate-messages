package service

type message struct {
	Index int    `json:"index"`
	Item  string `json:"item"`
}

var countries = []string{
	"Brasil",
	"Argentina",
	"Chile",
	"Uruguai",
	"Paraguai",
	"Bolívia",
	"Peru",
	"Colômbia",
	"Venezuela",
	"Equador",
	"Guiana",
	"Suriname",
	"Guiana Francesa",
	"Estados Unidos",
	"Canadá",
	"México",
	"Cuba",
	"República Dominicana",
	"Costa Rica",
	"Panamá",
	"China",
	"Japão",
	"Coreia do Sul",
	"Índia",
	"Rússia",
	"Alemanha",
	"França",
	"Itália",
	"Espanha",
	"Portugal",
}

func GenerateCountries() []message {
	var list []message
	for i, item := range countries {
		msm := message{
			Index: i,
			Item:  item,
		}
		list = append(list, msm)
	}
	return list
}
