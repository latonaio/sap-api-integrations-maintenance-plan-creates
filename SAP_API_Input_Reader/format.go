package sap_api_input_reader

import (
	"sap-api-integrations-maintenance-plan-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.MaintenancePlan
	return &requests.Header{
		MaintenancePlan:                data.MaintenancePlan,
		MaintenancePlanDesc:            data.MaintenancePlanDesc,
		CreationDate:                   data.CreationDate,
		LastChangeDate:                 data.LastChangeDate,
		MaintenanceStrategy:            data.MaintenanceStrategy,
		SchedulingDuration:             data.SchedulingDuration,
		SchedulingDurationUnit:         data.SchedulingDurationUnit,
		NumberOfMaintenanceItems:       data.NumberOfMaintenanceItems,
		CycleModificationRatio:         data.CycleModificationRatio,
		MaintPlanSchedgIndicator:       data.MaintPlanSchedgIndicator,
		MaintenancePlanInternalID:      data.MaintenancePlanInternalID,
		MaintenanceCall:                data.MaintenanceCall,
		MaintenancePlanCategory:        data.MaintenancePlanCategory,
		MaintPlanFreeDefinedAttrib:     data.MaintPlanFreeDefinedAttrib,
		BasicStartDate:                 data.BasicStartDate,
		SchedulingStartDate:            data.SchedulingStartDate,
		SchedulingStartTime:            data.SchedulingStartTime,
		MaintPlanStartCntrReadingValue: data.MaintPlanStartCntrReadingValue,
		MaintPlnStrtBufDurationInDays:  data.MaintPlnStrtBufDurationInDays,
		MaintPlanStartBufferUnit:       data.MaintPlanStartBufferUnit,
		FactoryCalendar:                data.FactoryCalendar,
		LateCompletionShiftInPercent:   data.LateCompletionShiftInPercent,
		LateCompletionTolerancePercent: data.LateCompletionTolerancePercent,
		EarlyCompletionShiftInPercent:  data.EarlyCompletionShiftInPercent,
		EarlyCompletionTolerancePct:    data.EarlyCompletionTolerancePct,
		MaintPlanLogicalOperatorCode:   data.MaintPlanLogicalOperatorCode,
		SchedulingEndDate:              data.SchedulingEndDate,
		MaintPlanEndCntrReadingValue:   data.MaintPlanEndCntrReadingValue,
		LastChangeDateTime:             data.LastChangeDateTime,
		MultipleCounterPlanShiftFactor: data.MultipleCounterPlanShiftFactor,
		MaintenanceLeadFloatInDays:     data.MaintenanceLeadFloatInDays,
		MaintenancePlanCallObject:      data.MaintenancePlanCallObject,
		MaintenancePlanSystemStatus:    data.MaintenancePlanSystemStatus,
	}
}