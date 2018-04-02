function WorldManager()
{
	var remoteIp = "182.254.152.149"
	//var remoteIp = "192.168.1.55"
	var loginUrl = "http://"+remoteIp+":12345/loginServer/GateServer?uId=";
	var ajax 
	var websocket
	var worldManagerObject = this
	
	var output
	var canvas 
	
	var myCircle				//自身控制对象
	var otherCircleArray = [];   //其它气泡空数组
	var robotCircleArray = [];   //机器气泡空数组
	
	var graphics				//graphics
	var canvasWidth = 800 , canvasHeight = 800
	
	var updateSpeed = 33;				//动画帧速度
	
	var robotCount = 0; 			//机器人数量
	
	function init()
	{
		initLogin()
		initOutput()
		initLister()
		initCanvas()
		initUpdate()
	}
	
	function initUpdate()
	{
		setInterval(update, updateSpeed)
		
		//setInterval(robotMove, 3000);			//5秒移动下机器人位置
		robotMove();
	}
	
	function update()
	{
		graphics.clearRect(0,0, canvasWidth, canvasHeight);    //清空 Canvas
		
		graphics.fillStyle="#ffffff";
		graphics.fillText("other : " + otherCircleArray.length ,10, 10);
		graphics.fillText("robot : " + robotCircleArray.length ,10, 30);
		
		
			
		for (var i = 0; i < otherCircleArray.length; i++)
		{
			//if(!existCirle(robotCircleArray , otherCircleArray[i].getId() ))
			//	otherCircleArray[i].update(graphics);
			if(existCirle(robotCircleArray , otherCircleArray[i].getId() ))
				otherCircleArray.splice(i , 1)
			else
				otherCircleArray[i].update(graphics);
		}
		
		
		for (var i = 0; i < robotCircleArray.length; i++)
			robotCircleArray[i].update(graphics);
		
		if (myCircle != null )
			myCircle.update(graphics);
	}
		
	function initCanvas()
	{
		canvas = document.getElementById("canvas")
		graphics = canvas.getContext('2d')
		
		
		window.addEventListener("resize",function(){
			
			 canvasHeight = 800;
             canvasWidth = 800;
            
		},false);	//自适应窗口
		
		
		//鼠标移动
        canvas.onmousemove = function(e)
		{
            mouseX = e.clientX;
            mouseY = e.clientY;
        }
		
		canvas.onmousedown = function(e)
		{
			x = Math.floor(getCanvasPos(canvas,e).x);
			y = Math.floor(getCanvasPos(canvas,e).y);
		
			if(myCircle!=null)
				myCircle.startMove(x ,y);
			
			//writeToScreen(x +" "+y)
		}
	}
	
	
	function robotMove()
	{
		  for (var i = 0; i < robotCircleArray.length; i++) 
		  {
			  var x = Math.floor(Math.random() * canvasWidth);
              var y = Math.floor(Math.random() * canvasHeight);
			  
              robotCircleArray[i].startMove(x ,y);
          }
	}
	
	function existCirle(array , id )
	{
		if(myCircle!=null)
		{
			if(myCircle.getId()==id)
				result=true;
		}
		
		for (var i = 0; i < array.length; i++) 
		{
			circle = array[i];
			if( circle.getId() == id)
				return true;
		}
		return false;
	}
	
	function initLister()
	{
		//window.addEventListener("load", init, false);  
	}
	
	function initOutput()
	{
		output = document.getElementById("output"); 	
	}
		
	function initLogin()
	{
		if(ajax == null)
			ajax = new MyAjax()
		
		document.getElementById("btnLogin").onclick=function()
		{
			uId = document.getElementById("txtUid").value
		
			ajax.send(loginUrl + uId , function(respText)
			{
				writeToScreen(respText)
				
				var jsonObject = JSON.parse(respText)
				var gateServer = JSON.parse(jsonObject.gateServer)
				
				websocket = new MyWebSocket(remoteIp  , gateServer.port , new MyWebSocketHandle(worldManagerObject , jsonObject , false))
			})
		}
		
		document.getElementById("btnChangeScene").onclick=function()
		{
			sceneId = document.getElementById("txtSceneId").value
			
			if( websocket == null)
				return 
			
			var message = "{\"protocolNo\":100001 , \"desc\":\"doLogin\"  , \"toSceneId\":" + sceneId + ",\"uId\": \""+uId+"\",\"id\": "+id+"}"
			websocket.send(message)
		}
		
		document.getElementById("btnRobot").onclick=function()
		{
			robotNum = document.getElementById("robotNum").value
			robotCount=robotNum;
			initRobot();
		}
		
	}
	
	
	function initRobot()
	{
		for (var i = 0; i < robotCount; i++) 
		{
			var userid=Math.floor(Math.random() * 1000000)
			var robotAjax = new MyAjax()
			
			robotAjax.send(loginUrl + userid , function(respText)
			{
				writeToScreen(respText)
				
				var jsonObject = JSON.parse(respText)
				var gateServer = JSON.parse(jsonObject.gateServer)
				var robotWebsocket = new MyWebSocket(remoteIp  , gateServer.port , new MyWebSocketHandle(worldManagerObject , jsonObject,true))
			});
		}
	}
	
	
	function getCanvasPos(canvas,e)  
	{
		//获取鼠标在canvas上的坐标  
		var rect = canvas.getBoundingClientRect();   
		return {   
		 x: e.clientX - rect.left * (canvas.width / rect.width),  
		 y: e.clientY - rect.top * (canvas.height / rect.height)  
	   };  
	}    
	
	this.Start = function()
	{
		//alert("world Start ");
	}
	
	
	this.createSelfPlay = function(websocket , o)
	{
		var selfPlay = o.selfPlay
		var robot = o.robot
		
		if (robot )
		{
			if(!existCirle(robotCircleArray , o.selfPlay.Id ))			//
				robotCircleArray.push(new Circle(websocket ,selfPlay.X,selfPlay.Y ,10,selfPlay.Id,true,"red"));
		}
		else
		{
			myCircle = new Circle(websocket , selfPlay.X , selfPlay.Y , 20,selfPlay.Id,false,"pink")
			
			if(otherCircleArray==null)
				otherCircleArray=[];
			
			if(otherCircleArray.length>0)
				otherCircleArray=[];
			
			
			if ( o.otherPlayList != null)
			{
				for (var i = 0; i < o.otherPlayList.length; i++) 
				{
					cirle = o.otherPlayList[i]
					if(!existCirle(otherCircleArray , cirle.Id ))
						otherCircleArray.push(new Circle(null ,cirle.X, cirle.Y ,10, cirle.Id,false,"red"));
				}
			}
		}
	}
		
	this.newAddUser = function(o)
	{
		id = o.Id
		
		if(!existCirle(otherCircleArray , id ))
			otherCircleArray.push(new Circle(null , o.X,o.Y ,10 , o.Id,false,"red"));
	}
		
	this.getPlayControl = function(id)
	{
		for (var i = 0; i < otherCircleArray.length; i++) 
		{
			otherCircle = otherCircleArray[i]
			if( otherCircle.getId() == id)
				return otherCircle
		}
		return null
	}
	
	this.exitScenePlay = function (id)
	{
		for (var i = 0; i < otherCircleArray.length; i++) 
		{
			otherCircle = otherCircleArray[i]
			
			if( otherCircle.getId() == id)
			{
				otherCircleArray.splice(i , 1)
				break;
			}
		}
	}
	
	this.writeToScreen = function(message) 
	{ 
        var pre = document.createElement("p"); 
        pre.style.wordWrap = "break-word"; 
        pre.innerHTML = message; output.appendChild(pre); 
	}
 
	init();
}	