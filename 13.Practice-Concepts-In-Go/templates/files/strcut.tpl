package modelss

{{ range . }}
// {{ .Desc }}
type {{ .Name }} struct {
    {{ range .Fields -}}
    {{ .Name }} {{ .Type }}
    {{ end }}
}
{{ end }}
