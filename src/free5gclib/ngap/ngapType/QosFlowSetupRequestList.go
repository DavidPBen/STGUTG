package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowSetupRequestList */
/* QosFlowSetupRequestItem */
type QosFlowSetupRequestList struct {
	List []QosFlowSetupRequestItem `aper:"valueExt,sizeLB:1,sizeUB:64"`
}