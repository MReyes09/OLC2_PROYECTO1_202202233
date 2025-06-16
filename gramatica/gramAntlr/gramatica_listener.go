// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gramAntlr // gramatica
import "github.com/antlr4-go/antlr/v4"

// gramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type gramaticaListener interface {
	antlr.ParseTreeListener

	// EnterInicio is called when entering the inicio production.
	EnterInicio(c *InicioContext)

	// EnterPrintStmt is called when entering the PrintStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterSwitchInstruccion is called when entering the SwitchInstruccion production.
	EnterSwitchInstruccion(c *SwitchInstruccionContext)

	// EnterSeccionInstruccion is called when entering the SeccionInstruccion production.
	EnterSeccionInstruccion(c *SeccionInstruccionContext)

	// EnterForStmt is called when entering the ForStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterVarDeclSliceStmt is called when entering the VarDeclSliceStmt production.
	EnterVarDeclSliceStmt(c *VarDeclSliceStmtContext)

	// EnterAsignStmt is called when entering the AsignStmt production.
	EnterAsignStmt(c *AsignStmtContext)

	// EnterVarDeclStmt is called when entering the VarDeclStmt production.
	EnterVarDeclStmt(c *VarDeclStmtContext)

	// EnterVarDeclStructStmt is called when entering the VarDeclStructStmt production.
	EnterVarDeclStructStmt(c *VarDeclStructStmtContext)

	// EnterVarStructDclStmt is called when entering the VarStructDclStmt production.
	EnterVarStructDclStmt(c *VarStructDclStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterFunctionStmt is called when entering the FunctionStmt production.
	EnterFunctionStmt(c *FunctionStmtContext)

	// EnterFunctionStructStmt is called when entering the FunctionStructStmt production.
	EnterFunctionStructStmt(c *FunctionStructStmtContext)

	// EnterCallFunctionStmt is called when entering the CallFunctionStmt production.
	EnterCallFunctionStmt(c *CallFunctionStmtContext)

	// EnterCallFunctionStructStmt is called when entering the CallFunctionStructStmt production.
	EnterCallFunctionStructStmt(c *CallFunctionStructStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterPrintln is called when entering the Println production.
	EnterPrintln(c *PrintlnContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

	// EnterIfOnly is called when entering the IfOnly production.
	EnterIfOnly(c *IfOnlyContext)

	// EnterIfAnidado is called when entering the IfAnidado production.
	EnterIfAnidado(c *IfAnidadoContext)

	// EnterSwitchStmt is called when entering the SwitchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterCase is called when entering the Case production.
	EnterCase(c *CaseContext)

	// EnterDefault is called when entering the Default production.
	EnterDefault(c *DefaultContext)

	// EnterBlockStmt is called when entering the blockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterForCondicion is called when entering the ForCondicion production.
	EnterForCondicion(c *ForCondicionContext)

	// EnterForAsignacion is called when entering the ForAsignacion production.
	EnterForAsignacion(c *ForAsignacionContext)

	// EnterForRange is called when entering the ForRange production.
	EnterForRange(c *ForRangeContext)

	// EnterVarDclWithTypeAndValue is called when entering the VarDclWithTypeAndValue production.
	EnterVarDclWithTypeAndValue(c *VarDclWithTypeAndValueContext)

	// EnterVarDclWithTypeOnly is called when entering the VarDclWithTypeOnly production.
	EnterVarDclWithTypeOnly(c *VarDclWithTypeOnlyContext)

	// EnterVarDclWithInference is called when entering the VarDclWithInference production.
	EnterVarDclWithInference(c *VarDclWithInferenceContext)

	// EnterSliceValores is called when entering the SliceValores production.
	EnterSliceValores(c *SliceValoresContext)

	// EnterSliceVacio is called when entering the SliceVacio production.
	EnterSliceVacio(c *SliceVacioContext)

	// EnterSliceDcl_Asign is called when entering the SliceDcl_Asign production.
	EnterSliceDcl_Asign(c *SliceDcl_AsignContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterNuevoSlice is called when entering the nuevoSlice production.
	EnterNuevoSlice(c *NuevoSliceContext)

	// EnterSliceContenido is called when entering the SliceContenido production.
	EnterSliceContenido(c *SliceContenidoContext)

	// EnterSliceContenidoSlice is called when entering the SliceContenidoSlice production.
	EnterSliceContenidoSlice(c *SliceContenidoSliceContext)

	// EnterDeclStructData is called when entering the DeclStructData production.
	EnterDeclStructData(c *DeclStructDataContext)

	// EnterStructVarTypeInference is called when entering the StructVarTypeInference production.
	EnterStructVarTypeInference(c *StructVarTypeInferenceContext)

	// EnterVarExpr is called when entering the varExpr production.
	EnterVarExpr(c *VarExprContext)

	// EnterVarAdd is called when entering the varAdd production.
	EnterVarAdd(c *VarAddContext)

	// EnterVarInc is called when entering the varInc production.
	EnterVarInc(c *VarIncContext)

	// EnterArrayAccess is called when entering the ArrayAccess production.
	EnterArrayAccess(c *ArrayAccessContext)

	// EnterStructAccessAsign is called when entering the StructAccessAsign production.
	EnterStructAccessAsign(c *StructAccessAsignContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterCallFunctionValue is called when entering the CallFunctionValue production.
	EnterCallFunctionValue(c *CallFunctionValueContext)

	// EnterLogical is called when entering the Logical production.
	EnterLogical(c *LogicalContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterStructAccess is called when entering the StructAccess production.
	EnterStructAccess(c *StructAccessContext)

	// EnterIdentifier is called when entering the Identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterChar is called when entering the Char production.
	EnterChar(c *CharContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterCallFunctionStructValue is called when entering the CallFunctionStructValue production.
	EnterCallFunctionStructValue(c *CallFunctionStructValueContext)

	// EnterArrayFindIndex is called when entering the ArrayFindIndex production.
	EnterArrayFindIndex(c *ArrayFindIndexContext)

	// EnterArrayAppend is called when entering the ArrayAppend production.
	EnterArrayAppend(c *ArrayAppendContext)

	// EnterEqualsNotEquals is called when entering the EqualsNotEquals production.
	EnterEqualsNotEquals(c *EqualsNotEqualsContext)

	// EnterIntToString is called when entering the IntToString production.
	EnterIntToString(c *IntToStringContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterArrayAccessSimple is called when entering the ArrayAccessSimple production.
	EnterArrayAccessSimple(c *ArrayAccessSimpleContext)

	// EnterArrayLength is called when entering the ArrayLength production.
	EnterArrayLength(c *ArrayLengthContext)

	// EnterMulDivModulo is called when entering the MulDivModulo production.
	EnterMulDivModulo(c *MulDivModuloContext)

	// EnterDouble is called when entering the Double production.
	EnterDouble(c *DoubleContext)

	// EnterInteger is called when entering the Integer production.
	EnterInteger(c *IntegerContext)

	// EnterNil is called when entering the Nil production.
	EnterNil(c *NilContext)

	// EnterMinorMajorEqual is called when entering the MinorMajorEqual production.
	EnterMinorMajorEqual(c *MinorMajorEqualContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterReflectType is called when entering the reflectType production.
	EnterReflectType(c *ReflectTypeContext)

	// EnterNegate is called when entering the Negate production.
	EnterNegate(c *NegateContext)

	// EnterArrayJoin is called when entering the ArrayJoin production.
	EnterArrayJoin(c *ArrayJoinContext)

	// EnterFloatToString is called when entering the floatToString production.
	EnterFloatToString(c *FloatToStringContext)

	// EnterPosicion is called when entering the posicion production.
	EnterPosicion(c *PosicionContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterBreak is called when entering the break production.
	EnterBreak(c *BreakContext)

	// EnterContinue is called when entering the continue production.
	EnterContinue(c *ContinueContext)

	// EnterFunciones is called when entering the Funciones production.
	EnterFunciones(c *FuncionesContext)

	// EnterFuncionesStructsNativas is called when entering the FuncionesStructsNativas production.
	EnterFuncionesStructsNativas(c *FuncionesStructsNativasContext)

	// EnterDefParams is called when entering the defParams production.
	EnterDefParams(c *DefParamsContext)

	// EnterCallFunction is called when entering the CallFunction production.
	EnterCallFunction(c *CallFunctionContext)

	// EnterCallFunctionStruct is called when entering the CallFunctionStruct production.
	EnterCallFunctionStruct(c *CallFunctionStructContext)

	// EnterValRet is called when entering the valRet production.
	EnterValRet(c *ValRetContext)

	// EnterRetorno is called when entering the retorno production.
	EnterRetorno(c *RetornoContext)

	// ExitInicio is called when exiting the inicio production.
	ExitInicio(c *InicioContext)

	// ExitPrintStmt is called when exiting the PrintStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitSwitchInstruccion is called when exiting the SwitchInstruccion production.
	ExitSwitchInstruccion(c *SwitchInstruccionContext)

	// ExitSeccionInstruccion is called when exiting the SeccionInstruccion production.
	ExitSeccionInstruccion(c *SeccionInstruccionContext)

	// ExitForStmt is called when exiting the ForStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitVarDeclSliceStmt is called when exiting the VarDeclSliceStmt production.
	ExitVarDeclSliceStmt(c *VarDeclSliceStmtContext)

	// ExitAsignStmt is called when exiting the AsignStmt production.
	ExitAsignStmt(c *AsignStmtContext)

	// ExitVarDeclStmt is called when exiting the VarDeclStmt production.
	ExitVarDeclStmt(c *VarDeclStmtContext)

	// ExitVarDeclStructStmt is called when exiting the VarDeclStructStmt production.
	ExitVarDeclStructStmt(c *VarDeclStructStmtContext)

	// ExitVarStructDclStmt is called when exiting the VarStructDclStmt production.
	ExitVarStructDclStmt(c *VarStructDclStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitFunctionStmt is called when exiting the FunctionStmt production.
	ExitFunctionStmt(c *FunctionStmtContext)

	// ExitFunctionStructStmt is called when exiting the FunctionStructStmt production.
	ExitFunctionStructStmt(c *FunctionStructStmtContext)

	// ExitCallFunctionStmt is called when exiting the CallFunctionStmt production.
	ExitCallFunctionStmt(c *CallFunctionStmtContext)

	// ExitCallFunctionStructStmt is called when exiting the CallFunctionStructStmt production.
	ExitCallFunctionStructStmt(c *CallFunctionStructStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitPrintln is called when exiting the Println production.
	ExitPrintln(c *PrintlnContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

	// ExitIfOnly is called when exiting the IfOnly production.
	ExitIfOnly(c *IfOnlyContext)

	// ExitIfAnidado is called when exiting the IfAnidado production.
	ExitIfAnidado(c *IfAnidadoContext)

	// ExitSwitchStmt is called when exiting the SwitchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitCase is called when exiting the Case production.
	ExitCase(c *CaseContext)

	// ExitDefault is called when exiting the Default production.
	ExitDefault(c *DefaultContext)

	// ExitBlockStmt is called when exiting the blockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitForCondicion is called when exiting the ForCondicion production.
	ExitForCondicion(c *ForCondicionContext)

	// ExitForAsignacion is called when exiting the ForAsignacion production.
	ExitForAsignacion(c *ForAsignacionContext)

	// ExitForRange is called when exiting the ForRange production.
	ExitForRange(c *ForRangeContext)

	// ExitVarDclWithTypeAndValue is called when exiting the VarDclWithTypeAndValue production.
	ExitVarDclWithTypeAndValue(c *VarDclWithTypeAndValueContext)

	// ExitVarDclWithTypeOnly is called when exiting the VarDclWithTypeOnly production.
	ExitVarDclWithTypeOnly(c *VarDclWithTypeOnlyContext)

	// ExitVarDclWithInference is called when exiting the VarDclWithInference production.
	ExitVarDclWithInference(c *VarDclWithInferenceContext)

	// ExitSliceValores is called when exiting the SliceValores production.
	ExitSliceValores(c *SliceValoresContext)

	// ExitSliceVacio is called when exiting the SliceVacio production.
	ExitSliceVacio(c *SliceVacioContext)

	// ExitSliceDcl_Asign is called when exiting the SliceDcl_Asign production.
	ExitSliceDcl_Asign(c *SliceDcl_AsignContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitNuevoSlice is called when exiting the nuevoSlice production.
	ExitNuevoSlice(c *NuevoSliceContext)

	// ExitSliceContenido is called when exiting the SliceContenido production.
	ExitSliceContenido(c *SliceContenidoContext)

	// ExitSliceContenidoSlice is called when exiting the SliceContenidoSlice production.
	ExitSliceContenidoSlice(c *SliceContenidoSliceContext)

	// ExitDeclStructData is called when exiting the DeclStructData production.
	ExitDeclStructData(c *DeclStructDataContext)

	// ExitStructVarTypeInference is called when exiting the StructVarTypeInference production.
	ExitStructVarTypeInference(c *StructVarTypeInferenceContext)

	// ExitVarExpr is called when exiting the varExpr production.
	ExitVarExpr(c *VarExprContext)

	// ExitVarAdd is called when exiting the varAdd production.
	ExitVarAdd(c *VarAddContext)

	// ExitVarInc is called when exiting the varInc production.
	ExitVarInc(c *VarIncContext)

	// ExitArrayAccess is called when exiting the ArrayAccess production.
	ExitArrayAccess(c *ArrayAccessContext)

	// ExitStructAccessAsign is called when exiting the StructAccessAsign production.
	ExitStructAccessAsign(c *StructAccessAsignContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitCallFunctionValue is called when exiting the CallFunctionValue production.
	ExitCallFunctionValue(c *CallFunctionValueContext)

	// ExitLogical is called when exiting the Logical production.
	ExitLogical(c *LogicalContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitStructAccess is called when exiting the StructAccess production.
	ExitStructAccess(c *StructAccessContext)

	// ExitIdentifier is called when exiting the Identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitChar is called when exiting the Char production.
	ExitChar(c *CharContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitCallFunctionStructValue is called when exiting the CallFunctionStructValue production.
	ExitCallFunctionStructValue(c *CallFunctionStructValueContext)

	// ExitArrayFindIndex is called when exiting the ArrayFindIndex production.
	ExitArrayFindIndex(c *ArrayFindIndexContext)

	// ExitArrayAppend is called when exiting the ArrayAppend production.
	ExitArrayAppend(c *ArrayAppendContext)

	// ExitEqualsNotEquals is called when exiting the EqualsNotEquals production.
	ExitEqualsNotEquals(c *EqualsNotEqualsContext)

	// ExitIntToString is called when exiting the IntToString production.
	ExitIntToString(c *IntToStringContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitArrayAccessSimple is called when exiting the ArrayAccessSimple production.
	ExitArrayAccessSimple(c *ArrayAccessSimpleContext)

	// ExitArrayLength is called when exiting the ArrayLength production.
	ExitArrayLength(c *ArrayLengthContext)

	// ExitMulDivModulo is called when exiting the MulDivModulo production.
	ExitMulDivModulo(c *MulDivModuloContext)

	// ExitDouble is called when exiting the Double production.
	ExitDouble(c *DoubleContext)

	// ExitInteger is called when exiting the Integer production.
	ExitInteger(c *IntegerContext)

	// ExitNil is called when exiting the Nil production.
	ExitNil(c *NilContext)

	// ExitMinorMajorEqual is called when exiting the MinorMajorEqual production.
	ExitMinorMajorEqual(c *MinorMajorEqualContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitReflectType is called when exiting the reflectType production.
	ExitReflectType(c *ReflectTypeContext)

	// ExitNegate is called when exiting the Negate production.
	ExitNegate(c *NegateContext)

	// ExitArrayJoin is called when exiting the ArrayJoin production.
	ExitArrayJoin(c *ArrayJoinContext)

	// ExitFloatToString is called when exiting the floatToString production.
	ExitFloatToString(c *FloatToStringContext)

	// ExitPosicion is called when exiting the posicion production.
	ExitPosicion(c *PosicionContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitBreak is called when exiting the break production.
	ExitBreak(c *BreakContext)

	// ExitContinue is called when exiting the continue production.
	ExitContinue(c *ContinueContext)

	// ExitFunciones is called when exiting the Funciones production.
	ExitFunciones(c *FuncionesContext)

	// ExitFuncionesStructsNativas is called when exiting the FuncionesStructsNativas production.
	ExitFuncionesStructsNativas(c *FuncionesStructsNativasContext)

	// ExitDefParams is called when exiting the defParams production.
	ExitDefParams(c *DefParamsContext)

	// ExitCallFunction is called when exiting the CallFunction production.
	ExitCallFunction(c *CallFunctionContext)

	// ExitCallFunctionStruct is called when exiting the CallFunctionStruct production.
	ExitCallFunctionStruct(c *CallFunctionStructContext)

	// ExitValRet is called when exiting the valRet production.
	ExitValRet(c *ValRetContext)

	// ExitRetorno is called when exiting the retorno production.
	ExitRetorno(c *RetornoContext)
}
