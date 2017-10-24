<!DOCTYPE html>

<html>
<head>
  <title>EasyNCV</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">

ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    overflow: hidden;
    background-color: #333;
}

li {
    float: left;
}

li a {
    display: block;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

/* Change the link color to #111 (black) on hover */
li a:hover {
    background-color: #111;
}


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
<ul>
  <li><a href="/">Home</a></li>
  <li><a href="/">Deploy</a></li>
  <li><a href="/stop">Undeploy</a></li>

</ul>
<br>
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
