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
            <header class="panel-heading"> 我的基金 总数：{{.countFunds}}
              <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              </span>
            </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="funds-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>基金名称</th>
                        <th>成立日期</th>
                        <th>规模(份)</th>
                        <th>市值(元)</th>
                        <th>净值</th>
                      </tr>
                    </thead>
                    <tbody>
                    
                    {{range $k,$v := .funds}}
                    <tr>
                      <td><a href="/funds/{{$v.Id}}">{{$v.Name}}</a></td>
                      <td>{{$v.CreateTime}}</td>
                      <td>{{$v.Quotas}}</td>
                      <td>{{$v.MarketValue}}</td>
                      <td>{{$v.NetValue}}</td>                     
                    </tr>
                    {{end}}
                    </tbody>
                  </table>
                </form>
				      </section>
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
