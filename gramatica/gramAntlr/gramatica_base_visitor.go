// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gramAntlr // gramatica
import "github.com/antlr4-go/antlr/v4"

type BasegramaticaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasegramaticaVisitor) VisitInicio(ctx *InicioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSwitchInstruccion(ctx *SwitchInstruccionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSeccionInstruccion(ctx *SeccionInstruccionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarDeclSliceStmt(ctx *VarDeclSliceStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsignStmt(ctx *AsignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarDeclStmt(ctx *VarDeclStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarDeclStructStmt(ctx *VarDeclStructStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarStructDclStmt(ctx *VarStructDclStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFunctionStmt(ctx *FunctionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFunctionStructStmt(ctx *FunctionStructStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCallFunctionStmt(ctx *CallFunctionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCallFunctionStructStmt(ctx *CallFunctionStructStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitPrintln(ctx *PrintlnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitPrint(ctx *PrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitIfOnly(ctx *IfOnlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitIfAnidado(ctx *IfAnidadoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCase(ctx *CaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDefault(ctx *DefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBlockStmt(ctx *BlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitForCondicion(ctx *ForCondicionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitForAsignacion(ctx *ForAsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitForRange(ctx *ForRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarDclWithTypeAndValue(ctx *VarDclWithTypeAndValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarDclWithTypeOnly(ctx *VarDclWithTypeOnlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarDclWithInference(ctx *VarDclWithInferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSliceValores(ctx *SliceValoresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSliceVacio(ctx *SliceVacioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSliceDcl_Asign(ctx *SliceDcl_AsignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitNuevoSlice(ctx *NuevoSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSliceContenido(ctx *SliceContenidoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSliceContenidoSlice(ctx *SliceContenidoSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDeclStructData(ctx *DeclStructDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitStructVarTypeInference(ctx *StructVarTypeInferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarExpr(ctx *VarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarAdd(ctx *VarAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitVarInc(ctx *VarIncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitArrayAccess(ctx *ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitStructAccessAsign(ctx *StructAccessAsignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCallFunctionValue(ctx *CallFunctionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLogical(ctx *LogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitStructAccess(ctx *StructAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitChar(ctx *CharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCallFunctionStructValue(ctx *CallFunctionStructValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitArrayFindIndex(ctx *ArrayFindIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitArrayAppend(ctx *ArrayAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitEqualsNotEquals(ctx *EqualsNotEqualsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitIntToString(ctx *IntToStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitArrayAccessSimple(ctx *ArrayAccessSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitArrayLength(ctx *ArrayLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitMulDivModulo(ctx *MulDivModuloContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDouble(ctx *DoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitNil(ctx *NilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitMinorMajorEqual(ctx *MinorMajorEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitReflectType(ctx *ReflectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitNegate(ctx *NegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitArrayJoin(ctx *ArrayJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFloatToString(ctx *FloatToStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitPosicion(ctx *PosicionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBreak(ctx *BreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitContinue(ctx *ContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFunciones(ctx *FuncionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFuncionesStructsNativas(ctx *FuncionesStructsNativasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDefParams(ctx *DefParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCallFunction(ctx *CallFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCallFunctionStruct(ctx *CallFunctionStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitValRet(ctx *ValRetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitRetorno(ctx *RetornoContext) interface{} {
	return v.VisitChildren(ctx)
}
