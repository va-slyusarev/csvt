<?xml version="1.0" encoding="UTF-8"?>
<object ifEmpty="true">
    {{- range . }}
    <Country code="{{.col2}}" nameShort="{{.col0}}" nameFull="{{.col1}}" alfa2="{{.col3}}" alfa3="{{.col4}}" id="{{GUID}}"/>
    {{- end }}
</object>