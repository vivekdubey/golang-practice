package fileexplorer

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Run before test")
	exitVal := m.Run()
	log.Println("Run after test")
	os.Exit(exitVal)
}
