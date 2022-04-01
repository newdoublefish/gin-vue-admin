package utils

type StaffStatusEnum uint

const(
	StaffStatusEmployed StaffStatusEnum = 0x01 // 在职
	StaffStatusUnEmployed StaffStatusEnum = 0x02 // 离职
)