<?xml version="1.0" encoding="UTF-8"?>
<object ifEmpty="true">
    {{- range . }}
    <Currency digitCode="{{.col0}}" symbolCode="{{.col1}}" name="{{.col2}}" id="{{GUID}}"/>
    {{- end }}
</object>