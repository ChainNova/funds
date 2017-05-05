<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
<link href="/static/css/clndr.css" rel="stylesheet">
<link href="/static/css/table-responsive.css" rel="stylesheet">
{{template "inc/meta.tpl" .}}
</head>
</head>
<body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <a class="toggle-btn">
        <i class="fa fa-bars"></i>
      </a>
      {{template "inc/user-info.tpl" .}}
      </div>
    <!-- header section end-->
    <!-- page heading start-->
    <!--<div class="page-heading">-->
    <!--Page Tittle goes here-->
    <!--</div>-->
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-md-12">

          <!-- 基金信息 -->
            <div class="col-md-8">
              <div class="panel">
                <div class="panel-body">
                  <div class="row">
                    <!-- 第一行 -->
                    <div class="col-md-12">
                      <div class="col-md-4">
                          <h2>{{.myfund.Name}}</h2>
                      </div>

                      <div class="col-md-3">
                        <div class="row">
                          <div class="col-md-12">
                            <h5>成立日期</h5>
                          </div>
                          <div class="col-md-12">
                            <ul class="p-info">
                              <div class="desk">{{.myfund.CreateTime}}</div>
                            </ul>
                          </div>
                        </tr>
                        </div>
                      </div>

                      <div class="col-md-2">
                        <div class="row">
                          <div class="col-md-12">
                            <h5>规模(份)</h5>
                          </div>
                          <div class="col-md-12">
                            <ul class="p-info">
                              <div class="desk">{{.myfund.Quotas}}</div>
                            </ul>
                          </div>
                        </tr>
                        </div>
                      </div>

                      <div class="col-md-3">
                        <div class="row">
                          <div class="col-md-12">
                            <h5>市值(元)</h5>
                          </div>
                          <div class="col-md-12">
                            <ul class="p-info">
                              <div class="desk">{{.myfund.MarketValue}}</div>
                            </ul>
                          </div>
                        </tr>
                        </div>
                      </div>

                      
                    </div>

                    <!-- 第二行 -->
                    <div class="col-md-12">
                      <div class="col-md-2">
                        <div class="row">
                          <div class="col-md-12">
                            <h5>最低投资额(元)</h5>
                          </div>
                          <div class="col-md-12">
                            <ul class="p-info">
                              <div class="desk">{{.myfund.ThresholdValue}}</div>
                            </ul>
                          </div>
                        </tr>
                        </div>
                      </div>

                      <div class="col-md-2">
                        <div class="row">
                          <div class="col-md-12">
                            <h5>净值</h5>
                          </div>
                          <div class="col-md-12">
                            <ul class="p-info">
                              <div class="desk" name="netvalue">{{.myfund.NetValue}}</div>
                            </ul>
                          </div>
                        </tr>
                        </div>
                      </div>

                      <div class="col-md-3">
                        <div class="row">
                          <div class="col-md-12">
                            <h5>单日变动</h5>
                          </div>
                          <div class="col-md-12">
                            <ul class="p-info">
                              <div class="desk">{{.myfund.NetDelta}}</div>
                            </ul>
                          </div>
                        </tr>
                        </div>
                      </div>

                      <div class="col-md-2">
                        <div class="row">
                          <div class="col-md-12">
                            <h5>持有份额</h5>
                          </div>
                          <div class="col-md-12">
                            <ul class="p-info">
                              <div class="desk" name="myquotas">{{.myfund.MyQuotas}}</div>
                            </ul>
                          </div>
                        </tr>
                        </div>
                      </div>

                      <div class="col-md-3">
                        <div class="row">
                          <div class="col-md-12">
                            <h5>参考市值</h5>
                          </div>
                          <div class="col-md-12">
                            <ul class="p-info">
                              <div class="desk" name="mymarketvalue">{{.myfund.MyMarketValue}}</div>
                            </ul>
                          </div>
                        </tr>
                        </div>
                      </div>

                    </div>

                  </div>  
                </div>
              </div>
            </div>

            <div class="col-md-4">
              <div class="panel">
                <div class="panel-body">
                  <div class="container col-md-12">
                    <ul class="nav nav-tabs" id="buyredeem-tab">
                      <li class="active">
                        <a data-toggle="tab" href="#buy">购买</a>
                      </li>
                      <li>
                        <a data-toggle="tab" href="#redeem">赎回</a>
                      </li>
                    </ul>
                  
                    <div class="tab-content">
                        <div id="buy" class="tab-pane fade in active ">
                          <div class="row">
                            <div class="col-md-12" style="height:8px;"></div>
                            <div class="col-md-12">
                              <p>投资金额：<lable class="pull-right" name="myaccount">可用余额：{{.myfund.MyBalance}}</lable></p>
                            </div>
                            <div class="col-md-12">
                              <form class="form-horizontal adminex-form" id="buyfund-form">
                                <input type="hidden" name="fundid" value="{{.fundid}}">
                                <input type="hidden" name="netvalue" value="{{.myfund.NetValue}}">
                                <div class="form-group">
                                  <div class="col-md-12">
                                    <input type="text" name="buycount" class="form-control" placeholder="请填购买份额">
                                  </div>
                                  <div class="col-md-12" style="height:8px;"></div>
                                  <div class="col-md-12">
                                    <button type="submit" class="btn btn-primary" id="buyredeem-btn">购买基金</button>
                                  </div>
                                </div>
                              </form>
                            </div>
                          </div>
                        </div>

                        <div id="redeem" class="tab-pane fade">
                          <div class="row">
                            <div class="col-md-12" style="height:8px;"></div>
                            <div class="col-md-12">
                              <p>赎回份额：<lable class="pull-right">可用份额：{{.myfund.MyQuotas}}</lable></p>
                            </div>
                            <div class="col-md-12">
                              <form class="form-horizontal adminex-form" id="redeemfund-form">
                                <input type="hidden" name="fundid" value="{{.fundid}}">
                                <div class="form-group">
                                  <div class="col-md-12">
                                    <input type="text" name="redeemcount" class="form-control" placeholder="请填赎回份额">
                                  </div>
                                  <div class="col-md-12" style="height:8px;"></div>
                                  <div class="col-md-12">
                                    <button type="submit" class="btn btn-primary" id="buyredeem-btn">赎回基金</button>
                                  </div>
                                </div>
                              </form>
                            </div>
                          </div>
                        </div>

                    </div>
                  </div>
                </div>
              </div>
            </div>

          <!--<div class="row">-->
            <div class="col-md-8">
              <div class="panel">
                <div class="panel-body">
                  <div class="profile-desk">
                    <h1>净值走势</h1>
                    <div id="netvaluechart" style="height:300px; min-width: 310px"></div>
                  </div>
                </div>
              </div>
            </div>
          <!--</div>-->

          <!--<div class="row">-->
            <div class="col-md-4">
              <div class="panel">
                <div class="panel-body">
                  <div class="profile-desk">
                    <h1>市场动态<a class="pull-right" style="font-size:16px;" href="#">更多</a></h1>
                    <table class="table table-bordered table-striped table-condensed cf">
                      <thead class="cf">
                        <tr>
                          <th class="numeric">序号</th>
                          <th class="numeric">规模</th>
                          <th>类型</th>
                        </tr>
                      </thead>
                      <tbody>
                      
                      {{range $k,$v := .markets}}
                      <tr>
                        <td>{{$v.Index}}</td>
                        <td>{{$v.Size}}</td>
                        <td>{{$v.Type}}</td>
                      </tr>
                      {{end}}
                      </tbody>
                      
                    </table>
                  </div>
                </div>
              </div>
            </div>
          <!--</div>-->

          <!--<div class="row">-->
            <div class="col-md-12">
              <div class="panel">
                <header class="panel-heading"> 基金公告 <span class="pull-right"> <a href="#">更多</a></span> </header>
                <div class="panel-body">
                  <ul class="activity-list">
                    {{range $k,$v := .notices}}
                    <li>
                      <div>
                        <i class="fa fa-caret-right" aria-hidden="true"></i>
                        <a href="#">{{$v.News}}</a>
                        <div class="pull-right text-muted">{{$v.Date}}</div>
                      </div>
                    </li>
                    {{end}}
                  </ul>
                </div>
              </div>
            </div>
          <!--</div>-->

        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
	{{template "inc/notice-dialog.tpl" .}}
    {{template "inc/foot-info.tpl" .}}	
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script src="/static/js/calendar/clndr.js"></script>
<script src="/static/js/calendar/evnt.calendar.init.js"></script>
<script src="/static/js/calendar/moment-2.2.1.js"></script>
<script src="/static/js/underscore-min.js"></script>
<script src="/static/js/highstock.js"></script>
<script src="/static/js/exporting.js"></script>
<script>
$(function(){
	$('#noticeModal').on('show.bs.modal', function (e) {
		$('#notice-box').html($(e.relatedTarget).attr('data-content'))
	});

  //$.getJSON('https://www.highcharts.com/samples/data/jsonp.php?filename=aapl-c.json&callback=?', function (data) {
    
        // Global options
        Highcharts.setOptions({
            lang:{
                rangeSelectorZoom: ''
            }
        });

        // Create the chart
        $('#netvaluechart').highcharts('StockChart', {
            credits: {  
                enabled: false  
            },

            exporting: {
                enabled: false
            },

            rangeSelector: {
                selected: 1,
                inputEnabled : false,
                buttons : [{
                    type : 'week',
                    count : 1,
                    text : '1周'
                }, {
                    type : 'month',
                    count : 1,
                    text : '1月'
                }, {
                    type : 'month',
                    count : 6,
                    text : '半年'
                }, {
                    type : 'yeal',
                    count : 3,
                    text : '1年'
                }, {
                    type : 'all',
                    count : 1,
                    text : '所有'
                }],
            },

            title: {
                text: '净值走势'
            },
            
            navigator : {
                enabled : false
            },

            series: [{
                name: '净值',
                data: {{$.netLog}},
                type: 'areaspline',
                threshold: null,
                tooltip: {
                    valueDecimals: 2,
                    dateTimeLabelFormats: "%Y-%m-%d"
                },
                fillColor: {
                    linearGradient: {
                        x1: 0,
                        y1: 0,
                        x2: 0,
                        y2: 1
                    },
                    stops: [
                        [0, Highcharts.getOptions().colors[0]],
                        [1, Highcharts.Color(Highcharts.getOptions().colors[0]).setOpacity(0).get('rgba')]
                    ]
                }
            }]
        });
   // });
})
</script>
</body>
</html>
