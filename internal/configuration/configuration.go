package configuration

import (
	"os"
)

type Configuration struct {
	RootPath string
}

func Read() *Configuration {
	rootPath, rootPathSet := os.LookupEnv("ROOT_PATH")
	if !rootPathSet {
		rootPath = "./data"
	}

	return &Configuration{RootPath: rootPath}
}

func (c *Configuration) Verify() error {
	//if !internal.dirExists(c.RootPath) {
	//	return errors.New("root path does not exist")
	//}
	return nil
}
