package main

import (
	"fmt"
	"log"
)

/**
ファクトリーパターンです。
S3、ローカルディスクを抽象化したパターンになります。
**/

type Storage interface {
	Save(filename string, data []byte) error
	GetURL(filename string) string
}

type S3Storage struct {
	Bucket string
	Region string
}

func (s *S3Storage) Save(f string, d []byte) error { /* S3へのアップロード処理 */ return nil }
func (s *S3Storage) GetURL(f string) string        { return "https://s3.amazonaws.com/" + f }

type LocalStorage struct {
	BasePath string
}

func (l *LocalStorage) Save(f string, d []byte) error { /* OSのファイル書き込み処理 */
	return nil
}
func (l *LocalStorage) GetURL(f string) string { return "/local/path/" + f }

type StorageType string

const (
	S3    StorageType = "s3"
	Local StorageType = "local"
)

func NewStorage(t StorageType, config map[string]string) (Storage, error) {
	switch t {
	case S3:
		return &S3Storage{
			Bucket: config["bucket"],
			Region: config["region"],
		}, nil
	case Local:
		return &LocalStorage{
			BasePath: config["base_path"],
		}, nil
	default:
		return nil, fmt.Errorf("storage type %s not supported", t)
	}
}

func runBusinessLogic(s Storage) {
	data := []byte("image-data-123")
	err := s.Save("avatar.png", data)
	if err != nil {
		log.Fatal(err)
	}
}

func Factory() {
	currentEnv := "local"

	configs := map[string]map[string]string{
		"s3":    {"bucket": "my-photo-bucket"},
		"local": {"base_path": "/var/www/uploads"},
	}

	storage, err := NewStorage(StorageType(currentEnv), configs[currentEnv])
	if err != nil {
		log.Fatalf("ストレージの初期化に失敗: %v", err)
	}

	runBusinessLogic(storage)
}
