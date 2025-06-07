// Generated from /home/myubuntu/Desktop/Compi2/OLC2_PROYECTO1_202202233/gramatica/gramatica.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link gramaticaParser}.
 */
public interface gramaticaListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#inicio}.
	 * @param ctx the parse tree
	 */
	void enterInicio(gramaticaParser.InicioContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#inicio}.
	 * @param ctx the parse tree
	 */
	void exitInicio(gramaticaParser.InicioContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PrintStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterPrintStmt(gramaticaParser.PrintStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PrintStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitPrintStmt(gramaticaParser.PrintStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IfStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterIfStmt(gramaticaParser.IfStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IfStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitIfStmt(gramaticaParser.IfStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SwitchInstruccion}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterSwitchInstruccion(gramaticaParser.SwitchInstruccionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SwitchInstruccion}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitSwitchInstruccion(gramaticaParser.SwitchInstruccionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SeccionInstruccion}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterSeccionInstruccion(gramaticaParser.SeccionInstruccionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SeccionInstruccion}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitSeccionInstruccion(gramaticaParser.SeccionInstruccionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterForStmt(gramaticaParser.ForStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitForStmt(gramaticaParser.ForStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDeclSliceStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclSliceStmt(gramaticaParser.VarDeclSliceStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDeclSliceStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclSliceStmt(gramaticaParser.VarDeclSliceStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterAsignStmt(gramaticaParser.AsignStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitAsignStmt(gramaticaParser.AsignStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDeclStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclStmt(gramaticaParser.VarDeclStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDeclStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclStmt(gramaticaParser.VarDeclStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDeclStructStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclStructStmt(gramaticaParser.VarDeclStructStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDeclStructStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclStructStmt(gramaticaParser.VarDeclStructStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarStructDclStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterVarStructDclStmt(gramaticaParser.VarStructDclStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarStructDclStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitVarStructDclStmt(gramaticaParser.VarStructDclStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BreakStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterBreakStmt(gramaticaParser.BreakStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BreakStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitBreakStmt(gramaticaParser.BreakStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ContinueStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterContinueStmt(gramaticaParser.ContinueStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ContinueStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitContinueStmt(gramaticaParser.ContinueStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FunctionStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterFunctionStmt(gramaticaParser.FunctionStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FunctionStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitFunctionStmt(gramaticaParser.FunctionStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FunctionStructStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterFunctionStructStmt(gramaticaParser.FunctionStructStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FunctionStructStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitFunctionStructStmt(gramaticaParser.FunctionStructStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CallFunctionStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterCallFunctionStmt(gramaticaParser.CallFunctionStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CallFunctionStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitCallFunctionStmt(gramaticaParser.CallFunctionStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CallFunctionStructStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterCallFunctionStructStmt(gramaticaParser.CallFunctionStructStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CallFunctionStructStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitCallFunctionStructStmt(gramaticaParser.CallFunctionStructStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ReturnStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterReturnStmt(gramaticaParser.ReturnStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ReturnStmt}
	 * labeled alternative in {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitReturnStmt(gramaticaParser.ReturnStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Println}
	 * labeled alternative in {@link gramaticaParser#imprimir}.
	 * @param ctx the parse tree
	 */
	void enterPrintln(gramaticaParser.PrintlnContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Println}
	 * labeled alternative in {@link gramaticaParser#imprimir}.
	 * @param ctx the parse tree
	 */
	void exitPrintln(gramaticaParser.PrintlnContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Print}
	 * labeled alternative in {@link gramaticaParser#imprimir}.
	 * @param ctx the parse tree
	 */
	void enterPrint(gramaticaParser.PrintContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Print}
	 * labeled alternative in {@link gramaticaParser#imprimir}.
	 * @param ctx the parse tree
	 */
	void exitPrint(gramaticaParser.PrintContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IfOnly}
	 * labeled alternative in {@link gramaticaParser#sIf}.
	 * @param ctx the parse tree
	 */
	void enterIfOnly(gramaticaParser.IfOnlyContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IfOnly}
	 * labeled alternative in {@link gramaticaParser#sIf}.
	 * @param ctx the parse tree
	 */
	void exitIfOnly(gramaticaParser.IfOnlyContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IfAnidado}
	 * labeled alternative in {@link gramaticaParser#sIf}.
	 * @param ctx the parse tree
	 */
	void enterIfAnidado(gramaticaParser.IfAnidadoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IfAnidado}
	 * labeled alternative in {@link gramaticaParser#sIf}.
	 * @param ctx the parse tree
	 */
	void exitIfAnidado(gramaticaParser.IfAnidadoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SwitchStmt}
	 * labeled alternative in {@link gramaticaParser#sSwitch}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmt(gramaticaParser.SwitchStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SwitchStmt}
	 * labeled alternative in {@link gramaticaParser#sSwitch}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmt(gramaticaParser.SwitchStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Case}
	 * labeled alternative in {@link gramaticaParser#cases}.
	 * @param ctx the parse tree
	 */
	void enterCase(gramaticaParser.CaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Case}
	 * labeled alternative in {@link gramaticaParser#cases}.
	 * @param ctx the parse tree
	 */
	void exitCase(gramaticaParser.CaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Default}
	 * labeled alternative in {@link gramaticaParser#cases}.
	 * @param ctx the parse tree
	 */
	void enterDefault(gramaticaParser.DefaultContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Default}
	 * labeled alternative in {@link gramaticaParser#cases}.
	 * @param ctx the parse tree
	 */
	void exitDefault(gramaticaParser.DefaultContext ctx);
	/**
	 * Enter a parse tree produced by the {@code blockStmt}
	 * labeled alternative in {@link gramaticaParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlockStmt(gramaticaParser.BlockStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code blockStmt}
	 * labeled alternative in {@link gramaticaParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlockStmt(gramaticaParser.BlockStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForCondicion}
	 * labeled alternative in {@link gramaticaParser#sFor}.
	 * @param ctx the parse tree
	 */
	void enterForCondicion(gramaticaParser.ForCondicionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForCondicion}
	 * labeled alternative in {@link gramaticaParser#sFor}.
	 * @param ctx the parse tree
	 */
	void exitForCondicion(gramaticaParser.ForCondicionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForAsignacion}
	 * labeled alternative in {@link gramaticaParser#sFor}.
	 * @param ctx the parse tree
	 */
	void enterForAsignacion(gramaticaParser.ForAsignacionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForAsignacion}
	 * labeled alternative in {@link gramaticaParser#sFor}.
	 * @param ctx the parse tree
	 */
	void exitForAsignacion(gramaticaParser.ForAsignacionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForRange}
	 * labeled alternative in {@link gramaticaParser#sFor}.
	 * @param ctx the parse tree
	 */
	void enterForRange(gramaticaParser.ForRangeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForRange}
	 * labeled alternative in {@link gramaticaParser#sFor}.
	 * @param ctx the parse tree
	 */
	void exitForRange(gramaticaParser.ForRangeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDclWithTypeAndValue}
	 * labeled alternative in {@link gramaticaParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterVarDclWithTypeAndValue(gramaticaParser.VarDclWithTypeAndValueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDclWithTypeAndValue}
	 * labeled alternative in {@link gramaticaParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitVarDclWithTypeAndValue(gramaticaParser.VarDclWithTypeAndValueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDclWithTypeOnly}
	 * labeled alternative in {@link gramaticaParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterVarDclWithTypeOnly(gramaticaParser.VarDclWithTypeOnlyContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDclWithTypeOnly}
	 * labeled alternative in {@link gramaticaParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitVarDclWithTypeOnly(gramaticaParser.VarDclWithTypeOnlyContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDclWithInference}
	 * labeled alternative in {@link gramaticaParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterVarDclWithInference(gramaticaParser.VarDclWithInferenceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDclWithInference}
	 * labeled alternative in {@link gramaticaParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitVarDclWithInference(gramaticaParser.VarDclWithInferenceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceValores}
	 * labeled alternative in {@link gramaticaParser#varDclSlice}.
	 * @param ctx the parse tree
	 */
	void enterSliceValores(gramaticaParser.SliceValoresContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceValores}
	 * labeled alternative in {@link gramaticaParser#varDclSlice}.
	 * @param ctx the parse tree
	 */
	void exitSliceValores(gramaticaParser.SliceValoresContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceVacio}
	 * labeled alternative in {@link gramaticaParser#varDclSlice}.
	 * @param ctx the parse tree
	 */
	void enterSliceVacio(gramaticaParser.SliceVacioContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceVacio}
	 * labeled alternative in {@link gramaticaParser#varDclSlice}.
	 * @param ctx the parse tree
	 */
	void exitSliceVacio(gramaticaParser.SliceVacioContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#assign}.
	 * @param ctx the parse tree
	 */
	void enterAssign(gramaticaParser.AssignContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#assign}.
	 * @param ctx the parse tree
	 */
	void exitAssign(gramaticaParser.AssignContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#nuevoSlice}.
	 * @param ctx the parse tree
	 */
	void enterNuevoSlice(gramaticaParser.NuevoSliceContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#nuevoSlice}.
	 * @param ctx the parse tree
	 */
	void exitNuevoSlice(gramaticaParser.NuevoSliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceContenido}
	 * labeled alternative in {@link gramaticaParser#contenidoSlice}.
	 * @param ctx the parse tree
	 */
	void enterSliceContenido(gramaticaParser.SliceContenidoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceContenido}
	 * labeled alternative in {@link gramaticaParser#contenidoSlice}.
	 * @param ctx the parse tree
	 */
	void exitSliceContenido(gramaticaParser.SliceContenidoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceContenidoSlice}
	 * labeled alternative in {@link gramaticaParser#contenidoSlice}.
	 * @param ctx the parse tree
	 */
	void enterSliceContenidoSlice(gramaticaParser.SliceContenidoSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceContenidoSlice}
	 * labeled alternative in {@link gramaticaParser#contenidoSlice}.
	 * @param ctx the parse tree
	 */
	void exitSliceContenidoSlice(gramaticaParser.SliceContenidoSliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclStructData}
	 * labeled alternative in {@link gramaticaParser#varDclStruct}.
	 * @param ctx the parse tree
	 */
	void enterDeclStructData(gramaticaParser.DeclStructDataContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclStructData}
	 * labeled alternative in {@link gramaticaParser#varDclStruct}.
	 * @param ctx the parse tree
	 */
	void exitDeclStructData(gramaticaParser.DeclStructDataContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructVarType}
	 * labeled alternative in {@link gramaticaParser#varStructDcl}.
	 * @param ctx the parse tree
	 */
	void enterStructVarType(gramaticaParser.StructVarTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructVarType}
	 * labeled alternative in {@link gramaticaParser#varStructDcl}.
	 * @param ctx the parse tree
	 */
	void exitStructVarType(gramaticaParser.StructVarTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructVarTypeInference}
	 * labeled alternative in {@link gramaticaParser#varStructDcl}.
	 * @param ctx the parse tree
	 */
	void enterStructVarTypeInference(gramaticaParser.StructVarTypeInferenceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructVarTypeInference}
	 * labeled alternative in {@link gramaticaParser#varStructDcl}.
	 * @param ctx the parse tree
	 */
	void exitStructVarTypeInference(gramaticaParser.StructVarTypeInferenceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code varExpr}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void enterVarExpr(gramaticaParser.VarExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code varExpr}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void exitVarExpr(gramaticaParser.VarExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code varAdd}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void enterVarAdd(gramaticaParser.VarAddContext ctx);
	/**
	 * Exit a parse tree produced by the {@code varAdd}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void exitVarAdd(gramaticaParser.VarAddContext ctx);
	/**
	 * Enter a parse tree produced by the {@code varInc}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void enterVarInc(gramaticaParser.VarIncContext ctx);
	/**
	 * Exit a parse tree produced by the {@code varInc}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void exitVarInc(gramaticaParser.VarIncContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArrayAccess}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void enterArrayAccess(gramaticaParser.ArrayAccessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArrayAccess}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void exitArrayAccess(gramaticaParser.ArrayAccessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructAccessAsign}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void enterStructAccessAsign(gramaticaParser.StructAccessAsignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructAccessAsign}
	 * labeled alternative in {@link gramaticaParser#varAsign}.
	 * @param ctx the parse tree
	 */
	void exitStructAccessAsign(gramaticaParser.StructAccessAsignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterParens(gramaticaParser.ParensContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitParens(gramaticaParser.ParensContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CallFunctionValue}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterCallFunctionValue(gramaticaParser.CallFunctionValueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CallFunctionValue}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitCallFunctionValue(gramaticaParser.CallFunctionValueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Logical}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterLogical(gramaticaParser.LogicalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Logical}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitLogical(gramaticaParser.LogicalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code String}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterString(gramaticaParser.StringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code String}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitString(gramaticaParser.StringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructAccess}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterStructAccess(gramaticaParser.StructAccessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructAccess}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitStructAccess(gramaticaParser.StructAccessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Identifier}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(gramaticaParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Identifier}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(gramaticaParser.IdentifierContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Char}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterChar(gramaticaParser.CharContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Char}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitChar(gramaticaParser.CharContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Boolean}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterBoolean(gramaticaParser.BooleanContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Boolean}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitBoolean(gramaticaParser.BooleanContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CallFunctionStructValue}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterCallFunctionStructValue(gramaticaParser.CallFunctionStructValueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CallFunctionStructValue}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitCallFunctionStructValue(gramaticaParser.CallFunctionStructValueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArrayFindIndex}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterArrayFindIndex(gramaticaParser.ArrayFindIndexContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArrayFindIndex}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitArrayFindIndex(gramaticaParser.ArrayFindIndexContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArrayAppend}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterArrayAppend(gramaticaParser.ArrayAppendContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArrayAppend}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitArrayAppend(gramaticaParser.ArrayAppendContext ctx);
	/**
	 * Enter a parse tree produced by the {@code EqualsNotEquals}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterEqualsNotEquals(gramaticaParser.EqualsNotEqualsContext ctx);
	/**
	 * Exit a parse tree produced by the {@code EqualsNotEquals}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitEqualsNotEquals(gramaticaParser.EqualsNotEqualsContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IntToString}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIntToString(gramaticaParser.IntToStringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntToString}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIntToString(gramaticaParser.IntToStringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(gramaticaParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(gramaticaParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArrayAccessSimple}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterArrayAccessSimple(gramaticaParser.ArrayAccessSimpleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArrayAccessSimple}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitArrayAccessSimple(gramaticaParser.ArrayAccessSimpleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArrayLength}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterArrayLength(gramaticaParser.ArrayLengthContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArrayLength}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitArrayLength(gramaticaParser.ArrayLengthContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDivModulo}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterMulDivModulo(gramaticaParser.MulDivModuloContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDivModulo}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitMulDivModulo(gramaticaParser.MulDivModuloContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Double}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterDouble(gramaticaParser.DoubleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Double}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitDouble(gramaticaParser.DoubleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Integer}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterInteger(gramaticaParser.IntegerContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Integer}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitInteger(gramaticaParser.IntegerContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Nil}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNil(gramaticaParser.NilContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Nil}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNil(gramaticaParser.NilContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MinorMajorEqual}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterMinorMajorEqual(gramaticaParser.MinorMajorEqualContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MinorMajorEqual}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitMinorMajorEqual(gramaticaParser.MinorMajorEqualContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Not}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNot(gramaticaParser.NotContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Not}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNot(gramaticaParser.NotContext ctx);
	/**
	 * Enter a parse tree produced by the {@code reflectType}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterReflectType(gramaticaParser.ReflectTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code reflectType}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitReflectType(gramaticaParser.ReflectTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Negate}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNegate(gramaticaParser.NegateContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Negate}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNegate(gramaticaParser.NegateContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArrayJoin}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterArrayJoin(gramaticaParser.ArrayJoinContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArrayJoin}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitArrayJoin(gramaticaParser.ArrayJoinContext ctx);
	/**
	 * Enter a parse tree produced by the {@code floatToString}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterFloatToString(gramaticaParser.FloatToStringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code floatToString}
	 * labeled alternative in {@link gramaticaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitFloatToString(gramaticaParser.FloatToStringContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#posicion}.
	 * @param ctx the parse tree
	 */
	void enterPosicion(gramaticaParser.PosicionContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#posicion}.
	 * @param ctx the parse tree
	 */
	void exitPosicion(gramaticaParser.PosicionContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(gramaticaParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(gramaticaParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#break}.
	 * @param ctx the parse tree
	 */
	void enterBreak(gramaticaParser.BreakContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#break}.
	 * @param ctx the parse tree
	 */
	void exitBreak(gramaticaParser.BreakContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#continue}.
	 * @param ctx the parse tree
	 */
	void enterContinue(gramaticaParser.ContinueContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#continue}.
	 * @param ctx the parse tree
	 */
	void exitContinue(gramaticaParser.ContinueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Funciones}
	 * labeled alternative in {@link gramaticaParser#functions}.
	 * @param ctx the parse tree
	 */
	void enterFunciones(gramaticaParser.FuncionesContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Funciones}
	 * labeled alternative in {@link gramaticaParser#functions}.
	 * @param ctx the parse tree
	 */
	void exitFunciones(gramaticaParser.FuncionesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionesStructsNativas}
	 * labeled alternative in {@link gramaticaParser#functionStruct}.
	 * @param ctx the parse tree
	 */
	void enterFuncionesStructsNativas(gramaticaParser.FuncionesStructsNativasContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionesStructsNativas}
	 * labeled alternative in {@link gramaticaParser#functionStruct}.
	 * @param ctx the parse tree
	 */
	void exitFuncionesStructsNativas(gramaticaParser.FuncionesStructsNativasContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#defParams}.
	 * @param ctx the parse tree
	 */
	void enterDefParams(gramaticaParser.DefParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#defParams}.
	 * @param ctx the parse tree
	 */
	void exitDefParams(gramaticaParser.DefParamsContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CallFunction}
	 * labeled alternative in {@link gramaticaParser#varCallStatement}.
	 * @param ctx the parse tree
	 */
	void enterCallFunction(gramaticaParser.CallFunctionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CallFunction}
	 * labeled alternative in {@link gramaticaParser#varCallStatement}.
	 * @param ctx the parse tree
	 */
	void exitCallFunction(gramaticaParser.CallFunctionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CallFunctionStruct}
	 * labeled alternative in {@link gramaticaParser#varCallFuncStruct}.
	 * @param ctx the parse tree
	 */
	void enterCallFunctionStruct(gramaticaParser.CallFunctionStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CallFunctionStruct}
	 * labeled alternative in {@link gramaticaParser#varCallFuncStruct}.
	 * @param ctx the parse tree
	 */
	void exitCallFunctionStruct(gramaticaParser.CallFunctionStructContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#valRet}.
	 * @param ctx the parse tree
	 */
	void enterValRet(gramaticaParser.ValRetContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#valRet}.
	 * @param ctx the parse tree
	 */
	void exitValRet(gramaticaParser.ValRetContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#retorno}.
	 * @param ctx the parse tree
	 */
	void enterRetorno(gramaticaParser.RetornoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#retorno}.
	 * @param ctx the parse tree
	 */
	void exitRetorno(gramaticaParser.RetornoContext ctx);
}