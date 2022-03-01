package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-plan-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
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
		// ToStrategyCycle:                data.ToStrategyCycle.ToStrategyCycleResults[0],
		// ToItem:                         data.ToItem.ToItemResults[0],
	}

	return header, nil
}
