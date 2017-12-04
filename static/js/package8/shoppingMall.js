$(function(){
	var userToken=$('meta[name=appToken]').attr('content');
	var packageId=$('meta[name=packageId]').attr('content');
	window.document.addEventListener('message',function(e){
	    var message = e.data;
	    if( message=="MD2" ){
	    	window.location.href="/signh5/gettask?Token="+userToken+"&PackageId="+packageId;
	    }
	});
	//获取融豆数量
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
			$("#peasCount").html(data.Score);
		}
	});
	//中奖快报轮播信息获取
	var msgIndex=0;
	$.ajax({
		type:"post",
		url:"/scoreexchange/lotteryreport",
		contentType:"application/json;charset=utf-8",
		data:{},
		cache:false,
		dataType:"text",
		success:function(data){
			if( data.ret!=200 ){
				return;
			}
			$.each(data.lotteryMsg, function(index){
				var msgLi="<li>"+data.lotteryMsg[index]+"</li>";
				$("#carousel").append(msgLi);
				msgIndex++;
			});
			//中奖名单轮播
		    var carousel=0;
			var timer=setInterval(function(){
				carousel++;
				if( carousel==msgIndex ){
					carousel=0;
					$("#carousel").css("top","0");
				}
				$("#carousel").animate({"top":-1.22*carousel+"rem"});
			},3000);
		}
	});
	//点击X弹窗消失
	$("#closeAlert").on("touchstart",function(){
		$("#noPeasAlert").hide();
	});
	//点击取消弹窗消失
	$("#cancel").on("click",function(){
		$("#noPeasAlert").hide();
	});
	//跳转到兑换记录页
	$("#recordPage").on("click",function(){
		window.location.href="/nobaseexchange/scoreexchangerecord?Token="+userToken+"&PackageId="+packageId;
	});
	//跳转到融豆抽奖页
	$("#peasCJ").on("click",function(){
		window.location.href="/nobaseexchange/showlotteryproducts?Token="+userToken+"&PackageId="+packageId;
	});
	//获取新口子
	$("#getNewCut").on("click",function(){
		window.postMessage("buy+getNewCut",'*');
	});
	//网贷征信查询
	$("#CreditReportingQueries").on("click",function(){
		window.postMessage("buy+CreditReportingQueries",'*');
	});
	//平台征信查询
	$("#PlatformCreditEnquiry").on("click",function(){
		window.postMessage("buy+PlatformCreditEnquiry",'*');
	});
	//每日任务
	$("#everydayTask").on("click",function(){
		window.postMessage("MD2",'*');
//		window.location.href="/signh5/gettask?Token="+userToken+"&PackageId="+packageId;
	});
	//1元话费券
	$("#yuan1").on("click",function(){
		window.postMessage("buy+yuan1",'*');
	});
	//10元话费券
	$("#yuan10").on("click",function(){
		window.postMessage("buy+yuan10",'*');
	});
});