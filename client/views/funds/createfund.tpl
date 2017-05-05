<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->      
      <!--search end-->
      {{template "inc/user-info.tpl" .}} 
    </div>
    <!-- header section end-->
    <!-- page heading start-->
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 创建基金
              <span class="tools pull-right"><a href="/funds/manage" class="fa fa-chevron-down">基金管理</a>
              </span>
            </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="createfund-form">
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">基金名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="fundname" class="form-control" placeholder="基金名称">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">初始基金数</label>
                  <div class="col-sm-10">
                    <input type="text" name="quotas" class="form-control" placeholder="初始基金数">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">初始资金数</label>
                  <div class="col-sm-10">
                    <input type="text" name="balance" class="form-control" placeholder="初始资金数">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">注册资金</label>
                  <div class="col-sm-10">
                    <input type="text" name="tbalance" class="form-control" placeholder="注册资金">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">注册时间</label>
                  <div class="col-sm-10">
                    <input type="text" name="ttime" class="form-control" placeholder="注册时间">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">入购起点</label>
                  <div class="col-sm-10">
                    <input type="text" name="tcount" class="form-control" placeholder="入购起点">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">限购单量</label>
                  <div class="col-sm-10">
                    <input type="text" name="tbuyper" class="form-control" placeholder="限购单量">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">限购总量</label>
                  <div class="col-sm-10">
                    <input type="text" name="tbuyall" class="form-control" placeholder="限购总量">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">基金净值</label>
                  <div class="col-sm-10">
                    <input type="text" name="netvalue" class="form-control" placeholder="基金净值">
                  </div>
                </div>

                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" value="{{.pro.Id}}">
                    <button type="submit" class="btn btn-primary">创建基金</button>
                  </div>
                </div>
              </form>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
</body>
</html>
