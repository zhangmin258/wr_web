$(function(){
	var userToken=$('meta[name=appToken]').attr('content');
	var packageId=$('meta[name=packageId]').attr('content');
	window.document.addEventListener('message',function(e){
	    var message = e.data;
	    if( message=="MD1" ){
	    	window.location.href="/signh5/gettask?Token="+userToken+"&PackageId="+packageId;
	    }
	    if( message=="SignCalendar" ){
	    	$("#calendar").show();
	    }
	});
	//阻止它默认事件
    document.addEventListener("touchmove",function(e){
        e.preventDefault();
    });
	//签到提醒
	var isRemind;
	$('#goRemind').on("click",function(){
		if( isRemind==false ){
			isRemind=true;
			$(this).attr("src","../../static/img/package2/alertOpen.png");
			tuiter(1);
		}else{
			isRemind=false;
			$(this).attr("src","../../static/img/package2/alertClose.png");
			tuiter(0);
		}
	});
	function tuiter(state){
		$.ajax({
			type:"post",
			url:"/sign/getremindstatus",
			contentType:"application/json;charset=utf-8",
			data:{status:state},
			dataType:"text",
			cache:false,
			success:function(data){
				if( data.ret==200 ){
					var alertMsg;
					if( data.status==0 ){
						alertMsg="签到提醒已关闭！";
					}else{
						alertMsg="签到提醒已开启！";
					}
					window.postMessage("Alert+"+alertMsg,'*'); 
				}else{
					window.postMessage("Alert+"+data.err,'*'); 
				}
			}
		});
	}
	//获取签到页面数据（是否签到、额外奖励数据）
	$.ajax({
		type:"post",
		url:"/sign/getsign",
		contentType:"application/json;charset=utf-8",
		data:{},
		cache:false,
		dataType:"text",
		success:function(data){
			if( data.ret!=200 ){
				return;
			}
			if( data.RemindStatus==0 ){
				isRemind=false;
				$("#goRemind").attr("src","../../static/img/package2/alertClose.png");
			}else{
				isRemind=true;
				$("#goRemind").attr("src","../../static/img/package2/alertOpen.png");
			}
			if( data.status==1 ){
				$("#signIn").attr("src","../../static/img/package2/signIned.png");
				$("#signIn").on("click",function(){
					window.postMessage("Alert+已签到",'*'); 
					return;
				});
			}else{
				$("#signIn").on("click",function(){
					$.ajax({
						type:"post",
						url:"/sign/addsign",
						contentType:"application/json;charset=utf-8",
						data:{},
						cache:false,
						dataType:"text",
						success:function(data){
							if( data.ret==200 ){
								window.postMessage("Sign+5",'*');
								$("#signIn").attr("src","../../static/img/package2/signIned.png");
								queryPeas();
								queryDays(2);
								getCalendar();
							}else{
								window.postMessage("Alert+"+data.msg,'*'); 
							}
						}
					});
				});
			}
			$.each(data.SignReward,function(index){
				if( data.SignReward[index].IsReceive==0 ){
					//未领取
					$("#navBtn img").eq(index).attr("src","../../static/img/package2/"+data.SignReward[index].SignCount+"CanGet.png");
				}else if( data.SignReward[index].IsReceive==1 ){
					//已领取
					$("#navBtn img").eq(index).attr("src","../../static/img/package2/"+data.SignReward[index].SignCount+"haveGet.png");
				}else if( data.SignReward[index].IsReceive==2 ){
					//未完成
					$("#navBtn img").eq(index).attr("src","../../static/img/package2/"+data.SignReward[index].SignCount+"Can'tGet.png");
				}
				$("#navBtn img").eq(index).attr("name",data.SignReward[index].SignCount);
			});
		}
	});
	//获取累计签到天数
	queryDays(1);
	function queryDays(key){
		$.ajax({
			type:"post",
			url:"/sign/getsigncount",
			contentType:"application/json;charset=utf-8",
			data:{},
			cache:false,
			dataType:"text",
			success:function(data){
				if( data.ret!=200 ){
					return;
				}
				$(".SignCount").html( data.SignCount );
				if( data.SignCount==0 ){
					$("#progressBar span").width("0rem");
				}else if( data.SignCount<3 ){
					$("#progressBar span").width("0.65rem");
				}else if( data.SignCount==3 ){
					$("#progressBar span").width("1.3rem");
					if( key==2 ){
						$("#navBtn img").eq(0).attr("src","../../static/img/package2/3CanGet.png");
					}
				}else if( data.SignCount<7 ){
					$("#progressBar span").width("2.75rem");
				}else if( data.SignCount==7 ){
					$("#progressBar span").width("4.2rem");
					if( key==2 ){
						$("#navBtn img").eq(1).attr("src","../../static/img/package2/7CanGet.png");
					}
				}else if( data.SignCount<12 ){
					$("#progressBar span").width("5.56rem");
				}else if( data.SignCount==12 ){
					$("#progressBar span").width("7.1rem");
					if( key==2 ){
						$("#navBtn img").eq(2).attr("src","../../static/img/package2/12CanGet.png");
					}
				}else if( data.SignCount<18 ){
					$("#progressBar span").width("8.55rem");
				}else if( data.SignCount==18 ){
					$("#progressBar span").width("10rem");
					if( key==2 ){
						$("#navBtn img").eq(3).attr("src","../../static/img/package2/18CanGet.png");
					}
				}else if( data.SignCount<25 ){
					$("#progressBar span").width("11.45rem");
				}else if( data.SignCount==25 ){
					$("#progressBar span").width("13.8rem");
					if( key==2 ){
						$("#navBtn img").eq(4).attr("src","../../static/img/package2/25CanGet.png");
					}
				}
			}
		});
	}
	//点击融豆领取
	$("#navBtn img").on("click",function(){
		var thisName=$(this).attr('name');
		var that=this;
		$.ajax({
			type:"post",
			url:"/sign/getextrareward",
			contentType:"application/json;charset=utf-8",
			data:{Id:$(this).index()+1},
			dataType:"text",
			cache:false,
			success:function(data){
				if( data.ret==200 ){
					window.postMessage("at+"+$(that).attr("pCount"),'*');
					$(that).attr("src","../../static/img/package2/"+thisName+"haveGet.png");
					queryPeas();
				}else{
					window.postMessage("Alert+"+data.msg,'*');
				}
			}
		});
	});
	//查询融豆数量
	queryPeas();
	function queryPeas(){
		$.ajax({
			type:"post",
			url:"/sign/getscore",
			contentType:"application/json;charset=utf-8",
			data:{},
			cache:false,
			dataType:"text",
			success:function(data){
				if( data.ret!=200 ){
					return;
				}
				$("#ScoreCount").html(data.Score);
			}
		});
	}
	//获取可完成任务数量
	$.ajax({
		type:"post",
		url:"/sign/getmissioncount",
		contentType:"application/json;charset=utf-8",
		data:{},
		cache:false,
		dataType:"text",
		success:function(data){
			if( data.ret!=200 ){
				return;
			}
			$("#taskCount").html(data.Count);
		}
	});
	//获取签到日历已经签到的日期数组
	getCalendar();
	function getCalendar(){
	    var dateArray = [];
	    $.ajax({
	    	type:"post",
	    	url:"/sign/getdays",
	    	contentType:"application/json;charset=utf-8",
	    	data:{},
	    	dataType:"text",
	    	cache:false,
	    	success:function(data){
	    		if( data.ret!=200 ){
	    			return;
	    		}
	    		$.each(data.Days, function(index){
	    			dateArray.push( data.Days[index] );
	    		});
	    		//生成日历网格
			    var qdLi = "";
			    for (var i = 0; i < 42; i++) {
			        qdLi += "<li><div class='qiandao-icon'></div><span></span></li>";
			    }
			    $("#qiandaoList").html(qdLi); 
			    //获取当前月的天数
			    var myDate = new Date();
			    $("#nowYM").html( myDate.getFullYear()+"年"+(myDate.getMonth()+1)+"月" );
			    var monthFirst = new Date(myDate.getFullYear(), parseInt(myDate.getMonth()), 1).getDay();
			    var d = new Date(myDate.getFullYear(), parseInt(myDate.getMonth() + 1), 0);
			    var totalDay = d.getDate();
			    //生成当月的日历且含已签到
			    var $dateLi = $("#qiandaoList").find("li");
			    $dateLi.eq( monthFirst+myDate.getDate()-1 ).find('div').css({background:"url(../../static/img/package2/nowDay.png) no-repeat center 0.05rem",backgroundSize:"1.1rem","display":"block"});
			    for (var i = 0; i < totalDay; i++) {
			        $dateLi.eq(i + monthFirst).find('span').html(parseInt(i + 1));
			        for (var j = 0; j < dateArray.length; j++) {
			            if (i == dateArray[j]-1) {
			                $dateLi.eq(i + monthFirst).find('div').css("display","block");
			            }
			        }
			    } 
	    	}
	    });
	}
	//调转到签到日历页面
	$("#goCalendar").on("touchstart",function(){
		window.location.href="/signh5/getsignin?Token="+userToken+"&PackageId="+packageId;
	});
	//日历页面点击消失
	$("#calendar").on("click",function(){
		$(this).hide();
	});
	//跳转到我的任务页面
	$("#goEarnPeas").on("click",function(){
		window.postMessage("MD1",'*');
//		window.location.href="/signh5/gettask?Token="+userToken+"&PackageId="+packageId;
	});
	//跳转到融豆商城
	$("#goShoppingMall").on("touchstart",function(){
		window.location.href="/nobaseexchange/showscoreexchangeproducts?Token="+userToken+"&PackageId="+packageId;
	});
});