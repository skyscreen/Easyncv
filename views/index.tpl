<!DOCTYPE html>

<html>
<head>
  <title>EasyNCV</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }






    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }


  </style>
</head>

<body>
<center/>
  <header>


    <h1 class="logo">NCV configuration example</h1>
    <br>
    <br>
    <div class="description">
      Please input values for hcl file
    </div>
  </header>
<form action="" method="post">
{{.Form | renderform}}
<br>
<br>
<br>
<input type=submit value='Deploy'/>
</form>
  <footer>
    <div class="author">

      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
