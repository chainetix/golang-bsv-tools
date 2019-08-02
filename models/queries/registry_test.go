package database

import (
	"fmt"
	"bytes"
	"strings"
	"testing"
	"reflect"
	"io/ioutil"
)

func TestRegistry(t *testing.T) {

	createTableBuf := bytes.NewBuffer(nil)
	createTableBuf.WriteString(
		CONST_PACKAGE_DECLARATION,
	)
	createTableBuf.WriteString(
`func (db *DB) CreateTables() error {
`,
	)

	array := []func()(map[interface{}]interface{}, string, string){
		Registry_Project,
		Registry_Network,
	}

	for _, f := range array {
		registry, imports, modelpackage := f()
		Render(t, createTableBuf, registry, imports, modelpackage)
	}

	createTableBuf.WriteString(
`	return nil
}`,
	)

	ioutil.WriteFile(
		"gen/createTables.go",
		createTableBuf.Bytes(),
		0777,
	)

}

func Render(t *testing.T, createTableBuf *bytes.Buffer, registry map[interface{}]interface{}, imports, modelpackage string) {

	for unknownPtr, unknown := range registry {

		name := reflect.TypeOf(unknownPtr).Elem().Name()

		fmt.Println("DOING", name)

		createTableBuf.WriteString(
			fmt.Sprintf(
				CONST_TEMPLATE_CREATETABLE,
				name,
			),
		)

		args := []string{}
		values := []string{}
		fields := []string{}
		sortFields := []string{}
		fieldTypes := []string{}

		m := map[string]string{}
		DeepFields(m, unknown)

		for k, v := range m {

			fields = append(
				fields,
				k,
			)
			sortFields = append(
				sortFields,
				k,
			)
			fieldTypes = append(
				fieldTypes,
				v,
			)
			values = append(
				values,
				fmt.Sprintf("$%v", len(fields)),
			)
			args = append(
				args,
				fmt.Sprintf("%s %s", k, v),
			)

			fmt.Println("ADDING", name, k, v)

		}

		fmt.Println(fields, fieldTypes)

		buf := bytes.NewBuffer(nil)
		buf.WriteString(CONST_PACKAGE_DECLARATION + fmt.Sprintf(imports, modelpackage))

		buf.WriteString(
			renderCreate(name, fields, fieldTypes),
		)

		buf.WriteString(
			renderInsert(name, args, sortFields, values),
		)
		buf.WriteString(
			renderDelete(name),
		)
		buf.WriteString(
			renderPut(name, args, sortFields, values),
		)
		buf.WriteString(
			renderScan(name, sortFields),
		)

		buf.WriteString(
			renderCountAll(name, sortFields),
		)

		buf.WriteString(
			renderAll(name, sortFields),
		)

		for x, _ := range fields {

			buf.WriteString(
				renderCount(name, fields[x], fieldTypes[x]),
			)

			buf.WriteString(
				renderGet(name, fields[x], fieldTypes[x], sortFields),
			)

			buf.WriteString(
				renderQuery(name, fields[x], fieldTypes[x], sortFields),
			)

			if fields[x] != "UID" {
				for y, fieldName := range fields {

					// dont to it to itself
					if fieldName == fields[x] {
						continue
					}

					buf.WriteString(
						renderGetAnd(name, fields[x], fieldName, fieldTypes[x], fieldTypes[y], sortFields),
					)

					buf.WriteString(
						renderQueryAnd(name, fields[x], fieldName, fieldTypes[x], fieldTypes[y], sortFields),
					)

				}
			}

		}

		ioutil.WriteFile(
			fmt.Sprintf("gen/%s.go", strings.ToLower(name)),
			buf.Bytes(),
			0777,
		)

		//fmt.Println(string(buf.Bytes()))

	}

}

func renderCreate(structName string, fields, fieldTypes []string) string {

	createMap := map[string]string{
		"UID": "%s UUID PRIMARY KEY DEFAULT gen_random_uuid()",
		"Salt": "%s UUID DEFAULT gen_random_uuid()",
		"Created": "%s TIMESTAMP DEFAULT current_timestamp()",
	}

	typesMap := map[string]string{
		"[]string": "STRING ARRAY",
		"float64": "FLOAT",
		"bool": "BOOL",
		"int": "INT",
		"string": "STRING",
		"map[string]interface {}": "JSON",
	}

	createFields := []string{}

	for x, field := range fields {

		pf := field

		field = fmt.Sprintf(`\"%s\"`, field)

		s, ok := createMap[pf]
		if ok {
			createFields = append(
				createFields,
				fmt.Sprintf(s, field),
			)
		} else {
			s, ok := typesMap[fieldTypes[x]]
			if ok {
				createFields = append(
					createFields,
					fmt.Sprintf("%s %s NOT NULL", field, s),
				)
			} else {
				createFields = append(
					createFields,
					fmt.Sprintf("%s %s NOT NULL", field, fieldTypes[x]),
				)
			}
		}

	}

	return fmt.Sprintf(
		CONST_TEMPLATE_SQL_CREATETABLE,
		structName,
		quoteTable(structName),
		strings.Join(createFields, ", "),
	)
}
