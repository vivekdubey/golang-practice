package fileexplorer

import (
	"fmt"
	"log"
	"os"
	"testing"
	"vivekdubey/golang-practice/internal/util"

	"github.com/stretchr/testify/require"
)

func testSetup() (*os.File, string) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		log.Fatalf("Failed to create temp file in the test setup step")
	}
	d, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatalf("Failed to create temp directory in the test setup step")
	}

	return f, d

}

// test if a file could be created with a given name
func TestCreateFile(t *testing.T) {
	f, d := testSetup()
	fn := f.Name()
	t.Logf("File: %v, Dir: %v", fn, d)
	err := CreateFile(fn)
	fmt.Printf("Error: %v", err)
	require.Error(t, err)
	err = CreateFile(d)
	require.Error(t, err)
	rf := util.RandomFileName(10)
	err = CreateFile(rf)
	fmt.Println(err)
	require.NoError(t, err)
	// cleanup
	os.Remove(fn)
	os.Remove(rf)
	os.RemoveAll(d)
}
