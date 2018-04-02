package main

import (
	"fmt"
	. "gameTools/game/tools"
	. "gameTools/game/tools/collect"
	. "gameTools/game/tools/config"
	. "gameTools/game/tools/db/mysql"
	. "gameTools/game/tools/json"
	. "gameTools/game/tools/net/sockets/socket"
	//. "lib/gameTools/game/tools/log"
	"strconv"
	"gameTools/game/tools/db/redis"
	"gameTools/game/tools/net/http"
	"gameTools/game/tools/logger"
	"gameTools/game/tools/net/sockets/websocket"
	"uuid-master"
	"time"
	"runtime"
	"sync"
	"gameTools/game/tools"
	"gameTools/game/tools/sorts"
)

//定义的一个动物类
type Animal struct {
	name       string
	x, y, w, h int
	sex        int
}

func (this *Animal) SetName(name string) string {
	this.name = name
	return this.name
}

func (this Animal) GetName() string {	return this.name}
func (this Animal) Action() {	fmt.Println("Animal Action ", this.name)}
func (this Animal) init() {	fmt.Println("Animal init")}
type Dog struct {	Animal}

func (this *Dog) Action() {
	content := "Dog Action " + this.GetName() //调用父类函数
	fmt.Println(content)
}

type People struct{}

func (p *People) ShowA() {	fmt.Println("showA")}
func (p *People) ShowB() {	fmt.Println("showB")}
type Teacher struct {	People}
func (t *Teacher) ShowB() {	fmt.Println("teacher showB")}

type People1 interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	//arraylist()
	//testMap()
	//testRandom();
	//testSocketServer();
	//testSocketClient()
	//testMysql()
	//testJson()
	//testLogFile()
	//testConfig()
	//testRedisPool()sa
	//testLogger()
	//testError()
	//testJsonObject()
	//testHttpServer()
	//testWebSocket()
	//testDeferCall()
	//testUUID()
	//testForeach()
	//testWaitGroup()
	//testExtend()
	testSorts()
	//testTime()
	//testSpeak()
	//testDeferFunc()

}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	fmt.Println("t : " , t)
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func testDeferFunc() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func testSpeak() {

}

func testTime() {
	fmt.Println(CurrentTimeString())
}

func testSorts() {

	var number int
	number = 2222222222115555
	fmt.Println("number : " , number)


	t1 := new(Test)
	t1.Id = 1

	t2 := new(Test)
	t2.Id = 2

	t3 := new(Test)
	t3.Id = 3

	arrayList := NewArrayList()
	arrayList.Add(t3)
	arrayList.Add(t1)
	arrayList.Add(t2)

	fmt.Println("sort 1 : " , arrayList)
	sorts.Sorts(arrayList , "Id" , "asc")
	fmt.Println("sort 2 : " , arrayList)


}

func testExtend() {

	//dog := new(Dog)
	//dog.SetName("dogs1")
	//dog.Action()

	t := Teacher{}
	t.ShowA()
}

func testWaitGroup() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("testWaitGroup")
}

func testForeach() {

	type student struct {
		Name string
		Age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k,v:=range m{
		println(k,"=>",v.Name)
	}
}

func testUUID() {
	hex := uuid.Rand().Hex()
	fmt.Println("hex : " , hex)
}

func testDeferCall() {
	fmt.Println("testDeferCall 1 ")
	defer func() { fmt.Println("打印前") }()
	fmt.Println("testDeferCall 2 ")
	defer func() { fmt.Println("打印中") }()
	fmt.Println("testDeferCall 3 ")
	defer func() { fmt.Println("打印后") }()
	fmt.Println("testDeferCall 4 ")

	panic("触发异常")
}

func testWebSocket() {
	ws := websocket.NewWebSocketServer(12345,  new(websocket.WebSocketHandle))
	ws.Start()
}

func testJsonObject() {

	jsonObject := NewJSONObject()
	jsonObject.Put("a", "a")
	jsonObject.Put("b", "b")
	jsonObject.Put("c", 11)

	arrayList := NewArrayList()
	for i := 0 ; i < 10 ; i ++{
		arrayList.Add(jsonObject)
	}

	fmt.Println(arrayList)

	/*
	o := NewJSONObject()
	o.Put("a", "a")
	o.Put("b", "b")
	o.Put("c", 11)
	o.Put("i", 2)
	o.Put("d", "这是中文")
	o.Put("obj" , "{\"name\":\"张三\", \"id\":11 }")
	array := make([]int, 10 )
	array[1] = 11
	array[2] = 22

	o.Put("array" , array)

	s := o.ToJSONString()
	fmt.Println(s)

	o = ParseJSONObject(s)
	arrayObj := o.GetArray("array")
	obj := o.GetObject("obj" , new(Test)).(*Test)
	stringOb := o.GetString("d")
	intob := o.GetInt("i")

	fmt.Println(stringOb , " , " , intob , " , " , arrayObj[2] , " , " , obj.Name)
	*/
}

func testErr(index int) {

	//defer Exception(func(err interface{}) {
//		fmt.Println("Exception catch ok testErr" , err)
//	})

	panic("testErr")

}

func testErr1(index int) {

	defer Exception(func(err interface{}) {
		fmt.Println("Exception catch ok testErr1 " , err)
	})
	testErr(1)
}


func testError() {

	/*
	var errstring string
	errstring = GetStackString()

	fmt.Println("defer testError"  , errstring)
	*/

	/*
	Try(func() {
		panic("foo")
	}, func(e interface{}) {
		print("testError" , e)
	})
	*/


	testErr1(1)
	tools.DoCmd(func(text string) {
		fmt.Println("text")
	})


	/*
	value := 111
	zero := 0
	value = value / zero
	fmt.Println("execute")
	*/
}


func testLogger(){

	log := logger.NewFileLogger("create")
	for i := 0 ; i < 1000; i++ {
		log.Writer("如果 doParse panic了，复原块会设置返回值为 nil ——延迟函数可以修改带名的返回值。它然后通过断定 err 的赋值是类型 Error 来检查问题出自语法分析。如果不是，类型断言会失败，导致一个运行态错误，继续堆栈退绕，就好像无事发生一样。这个检查意味着如果未曾预料的事情发生了，例如数组下标越界，代码会失败，尽管我们用了panic 和 recover 出来用户触发的错误。有了这种出错处理，error 方法能轻易的报错，而不需担心自己动手退绕堆栈。种有用的模式只应在一个包的内部使用。 Parse 将其内部的 panic 调用转为 os.Error 值；不把 panic 暴露给客户。这个好规则值得效法。" + strconv.Itoa(i))
	}

	log1 := logger.NewFileLogger("longin")
	for i := 0 ; i < 1000; i++ {
		log1.Writer("如果 doParse panic了，复原块会设置返回值为 nil ——延迟函数可以修改带名的返回值。它然后通过断定 err 的赋值是类型 Error 来检查问题出自语法分析。如果不是，类型断言会失败，导致一个运行态错误，继续堆栈退绕，就好像无事发生一样。这个检查意味着如果未曾预料的事情发生了，例如数组下标越界，代码会失败，尽管我们用了panic 和 recover 出来用户触发的错误。有了这种出错处理，error 方法能轻易的报错，而不需担心自己动手退绕堆栈。种有用的模式只应在一个包的内部使用。 Parse 将其内部的 panic 调用转为 os.Error 值；不把 panic 暴露给客户。这个好规则值得效法。"+ strconv.Itoa(i))
	}
}

func testHttpServer() {
	listValue := NewArrayList()
	listValue.Add(new(http.HttpHandle))

	httpServer := http.NewHttpServer("goHttp" , 12345)
	httpServer.RegisterHandles(listValue)

	httpServer.Start()
}

func testRedisPool() {
	redisPool := redis.NewRedisPool("182.254.152.149:6379" , "test_gamechef_redis_*QWERTasdfzxcv%$@" )

	result := redisPool.Do("hgetall","playInfo-10010000000007")
	fmt.Println("testRedisPool result " , result)

	result = redisPool.Do("hmget","playInfo-10010000000016" , "name" , "headIco")
	fmt.Println("testRedisPool result " , result)
}

func testConfig() {
	config := NewConfig("")
	//config.GetValue()

	fmt.Println(config)
}

func testJson() {
	jsonString := "{\"name\":\"张在\", \"id\":15}"

	test := JSONParseObject(jsonString, new(Test)).(*Test)

	jsonObject := JSONParseJson(jsonString)
	jsonName := jsonObject.GetValue("name")

	fmt.Printf("jsonName = %s , test.name = %s \n", jsonName, test.Name)

	arrayList := NewArrayList()

	arrayList.Add(test)
	arrayList.Add(test)

	t := arrayList.Get(0)

	jsonName = t.(*Test).Name

	test1 := new(Test)
	test1.Id = 123
	test1.Name = "苛"

	arrayList1 := NewArrayList()
	arrayList1.Add(test1)

	test2 := test1.Clone().(*Test)
	test2.Name = "在"

	arrayList1.Add(test2)

	arrayListString := JSONParseString(arrayList1.Value())

	fmt.Printf("arrayListString %s , %s \n ", arrayListString, jsonName)

	arrayList = JSONParseArray(arrayListString, new(Test))
	fmt.Println("arrayList = ", arrayList)
}

func testMysql() {

	mysql := NewMysql("192.168.1.131", "3306", "innertest01", "innertest01", "doraemon_logic_1001")
	//test := new(mysql2.Test)
	test := &Test{5, "a1sa55555"}
	test1 := &Test{6, "a4sa"}

	//mysql.InsertObject(test )

	list := new(ArrayList)
	list.Add(test)
	list.Add(test1)
	list.Add(test1)
	list.Add(test1)

	mysql.InsertObject(list.Value()...)
	//fmt.Println(list.String())
	type e = Entity

	queryList := mysql.Query("select * from Test where `name` = ? or `name` = ? or `name` = ?", &Entity{}, "a1sa", 1, 3)
	o := queryList.Get(0)
	fmt.Printf("entity.id = %d \n ", o.(*Entity).Id)
	fmt.Println(queryList)

	//mysql.UpdateObject(test)

	row := mysql.Execute("delete from Test where id = ?" , 1)
	fmt.Println(" row " , row)

	//reflect.New()

}

func testSocketClient() {
	client := NewSocketClient("127.0.0.1:22345", NewStringDeEncode(), func(conn *Connection, msgObject interface{}) {
		fmt.Println("testSocketClient msgObject = ", msgObject)
	})



	for i := 0; i < 2; i++ {

		o := NewJSONObject()
		o.Put("a", "a")
		o.Put("b", "b")
		o.Put("c", 11)
		o.Put("i", i)
		o.Put("d", "这是中文")

	//	data := []byte("hello world 这是中文啊 i = " + strconv.Itoa(i))

		data := []byte(o.ToJSONString())
		dataLength := len(data)

		buffer := make([]byte, 0)
		buffer = append(buffer, IntToBytes(dataLength)...)
		buffer = append(buffer, data...)

		client.Send(buffer)

		fmt.Println("i = ", i, " dataLength = ", dataLength, " buffer = ", buffer)
		//time.Sleep(0.1 * 1e9)

	}

	//client.Close()

	fmt.Println(" sleep 5 sce")
	time.Sleep(50 * 1e9)
}

func testSocketServer() {

	server := NewSocketServer(2451, NewStringDeEncode(), new(SocketHandle))

	server.Start()
}

func testRandom() {
	for i := 0; i < 30; i++ {
		no := ToolsRandom(0, 10)
		fmt.Println(no)
	}
}

func testMap() {
	localMap := new(Map)
	localMap.Put("zzb", 23123)
	localMap.Put("a", "a")
	localMap.Put("b", "b")
	localMap.Put(111, 222)

	localMap.Print()

	localMap.Remove("a")

	localMap.Print()

	for k, v := range localMap.Value() {
		fmt.Println(k, v)
	}

	fmt.Println("======================")
	localMap.Pool(func(key interface{}, value interface{}) {
		fmt.Println(key, value)
	})

}

func arraylist() {
	list := new(ArrayList)
	list1 := new(ArrayList)
	list.Add("2222")
	list.Add("dddd")
	list.Add(45)
	list1.Add("sssss")

	list.Print()
	fmt.Println("**********************")
	list1.Print()

	list.Remove(1)
	list.Print()

	object := list.Get(1)

	fmt.Println(object)
	fmt.Println("**********************")
	for e, v := range list.Value() {
		fmt.Print(e)
		fmt.Print("   ")
		fmt.Println(v)
	}

	fmt.Println("======================")
	list.Pool(func(index interface{}, value interface{}) {
		//fmt.Printf("Pool index %s value %s " , index ,value)
		fmt.Println(value)
	})

	test := new(Test)
	test.Name = "王"
	test.Id = 1111

	objectList := new(ArrayList)
	objectList.Add(test)
	objectList.Add(test)

	t := objectList.Get(0)
	fmt.Println(t.(*Test).Name)
}
