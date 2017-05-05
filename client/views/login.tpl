<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<style>
.form-signin .help-block{color:#a94442;}
</style>
</head><body class="login-body">
<div class="container">
  <form class="form-signin" id="login-form">
    <div class="form-signin-heading text-center">
      <h1 class="sign-title">请输入用户名和密码登录系统</h1>
      <img src="/static/img/logo.png" alt="基金管理系统" style="width:120px;"/> </div>
    <div class="login-wrap">
      <input type="text" class="form-control" name="username" placeholder="请填写用户名" autofocus>
      <input type="password" class="form-control" name="password" placeholder="请填写密码">
      <button class="btn btn-lg btn-login btn-block" type="submit"> 登录</button>
    </div>
  </form>
</div>
{{template "inc/foot.tpl" .}}
</body>
</html>

