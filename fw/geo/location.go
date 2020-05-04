package geo

type Continent struct {
	Code string
	Name string
}

type Country struct {
	Code string
	Name string
}

type Region struct {
	Code string
	Name string
}

type Currency struct {
	Code string
	Name string
}

type Language struct {
	Code string
	Name string
}

type Location struct {
	Continent       Continent
	Country         Country
	Region          Region
	City            string
	Currency        Currency
	Languages       []Language
	IsEuropeanUnion bool
}
