package utils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AkashGit21/typeface-assignment/models"
	_ "github.com/go-sql-driver/mysql"
)

type PersistenceDBLayer struct {
	*sql.DB
}

type MetadataOps interface {
	SaveRecord(record models.Metadata) (int64, error)
	UpdateRecord(id int64, record models.Metadata) error
	FetchRecords() ([]models.Metadata, error)
	GetRecord(id int64) (*models.Metadata, error)
	DeleteRecord(id int64) error
}

func NewPersistenceDBLayer() (MetadataOps, error) {
	database := GetEnvValue("METADATA_DATABASE", "dbname")
	username := GetEnvValue("METADATA_USERNAME", "app-username")
	password := GetEnvValue("METADATA_PASSWORD", "app-password")
	host := GetEnvValue("METADATA_HOST", "dbhost")
	port := GetEnvValue("METADATA_PORT", "3306")

	// Create a DSN (Data Source Name) for the MySQL connection.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	// Open a connection to the MySQL database.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		ErrorLog("could not open database connection: ", err)
	}
	// defer db.Close()
	return &PersistenceDBLayer{
		db,
	}, nil
}

func (pdb *PersistenceDBLayer) SaveRecord(record models.Metadata) (int64, error) {
	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Insert a row into the "files" table.
	stmt, err := pdb.Prepare("INSERT INTO file_metadata (filename, size_in_bytes, s3_object_key, description, mime_type, status) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return int64(-1), err
	}
	defer stmt.Close()

	// Execute the SQL statement to insert the new row
	res, err := stmt.Exec(record.Filename, record.SizeInBytes, record.S3ObjectKey, record.Description, record.MimeType, record.Status)
	if err != nil {
		return int64(-1), err
	}

	return res.LastInsertId()
}

func (pdb *PersistenceDBLayer) UpdateRecord(id int64, record models.Metadata) error {
	return nil
}

func (pdb *PersistenceDBLayer) FetchRecords() ([]models.Metadata, error) {
	return nil, nil
}

func (pdb *PersistenceDBLayer) GetRecord(id int64) (*models.Metadata, error) {
	return nil, nil
}

func (pdb *PersistenceDBLayer) DeleteRecord(id int64) error {
	return nil
}
