$(function(){
	//阻止它默认事件
    document.addEventListener("touchmove",function(e){
        e.preventDefault();
    });
    //整个页面滚动
	mainScroll();
	var mScroll;
	function mainScroll(){
	   	mScroll = new IScroll("#container",{
	   	   	preventDefault:true,
	   	   	click:true,
			scrollWheel:true,
			fadeScrollbar:true
	   	})
	}
    $("#container").on("touchmove",function(){
        //下拉刷新
        if(mScroll.y>20){
         	
        }
        //上拉刷新,分页
        if(mScroll.y<mScroll.maxScrollY-20){

        }
    });
	//兑换记录展示
	$.ajax({
		type:"post",
		url:"/scoreexchange/getuserslotteryrecord",
		contentType:"application/json;charset=utf-8",
		data:{},
		cache:false,
		dataType:"text",
		success:function(data){
			if( data.ret!=200 ){
				return;
			}
			if( data.lotteryRecord==null ){
				window.postMessage("Alert+您还没有中奖记录！",'*');
				return;
			}
			$.each(data.lotteryRecord, function(index){
				var timeStr=data.lotteryRecord[index].CreateTime.replace("T"," ").substring(0,19);
				var strLi="<li>"+
						"<b>"+data.lotteryRecord[index].Content+"</b>"+
						"<span>"+timeStr+"</span>"+
					"</li>";
				$("#recordList").append(strLi);
			});
			mScroll.refresh();
		}
	});

	
	
	
	


	
});

