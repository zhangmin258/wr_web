$(function(){
	var userToken=$('meta[name=appToken]').attr('content');
	var packageId=$('meta[name=packageId]').attr('content');
	window.document.addEventListener('message',function(e){
	    var message = e.data;
	    if( message=="MD1" ){
	    	window.location.href="/signh5/gettask?Token="+userToken+"&PackageId="+packageId;
	    }
	    if( message=="SignCalendar" ){
	    	window.location.href="/signh5/getsignin?Token="+userToken+"&PackageId="+packageId;
	    }
	});
	//签到提醒
	var isRemind;
	$('#goRemind').on("click",function(){
		if( isRemind==false ){
			isRemind=true;
			$(this).attr("src","../../static/img/package8/alertOpen.png");
			tuiter(1);
		}else{
			isRemind=false;
			$(this).attr("src","../../static/img/package8/alertClose.png");
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
				$("#goRemind").attr("src","../../static/img/package8/alertClose.png");
			}else{
				isRemind=true;
				$("#goRemind").attr("src","../../static/img/package8/alertOpen.png");
			}
			if( data.status==1 ){
				$("#signIn").attr("src","../../static/img/package8/signIned.png");
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
								$("#signIn").attr("src","../../static/img/package8/signIned.png");
								queryPeas();
								queryDays();
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
					$("#navBtn div").eq(index).find('img').attr("src","../../static/img/package8/"+data.SignReward[index].SignCount+"CanGet.png");
				}else if( data.SignReward[index].IsReceive==1 ){
					//已领取
					$("#navBtn div").eq(index).find('img').attr("src","../../static/img/package8/"+data.SignReward[index].SignCount+"haveGet.png");
				}else if( data.SignReward[index].IsReceive==2 ){
					//未完成
					$("#navBtn div").eq(index).find('img').attr("src","../../static/img/package8/"+data.SignReward[index].SignCount+"Can'tGet.png");
				}
				$("#navBtn div").eq(index).find('img').attr("name",data.SignReward[index].SignCount);
			});
		}
	});
	//获取累计签到天数
	queryDays();
	function queryDays(){
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
				$("#SignCount").html( data.SignCount );
			}
		});
	}
	//点击融豆领取
	$("#navBtn div").on("click",function(){
		var thisName=$(this).find('img').attr('name');
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
					queryPeas();
					$(that).find('img').attr("src","../../static/img/package8/"+thisName+"haveGet.png");
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
	//调转到签到日历页面
	$("#goCalendar").on("touchstart",function(){
		window.location.href="/signh5/getsignin?Token="+userToken+"&PackageId="+packageId;
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