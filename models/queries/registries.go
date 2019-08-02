package database

import (
	"fmt"
	"sort"
	"strings"
	"reflect"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
	netmodels "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/network"
)

func quoteTable(s string) string {
	return fmt.Sprintf(`" + db.dbName + ".\"%s\"`, s)
}

func quote(s string) string {
	return fmt.Sprintf(`\"%s\"`, s)
}

func DeepFields(m map[string]string, iface interface{}) {

	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	for i := 0; i < ift.NumField(); i++ {

		f := ifv.Field(i)
		k := ifv.Type().Field(i).Name
		t := ifv.Type().Field(i).Type

		switch f.Kind() {
		case reflect.Struct:
			if k == "Created" {
				break
			}
			DeepFields(m, f.Interface())
			continue
		}

		v := fmt.Sprintf("%v", t)
		m[k] = v

	}
}

func Registry_Project() (map[interface{}]interface{}, string, string) {
	return map[interface{}]interface{}{
//		&models.VerboseBlock{}: models.VerboseBlock{},
		&models.ApiToken{}: models.ApiToken{},
		&models.Project{}: models.Project{},
		&models.Asset{}: models.Asset{},
		&models.User{}: models.User{},
		&models.Agent{}: models.Agent{},
		&models.UserGroup{}: models.UserGroup{},
		&models.Wallet{}: models.Wallet{},
		&models.Address{}: models.Address{},
		&models.Stream{}: models.Stream{},
		&models.Currency{}: models.Currency{},
		&models.Permission{}: models.Permission{},
		&models.StreamItem{}: models.StreamItem{},
		&models.Exchange{}: models.Exchange{},
		&models.ProjectStats{}: models.ProjectStats{},
		&models.BillingRecord{}: models.BillingRecord{},
		&models.Input{}: models.Input{},
		&models.Output{}: models.Output{},
		&models.FeedMessage{}: models.FeedMessage{},
		&models.Transaction{}: models.Transaction{},
	}, CONST_PACKAGE_IMPORTS_PROJECT, "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
}

func Registry_Network() (map[interface{}]interface{}, string, string) {
	return map[interface{}]interface{}{
		&netmodels.Network{}: netmodels.Network{},
		&netmodels.Node{}: netmodels.Node{},
	}, CONST_PACKAGE_IMPORTS_NETWORK, "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/network"
}

func renderPut(structName string, args, fields, values []string) string {

	sort.Strings(fields)
	sort.Strings(args)

	filter := map[string]bool{
		"uid": true,
		"created": true,
		"salt": true,
	}

	f := []string{}
	a := []string{}
	for x, _ := range fields {
		if filter[strings.ToLower(fields[x])] {
			continue
		}
		f = append(f, fields[x])
		a = append(a, args[x])
	}

	ff := []string{}
	for _, x := range f {
		ff = append(
			ff,
			fmt.Sprintf("model.%s", x),
		)
	}

	fff := []string{}
	for _, x := range f {
		fff = append(
			fff,
			quote(x),
		)
	}

	scans := []string{}
	for x, _ := range fields {
		scans = append(
			scans,
			fmt.Sprintf("&model.%s", fields[x]),
		)
	}

	scanFields := []string{}
	for x, _ := range fields {
		scanFields = append(
			scanFields,
			quote(fields[x]),
		)
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_PUT,
		structName,
		fmt.Sprintf("model *models.%s", structName),
		//
		quoteTable(structName),
		strings.Join(fff, ", "),
		strings.Join(values[:len(f)], ", "),
		strings.Join(scanFields, ", "),
		strings.Join(ff, ", "),
		strings.Join(scans, ", "),
	)
}

func renderInsert(structName string, args, fields, values []string) string {

	sort.Strings(fields)
	sort.Strings(args)

	filter := map[string]bool{
		"uid": true,
		"created": true,
		"salt": true,
	}

	f := []string{}
	ff := []string{}
	a := []string{}
	for x, _ := range fields {
		if filter[strings.ToLower(fields[x])] {
			continue
		}
		f = append(f, fields[x])
		ff = append(ff, quote(fields[x]))
		a = append(a, args[x])
	}

	scans := []string{}
	for x, _ := range fields {
		scans = append(
			scans,
			fmt.Sprintf("&row.%s", fields[x]),
		)
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_INSERT,
		structName,
		strings.Join(a, ", "),
		structName,
		structName,
		quoteTable(structName),
		strings.Join(ff, ", "),
		strings.Join(values[:len(f)], ", "),
		strings.Join(f, ", "),
		strings.Join(f, ", "),
		strings.Join(scans, ", "),
	)
}

func renderDelete(structName string) string {

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_DELETE,
		structName,
		quoteTable(structName),
		quote("UID"),
	)
}

func renderCount(structName, field, fieldType string) string {

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_COUNT,
		structName,
		strings.Title(field),
		fieldType,
		quote(structName),
		quote(field),
	)
}

func renderCountAll(structName string, sortFields []string) string {

	sort.Strings(sortFields)

	sf := make([]string, len(sortFields))
	for x, sortField := range sortFields {
		sf[x] = fmt.Sprintf("&%s.%s", structName, strings.Title(sortField))
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_COUNTALL,
		structName,
		quoteTable(structName),
	)
}

func renderQuery(structName, field, fieldType string, sortFields []string) string {

	sort.Strings(sortFields)

	sf := make([]string, len(sortFields))
	for x, sortField := range sortFields {
		sf[x] = quote(sortField)
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_QUERY,
		structName,
		strings.Title(field),
		fieldType,
		structName,
		strings.Join(sf, ", "),
		quoteTable(structName),
		quote(field),
		structName,
	)
}

func renderQueryAnd(structName, field1, field2, fieldType1, fieldType2 string, sortFields []string) string {

	sort.Strings(sortFields)

	sf := make([]string, len(sortFields))
	for x, sortField := range sortFields {
		sf[x] = quote(sortField)
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_QUERYAND,
		structName,
		strings.Title(field1),
		strings.Title(field2),
		fieldType1,
		fieldType2,
		structName,
		strings.Join(sf, ", "),
		quoteTable(structName),
		quote(field1),
		quote(field2),
		structName,
	)
}

func renderAll(structName string, sortFields []string) string {

	sort.Strings(sortFields)

	sf := make([]string, len(sortFields))
	for x, sortField := range sortFields {
		sf[x] = quote(sortField)
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_ALL,
		structName,
		structName,
		strings.Join(sf, ", "),
		quoteTable(structName),
		structName,
	)
}

func renderGet(structName, field, fieldType string, sortFields []string) string {

	sort.Strings(sortFields)

	sf := make([]string, len(sortFields))
	for x, sortField := range sortFields {
		sf[x] = quote(sortField)
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_GET,
		structName,
		strings.Title(field),
		fieldType,
		structName,
		strings.Join(sf, ", "),
		quoteTable(structName),
		quote(field),
		structName,
	)
}

func renderGetAnd(structName, field1, field2, fieldType1, fieldType2 string, sortFields []string) string {

	sort.Strings(sortFields)

	sf := make([]string, len(sortFields))
	for x, sortField := range sortFields {
		sf[x] = quote(sortField)
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_GETAND,
		structName,
		strings.Title(field1),
		strings.Title(field2),
		fieldType1,
		fieldType2,
		structName,
		strings.Join(sf, ", "),
		quoteTable(structName),
		quote(field1),
		quote(field2),
		structName,
	)
}

func renderScan(structName string, sortFields []string) string {

	sort.Strings(sortFields)

	sf := make([]string, len(sortFields))
	for x, sortField := range sortFields {
		sf[x] = fmt.Sprintf("&result.%s", sortField)
	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_ROWSCAN,
		structName,
		structName,
		structName,
		structName,
		strings.Join(sf, ", "),
	)
}
