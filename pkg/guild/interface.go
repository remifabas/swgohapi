package guild

type Service interface {
	GetPlayers() ([]Player, error)
}

func NewService(path string) Service {
	return &DefaultService{filePath: path}
}

type Player struct {
	AllyCode     string `yaml:"ally_code"`
	GameName     string `yaml:"game_name"`
	ExportedName string `yaml:"exported_name"`
}
