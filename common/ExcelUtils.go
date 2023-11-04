package common

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// 导出文件
func ExportFile(rows [][]interface{}, title []interface{}, path, fileName string) {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	_ = file.SetSheetName("Sheet1", "表1")
	writer, _ := file.NewStreamWriter("表1")
	_ = writer.SetRow("A1", title)

	cellIndex := 1
	for _, row := range rows {
		cellIndex = cellIndex + 1
		cell, _ := excelize.CoordinatesToCellName(1, cellIndex)
		fmt.Printf("行：%s, 写入数据:%s \n", cell, row)
		err := writer.SetRow(cell, row)
		if err != nil {
			fmt.Printf("行:%s, 写入行错误:%s \n", cell, err.Error())
			continue
		}
	}
	_ = writer.Flush()
	_ = file.SaveAs(path + fileName)
}
