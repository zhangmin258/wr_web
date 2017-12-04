$(function(){
	var userToken=$('meta[name=appToken]').attr('content');
	var packageId=$('meta[name=packageId]').attr('content');
	window.document.addEventListener('message',function(e){
	    var message = e.data;
	    if( message=="Record" ){
	    	window.location.href="/nobaseexchange/showlotteryrecord?Token="+userToken+"&PackageId="+packageId;
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
				$("#lotteryWinList").append(msgLi);
				msgIndex++;
			});
			//中奖名单轮播
		    var carousel=0;
			var timer=setInterval(function(){
				carousel++;
				if( carousel==msgIndex ){
					carousel=0;
					$("#lotteryWinList").css("top","0");
				}
				$("#lotteryWinList").animate({"top":-1.12*carousel+"rem"});
			},3000);
		}
	});
	//融豆抽奖
	var canCJ=true;
    $("#Tpointer").click(function(){
    	if( canCJ==false ){
    		return;
    	}
    	canCJ=false;
    	$.ajax({
			type:"post",
			url:"/scoreexchange/lotteryresult",
			contentType:"application/json;charset=utf-8",
			data:{},
			async:false,
			cache:false,
			dataType:"text",
			success:function(data){
				if( data.ret!=200 ){
					if( data.msg=="您的融豆余额不足！" ){
						$("#earnPeas").html('去赚融豆');
						$("#winMsg").html( data.msg );
			            $("#winAlert").show();
					}else{
						window.postMessage("Alert+"+data.msg,'*');
					}
					canCJ=true;
			        return;
				}
				var rag;
				var alertMsg=data.msg;
				switch(data.LotteryId){
					case 1:
						rag=7;
//						alertMsg='100融豆';
						break;
					case 2:
						rag=0;
//						alertMsg='200融豆';
						break;
					case 3:
						rag=6;
//						alertMsg='5元';
						break;
					case 4:
						rag=3;
//						alertMsg='10元';
						break;
					case 5:
						rag=4;
//						alertMsg='1000元';
						break;
					case 6:
						rag=2;
//						alertMsg='苹果7';
						break;
					case 7:
						rag=1;
//						alertMsg='谢谢参与';
						break;
					case 8:
						rag=5;
//						alertMsg='谢谢参与';
						break;
				}
				$("#turntable").rotate({
		          	duration:3000,
		            angle:0,
		            animateTo:5760+22.5+45*rag,
		            easing:function(x,t,b,c,d){ return -c * ((t=t/d-1)*t*t*t - 1) + b; },
		            callback: function () {
		             	setTimeout(function(){
		             		$("#earnPeas").html('查看');
		             		$("#winMsg").html( alertMsg );
		             		$("#winAlert").show();
		             		canCJ=true;
		             	},500);
		            }
		        });
			}
		});
    });
    //去赚融豆或去查看中奖记录
    $("#earnPeas").on("click",function(){
    	if( $(this).html()=="去赚融豆" ){
    		window.location.href="/signhome/getsignhome?Token="+userToken+"&PackageId="+packageId;
    	}else if( $(this).html()=="查看" ){
    		window.location.href="/nobaseexchange/showlotteryrecord?Token="+userToken+"&PackageId="+packageId;
    	}
    });
    //点击X弹窗消失
	$("#closeAlert").on("touchstart",function(){
		$("#winAlert").hide();
	});
	//点击取消弹窗消失
	$("#cancel").on("click",function(){
		$("#winAlert").hide();
	});
	
});