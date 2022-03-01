package responses

type ToStrategyCycle struct {
	D struct {
			MaintenancePlan                string `json:"MaintenancePlan"`
			MaintenancePlanCycle           string `json:"MaintenancePlanCycle"`
			MaintenanceStrategy            string `json:"MaintenanceStrategy"`
			MaintPlanCycRcrrcIntervalQty   string `json:"MaintPlanCycRcrrcIntervalQty"`
			MaintPlanCycRcrrcIntervalUnit  string `json:"MaintPlanCycRcrrcIntervalUnit"`
			MaintPlanCycleDesc             string `json:"MaintPlanCycleDesc"`
			MaintPlanCycleStartOffsetValue string `json:"MaintPlanCycleStartOffsetValue"`
	} `json:"d"`
}
