# CSV Transformer

Utility for converting data from CSV to the format specified by the go template.

```
Usage of csvt:
  -skip
        Skip header line. (default true)
  -src string
        Input data file path. (default "in.csv")
  -tar string
        Out data file path. (default "out.txt")
  -tpl string
        Go template file. (default "out.tpl")
```

## csv
The csv file uses a field separator `,`

## helper functions
`GUID` - randomly generated hexadecimal string of the form `ac1b73df-5e35-f216-b6ef-4957b9bc5781`

## usage
Data of csv is passed to the template as `[]map[string]string`.
The value key in each column is `map[string]string` - `colX`, where `X` is the sequence number starting from `0`

### example of usage in a template

csv
```csv
ISO_DIG,ISO_LAT3,NAME_RUS
008,ALL,ЛЕК
012,DZD,АЛЖИРСКИЙ ДИНАР
032,ARS,АРГЕНТИНСКОЕ ПЕСО
```

tpl
```gotemplate
<?xml version="1.0" encoding="UTF-8"?>
<object ifEmpty="true">
    {{- range . }}
    <Currency digitCode="{{.col0}}" symbolCode="{{.col1}}" name="{{.col2}}" id="{{GUID}}"/>
    {{- end }}
</object>
```

out
```xml
<?xml version="1.0" encoding="UTF-8"?>
<object ifEmpty="true">
    <Currency digitCode="008" symbolCode="ALL" name="ЛЕК" id="e3120fa7-a818-4adc-3916-edae2397a16d"/>
    <Currency digitCode="012" symbolCode="DZD" name="АЛЖИРСКИЙ ДИНАР" id="d1d58de5-a30a-2a17-c432-049d0c5a7e72"/>
    <Currency digitCode="032" symbolCode="ARS" name="АРГЕНТИНСКОЕ ПЕСО" id="1a947ef9-63ab-c1f3-ccfd-57bba4087c73"/>
</object>
```

