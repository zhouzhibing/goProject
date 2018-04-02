package mysql

import (
	"database/sql"
	"fmt"
	"reflect"
	//"strings"
	//"strings"
	//"game/tools"
	_ "go-sql-driver-mysql"
	"strings"
	"gameTools/game/tools/collect"
	"go-goinggo-mapstructure"
	"gameTools/game/tools/inter"
	"gameTools/game/tools"
)

const CHECK_CONNECTION_GAP_TIME = 7 		//检测数据库连接间隔时间戳，（单位：小时)

type Mysql struct {
	lastExecuteTime int64			//最后一新执行的时间戳
	dataBase *sql.DB
	url string
	port string
	username string
	password string
	dbname string
}

func NewMysql(url string, port string, username string, password string, dbname string) *Mysql {
	this := new(Mysql)
	this.url = url
	this.port = port
	this.username = username
	this.password = password
	this.dbname = dbname

	this.connection(url , port , username , password , dbname)
	return this
}

/**
其中连接参数可以有如下几种形式：
user@unix(/path/to/socket)/dbname?charset=utf8
user:password@tcp(localhost:5555)/dbname?charset=utf8
user:password@/dbname
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
*/
func (this *Mysql) connection(url string, port string, username string, password string, dbname string) {

	if this.dataBase != nil {
		this.dataBase.Close();
		this.dataBase = nil;
		//fmt.Println("Sql database already connection")
		//return
	}

	connectionString := username + ":" + password + "@tcp(" + url + ":" + port + ")/" + dbname + "?charset=utf8"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println("Sql open error " + err.Error())
	}
	this.dataBase = db
	this.dataBase.SetMaxOpenConns(100)
	this.dataBase.SetMaxIdleConns(10)

	this.lastExecuteTime = tools.Millisecond()
}

func (this *Mysql) reConnection() {

	this.connection(this.url , this.port , this.username , this.password , this.dbname)

}

func (this *Mysql) InsertObject(objectArray ...interface{}) {

	this.checkConnection()

	var sqlString string
	var typeElem reflect.Value
	var typeObject reflect.Type
	var fieldString string                   //记录的列名字符串
	var valueString string                   //记录的值字符串
	var tableString string                   //记录的表名称
	var fieldNum int                         //对应对象的属性数量
	var objectArrayLength = len(objectArray) //要插入对象的数组长度
	var valueArrayList collect.ArrayList     //记录要插入的值列表，用于后面参数传入
	var firstRefValueObject reflect.Value    //传入要插入的列表数组中第一个对象的反射对象

	for index, object := range objectArray {

		typeElem = reflect.ValueOf(object).Elem()
		typeObject = typeElem.Type()

		if index == 0 { //如果第一个对象，则首先去确定下列名字符串

			firstRefValueObject = typeElem //记录第一个反射对象
			tableString = typeElem.Type().Name()
			fieldNum = typeElem.NumField()

			for i := 0; i < fieldNum; i++ {
				fieldObject := typeObject.Field(i)
				if equalsId(typeObject.Field(i).Name) {
					continue
				}
				if i < fieldNum-1 {
					fieldString += fieldObject.Name + ","
				} else {
					fieldString += fieldObject.Name
				}
			}
		}

		valueString += "("
		for i := 0; i < fieldNum; i++ {
			if equalsId(typeObject.Field(i).Name) {
				continue
			}
			if i < fieldNum-1 {
				valueString += " ? ,"
				valueArrayList.Add(typeElem.Field(i).Interface())
			} else {
				valueString += " ? "
				valueArrayList.Add(typeElem.Field(i).Interface())
			}
		}

		if index < objectArrayLength-1 {
			valueString += ") ,"
		} else {
			valueString += ")"
		}
	}

	sqlString = "Insert into " + tableString + " (" + fieldString + ") values  " + valueString + " "

	stmt, err := this.dataBase.Prepare(sqlString)
	if err != nil {
		fmt.Println("Sql InsertObject prepare error " + err.Error())
		return
	}

	var res sql.Result

	res, err = stmt.Exec(valueArrayList.Value()...)

	if err != nil {
		fmt.Println("Sql exec error " + err.Error())
		return
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Sql lastInsertId error " + err.Error())
		return
	}

	if objectArrayLength == 1 { //如果只插入一个对象，则把最后插入的ID，设置该对象上
		value := firstRefValueObject.FieldByName("Id")
		value.SetInt(insertId)
	}

	fmt.Println(insertId, sqlString)

	this.lastExecuteTime = tools.Millisecond()
}


/**
* queryString 输出要查询的语句字符串
* cloneObj 传入要克隆实体类对象的实现接口
* params 查询语句字符串的参数列表
* return ArrayList 要返回的传入的克隆实体类对象列表
*/
func (this *Mysql) Query(queryString string, cloneObj  inter.IClone , params ...interface{} ) *collect.ArrayList {

		this.checkConnection()

		stmt, err := this.dataBase.Prepare(queryString)
		if err != nil {
			fmt.Println("Sql Query prepare error " + err.Error())
			tools.Exception(nil)

			this.reConnection()
			return nil
		}
		rows , err := stmt.Query(params...)

		if err != nil {
			fmt.Println("Sql Query Query error " + err.Error())
			tools.Exception(nil)

			this.reConnection()
			return nil
		}

		defer rows.Close()
		columns, err := rows.Columns()
		if err != nil {
			return nil
		}

		resultList := new(collect.ArrayList)

		count := len(columns)
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		for rows.Next() {
			for i := 0; i < count; i++ {
				valuePtrs[i] = &values[i]
			}
			rows.Scan(valuePtrs...)

			rowMap := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]
				b, ok := val.([]byte)
				if ok {
					v = string(b)
				} else {
					v = val
				}
				rowMap[col] = v
			}

			object := cloneObj.Clone()
			mapstructure.Decode(rowMap, object)

			resultList.Add(object)
		}

		this.lastExecuteTime = tools.Millisecond()
		return resultList
}

//比较是否为id字段
func equalsId(fieldName string) bool {

	if strings.EqualFold(fieldName, "ID") {
		return true
	}
	return false
}

func (this *Mysql) UpdateObjects(objectArray ...interface{}) int64{
	return 0
}

func (this *Mysql) UpdateObject(object interface{}) int64{

	this.checkConnection()

	refObject := reflect.ValueOf(object)
	refObjectElem := refObject.Elem()
	refObjectType := refObjectElem.Type()

	tableString := strings.Split(refObjectType.String(), ".")[1]
	fieldNum := refObjectElem.NumField()

	updateString := "Update `" + tableString + "` set "

	var fieldValue interface{}
	var fieldName string
	var idValue int64

	fieldValueList := collect.NewArrayList()

	for i := 0 ; i < fieldNum ; i ++{
		refObjectType.Field(i)
		fieldName = refObjectType.Field(i).Name
		fieldValue = refObjectElem.Field(i).Interface()

		if equalsId(fieldName) {
			idValue =  fieldValue.(int64)
			continue
		}

		if i < fieldNum-1 {
			updateString += fieldName + " = ? ,  "
		} else {
			updateString += fieldName + " = ? "
		}

		fieldValueList.Add(fieldValue)
	}

	updateString += " where id = " + tools.IntToString(idValue)

	stmt, err := this.dataBase.Prepare(updateString)
	if err != nil {
		fmt.Println("Sql UpdateObject.Prepare prepare error " + err.Error() , " updateString : " , updateString)
		this.reConnection()
	}

	res , err := stmt.Exec(fieldValueList.Value()...)

	if err != nil {
		fmt.Println("Sql UpdateObject.Exec Query error " + err.Error())
		this.reConnection()
	}

	num, err := res.RowsAffected()

	if err != nil {
		fmt.Println("Sql UpdateObject.RowsAffected Query error " + err.Error())
	}

	//fmt.Println("updateString = " , updateString , " res " , res)
	//fmt.Println("updateString = " , updateString , "  rows " , rows)
	this.lastExecuteTime = tools.Millisecond()

	return num
}

func (this *Mysql) delete(tableName string, args ... interface{}) int64{

	this.checkConnection()

	var queryString string

	stmt, err := this.dataBase.Prepare(queryString)
	if err != nil {
		fmt.Println("Sql Delete.Prepare prepare error " + err.Error())
	}

	res , err := stmt.Exec(args...)

	if err != nil {
		fmt.Println("Sql Delete.Exec Query error " + err.Error())
	}

	num, err := res.RowsAffected()

	if err != nil {
		fmt.Println("Sql Delete.RowsAffected Query error " + err.Error())
	}

	this.lastExecuteTime = tools.Millisecond()

	return num
}

/*
* 执行传入的sql语句，返回影响的行数 */
func (this *Mysql) Execute(sqlString string , args ... interface{}) int64 {

	this.checkConnection()

	stmt, err := this.dataBase.Prepare(sqlString)

	if err != nil {
		fmt.Println("Sql Execute.Prepare prepare error " + err.Error())
	}

	res , err := stmt.Exec(args...)

	if err != nil {
		fmt.Println("Sql Execute.Exec Query error " + err.Error())
	}

	rowAffect , _:= res.RowsAffected()
	//lastInsert , _ := res.LastInsertId()

	this.lastExecuteTime = tools.Millisecond()

	return rowAffect
}

func (this * Mysql) checkConnection(){

	nowTime := tools.Millisecond()
	lastHourTime := (nowTime - this.lastExecuteTime) / 1000 / 60 / 60

	//fmt.Println("this.lastExecuteTime : " , this.lastExecuteTime , " lastHourTime : " , lastHourTime , " CHECK_CONNECTION_GAP_TIME : " , CHECK_CONNECTION_GAP_TIME)
	if  lastHourTime >= CHECK_CONNECTION_GAP_TIME{
		this.reConnection()
	}
}

//
//func findFieldByIdName(firstRefValueObject reflect.Value)  reflect.Value{
//	value := firstRefValueObject.FieldByName("Id")
//	if value.Pointer() {
//		return value
//	}
//	return nil
//}
