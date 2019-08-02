package database

import (
	"fmt"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/network"
)

func (db *DB) CreateTableNetwork() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Network\" (\"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"RPCUser\" STRING NOT NULL, \"ChainName\" STRING NOT NULL, \"Public\" BOOL NOT NULL, \"Shared\" BOOL NOT NULL, \"DefaultRpcPort\" INT NOT NULL, \"AddressPubkeyhashVersion\" STRING NOT NULL, \"AddressChecksumValue\" STRING NOT NULL, \"BurnAddress\" STRING NOT NULL, \"Model\" INT NOT NULL, \"WizardKey\" STRING NOT NULL, \"RPCPassword\" STRING NOT NULL, \"PrivateKeyVersion\" STRING NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertNetwork(AddressChecksumValue string, AddressPubkeyhashVersion string, BurnAddress string, ChainName string, DefaultRpcPort int, Model int, PrivateKeyVersion string, Public bool, RPCPassword string, RPCUser string, Shared bool, WizardKey string) (*models.Network, error) {
	row := &models.Network{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Network\" (\"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"WizardKey\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING AddressChecksumValue, AddressPubkeyhashVersion, BurnAddress, ChainName, DefaultRpcPort, Model, PrivateKeyVersion, Public, RPCPassword, RPCUser, Shared, WizardKey;", AddressChecksumValue, AddressPubkeyhashVersion, BurnAddress, ChainName, DefaultRpcPort, Model, PrivateKeyVersion, Public, RPCPassword, RPCUser, Shared, WizardKey).Scan(&row.AddressChecksumValue, &row.AddressPubkeyhashVersion, &row.BurnAddress, &row.ChainName, &row.Created, &row.DefaultRpcPort, &row.Model, &row.PrivateKeyVersion, &row.Public, &row.RPCPassword, &row.RPCUser, &row.Shared, &row.UID, &row.WizardKey); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteNetwork(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Network\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutNetwork(model *models.Network) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Network\" (\"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"WizardKey\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\";", model.AddressChecksumValue, model.AddressPubkeyhashVersion, model.BurnAddress, model.ChainName, model.DefaultRpcPort, model.Model, model.PrivateKeyVersion, model.Public, model.RPCPassword, model.RPCUser, model.Shared, model.WizardKey).Scan(&model.AddressChecksumValue, &model.AddressPubkeyhashVersion, &model.BurnAddress, &model.ChainName, &model.Created, &model.DefaultRpcPort, &model.Model, &model.PrivateKeyVersion, &model.Public, &model.RPCPassword, &model.RPCUser, &model.Shared, &model.UID, &model.WizardKey); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToNetwork(rows *sql.Rows) ([]*models.Network, error) {
	defer rows.Close()
	results := []*models.Network{}
	for rows.Next() {
		result := &models.Network{}
		err := rows.Scan(&result.AddressChecksumValue, &result.AddressPubkeyhashVersion, &result.BurnAddress, &result.ChainName, &result.Created, &result.DefaultRpcPort, &result.Model, &result.PrivateKeyVersion, &result.Public, &result.RPCPassword, &result.RPCUser, &result.Shared, &result.UID, &result.WizardKey)
		if err != nil {
			return nil, err
		}
		results = append(
			results,
			result,
		)
	}
	return results, nil
}

func (db *DB) CountNetwork() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Network\";")
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) QueryNetwork() ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByUID(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByUID(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByCreated(value int64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByCreated(value int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreated(value int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndUID(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndUID(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndRPCUser(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndRPCUser(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndChainName(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndChainName(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndPublic(value1 int64, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndPublic(value1 int64, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndShared(value1 int64, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndShared(value1 int64, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndDefaultRpcPort(value1 int64, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndDefaultRpcPort(value1 int64, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndAddressPubkeyhashVersion(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndAddressPubkeyhashVersion(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndAddressChecksumValue(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndAddressChecksumValue(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndBurnAddress(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndBurnAddress(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndModel(value1 int64, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndModel(value1 int64, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndWizardKey(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndWizardKey(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndRPCPassword(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndRPCPassword(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByCreatedAndPrivateKeyVersion(value1 int64, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByCreatedAndPrivateKeyVersion(value1 int64, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Created\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByRPCUser(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"RPCUser\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByRPCUser(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUser(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndUID(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndUID(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndCreated(value1 string, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndCreated(value1 string, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndChainName(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndChainName(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndPublic(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndPublic(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndShared(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndShared(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndDefaultRpcPort(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndDefaultRpcPort(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndAddressPubkeyhashVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndAddressPubkeyhashVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndAddressChecksumValue(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndAddressChecksumValue(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndBurnAddress(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndBurnAddress(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndModel(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndModel(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndWizardKey(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndWizardKey(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndRPCPassword(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndRPCPassword(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCUserAndPrivateKeyVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCUserAndPrivateKeyVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCUser\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByChainName(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"ChainName\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByChainName(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainName(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndUID(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndUID(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndCreated(value1 string, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndCreated(value1 string, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndRPCUser(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndRPCUser(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndPublic(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndPublic(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndShared(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndShared(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndDefaultRpcPort(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndDefaultRpcPort(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndAddressPubkeyhashVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndAddressPubkeyhashVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndAddressChecksumValue(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndAddressChecksumValue(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndBurnAddress(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndBurnAddress(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndModel(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndModel(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndWizardKey(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndWizardKey(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndRPCPassword(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndRPCPassword(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByChainNameAndPrivateKeyVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByChainNameAndPrivateKeyVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"ChainName\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByPublic(value bool) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"Public\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByPublic(value bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublic(value bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndUID(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndUID(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndCreated(value1 bool, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndCreated(value1 bool, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndRPCUser(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndRPCUser(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndChainName(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndChainName(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndShared(value1 bool, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndShared(value1 bool, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndDefaultRpcPort(value1 bool, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndDefaultRpcPort(value1 bool, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndAddressPubkeyhashVersion(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndAddressPubkeyhashVersion(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndAddressChecksumValue(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndAddressChecksumValue(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndBurnAddress(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndBurnAddress(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndModel(value1 bool, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndModel(value1 bool, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndWizardKey(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndWizardKey(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndRPCPassword(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndRPCPassword(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPublicAndPrivateKeyVersion(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPublicAndPrivateKeyVersion(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Public\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByShared(value bool) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"Shared\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByShared(value bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByShared(value bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndUID(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndUID(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndCreated(value1 bool, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndCreated(value1 bool, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndRPCUser(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndRPCUser(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndChainName(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndChainName(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndPublic(value1 bool, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndPublic(value1 bool, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndDefaultRpcPort(value1 bool, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndDefaultRpcPort(value1 bool, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndAddressPubkeyhashVersion(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndAddressPubkeyhashVersion(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndAddressChecksumValue(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndAddressChecksumValue(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndBurnAddress(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndBurnAddress(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndModel(value1 bool, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndModel(value1 bool, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndWizardKey(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndWizardKey(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndRPCPassword(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndRPCPassword(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkBySharedAndPrivateKeyVersion(value1 bool, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkBySharedAndPrivateKeyVersion(value1 bool, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Shared\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByDefaultRpcPort(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"DefaultRpcPort\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByDefaultRpcPort(value int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPort(value int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndUID(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndUID(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndCreated(value1 int, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndCreated(value1 int, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndRPCUser(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndRPCUser(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndChainName(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndChainName(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndPublic(value1 int, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndPublic(value1 int, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndShared(value1 int, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndShared(value1 int, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndAddressPubkeyhashVersion(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndAddressPubkeyhashVersion(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndAddressChecksumValue(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndAddressChecksumValue(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndBurnAddress(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndBurnAddress(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndModel(value1 int, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndModel(value1 int, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndWizardKey(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndWizardKey(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndRPCPassword(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndRPCPassword(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByDefaultRpcPortAndPrivateKeyVersion(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByDefaultRpcPortAndPrivateKeyVersion(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"DefaultRpcPort\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByAddressPubkeyhashVersion(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"AddressPubkeyhashVersion\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByAddressPubkeyhashVersion(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersion(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndUID(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndUID(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndCreated(value1 string, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndCreated(value1 string, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndRPCUser(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndRPCUser(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndChainName(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndChainName(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndPublic(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndPublic(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndShared(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndShared(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndDefaultRpcPort(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndDefaultRpcPort(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndAddressChecksumValue(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndAddressChecksumValue(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndBurnAddress(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndBurnAddress(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndModel(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndModel(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndWizardKey(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndWizardKey(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndRPCPassword(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndRPCPassword(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressPubkeyhashVersionAndPrivateKeyVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressPubkeyhashVersionAndPrivateKeyVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressPubkeyhashVersion\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByAddressChecksumValue(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"AddressChecksumValue\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByAddressChecksumValue(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValue(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndUID(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndUID(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndCreated(value1 string, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndCreated(value1 string, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndRPCUser(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndRPCUser(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndChainName(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndChainName(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndPublic(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndPublic(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndShared(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndShared(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndDefaultRpcPort(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndDefaultRpcPort(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndAddressPubkeyhashVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndAddressPubkeyhashVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndBurnAddress(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndBurnAddress(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndModel(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndModel(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndWizardKey(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndWizardKey(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndRPCPassword(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndRPCPassword(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByAddressChecksumValueAndPrivateKeyVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByAddressChecksumValueAndPrivateKeyVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"AddressChecksumValue\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByBurnAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"BurnAddress\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByBurnAddress(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddress(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndUID(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndUID(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndCreated(value1 string, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndCreated(value1 string, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndRPCUser(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndRPCUser(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndChainName(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndChainName(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndPublic(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndPublic(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndShared(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndShared(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndDefaultRpcPort(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndDefaultRpcPort(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndAddressPubkeyhashVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndAddressPubkeyhashVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndAddressChecksumValue(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndAddressChecksumValue(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndModel(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndModel(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndWizardKey(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndWizardKey(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndRPCPassword(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndRPCPassword(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByBurnAddressAndPrivateKeyVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByBurnAddressAndPrivateKeyVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"BurnAddress\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByModel(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"Model\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByModel(value int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModel(value int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndUID(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndUID(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndCreated(value1 int, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndCreated(value1 int, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndRPCUser(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndRPCUser(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndChainName(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndChainName(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndPublic(value1 int, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndPublic(value1 int, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndShared(value1 int, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndShared(value1 int, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndDefaultRpcPort(value1 int, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndDefaultRpcPort(value1 int, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndAddressPubkeyhashVersion(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndAddressPubkeyhashVersion(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndAddressChecksumValue(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndAddressChecksumValue(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndBurnAddress(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndBurnAddress(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndWizardKey(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndWizardKey(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndRPCPassword(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndRPCPassword(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByModelAndPrivateKeyVersion(value1 int, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByModelAndPrivateKeyVersion(value1 int, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"Model\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByWizardKey(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"WizardKey\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByWizardKey(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKey(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndUID(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndUID(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndCreated(value1 string, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndCreated(value1 string, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndRPCUser(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndRPCUser(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndChainName(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndChainName(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndPublic(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndPublic(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndShared(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndShared(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndDefaultRpcPort(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndDefaultRpcPort(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndAddressPubkeyhashVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndAddressPubkeyhashVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndAddressChecksumValue(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndAddressChecksumValue(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndBurnAddress(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndBurnAddress(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndModel(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndModel(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndRPCPassword(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndRPCPassword(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByWizardKeyAndPrivateKeyVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByWizardKeyAndPrivateKeyVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"WizardKey\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByRPCPassword(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"RPCPassword\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByRPCPassword(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPassword(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndUID(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndUID(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndCreated(value1 string, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndCreated(value1 string, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndRPCUser(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndRPCUser(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndChainName(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndChainName(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndPublic(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndPublic(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndShared(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndShared(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndDefaultRpcPort(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndDefaultRpcPort(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndAddressPubkeyhashVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndAddressPubkeyhashVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndAddressChecksumValue(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndAddressChecksumValue(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndBurnAddress(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndBurnAddress(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndModel(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndModel(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndWizardKey(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndWizardKey(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByRPCPasswordAndPrivateKeyVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"PrivateKeyVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByRPCPasswordAndPrivateKeyVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"RPCPassword\" = $1 AND \"PrivateKeyVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) CountNetworkByPrivateKeyVersion(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Network\" WHERE \"PrivateKeyVersion\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetNetworkByPrivateKeyVersion(value string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersion(value string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndUID(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndUID(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndCreated(value1 string, value2 int64) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndCreated(value1 string, value2 int64) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndRPCUser(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndRPCUser(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndChainName(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"ChainName\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndChainName(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"ChainName\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndPublic(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndPublic(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndShared(value1 string, value2 bool) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"Shared\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndShared(value1 string, value2 bool) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"Shared\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndDefaultRpcPort(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"DefaultRpcPort\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndDefaultRpcPort(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"DefaultRpcPort\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndAddressPubkeyhashVersion(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"AddressPubkeyhashVersion\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndAddressPubkeyhashVersion(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"AddressPubkeyhashVersion\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndAddressChecksumValue(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"AddressChecksumValue\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndAddressChecksumValue(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"AddressChecksumValue\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndBurnAddress(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndBurnAddress(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndModel(value1 string, value2 int) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndModel(value1 string, value2 int) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndWizardKey(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"WizardKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndWizardKey(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"WizardKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

func (db *DB) GetNetworkByPrivateKeyVersionAndRPCPassword(value1 string, value2 string) (bool, *models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNetwork(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNetworkByPrivateKeyVersionAndRPCPassword(value1 string, value2 string) ([]*models.Network, error) {
	rows, err := db.DoQuery("SELECT \"AddressChecksumValue\", \"AddressPubkeyhashVersion\", \"BurnAddress\", \"ChainName\", \"Created\", \"DefaultRpcPort\", \"Model\", \"PrivateKeyVersion\", \"Public\", \"RPCPassword\", \"RPCUser\", \"Shared\", \"UID\", \"WizardKey\" FROM " + db.dbName + ".\"Network\" WHERE \"PrivateKeyVersion\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNetwork(rows)
}

