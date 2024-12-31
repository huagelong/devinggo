#### 结构体标签

- excelName 列名
- excelIndex 列序号
- toExcelFormat 列转excel函数名称
- toDataFormat excel转data函数名称
- excelColWidth 列宽度

#### 导出示例
1. 结构体定义
```
type NameStruct struct{
   Name string `excelName:"姓名" excelIndex:"1" excelColWidth:"30"`
   Age string `excelName:"年龄" excelIndex:"3"`
   Sex int `excelName:"性别" excelIndex:"1" toExcelFormat:"ToExcelSexFormat"`
}

func (n NameStruct) ToExcelSexFormat() string{
    if n.Sex == 0 {
      return "女"
    }
   return "男"
}
```
2. 输出
````
func main() {
    //创建数据源
   data := createData()
   //创建导出对象
   export := excel.NewExcelExport("test", NameStruct{})
   //销毁对象
   defer export.Close()
   //导出
   err = export.ExportSmallExcelByStruct(data).WriteInFileName("test.xlsx").Error()
   if err != nil {
      fmt.Println("生成失败", err.Error())
   }
}

func createData() []NameStruct {
   var names []NameStruct
   for i := 0; i < 10; i++ {
      names = append(names, NameStruct{name: "hlr" + strconv.Itoa(i), age: strconv.Itoa(i),Sex: i})
   }
   return names
}

````
3. 也可指定输出模拟分页
````
func main() {
    //创建数据源
   data := createData()
   //创建导出对象
   export := excel.NewExcelExport("test", NameStruct{})
   //销毁对象
   defer export.Close()
   //导出 模拟分页
    for i := 0; i < 10; i++ {
       export.ExportData(result, i*10)
    }
   err = export.WriteInFileName("test.xlsx").Error()
   if err != nil {
      fmt.Println("生成失败", err.Error())
   }
}
````
#### 导入实例
1. 导入数据到结构体
````
func main() {
    //接受数据
    var result []NameStruct
    //创建导入对象
   importFile := excel.NewExcelImportFile("111.xlsx", NameStruct{})
   //对象销毁
   defer importFile.Close()
   
   //数据填充 
   err := importFile.ImportDataToStruct(&result).Error()
    //数据显示
   if err != nil {
      fmt.Println("生成失败", err.Error())
   } else {
      marshal, _ := json.Marshal(result)
      fmt.Println(string(marshal))
   }
}
````
2. 逐行获取数据读取
```
func main() {
    //接受数据
    var result []NameStruct
    //创建导入对象
   importFile := excel.NewExcelImportFile("111.xlsx", NameStruct{})
   //对象销毁
   defer importFile.Close()
   // 方式二 逐行遍历
   err := importFile.ImportRead(func(row NameStruct) {
      fmt.Println(row.Name)
   }).Error()

}
```
#### 定义表头样式
1. 代码示例
```
export := excel.NewExcelExport("test", NameStruct{})
//获取默认表头样式
header := excel.CreateDefaultHeader()
//编辑默认header
// .......

export.SetHeadStyle(header)

```
2. 默认样式对象
```
{
   Border: []excelize.Border{
      {Type: "left", Color: "050505", Style: 1},
      {Type: "top", Color: "050505", Style: 1},
      {Type: "bottom", Color: "050505", Style: 1},
      {Type: "right", Color: "050505", Style: 1},
   },
   Fill: excelize.Fill{Type: "gradient", Color: []string{"#a6a6a6", "#a6a6a6"}, Shading: 1},
   Font: nil,
   Alignment: &excelize.Alignment{
      Horizontal:      "center",
      Indent:          1,
      JustifyLastLine: true,
      ReadingOrder:    0,
      RelativeIndent:  1,
      ShrinkToFit:     false,
      TextRotation:    0,
      Vertical:        "",
      WrapText:        false,
   },
   Protection:    nil,
   NumFmt:        0,
   DecimalPlaces: 0,
   CustomNumFmt:  nil,
   Lang:          "",
   NegRed:        false,
}
```