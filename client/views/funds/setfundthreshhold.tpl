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
            <header class="panel-heading"> 设置基金限制
              <span class="tools pull-right"><a href="/funds/manage" class="fa fa-chevron-down">基金管理</a>
              </span>
            </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="setfundthreshhold-form">
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">基金名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="fundname" class="form-control" readonly value="{{.fund.Name}}">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">注册资金</label>
                  <div class="col-sm-10">
                    <input type="text" name="tbalance" class="form-control" value="{{.fund.PartnerAssets}}">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">注册时间</label>
                  <div class="col-sm-10">
                    <input type="text" name="ttime" class="form-control" value="{{.fund.PartnerTime}}">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">入购起点</label>
                  <div class="col-sm-10">
                    <input type="text" name="tcount" class="form-control" value="{{.fund.BuyStart}}">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">限购单量</label>
                  <div class="col-sm-10">
                    <input type="text" name="tbuyper" class="form-control" value="{{.fund.BuyPer}}">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">限购总量</label>
                  <div class="col-sm-10">
                    <input type="text" name="tbuyall" class="form-control" value="{{.fund.BuyAll}}">
                  </div>
                </div>

                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" value="{{.pro.Id}}">
                    <button type="submit" class="btn btn-primary">设置基金限制</button>
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
