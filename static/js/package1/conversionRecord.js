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
		url:"/scoreexchange/showscoreexchangerecord",
		contentType:"application/json;charset=utf-8",
		data:{},
		dataType:"text",
		cache:false,
		success:function(data){
			if( data.ret!=200 ){
				return;
			}
			if( data.showSeRecord==null ){
				window.postMessage("Alert+您没有商品兑换记录！",'*');
				return;
			}
			$.each(data.showSeRecord, function(index){
				var timeStr=data.showSeRecord[index].CreateTime.substring(0,10);
				var strLi="<li>"+
					"<dl>"+
						"<dt><img src='"+data.showSeRecord[index].ImgUrl+"'/></dt>"+
						"<dd>"+data.showSeRecord[index].Title+"</dd>"+
						"<dd>"+data.showSeRecord[index].Pay+"</dd>"+
						"<dd>"+timeStr+"</dd>"+
					"</dl>"+
				"</li>";
				$("#recordList").append(strLi);
			});
			mScroll.refresh();
		}
	});

	
	
	
	


	
});

