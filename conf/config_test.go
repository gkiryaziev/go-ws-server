package conf

import "testing"

func TestLoad(t *testing.T) {
	// read and parse yaml file
	config, err := NewConfig("../config.yaml").Load()
	if err != nil {
		t.Fatal(err)
	}

	// check parameters
	switch {
	case config.Debug != true && config.Debug != false:
		t.Error("Error, Debug =", config.Debug)
	case config.Server.Host == "":
		t.Error("Error, Server is empty.")
	case config.Server.Port == "":
		t.Error("Error, Port is empty.")
	}
}
