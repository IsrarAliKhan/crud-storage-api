package storage

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	storage_model "crud-storage-api/shared/storage/models"
)

var FindIdRegExp = regexp.MustCompile(`^[0-9]+`)

type FileStorage struct {
	filename string
	orm      storage_model.StorageORM
}

func NewFileStorage(filename string, orm storage_model.StorageORM) *FileStorage {
	return &FileStorage{filename, orm}
}

func (fs FileStorage) Get(id uint64) (storage_model.StorageORM, error) {
	// open the file
	file, err := os.Open(fs.filename)
	if err != nil {
		log.Fatal(err)
		return fs.orm, err
	}
	defer file.Close()

	// read the file to check if the record exists
	var record string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if FindIdRegExp.FindString(scanner.Text()) == fmt.Sprintf("%v", id) {
			record = scanner.Text()
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return fs.orm, err
	}

	// check if the record exists
	if record == "" {
		return nil, fmt.Errorf("record not found")
	}

	// respond
	return fs.orm.Parse(record)
}

func (fs FileStorage) Save(m storage_model.StorageORM) (uint64, error) {
	// open the file
	file, err := os.OpenFile(fs.filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer file.Close()

	// read the file to check if the record already exists
	id := fmt.Sprintf("%v", m.GetId())
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if FindIdRegExp.FindString(scanner.Text()) == id {
			return 0, fmt.Errorf("record already exists")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return 0, err
	}

	// save the new record int the file
	_, err = file.WriteString(m.String())
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return m.GetId(), nil
}

func (fs FileStorage) Update(m storage_model.StorageORM) error {
	// open the file
	f, err := os.Open(fs.filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)
	id := fmt.Sprintf("%v", m.GetId())

	// read the file and write it to the buffer with the updated record
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if FindIdRegExp.FindString(scanner.Text()) != id {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
				return err
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
				return err
			}
		} else {
			_, err = buf.WriteString(m.String())
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	// write the buffer to the file
	err = os.WriteFile(fs.filename, buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (fs FileStorage) Delete(id uint64) error {
	// open the file
	f, err := os.Open(fs.filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)
	dId := fmt.Sprintf("%v", id)

	// read the file and write it to the buffer except for the record to be removed
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if FindIdRegExp.FindString(scanner.Text()) != dId {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
				return err
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	// write the buffer to the file
	err = os.WriteFile(fs.filename, buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
