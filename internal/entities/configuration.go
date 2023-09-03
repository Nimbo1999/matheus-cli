package entities

type Configuration struct {
	Id       string
	FilePath string
}

func NewConfiguration(id, filePath string) *Configuration {
	return &Configuration{
		Id:       id,
		FilePath: filePath,
	}
}
