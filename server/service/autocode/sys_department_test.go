package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode/response"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestSysDepartmentService_GetDepartmentTreeMap(t *testing.T) {
	dsn := "root:Aa@6447985@tcp(10.0.0.69:13306)/qmPlus?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err==nil{
		global.GVA_DB = db
	}else{
		return
	}
	tests := []struct {
		name        string
		wantErr     error
		wantTreeMap map[uint][]response.DepartmentTreeResponse
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SysDpService := &SysDepartmentService{}
			_, departments := SysDpService.GetDepartmentTree()
			println(departments)
			//gotErr, gotTreeMap := SysDpService.GetDepartmentTreeMap()
			//if !reflect.DeepEqual(gotErr, tt.wantErr) {
			//	t.Errorf("GetDepartmentTreeMap() gotErr = %v, want %v", gotErr, tt.wantErr)
			//}
			//if !reflect.DeepEqual(gotTreeMap, tt.wantTreeMap) {
			//	t.Errorf("GetDepartmentTreeMap() gotTreeMap = %v, want %v", gotTreeMap, tt.wantTreeMap)
			//}
		})
	}
}
