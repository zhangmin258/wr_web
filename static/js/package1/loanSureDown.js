$(function(){
	$("#linkUrl").attr("href","/rule/counselteams?PackageId="+$('meta[name=packageId]').attr('content'));
	//checkbox样式点击切换
	var isSelected=false;
	$("#checkedImg").on("touchstart",function(){
		if( isSelected==true ){
			$(this).attr("src","../../static/img/package1/checked.png");
			isSelected=false;
			document.getElementById("nowInvite").disabled="disabled";
			$("#nowInvite").addClass("disabledSty");
		}else{
			$(this).attr("src","../../static/img/package1/checkedSelected.png");
			document.getElementById("nowInvite").disabled="";
			isSelected=true;
			$("#nowInvite").removeClass("disabledSty");
		}
	});
	//轮询点击消失
	$("#closeWindow").on("touchstart",function(){
		$("header").hide();
	});
	//立即邀请
	$("#nowInvite").on("click",function(){
		if( !isSelected ){
			alert("请阅读并同意《贷款咨询服务协议》");
			return;
		}
//		alert("专属贷款顾问稍后会跟您电话联系，请保持通讯通畅！");
		window.postMessage("buy+LoanSureDown",'*');
		
	});
	
	
});

