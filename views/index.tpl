{{ template "header.tpl" . }}
<form action="" method="post">
{{.Form | renderform}}
<br>
<br>
<br>
<input type=submit value='Deploy'/>
</form>
{{ template "tail.tpl" . }}
