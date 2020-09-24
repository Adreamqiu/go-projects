package sql2struct

import (
	"fmt"
	"text/template"
	"os"
	"github.com/go-projects/tour/internal/word"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range.Columns}} {{ $length := len .Comment}} {{if gt $length 0 }}  // {{.Comment}} {{else}}  // {{.Name}} {{end}}
   {{ $typeLen := len .Type}} {{if gt $typeLen 0 }}{{.Name | ToCamelCase}} {{.Type}}  {{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
    return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name string
	Type string
	Tag string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn {
			Name: column.ColumnName,
			Type: DBTypeToStructType[column.DataType],
			Tag: fmt.Sprintf("`json:"+"%s"+"`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns: tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}



type SkuData struct {
  // id
Id int32  `json:id`
  // 点位编号
   MacNo int32  `json:MacNo`
  // 上线第几周
   Week int8  `json:week`
  // 商品ID
   GoodsID int32  `json:GoodsID`
  // 销量
   Sale int32  `json:Sale`
}

func (model SkuData) TableName() string {
        return "sku_data"
}