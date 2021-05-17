package config

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// dbConfig represents the config which is used when DB is initialized
type dbConfig struct {
	ID       int    `json:"id"`
	Checksum string `json:"checksum"`
}

// DB represents the application server database (json file)
type DB struct {
	DbPath       string `envconfig:"DB_PATH" required:"true"`
	DbConfigPath string `envconfig:"DB_CONFIG_PATH" required:"true"`
	cfg          dbConfig
	db           []byte
}

// NewDatabase load configuration from env vars
func NewDatabase() (DB, error) {
	d := DB{}
	return d, envconfig.Process("", &d)
}

// Start starts and initializes the file database
func (d *DB) Start() error {
	bs, err := d.read(d.DbConfigPath)
	if err != nil {
		return fmt.Errorf("could not read db config contents %s", err)
	}
	var cfg dbConfig
	if len(bs) == 0 {
		bs = []byte("{}")
	}
	err = json.Unmarshal(bs, &cfg)
	if err != nil {
		return fmt.Errorf("could not unmarshall db config %s", err)
	}

	bs, err = d.read(d.DbPath)
	if err != nil {
		return fmt.Errorf("could not read db contents %s", err)
	}
	d.db = bs
	if d.cfg.Checksum == "" {
		checksum, err := genChecksum(bytes.NewReader(bs))
		if err != nil {
			return err
		}
		cfg.Checksum = checksum
	}
	d.cfg = cfg
	return nil
}

func genChecksum(r io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, r); err != nil {
		return "", fmt.Errorf("could not copy db contents %s", err)
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

func (d *DB) read(path string) ([]byte, error) {
	dbFile, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if errors.Is(err, os.ErrNotExist) {
		dbFile, err = os.Create(path)
	}
	if err != nil {
		return nil, fmt.Errorf("could not open or create db file %s", err.Error())
	}
	return ioutil.ReadAll(dbFile)
}
