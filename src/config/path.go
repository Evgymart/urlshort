package config

type Path struct {
	AppRoot string
}

var (
	path *Path = nil
)

func InitPath(initPath *Path) {
	if path == nil {
		path = initPath
	}
}

func GetPath() *Path {
	return path
}
