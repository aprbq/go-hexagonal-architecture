package repository

import "github.com/jmoiron/sqlx"

// ไม่ expose
type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(acc Account) (*Account, error) {
	query := `INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING account_id`

	err := r.db.QueryRow(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	).Scan(&acc.AccountID)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id=$1"
	accounts := []Account{}
	err := r.db.Select(&accounts, query, customerID)

	if err != nil {
		return nil, err
	}
	return accounts, nil
}
