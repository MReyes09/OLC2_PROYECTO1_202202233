// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gramAntlr // gramatica
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by gramaticaParser.
type gramaticaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by gramaticaParser#inicio.
	VisitInicio(ctx *InicioContext) interface{}

	// Visit a parse tree produced by gramaticaParser#PrintStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#IfStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SwitchInstruccion.
	VisitSwitchInstruccion(ctx *SwitchInstruccionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SeccionInstruccion.
	VisitSeccionInstruccion(ctx *SeccionInstruccionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ForStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#VarDeclSliceStmt.
	VisitVarDeclSliceStmt(ctx *VarDeclSliceStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#AsignStmt.
	VisitAsignStmt(ctx *AsignStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#VarDeclStmt.
	VisitVarDeclStmt(ctx *VarDeclStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#VarDeclStructStmt.
	VisitVarDeclStructStmt(ctx *VarDeclStructStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#VarStructDclStmt.
	VisitVarStructDclStmt(ctx *VarStructDclStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ContinueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#FunctionStmt.
	VisitFunctionStmt(ctx *FunctionStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#FunctionStructStmt.
	VisitFunctionStructStmt(ctx *FunctionStructStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#CallFunctionStmt.
	VisitCallFunctionStmt(ctx *CallFunctionStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#CallFunctionStructStmt.
	VisitCallFunctionStructStmt(ctx *CallFunctionStructStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Println.
	VisitPrintln(ctx *PrintlnContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by gramaticaParser#IfOnly.
	VisitIfOnly(ctx *IfOnlyContext) interface{}

	// Visit a parse tree produced by gramaticaParser#IfAnidado.
	VisitIfAnidado(ctx *IfAnidadoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SwitchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Default.
	VisitDefault(ctx *DefaultContext) interface{}

	// Visit a parse tree produced by gramaticaParser#blockStmt.
	VisitBlockStmt(ctx *BlockStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ForCondicion.
	VisitForCondicion(ctx *ForCondicionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ForAsignacion.
	VisitForAsignacion(ctx *ForAsignacionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ForRange.
	VisitForRange(ctx *ForRangeContext) interface{}

	// Visit a parse tree produced by gramaticaParser#VarDclWithTypeAndValue.
	VisitVarDclWithTypeAndValue(ctx *VarDclWithTypeAndValueContext) interface{}

	// Visit a parse tree produced by gramaticaParser#VarDclWithTypeOnly.
	VisitVarDclWithTypeOnly(ctx *VarDclWithTypeOnlyContext) interface{}

	// Visit a parse tree produced by gramaticaParser#VarDclWithInference.
	VisitVarDclWithInference(ctx *VarDclWithInferenceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SliceValores.
	VisitSliceValores(ctx *SliceValoresContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SliceVacio.
	VisitSliceVacio(ctx *SliceVacioContext) interface{}

	// Visit a parse tree produced by gramaticaParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by gramaticaParser#nuevoSlice.
	VisitNuevoSlice(ctx *NuevoSliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SliceContenido.
	VisitSliceContenido(ctx *SliceContenidoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SliceContenidoSlice.
	VisitSliceContenidoSlice(ctx *SliceContenidoSliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#DeclStructData.
	VisitDeclStructData(ctx *DeclStructDataContext) interface{}

	// Visit a parse tree produced by gramaticaParser#StructVarType.
	VisitStructVarType(ctx *StructVarTypeContext) interface{}

	// Visit a parse tree produced by gramaticaParser#StructVarTypeInference.
	VisitStructVarTypeInference(ctx *StructVarTypeInferenceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#varExpr.
	VisitVarExpr(ctx *VarExprContext) interface{}

	// Visit a parse tree produced by gramaticaParser#varAdd.
	VisitVarAdd(ctx *VarAddContext) interface{}

	// Visit a parse tree produced by gramaticaParser#varInc.
	VisitVarInc(ctx *VarIncContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ArrayAccess.
	VisitArrayAccess(ctx *ArrayAccessContext) interface{}

	// Visit a parse tree produced by gramaticaParser#StructAccessAsign.
	VisitStructAccessAsign(ctx *StructAccessAsignContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by gramaticaParser#CallFunctionValue.
	VisitCallFunctionValue(ctx *CallFunctionValueContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Logical.
	VisitLogical(ctx *LogicalContext) interface{}

	// Visit a parse tree produced by gramaticaParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by gramaticaParser#StructAccess.
	VisitStructAccess(ctx *StructAccessContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Char.
	VisitChar(ctx *CharContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by gramaticaParser#CallFunctionStructValue.
	VisitCallFunctionStructValue(ctx *CallFunctionStructValueContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ArrayFindIndex.
	VisitArrayFindIndex(ctx *ArrayFindIndexContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ArrayAppend.
	VisitArrayAppend(ctx *ArrayAppendContext) interface{}

	// Visit a parse tree produced by gramaticaParser#EqualsNotEquals.
	VisitEqualsNotEquals(ctx *EqualsNotEqualsContext) interface{}

	// Visit a parse tree produced by gramaticaParser#IntToString.
	VisitIntToString(ctx *IntToStringContext) interface{}

	// Visit a parse tree produced by gramaticaParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ArrayAccessSimple.
	VisitArrayAccessSimple(ctx *ArrayAccessSimpleContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ArrayLength.
	VisitArrayLength(ctx *ArrayLengthContext) interface{}

	// Visit a parse tree produced by gramaticaParser#MulDivModulo.
	VisitMulDivModulo(ctx *MulDivModuloContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Double.
	VisitDouble(ctx *DoubleContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Nil.
	VisitNil(ctx *NilContext) interface{}

	// Visit a parse tree produced by gramaticaParser#MinorMajorEqual.
	VisitMinorMajorEqual(ctx *MinorMajorEqualContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by gramaticaParser#reflectType.
	VisitReflectType(ctx *ReflectTypeContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Negate.
	VisitNegate(ctx *NegateContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ArrayJoin.
	VisitArrayJoin(ctx *ArrayJoinContext) interface{}

	// Visit a parse tree produced by gramaticaParser#floatToString.
	VisitFloatToString(ctx *FloatToStringContext) interface{}

	// Visit a parse tree produced by gramaticaParser#posicion.
	VisitPosicion(ctx *PosicionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by gramaticaParser#break.
	VisitBreak(ctx *BreakContext) interface{}

	// Visit a parse tree produced by gramaticaParser#continue.
	VisitContinue(ctx *ContinueContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Funciones.
	VisitFunciones(ctx *FuncionesContext) interface{}

	// Visit a parse tree produced by gramaticaParser#FuncionesStructsNativas.
	VisitFuncionesStructsNativas(ctx *FuncionesStructsNativasContext) interface{}

	// Visit a parse tree produced by gramaticaParser#defParams.
	VisitDefParams(ctx *DefParamsContext) interface{}

	// Visit a parse tree produced by gramaticaParser#CallFunction.
	VisitCallFunction(ctx *CallFunctionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#CallFunctionStruct.
	VisitCallFunctionStruct(ctx *CallFunctionStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#valRet.
	VisitValRet(ctx *ValRetContext) interface{}

	// Visit a parse tree produced by gramaticaParser#retorno.
	VisitRetorno(ctx *RetornoContext) interface{}
}
