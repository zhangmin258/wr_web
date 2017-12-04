$(function() {
	//获取累计签到天数
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
	//获取签到日历已经签到的日期数组
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
		    var monthFirst = new Date(myDate.getFullYear(), parseInt(myDate.getMonth()), 1).getDay();
		    var d = new Date(myDate.getFullYear(), parseInt(myDate.getMonth() + 1), 0);
		    var totalDay = d.getDate();
		    //生成当月的日历且含已签到
		    var $dateLi = $("#qiandaoList").find("li");
		    for (var i = 0; i < totalDay; i++) {
		        $dateLi.eq(i + monthFirst).find('span').html(parseInt(i + 1));
		        for (var j = 0; j < dateArray.length; j++) {
		            if (i == dateArray[j]-1) {
		                $dateLi.eq(i + monthFirst).find('div').css("display","block");
		                $dateLi.eq(i + monthFirst).find('span').css("color","#fff");
		            }
		        }
		    } 
    	}
    });
    
	

});