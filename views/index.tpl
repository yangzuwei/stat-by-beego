<!DOCTYPE html>
<html>
<head>
<!-- 新 Bootstrap 核心 CSS 文件 -->
<link href="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
 
<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
<script src="https://cdn.staticfile.org/jquery/2.1.1/jquery.min.js"></script>
 
<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>简报</title>

</head>
<body class="container">
<header>
<h1 class="text-center">统计简报</h1>
</header>
<div class="panel panel-default">
    <div class="panel-body">
        截至{{.Current}}完成情况统计
    </div>
</div>

  <div class="progress progress-striped active">
	<div class="progress-bar progress-bar-success" role="progressbar"
		 aria-valuenow="{{mul 100 (div .HasHandle .TotalNum)}}" aria-valuemin="0" aria-valuemax="100"
		 style="width: {{mul 100 (div .HasHandle .TotalNum)}}%;">
		<span class="show">当前学籍总人数为:{{.TotalNum}}其中已经处理了{{.HasHandle}}</span>
	</div>
</div>
<table class="table">
  <caption>发送情况（包括补和重复再洗的）一共：{{.Actually}}</caption>
  <thead>
    <tr>
      <th>日期</th>
      <th>数量</th>
    </tr>
  </thead>
  <tbody>
{{range $v := .Accounts}}
    <tr>
      <td>{{$v.Date}}</td>
      <td>{{$v.Num}}</td>
    </tr>
{{end}}
  </tbody>
</table>

  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> 
	<br>
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
