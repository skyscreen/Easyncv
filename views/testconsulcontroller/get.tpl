{{ template "../header.tpl" . }}




<form action="" method="post">
{{.Form | renderform}}
<br>
<br>
<br>
<input type=submit value='Test'/>
</form>

<br>
<center/>history record <br>

<table border="1">
<tr><td align=center>request server</td><td align=center>description</td><td align=center>Consul Server</td></tr>
{{range $k := .ConsulRecs}}
	<tr><td align=center><b>{{$k.Name}} </b> </td><td align=center> {{$k.Describtion}} </td><td align=center> {{$k.ServerName}} </td></tr>
{{end}}
<table>
{{ template "../tail.tpl" . }}
