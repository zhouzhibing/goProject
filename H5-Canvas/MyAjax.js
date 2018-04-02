function MyAjax(ip , port)
{
	var ajaxObject ;
	
	//创建ajax对象
	function creatAjax()
	{
	  if (window.ActiveXObject) {//如果是IE
		  ajaxObject = new ActiveXObject("Microsoft.XMLHTTP");
	  } else if (window.XMLHttpRequest) {
			ajaxObject = new XMLHttpRequest(); //实例化一个xmlHttpReg
	  }
	  
	  return ajaxObject
	}
	
	this.send = function(url , callback)
	{
		  if (ajaxObject != null) 
		  {
			  ajaxObject.open("get", url, true);
			  ajaxObject.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
			  ajaxObject.send();
			  ajaxObject.onreadystatechange = function(e)
			  {
				  //alert("readyState = " + ajaxObject.readyState  + " status = " + ajaxObject.status);
				  if (ajaxObject.readyState == 4 && ajaxObject.status == 200)
				  {
					  callback(ajaxObject.responseText)
					  //alert(ajaxObject.responseText);
				  }
			  }
		  }
	}
	
	creatAjax();
}	