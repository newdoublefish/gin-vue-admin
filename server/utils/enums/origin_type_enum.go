package utils

type OriginTypeEnum uint

const(
	OriginTypeBuildIn OriginTypeEnum = 0x01 // 内部
	OriginTypeAttendance OriginTypeEnum = 0x02 // 考勤
	OriginTypeErp OriginTypeEnum = 0x03 // erp
)
