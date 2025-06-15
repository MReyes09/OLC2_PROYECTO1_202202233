// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gramAntlr // gramatica
import "github.com/antlr4-go/antlr/v4"

// BasegramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type BasegramaticaListener struct{}

var _ gramaticaListener = &BasegramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInicio is called when production inicio is entered.
func (s *BasegramaticaListener) EnterInicio(ctx *InicioContext) {}

// ExitInicio is called when production inicio is exited.
func (s *BasegramaticaListener) ExitInicio(ctx *InicioContext) {}

// EnterPrintStmt is called when production PrintStmt is entered.
func (s *BasegramaticaListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production PrintStmt is exited.
func (s *BasegramaticaListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterIfStmt is called when production IfStmt is entered.
func (s *BasegramaticaListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (s *BasegramaticaListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterSwitchInstruccion is called when production SwitchInstruccion is entered.
func (s *BasegramaticaListener) EnterSwitchInstruccion(ctx *SwitchInstruccionContext) {}

// ExitSwitchInstruccion is called when production SwitchInstruccion is exited.
func (s *BasegramaticaListener) ExitSwitchInstruccion(ctx *SwitchInstruccionContext) {}

// EnterSeccionInstruccion is called when production SeccionInstruccion is entered.
func (s *BasegramaticaListener) EnterSeccionInstruccion(ctx *SeccionInstruccionContext) {}

// ExitSeccionInstruccion is called when production SeccionInstruccion is exited.
func (s *BasegramaticaListener) ExitSeccionInstruccion(ctx *SeccionInstruccionContext) {}

// EnterForStmt is called when production ForStmt is entered.
func (s *BasegramaticaListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production ForStmt is exited.
func (s *BasegramaticaListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterVarDeclSliceStmt is called when production VarDeclSliceStmt is entered.
func (s *BasegramaticaListener) EnterVarDeclSliceStmt(ctx *VarDeclSliceStmtContext) {}

// ExitVarDeclSliceStmt is called when production VarDeclSliceStmt is exited.
func (s *BasegramaticaListener) ExitVarDeclSliceStmt(ctx *VarDeclSliceStmtContext) {}

// EnterAsignStmt is called when production AsignStmt is entered.
func (s *BasegramaticaListener) EnterAsignStmt(ctx *AsignStmtContext) {}

// ExitAsignStmt is called when production AsignStmt is exited.
func (s *BasegramaticaListener) ExitAsignStmt(ctx *AsignStmtContext) {}

// EnterVarDeclStmt is called when production VarDeclStmt is entered.
func (s *BasegramaticaListener) EnterVarDeclStmt(ctx *VarDeclStmtContext) {}

// ExitVarDeclStmt is called when production VarDeclStmt is exited.
func (s *BasegramaticaListener) ExitVarDeclStmt(ctx *VarDeclStmtContext) {}

// EnterVarDeclStructStmt is called when production VarDeclStructStmt is entered.
func (s *BasegramaticaListener) EnterVarDeclStructStmt(ctx *VarDeclStructStmtContext) {}

// ExitVarDeclStructStmt is called when production VarDeclStructStmt is exited.
func (s *BasegramaticaListener) ExitVarDeclStructStmt(ctx *VarDeclStructStmtContext) {}

// EnterVarStructDclStmt is called when production VarStructDclStmt is entered.
func (s *BasegramaticaListener) EnterVarStructDclStmt(ctx *VarStructDclStmtContext) {}

// ExitVarStructDclStmt is called when production VarStructDclStmt is exited.
func (s *BasegramaticaListener) ExitVarStructDclStmt(ctx *VarStructDclStmtContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BasegramaticaListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BasegramaticaListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BasegramaticaListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BasegramaticaListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterFunctionStmt is called when production FunctionStmt is entered.
func (s *BasegramaticaListener) EnterFunctionStmt(ctx *FunctionStmtContext) {}

// ExitFunctionStmt is called when production FunctionStmt is exited.
func (s *BasegramaticaListener) ExitFunctionStmt(ctx *FunctionStmtContext) {}

// EnterFunctionStructStmt is called when production FunctionStructStmt is entered.
func (s *BasegramaticaListener) EnterFunctionStructStmt(ctx *FunctionStructStmtContext) {}

// ExitFunctionStructStmt is called when production FunctionStructStmt is exited.
func (s *BasegramaticaListener) ExitFunctionStructStmt(ctx *FunctionStructStmtContext) {}

// EnterCallFunctionStmt is called when production CallFunctionStmt is entered.
func (s *BasegramaticaListener) EnterCallFunctionStmt(ctx *CallFunctionStmtContext) {}

// ExitCallFunctionStmt is called when production CallFunctionStmt is exited.
func (s *BasegramaticaListener) ExitCallFunctionStmt(ctx *CallFunctionStmtContext) {}

// EnterCallFunctionStructStmt is called when production CallFunctionStructStmt is entered.
func (s *BasegramaticaListener) EnterCallFunctionStructStmt(ctx *CallFunctionStructStmtContext) {}

// ExitCallFunctionStructStmt is called when production CallFunctionStructStmt is exited.
func (s *BasegramaticaListener) ExitCallFunctionStructStmt(ctx *CallFunctionStructStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BasegramaticaListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BasegramaticaListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterPrintln is called when production Println is entered.
func (s *BasegramaticaListener) EnterPrintln(ctx *PrintlnContext) {}

// ExitPrintln is called when production Println is exited.
func (s *BasegramaticaListener) ExitPrintln(ctx *PrintlnContext) {}

// EnterPrint is called when production Print is entered.
func (s *BasegramaticaListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production Print is exited.
func (s *BasegramaticaListener) ExitPrint(ctx *PrintContext) {}

// EnterIfOnly is called when production IfOnly is entered.
func (s *BasegramaticaListener) EnterIfOnly(ctx *IfOnlyContext) {}

// ExitIfOnly is called when production IfOnly is exited.
func (s *BasegramaticaListener) ExitIfOnly(ctx *IfOnlyContext) {}

// EnterIfAnidado is called when production IfAnidado is entered.
func (s *BasegramaticaListener) EnterIfAnidado(ctx *IfAnidadoContext) {}

// ExitIfAnidado is called when production IfAnidado is exited.
func (s *BasegramaticaListener) ExitIfAnidado(ctx *IfAnidadoContext) {}

// EnterSwitchStmt is called when production SwitchStmt is entered.
func (s *BasegramaticaListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production SwitchStmt is exited.
func (s *BasegramaticaListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterCase is called when production Case is entered.
func (s *BasegramaticaListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production Case is exited.
func (s *BasegramaticaListener) ExitCase(ctx *CaseContext) {}

// EnterDefault is called when production Default is entered.
func (s *BasegramaticaListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production Default is exited.
func (s *BasegramaticaListener) ExitDefault(ctx *DefaultContext) {}

// EnterBlockStmt is called when production blockStmt is entered.
func (s *BasegramaticaListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production blockStmt is exited.
func (s *BasegramaticaListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterForCondicion is called when production ForCondicion is entered.
func (s *BasegramaticaListener) EnterForCondicion(ctx *ForCondicionContext) {}

// ExitForCondicion is called when production ForCondicion is exited.
func (s *BasegramaticaListener) ExitForCondicion(ctx *ForCondicionContext) {}

// EnterForAsignacion is called when production ForAsignacion is entered.
func (s *BasegramaticaListener) EnterForAsignacion(ctx *ForAsignacionContext) {}

// ExitForAsignacion is called when production ForAsignacion is exited.
func (s *BasegramaticaListener) ExitForAsignacion(ctx *ForAsignacionContext) {}

// EnterForRange is called when production ForRange is entered.
func (s *BasegramaticaListener) EnterForRange(ctx *ForRangeContext) {}

// ExitForRange is called when production ForRange is exited.
func (s *BasegramaticaListener) ExitForRange(ctx *ForRangeContext) {}

// EnterVarDclWithTypeAndValue is called when production VarDclWithTypeAndValue is entered.
func (s *BasegramaticaListener) EnterVarDclWithTypeAndValue(ctx *VarDclWithTypeAndValueContext) {}

// ExitVarDclWithTypeAndValue is called when production VarDclWithTypeAndValue is exited.
func (s *BasegramaticaListener) ExitVarDclWithTypeAndValue(ctx *VarDclWithTypeAndValueContext) {}

// EnterVarDclWithTypeOnly is called when production VarDclWithTypeOnly is entered.
func (s *BasegramaticaListener) EnterVarDclWithTypeOnly(ctx *VarDclWithTypeOnlyContext) {}

// ExitVarDclWithTypeOnly is called when production VarDclWithTypeOnly is exited.
func (s *BasegramaticaListener) ExitVarDclWithTypeOnly(ctx *VarDclWithTypeOnlyContext) {}

// EnterVarDclWithInference is called when production VarDclWithInference is entered.
func (s *BasegramaticaListener) EnterVarDclWithInference(ctx *VarDclWithInferenceContext) {}

// ExitVarDclWithInference is called when production VarDclWithInference is exited.
func (s *BasegramaticaListener) ExitVarDclWithInference(ctx *VarDclWithInferenceContext) {}

// EnterSliceValores is called when production SliceValores is entered.
func (s *BasegramaticaListener) EnterSliceValores(ctx *SliceValoresContext) {}

// ExitSliceValores is called when production SliceValores is exited.
func (s *BasegramaticaListener) ExitSliceValores(ctx *SliceValoresContext) {}

// EnterSliceVacio is called when production SliceVacio is entered.
func (s *BasegramaticaListener) EnterSliceVacio(ctx *SliceVacioContext) {}

// ExitSliceVacio is called when production SliceVacio is exited.
func (s *BasegramaticaListener) ExitSliceVacio(ctx *SliceVacioContext) {}

// EnterSliceDcl_Asign is called when production SliceDcl_Asign is entered.
func (s *BasegramaticaListener) EnterSliceDcl_Asign(ctx *SliceDcl_AsignContext) {}

// ExitSliceDcl_Asign is called when production SliceDcl_Asign is exited.
func (s *BasegramaticaListener) ExitSliceDcl_Asign(ctx *SliceDcl_AsignContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasegramaticaListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasegramaticaListener) ExitAssign(ctx *AssignContext) {}

// EnterNuevoSlice is called when production nuevoSlice is entered.
func (s *BasegramaticaListener) EnterNuevoSlice(ctx *NuevoSliceContext) {}

// ExitNuevoSlice is called when production nuevoSlice is exited.
func (s *BasegramaticaListener) ExitNuevoSlice(ctx *NuevoSliceContext) {}

// EnterSliceContenido is called when production SliceContenido is entered.
func (s *BasegramaticaListener) EnterSliceContenido(ctx *SliceContenidoContext) {}

// ExitSliceContenido is called when production SliceContenido is exited.
func (s *BasegramaticaListener) ExitSliceContenido(ctx *SliceContenidoContext) {}

// EnterSliceContenidoSlice is called when production SliceContenidoSlice is entered.
func (s *BasegramaticaListener) EnterSliceContenidoSlice(ctx *SliceContenidoSliceContext) {}

// ExitSliceContenidoSlice is called when production SliceContenidoSlice is exited.
func (s *BasegramaticaListener) ExitSliceContenidoSlice(ctx *SliceContenidoSliceContext) {}

// EnterDeclStructData is called when production DeclStructData is entered.
func (s *BasegramaticaListener) EnterDeclStructData(ctx *DeclStructDataContext) {}

// ExitDeclStructData is called when production DeclStructData is exited.
func (s *BasegramaticaListener) ExitDeclStructData(ctx *DeclStructDataContext) {}

// EnterStructVarTypeInference is called when production StructVarTypeInference is entered.
func (s *BasegramaticaListener) EnterStructVarTypeInference(ctx *StructVarTypeInferenceContext) {}

// ExitStructVarTypeInference is called when production StructVarTypeInference is exited.
func (s *BasegramaticaListener) ExitStructVarTypeInference(ctx *StructVarTypeInferenceContext) {}

// EnterVarExpr is called when production varExpr is entered.
func (s *BasegramaticaListener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production varExpr is exited.
func (s *BasegramaticaListener) ExitVarExpr(ctx *VarExprContext) {}

// EnterVarAdd is called when production varAdd is entered.
func (s *BasegramaticaListener) EnterVarAdd(ctx *VarAddContext) {}

// ExitVarAdd is called when production varAdd is exited.
func (s *BasegramaticaListener) ExitVarAdd(ctx *VarAddContext) {}

// EnterVarInc is called when production varInc is entered.
func (s *BasegramaticaListener) EnterVarInc(ctx *VarIncContext) {}

// ExitVarInc is called when production varInc is exited.
func (s *BasegramaticaListener) ExitVarInc(ctx *VarIncContext) {}

// EnterArrayAccess is called when production ArrayAccess is entered.
func (s *BasegramaticaListener) EnterArrayAccess(ctx *ArrayAccessContext) {}

// ExitArrayAccess is called when production ArrayAccess is exited.
func (s *BasegramaticaListener) ExitArrayAccess(ctx *ArrayAccessContext) {}

// EnterStructAccessAsign is called when production StructAccessAsign is entered.
func (s *BasegramaticaListener) EnterStructAccessAsign(ctx *StructAccessAsignContext) {}

// ExitStructAccessAsign is called when production StructAccessAsign is exited.
func (s *BasegramaticaListener) ExitStructAccessAsign(ctx *StructAccessAsignContext) {}

// EnterParens is called when production Parens is entered.
func (s *BasegramaticaListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BasegramaticaListener) ExitParens(ctx *ParensContext) {}

// EnterCallFunctionValue is called when production CallFunctionValue is entered.
func (s *BasegramaticaListener) EnterCallFunctionValue(ctx *CallFunctionValueContext) {}

// ExitCallFunctionValue is called when production CallFunctionValue is exited.
func (s *BasegramaticaListener) ExitCallFunctionValue(ctx *CallFunctionValueContext) {}

// EnterLogical is called when production Logical is entered.
func (s *BasegramaticaListener) EnterLogical(ctx *LogicalContext) {}

// ExitLogical is called when production Logical is exited.
func (s *BasegramaticaListener) ExitLogical(ctx *LogicalContext) {}

// EnterString is called when production String is entered.
func (s *BasegramaticaListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BasegramaticaListener) ExitString(ctx *StringContext) {}

// EnterStructAccess is called when production StructAccess is entered.
func (s *BasegramaticaListener) EnterStructAccess(ctx *StructAccessContext) {}

// ExitStructAccess is called when production StructAccess is exited.
func (s *BasegramaticaListener) ExitStructAccess(ctx *StructAccessContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BasegramaticaListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BasegramaticaListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterChar is called when production Char is entered.
func (s *BasegramaticaListener) EnterChar(ctx *CharContext) {}

// ExitChar is called when production Char is exited.
func (s *BasegramaticaListener) ExitChar(ctx *CharContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BasegramaticaListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BasegramaticaListener) ExitBoolean(ctx *BooleanContext) {}

// EnterCallFunctionStructValue is called when production CallFunctionStructValue is entered.
func (s *BasegramaticaListener) EnterCallFunctionStructValue(ctx *CallFunctionStructValueContext) {}

// ExitCallFunctionStructValue is called when production CallFunctionStructValue is exited.
func (s *BasegramaticaListener) ExitCallFunctionStructValue(ctx *CallFunctionStructValueContext) {}

// EnterArrayFindIndex is called when production ArrayFindIndex is entered.
func (s *BasegramaticaListener) EnterArrayFindIndex(ctx *ArrayFindIndexContext) {}

// ExitArrayFindIndex is called when production ArrayFindIndex is exited.
func (s *BasegramaticaListener) ExitArrayFindIndex(ctx *ArrayFindIndexContext) {}

// EnterArrayAppend is called when production ArrayAppend is entered.
func (s *BasegramaticaListener) EnterArrayAppend(ctx *ArrayAppendContext) {}

// ExitArrayAppend is called when production ArrayAppend is exited.
func (s *BasegramaticaListener) ExitArrayAppend(ctx *ArrayAppendContext) {}

// EnterEqualsNotEquals is called when production EqualsNotEquals is entered.
func (s *BasegramaticaListener) EnterEqualsNotEquals(ctx *EqualsNotEqualsContext) {}

// ExitEqualsNotEquals is called when production EqualsNotEquals is exited.
func (s *BasegramaticaListener) ExitEqualsNotEquals(ctx *EqualsNotEqualsContext) {}

// EnterIntToString is called when production IntToString is entered.
func (s *BasegramaticaListener) EnterIntToString(ctx *IntToStringContext) {}

// ExitIntToString is called when production IntToString is exited.
func (s *BasegramaticaListener) ExitIntToString(ctx *IntToStringContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BasegramaticaListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BasegramaticaListener) ExitAddSub(ctx *AddSubContext) {}

// EnterArrayAccessSimple is called when production ArrayAccessSimple is entered.
func (s *BasegramaticaListener) EnterArrayAccessSimple(ctx *ArrayAccessSimpleContext) {}

// ExitArrayAccessSimple is called when production ArrayAccessSimple is exited.
func (s *BasegramaticaListener) ExitArrayAccessSimple(ctx *ArrayAccessSimpleContext) {}

// EnterArrayLength is called when production ArrayLength is entered.
func (s *BasegramaticaListener) EnterArrayLength(ctx *ArrayLengthContext) {}

// ExitArrayLength is called when production ArrayLength is exited.
func (s *BasegramaticaListener) ExitArrayLength(ctx *ArrayLengthContext) {}

// EnterMulDivModulo is called when production MulDivModulo is entered.
func (s *BasegramaticaListener) EnterMulDivModulo(ctx *MulDivModuloContext) {}

// ExitMulDivModulo is called when production MulDivModulo is exited.
func (s *BasegramaticaListener) ExitMulDivModulo(ctx *MulDivModuloContext) {}

// EnterDouble is called when production Double is entered.
func (s *BasegramaticaListener) EnterDouble(ctx *DoubleContext) {}

// ExitDouble is called when production Double is exited.
func (s *BasegramaticaListener) ExitDouble(ctx *DoubleContext) {}

// EnterInteger is called when production Integer is entered.
func (s *BasegramaticaListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production Integer is exited.
func (s *BasegramaticaListener) ExitInteger(ctx *IntegerContext) {}

// EnterNil is called when production Nil is entered.
func (s *BasegramaticaListener) EnterNil(ctx *NilContext) {}

// ExitNil is called when production Nil is exited.
func (s *BasegramaticaListener) ExitNil(ctx *NilContext) {}

// EnterMinorMajorEqual is called when production MinorMajorEqual is entered.
func (s *BasegramaticaListener) EnterMinorMajorEqual(ctx *MinorMajorEqualContext) {}

// ExitMinorMajorEqual is called when production MinorMajorEqual is exited.
func (s *BasegramaticaListener) ExitMinorMajorEqual(ctx *MinorMajorEqualContext) {}

// EnterNot is called when production Not is entered.
func (s *BasegramaticaListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BasegramaticaListener) ExitNot(ctx *NotContext) {}

// EnterReflectType is called when production reflectType is entered.
func (s *BasegramaticaListener) EnterReflectType(ctx *ReflectTypeContext) {}

// ExitReflectType is called when production reflectType is exited.
func (s *BasegramaticaListener) ExitReflectType(ctx *ReflectTypeContext) {}

// EnterNegate is called when production Negate is entered.
func (s *BasegramaticaListener) EnterNegate(ctx *NegateContext) {}

// ExitNegate is called when production Negate is exited.
func (s *BasegramaticaListener) ExitNegate(ctx *NegateContext) {}

// EnterArrayJoin is called when production ArrayJoin is entered.
func (s *BasegramaticaListener) EnterArrayJoin(ctx *ArrayJoinContext) {}

// ExitArrayJoin is called when production ArrayJoin is exited.
func (s *BasegramaticaListener) ExitArrayJoin(ctx *ArrayJoinContext) {}

// EnterFloatToString is called when production floatToString is entered.
func (s *BasegramaticaListener) EnterFloatToString(ctx *FloatToStringContext) {}

// ExitFloatToString is called when production floatToString is exited.
func (s *BasegramaticaListener) ExitFloatToString(ctx *FloatToStringContext) {}

// EnterPosicion is called when production posicion is entered.
func (s *BasegramaticaListener) EnterPosicion(ctx *PosicionContext) {}

// ExitPosicion is called when production posicion is exited.
func (s *BasegramaticaListener) ExitPosicion(ctx *PosicionContext) {}

// EnterType is called when production type is entered.
func (s *BasegramaticaListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasegramaticaListener) ExitType(ctx *TypeContext) {}

// EnterBreak is called when production break is entered.
func (s *BasegramaticaListener) EnterBreak(ctx *BreakContext) {}

// ExitBreak is called when production break is exited.
func (s *BasegramaticaListener) ExitBreak(ctx *BreakContext) {}

// EnterContinue is called when production continue is entered.
func (s *BasegramaticaListener) EnterContinue(ctx *ContinueContext) {}

// ExitContinue is called when production continue is exited.
func (s *BasegramaticaListener) ExitContinue(ctx *ContinueContext) {}

// EnterFunciones is called when production Funciones is entered.
func (s *BasegramaticaListener) EnterFunciones(ctx *FuncionesContext) {}

// ExitFunciones is called when production Funciones is exited.
func (s *BasegramaticaListener) ExitFunciones(ctx *FuncionesContext) {}

// EnterFuncionesStructsNativas is called when production FuncionesStructsNativas is entered.
func (s *BasegramaticaListener) EnterFuncionesStructsNativas(ctx *FuncionesStructsNativasContext) {}

// ExitFuncionesStructsNativas is called when production FuncionesStructsNativas is exited.
func (s *BasegramaticaListener) ExitFuncionesStructsNativas(ctx *FuncionesStructsNativasContext) {}

// EnterDefParams is called when production defParams is entered.
func (s *BasegramaticaListener) EnterDefParams(ctx *DefParamsContext) {}

// ExitDefParams is called when production defParams is exited.
func (s *BasegramaticaListener) ExitDefParams(ctx *DefParamsContext) {}

// EnterCallFunction is called when production CallFunction is entered.
func (s *BasegramaticaListener) EnterCallFunction(ctx *CallFunctionContext) {}

// ExitCallFunction is called when production CallFunction is exited.
func (s *BasegramaticaListener) ExitCallFunction(ctx *CallFunctionContext) {}

// EnterCallFunctionStruct is called when production CallFunctionStruct is entered.
func (s *BasegramaticaListener) EnterCallFunctionStruct(ctx *CallFunctionStructContext) {}

// ExitCallFunctionStruct is called when production CallFunctionStruct is exited.
func (s *BasegramaticaListener) ExitCallFunctionStruct(ctx *CallFunctionStructContext) {}

// EnterValRet is called when production valRet is entered.
func (s *BasegramaticaListener) EnterValRet(ctx *ValRetContext) {}

// ExitValRet is called when production valRet is exited.
func (s *BasegramaticaListener) ExitValRet(ctx *ValRetContext) {}

// EnterRetorno is called when production retorno is entered.
func (s *BasegramaticaListener) EnterRetorno(ctx *RetornoContext) {}

// ExitRetorno is called when production retorno is exited.
func (s *BasegramaticaListener) ExitRetorno(ctx *RetornoContext) {}
