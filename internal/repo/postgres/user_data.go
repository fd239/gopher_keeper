package postgres

import (
	"github.com/fd239/gopher_keeper/internal/models"
	"github.com/fd239/gopher_keeper/internal/repo"
	"github.com/fd239/gopher_keeper/pkg/crypt"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type userDataRepo struct {
	db    *sqlx.DB
	crypt *crypt.CipherCrypt
}

func NewUserDataRepo(db *sqlx.DB, crypt *crypt.CipherCrypt) repo.UsersDataRepo {
	return &userDataRepo{db: db, crypt: crypt}
}

const saveTextStmt = `INSERT INTO users_data (text, user_id, meta, type) VALUES ($1, $2, $3, $4) RETURNING id`

// SaveText implements text data save to postgres
func (r *userDataRepo) SaveText(dataText *models.DataText, userId uuid.UUID) (textId uuid.UUID, err error) {
	err = r.db.QueryRowx(
		saveTextStmt,
		dataText.Text,
		userId,
		dataText.Meta,
		dataText.Type,
	).Scan(&textId)

	if err != nil {
		return uuid.UUID{}, err
	}

	return
}

const selectTextStmt = `SELECT meta, text FROM users_data WHERE id=$1`

// SaveText implements text data save to postgres
func (r *userDataRepo) GetText(id uuid.UUID) (*models.DataText, error) {
	var dataText models.DataText
	err := r.db.Get(&dataText, selectTextStmt, id)
	return &dataText, err
}

const saveCardStmt = `INSERT INTO users_data (number, user_id, meta, type) VALUES ($1, $2, $3, $4) RETURNING id`

func (r *userDataRepo) SaveCard(dataCard *models.DataCard, userId uuid.UUID) (cardId uuid.UUID, err error) {
	cardNumber, err := r.crypt.Encrypt(dataCard.Number)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = r.db.QueryRowx(
		saveCardStmt,
		cardNumber,
		userId,
		dataCard.Meta,
		dataCard.Type,
	).Scan(&cardId)

	if err != nil {
		return uuid.UUID{}, err
	}

	return
}

const selectCardStmt = `SELECT meta, number FROM users_data WHERE id=$1`

func (r *userDataRepo) GetCard(id uuid.UUID) (*models.DataCard, error) {
	var dataCard models.DataCard
	err := r.db.Get(&dataCard, selectCardStmt, id)

	decryptedCardNumber, err := r.crypt.Decrypt(dataCard.Number)
	if err != nil {
		return nil, err
	}

	dataCard.Number = decryptedCardNumber

	return &dataCard, nil
}
