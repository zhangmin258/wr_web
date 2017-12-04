$(function(){
	//获取任务信息列表
	$.ajax({
		type:"post",
		url:"/sign/getmission",
		contentType:"application/json;charset=utf-8",
		data:{},
		cache:false,
		dataType:"text",
		success:function(data){
			if( data.ret!=200 ){
				return;
			}
			//成长任务提示
			if( data.PointStatus==1 ){
				$("#PointStatus").show();
			}
			$.each(data.DailyMission,function(index){
				var liStr="<li>"+
					"<dl>"+
						"<dt><img src='"+data.DailyMission[index].ImgUrl+"'/></dt>"+
						"<dd>"+data.DailyMission[index].Title+"</dd>"+
						"<dd>"+data.DailyMission[index].Content+"</dd>"+
					"</dl>"+
					"<p class='comStatic'>"+
						"<b></b>"+
						"<span>"+data.DailyMission[index].ProgressStatus+"</span>/1"+
					"</p>"+
					"<p class='addPeas'>"+
						"<span>+"+data.DailyMission[index].MissionAward+"</span>"+
						"<img src='../../static/img/package2/peasYellow.png' />"+
					"</p>"+
					"<button onclick='getPeas(this,"+data.DailyMission[index].MissionType+","+data.DailyMission[index].Id+")' mystatus='"+data.DailyMission[index].Status+"' gotitle='"+data.DailyMission[index].Title+"' peas='"+data.DailyMission[index].MissionAward+"'></button>"+
				"</li>";
				$("#everyDayTask").append(liStr);
			});
			$.each(data.GrowMission,function(index){
				var liStr="<li>"+
					"<dl>"+
						"<dt><img src='"+data.GrowMission[index].ImgUrl+"'/></dt>"+
						"<dd>"+data.GrowMission[index].Title+"</dd>"+
						"<dd>"+data.GrowMission[index].Content+"</dd>"+
					"</dl>"+
					"<p class='comStatic'>"+
						"<b></b>"+
						"<span>"+data.GrowMission[index].ProgressStatus+"</span>/1"+
					"</p>"+
					"<p class='addPeas'>"+
						"<span>+"+data.GrowMission[index].MissionAward+"</span>"+
						"<img src='../../static/img/package2/peasYellow.png' />"+
					"</p>"+
					"<button onclick='getPeas(this,"+data.GrowMission[index].MissionType+","+data.GrowMission[index].Id+")' mystatus='"+data.GrowMission[index].Status+"' gotitle='"+data.GrowMission[index].Title+"' peas='"+data.GrowMission[index].MissionAward+"'></button>"+
				"</li>";
				$("#growthTask").append(liStr);
			});
			//任务完成状态效果呈现
			for( var i=0; i<$(".comStatic span").length; i++ ){
				if( $(".comStatic span").eq(i).html()==1 ){
					$(".comStatic span").eq(i).parent().find('b').addClass('bActive');
				}
				var getStatus=$(".comStatic span").eq(i).parent().parent().find('button').attr('mystatus');
				if( getStatus==0 ){
					$(".comStatic span").eq(i).parent().parent().find('button').html('立即领取');
					$(".comStatic span").eq(i).parent().parent().find('button').css("background","#d0aa74");
				}else if( getStatus==1 ){
					$(".comStatic span").eq(i).parent().parent().find('button').html('已领取');
					$(".comStatic span").eq(i).parent().parent().find('button').attr("disabled","disabled");
					$(".comStatic span").eq(i).parent().parent().find('button').css("background","#c7c7c7");
				}else if( getStatus==2 ){
					$(".comStatic span").eq(i).parent().parent().find('button').html('去完成');
					$(".comStatic span").eq(i).parent().parent().find('button').css("background","#ff8f00");
				}else if( getStatus==3 ){
					$(".comStatic span").eq(i).parent().parent().find('button').html('完成中');
					$(".comStatic span").eq(i).parent().parent().find('button').attr("disabled","disabled");
					$(".comStatic span").eq(i).parent().parent().find('button').css("background","#ccc");
				}
			}
		}
	});
	//任务tab切换
	$("#tab span").on("click",function(){
		$("#tab span").removeClass("tabActive");
		$(this).addClass("tabActive");
		if( $(this).html()=="每日任务" ){
			$("#everyDayTask").show();
			$("#growthTask").hide();
		}else{
			$("#growthTask").show();
			$("#everyDayTask").hide();
		}
	});
	
});
function getPeas(that,missionType,taskId){
	if( $(that).html()=="去完成" ){
		if( $(that).attr('gotitle')=="移动运营认证" ){
			$.ajax({
				type:"post",
				url:"/sign/getfacecount",
				contentType:"application/json;charset=utf-8",
				data:{},
				cache:false,
				dataType:"text",
				success:function(data){
					if( data.ret!=200 ){
						return;
					}
					if( data.status==0 ){
						window.postMessage("Alert+先进行人脸识别",'*');
					}else{
						window.postMessage("task+"+$(that).attr('gotitle'),'*');
					}
				}
			});
		}else{
			window.postMessage("task+"+$(that).attr('gotitle'),'*');
		}
		return;
	}
	if( missionType==1 ){
		$.ajax({
			type:"post",
			url:"/sign/getdailyreward",
			contentType:"application/json;charset=utf-8",
			data:{MissionId:taskId},
			dataType:"text",
			cache:false,
			success:function(data){
				if( data.ret==200 ){
					window.postMessage("at+"+$(that).attr("peas"),'*');
					$(that).attr("disabled","disabled");
					$(that).html("已领取");
					$(that).css("background","#c7c7c7");
				}else{
					window.postMessage("Alert+"+data.msg,'*');
				}
			}
		});
	}else if( missionType==0 ){
		$.ajax({
			type:"post",
			url:"/sign/getgrowreward",
			contentType:"application/json;charset=utf-8",
			data:{MissionId:taskId},
			dataType:"text",
			cache:false,
			success:function(data){
				if( data.ret==200 ){
					window.postMessage("at+"+$(that).attr("peas"),'*');
					$(that).attr("disabled","disabled");
					$(that).html("已领取");
					$(that).css("background","#c7c7c7");
					if( data.PointStatus==0 ){
						$("#PointStatus").hide();
					}
				}else{
					window.postMessage("Alert+"+data.msg,'*');
				}
			}
		});
	}
}
