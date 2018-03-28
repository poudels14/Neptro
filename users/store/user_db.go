package store

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/poudels14/Neptro/utils"
)

const (
	mysqlDriver = "mysql"
)

type UserDb struct {
	db *sql.DB
}

func InitializeUserStore() (UserStore, error) {
	utils.LoadAllKeys()

	dbUser := os.Getenv("TEST_DB_USER")
	dbPass := os.Getenv("TEST_DB_PASS")
	dbName := os.Getenv("TEST_DB_NAME")

	return InitializeUserDb(dbUser, dbPass, dbName)
}

func InitializeUserDb(dbUser, dbPwd, dbName string) (UserStore, error) {
	dataSource := fmt.Sprintf("%s:%s@/%s?parseTime=true", dbUser, dbPwd, dbName)
	db, err := sql.Open(mysqlDriver, dataSource)
	if err != nil {
		return nil, err
	}

	return &UserDb{db}, nil
}

// Creates a new user row in the twitter_users table
func (m *UserDb) Create(user User) (int64, error) {
	stmt, err := m.db.Prepare(`
      INSERT INTO users (first_name, last_name, email, password_hash,
				phone, country_code, created_at)
      VALUES(?, ?, ?, ?, ?, ?, ?);
    `)
	if err != nil {
		return 0, err
	}

	currentTime := time.Now()
	res, err := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		user.Phone,
		user.CountryCode,
		currentTime.Unix())

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Updates the given user
func (m *UserDb) Update(user User) error {
	stmt, err := m.db.Prepare(`
      UPDATE users
			SET first_name = ?, last_name = ?, email = ?, phone = ?,
        country_code = ?, last_signed_in = ?, last_signed_in_ip = ?
			WHERE id = ?;
    `)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.CountryCode,
		user.LastSignedIn,
		user.LastSignedInIP,
		user.ID)

	if err != nil {
		return err
	}

	return nil
}

// Deletes the user corresponding to the given id
func (m *UserDb) Deactivate(id int64) error {
	stmt, err := m.db.Prepare(`
      UPDATE users
			SET archived_at = ?
			WHERE id = ?;
    `)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	_, err = stmt.Exec(currentTime.Unix(), id)
	if err != nil {
		return err
	}

	return nil
}

// Retrieves the user corresponding to the given id
func (m *UserDb) Get(id int64) (*User, error) {
	row := m.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	user := User{}
	var createdAt int64
	var lastSignedIn, archivedAt sql.NullInt64
	var lastSignedInIP sql.NullString

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone,
		&user.CountryCode, &user.PasswordHash, &createdAt, &lastSignedIn, &lastSignedInIP, &archivedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	user.CreatedAt = time.Unix(createdAt, 0)

	if lastSignedInIP.Valid {
		user.LastSignedInIP = &lastSignedInIP.String
	}

	if lastSignedIn.Valid {
		lastSignedInTime := time.Unix(lastSignedIn.Int64, 0)
		user.LastSignedIn = &lastSignedInTime
	}

	if archivedAt.Valid {
		archivedAtTime := time.Unix(archivedAt.Int64, 0)
		user.ArchivedAt = &archivedAtTime
	}

	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

// Retrieves the total number for rows in the twitter_users table
func (m *UserDb) Count() (int64, error) {
	row := m.db.QueryRow("SELECT COUNT(*) FROM users")

	var count int64
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}
