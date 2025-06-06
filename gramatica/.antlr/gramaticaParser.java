// Generated from /home/myubuntu/Desktop/Compi2/OLC2_PROYECTO1_202202233/gramatica/gramatica.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class gramaticaParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		T__45=46, T__46=47, T__47=48, T__48=49, T__49=50, T__50=51, T__51=52, 
		T__52=53, T__53=54, T__54=55, T__55=56, T__56=57, T__57=58, T__58=59, 
		T__59=60, T__60=61, INT=62, DOUBLE=63, CHAR=64, STRING=65, BOOL=66, BLANCOS=67, 
		ID_VARIABLE=68, COMENTARIOLINEA=69, COMENTARIOMULTILINEA=70;
	public static final int
		RULE_inicio = 0, RULE_instrucciones = 1, RULE_imprimir = 2, RULE_sIf = 3, 
		RULE_sSwitch = 4, RULE_cases = 5, RULE_block = 6, RULE_sFor = 7, RULE_varDcl = 8, 
		RULE_varDclSlice = 9, RULE_assign = 10, RULE_nuevoSlice = 11, RULE_contenidoSlice = 12, 
		RULE_varDclStruct = 13, RULE_varStructDcl = 14, RULE_varAsign = 15, RULE_expr = 16, 
		RULE_posicion = 17, RULE_type = 18, RULE_break = 19, RULE_continue = 20, 
		RULE_functions = 21, RULE_functionStruct = 22, RULE_defParams = 23, RULE_varCallStatement = 24, 
		RULE_varCallFuncStruct = 25, RULE_valRet = 26, RULE_retorno = 27;
	private static String[] makeRuleNames() {
		return new String[] {
			"inicio", "instrucciones", "imprimir", "sIf", "sSwitch", "cases", "block", 
			"sFor", "varDcl", "varDclSlice", "assign", "nuevoSlice", "contenidoSlice", 
			"varDclStruct", "varStructDcl", "varAsign", "expr", "posicion", "type", 
			"break", "continue", "functions", "functionStruct", "defParams", "varCallStatement", 
			"varCallFuncStruct", "valRet", "retorno"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "'println('", "','", "')'", "';'", "'print('", "'if'", 
			"'else'", "'switch'", "'case'", "':'", "'default:'", "'for'", "':='", 
			"'range'", "'mut'", "'='", "'var'", "'[]'", "'type'", "'struct'", "'+='", 
			"'-='", "'++'", "'--'", "'['", "']'", "'.'", "'-'", "'!'", "'*'", "'/'", 
			"'%'", "'+'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'", "'||'", 
			"'nil'", "'('", "'slices.Index('", "'strings.Join('", "'len('", "'append('", 
			"'strconv.Atoi('", "'strconv.ParseFloat('", "'reflect.TypeOf('", "'int'", 
			"'float64'", "'string'", "'bool'", "'rune'", "'break'", "'continue'", 
			"'func'", "'return'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, "INT", "DOUBLE", "CHAR", "STRING", "BOOL", "BLANCOS", "ID_VARIABLE", 
			"COMENTARIOLINEA", "COMENTARIOMULTILINEA"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "gramatica.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public gramaticaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InicioContext extends ParserRuleContext {
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public InicioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inicio; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterInicio(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitInicio(this);
		}
	}

	public final InicioContext inicio() throws RecognitionException {
		InicioContext _localctx = new InicioContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_inicio);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4323455642282640778L) != 0) || _la==ID_VARIABLE) {
				{
				{
				setState(56);
				instrucciones();
				}
				}
				setState(61);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstruccionesContext extends ParserRuleContext {
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	 
		public InstruccionesContext() { }
		public void copyFrom(InstruccionesContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SeccionInstruccionContext extends InstruccionesContext {
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public SeccionInstruccionContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterSeccionInstruccion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitSeccionInstruccion(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionStmtContext extends InstruccionesContext {
		public VarCallStatementContext varCallStatement() {
			return getRuleContext(VarCallStatementContext.class,0);
		}
		public CallFunctionStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterCallFunctionStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitCallFunctionStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintStmtContext extends InstruccionesContext {
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
		}
		public PrintStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterPrintStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitPrintStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarStructDclStmtContext extends InstruccionesContext {
		public VarStructDclContext varStructDcl() {
			return getRuleContext(VarStructDclContext.class,0);
		}
		public VarStructDclStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarStructDclStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarStructDclStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignStmtContext extends InstruccionesContext {
		public VarAsignContext varAsign() {
			return getRuleContext(VarAsignContext.class,0);
		}
		public AsignStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterAsignStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitAsignStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionStructStmtContext extends InstruccionesContext {
		public VarCallFuncStructContext varCallFuncStruct() {
			return getRuleContext(VarCallFuncStructContext.class,0);
		}
		public CallFunctionStructStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterCallFunctionStructStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitCallFunctionStructStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStmtContext extends InstruccionesContext {
		public ContinueContext continue_() {
			return getRuleContext(ContinueContext.class,0);
		}
		public ContinueStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterContinueStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitContinueStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDeclStructStmtContext extends InstruccionesContext {
		public VarDclStructContext varDclStruct() {
			return getRuleContext(VarDclStructContext.class,0);
		}
		public VarDeclStructStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarDeclStructStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarDeclStructStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStmtContext extends InstruccionesContext {
		public SIfContext sIf() {
			return getRuleContext(SIfContext.class,0);
		}
		public IfStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterIfStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitIfStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionStmtContext extends InstruccionesContext {
		public FunctionsContext functions() {
			return getRuleContext(FunctionsContext.class,0);
		}
		public FunctionStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterFunctionStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitFunctionStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionStructStmtContext extends InstruccionesContext {
		public FunctionStructContext functionStruct() {
			return getRuleContext(FunctionStructContext.class,0);
		}
		public FunctionStructStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterFunctionStructStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitFunctionStructStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDeclSliceStmtContext extends InstruccionesContext {
		public VarDclSliceContext varDclSlice() {
			return getRuleContext(VarDclSliceContext.class,0);
		}
		public VarDeclSliceStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarDeclSliceStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarDeclSliceStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDeclStmtContext extends InstruccionesContext {
		public VarDclContext varDcl() {
			return getRuleContext(VarDclContext.class,0);
		}
		public VarDeclStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarDeclStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarDeclStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends InstruccionesContext {
		public BreakContext break_() {
			return getRuleContext(BreakContext.class,0);
		}
		public BreakStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterBreakStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitBreakStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchInstruccionContext extends InstruccionesContext {
		public SSwitchContext sSwitch() {
			return getRuleContext(SSwitchContext.class,0);
		}
		public SwitchInstruccionContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterSwitchInstruccion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitSwitchInstruccion(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStmtContext extends InstruccionesContext {
		public SForContext sFor() {
			return getRuleContext(SForContext.class,0);
		}
		public ForStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterForStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitForStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends InstruccionesContext {
		public RetornoContext retorno() {
			return getRuleContext(RetornoContext.class,0);
		}
		public ReturnStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterReturnStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitReturnStmt(this);
		}
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);
		int _la;
		try {
			setState(86);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new PrintStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				imprimir();
				}
				break;
			case 2:
				_localctx = new IfStmtContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(63);
				sIf();
				}
				break;
			case 3:
				_localctx = new SwitchInstruccionContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(64);
				sSwitch();
				}
				break;
			case 4:
				_localctx = new SeccionInstruccionContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(65);
				match(T__0);
				setState(69);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4323455642282640778L) != 0) || _la==ID_VARIABLE) {
					{
					{
					setState(66);
					instrucciones();
					}
					}
					setState(71);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(72);
				match(T__1);
				}
				break;
			case 5:
				_localctx = new ForStmtContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(73);
				sFor();
				}
				break;
			case 6:
				_localctx = new VarDeclSliceStmtContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(74);
				varDclSlice();
				}
				break;
			case 7:
				_localctx = new AsignStmtContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(75);
				varAsign();
				}
				break;
			case 8:
				_localctx = new VarDeclStmtContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(76);
				varDcl();
				}
				break;
			case 9:
				_localctx = new VarDeclStructStmtContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(77);
				varDclStruct();
				}
				break;
			case 10:
				_localctx = new VarStructDclStmtContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(78);
				varStructDcl();
				}
				break;
			case 11:
				_localctx = new BreakStmtContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(79);
				break_();
				}
				break;
			case 12:
				_localctx = new ContinueStmtContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(80);
				continue_();
				}
				break;
			case 13:
				_localctx = new FunctionStmtContext(_localctx);
				enterOuterAlt(_localctx, 13);
				{
				setState(81);
				functions();
				}
				break;
			case 14:
				_localctx = new FunctionStructStmtContext(_localctx);
				enterOuterAlt(_localctx, 14);
				{
				setState(82);
				functionStruct();
				}
				break;
			case 15:
				_localctx = new CallFunctionStmtContext(_localctx);
				enterOuterAlt(_localctx, 15);
				{
				setState(83);
				varCallStatement();
				}
				break;
			case 16:
				_localctx = new CallFunctionStructStmtContext(_localctx);
				enterOuterAlt(_localctx, 16);
				{
				setState(84);
				varCallFuncStruct();
				}
				break;
			case 17:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 17);
				{
				setState(85);
				retorno();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ImprimirContext extends ParserRuleContext {
		public ImprimirContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imprimir; }
	 
		public ImprimirContext() { }
		public void copyFrom(ImprimirContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintContext extends ImprimirContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public PrintContext(ImprimirContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterPrint(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitPrint(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintlnContext extends ImprimirContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public PrintlnContext(ImprimirContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterPrintln(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitPrintln(this);
		}
	}

	public final ImprimirContext imprimir() throws RecognitionException {
		ImprimirContext _localctx = new ImprimirContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_imprimir);
		int _la;
		try {
			setState(118);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				_localctx = new PrintlnContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(88);
				match(T__2);
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 30)) & ~0x3f) == 0 && ((1L << (_la - 30)) & 408030265347L) != 0)) {
					{
					setState(89);
					expr(0);
					setState(94);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__3) {
						{
						{
						setState(90);
						match(T__3);
						setState(91);
						expr(0);
						}
						}
						setState(96);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(99);
				match(T__4);
				setState(101);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__5) {
					{
					setState(100);
					match(T__5);
					}
				}

				}
				break;
			case T__6:
				_localctx = new PrintContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				match(T__6);
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 30)) & ~0x3f) == 0 && ((1L << (_la - 30)) & 408030265347L) != 0)) {
					{
					setState(104);
					expr(0);
					setState(109);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__3) {
						{
						{
						setState(105);
						match(T__3);
						setState(106);
						expr(0);
						}
						}
						setState(111);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(114);
				match(T__4);
				setState(116);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__5) {
					{
					setState(115);
					match(T__5);
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SIfContext extends ParserRuleContext {
		public SIfContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sIf; }
	 
		public SIfContext() { }
		public void copyFrom(SIfContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfAnidadoContext extends SIfContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public SIfContext sIf() {
			return getRuleContext(SIfContext.class,0);
		}
		public IfAnidadoContext(SIfContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterIfAnidado(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitIfAnidado(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfOnlyContext extends SIfContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public IfOnlyContext(SIfContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterIfOnly(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitIfOnly(this);
		}
	}

	public final SIfContext sIf() throws RecognitionException {
		SIfContext _localctx = new SIfContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_sIf);
		int _la;
		try {
			setState(133);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				_localctx = new IfOnlyContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(T__7);
				setState(121);
				expr(0);
				setState(122);
				block();
				setState(125);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__8) {
					{
					setState(123);
					match(T__8);
					setState(124);
					block();
					}
				}

				}
				break;
			case 2:
				_localctx = new IfAnidadoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(127);
				match(T__7);
				setState(128);
				expr(0);
				setState(129);
				block();
				setState(130);
				match(T__8);
				setState(131);
				sIf();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SSwitchContext extends ParserRuleContext {
		public SSwitchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sSwitch; }
	 
		public SSwitchContext() { }
		public void copyFrom(SSwitchContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtContext extends SSwitchContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public CasesContext cases() {
			return getRuleContext(CasesContext.class,0);
		}
		public SwitchStmtContext(SSwitchContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterSwitchStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitSwitchStmt(this);
		}
	}

	public final SSwitchContext sSwitch() throws RecognitionException {
		SSwitchContext _localctx = new SSwitchContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_sSwitch);
		try {
			_localctx = new SwitchStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(T__9);
			setState(136);
			expr(0);
			setState(137);
			match(T__0);
			setState(138);
			cases();
			setState(139);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CasesContext extends ParserRuleContext {
		public CasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cases; }
	 
		public CasesContext() { }
		public void copyFrom(CasesContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefaultContext extends CasesContext {
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public DefaultContext(CasesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterDefault(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitDefault(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CaseContext extends CasesContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public CasesContext cases() {
			return getRuleContext(CasesContext.class,0);
		}
		public CaseContext(CasesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterCase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitCase(this);
		}
	}

	public final CasesContext cases() throws RecognitionException {
		CasesContext _localctx = new CasesContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_cases);
		int _la;
		try {
			setState(160);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__10:
				_localctx = new CaseContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(141);
				match(T__10);
				setState(142);
				expr(0);
				setState(143);
				match(T__11);
				setState(147);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4323455642282640778L) != 0) || _la==ID_VARIABLE) {
					{
					{
					setState(144);
					instrucciones();
					}
					}
					setState(149);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__10 || _la==T__12) {
					{
					setState(150);
					cases();
					}
				}

				}
				break;
			case T__12:
				_localctx = new DefaultContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(153);
				match(T__12);
				setState(157);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4323455642282640778L) != 0) || _la==ID_VARIABLE) {
					{
					{
					setState(154);
					instrucciones();
					}
					}
					setState(159);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	 
		public BlockContext() { }
		public void copyFrom(BlockContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BlockStmtContext extends BlockContext {
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public BlockStmtContext(BlockContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterBlockStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitBlockStmt(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_block);
		int _la;
		try {
			_localctx = new BlockStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(T__0);
			setState(166);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4323455642282640778L) != 0) || _la==ID_VARIABLE) {
				{
				{
				setState(163);
				instrucciones();
				}
				}
				setState(168);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(169);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SForContext extends ParserRuleContext {
		public SForContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sFor; }
	 
		public SForContext() { }
		public void copyFrom(SForContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForCondicionContext extends SForContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForCondicionContext(SForContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterForCondicion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitForCondicion(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForRangeContext extends SForContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForRangeContext(SForContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterForRange(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitForRange(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForAsignacionContext extends SForContext {
		public VarDclContext varDcl() {
			return getRuleContext(VarDclContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarAsignContext varAsign() {
			return getRuleContext(VarAsignContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForAsignacionContext(SForContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterForAsignacion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitForAsignacion(this);
		}
	}

	public final SForContext sFor() throws RecognitionException {
		SForContext _localctx = new SForContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_sFor);
		try {
			setState(191);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				_localctx = new ForCondicionContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(171);
				match(T__13);
				setState(172);
				expr(0);
				setState(173);
				block();
				}
				break;
			case 2:
				_localctx = new ForAsignacionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(175);
				match(T__13);
				setState(176);
				varDcl();
				setState(177);
				match(T__5);
				setState(178);
				expr(0);
				setState(179);
				match(T__5);
				setState(180);
				varAsign();
				setState(181);
				block();
				}
				break;
			case 3:
				_localctx = new ForRangeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(183);
				match(T__13);
				setState(184);
				match(ID_VARIABLE);
				setState(185);
				match(T__3);
				setState(186);
				match(ID_VARIABLE);
				setState(187);
				match(T__14);
				setState(188);
				match(T__15);
				setState(189);
				match(ID_VARIABLE);
				setState(190);
				block();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarDclContext extends ParserRuleContext {
		public VarDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDcl; }
	 
		public VarDclContext() { }
		public void copyFrom(VarDclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDclWithTypeAndValueContext extends VarDclContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarDclWithTypeAndValueContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarDclWithTypeAndValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarDclWithTypeAndValue(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDclWithInferenceContext extends VarDclContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarDclWithInferenceContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarDclWithInference(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarDclWithInference(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDclWithTypeOnlyContext extends VarDclContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public VarDclWithTypeOnlyContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarDclWithTypeOnly(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarDclWithTypeOnly(this);
		}
	}

	public final VarDclContext varDcl() throws RecognitionException {
		VarDclContext _localctx = new VarDclContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_varDcl);
		try {
			setState(206);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				_localctx = new VarDclWithTypeAndValueContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				match(T__16);
				setState(194);
				match(ID_VARIABLE);
				setState(195);
				type();
				setState(196);
				match(T__17);
				setState(197);
				expr(0);
				}
				break;
			case 2:
				_localctx = new VarDclWithTypeOnlyContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(199);
				match(T__16);
				setState(200);
				match(ID_VARIABLE);
				setState(201);
				type();
				}
				break;
			case 3:
				_localctx = new VarDclWithInferenceContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(202);
				match(T__16);
				setState(203);
				match(ID_VARIABLE);
				setState(204);
				match(T__14);
				setState(205);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarDclSliceContext extends ParserRuleContext {
		public VarDclSliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDclSlice; }
	 
		public VarDclSliceContext() { }
		public void copyFrom(VarDclSliceContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceValoresContext extends VarDclSliceContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public AssignContext assign() {
			return getRuleContext(AssignContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ContenidoSliceContext contenidoSlice() {
			return getRuleContext(ContenidoSliceContext.class,0);
		}
		public List<NuevoSliceContext> nuevoSlice() {
			return getRuleContexts(NuevoSliceContext.class);
		}
		public NuevoSliceContext nuevoSlice(int i) {
			return getRuleContext(NuevoSliceContext.class,i);
		}
		public SliceValoresContext(VarDclSliceContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterSliceValores(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitSliceValores(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceVacioContext extends VarDclSliceContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public List<NuevoSliceContext> nuevoSlice() {
			return getRuleContexts(NuevoSliceContext.class);
		}
		public NuevoSliceContext nuevoSlice(int i) {
			return getRuleContext(NuevoSliceContext.class,i);
		}
		public SliceVacioContext(VarDclSliceContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterSliceVacio(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitSliceVacio(this);
		}
	}

	public final VarDclSliceContext varDclSlice() throws RecognitionException {
		VarDclSliceContext _localctx = new VarDclSliceContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_varDclSlice);
		int _la;
		try {
			setState(229);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID_VARIABLE:
				_localctx = new SliceValoresContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(208);
				match(ID_VARIABLE);
				setState(209);
				assign();
				setState(211); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(210);
					nuevoSlice();
					}
					}
					setState(213); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__19 );
				setState(215);
				type();
				setState(216);
				match(T__0);
				setState(217);
				contenidoSlice();
				setState(218);
				match(T__1);
				}
				break;
			case T__18:
				_localctx = new SliceVacioContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(220);
				match(T__18);
				setState(221);
				match(ID_VARIABLE);
				setState(223); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(222);
					nuevoSlice();
					}
					}
					setState(225); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__19 );
				setState(227);
				type();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignContext extends ParserRuleContext {
		public AssignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterAssign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitAssign(this);
		}
	}

	public final AssignContext assign() throws RecognitionException {
		AssignContext _localctx = new AssignContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_assign);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			_la = _input.LA(1);
			if ( !(_la==T__14 || _la==T__17) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NuevoSliceContext extends ParserRuleContext {
		public NuevoSliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nuevoSlice; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterNuevoSlice(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitNuevoSlice(this);
		}
	}

	public final NuevoSliceContext nuevoSlice() throws RecognitionException {
		NuevoSliceContext _localctx = new NuevoSliceContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_nuevoSlice);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(233);
			match(T__19);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ContenidoSliceContext extends ParserRuleContext {
		public ContenidoSliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_contenidoSlice; }
	 
		public ContenidoSliceContext() { }
		public void copyFrom(ContenidoSliceContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceContenidoSliceContext extends ContenidoSliceContext {
		public List<ContenidoSliceContext> contenidoSlice() {
			return getRuleContexts(ContenidoSliceContext.class);
		}
		public ContenidoSliceContext contenidoSlice(int i) {
			return getRuleContext(ContenidoSliceContext.class,i);
		}
		public SliceContenidoSliceContext(ContenidoSliceContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterSliceContenidoSlice(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitSliceContenidoSlice(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceContenidoContext extends ContenidoSliceContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public SliceContenidoContext(ContenidoSliceContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterSliceContenido(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitSliceContenido(this);
		}
	}

	public final ContenidoSliceContext contenidoSlice() throws RecognitionException {
		ContenidoSliceContext _localctx = new ContenidoSliceContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_contenidoSlice);
		int _la;
		try {
			setState(258);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__29:
			case T__30:
			case T__43:
			case T__44:
			case T__45:
			case T__46:
			case T__47:
			case T__48:
			case T__49:
			case T__50:
			case T__51:
			case INT:
			case DOUBLE:
			case CHAR:
			case STRING:
			case BOOL:
			case ID_VARIABLE:
				_localctx = new SliceContenidoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(235);
				expr(0);
				setState(240);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(236);
					match(T__3);
					setState(237);
					expr(0);
					}
					}
					setState(242);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case T__0:
				_localctx = new SliceContenidoSliceContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(243);
				match(T__0);
				setState(244);
				contenidoSlice();
				setState(245);
				match(T__1);
				setState(255);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(246);
					match(T__3);
					setState(251);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==T__0) {
						{
						setState(247);
						match(T__0);
						setState(248);
						contenidoSlice();
						setState(249);
						match(T__1);
						}
					}

					}
					}
					setState(257);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarDclStructContext extends ParserRuleContext {
		public VarDclStructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDclStruct; }
	 
		public VarDclStructContext() { }
		public void copyFrom(VarDclStructContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclStructDataContext extends VarDclStructContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public DeclStructDataContext(VarDclStructContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterDeclStructData(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitDeclStructData(this);
		}
	}

	public final VarDclStructContext varDclStruct() throws RecognitionException {
		VarDclStructContext _localctx = new VarDclStructContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_varDclStruct);
		int _la;
		try {
			_localctx = new DeclStructDataContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__20) {
				{
				setState(260);
				match(T__20);
				}
			}

			setState(264);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__21) {
				{
				setState(263);
				match(T__21);
				}
			}

			setState(266);
			match(ID_VARIABLE);
			setState(268);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__21) {
				{
				setState(267);
				match(T__21);
				}
			}

			setState(270);
			match(T__0);
			setState(276); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(271);
				match(ID_VARIABLE);
				setState(272);
				type();
				setState(274);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__5) {
					{
					setState(273);
					match(T__5);
					}
				}

				}
				}
				setState(278); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID_VARIABLE );
			setState(280);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarStructDclContext extends ParserRuleContext {
		public VarStructDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varStructDcl; }
	 
		public VarStructDclContext() { }
		public void copyFrom(VarStructDclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructVarTypeContext extends VarStructDclContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public StructVarTypeContext(VarStructDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterStructVarType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitStructVarType(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructVarTypeInferenceContext extends VarStructDclContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public StructVarTypeInferenceContext(VarStructDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterStructVarTypeInference(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitStructVarTypeInference(this);
		}
	}

	public final VarStructDclContext varStructDcl() throws RecognitionException {
		VarStructDclContext _localctx = new VarStructDclContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_varStructDcl);
		int _la;
		try {
			setState(318);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				_localctx = new StructVarTypeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(282);
				match(ID_VARIABLE);
				setState(283);
				match(ID_VARIABLE);
				setState(284);
				match(T__17);
				setState(285);
				match(T__0);
				setState(286);
				match(ID_VARIABLE);
				setState(287);
				match(T__11);
				setState(288);
				expr(0);
				setState(293); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(289);
					match(T__3);
					setState(290);
					match(ID_VARIABLE);
					setState(291);
					match(T__11);
					setState(292);
					expr(0);
					}
					}
					setState(295); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__3 );
				setState(297);
				match(T__1);
				}
				break;
			case 2:
				_localctx = new StructVarTypeInferenceContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(299);
				match(ID_VARIABLE);
				setState(300);
				match(T__14);
				setState(301);
				match(ID_VARIABLE);
				setState(302);
				match(T__0);
				setState(303);
				match(ID_VARIABLE);
				setState(304);
				match(T__11);
				setState(305);
				expr(0);
				setState(312); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(306);
					match(T__3);
					setState(310);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==ID_VARIABLE) {
						{
						setState(307);
						match(ID_VARIABLE);
						setState(308);
						match(T__11);
						setState(309);
						expr(0);
						}
					}

					}
					}
					setState(314); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__3 );
				setState(316);
				match(T__1);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarAsignContext extends ParserRuleContext {
		public VarAsignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varAsign; }
	 
		public VarAsignContext() { }
		public void copyFrom(VarAsignContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarIncContext extends VarAsignContext {
		public Token op;
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public VarIncContext(VarAsignContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarInc(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarInc(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarExprContext extends VarAsignContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarExprContext(VarAsignContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarExpr(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayAccessContext extends VarAsignContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ArrayAccessContext(VarAsignContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterArrayAccess(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitArrayAccess(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAccessAsignContext extends VarAsignContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StructAccessAsignContext(VarAsignContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterStructAccessAsign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitStructAccessAsign(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarAddContext extends VarAsignContext {
		public Token op;
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarAddContext(VarAsignContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterVarAdd(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitVarAdd(this);
		}
	}

	public final VarAsignContext varAsign() throws RecognitionException {
		VarAsignContext _localctx = new VarAsignContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_varAsign);
		int _la;
		try {
			setState(349);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				_localctx = new VarExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(320);
				match(ID_VARIABLE);
				setState(321);
				match(T__17);
				setState(322);
				expr(0);
				}
				break;
			case 2:
				_localctx = new VarAddContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(323);
				match(ID_VARIABLE);
				setState(324);
				((VarAddContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__22 || _la==T__23) ) {
					((VarAddContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(325);
				expr(0);
				}
				break;
			case 3:
				_localctx = new VarIncContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(326);
				match(ID_VARIABLE);
				setState(327);
				((VarIncContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__24 || _la==T__25) ) {
					((VarIncContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 4:
				_localctx = new ArrayAccessContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(328);
				match(ID_VARIABLE);
				setState(333); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(329);
					match(T__26);
					setState(330);
					expr(0);
					setState(331);
					match(T__27);
					}
					}
					setState(335); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__26 );
				setState(337);
				match(T__17);
				setState(338);
				expr(0);
				}
				break;
			case 5:
				_localctx = new StructAccessAsignContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(340);
				match(ID_VARIABLE);
				setState(343); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(341);
					match(T__28);
					setState(342);
					match(ID_VARIABLE);
					}
					}
					setState(345); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__28 );
				setState(347);
				match(T__17);
				setState(348);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParensContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ParensContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterParens(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitParens(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionValueContext extends ExprContext {
		public VarCallStatementContext varCallStatement() {
			return getRuleContext(VarCallStatementContext.class,0);
		}
		public CallFunctionValueContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterCallFunctionValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitCallFunctionValue(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LogicalContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LogicalContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterLogical(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitLogical(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringContext extends ExprContext {
		public TerminalNode STRING() { return getToken(gramaticaParser.STRING, 0); }
		public StringContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterString(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitString(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAccessContext extends ExprContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public StructAccessContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterStructAccess(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitStructAccess(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public IdentifierContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitIdentifier(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CharContext extends ExprContext {
		public TerminalNode CHAR() { return getToken(gramaticaParser.CHAR, 0); }
		public CharContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterChar(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitChar(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BooleanContext extends ExprContext {
		public TerminalNode BOOL() { return getToken(gramaticaParser.BOOL, 0); }
		public BooleanContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterBoolean(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitBoolean(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionStructValueContext extends ExprContext {
		public VarCallFuncStructContext varCallFuncStruct() {
			return getRuleContext(VarCallFuncStructContext.class,0);
		}
		public CallFunctionStructValueContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterCallFunctionStructValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitCallFunctionStructValue(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayFindIndexContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ArrayFindIndexContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterArrayFindIndex(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitArrayFindIndex(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayAppendContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ArrayAppendContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterArrayAppend(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitArrayAppend(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EqualsNotEqualsContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public EqualsNotEqualsContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterEqualsNotEquals(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitEqualsNotEquals(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntToStringContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IntToStringContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterIntToString(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitIntToString(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AddSubContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public AddSubContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterAddSub(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitAddSub(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayAccessSimpleContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ArrayAccessSimpleContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterArrayAccessSimple(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitArrayAccessSimple(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayLengthContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public List<PosicionContext> posicion() {
			return getRuleContexts(PosicionContext.class);
		}
		public PosicionContext posicion(int i) {
			return getRuleContext(PosicionContext.class,i);
		}
		public ArrayLengthContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterArrayLength(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitArrayLength(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MulDivModuloContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public MulDivModuloContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterMulDivModulo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitMulDivModulo(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DoubleContext extends ExprContext {
		public TerminalNode DOUBLE() { return getToken(gramaticaParser.DOUBLE, 0); }
		public DoubleContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterDouble(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitDouble(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntegerContext extends ExprContext {
		public TerminalNode INT() { return getToken(gramaticaParser.INT, 0); }
		public IntegerContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterInteger(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitInteger(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NilContext extends ExprContext {
		public NilContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterNil(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitNil(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MinorMajorEqualContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public MinorMajorEqualContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterMinorMajorEqual(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitMinorMajorEqual(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NotContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NotContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterNot(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitNot(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReflectTypeContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReflectTypeContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterReflectType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitReflectType(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NegateContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NegateContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterNegate(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitNegate(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayJoinContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ArrayJoinContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterArrayJoin(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitArrayJoin(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatToStringContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public FloatToStringContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterFloatToString(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitFloatToString(this);
		}
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(427);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
			case 1:
				{
				_localctx = new NegateContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(352);
				match(T__29);
				setState(353);
				expr(26);
				}
				break;
			case 2:
				{
				_localctx = new NotContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(354);
				match(T__30);
				setState(355);
				expr(25);
				}
				break;
			case 3:
				{
				_localctx = new IntegerContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(356);
				match(INT);
				}
				break;
			case 4:
				{
				_localctx = new DoubleContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(357);
				match(DOUBLE);
				}
				break;
			case 5:
				{
				_localctx = new StringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(358);
				match(STRING);
				}
				break;
			case 6:
				{
				_localctx = new BooleanContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(359);
				match(BOOL);
				}
				break;
			case 7:
				{
				_localctx = new IdentifierContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(360);
				match(ID_VARIABLE);
				}
				break;
			case 8:
				{
				_localctx = new CharContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(361);
				match(CHAR);
				}
				break;
			case 9:
				{
				_localctx = new NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(362);
				match(T__43);
				}
				break;
			case 10:
				{
				_localctx = new ParensContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(363);
				match(T__44);
				setState(364);
				expr(0);
				setState(365);
				match(T__4);
				}
				break;
			case 11:
				{
				_localctx = new ArrayAccessSimpleContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(367);
				match(ID_VARIABLE);
				setState(372); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(368);
						match(T__26);
						setState(369);
						expr(0);
						setState(370);
						match(T__27);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(374); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,38,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 12:
				{
				_localctx = new ArrayFindIndexContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(376);
				match(T__45);
				setState(377);
				match(ID_VARIABLE);
				setState(378);
				match(T__3);
				setState(379);
				expr(0);
				setState(380);
				match(T__4);
				}
				break;
			case 13:
				{
				_localctx = new ArrayJoinContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(382);
				match(T__46);
				setState(383);
				match(ID_VARIABLE);
				setState(384);
				match(T__3);
				setState(385);
				expr(0);
				setState(386);
				match(T__4);
				}
				break;
			case 14:
				{
				_localctx = new ArrayLengthContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(388);
				match(T__47);
				setState(389);
				match(ID_VARIABLE);
				setState(393);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__26) {
					{
					{
					setState(390);
					posicion();
					}
					}
					setState(395);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(396);
				match(T__4);
				}
				break;
			case 15:
				{
				_localctx = new ArrayAppendContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(397);
				match(T__48);
				setState(398);
				match(ID_VARIABLE);
				setState(399);
				match(T__3);
				setState(400);
				expr(0);
				setState(401);
				match(T__4);
				}
				break;
			case 16:
				{
				_localctx = new IntToStringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(403);
				match(T__49);
				setState(404);
				expr(0);
				setState(405);
				match(T__4);
				}
				break;
			case 17:
				{
				_localctx = new FloatToStringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(407);
				match(T__50);
				setState(408);
				expr(0);
				setState(409);
				match(T__4);
				}
				break;
			case 18:
				{
				_localctx = new ReflectTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(411);
				match(T__51);
				setState(412);
				expr(0);
				setState(413);
				match(T__4);
				}
				break;
			case 19:
				{
				_localctx = new StructAccessContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(415);
				match(ID_VARIABLE);
				setState(418); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(416);
						match(T__28);
						setState(417);
						match(ID_VARIABLE);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(420); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(423);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
				case 1:
					{
					setState(422);
					match(T__5);
					}
					break;
				}
				}
				break;
			case 20:
				{
				_localctx = new CallFunctionValueContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(425);
				varCallStatement();
				}
				break;
			case 21:
				{
				_localctx = new CallFunctionStructValueContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(426);
				varCallFuncStruct();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(446);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,44,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(444);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
					case 1:
						{
						_localctx = new MulDivModuloContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(429);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(430);
						((MulDivModuloContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 30064771072L) != 0)) ) {
							((MulDivModuloContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(431);
						expr(25);
						}
						break;
					case 2:
						{
						_localctx = new AddSubContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(432);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(433);
						((AddSubContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__29 || _la==T__34) ) {
							((AddSubContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(434);
						expr(24);
						}
						break;
					case 3:
						{
						_localctx = new MinorMajorEqualContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(435);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(436);
						((MinorMajorEqualContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1030792151040L) != 0)) ) {
							((MinorMajorEqualContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(437);
						expr(23);
						}
						break;
					case 4:
						{
						_localctx = new EqualsNotEqualsContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(438);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(439);
						((EqualsNotEqualsContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__39 || _la==T__40) ) {
							((EqualsNotEqualsContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(440);
						expr(22);
						}
						break;
					case 5:
						{
						_localctx = new LogicalContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(441);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(442);
						((LogicalContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__41 || _la==T__42) ) {
							((LogicalContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(443);
						expr(21);
						}
						break;
					}
					} 
				}
				setState(448);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,44,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PosicionContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public PosicionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_posicion; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterPosicion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitPosicion(this);
		}
	}

	public final PosicionContext posicion() throws RecognitionException {
		PosicionContext _localctx = new PosicionContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_posicion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(449);
			match(T__26);
			setState(450);
			expr(0);
			setState(451);
			match(T__27);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitType(this);
		}
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(453);
			_la = _input.LA(1);
			if ( !(((((_la - 22)) & ~0x3f) == 0 && ((1L << (_la - 22)) & 70435316170753L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BreakContext extends ParserRuleContext {
		public BreakContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterBreak(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitBreak(this);
		}
	}

	public final BreakContext break_() throws RecognitionException {
		BreakContext _localctx = new BreakContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_break);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(455);
			match(T__57);
			setState(457);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(456);
				match(T__5);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ContinueContext extends ParserRuleContext {
		public ContinueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterContinue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitContinue(this);
		}
	}

	public final ContinueContext continue_() throws RecognitionException {
		ContinueContext _localctx = new ContinueContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_continue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(459);
			match(T__58);
			setState(461);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(460);
				match(T__5);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionsContext extends ParserRuleContext {
		public FunctionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functions; }
	 
		public FunctionsContext() { }
		public void copyFrom(FunctionsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionesContext extends FunctionsContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public ValRetContext valRet() {
			return getRuleContext(ValRetContext.class,0);
		}
		public FuncionesContext(FunctionsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterFunciones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitFunciones(this);
		}
	}

	public final FunctionsContext functions() throws RecognitionException {
		FunctionsContext _localctx = new FunctionsContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_functions);
		int _la;
		try {
			_localctx = new FuncionesContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(463);
			match(T__59);
			setState(464);
			match(ID_VARIABLE);
			setState(465);
			match(T__44);
			setState(476);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID_VARIABLE) {
				{
				setState(466);
				match(ID_VARIABLE);
				setState(467);
				type();
				setState(473);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(468);
					match(T__3);
					setState(469);
					match(ID_VARIABLE);
					setState(470);
					type();
					}
					}
					setState(475);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(478);
			match(T__4);
			setState(480);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 22)) & ~0x3f) == 0 && ((1L << (_la - 22)) & 70435316170753L) != 0)) {
				{
				setState(479);
				valRet();
				}
			}

			setState(482);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionStructContext extends ParserRuleContext {
		public FunctionStructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionStruct; }
	 
		public FunctionStructContext() { }
		public void copyFrom(FunctionStructContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionesStructsNativasContext extends FunctionStructContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public DefParamsContext defParams() {
			return getRuleContext(DefParamsContext.class,0);
		}
		public ValRetContext valRet() {
			return getRuleContext(ValRetContext.class,0);
		}
		public FuncionesStructsNativasContext(FunctionStructContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterFuncionesStructsNativas(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitFuncionesStructsNativas(this);
		}
	}

	public final FunctionStructContext functionStruct() throws RecognitionException {
		FunctionStructContext _localctx = new FunctionStructContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_functionStruct);
		int _la;
		try {
			_localctx = new FuncionesStructsNativasContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(484);
			match(T__59);
			setState(485);
			match(T__44);
			setState(486);
			match(ID_VARIABLE);
			setState(487);
			match(ID_VARIABLE);
			setState(488);
			match(T__4);
			setState(489);
			match(ID_VARIABLE);
			setState(490);
			match(T__44);
			setState(492);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID_VARIABLE) {
				{
				setState(491);
				defParams();
				}
			}

			setState(494);
			match(T__4);
			setState(496);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 22)) & ~0x3f) == 0 && ((1L << (_la - 22)) & 70435316170753L) != 0)) {
				{
				setState(495);
				valRet();
				}
			}

			setState(498);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DefParamsContext extends ParserRuleContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public DefParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defParams; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterDefParams(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitDefParams(this);
		}
	}

	public final DefParamsContext defParams() throws RecognitionException {
		DefParamsContext _localctx = new DefParamsContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_defParams);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(500);
			match(ID_VARIABLE);
			setState(501);
			type();
			setState(507);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(502);
				match(T__3);
				setState(503);
				match(ID_VARIABLE);
				setState(504);
				type();
				}
				}
				setState(509);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarCallStatementContext extends ParserRuleContext {
		public VarCallStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varCallStatement; }
	 
		public VarCallStatementContext() { }
		public void copyFrom(VarCallStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionContext extends VarCallStatementContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public CallFunctionContext(VarCallStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterCallFunction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitCallFunction(this);
		}
	}

	public final VarCallStatementContext varCallStatement() throws RecognitionException {
		VarCallStatementContext _localctx = new VarCallStatementContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_varCallStatement);
		int _la;
		try {
			_localctx = new CallFunctionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(510);
			match(ID_VARIABLE);
			setState(511);
			match(T__44);
			setState(520);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 30)) & ~0x3f) == 0 && ((1L << (_la - 30)) & 408030265347L) != 0)) {
				{
				setState(512);
				expr(0);
				setState(517);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(513);
					match(T__3);
					setState(514);
					expr(0);
					}
					}
					setState(519);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(522);
			match(T__4);
			setState(524);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,55,_ctx) ) {
			case 1:
				{
				setState(523);
				match(T__5);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarCallFuncStructContext extends ParserRuleContext {
		public VarCallFuncStructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varCallFuncStruct; }
	 
		public VarCallFuncStructContext() { }
		public void copyFrom(VarCallFuncStructContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionStructContext extends VarCallFuncStructContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public CallFunctionStructContext(VarCallFuncStructContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterCallFunctionStruct(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitCallFunctionStruct(this);
		}
	}

	public final VarCallFuncStructContext varCallFuncStruct() throws RecognitionException {
		VarCallFuncStructContext _localctx = new VarCallFuncStructContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_varCallFuncStruct);
		int _la;
		try {
			_localctx = new CallFunctionStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(526);
			match(ID_VARIABLE);
			setState(527);
			match(T__28);
			setState(528);
			match(ID_VARIABLE);
			setState(529);
			match(T__44);
			setState(538);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 30)) & ~0x3f) == 0 && ((1L << (_la - 30)) & 408030265347L) != 0)) {
				{
				setState(530);
				expr(0);
				setState(535);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(531);
					match(T__3);
					setState(532);
					expr(0);
					}
					}
					setState(537);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(540);
			match(T__4);
			setState(542);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,58,_ctx) ) {
			case 1:
				{
				setState(541);
				match(T__5);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValRetContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ValRetContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valRet; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterValRet(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitValRet(this);
		}
	}

	public final ValRetContext valRet() throws RecognitionException {
		ValRetContext _localctx = new ValRetContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_valRet);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(544);
			type();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RetornoContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public RetornoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_retorno; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).enterRetorno(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof gramaticaListener ) ((gramaticaListener)listener).exitRetorno(this);
		}
	}

	public final RetornoContext retorno() throws RecognitionException {
		RetornoContext _localctx = new RetornoContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_retorno);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(546);
			match(T__60);
			setState(548);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,59,_ctx) ) {
			case 1:
				{
				setState(547);
				expr(0);
				}
				break;
			}
			setState(551);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(550);
				match(T__5);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 16:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 24);
		case 1:
			return precpred(_ctx, 23);
		case 2:
			return precpred(_ctx, 22);
		case 3:
			return precpred(_ctx, 21);
		case 4:
			return precpred(_ctx, 20);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001F\u022a\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0001\u0000\u0005\u0000:\b\u0000\n\u0000\f\u0000=\t\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0005\u0001D\b\u0001"+
		"\n\u0001\f\u0001G\t\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001W\b\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002]\b\u0002"+
		"\n\u0002\f\u0002`\t\u0002\u0003\u0002b\b\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002f\b\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0005\u0002l\b\u0002\n\u0002\f\u0002o\t\u0002\u0003\u0002q\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002u\b\u0002\u0003\u0002w\b\u0002\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003~\b\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0003\u0003\u0086\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0005\u0005\u0092\b\u0005\n\u0005\f\u0005\u0095\t\u0005\u0001\u0005\u0003"+
		"\u0005\u0098\b\u0005\u0001\u0005\u0001\u0005\u0005\u0005\u009c\b\u0005"+
		"\n\u0005\f\u0005\u009f\t\u0005\u0003\u0005\u00a1\b\u0005\u0001\u0006\u0001"+
		"\u0006\u0005\u0006\u00a5\b\u0006\n\u0006\f\u0006\u00a8\t\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u00c0\b\u0007\u0001\b"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0003\b\u00cf\b\b\u0001\t\u0001\t\u0001\t\u0004"+
		"\t\u00d4\b\t\u000b\t\f\t\u00d5\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t"+
		"\u0001\t\u0001\t\u0001\t\u0004\t\u00e0\b\t\u000b\t\f\t\u00e1\u0001\t\u0001"+
		"\t\u0003\t\u00e6\b\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001"+
		"\f\u0001\f\u0005\f\u00ef\b\f\n\f\f\f\u00f2\t\f\u0001\f\u0001\f\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u00fc\b\f\u0005\f\u00fe"+
		"\b\f\n\f\f\f\u0101\t\f\u0003\f\u0103\b\f\u0001\r\u0003\r\u0106\b\r\u0001"+
		"\r\u0003\r\u0109\b\r\u0001\r\u0001\r\u0003\r\u010d\b\r\u0001\r\u0001\r"+
		"\u0001\r\u0001\r\u0003\r\u0113\b\r\u0004\r\u0115\b\r\u000b\r\f\r\u0116"+
		"\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0004\u000e\u0126\b\u000e\u000b\u000e\f\u000e\u0127\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0003\u000e\u0137\b\u000e\u0004\u000e\u0139\b\u000e\u000b\u000e\f\u000e"+
		"\u013a\u0001\u000e\u0001\u000e\u0003\u000e\u013f\b\u000e\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0004"+
		"\u000f\u014e\b\u000f\u000b\u000f\f\u000f\u014f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0004\u000f\u0158\b\u000f"+
		"\u000b\u000f\f\u000f\u0159\u0001\u000f\u0001\u000f\u0003\u000f\u015e\b"+
		"\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0004\u0010\u0175\b\u0010\u000b"+
		"\u0010\f\u0010\u0176\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u0188"+
		"\b\u0010\n\u0010\f\u0010\u018b\t\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0004\u0010\u01a3\b\u0010\u000b\u0010\f\u0010\u01a4\u0001"+
		"\u0010\u0003\u0010\u01a8\b\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u01ac"+
		"\b\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u01bd\b\u0010\n"+
		"\u0010\f\u0010\u01c0\t\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0003\u0013\u01ca"+
		"\b\u0013\u0001\u0014\u0001\u0014\u0003\u0014\u01ce\b\u0014\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0005\u0015\u01d8\b\u0015\n\u0015\f\u0015\u01db\t\u0015\u0003"+
		"\u0015\u01dd\b\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u01e1\b\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u01ed\b\u0016"+
		"\u0001\u0016\u0001\u0016\u0003\u0016\u01f1\b\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0005\u0017"+
		"\u01fa\b\u0017\n\u0017\f\u0017\u01fd\t\u0017\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0005\u0018\u0204\b\u0018\n\u0018\f\u0018"+
		"\u0207\t\u0018\u0003\u0018\u0209\b\u0018\u0001\u0018\u0001\u0018\u0003"+
		"\u0018\u020d\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u0216\b\u0019\n\u0019\f\u0019"+
		"\u0219\t\u0019\u0003\u0019\u021b\b\u0019\u0001\u0019\u0001\u0019\u0003"+
		"\u0019\u021f\b\u0019\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0003"+
		"\u001b\u0225\b\u001b\u0001\u001b\u0003\u001b\u0228\b\u001b\u0001\u001b"+
		"\u0000\u0001 \u001c\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014"+
		"\u0016\u0018\u001a\u001c\u001e \"$&(*,.0246\u0000\t\u0002\u0000\u000f"+
		"\u000f\u0012\u0012\u0001\u0000\u0017\u0018\u0001\u0000\u0019\u001a\u0001"+
		"\u0000 \"\u0002\u0000\u001e\u001e##\u0001\u0000$\'\u0001\u0000()\u0001"+
		"\u0000*+\u0003\u0000\u0016\u001659DD\u0274\u0000;\u0001\u0000\u0000\u0000"+
		"\u0002V\u0001\u0000\u0000\u0000\u0004v\u0001\u0000\u0000\u0000\u0006\u0085"+
		"\u0001\u0000\u0000\u0000\b\u0087\u0001\u0000\u0000\u0000\n\u00a0\u0001"+
		"\u0000\u0000\u0000\f\u00a2\u0001\u0000\u0000\u0000\u000e\u00bf\u0001\u0000"+
		"\u0000\u0000\u0010\u00ce\u0001\u0000\u0000\u0000\u0012\u00e5\u0001\u0000"+
		"\u0000\u0000\u0014\u00e7\u0001\u0000\u0000\u0000\u0016\u00e9\u0001\u0000"+
		"\u0000\u0000\u0018\u0102\u0001\u0000\u0000\u0000\u001a\u0105\u0001\u0000"+
		"\u0000\u0000\u001c\u013e\u0001\u0000\u0000\u0000\u001e\u015d\u0001\u0000"+
		"\u0000\u0000 \u01ab\u0001\u0000\u0000\u0000\"\u01c1\u0001\u0000\u0000"+
		"\u0000$\u01c5\u0001\u0000\u0000\u0000&\u01c7\u0001\u0000\u0000\u0000("+
		"\u01cb\u0001\u0000\u0000\u0000*\u01cf\u0001\u0000\u0000\u0000,\u01e4\u0001"+
		"\u0000\u0000\u0000.\u01f4\u0001\u0000\u0000\u00000\u01fe\u0001\u0000\u0000"+
		"\u00002\u020e\u0001\u0000\u0000\u00004\u0220\u0001\u0000\u0000\u00006"+
		"\u0222\u0001\u0000\u0000\u00008:\u0003\u0002\u0001\u000098\u0001\u0000"+
		"\u0000\u0000:=\u0001\u0000\u0000\u0000;9\u0001\u0000\u0000\u0000;<\u0001"+
		"\u0000\u0000\u0000<\u0001\u0001\u0000\u0000\u0000=;\u0001\u0000\u0000"+
		"\u0000>W\u0003\u0004\u0002\u0000?W\u0003\u0006\u0003\u0000@W\u0003\b\u0004"+
		"\u0000AE\u0005\u0001\u0000\u0000BD\u0003\u0002\u0001\u0000CB\u0001\u0000"+
		"\u0000\u0000DG\u0001\u0000\u0000\u0000EC\u0001\u0000\u0000\u0000EF\u0001"+
		"\u0000\u0000\u0000FH\u0001\u0000\u0000\u0000GE\u0001\u0000\u0000\u0000"+
		"HW\u0005\u0002\u0000\u0000IW\u0003\u000e\u0007\u0000JW\u0003\u0012\t\u0000"+
		"KW\u0003\u001e\u000f\u0000LW\u0003\u0010\b\u0000MW\u0003\u001a\r\u0000"+
		"NW\u0003\u001c\u000e\u0000OW\u0003&\u0013\u0000PW\u0003(\u0014\u0000Q"+
		"W\u0003*\u0015\u0000RW\u0003,\u0016\u0000SW\u00030\u0018\u0000TW\u0003"+
		"2\u0019\u0000UW\u00036\u001b\u0000V>\u0001\u0000\u0000\u0000V?\u0001\u0000"+
		"\u0000\u0000V@\u0001\u0000\u0000\u0000VA\u0001\u0000\u0000\u0000VI\u0001"+
		"\u0000\u0000\u0000VJ\u0001\u0000\u0000\u0000VK\u0001\u0000\u0000\u0000"+
		"VL\u0001\u0000\u0000\u0000VM\u0001\u0000\u0000\u0000VN\u0001\u0000\u0000"+
		"\u0000VO\u0001\u0000\u0000\u0000VP\u0001\u0000\u0000\u0000VQ\u0001\u0000"+
		"\u0000\u0000VR\u0001\u0000\u0000\u0000VS\u0001\u0000\u0000\u0000VT\u0001"+
		"\u0000\u0000\u0000VU\u0001\u0000\u0000\u0000W\u0003\u0001\u0000\u0000"+
		"\u0000Xa\u0005\u0003\u0000\u0000Y^\u0003 \u0010\u0000Z[\u0005\u0004\u0000"+
		"\u0000[]\u0003 \u0010\u0000\\Z\u0001\u0000\u0000\u0000]`\u0001\u0000\u0000"+
		"\u0000^\\\u0001\u0000\u0000\u0000^_\u0001\u0000\u0000\u0000_b\u0001\u0000"+
		"\u0000\u0000`^\u0001\u0000\u0000\u0000aY\u0001\u0000\u0000\u0000ab\u0001"+
		"\u0000\u0000\u0000bc\u0001\u0000\u0000\u0000ce\u0005\u0005\u0000\u0000"+
		"df\u0005\u0006\u0000\u0000ed\u0001\u0000\u0000\u0000ef\u0001\u0000\u0000"+
		"\u0000fw\u0001\u0000\u0000\u0000gp\u0005\u0007\u0000\u0000hm\u0003 \u0010"+
		"\u0000ij\u0005\u0004\u0000\u0000jl\u0003 \u0010\u0000ki\u0001\u0000\u0000"+
		"\u0000lo\u0001\u0000\u0000\u0000mk\u0001\u0000\u0000\u0000mn\u0001\u0000"+
		"\u0000\u0000nq\u0001\u0000\u0000\u0000om\u0001\u0000\u0000\u0000ph\u0001"+
		"\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qr\u0001\u0000\u0000\u0000"+
		"rt\u0005\u0005\u0000\u0000su\u0005\u0006\u0000\u0000ts\u0001\u0000\u0000"+
		"\u0000tu\u0001\u0000\u0000\u0000uw\u0001\u0000\u0000\u0000vX\u0001\u0000"+
		"\u0000\u0000vg\u0001\u0000\u0000\u0000w\u0005\u0001\u0000\u0000\u0000"+
		"xy\u0005\b\u0000\u0000yz\u0003 \u0010\u0000z}\u0003\f\u0006\u0000{|\u0005"+
		"\t\u0000\u0000|~\u0003\f\u0006\u0000}{\u0001\u0000\u0000\u0000}~\u0001"+
		"\u0000\u0000\u0000~\u0086\u0001\u0000\u0000\u0000\u007f\u0080\u0005\b"+
		"\u0000\u0000\u0080\u0081\u0003 \u0010\u0000\u0081\u0082\u0003\f\u0006"+
		"\u0000\u0082\u0083\u0005\t\u0000\u0000\u0083\u0084\u0003\u0006\u0003\u0000"+
		"\u0084\u0086\u0001\u0000\u0000\u0000\u0085x\u0001\u0000\u0000\u0000\u0085"+
		"\u007f\u0001\u0000\u0000\u0000\u0086\u0007\u0001\u0000\u0000\u0000\u0087"+
		"\u0088\u0005\n\u0000\u0000\u0088\u0089\u0003 \u0010\u0000\u0089\u008a"+
		"\u0005\u0001\u0000\u0000\u008a\u008b\u0003\n\u0005\u0000\u008b\u008c\u0005"+
		"\u0002\u0000\u0000\u008c\t\u0001\u0000\u0000\u0000\u008d\u008e\u0005\u000b"+
		"\u0000\u0000\u008e\u008f\u0003 \u0010\u0000\u008f\u0093\u0005\f\u0000"+
		"\u0000\u0090\u0092\u0003\u0002\u0001\u0000\u0091\u0090\u0001\u0000\u0000"+
		"\u0000\u0092\u0095\u0001\u0000\u0000\u0000\u0093\u0091\u0001\u0000\u0000"+
		"\u0000\u0093\u0094\u0001\u0000\u0000\u0000\u0094\u0097\u0001\u0000\u0000"+
		"\u0000\u0095\u0093\u0001\u0000\u0000\u0000\u0096\u0098\u0003\n\u0005\u0000"+
		"\u0097\u0096\u0001\u0000\u0000\u0000\u0097\u0098\u0001\u0000\u0000\u0000"+
		"\u0098\u00a1\u0001\u0000\u0000\u0000\u0099\u009d\u0005\r\u0000\u0000\u009a"+
		"\u009c\u0003\u0002\u0001\u0000\u009b\u009a\u0001\u0000\u0000\u0000\u009c"+
		"\u009f\u0001\u0000\u0000\u0000\u009d\u009b\u0001\u0000\u0000\u0000\u009d"+
		"\u009e\u0001\u0000\u0000\u0000\u009e\u00a1\u0001\u0000\u0000\u0000\u009f"+
		"\u009d\u0001\u0000\u0000\u0000\u00a0\u008d\u0001\u0000\u0000\u0000\u00a0"+
		"\u0099\u0001\u0000\u0000\u0000\u00a1\u000b\u0001\u0000\u0000\u0000\u00a2"+
		"\u00a6\u0005\u0001\u0000\u0000\u00a3\u00a5\u0003\u0002\u0001\u0000\u00a4"+
		"\u00a3\u0001\u0000\u0000\u0000\u00a5\u00a8\u0001\u0000\u0000\u0000\u00a6"+
		"\u00a4\u0001\u0000\u0000\u0000\u00a6\u00a7\u0001\u0000\u0000\u0000\u00a7"+
		"\u00a9\u0001\u0000\u0000\u0000\u00a8\u00a6\u0001\u0000\u0000\u0000\u00a9"+
		"\u00aa\u0005\u0002\u0000\u0000\u00aa\r\u0001\u0000\u0000\u0000\u00ab\u00ac"+
		"\u0005\u000e\u0000\u0000\u00ac\u00ad\u0003 \u0010\u0000\u00ad\u00ae\u0003"+
		"\f\u0006\u0000\u00ae\u00c0\u0001\u0000\u0000\u0000\u00af\u00b0\u0005\u000e"+
		"\u0000\u0000\u00b0\u00b1\u0003\u0010\b\u0000\u00b1\u00b2\u0005\u0006\u0000"+
		"\u0000\u00b2\u00b3\u0003 \u0010\u0000\u00b3\u00b4\u0005\u0006\u0000\u0000"+
		"\u00b4\u00b5\u0003\u001e\u000f\u0000\u00b5\u00b6\u0003\f\u0006\u0000\u00b6"+
		"\u00c0\u0001\u0000\u0000\u0000\u00b7\u00b8\u0005\u000e\u0000\u0000\u00b8"+
		"\u00b9\u0005D\u0000\u0000\u00b9\u00ba\u0005\u0004\u0000\u0000\u00ba\u00bb"+
		"\u0005D\u0000\u0000\u00bb\u00bc\u0005\u000f\u0000\u0000\u00bc\u00bd\u0005"+
		"\u0010\u0000\u0000\u00bd\u00be\u0005D\u0000\u0000\u00be\u00c0\u0003\f"+
		"\u0006\u0000\u00bf\u00ab\u0001\u0000\u0000\u0000\u00bf\u00af\u0001\u0000"+
		"\u0000\u0000\u00bf\u00b7\u0001\u0000\u0000\u0000\u00c0\u000f\u0001\u0000"+
		"\u0000\u0000\u00c1\u00c2\u0005\u0011\u0000\u0000\u00c2\u00c3\u0005D\u0000"+
		"\u0000\u00c3\u00c4\u0003$\u0012\u0000\u00c4\u00c5\u0005\u0012\u0000\u0000"+
		"\u00c5\u00c6\u0003 \u0010\u0000\u00c6\u00cf\u0001\u0000\u0000\u0000\u00c7"+
		"\u00c8\u0005\u0011\u0000\u0000\u00c8\u00c9\u0005D\u0000\u0000\u00c9\u00cf"+
		"\u0003$\u0012\u0000\u00ca\u00cb\u0005\u0011\u0000\u0000\u00cb\u00cc\u0005"+
		"D\u0000\u0000\u00cc\u00cd\u0005\u000f\u0000\u0000\u00cd\u00cf\u0003 \u0010"+
		"\u0000\u00ce\u00c1\u0001\u0000\u0000\u0000\u00ce\u00c7\u0001\u0000\u0000"+
		"\u0000\u00ce\u00ca\u0001\u0000\u0000\u0000\u00cf\u0011\u0001\u0000\u0000"+
		"\u0000\u00d0\u00d1\u0005D\u0000\u0000\u00d1\u00d3\u0003\u0014\n\u0000"+
		"\u00d2\u00d4\u0003\u0016\u000b\u0000\u00d3\u00d2\u0001\u0000\u0000\u0000"+
		"\u00d4\u00d5\u0001\u0000\u0000\u0000\u00d5\u00d3\u0001\u0000\u0000\u0000"+
		"\u00d5\u00d6\u0001\u0000\u0000\u0000\u00d6\u00d7\u0001\u0000\u0000\u0000"+
		"\u00d7\u00d8\u0003$\u0012\u0000\u00d8\u00d9\u0005\u0001\u0000\u0000\u00d9"+
		"\u00da\u0003\u0018\f\u0000\u00da\u00db\u0005\u0002\u0000\u0000\u00db\u00e6"+
		"\u0001\u0000\u0000\u0000\u00dc\u00dd\u0005\u0013\u0000\u0000\u00dd\u00df"+
		"\u0005D\u0000\u0000\u00de\u00e0\u0003\u0016\u000b\u0000\u00df\u00de\u0001"+
		"\u0000\u0000\u0000\u00e0\u00e1\u0001\u0000\u0000\u0000\u00e1\u00df\u0001"+
		"\u0000\u0000\u0000\u00e1\u00e2\u0001\u0000\u0000\u0000\u00e2\u00e3\u0001"+
		"\u0000\u0000\u0000\u00e3\u00e4\u0003$\u0012\u0000\u00e4\u00e6\u0001\u0000"+
		"\u0000\u0000\u00e5\u00d0\u0001\u0000\u0000\u0000\u00e5\u00dc\u0001\u0000"+
		"\u0000\u0000\u00e6\u0013\u0001\u0000\u0000\u0000\u00e7\u00e8\u0007\u0000"+
		"\u0000\u0000\u00e8\u0015\u0001\u0000\u0000\u0000\u00e9\u00ea\u0005\u0014"+
		"\u0000\u0000\u00ea\u0017\u0001\u0000\u0000\u0000\u00eb\u00f0\u0003 \u0010"+
		"\u0000\u00ec\u00ed\u0005\u0004\u0000\u0000\u00ed\u00ef\u0003 \u0010\u0000"+
		"\u00ee\u00ec\u0001\u0000\u0000\u0000\u00ef\u00f2\u0001\u0000\u0000\u0000"+
		"\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f0\u00f1\u0001\u0000\u0000\u0000"+
		"\u00f1\u0103\u0001\u0000\u0000\u0000\u00f2\u00f0\u0001\u0000\u0000\u0000"+
		"\u00f3\u00f4\u0005\u0001\u0000\u0000\u00f4\u00f5\u0003\u0018\f\u0000\u00f5"+
		"\u00ff\u0005\u0002\u0000\u0000\u00f6\u00fb\u0005\u0004\u0000\u0000\u00f7"+
		"\u00f8\u0005\u0001\u0000\u0000\u00f8\u00f9\u0003\u0018\f\u0000\u00f9\u00fa"+
		"\u0005\u0002\u0000\u0000\u00fa\u00fc\u0001\u0000\u0000\u0000\u00fb\u00f7"+
		"\u0001\u0000\u0000\u0000\u00fb\u00fc\u0001\u0000\u0000\u0000\u00fc\u00fe"+
		"\u0001\u0000\u0000\u0000\u00fd\u00f6\u0001\u0000\u0000\u0000\u00fe\u0101"+
		"\u0001\u0000\u0000\u0000\u00ff\u00fd\u0001\u0000\u0000\u0000\u00ff\u0100"+
		"\u0001\u0000\u0000\u0000\u0100\u0103\u0001\u0000\u0000\u0000\u0101\u00ff"+
		"\u0001\u0000\u0000\u0000\u0102\u00eb\u0001\u0000\u0000\u0000\u0102\u00f3"+
		"\u0001\u0000\u0000\u0000\u0103\u0019\u0001\u0000\u0000\u0000\u0104\u0106"+
		"\u0005\u0015\u0000\u0000\u0105\u0104\u0001\u0000\u0000\u0000\u0105\u0106"+
		"\u0001\u0000\u0000\u0000\u0106\u0108\u0001\u0000\u0000\u0000\u0107\u0109"+
		"\u0005\u0016\u0000\u0000\u0108\u0107\u0001\u0000\u0000\u0000\u0108\u0109"+
		"\u0001\u0000\u0000\u0000\u0109\u010a\u0001\u0000\u0000\u0000\u010a\u010c"+
		"\u0005D\u0000\u0000\u010b\u010d\u0005\u0016\u0000\u0000\u010c\u010b\u0001"+
		"\u0000\u0000\u0000\u010c\u010d\u0001\u0000\u0000\u0000\u010d\u010e\u0001"+
		"\u0000\u0000\u0000\u010e\u0114\u0005\u0001\u0000\u0000\u010f\u0110\u0005"+
		"D\u0000\u0000\u0110\u0112\u0003$\u0012\u0000\u0111\u0113\u0005\u0006\u0000"+
		"\u0000\u0112\u0111\u0001\u0000\u0000\u0000\u0112\u0113\u0001\u0000\u0000"+
		"\u0000\u0113\u0115\u0001\u0000\u0000\u0000\u0114\u010f\u0001\u0000\u0000"+
		"\u0000\u0115\u0116\u0001\u0000\u0000\u0000\u0116\u0114\u0001\u0000\u0000"+
		"\u0000\u0116\u0117\u0001\u0000\u0000\u0000\u0117\u0118\u0001\u0000\u0000"+
		"\u0000\u0118\u0119\u0005\u0002\u0000\u0000\u0119\u001b\u0001\u0000\u0000"+
		"\u0000\u011a\u011b\u0005D\u0000\u0000\u011b\u011c\u0005D\u0000\u0000\u011c"+
		"\u011d\u0005\u0012\u0000\u0000\u011d\u011e\u0005\u0001\u0000\u0000\u011e"+
		"\u011f\u0005D\u0000\u0000\u011f\u0120\u0005\f\u0000\u0000\u0120\u0125"+
		"\u0003 \u0010\u0000\u0121\u0122\u0005\u0004\u0000\u0000\u0122\u0123\u0005"+
		"D\u0000\u0000\u0123\u0124\u0005\f\u0000\u0000\u0124\u0126\u0003 \u0010"+
		"\u0000\u0125\u0121\u0001\u0000\u0000\u0000\u0126\u0127\u0001\u0000\u0000"+
		"\u0000\u0127\u0125\u0001\u0000\u0000\u0000\u0127\u0128\u0001\u0000\u0000"+
		"\u0000\u0128\u0129\u0001\u0000\u0000\u0000\u0129\u012a\u0005\u0002\u0000"+
		"\u0000\u012a\u013f\u0001\u0000\u0000\u0000\u012b\u012c\u0005D\u0000\u0000"+
		"\u012c\u012d\u0005\u000f\u0000\u0000\u012d\u012e\u0005D\u0000\u0000\u012e"+
		"\u012f\u0005\u0001\u0000\u0000\u012f\u0130\u0005D\u0000\u0000\u0130\u0131"+
		"\u0005\f\u0000\u0000\u0131\u0138\u0003 \u0010\u0000\u0132\u0136\u0005"+
		"\u0004\u0000\u0000\u0133\u0134\u0005D\u0000\u0000\u0134\u0135\u0005\f"+
		"\u0000\u0000\u0135\u0137\u0003 \u0010\u0000\u0136\u0133\u0001\u0000\u0000"+
		"\u0000\u0136\u0137\u0001\u0000\u0000\u0000\u0137\u0139\u0001\u0000\u0000"+
		"\u0000\u0138\u0132\u0001\u0000\u0000\u0000\u0139\u013a\u0001\u0000\u0000"+
		"\u0000\u013a\u0138\u0001\u0000\u0000\u0000\u013a\u013b\u0001\u0000\u0000"+
		"\u0000\u013b\u013c\u0001\u0000\u0000\u0000\u013c\u013d\u0005\u0002\u0000"+
		"\u0000\u013d\u013f\u0001\u0000\u0000\u0000\u013e\u011a\u0001\u0000\u0000"+
		"\u0000\u013e\u012b\u0001\u0000\u0000\u0000\u013f\u001d\u0001\u0000\u0000"+
		"\u0000\u0140\u0141\u0005D\u0000\u0000\u0141\u0142\u0005\u0012\u0000\u0000"+
		"\u0142\u015e\u0003 \u0010\u0000\u0143\u0144\u0005D\u0000\u0000\u0144\u0145"+
		"\u0007\u0001\u0000\u0000\u0145\u015e\u0003 \u0010\u0000\u0146\u0147\u0005"+
		"D\u0000\u0000\u0147\u015e\u0007\u0002\u0000\u0000\u0148\u014d\u0005D\u0000"+
		"\u0000\u0149\u014a\u0005\u001b\u0000\u0000\u014a\u014b\u0003 \u0010\u0000"+
		"\u014b\u014c\u0005\u001c\u0000\u0000\u014c\u014e\u0001\u0000\u0000\u0000"+
		"\u014d\u0149\u0001\u0000\u0000\u0000\u014e\u014f\u0001\u0000\u0000\u0000"+
		"\u014f\u014d\u0001\u0000\u0000\u0000\u014f\u0150\u0001\u0000\u0000\u0000"+
		"\u0150\u0151\u0001\u0000\u0000\u0000\u0151\u0152\u0005\u0012\u0000\u0000"+
		"\u0152\u0153\u0003 \u0010\u0000\u0153\u015e\u0001\u0000\u0000\u0000\u0154"+
		"\u0157\u0005D\u0000\u0000\u0155\u0156\u0005\u001d\u0000\u0000\u0156\u0158"+
		"\u0005D\u0000\u0000\u0157\u0155\u0001\u0000\u0000\u0000\u0158\u0159\u0001"+
		"\u0000\u0000\u0000\u0159\u0157\u0001\u0000\u0000\u0000\u0159\u015a\u0001"+
		"\u0000\u0000\u0000\u015a\u015b\u0001\u0000\u0000\u0000\u015b\u015c\u0005"+
		"\u0012\u0000\u0000\u015c\u015e\u0003 \u0010\u0000\u015d\u0140\u0001\u0000"+
		"\u0000\u0000\u015d\u0143\u0001\u0000\u0000\u0000\u015d\u0146\u0001\u0000"+
		"\u0000\u0000\u015d\u0148\u0001\u0000\u0000\u0000\u015d\u0154\u0001\u0000"+
		"\u0000\u0000\u015e\u001f\u0001\u0000\u0000\u0000\u015f\u0160\u0006\u0010"+
		"\uffff\uffff\u0000\u0160\u0161\u0005\u001e\u0000\u0000\u0161\u01ac\u0003"+
		" \u0010\u001a\u0162\u0163\u0005\u001f\u0000\u0000\u0163\u01ac\u0003 \u0010"+
		"\u0019\u0164\u01ac\u0005>\u0000\u0000\u0165\u01ac\u0005?\u0000\u0000\u0166"+
		"\u01ac\u0005A\u0000\u0000\u0167\u01ac\u0005B\u0000\u0000\u0168\u01ac\u0005"+
		"D\u0000\u0000\u0169\u01ac\u0005@\u0000\u0000\u016a\u01ac\u0005,\u0000"+
		"\u0000\u016b\u016c\u0005-\u0000\u0000\u016c\u016d\u0003 \u0010\u0000\u016d"+
		"\u016e\u0005\u0005\u0000\u0000\u016e\u01ac\u0001\u0000\u0000\u0000\u016f"+
		"\u0174\u0005D\u0000\u0000\u0170\u0171\u0005\u001b\u0000\u0000\u0171\u0172"+
		"\u0003 \u0010\u0000\u0172\u0173\u0005\u001c\u0000\u0000\u0173\u0175\u0001"+
		"\u0000\u0000\u0000\u0174\u0170\u0001\u0000\u0000\u0000\u0175\u0176\u0001"+
		"\u0000\u0000\u0000\u0176\u0174\u0001\u0000\u0000\u0000\u0176\u0177\u0001"+
		"\u0000\u0000\u0000\u0177\u01ac\u0001\u0000\u0000\u0000\u0178\u0179\u0005"+
		".\u0000\u0000\u0179\u017a\u0005D\u0000\u0000\u017a\u017b\u0005\u0004\u0000"+
		"\u0000\u017b\u017c\u0003 \u0010\u0000\u017c\u017d\u0005\u0005\u0000\u0000"+
		"\u017d\u01ac\u0001\u0000\u0000\u0000\u017e\u017f\u0005/\u0000\u0000\u017f"+
		"\u0180\u0005D\u0000\u0000\u0180\u0181\u0005\u0004\u0000\u0000\u0181\u0182"+
		"\u0003 \u0010\u0000\u0182\u0183\u0005\u0005\u0000\u0000\u0183\u01ac\u0001"+
		"\u0000\u0000\u0000\u0184\u0185\u00050\u0000\u0000\u0185\u0189\u0005D\u0000"+
		"\u0000\u0186\u0188\u0003\"\u0011\u0000\u0187\u0186\u0001\u0000\u0000\u0000"+
		"\u0188\u018b\u0001\u0000\u0000\u0000\u0189\u0187\u0001\u0000\u0000\u0000"+
		"\u0189\u018a\u0001\u0000\u0000\u0000\u018a\u018c\u0001\u0000\u0000\u0000"+
		"\u018b\u0189\u0001\u0000\u0000\u0000\u018c\u01ac\u0005\u0005\u0000\u0000"+
		"\u018d\u018e\u00051\u0000\u0000\u018e\u018f\u0005D\u0000\u0000\u018f\u0190"+
		"\u0005\u0004\u0000\u0000\u0190\u0191\u0003 \u0010\u0000\u0191\u0192\u0005"+
		"\u0005\u0000\u0000\u0192\u01ac\u0001\u0000\u0000\u0000\u0193\u0194\u0005"+
		"2\u0000\u0000\u0194\u0195\u0003 \u0010\u0000\u0195\u0196\u0005\u0005\u0000"+
		"\u0000\u0196\u01ac\u0001\u0000\u0000\u0000\u0197\u0198\u00053\u0000\u0000"+
		"\u0198\u0199\u0003 \u0010\u0000\u0199\u019a\u0005\u0005\u0000\u0000\u019a"+
		"\u01ac\u0001\u0000\u0000\u0000\u019b\u019c\u00054\u0000\u0000\u019c\u019d"+
		"\u0003 \u0010\u0000\u019d\u019e\u0005\u0005\u0000\u0000\u019e\u01ac\u0001"+
		"\u0000\u0000\u0000\u019f\u01a2\u0005D\u0000\u0000\u01a0\u01a1\u0005\u001d"+
		"\u0000\u0000\u01a1\u01a3\u0005D\u0000\u0000\u01a2\u01a0\u0001\u0000\u0000"+
		"\u0000\u01a3\u01a4\u0001\u0000\u0000\u0000\u01a4\u01a2\u0001\u0000\u0000"+
		"\u0000\u01a4\u01a5\u0001\u0000\u0000\u0000\u01a5\u01a7\u0001\u0000\u0000"+
		"\u0000\u01a6\u01a8\u0005\u0006\u0000\u0000\u01a7\u01a6\u0001\u0000\u0000"+
		"\u0000\u01a7\u01a8\u0001\u0000\u0000\u0000\u01a8\u01ac\u0001\u0000\u0000"+
		"\u0000\u01a9\u01ac\u00030\u0018\u0000\u01aa\u01ac\u00032\u0019\u0000\u01ab"+
		"\u015f\u0001\u0000\u0000\u0000\u01ab\u0162\u0001\u0000\u0000\u0000\u01ab"+
		"\u0164\u0001\u0000\u0000\u0000\u01ab\u0165\u0001\u0000\u0000\u0000\u01ab"+
		"\u0166\u0001\u0000\u0000\u0000\u01ab\u0167\u0001\u0000\u0000\u0000\u01ab"+
		"\u0168\u0001\u0000\u0000\u0000\u01ab\u0169\u0001\u0000\u0000\u0000\u01ab"+
		"\u016a\u0001\u0000\u0000\u0000\u01ab\u016b\u0001\u0000\u0000\u0000\u01ab"+
		"\u016f\u0001\u0000\u0000\u0000\u01ab\u0178\u0001\u0000\u0000\u0000\u01ab"+
		"\u017e\u0001\u0000\u0000\u0000\u01ab\u0184\u0001\u0000\u0000\u0000\u01ab"+
		"\u018d\u0001\u0000\u0000\u0000\u01ab\u0193\u0001\u0000\u0000\u0000\u01ab"+
		"\u0197\u0001\u0000\u0000\u0000\u01ab\u019b\u0001\u0000\u0000\u0000\u01ab"+
		"\u019f\u0001\u0000\u0000\u0000\u01ab\u01a9\u0001\u0000\u0000\u0000\u01ab"+
		"\u01aa\u0001\u0000\u0000\u0000\u01ac\u01be\u0001\u0000\u0000\u0000\u01ad"+
		"\u01ae\n\u0018\u0000\u0000\u01ae\u01af\u0007\u0003\u0000\u0000\u01af\u01bd"+
		"\u0003 \u0010\u0019\u01b0\u01b1\n\u0017\u0000\u0000\u01b1\u01b2\u0007"+
		"\u0004\u0000\u0000\u01b2\u01bd\u0003 \u0010\u0018\u01b3\u01b4\n\u0016"+
		"\u0000\u0000\u01b4\u01b5\u0007\u0005\u0000\u0000\u01b5\u01bd\u0003 \u0010"+
		"\u0017\u01b6\u01b7\n\u0015\u0000\u0000\u01b7\u01b8\u0007\u0006\u0000\u0000"+
		"\u01b8\u01bd\u0003 \u0010\u0016\u01b9\u01ba\n\u0014\u0000\u0000\u01ba"+
		"\u01bb\u0007\u0007\u0000\u0000\u01bb\u01bd\u0003 \u0010\u0015\u01bc\u01ad"+
		"\u0001\u0000\u0000\u0000\u01bc\u01b0\u0001\u0000\u0000\u0000\u01bc\u01b3"+
		"\u0001\u0000\u0000\u0000\u01bc\u01b6\u0001\u0000\u0000\u0000\u01bc\u01b9"+
		"\u0001\u0000\u0000\u0000\u01bd\u01c0\u0001\u0000\u0000\u0000\u01be\u01bc"+
		"\u0001\u0000\u0000\u0000\u01be\u01bf\u0001\u0000\u0000\u0000\u01bf!\u0001"+
		"\u0000\u0000\u0000\u01c0\u01be\u0001\u0000\u0000\u0000\u01c1\u01c2\u0005"+
		"\u001b\u0000\u0000\u01c2\u01c3\u0003 \u0010\u0000\u01c3\u01c4\u0005\u001c"+
		"\u0000\u0000\u01c4#\u0001\u0000\u0000\u0000\u01c5\u01c6\u0007\b\u0000"+
		"\u0000\u01c6%\u0001\u0000\u0000\u0000\u01c7\u01c9\u0005:\u0000\u0000\u01c8"+
		"\u01ca\u0005\u0006\u0000\u0000\u01c9\u01c8\u0001\u0000\u0000\u0000\u01c9"+
		"\u01ca\u0001\u0000\u0000\u0000\u01ca\'\u0001\u0000\u0000\u0000\u01cb\u01cd"+
		"\u0005;\u0000\u0000\u01cc\u01ce\u0005\u0006\u0000\u0000\u01cd\u01cc\u0001"+
		"\u0000\u0000\u0000\u01cd\u01ce\u0001\u0000\u0000\u0000\u01ce)\u0001\u0000"+
		"\u0000\u0000\u01cf\u01d0\u0005<\u0000\u0000\u01d0\u01d1\u0005D\u0000\u0000"+
		"\u01d1\u01dc\u0005-\u0000\u0000\u01d2\u01d3\u0005D\u0000\u0000\u01d3\u01d9"+
		"\u0003$\u0012\u0000\u01d4\u01d5\u0005\u0004\u0000\u0000\u01d5\u01d6\u0005"+
		"D\u0000\u0000\u01d6\u01d8\u0003$\u0012\u0000\u01d7\u01d4\u0001\u0000\u0000"+
		"\u0000\u01d8\u01db\u0001\u0000\u0000\u0000\u01d9\u01d7\u0001\u0000\u0000"+
		"\u0000\u01d9\u01da\u0001\u0000\u0000\u0000\u01da\u01dd\u0001\u0000\u0000"+
		"\u0000\u01db\u01d9\u0001\u0000\u0000\u0000\u01dc\u01d2\u0001\u0000\u0000"+
		"\u0000\u01dc\u01dd\u0001\u0000\u0000\u0000\u01dd\u01de\u0001\u0000\u0000"+
		"\u0000\u01de\u01e0\u0005\u0005\u0000\u0000\u01df\u01e1\u00034\u001a\u0000"+
		"\u01e0\u01df\u0001\u0000\u0000\u0000\u01e0\u01e1\u0001\u0000\u0000\u0000"+
		"\u01e1\u01e2\u0001\u0000\u0000\u0000\u01e2\u01e3\u0003\f\u0006\u0000\u01e3"+
		"+\u0001\u0000\u0000\u0000\u01e4\u01e5\u0005<\u0000\u0000\u01e5\u01e6\u0005"+
		"-\u0000\u0000\u01e6\u01e7\u0005D\u0000\u0000\u01e7\u01e8\u0005D\u0000"+
		"\u0000\u01e8\u01e9\u0005\u0005\u0000\u0000\u01e9\u01ea\u0005D\u0000\u0000"+
		"\u01ea\u01ec\u0005-\u0000\u0000\u01eb\u01ed\u0003.\u0017\u0000\u01ec\u01eb"+
		"\u0001\u0000\u0000\u0000\u01ec\u01ed\u0001\u0000\u0000\u0000\u01ed\u01ee"+
		"\u0001\u0000\u0000\u0000\u01ee\u01f0\u0005\u0005\u0000\u0000\u01ef\u01f1"+
		"\u00034\u001a\u0000\u01f0\u01ef\u0001\u0000\u0000\u0000\u01f0\u01f1\u0001"+
		"\u0000\u0000\u0000\u01f1\u01f2\u0001\u0000\u0000\u0000\u01f2\u01f3\u0003"+
		"\f\u0006\u0000\u01f3-\u0001\u0000\u0000\u0000\u01f4\u01f5\u0005D\u0000"+
		"\u0000\u01f5\u01fb\u0003$\u0012\u0000\u01f6\u01f7\u0005\u0004\u0000\u0000"+
		"\u01f7\u01f8\u0005D\u0000\u0000\u01f8\u01fa\u0003$\u0012\u0000\u01f9\u01f6"+
		"\u0001\u0000\u0000\u0000\u01fa\u01fd\u0001\u0000\u0000\u0000\u01fb\u01f9"+
		"\u0001\u0000\u0000\u0000\u01fb\u01fc\u0001\u0000\u0000\u0000\u01fc/\u0001"+
		"\u0000\u0000\u0000\u01fd\u01fb\u0001\u0000\u0000\u0000\u01fe\u01ff\u0005"+
		"D\u0000\u0000\u01ff\u0208\u0005-\u0000\u0000\u0200\u0205\u0003 \u0010"+
		"\u0000\u0201\u0202\u0005\u0004\u0000\u0000\u0202\u0204\u0003 \u0010\u0000"+
		"\u0203\u0201\u0001\u0000\u0000\u0000\u0204\u0207\u0001\u0000\u0000\u0000"+
		"\u0205\u0203\u0001\u0000\u0000\u0000\u0205\u0206\u0001\u0000\u0000\u0000"+
		"\u0206\u0209\u0001\u0000\u0000\u0000\u0207\u0205\u0001\u0000\u0000\u0000"+
		"\u0208\u0200\u0001\u0000\u0000\u0000\u0208\u0209\u0001\u0000\u0000\u0000"+
		"\u0209\u020a\u0001\u0000\u0000\u0000\u020a\u020c\u0005\u0005\u0000\u0000"+
		"\u020b\u020d\u0005\u0006\u0000\u0000\u020c\u020b\u0001\u0000\u0000\u0000"+
		"\u020c\u020d\u0001\u0000\u0000\u0000\u020d1\u0001\u0000\u0000\u0000\u020e"+
		"\u020f\u0005D\u0000\u0000\u020f\u0210\u0005\u001d\u0000\u0000\u0210\u0211"+
		"\u0005D\u0000\u0000\u0211\u021a\u0005-\u0000\u0000\u0212\u0217\u0003 "+
		"\u0010\u0000\u0213\u0214\u0005\u0004\u0000\u0000\u0214\u0216\u0003 \u0010"+
		"\u0000\u0215\u0213\u0001\u0000\u0000\u0000\u0216\u0219\u0001\u0000\u0000"+
		"\u0000\u0217\u0215\u0001\u0000\u0000\u0000\u0217\u0218\u0001\u0000\u0000"+
		"\u0000\u0218\u021b\u0001\u0000\u0000\u0000\u0219\u0217\u0001\u0000\u0000"+
		"\u0000\u021a\u0212\u0001\u0000\u0000\u0000\u021a\u021b\u0001\u0000\u0000"+
		"\u0000\u021b\u021c\u0001\u0000\u0000\u0000\u021c\u021e\u0005\u0005\u0000"+
		"\u0000\u021d\u021f\u0005\u0006\u0000\u0000\u021e\u021d\u0001\u0000\u0000"+
		"\u0000\u021e\u021f\u0001\u0000\u0000\u0000\u021f3\u0001\u0000\u0000\u0000"+
		"\u0220\u0221\u0003$\u0012\u0000\u02215\u0001\u0000\u0000\u0000\u0222\u0224"+
		"\u0005=\u0000\u0000\u0223\u0225\u0003 \u0010\u0000\u0224\u0223\u0001\u0000"+
		"\u0000\u0000\u0224\u0225\u0001\u0000\u0000\u0000\u0225\u0227\u0001\u0000"+
		"\u0000\u0000\u0226\u0228\u0005\u0006\u0000\u0000\u0227\u0226\u0001\u0000"+
		"\u0000\u0000\u0227\u0228\u0001\u0000\u0000\u0000\u02287\u0001\u0000\u0000"+
		"\u0000=;EV^aemptv}\u0085\u0093\u0097\u009d\u00a0\u00a6\u00bf\u00ce\u00d5"+
		"\u00e1\u00e5\u00f0\u00fb\u00ff\u0102\u0105\u0108\u010c\u0112\u0116\u0127"+
		"\u0136\u013a\u013e\u014f\u0159\u015d\u0176\u0189\u01a4\u01a7\u01ab\u01bc"+
		"\u01be\u01c9\u01cd\u01d9\u01dc\u01e0\u01ec\u01f0\u01fb\u0205\u0208\u020c"+
		"\u0217\u021a\u021e\u0224\u0227";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}