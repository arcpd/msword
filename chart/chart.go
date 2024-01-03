package chart

import (
	_d "fmt"
	_a "github.com/arcpd/msword"
	_df "github.com/arcpd/msword/color"
	_gb "github.com/arcpd/msword/drawing"
	_e "github.com/arcpd/msword/measurement"
	_gd "github.com/arcpd/msword/schema/soo/dml"
	_c "github.com/arcpd/msword/schema/soo/dml/chart"
	_g "math/rand"
)

func (_bga SurfaceChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _bga._ccda.Cat == nil {
		_bga._ccda.Cat = _c.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_bga._ccda.Cat)
}

// CategoryAxis returns the category data source.
func (_db AreaChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _db._be.Cat == nil {
		_db._be.Cat = _c.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_db._be.Cat)
}
func (_fca AreaChart) AddAxis(axis Axis) {
	_af := _c.NewCT_UnsignedInt()
	_af.ValAttr = axis.AxisID()
	_fca._gdb.AxId = append(_fca._gdb.AxId, _af)
}

// SetText sets the series text
func (_egg ScatterChartSeries) SetText(s string) {
	_egg._bcb.Tx = _c.NewCT_SerTx()
	_egg._bcb.Tx.Choice.V = &s
}
func (_ebf ScatterChart) InitializeDefaults() {
	_ebf._dfb.ScatterStyle.ValAttr = _c.ST_ScatterStyleMarker
}

// SetIndex sets the index of the series
func (_fbf SurfaceChartSeries) SetIndex(idx uint32) { _fbf._ccda.Idx.ValAttr = idx }

// Values returns the value data source.
func (_cgdc BarChartSeries) Values() NumberDataSource {
	if _cgdc._bge.Val == nil {
		_cgdc._bge.Val = _c.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_cgdc._bge.Val)
}

// X returns the inner wrapped XML type.
func (_afgc RadarChart) X() *_c.CT_RadarChart { return _afgc._ecg }

// AddSurfaceChart adds a new surface chart to a chart.
func (_ebe Chart) AddSurfaceChart() SurfaceChart {
	_aea := _c.NewCT_PlotAreaChoice()
	_ebe._cba.Chart.PlotArea.Choice = append(_ebe._cba.Chart.PlotArea.Choice, _aea)
	_aea.SurfaceChart = _c.NewCT_SurfaceChart()
	_ce(_ebe._cba.Chart)
	_ebe._cba.Chart.View3D.RotX.ValAttr = _a.Int8(90)
	_ebe._cba.Chart.View3D.RotY.ValAttr = _a.Uint16(0)
	_ebe._cba.Chart.View3D.Perspective = _c.NewCT_Perspective()
	_ebe._cba.Chart.View3D.Perspective.ValAttr = _a.Uint8(0)
	_eeab := SurfaceChart{_afcaf: _aea.SurfaceChart}
	_eeab.InitializeDefaults()
	return _eeab
}

// AddSeries adds a default series to an Pie chart.
func (_aed PieOfPieChart) AddSeries() PieChartSeries {
	_cgfc := _c.NewCT_PieSer()
	_aed._ecfb.Ser = append(_aed._ecfb.Ser, _cgfc)
	_cgfc.Idx.ValAttr = uint32(len(_aed._ecfb.Ser) - 1)
	_cgfc.Order.ValAttr = uint32(len(_aed._ecfb.Ser) - 1)
	_gfd := PieChartSeries{_cgfc}
	_gfd.InitializeDefaults()
	return _gfd
}

// InitializeDefaults the bar chart to its defaults
func (_ccgg RadarChart) InitializeDefaults() { _ccgg._ecg.RadarStyle.ValAttr = _c.ST_RadarStyleMarker }

// X returns the inner wrapped XML type.
func (_afgg DoughnutChart) X() *_c.CT_DoughnutChart { return _afgg._ggc }

// AddSeries adds a default series to a Surface chart.
func (_faf Surface3DChart) AddSeries() SurfaceChartSeries {
	_ffc := _faf.nextColor(len(_faf._bfe.Ser))
	_cdd := _c.NewCT_SurfaceSer()
	_faf._bfe.Ser = append(_faf._bfe.Ser, _cdd)
	_cdd.Idx.ValAttr = uint32(len(_faf._bfe.Ser) - 1)
	_cdd.Order.ValAttr = uint32(len(_faf._bfe.Ser) - 1)
	_ceb := SurfaceChartSeries{_cdd}
	_ceb.InitializeDefaults()
	_ceb.Properties().LineProperties().SetSolidFill(_ffc)
	return _ceb
}
func (_ecc DateAxis) SetTickLabelPosition(p _c.ST_TickLblPos) {
	if p == _c.ST_TickLblPosUnset {
		_ecc._dgc.TickLblPos = nil
	} else {
		_ecc._dgc.TickLblPos = _c.NewCT_TickLblPos()
		_ecc._dgc.TickLblPos.ValAttr = p
	}
}

// AreaChart is an area chart that has a shaded area underneath a curve.
type AreaChart struct {
	chartBase
	_gdb *_c.CT_AreaChart
}

// InitializeDefaults the Stock chart to its defaults
func (_cfeg StockChart) InitializeDefaults() {
	_cfeg._egee.HiLowLines = _c.NewCT_ChartLines()
	_cfeg._egee.UpDownBars = _c.NewCT_UpDownBars()
	_cfeg._egee.UpDownBars.GapWidth = _c.NewCT_GapAmount()
	_cfeg._egee.UpDownBars.GapWidth.ValAttr = &_c.ST_GapAmount{}
	_cfeg._egee.UpDownBars.GapWidth.ValAttr.ST_GapAmountUShort = _a.Uint16(150)
	_cfeg._egee.UpDownBars.UpBars = _c.NewCT_UpDownBar()
	_cfeg._egee.UpDownBars.DownBars = _c.NewCT_UpDownBar()
}
func (_dfc ValueAxis) SetMajorTickMark(m _c.ST_TickMark) {
	if m == _c.ST_TickMarkUnset {
		_dfc._gec.MajorTickMark = nil
	} else {
		_dfc._gec.MajorTickMark = _c.NewCT_TickMark()
		_dfc._gec.MajorTickMark.ValAttr = m
	}
}
func (_ffg DataLabels) SetShowLegendKey(b bool) {
	_ffg.ensureChoice()
	_ffg._ecb.Choice.ShowLegendKey = _c.NewCT_Boolean()
	_ffg._ecb.Choice.ShowLegendKey.ValAttr = _a.Bool(b)
}

type LineChart struct {
	chartBase
	_fdf *_c.CT_LineChart
}

// Properties returns the bar chart series shape properties.
func (_gbc AreaChartSeries) Properties() _gb.ShapeProperties {
	if _gbc._be.SpPr == nil {
		_gbc._be.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_gbc._be.SpPr)
}

type ValueAxis struct{ _gec *_c.CT_ValAx }

func (_caa CategoryAxis) MajorGridLines() GridLines {
	if _caa._ac.MajorGridlines == nil {
		_caa._ac.MajorGridlines = _c.NewCT_ChartLines()
	}
	return GridLines{_caa._ac.MajorGridlines}
}
func (_gdac ValueAxis) SetMinorTickMark(m _c.ST_TickMark) {
	if m == _c.ST_TickMarkUnset {
		_gdac._gec.MinorTickMark = nil
	} else {
		_gdac._gec.MinorTickMark = _c.NewCT_TickMark()
		_gdac._gec.MinorTickMark.ValAttr = m
	}
}
func (_aegd DataLabels) SetShowSeriesName(b bool) {
	_aegd.ensureChoice()
	_aegd._ecb.Choice.ShowSerName = _c.NewCT_Boolean()
	_aegd._ecb.Choice.ShowSerName.ValAttr = _a.Bool(b)
}

// X returns the inner wrapped XML type.
func (_eagg SurfaceChart) X() *_c.CT_SurfaceChart { return _eagg._afcaf }
func (_ggea DateAxis) SetCrosses(axis Axis)       { _ggea._dgc.CrossAx.ValAttr = axis.AxisID() }
func (_dc CategoryAxis) AxisID() uint32           { return _dc._ac.AxId.ValAttr }

// X returns the inner wrapped XML type.
func (_adbf Surface3DChart) X() *_c.CT_Surface3DChart { return _adbf._bfe }

// X returns the inner wrapped XML type.
func (_dbc LineChartSeries) X() *_c.CT_LineSer { return _dbc._cegd }

// Order returns the order of the series
func (_bb LineChartSeries) Order() uint32 { return _bb._cegd.Order.ValAttr }
func (_fgdf Title) RunProperties() _gb.RunProperties {
	if _fgdf._eac.Tx == nil {
		_fgdf.SetText("")
	}
	if _fgdf._eac.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr == nil {
		_fgdf._eac.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr = _gd.NewCT_TextCharacterProperties()
	}
	return _gb.MakeRunProperties(_fgdf._eac.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr)
}
func (_cfa ValueAxis) SetPosition(p _c.ST_AxPos) {
	_cfa._gec.AxPos = _c.NewCT_AxPos()
	_cfa._gec.AxPos.ValAttr = p
}

// SetLabelReference is used to set the source data to a range of cells
// containing strings.
func (_afg CategoryAxisDataSource) SetLabelReference(s string) {
	_afg._fcbg.Choice = _c.NewCT_AxDataSourceChoice()
	_afg._fcbg.Choice.StrRef = _c.NewCT_StrRef()
	_afg._fcbg.Choice.StrRef.F = s
}

// InitializeDefaults the Bubble chart to its defaults
func (_bca BubbleChart) InitializeDefaults() {}

// SetValues sets values directly on a source.
func (_gdf NumberDataSource) SetValues(v []float64) {
	_gdf.ensureChoice()
	_gdf._ggb.Choice.NumRef = nil
	_gdf._ggb.Choice.NumLit = _c.NewCT_NumData()
	_gdf._ggb.Choice.NumLit.PtCount = _c.NewCT_UnsignedInt()
	_gdf._ggb.Choice.NumLit.PtCount.ValAttr = uint32(len(v))
	for _baa, _gdd := range v {
		_gdf._ggb.Choice.NumLit.Pt = append(_gdf._ggb.Choice.NumLit.Pt, &_c.CT_NumVal{IdxAttr: uint32(_baa), V: _d.Sprintf("\u0025\u0067", _gdd)})
	}
}

// SetDirection changes the direction of the bar chart (bar or column).
func (_edc BarChart) SetDirection(d _c.ST_BarDir) { _edc._aad.BarDir.ValAttr = d }
func (_eeg DataLabels) SetShowValue(b bool) {
	_eeg.ensureChoice()
	_eeg._ecb.Choice.ShowVal = _c.NewCT_Boolean()
	_eeg._ecb.Choice.ShowVal.ValAttr = _a.Bool(b)
}

// Properties returns the chart's shape properties.
func (_fcf Chart) Properties() _gb.ShapeProperties {
	if _fcf._cba.SpPr == nil {
		_fcf._cba.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_fcf._cba.SpPr)
}

// AddLine3DChart adds a new 3D line chart to a chart.
func (_fee Chart) AddLine3DChart() Line3DChart {
	_ce(_fee._cba.Chart)
	_gbcd := _c.NewCT_PlotAreaChoice()
	_fee._cba.Chart.PlotArea.Choice = append(_fee._cba.Chart.PlotArea.Choice, _gbcd)
	_gbcd.Line3DChart = _c.NewCT_Line3DChart()
	_gbcd.Line3DChart.Grouping = _c.NewCT_Grouping()
	_gbcd.Line3DChart.Grouping.ValAttr = _c.ST_GroupingStandard
	return Line3DChart{_afcf: _gbcd.Line3DChart}
}

// SetType sets the type the secone pie to either pie or bar
func (_ced PieOfPieChart) SetType(t _c.ST_OfPieType) { _ced._ecfb.OfPieType.ValAttr = t }
func (_cb CategoryAxis) Properties() _gb.ShapeProperties {
	if _cb._ac.SpPr == nil {
		_cb._ac.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_cb._ac.SpPr)
}
func (_abg ValueAxis) Properties() _gb.ShapeProperties {
	if _abg._gec.SpPr == nil {
		_abg._gec.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_abg._gec.SpPr)
}

// BubbleChart is a 2D Bubble chart.
type BubbleChart struct {
	chartBase
	_fcb *_c.CT_BubbleChart
}

// AddLineChart adds a new line chart to a chart.
func (_efa Chart) AddLineChart() LineChart {
	_abb := _c.NewCT_PlotAreaChoice()
	_efa._cba.Chart.PlotArea.Choice = append(_efa._cba.Chart.PlotArea.Choice, _abb)
	_abb.LineChart = _c.NewCT_LineChart()
	_abb.LineChart.Grouping = _c.NewCT_Grouping()
	_abb.LineChart.Grouping.ValAttr = _c.ST_GroupingStandard
	return LineChart{_fdf: _abb.LineChart}
}
func MakeTitle(x *_c.CT_Title) Title { return Title{x} }

// AddSeries adds a default series to a bar chart.
func (_fe BarChart) AddSeries() BarChartSeries {
	_cfb := _fe.nextColor(len(_fe._aad.Ser))
	_gdcg := _c.NewCT_BarSer()
	_fe._aad.Ser = append(_fe._aad.Ser, _gdcg)
	_gdcg.Idx.ValAttr = uint32(len(_fe._aad.Ser) - 1)
	_gdcg.Order.ValAttr = uint32(len(_fe._aad.Ser) - 1)
	_aeb := BarChartSeries{_gdcg}
	_aeb.InitializeDefaults()
	_aeb.Properties().SetSolidFill(_cfb)
	return _aeb
}

// BarChartSeries is a series to be used on a bar chart.
type BarChartSeries struct{ _bge *_c.CT_BarSer }

// RadarChartSeries is a series to be used on an Radar chart.
type RadarChartSeries struct{ _aac *_c.CT_RadarSer }

// PieOfPieChart is a Pie chart with an extra Pie chart.
type PieOfPieChart struct {
	chartBase
	_ecfb *_c.CT_OfPieChart
}

// InitializeDefaults initializes an Pie series to the default values.
func (_caaa PieChartSeries) InitializeDefaults() {}

// InitializeDefaults initializes an Radar series to the default values.
func (_ffe RadarChartSeries) InitializeDefaults() {}

// X returns the inner wrapped XML type.
func (_dfag Title) X() *_c.CT_Title { return _dfag._eac }

// Values returns the bubble size data source.
func (_ef BubbleChartSeries) BubbleSizes() NumberDataSource {
	if _ef._ec.BubbleSize == nil {
		_ef._ec.BubbleSize = _c.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_ef._ec.BubbleSize)
}

// SetText sets the series text.
func (_baeg RadarChartSeries) SetText(s string) {
	_baeg._aac.Tx = _c.NewCT_SerTx()
	_baeg._aac.Tx.Choice.V = &s
}
func (_bd Bar3DChart) AddAxis(axis Axis) {
	_cc := _c.NewCT_UnsignedInt()
	_cc.ValAttr = axis.AxisID()
	_bd._cgd.AxId = append(_bd._cgd.AxId, _cc)
}

// Marker returns the marker properties.
func (_fcfa LineChartSeries) Marker() Marker {
	if _fcfa._cegd.Marker == nil {
		_fcfa._cegd.Marker = _c.NewCT_Marker()
	}
	return MakeMarker(_fcfa._cegd.Marker)
}
func (_ebb LineChartSeries) SetSmooth(b bool) {
	_ebb._cegd.Smooth = _c.NewCT_Boolean()
	_ebb._cegd.Smooth.ValAttr = &b
}

type GridLines struct{ _afca *_c.CT_ChartLines }

// PieChart is a Pie chart.
type PieChart struct {
	chartBase
	_efba *_c.CT_PieChart
}

// X returns the inner wrapped XML type.
func (_gcb AreaChartSeries) X() *_c.CT_AreaSer { return _gcb._be }

type Legend struct{ _gf *_c.CT_Legend }

// SetOrder sets the order of the series
func (_gaca LineChartSeries) SetOrder(idx uint32) { _gaca._cegd.Order.ValAttr = idx }
func (_cgge SurfaceChartSeries) InitializeDefaults() {
	_cgge.Properties().LineProperties().SetWidth(1 * _e.Point)
	_cgge.Properties().LineProperties().SetSolidFill(_df.Black)
	_cgge.Properties().LineProperties().SetJoin(_gb.LineJoinRound)
}
func (_ega ScatterChartSeries) SetSmooth(b bool) {
	_ega._bcb.Smooth = _c.NewCT_Boolean()
	_ega._bcb.Smooth.ValAttr = &b
}

// SetText sets the series text.
func (_fabb PieChartSeries) SetText(s string) {
	_fabb._aab.Tx = _c.NewCT_SerTx()
	_fabb._aab.Tx.Choice.V = &s
}

// SurfaceChart is a 3D surface chart, viewed from the top-down.
type SurfaceChart struct {
	chartBase
	_afcaf *_c.CT_SurfaceChart
}

func (_fbc LineChartSeries) Values() NumberDataSource {
	if _fbc._cegd.Val == nil {
		_fbc._cegd.Val = _c.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_fbc._cegd.Val)
}

// LineChartSeries is the data series for a line chart.
type LineChartSeries struct{ _cegd *_c.CT_LineSer }

// Properties returns the line chart series shape properties.
func (_bee ScatterChartSeries) Properties() _gb.ShapeProperties {
	if _bee._bcb.SpPr == nil {
		_bee._bcb.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_bee._bcb.SpPr)
}

// X returns the inner wrapped XML type.
func (_aeda RadarChartSeries) X() *_c.CT_RadarSer { return _aeda._aac }
func (_agea DataLabels) SetPosition(p _c.ST_DLblPos) {
	_agea.ensureChoice()
	_agea._ecb.Choice.DLblPos = _c.NewCT_DLblPos()
	_agea._ecb.Choice.DLblPos.ValAttr = p
}

// InitializeDefaults initializes a Bubble chart series to the default values.
func (_ag BubbleChartSeries) InitializeDefaults() {}

// SetOrder sets the order of the series
func (_cedf SurfaceChartSeries) SetOrder(idx uint32) { _cedf._ccda.Order.ValAttr = idx }

// PieChartSeries is a series to be used on an Pie chart.
type PieChartSeries struct{ _aab *_c.CT_PieSer }

// X returns the inner wrapped XML type.
func (_ea Bar3DChart) X() *_c.CT_Bar3DChart { return _ea._cgd }

var NullAxis Axis = nullAxis(0)

// Index returns the index of the series
func (_aege SurfaceChartSeries) Index() uint32 { return _aege._ccda.Idx.ValAttr }

// AddRadarChart adds a new radar chart to a chart.
func (_afbb Chart) AddRadarChart() RadarChart {
	_dg := _c.NewCT_PlotAreaChoice()
	_afbb._cba.Chart.PlotArea.Choice = append(_afbb._cba.Chart.PlotArea.Choice, _dg)
	_dg.RadarChart = _c.NewCT_RadarChart()
	_cee := RadarChart{_ecg: _dg.RadarChart}
	_cee.InitializeDefaults()
	return _cee
}

// X returns the inner wrapped XML type.
func (_gc AreaChart) X() *_c.CT_AreaChart { return _gc._gdb }
func (_ege ScatterChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _ege._bcb.XVal == nil {
		_ege._bcb.XVal = _c.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_ege._bcb.XVal)
}
func MakeNumberDataSource(x *_c.CT_NumDataSource) NumberDataSource { return NumberDataSource{x} }
func (_aegc DataLabels) SetShowLeaderLines(b bool) {
	_aegc.ensureChoice()
	_aegc._ecb.Choice.ShowLeaderLines = _c.NewCT_Boolean()
	_aegc._ecb.Choice.ShowLeaderLines.ValAttr = _a.Bool(b)
}

// AddBubbleChart adds a new bubble chart.
func (_de Chart) AddBubbleChart() BubbleChart {
	_dfg := _c.NewCT_PlotAreaChoice()
	_de._cba.Chart.PlotArea.Choice = append(_de._cba.Chart.PlotArea.Choice, _dfg)
	_dfg.BubbleChart = _c.NewCT_BubbleChart()
	_dcc := BubbleChart{_fcb: _dfg.BubbleChart}
	_dcc.InitializeDefaults()
	return _dcc
}
func MakeSeriesAxis(x *_c.CT_SerAx) SeriesAxis { return SeriesAxis{x} }

// BubbleChartSeries is a series to be used on a Bubble chart.
type BubbleChartSeries struct{ _ec *_c.CT_BubbleSer }

// AddSeries adds a default series to a Surface chart.
func (_gee SurfaceChart) AddSeries() SurfaceChartSeries {
	_beb := _gee.nextColor(len(_gee._afcaf.Ser))
	_feda := _c.NewCT_SurfaceSer()
	_gee._afcaf.Ser = append(_gee._afcaf.Ser, _feda)
	_feda.Idx.ValAttr = uint32(len(_gee._afcaf.Ser) - 1)
	_feda.Order.ValAttr = uint32(len(_gee._afcaf.Ser) - 1)
	_acgc := SurfaceChartSeries{_feda}
	_acgc.InitializeDefaults()
	_acgc.Properties().LineProperties().SetSolidFill(_beb)
	return _acgc
}

// Properties returns the bar chart series shape properties.
func (_beg BarChartSeries) Properties() _gb.ShapeProperties {
	if _beg._bge.SpPr == nil {
		_beg._bge.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_beg._bge.SpPr)
}

// Properties returns the bar chart series shape properties.
func (_geb RadarChartSeries) Properties() _gb.ShapeProperties {
	if _geb._aac.SpPr == nil {
		_geb._aac.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_geb._aac.SpPr)
}

// AddPie3DChart adds a new pie chart to a chart.
func (_fga Chart) AddPie3DChart() Pie3DChart {
	_ce(_fga._cba.Chart)
	_ff := _c.NewCT_PlotAreaChoice()
	_fga._cba.Chart.PlotArea.Choice = append(_fga._cba.Chart.PlotArea.Choice, _ff)
	_ff.Pie3DChart = _c.NewCT_Pie3DChart()
	_edb := Pie3DChart{_ebc: _ff.Pie3DChart}
	_edb.InitializeDefaults()
	return _edb
}

// DoughnutChart is a Doughnut chart.
type DoughnutChart struct {
	chartBase
	_ggc *_c.CT_DoughnutChart
}

func (_gac nullAxis) AxisID() uint32 { return 0 }
func (_bbd StockChart) AddAxis(axis Axis) {
	_dbcc := _c.NewCT_UnsignedInt()
	_dbcc.ValAttr = axis.AxisID()
	_bbd._egee.AxId = append(_bbd._egee.AxId, _dbcc)
}
func (_abf CategoryAxis) InitializeDefaults() {
	_abf.SetPosition(_c.ST_AxPosB)
	_abf.SetMajorTickMark(_c.ST_TickMarkOut)
	_abf.SetMinorTickMark(_c.ST_TickMarkIn)
	_abf.SetTickLabelPosition(_c.ST_TickLblPosNextTo)
	_abf.MajorGridLines().Properties().LineProperties().SetSolidFill(_df.LightGray)
	_abf.Properties().LineProperties().SetSolidFill(_df.Black)
}

// AddSeries adds a default series to an Doughnut chart.
func (_cbd DoughnutChart) AddSeries() PieChartSeries {
	_bdd := _c.NewCT_PieSer()
	_cbd._ggc.Ser = append(_cbd._ggc.Ser, _bdd)
	_bdd.Idx.ValAttr = uint32(len(_cbd._ggc.Ser) - 1)
	_bdd.Order.ValAttr = uint32(len(_cbd._ggc.Ser) - 1)
	_cff := PieChartSeries{_bdd}
	_cff.InitializeDefaults()
	return _cff
}

// InitializeDefaults the bar chart to its defaults
func (_baf PieChart) InitializeDefaults() {
	_baf._efba.VaryColors = _c.NewCT_Boolean()
	_baf._efba.VaryColors.ValAttr = _a.Bool(true)
}

// X returns the inner wrapped XML type.
func (_fdb LineChart) X() *_c.CT_LineChart { return _fdb._fdf }

// Index returns the index of the series
func (_gdad LineChartSeries) Index() uint32 { return _gdad._cegd.Idx.ValAttr }

// AddScatterChart adds a scatter (X/Y) chart.
func (_bdab Chart) AddScatterChart() ScatterChart {
	_fea := _c.NewCT_PlotAreaChoice()
	_bdab._cba.Chart.PlotArea.Choice = append(_bdab._cba.Chart.PlotArea.Choice, _fea)
	_fea.ScatterChart = _c.NewCT_ScatterChart()
	_afe := ScatterChart{_dfb: _fea.ScatterChart}
	_afe.InitializeDefaults()
	return _afe
}

// CategoryAxis returns the category data source.
func (_bfa RadarChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _bfa._aac.Cat == nil {
		_bfa._aac.Cat = _c.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_bfa._aac.Cat)
}
func (_cgf Area3DChart) AddAxis(axis Axis) {
	_bc := _c.NewCT_UnsignedInt()
	_bc.ValAttr = axis.AxisID()
	_cgf._ga.AxId = append(_cgf._ga.AxId, _bc)
}
func (_feab DateAxis) AxisID() uint32 { return _feab._dgc.AxId.ValAttr }

// Axis is the interface implemented by different axes when assigning to a
// chart.
type Axis interface{ AxisID() uint32 }

// X returns the inner wrapped XML type.
func (_gbb BarChartSeries) X() *_c.CT_BarSer { return _gbb._bge }
func (_feee LineChartSeries) InitializeDefaults() {
	_feee.Properties().LineProperties().SetWidth(1 * _e.Point)
	_feee.Properties().LineProperties().SetSolidFill(_df.Black)
	_feee.Properties().LineProperties().SetJoin(_gb.LineJoinRound)
	_feee.Marker().SetSymbol(_c.ST_MarkerStyleNone)
	_feee.Labels().SetShowLegendKey(false)
	_feee.Labels().SetShowValue(false)
	_feee.Labels().SetShowPercent(false)
	_feee.Labels().SetShowCategoryName(false)
	_feee.Labels().SetShowSeriesName(false)
	_feee.Labels().SetShowLeaderLines(false)
}
func (_febb Title) SetText(s string) {
	if _febb._eac.Tx == nil {
		_febb._eac.Tx = _c.NewCT_Tx()
	}
	if _febb._eac.Tx.Choice.Rich == nil {
		_febb._eac.Tx.Choice.Rich = _gd.NewCT_TextBody()
	}
	var _gebc *_gd.CT_TextParagraph
	if len(_febb._eac.Tx.Choice.Rich.P) == 0 {
		_gebc = _gd.NewCT_TextParagraph()
		_febb._eac.Tx.Choice.Rich.P = []*_gd.CT_TextParagraph{_gebc}
	} else {
		_gebc = _febb._eac.Tx.Choice.Rich.P[0]
	}
	var _bec *_gd.EG_TextRun
	if len(_gebc.EG_TextRun) == 0 {
		_bec = _gd.NewEG_TextRun()
		_gebc.EG_TextRun = []*_gd.EG_TextRun{_bec}
	} else {
		_bec = _gebc.EG_TextRun[0]
	}
	if _bec.R == nil {
		_bec.R = _gd.NewCT_RegularTextRun()
	}
	_bec.R.T = s
}
func (_eab NumberDataSource) ensureChoice() {
	if _eab._ggb.Choice == nil {
		_eab._ggb.Choice = _c.NewCT_NumDataSourceChoice()
	}
}

// SetHoleSize controls the hole size in the pie chart and is measured in percent.
func (_acfd DoughnutChart) SetHoleSize(pct uint8) {
	if _acfd._ggc.HoleSize == nil {
		_acfd._ggc.HoleSize = _c.NewCT_HoleSize()
	}
	if _acfd._ggc.HoleSize.ValAttr == nil {
		_acfd._ggc.HoleSize.ValAttr = &_c.ST_HoleSize{}
	}
	_acfd._ggc.HoleSize.ValAttr.ST_HoleSizeUByte = &pct
}

// AddCategoryAxis adds a category axis.
func (_cdg Chart) AddCategoryAxis() CategoryAxis {
	_ebd := _c.NewCT_CatAx()
	if _cdg._cba.Chart.PlotArea.CChoice == nil {
		_cdg._cba.Chart.PlotArea.CChoice = _c.NewCT_PlotAreaChoice1()
	}
	_ebd.AxId = _c.NewCT_UnsignedInt()
	_ebd.AxId.ValAttr = 0x7FFFFFFF & _g.Uint32()
	_cdg._cba.Chart.PlotArea.CChoice.CatAx = append(_cdg._cba.Chart.PlotArea.CChoice.CatAx, _ebd)
	_ebd.Auto = _c.NewCT_Boolean()
	_ebd.Auto.ValAttr = _a.Bool(true)
	_ebd.Delete = _c.NewCT_Boolean()
	_ebd.Delete.ValAttr = _a.Bool(false)
	_cfbe := MakeCategoryAxis(_ebd)
	_cfbe.InitializeDefaults()
	return _cfbe
}

type ScatterChart struct {
	chartBase
	_dfb *_c.CT_ScatterChart
}

// X returns the inner wrapped XML type.
func (_bcbf ValueAxis) X() *_c.CT_ValAx { return _bcbf._gec }

// AddStockChart adds a new stock chart.
func (_egf Chart) AddStockChart() StockChart {
	_ada := _c.NewCT_PlotAreaChoice()
	_egf._cba.Chart.PlotArea.Choice = append(_egf._cba.Chart.PlotArea.Choice, _ada)
	_ada.StockChart = _c.NewCT_StockChart()
	_ebg := StockChart{_egee: _ada.StockChart}
	_ebg.InitializeDefaults()
	return _ebg
}

// AddSeries adds a default series to an area chart.
func (_dfd AreaChart) AddSeries() AreaChartSeries {
	_ed := _dfd.nextColor(len(_dfd._gdb.Ser))
	_fc := _c.NewCT_AreaSer()
	_dfd._gdb.Ser = append(_dfd._gdb.Ser, _fc)
	_fc.Idx.ValAttr = uint32(len(_dfd._gdb.Ser) - 1)
	_fc.Order.ValAttr = uint32(len(_dfd._gdb.Ser) - 1)
	_ae := AreaChartSeries{_fc}
	_ae.InitializeDefaults()
	_ae.Properties().SetSolidFill(_ed)
	return _ae
}
func (_aae SurfaceChart) InitializeDefaults() {
	_aae._afcaf.Wireframe = _c.NewCT_Boolean()
	_aae._afcaf.Wireframe.ValAttr = _a.Bool(false)
	_aae._afcaf.BandFmts = _c.NewCT_BandFmts()
	for _acb := 0; _acb < 15; _acb++ {
		_bbc := _c.NewCT_BandFmt()
		_bbc.Idx.ValAttr = uint32(_acb)
		_bbc.SpPr = _gd.NewCT_ShapeProperties()
		_ebfg := _gb.MakeShapeProperties(_bbc.SpPr)
		_ebfg.SetSolidFill(_aae.nextColor(_acb))
		_aae._afcaf.BandFmts.BandFmt = append(_aae._afcaf.BandFmts.BandFmt, _bbc)
	}
}

// X returns the inner wrapped XML type.
func (_ccf PieChartSeries) X() *_c.CT_PieSer { return _ccf._aab }

type DataLabels struct{ _ecb *_c.CT_DLbls }

// AddArea3DChart adds a new area chart to a chart.
func (_adb Chart) AddArea3DChart() Area3DChart {
	_ce(_adb._cba.Chart)
	_efaa := _c.NewCT_PlotAreaChoice()
	_adb._cba.Chart.PlotArea.Choice = append(_adb._cba.Chart.PlotArea.Choice, _efaa)
	_efaa.Area3DChart = _c.NewCT_Area3DChart()
	_fdce := Area3DChart{_ga: _efaa.Area3DChart}
	_fdce.InitializeDefaults()
	return _fdce
}
func (_gdcd ValueAxis) SetTickLabelPosition(p _c.ST_TickLblPos) {
	if p == _c.ST_TickLblPosUnset {
		_gdcd._gec.TickLblPos = nil
	} else {
		_gdcd._gec.TickLblPos = _c.NewCT_TickLblPos()
		_gdcd._gec.TickLblPos.ValAttr = p
	}
}

// AddSeries adds a default series to a Stock chart.
func (_gbec StockChart) AddSeries() LineChartSeries {
	_abag := _c.NewCT_LineSer()
	_gbec._egee.Ser = append(_gbec._egee.Ser, _abag)
	_abag.Idx.ValAttr = uint32(len(_gbec._egee.Ser) - 1)
	_abag.Order.ValAttr = uint32(len(_gbec._egee.Ser) - 1)
	_aedg := LineChartSeries{_abag}
	_aedg.Values().CreateEmptyNumberCache()
	_aedg.Properties().LineProperties().SetNoFill()
	return _aedg
}
func (_agc DateAxis) SetMajorTickMark(m _c.ST_TickMark) {
	if m == _c.ST_TickMarkUnset {
		_agc._dgc.MajorTickMark = nil
	} else {
		_agc._dgc.MajorTickMark = _c.NewCT_TickMark()
		_agc._dgc.MajorTickMark.ValAttr = m
	}
}
func (_fbd SurfaceChartSeries) Values() NumberDataSource {
	if _fbd._ccda.Val == nil {
		_fbd._ccda.Val = _c.NewCT_NumDataSource()
	}
	_bea := MakeNumberDataSource(_fbd._ccda.Val)
	_bea.CreateEmptyNumberCache()
	return _bea
}

// InitializeDefaults the bar chart to its defaults
func (_f AreaChart) InitializeDefaults() {}

// X returns the inner wrapped XML type.
func (_ddd BubbleChart) X() *_c.CT_BubbleChart { return _ddd._fcb }

// SetText sets the series text.
func (_cga BarChartSeries) SetText(s string) {
	_cga._bge.Tx = _c.NewCT_SerTx()
	_cga._bge.Tx.Choice.V = &s
}

// InitializeDefaults initializes an area series to the default values.
func (_fg AreaChartSeries) InitializeDefaults() {}

// SetIndex sets the index of the series
func (_cedd ScatterChartSeries) SetIndex(idx uint32) { _cedd._bcb.Idx.ValAttr = idx }
func (_fae Legend) InitializeDefaults() {
	_fae.SetPosition(_c.ST_LegendPosR)
	_fae.SetOverlay(false)
	_fae.Properties().SetNoFill()
	_fae.Properties().LineProperties().SetNoFill()
}

// X returns the inner wrapped XML type.
func (_bcbc SurfaceChartSeries) X() *_c.CT_SurfaceSer { return _bcbc._ccda }

// Pie3DChart is a Pie3D chart.
type Pie3DChart struct {
	chartBase
	_ebc *_c.CT_Pie3DChart
}

// AddTitle sets a new title on the chart.
func (_cfe Chart) AddTitle() Title {
	_cfe._cba.Chart.Title = _c.NewCT_Title()
	_cfe._cba.Chart.Title.Overlay = _c.NewCT_Boolean()
	_cfe._cba.Chart.Title.Overlay.ValAttr = _a.Bool(false)
	_cfe._cba.Chart.AutoTitleDeleted = _c.NewCT_Boolean()
	_cfe._cba.Chart.AutoTitleDeleted.ValAttr = _a.Bool(false)
	_dacb := MakeTitle(_cfe._cba.Chart.Title)
	_dacb.InitializeDefaults()
	return _dacb
}

type Title struct{ _eac *_c.CT_Title }

func (_fcc Marker) SetSymbol(s _c.ST_MarkerStyle) {
	if s == _c.ST_MarkerStyleUnset {
		_fcc._acc.Symbol = nil
	} else {
		_fcc._acc.Symbol = _c.NewCT_MarkerStyle()
		_fcc._acc.Symbol.ValAttr = s
	}
}

// SetDirection changes the direction of the bar chart (bar or column).
func (_aa Bar3DChart) SetDirection(d _c.ST_BarDir) { _aa._cgd.BarDir.ValAttr = d }
func (_bae Marker) SetSize(sz uint8) {
	_bae._acc.Size = _c.NewCT_MarkerSize()
	_bae._acc.Size.ValAttr = &sz
}

var _fab = []_df.Color{_df.RGB(0x33, 0x66, 0xcc), _df.RGB(0xDC, 0x39, 0x12), _df.RGB(0xFF, 0x99, 0x00), _df.RGB(0x10, 0x96, 0x18), _df.RGB(0x99, 0x00, 0x99), _df.RGB(0x3B, 0x3E, 0xAC), _df.RGB(0x00, 0x99, 0xC6), _df.RGB(0xDD, 0x44, 0x77), _df.RGB(0x66, 0xAA, 0x00), _df.RGB(0xB8, 0x2E, 0x2E), _df.RGB(0x31, 0x63, 0x95), _df.RGB(0x99, 0x44, 0x99), _df.RGB(0x22, 0xAA, 0x99), _df.RGB(0xAA, 0xAA, 0x11), _df.RGB(0x66, 0x33, 0xCC), _df.RGB(0xE6, 0x73, 0x00), _df.RGB(0x8B, 0x07, 0x07), _df.RGB(0x32, 0x92, 0x62), _df.RGB(0x55, 0x74, 0xA6), _df.RGB(0x3B, 0x3E, 0xAC)}

func (_cfbf Legend) Properties() _gb.ShapeProperties {
	if _cfbf._gf.SpPr == nil {
		_cfbf._gf.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_cfbf._gf.SpPr)
}

// AddAxis adds an axis to a Scatter chart.
func (_fec ScatterChart) AddAxis(axis Axis) {
	_adac := _c.NewCT_UnsignedInt()
	_adac.ValAttr = axis.AxisID()
	_fec._dfb.AxId = append(_fec._dfb.AxId, _adac)
}
func (_gbde CategoryAxis) SetMajorTickMark(m _c.ST_TickMark) {
	if m == _c.ST_TickMarkUnset {
		_gbde._ac.MajorTickMark = nil
	} else {
		_gbde._ac.MajorTickMark = _c.NewCT_TickMark()
		_gbde._ac.MajorTickMark.ValAttr = m
	}
}
func (_cecb ValueAxis) AxisID() uint32 { return _cecb._gec.AxId.ValAttr }
func (_eba DataLabels) SetShowPercent(b bool) {
	_eba.ensureChoice()
	_eba._ecb.Choice.ShowPercent = _c.NewCT_Boolean()
	_eba._ecb.Choice.ShowPercent.ValAttr = _a.Bool(b)
}

// X returns the inner wrapped XML type.
func (_bcg Line3DChart) X() *_c.CT_Line3DChart { return _bcg._afcf }
func (_fdba RadarChart) AddAxis(axis Axis) {
	_bccd := _c.NewCT_UnsignedInt()
	_bccd.ValAttr = axis.AxisID()
	_fdba._ecg.AxId = append(_fdba._ecg.AxId, _bccd)
}

// InitializeDefaults initializes a bar chart series to the default values.
func (_gcd BarChartSeries) InitializeDefaults() {}

// Properties returns the bar chart series shape properties.
func (_bddd PieChartSeries) Properties() _gb.ShapeProperties {
	if _bddd._aab.SpPr == nil {
		_bddd._aab.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_bddd._aab.SpPr)
}
func (_ceea GridLines) Properties() _gb.ShapeProperties {
	if _ceea._afca.SpPr == nil {
		_ceea._afca.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_ceea._afca.SpPr)
}
func (_fbg Chart) AddSeriesAxis() SeriesAxis {
	_ebgd := _c.NewCT_SerAx()
	if _fbg._cba.Chart.PlotArea.CChoice == nil {
		_fbg._cba.Chart.PlotArea.CChoice = _c.NewCT_PlotAreaChoice1()
	}
	_ebgd.AxId = _c.NewCT_UnsignedInt()
	_ebgd.AxId.ValAttr = 0x7FFFFFFF & _g.Uint32()
	_fbg._cba.Chart.PlotArea.CChoice.SerAx = append(_fbg._cba.Chart.PlotArea.CChoice.SerAx, _ebgd)
	_ebgd.Delete = _c.NewCT_Boolean()
	_ebgd.Delete.ValAttr = _a.Bool(false)
	_egd := MakeSeriesAxis(_ebgd)
	_egd.InitializeDefaults()
	return _egd
}

// InitializeDefaults the bar chart to its defaults
func (_eb Area3DChart) InitializeDefaults() {}
func MakeLegend(l *_c.CT_Legend) Legend     { return Legend{l} }
func (_ccde SeriesAxis) AxisID() uint32     { return _ccde._gbce.AxId.ValAttr }

// X returns the inner wrapped XML type.
func (_eed Marker) X() *_c.CT_Marker { return _eed._acc }

// InitializeDefaults the bar chart to its defaults
func (_ccd Pie3DChart) InitializeDefaults() {
	_ccd._ebc.VaryColors = _c.NewCT_Boolean()
	_ccd._ebc.VaryColors.ValAttr = _a.Bool(true)
}

// AddAxis adds an axis to a line chart.
func (_eccc Line3DChart) AddAxis(axis Axis) {
	_efb := _c.NewCT_UnsignedInt()
	_efb.ValAttr = axis.AxisID()
	_eccc._afcf.AxId = append(_eccc._afcf.AxId, _efb)
}
func (_gae DataLabels) SetShowCategoryName(b bool) {
	_gae.ensureChoice()
	_gae._ecb.Choice.ShowCatName = _c.NewCT_Boolean()
	_gae._ecb.Choice.ShowCatName.ValAttr = _a.Bool(b)
}

// SetIndex sets the index of the series
func (_bdac LineChartSeries) SetIndex(idx uint32) { _bdac._cegd.Idx.ValAttr = idx }

type nullAxis byte

// X returns the inner wrapped XML type.
func (_bdabc GridLines) X() *_c.CT_ChartLines { return _bdabc._afca }

// X returns the inner wrapped XML type.
func (_egb PieChart) X() *_c.CT_PieChart { return _egb._efba }
func (_acd ValueAxis) MajorGridLines() GridLines {
	if _acd._gec.MajorGridlines == nil {
		_acd._gec.MajorGridlines = _c.NewCT_ChartLines()
	}
	return GridLines{_acd._gec.MajorGridlines}
}

// AddDateAxis adds a value axis to the chart.
func (_eag Chart) AddDateAxis() DateAxis {
	_fdcg := _c.NewCT_DateAx()
	if _eag._cba.Chart.PlotArea.CChoice == nil {
		_eag._cba.Chart.PlotArea.CChoice = _c.NewCT_PlotAreaChoice1()
	}
	_fdcg.AxId = _c.NewCT_UnsignedInt()
	_fdcg.AxId.ValAttr = 0x7FFFFFFF & _g.Uint32()
	_eag._cba.Chart.PlotArea.CChoice.DateAx = append(_eag._cba.Chart.PlotArea.CChoice.DateAx, _fdcg)
	_fdcg.Delete = _c.NewCT_Boolean()
	_fdcg.Delete.ValAttr = _a.Bool(false)
	_fdcg.Scaling = _c.NewCT_Scaling()
	_fdcg.Scaling.Orientation = _c.NewCT_Orientation()
	_fdcg.Scaling.Orientation.ValAttr = _c.ST_OrientationMinMax
	_fdcg.Choice = &_c.EG_AxSharedChoice{}
	_fdcg.Choice.Crosses = _c.NewCT_Crosses()
	_fdcg.Choice.Crosses.ValAttr = _c.ST_CrossesAutoZero
	_cdb := DateAxis{_fdcg}
	_cdb.MajorGridLines().Properties().LineProperties().SetSolidFill(_df.LightGray)
	_cdb.SetMajorTickMark(_c.ST_TickMarkOut)
	_cdb.SetMinorTickMark(_c.ST_TickMarkIn)
	_cdb.SetTickLabelPosition(_c.ST_TickLblPosNextTo)
	_cdb.Properties().LineProperties().SetSolidFill(_df.Black)
	_cdb.SetPosition(_c.ST_AxPosL)
	return _cdb
}

// InitializeDefaults the bar chart to its defaults
func (_eeb PieOfPieChart) InitializeDefaults() {
	_eeb._ecfb.VaryColors = _c.NewCT_Boolean()
	_eeb._ecfb.VaryColors.ValAttr = _a.Bool(true)
	_eeb.SetType(_c.ST_OfPieTypePie)
	_eeb._ecfb.SecondPieSize = _c.NewCT_SecondPieSize()
	_eeb._ecfb.SecondPieSize.ValAttr = &_c.ST_SecondPieSize{}
	_eeb._ecfb.SecondPieSize.ValAttr.ST_SecondPieSizeUShort = _a.Uint16(75)
	_dbcb := _c.NewCT_ChartLines()
	_dbcb.SpPr = _gd.NewCT_ShapeProperties()
	_aebg := _gb.MakeShapeProperties(_dbcb.SpPr)
	_aebg.LineProperties().SetSolidFill(_df.Auto)
	_eeb._ecfb.SerLines = append(_eeb._ecfb.SerLines, _dbcb)
}

// SetOrder sets the order of the series
func (_gbf ScatterChartSeries) SetOrder(idx uint32) { _gbf._bcb.Order.ValAttr = idx }

// AddPieOfPieChart adds a new pie chart to a chart.
func (_acg Chart) AddPieOfPieChart() PieOfPieChart {
	_gbe := _c.NewCT_PlotAreaChoice()
	_acg._cba.Chart.PlotArea.Choice = append(_acg._cba.Chart.PlotArea.Choice, _gbe)
	_gbe.OfPieChart = _c.NewCT_OfPieChart()
	_adbb := PieOfPieChart{_ecfb: _gbe.OfPieChart}
	_adbb.InitializeDefaults()
	return _adbb
}

// X returns the inner wrapped XML type.
func (_efbc ScatterChart) X() *_c.CT_ScatterChart { return _efbc._dfb }

// X returns the inner wrapped XML type.
func (_abe BubbleChartSeries) X() *_c.CT_BubbleSer { return _abe._ec }
func (_eeaf DateAxis) Properties() _gb.ShapeProperties {
	if _eeaf._dgc.SpPr == nil {
		_eeaf._dgc.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_eeaf._dgc.SpPr)
}

// Chart is a generic chart.
type Chart struct{ _cba *_c.ChartSpace }
type CategoryAxis struct{ _ac *_c.CT_CatAx }

func MakeValueAxis(x *_c.CT_ValAx) ValueAxis { return ValueAxis{x} }

// Values returns the value data source.
func (_gda AreaChartSeries) Values() NumberDataSource {
	if _gda._be.Val == nil {
		_gda._be.Val = _c.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_gda._be.Val)
}

type chartBase struct{}
type SurfaceChartSeries struct{ _ccda *_c.CT_SurfaceSer }

// Properties returns the Bubble chart series shape properties.
func (_fa BubbleChartSeries) Properties() _gb.ShapeProperties {
	if _fa._ec.SpPr == nil {
		_fa._ec.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_fa._ec.SpPr)
}

type NumberDataSource struct{ _ggb *_c.CT_NumDataSource }

// AddSeries adds a default series to an Radar chart.
func (_egc RadarChart) AddSeries() RadarChartSeries {
	_bbe := _egc.nextColor(len(_egc._ecg.Ser))
	_gbbf := _c.NewCT_RadarSer()
	_egc._ecg.Ser = append(_egc._ecg.Ser, _gbbf)
	_gbbf.Idx.ValAttr = uint32(len(_egc._ecg.Ser) - 1)
	_gbbf.Order.ValAttr = uint32(len(_egc._ecg.Ser) - 1)
	_bbea := RadarChartSeries{_gbbf}
	_bbea.InitializeDefaults()
	_bbea.Properties().SetSolidFill(_bbe)
	return _bbea
}

// InitializeDefaults the bar chart to its defaults
func (_da Bar3DChart) InitializeDefaults() { _da.SetDirection(_c.ST_BarDirCol) }

// AddSeries adds a default series to a line chart.
func (_dfa Line3DChart) AddSeries() LineChartSeries {
	_ffb := _dfa.nextColor(len(_dfa._afcf.Ser))
	_afbbc := _c.NewCT_LineSer()
	_dfa._afcf.Ser = append(_dfa._afcf.Ser, _afbbc)
	_afbbc.Idx.ValAttr = uint32(len(_dfa._afcf.Ser) - 1)
	_afbbc.Order.ValAttr = uint32(len(_dfa._afcf.Ser) - 1)
	_fcba := LineChartSeries{_afbbc}
	_fcba.InitializeDefaults()
	_fcba.Properties().LineProperties().SetSolidFill(_ffb)
	_fcba.Properties().SetSolidFill(_ffb)
	return _fcba
}

// InitializeDefaults the bar chart to its defaults
func (_fde DoughnutChart) InitializeDefaults() {
	_fde._ggc.VaryColors = _c.NewCT_Boolean()
	_fde._ggc.VaryColors.ValAttr = _a.Bool(true)
	_fde._ggc.HoleSize = _c.NewCT_HoleSize()
	_fde._ggc.HoleSize.ValAttr = &_c.ST_HoleSize{}
	_fde._ggc.HoleSize.ValAttr.ST_HoleSizeUByte = _a.Uint8(50)
}
func (_gaf SeriesAxis) InitializeDefaults() {}

// X returns the inner wrapped XML type.
func (_cgad ScatterChartSeries) X() *_c.CT_ScatterSer { return _cgad._bcb }
func (_edff LineChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _edff._cegd.Cat == nil {
		_edff._cegd.Cat = _c.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_edff._cegd.Cat)
}

// AddDoughnutChart adds a new doughnut (pie with a hole in the center) chart to a chart.
func (_gbbc Chart) AddDoughnutChart() DoughnutChart {
	_ceg := _c.NewCT_PlotAreaChoice()
	_gbbc._cba.Chart.PlotArea.Choice = append(_gbbc._cba.Chart.PlotArea.Choice, _ceg)
	_ceg.DoughnutChart = _c.NewCT_DoughnutChart()
	_gbcf := DoughnutChart{_ggc: _ceg.DoughnutChart}
	_gbcf.InitializeDefaults()
	return _gbcf
}

// Values returns the value data source.
func (_eec PieChartSeries) Values() NumberDataSource {
	if _eec._aab.Val == nil {
		_eec._aab.Val = _c.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_eec._aab.Val)
}
func (_fgc CategoryAxis) SetMinorTickMark(m _c.ST_TickMark) {
	if m == _c.ST_TickMarkUnset {
		_fgc._ac.MinorTickMark = nil
	} else {
		_fgc._ac.MinorTickMark = _c.NewCT_TickMark()
		_fgc._ac.MinorTickMark.ValAttr = m
	}
}

// X returns the inner wrapped XML type.
func (_cafg SeriesAxis) X() *_c.CT_SerAx { return _cafg._gbce }

// Bar3DChart is a 3D bar chart.
type Bar3DChart struct {
	chartBase
	_cgd *_c.CT_Bar3DChart
}

func MakeChart(x *_c.ChartSpace) Chart { return Chart{x} }

// CreateEmptyNumberCache creates an empty number cache, which is used sometimes
// to increase file format compatibility.  It should actually contain the
// computed cell data, but just creating an empty one is good enough.
func (_fef NumberDataSource) CreateEmptyNumberCache() {
	_fef.ensureChoice()
	if _fef._ggb.Choice.NumRef == nil {
		_fef._ggb.Choice.NumRef = _c.NewCT_NumRef()
	}
	_fef._ggb.Choice.NumLit = nil
	_fef._ggb.Choice.NumRef.NumCache = _c.NewCT_NumData()
	_fef._ggb.Choice.NumRef.NumCache.PtCount = _c.NewCT_UnsignedInt()
	_fef._ggb.Choice.NumRef.NumCache.PtCount.ValAttr = 0
}

// X returns the inner wrapped XML type.
func (_cce PieOfPieChart) X() *_c.CT_OfPieChart { return _cce._ecfb }

// AddSeries adds a default series to an area chart.
func (_bg Area3DChart) AddSeries() AreaChartSeries {
	_ad := _bg.nextColor(len(_bg._ga.Ser))
	_gbd := _c.NewCT_AreaSer()
	_bg._ga.Ser = append(_bg._ga.Ser, _gbd)
	_gbd.Idx.ValAttr = uint32(len(_bg._ga.Ser) - 1)
	_gbd.Order.ValAttr = uint32(len(_bg._ga.Ser) - 1)
	_ee := AreaChartSeries{_gbd}
	_ee.InitializeDefaults()
	_ee.Properties().SetSolidFill(_ad)
	return _ee
}

// AddSeries adds a default series to a Bubble chart.
func (_bda BubbleChart) AddSeries() BubbleChartSeries {
	_afd := _bda.nextColor(len(_bda._fcb.Ser))
	_dff := _c.NewCT_BubbleSer()
	_bda._fcb.Ser = append(_bda._fcb.Ser, _dff)
	_dff.Idx.ValAttr = uint32(len(_bda._fcb.Ser) - 1)
	_dff.Order.ValAttr = uint32(len(_bda._fcb.Ser) - 1)
	_fgd := BubbleChartSeries{_dff}
	_fgd.InitializeDefaults()
	_fgd.Properties().SetSolidFill(_afd)
	return _fgd
}

// InitializeDefaults the bar chart to its defaults
func (_cf BarChart) InitializeDefaults() { _cf.SetDirection(_c.ST_BarDirCol) }

// SetText sets the series text
func (_fcad SurfaceChartSeries) SetText(s string) {
	_fcad._ccda.Tx = _c.NewCT_SerTx()
	_fcad._ccda.Tx.Choice.V = &s
}

// Labels returns the data label properties.
func (_dee LineChartSeries) Labels() DataLabels {
	if _dee._cegd.DLbls == nil {
		_dee._cegd.DLbls = _c.NewCT_DLbls()
	}
	return MakeDataLabels(_dee._cegd.DLbls)
}

// Labels returns the data label properties.
func (_bac ScatterChartSeries) Labels() DataLabels {
	if _bac._bcb.DLbls == nil {
		_bac._bcb.DLbls = _c.NewCT_DLbls()
	}
	return MakeDataLabels(_bac._bcb.DLbls)
}
func _ce(_dcd *_c.CT_Chart) {
	_dcd.View3D = _c.NewCT_View3D()
	_dcd.View3D.RotX = _c.NewCT_RotX()
	_dcd.View3D.RotX.ValAttr = _a.Int8(15)
	_dcd.View3D.RotY = _c.NewCT_RotY()
	_dcd.View3D.RotY.ValAttr = _a.Uint16(20)
	_dcd.View3D.RAngAx = _c.NewCT_Boolean()
	_dcd.View3D.RAngAx.ValAttr = _a.Bool(false)
	_dcd.Floor = _c.NewCT_Surface()
	_dcd.Floor.Thickness = _c.NewCT_Thickness()
	_dcd.Floor.Thickness.ValAttr.Uint32 = _a.Uint32(0)
	_dcd.SideWall = _c.NewCT_Surface()
	_dcd.SideWall.Thickness = _c.NewCT_Thickness()
	_dcd.SideWall.Thickness.ValAttr.Uint32 = _a.Uint32(0)
	_dcd.BackWall = _c.NewCT_Surface()
	_dcd.BackWall.Thickness = _c.NewCT_Thickness()
	_dcd.BackWall.Thickness.ValAttr.Uint32 = _a.Uint32(0)
}
func (_bbg NumberDataSource) SetReference(s string) {
	_bbg.ensureChoice()
	if _bbg._ggb.Choice.NumRef == nil {
		_bbg._ggb.Choice.NumRef = _c.NewCT_NumRef()
	}
	_bbg._ggb.Choice.NumRef.F = s
}

// Surface3DChart is a 3D view of a surface chart.
type Surface3DChart struct {
	chartBase
	_bfe *_c.CT_Surface3DChart
}

// AddSeries adds a default series to a bar chart.
func (_cgb Bar3DChart) AddSeries() BarChartSeries {
	_gdae := _cgb.nextColor(len(_cgb._cgd.Ser))
	_ab := _c.NewCT_BarSer()
	_cgb._cgd.Ser = append(_cgb._cgd.Ser, _ab)
	_ab.Idx.ValAttr = uint32(len(_cgb._cgd.Ser) - 1)
	_ab.Order.ValAttr = uint32(len(_cgb._cgd.Ser) - 1)
	_gdc := BarChartSeries{_ab}
	_gdc.InitializeDefaults()
	_gdc.Properties().SetSolidFill(_gdae)
	return _gdc
}

// Values returns the value data source.
func (_cd BubbleChartSeries) Values() NumberDataSource {
	if _cd._ec.YVal == nil {
		_cd._ec.YVal = _c.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_cd._ec.YVal)
}

// AddValueAxis adds a value axis to the chart.
func (_ede Chart) AddValueAxis() ValueAxis {
	_gge := _c.NewCT_ValAx()
	if _ede._cba.Chart.PlotArea.CChoice == nil {
		_ede._cba.Chart.PlotArea.CChoice = _c.NewCT_PlotAreaChoice1()
	}
	_gge.AxId = _c.NewCT_UnsignedInt()
	_gge.AxId.ValAttr = 0x7FFFFFFF & _g.Uint32()
	_ede._cba.Chart.PlotArea.CChoice.ValAx = append(_ede._cba.Chart.PlotArea.CChoice.ValAx, _gge)
	_gge.Delete = _c.NewCT_Boolean()
	_gge.Delete.ValAttr = _a.Bool(false)
	_gge.Scaling = _c.NewCT_Scaling()
	_gge.Scaling.Orientation = _c.NewCT_Orientation()
	_gge.Scaling.Orientation.ValAttr = _c.ST_OrientationMinMax
	_gge.Choice = &_c.EG_AxSharedChoice{}
	_gge.Choice.Crosses = _c.NewCT_Crosses()
	_gge.Choice.Crosses.ValAttr = _c.ST_CrossesAutoZero
	_gge.CrossBetween = _c.NewCT_CrossBetween()
	_gge.CrossBetween.ValAttr = _c.ST_CrossBetweenBetween
	_dcg := MakeValueAxis(_gge)
	_dcg.MajorGridLines().Properties().LineProperties().SetSolidFill(_df.LightGray)
	_dcg.SetMajorTickMark(_c.ST_TickMarkOut)
	_dcg.SetMinorTickMark(_c.ST_TickMarkIn)
	_dcg.SetTickLabelPosition(_c.ST_TickLblPosNextTo)
	_dcg.Properties().LineProperties().SetSolidFill(_df.Black)
	_dcg.SetPosition(_c.ST_AxPosL)
	return _dcg
}

// Index returns the index of the series
func (_eff ScatterChartSeries) Index() uint32 { return _eff._bcb.Idx.ValAttr }

// RemoveLegend removes the legend if the chart has one.
func (_cea Chart) RemoveLegend() { _cea._cba.Chart.Legend = nil }

// AddAreaChart adds a new area chart to a chart.
func (_aba Chart) AddAreaChart() AreaChart {
	_ccg := _c.NewCT_PlotAreaChoice()
	_aba._cba.Chart.PlotArea.Choice = append(_aba._cba.Chart.PlotArea.Choice, _ccg)
	_ccg.AreaChart = _c.NewCT_AreaChart()
	_bdb := AreaChart{_gdb: _ccg.AreaChart}
	_bdb.InitializeDefaults()
	return _bdb
}

// CategoryAxisDataSource specifies the data for an axis.  It's commonly used with
// SetReference to set the axis data to a range of cells.
type CategoryAxisDataSource struct{ _fcbg *_c.CT_AxDataSource }

func (_fed Legend) SetPosition(p _c.ST_LegendPos) {
	if p == _c.ST_LegendPosUnset {
		_fed._gf.LegendPos = nil
	} else {
		_fed._gf.LegendPos = _c.NewCT_LegendPos()
		_fed._gf.LegendPos.ValAttr = p
	}
}

// X returns the inner wrapped XML type.
func (_dd BarChart) X() *_c.CT_BarChart { return _dd._aad }

// Values returns the value data source.
func (_faed RadarChartSeries) Values() NumberDataSource {
	if _faed._aac.Val == nil {
		_faed._aac.Val = _c.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_faed._aac.Val)
}
func (_dabd ValueAxis) SetCrosses(axis Axis)   { _dabd._gec.CrossAx.ValAttr = axis.AxisID() }
func MakeDataLabels(x *_c.CT_DLbls) DataLabels { return DataLabels{x} }
func (_baab SeriesAxis) SetCrosses(axis Axis)  { _baab._gbce.CrossAx.ValAttr = axis.AxisID() }
func MakeMarker(x *_c.CT_Marker) Marker        { return Marker{x} }
func (_cac ScatterChartSeries) Values() NumberDataSource {
	if _cac._bcb.YVal == nil {
		_cac._bcb.YVal = _c.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_cac._bcb.YVal)
}

// CategoryAxis returns the category data source.
func (_bcc BarChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _bcc._bge.Cat == nil {
		_bcc._bge.Cat = _c.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_bcc._bge.Cat)
}

type SeriesAxis struct{ _gbce *_c.CT_SerAx }

// AddPieChart adds a new pie chart to a chart.
func (_gg Chart) AddPieChart() PieChart {
	_fb := _c.NewCT_PlotAreaChoice()
	_gg._cba.Chart.PlotArea.Choice = append(_gg._cba.Chart.PlotArea.Choice, _fb)
	_fb.PieChart = _c.NewCT_PieChart()
	_cgg := PieChart{_efba: _fb.PieChart}
	_cgg.InitializeDefaults()
	return _cgg
}

// X returns the inner wrapped XML type.
func (_dfdf DateAxis) X() *_c.CT_DateAx { return _dfdf._dgc }

// Marker returns the marker properties.
func (_cbb ScatterChartSeries) Marker() Marker {
	if _cbb._bcb.Marker == nil {
		_cbb._bcb.Marker = _c.NewCT_Marker()
	}
	return MakeMarker(_cbb._bcb.Marker)
}

// AddSeries adds a default series to a line chart.
func (_cec LineChart) AddSeries() LineChartSeries {
	_ba := _cec.nextColor(len(_cec._fdf.Ser))
	_gcdc := _c.NewCT_LineSer()
	_cec._fdf.Ser = append(_cec._fdf.Ser, _gcdc)
	_gcdc.Idx.ValAttr = uint32(len(_cec._fdf.Ser) - 1)
	_gcdc.Order.ValAttr = uint32(len(_cec._fdf.Ser) - 1)
	_cda := LineChartSeries{_gcdc}
	_cda.InitializeDefaults()
	_cda.Properties().LineProperties().SetSolidFill(_ba)
	return _cda
}
func (_fdcec DateAxis) MajorGridLines() GridLines {
	if _fdcec._dgc.MajorGridlines == nil {
		_fdcec._dgc.MajorGridlines = _c.NewCT_ChartLines()
	}
	return GridLines{_fdcec._dgc.MajorGridlines}
}

// Properties returns the line chart series shape properties.
func (_add SurfaceChartSeries) Properties() _gb.ShapeProperties {
	if _add._ccda.SpPr == nil {
		_add._ccda.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_add._ccda.SpPr)
}

// AddSurface3DChart adds a new 3D surface chart to a chart.
func (_ggd Chart) AddSurface3DChart() Surface3DChart {
	_dgb := _c.NewCT_PlotAreaChoice()
	_ggd._cba.Chart.PlotArea.Choice = append(_ggd._cba.Chart.PlotArea.Choice, _dgb)
	_dgb.Surface3DChart = _c.NewCT_Surface3DChart()
	_ce(_ggd._cba.Chart)
	_abfd := Surface3DChart{_bfe: _dgb.Surface3DChart}
	_abfd.InitializeDefaults()
	return _abfd
}

type DateAxis struct{ _dgc *_c.CT_DateAx }

func MakeCategoryAxis(x *_c.CT_CatAx) CategoryAxis { return CategoryAxis{x} }
func (_cfff Title) ParagraphProperties() _gb.ParagraphProperties {
	if _cfff._eac.Tx == nil {
		_cfff.SetText("")
	}
	if _cfff._eac.Tx.Choice.Rich.P[0].PPr == nil {
		_cfff._eac.Tx.Choice.Rich.P[0].PPr = _gd.NewCT_TextParagraphProperties()
	}
	return _gb.MakeParagraphProperties(_cfff._eac.Tx.Choice.Rich.P[0].PPr)
}

// Area3DChart is an area chart that has a shaded area underneath a curve.
type Area3DChart struct {
	chartBase
	_ga *_c.CT_Area3DChart
}

// AddSeries adds a default series to an Pie3D chart.
func (_ecca Pie3DChart) AddSeries() PieChartSeries {
	_fbb := _c.NewCT_PieSer()
	_ecca._ebc.Ser = append(_ecca._ebc.Ser, _fbb)
	_fbb.Idx.ValAttr = uint32(len(_ecca._ebc.Ser) - 1)
	_fbb.Order.ValAttr = uint32(len(_ecca._ebc.Ser) - 1)
	_cege := PieChartSeries{_fbb}
	_cege.InitializeDefaults()
	return _cege
}

// RadarChart is an Radar chart that has a shaded Radar underneath a curve.
type RadarChart struct {
	chartBase
	_ecg *_c.CT_RadarChart
}

// StockChart is a 2D Stock chart.
type StockChart struct {
	chartBase
	_egee *_c.CT_StockChart
}

// BarChart is a 2D bar chart.
type BarChart struct {
	chartBase
	_aad *_c.CT_BarChart
}

// CategoryAxis returns the category data source.
func (_gdba BubbleChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _gdba._ec.XVal == nil {
		_gdba._ec.XVal = _c.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_gdba._ec.XVal)
}

// AddAxis adds an axis to a Surface chart.
func (_adbg SurfaceChart) AddAxis(axis Axis) {
	_bgc := _c.NewCT_UnsignedInt()
	_bgc.ValAttr = axis.AxisID()
	_adbg._afcaf.AxId = append(_adbg._afcaf.AxId, _bgc)
}

// AddAxis adds an axis to a line chart.
func (_cggc LineChart) AddAxis(axis Axis) {
	_dcgf := _c.NewCT_UnsignedInt()
	_dcgf.ValAttr = axis.AxisID()
	_cggc._fdf.AxId = append(_cggc._fdf.AxId, _dcgf)
}

// AddLegend adds a legend to a chart, replacing any existing legend.
func (_ceec Chart) AddLegend() Legend {
	_ceec._cba.Chart.Legend = _c.NewCT_Legend()
	_caf := MakeLegend(_ceec._cba.Chart.Legend)
	_caf.InitializeDefaults()
	return _caf
}
func (_acf chartBase) nextColor(_ccb int) _df.Color { return _fab[_ccb%len(_fab)] }

// RemoveTitle removes any existing title from the chart.
func (_afc Chart) RemoveTitle() {
	_afc._cba.Chart.Title = nil
	_afc._cba.Chart.AutoTitleDeleted = _c.NewCT_Boolean()
	_afc._cba.Chart.AutoTitleDeleted.ValAttr = _a.Bool(true)
}

// AddSeries adds a default series to a Scatter chart.
func (_cdc ScatterChart) AddSeries() ScatterChartSeries {
	_ead := _cdc.nextColor(len(_cdc._dfb.Ser))
	_eabd := _c.NewCT_ScatterSer()
	_cdc._dfb.Ser = append(_cdc._dfb.Ser, _eabd)
	_eabd.Idx.ValAttr = uint32(len(_cdc._dfb.Ser) - 1)
	_eabd.Order.ValAttr = uint32(len(_cdc._dfb.Ser) - 1)
	_gca := ScatterChartSeries{_eabd}
	_gca.InitializeDefaults()
	_gca.Marker().Properties().LineProperties().SetSolidFill(_ead)
	_gca.Marker().Properties().SetSolidFill(_ead)
	return _gca
}

// SetText sets the series text.
func (_bf BubbleChartSeries) SetText(s string) {
	_bf._ec.Tx = _c.NewCT_SerTx()
	_bf._ec.Tx.Choice.V = &s
}

// Order returns the order of the series
func (_adbfa SurfaceChartSeries) Order() uint32 { return _adbfa._ccda.Order.ValAttr }

// X returns the inner wrapped XML type.
func (_cg Area3DChart) X() *_c.CT_Area3DChart { return _cg._ga }

// AddBar3DChart adds a new 3D bar chart to a chart.
func (_bce Chart) AddBar3DChart() Bar3DChart {
	_ce(_bce._cba.Chart)
	_bgf := _c.NewCT_PlotAreaChoice()
	_bce._cba.Chart.PlotArea.Choice = append(_bce._cba.Chart.PlotArea.Choice, _bgf)
	_bgf.Bar3DChart = _c.NewCT_Bar3DChart()
	_bgf.Bar3DChart.Grouping = _c.NewCT_BarGrouping()
	_bgf.Bar3DChart.Grouping.ValAttr = _c.ST_BarGroupingStandard
	_aeg := Bar3DChart{_cgd: _bgf.Bar3DChart}
	_aeg.InitializeDefaults()
	return _aeg
}

// ScatterChartSeries is the data series for a scatter chart.
type ScatterChartSeries struct{ _bcb *_c.CT_ScatterSer }

func (_afb BubbleChart) AddAxis(axis Axis) {
	_ge := _c.NewCT_UnsignedInt()
	_ge.ValAttr = axis.AxisID()
	_afb._fcb.AxId = append(_afb._fcb.AxId, _ge)
}

// X returns the inner wrapped XML type.
func (_agd StockChart) X() *_c.CT_StockChart { return _agd._egee }

// AddBarChart adds a new bar chart to a chart.
func (_eaa Chart) AddBarChart() BarChart {
	_feed := _c.NewCT_PlotAreaChoice()
	_eaa._cba.Chart.PlotArea.Choice = append(_eaa._cba.Chart.PlotArea.Choice, _feed)
	_feed.BarChart = _c.NewCT_BarChart()
	_feed.BarChart.Grouping = _c.NewCT_BarGrouping()
	_feed.BarChart.Grouping.ValAttr = _c.ST_BarGroupingStandard
	_dab := BarChart{_aad: _feed.BarChart}
	_dab.InitializeDefaults()
	return _dab
}

// X returns the inner wrapped XML type.
func (_ged Legend) X() *_c.CT_Legend { return _ged._gf }
func (_ca BarChart) AddAxis(axis Axis) {
	_fd := _c.NewCT_UnsignedInt()
	_fd.ValAttr = axis.AxisID()
	_ca._aad.AxId = append(_ca._aad.AxId, _fd)
}
func (_fdg Surface3DChart) InitializeDefaults() {
	_fdg._bfe.Wireframe = _c.NewCT_Boolean()
	_fdg._bfe.Wireframe.ValAttr = _a.Bool(false)
	_fdg._bfe.BandFmts = _c.NewCT_BandFmts()
	for _gddd := 0; _gddd < 15; _gddd++ {
		_eae := _c.NewCT_BandFmt()
		_eae.Idx.ValAttr = uint32(_gddd)
		_eae.SpPr = _gd.NewCT_ShapeProperties()
		_ece := _gb.MakeShapeProperties(_eae.SpPr)
		_ece.SetSolidFill(_fdg.nextColor(_gddd))
		_fdg._bfe.BandFmts.BandFmt = append(_fdg._bfe.BandFmts.BandFmt, _eae)
	}
}
func (_edf CategoryAxis) SetPosition(p _c.ST_AxPos) {
	_edf._ac.AxPos = _c.NewCT_AxPos()
	_edf._ac.AxPos.ValAttr = p
}
func (_eg CategoryAxis) SetTickLabelPosition(p _c.ST_TickLblPos) {
	if p == _c.ST_TickLblPosUnset {
		_eg._ac.TickLblPos = nil
	} else {
		_eg._ac.TickLblPos = _c.NewCT_TickLblPos()
		_eg._ac.TickLblPos.ValAttr = p
	}
}
func (_afa Title) InitializeDefaults() {
	_afa.SetText("\u0054\u0069\u0074l\u0065")
	_afa.RunProperties().SetSize(16 * _e.Point)
	_afa.RunProperties().SetSolidFill(_df.Black)
	_afa.RunProperties().SetFont("\u0043\u0061\u006c\u0069\u0062\u0020\u0072\u0069")
	_afa.RunProperties().SetBold(false)
}

// SetValues is used to set the source data to a set of values.
func (_fdc CategoryAxisDataSource) SetValues(v []string) {
	_fdc._fcbg.Choice = _c.NewCT_AxDataSourceChoice()
	_fdc._fcbg.Choice.StrLit = _c.NewCT_StrData()
	_fdc._fcbg.Choice.StrLit.PtCount = _c.NewCT_UnsignedInt()
	_fdc._fcbg.Choice.StrLit.PtCount.ValAttr = uint32(len(v))
	for _aade, _age := range v {
		_fdc._fcbg.Choice.StrLit.Pt = append(_fdc._fcbg.Choice.StrLit.Pt, &_c.CT_StrVal{IdxAttr: uint32(_aade), V: _age})
	}
}

// SetDisplayBlanksAs controls how missing values are displayed.
func (_fag Chart) SetDisplayBlanksAs(v _c.ST_DispBlanksAs) {
	_fag._cba.Chart.DispBlanksAs = _c.NewCT_DispBlanksAs()
	_fag._cba.Chart.DispBlanksAs.ValAttr = v
}

// Order returns the order of the series
func (_agg ScatterChartSeries) Order() uint32 { return _agg._bcb.Order.ValAttr }
func (_dacg ScatterChartSeries) InitializeDefaults() {
	_dacg.Properties().LineProperties().SetNoFill()
	_dacg.Marker().SetSymbol(_c.ST_MarkerStyleAuto)
	_dacg.Labels().SetShowLegendKey(false)
	_dacg.Labels().SetShowValue(true)
	_dacg.Labels().SetShowPercent(false)
	_dacg.Labels().SetShowCategoryName(false)
	_dacg.Labels().SetShowSeriesName(false)
	_dacg.Labels().SetShowLeaderLines(false)
}

type Line3DChart struct {
	chartBase
	_afcf *_c.CT_Line3DChart
}

func (_aaf Marker) Properties() _gb.ShapeProperties {
	if _aaf._acc.SpPr == nil {
		_aaf._acc.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_aaf._acc.SpPr)
}

// Properties returns the line chart series shape properties.
func (_ffd LineChartSeries) Properties() _gb.ShapeProperties {
	if _ffd._cegd.SpPr == nil {
		_ffd._cegd.SpPr = _gd.NewCT_ShapeProperties()
	}
	return _gb.MakeShapeProperties(_ffd._cegd.SpPr)
}
func (_cge DateAxis) SetMinorTickMark(m _c.ST_TickMark) {
	if m == _c.ST_TickMarkUnset {
		_cge._dgc.MinorTickMark = nil
	} else {
		_cge._dgc.MinorTickMark = _c.NewCT_TickMark()
		_cge._dgc.MinorTickMark.ValAttr = m
	}
}

// AddSeries adds a default series to an Pie chart.
func (_bfc PieChart) AddSeries() PieChartSeries {
	_bcac := _c.NewCT_PieSer()
	_bfc._efba.Ser = append(_bfc._efba.Ser, _bcac)
	_bcac.Idx.ValAttr = uint32(len(_bfc._efba.Ser) - 1)
	_bcac.Order.ValAttr = uint32(len(_bfc._efba.Ser) - 1)
	_ecf := PieChartSeries{_bcac}
	_ecf.InitializeDefaults()
	return _ecf
}

// MakeAxisDataSource constructs an AxisDataSource wrapper.
func MakeAxisDataSource(x *_c.CT_AxDataSource) CategoryAxisDataSource {
	return CategoryAxisDataSource{x}
}
func (_dac CategoryAxis) SetCrosses(axis Axis) {
	_dac._ac.Choice = _c.NewEG_AxSharedChoice()
	_dac._ac.Choice.Crosses = _c.NewCT_Crosses()
	_dac._ac.Choice.Crosses.ValAttr = _c.ST_CrossesAutoZero
	_dac._ac.CrossAx.ValAttr = axis.AxisID()
}

// AddAxis adds an axis to a Surface chart.
func (_feb Surface3DChart) AddAxis(axis Axis) {
	_gba := _c.NewCT_UnsignedInt()
	_gba.ValAttr = axis.AxisID()
	_feb._bfe.AxId = append(_feb._bfe.AxId, _gba)
}

// CategoryAxis returns the category data source.
func (_gdab PieChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _gdab._aab.Cat == nil {
		_gdab._aab.Cat = _c.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_gdab._aab.Cat)
}

// SetText sets the series text
func (_dcgb LineChartSeries) SetText(s string) {
	_dcgb._cegd.Tx = _c.NewCT_SerTx()
	_dcgb._cegd.Tx.Choice.V = &s
}

// SetText sets the series text.
func (_bgd AreaChartSeries) SetText(s string) {
	_bgd._be.Tx = _c.NewCT_SerTx()
	_bgd._be.Tx.Choice.V = &s
}

// X returns the inner wrapped XML type.
func (_ebce Pie3DChart) X() *_c.CT_Pie3DChart { return _ebce._ebc }
func (_eaae DataLabels) ensureChoice() {
	if _eaae._ecb.Choice == nil {
		_eaae._ecb.Choice = _c.NewCT_DLblsChoice()
	}
}

// SetNumberReference is used to set the source data to a range of cells containing
// numbers.
func (_eea CategoryAxisDataSource) SetNumberReference(s string) {
	_eea._fcbg.Choice = _c.NewCT_AxDataSourceChoice()
	_eea._fcbg.Choice.NumRef = _c.NewCT_NumRef()
	_eea._fcbg.Choice.NumRef.F = s
}

type Marker struct{ _acc *_c.CT_Marker }

func (_ceeb Legend) SetOverlay(b bool) {
	_ceeb._gf.Overlay = _c.NewCT_Boolean()
	_ceeb._gf.Overlay.ValAttr = _a.Bool(b)
}

// X returns the inner wrapped XML type.
func (_cfbd Chart) X() *_c.ChartSpace { return _cfbd._cba }

// SetExplosion sets the value that the segements of the pie are 'exploded' by
func (_fbba PieChartSeries) SetExplosion(v uint32) {
	_fbba._aab.Explosion = _c.NewCT_UnsignedInt()
	_fbba._aab.Explosion.ValAttr = v
}
func (_dca DateAxis) SetPosition(p _c.ST_AxPos) {
	_dca._dgc.AxPos = _c.NewCT_AxPos()
	_dca._dgc.AxPos.ValAttr = p
}

// AreaChartSeries is a series to be used on an area chart.
type AreaChartSeries struct{ _be *_c.CT_AreaSer }
