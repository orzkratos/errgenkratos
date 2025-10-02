{{ range .Errors }}
{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Is{{.CamelValue}}(err error) bool {
	return errgenkratos.IsError(err, {{if .Name}}{{ .Name }}_{{end}}{{ .Value }}, {{ .HTTPCode }})
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Error{{ .CamelValue }}(format string, args ...interface{}) *errors.Error {
	return errgenkratos.NewError({{ .HTTPCode }}, {{if .Name}}{{ .Name }}_{{end}}{{ .Value }}, format, args...)
}

{{- end }}