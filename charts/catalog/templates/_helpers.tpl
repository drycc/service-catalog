{{/* vim: set filetype=mustache: */}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "image" -}}
{{- if (.Values.image) }}
{{- printf "%s" .Values.image -}}
{{- else }}
{{- printf "%s/%s/service-catalog:%s" .Values.imageRegistry .Values.imageOrg .Values.imageTag -}}
{{- end }}
{{- end -}}