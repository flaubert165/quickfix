//Package crossordercancelreplacerequest msg type = t.
package crossordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix50sp2/displayinstruction"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/peginstructions"
	"github.com/quickfixgo/quickfix/fix50sp2/rootparties"
	"github.com/quickfixgo/quickfix/fix50sp2/sidecrossordmodgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/strategyparametersgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/triggeringinstruction"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a CrossOrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"t"`
	fixt11.Header
	//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//CrossID is a required field for CrossOrderCancelReplaceRequest.
	CrossID string `fix:"548"`
	//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
	OrigCrossID string `fix:"551"`
	//CrossType is a required field for CrossOrderCancelReplaceRequest.
	CrossType int `fix:"549"`
	//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
	CrossPrioritization int `fix:"550"`
	//SideCrossOrdModGrp Component
	sidecrossordmodgrp.SideCrossOrdModGrp
	//Instrument Component
	instrument.Instrument
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
	//SettlType is a non-required field for CrossOrderCancelReplaceRequest.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for CrossOrderCancelReplaceRequest.
	SettlDate *string `fix:"64"`
	//HandlInst is a non-required field for CrossOrderCancelReplaceRequest.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//TrdgSesGrp Component
	trdgsesgrp.TrdgSesGrp
	//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
	ProcessCode *string `fix:"81"`
	//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
	PrevClosePx *float64 `fix:"140"`
	//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for CrossOrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//Stipulations Component
	stipulations.Stipulations
	//OrdType is a required field for CrossOrderCancelReplaceRequest.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
	PriceType *int `fix:"423"`
	//Price is a non-required field for CrossOrderCancelReplaceRequest.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData Component
	yielddata.YieldData
	//Currency is a non-required field for CrossOrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
	ComplianceID *string `fix:"376"`
	//IOIID is a non-required field for CrossOrderCancelReplaceRequest.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
	GTBookingInst *int `fix:"427"`
	//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
	MaxShow *float64 `fix:"210"`
	//PegInstructions Component
	peginstructions.PegInstructions
	//DiscretionInstructions Component
	discretioninstructions.DiscretionInstructions
	//TargetStrategy is a non-required field for CrossOrderCancelReplaceRequest.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for CrossOrderCancelReplaceRequest.
	ParticipationRate *float64 `fix:"849"`
	//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for CrossOrderCancelReplaceRequest.
	Designation *string `fix:"494"`
	//StrategyParametersGrp Component
	strategyparametersgrp.StrategyParametersGrp
	//HostCrossID is a non-required field for CrossOrderCancelReplaceRequest.
	HostCrossID *string `fix:"961"`
	//TransBkdTime is a non-required field for CrossOrderCancelReplaceRequest.
	TransBkdTime *time.Time `fix:"483"`
	//RootParties Component
	rootparties.RootParties
	//MatchIncrement is a non-required field for CrossOrderCancelReplaceRequest.
	MatchIncrement *float64 `fix:"1089"`
	//MaxPriceLevels is a non-required field for CrossOrderCancelReplaceRequest.
	MaxPriceLevels *int `fix:"1090"`
	//DisplayInstruction Component
	displayinstruction.DisplayInstruction
	//PriceProtectionScope is a non-required field for CrossOrderCancelReplaceRequest.
	PriceProtectionScope *string `fix:"1092"`
	//TriggeringInstruction Component
	triggeringinstruction.TriggeringInstruction
	//ExDestinationIDSource is a non-required field for CrossOrderCancelReplaceRequest.
	ExDestinationIDSource *string `fix:"1133"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)                  { m.OrderID = &v }
func (m *Message) SetCrossID(v string)                  { m.CrossID = v }
func (m *Message) SetOrigCrossID(v string)              { m.OrigCrossID = v }
func (m *Message) SetCrossType(v int)                   { m.CrossType = v }
func (m *Message) SetCrossPrioritization(v int)         { m.CrossPrioritization = v }
func (m *Message) SetSettlType(v string)                { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                { m.SettlDate = &v }
func (m *Message) SetHandlInst(v string)                { m.HandlInst = &v }
func (m *Message) SetExecInst(v string)                 { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                  { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)            { m.ExDestination = &v }
func (m *Message) SetProcessCode(v string)              { m.ProcessCode = &v }
func (m *Message) SetPrevClosePx(v float64)             { m.PrevClosePx = &v }
func (m *Message) SetLocateReqd(v bool)                 { m.LocateReqd = &v }
func (m *Message) SetTransactTime(v time.Time)          { m.TransactTime = v }
func (m *Message) SetOrdType(v string)                  { m.OrdType = v }
func (m *Message) SetPriceType(v int)                   { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                   { m.Price = &v }
func (m *Message) SetStopPx(v float64)                  { m.StopPx = &v }
func (m *Message) SetCurrency(v string)                 { m.Currency = &v }
func (m *Message) SetComplianceID(v string)             { m.ComplianceID = &v }
func (m *Message) SetIOIID(v string)                    { m.IOIID = &v }
func (m *Message) SetQuoteID(v string)                  { m.QuoteID = &v }
func (m *Message) SetTimeInForce(v string)              { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)         { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)               { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)            { m.ExpireTime = &v }
func (m *Message) SetGTBookingInst(v int)               { m.GTBookingInst = &v }
func (m *Message) SetMaxShow(v float64)                 { m.MaxShow = &v }
func (m *Message) SetTargetStrategy(v int)              { m.TargetStrategy = &v }
func (m *Message) SetTargetStrategyParameters(v string) { m.TargetStrategyParameters = &v }
func (m *Message) SetParticipationRate(v float64)       { m.ParticipationRate = &v }
func (m *Message) SetCancellationRights(v string)       { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)    { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                 { m.RegistID = &v }
func (m *Message) SetDesignation(v string)              { m.Designation = &v }
func (m *Message) SetHostCrossID(v string)              { m.HostCrossID = &v }
func (m *Message) SetTransBkdTime(v time.Time)          { m.TransBkdTime = &v }
func (m *Message) SetMatchIncrement(v float64)          { m.MatchIncrement = &v }
func (m *Message) SetMaxPriceLevels(v int)              { m.MaxPriceLevels = &v }
func (m *Message) SetPriceProtectionScope(v string)     { m.PriceProtectionScope = &v }
func (m *Message) SetExDestinationIDSource(v string)    { m.ExDestinationIDSource = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "t", r
}
