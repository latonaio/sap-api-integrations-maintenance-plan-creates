package sap_api_output_formatter

type MaintenancePlan struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaintenancePlan      string `json:"maintenance_plan_code"`
	Deleted       string `json:"deleted"`
}

type Header struct {
   MaintenancePlan                string      `json:"MaintenancePlan"`
   MaintenancePlanDesc            string      `json:"MaintenancePlanDesc"`
   CreationDate                   string      `json:"CreationDate"`
   LastChangeDate                 string      `json:"LastChangeDate"`
   MaintenanceStrategy            string      `json:"MaintenanceStrategy"`
   SchedulingDuration             string      `json:"SchedulingDuration"`
   SchedulingDurationUnit         string      `json:"SchedulingDurationUnit"`
   NumberOfMaintenanceItems       string      `json:"NumberOfMaintenanceItems"`
   CycleModificationRatio         string      `json:"CycleModificationRatio"`
   MaintPlanSchedgIndicator       string      `json:"MaintPlanSchedgIndicator"`
   AuthorizationGroup             string      `json:"AuthorizationGroup"`
   MaintenancePlanInternalID      string      `json:"MaintenancePlanInternalID"`
   MaintenanceCall                int         `json:"MaintenanceCall"`
   MaintenancePlanCategory        string      `json:"MaintenancePlanCategory"`
   MaintPlanFreeDefinedAttrib     string      `json:"MaintPlanFreeDefinedAttrib"`
   BasicStartDate                 string      `json:"BasicStartDate"`
   SchedulingStartDate            string      `json:"SchedulingStartDate"`
   SchedulingStartTime            string      `json:"SchedulingStartTime"`
   MaintPlanStartCntrReadingValue string      `json:"MaintPlanStartCntrReadingValue"`
   MaintPlnStrtBufDurationInDays  string      `json:"MaintPlnStrtBufDurationInDays"`
   MaintPlanStartBufferUnit       string      `json:"MaintPlanStartBufferUnit"`
   FactoryCalendar                string      `json:"FactoryCalendar"`
   LateCompletionShiftInPercent   string      `json:"LateCompletionShiftInPercent"`
   LateCompletionTolerancePercent string      `json:"LateCompletionTolerancePercent"`
   EarlyCompletionShiftInPercent  string      `json:"EarlyCompletionShiftInPercent"`
   EarlyCompletionTolerancePct    string      `json:"EarlyCompletionTolerancePct"`
   MaintPlanLogicalOperatorCode   string      `json:"MaintPlanLogicalOperatorCode"`
   SchedulingEndDate              string      `json:"SchedulingEndDate"`
   MaintPlanEndCntrReadingValue   string      `json:"MaintPlanEndCntrReadingValue"`
   LastChangeDateTime             string      `json:"LastChangeDateTime"`
   MultipleCounterPlanShiftFactor string      `json:"MultipleCounterPlanShiftFactor"`
   MaintenanceLeadFloatInDays     string      `json:"MaintenanceLeadFloatInDays"`
   MaintenancePlanCallObject      string      `json:"MaintenancePlanCallObject"`
   MaintenancePlanSystemStatus    string      `json:"MaintenancePlanSystemStatus"`
}