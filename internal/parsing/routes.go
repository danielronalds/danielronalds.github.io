package parsing

import "errors"

type Route struct {
	Src string
	Dest string
}

func ParseRoutes(dir string) ([]Route, error) {
	return nil, errors.New("not implemented")
}
