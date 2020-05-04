package geo

type Geo interface {
	GetLocation(ipAddress string) (Location, error)
}
