{{ template "../header.tpl" . }}


<form action="" method="post">
{{.Form | renderform}}
<br>
<br>
<br>
<input type=submit value='Undeploy'/>
</form>
{{ template "../tail.tpl" . }}