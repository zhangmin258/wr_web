$(function() {
	$("#container li").on("click",function(){
		$(this).toggleClass("divDisplay");
		if( $(this).hasClass("divDisplay") ){
			$(this).find("img").attr("src","../../static/img/package8/turndownBlue.png");
		}else{
			$(this).find("img").attr("src","../../static/img/package8/turnrightBlue.png");
		}
	});
    
});