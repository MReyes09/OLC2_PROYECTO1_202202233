// Generated from /home/mati/Escritorio/OLC2Vacas/OLC2_PROYECTO1_202202233/gramatica/gramatica.g4 by ANTLR 4.13.1
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
		T__59=60, INT=61, DOUBLE=62, CHAR=63, STRING=64, BOOL=65, BLANCOS=66, 
		ID_VARIABLE=67, COMENTARIOLINEA=68, COMENTARIOMULTILINEA=69;
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
			"'else'", "'switch'", "'case'", "':'", "'default:'", "'for'", "'in'", 
			"'mut'", "':='", "'='", "'[]'", "'type'", "'struct'", "'+='", "'-='", 
			"'++'", "'--'", "'['", "']'", "'.'", "'-'", "'!'", "'*'", "'/'", "'%'", 
			"'+'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'", "'||'", 
			"'nil'", "'('", "'indexOf('", "'join('", "'len('", "'append('", "'Atoi('", 
			"'parseFloat('", "'typeOf('", "'int'", "'float64'", "'string'", "'bool'", 
			"'rune'", "'break'", "'continue'", "'func'", "'return'"
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
			null, "INT", "DOUBLE", "CHAR", "STRING", "BOOL", "BLANCOS", "ID_VARIABLE", 
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
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2161727821141067146L) != 0) || _la==ID_VARIABLE) {
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionStmtContext extends InstruccionesContext {
		public VarCallStatementContext varCallStatement() {
			return getRuleContext(VarCallStatementContext.class,0);
		}
		public CallFunctionStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintStmtContext extends InstruccionesContext {
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
		}
		public PrintStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarStructDclStmtContext extends InstruccionesContext {
		public VarStructDclContext varStructDcl() {
			return getRuleContext(VarStructDclContext.class,0);
		}
		public VarStructDclStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignStmtContext extends InstruccionesContext {
		public VarAsignContext varAsign() {
			return getRuleContext(VarAsignContext.class,0);
		}
		public AsignStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionStructStmtContext extends InstruccionesContext {
		public VarCallFuncStructContext varCallFuncStruct() {
			return getRuleContext(VarCallFuncStructContext.class,0);
		}
		public CallFunctionStructStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStmtContext extends InstruccionesContext {
		public ContinueContext continue_() {
			return getRuleContext(ContinueContext.class,0);
		}
		public ContinueStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDeclStructStmtContext extends InstruccionesContext {
		public VarDclStructContext varDclStruct() {
			return getRuleContext(VarDclStructContext.class,0);
		}
		public VarDeclStructStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStmtContext extends InstruccionesContext {
		public SIfContext sIf() {
			return getRuleContext(SIfContext.class,0);
		}
		public IfStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionStmtContext extends InstruccionesContext {
		public FunctionsContext functions() {
			return getRuleContext(FunctionsContext.class,0);
		}
		public FunctionStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionStructStmtContext extends InstruccionesContext {
		public FunctionStructContext functionStruct() {
			return getRuleContext(FunctionStructContext.class,0);
		}
		public FunctionStructStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDeclSliceStmtContext extends InstruccionesContext {
		public VarDclSliceContext varDclSlice() {
			return getRuleContext(VarDclSliceContext.class,0);
		}
		public VarDeclSliceStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDeclStmtContext extends InstruccionesContext {
		public VarDclContext varDcl() {
			return getRuleContext(VarDclContext.class,0);
		}
		public VarDeclStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends InstruccionesContext {
		public BreakContext break_() {
			return getRuleContext(BreakContext.class,0);
		}
		public BreakStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchInstruccionContext extends InstruccionesContext {
		public SSwitchContext sSwitch() {
			return getRuleContext(SSwitchContext.class,0);
		}
		public SwitchInstruccionContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStmtContext extends InstruccionesContext {
		public SForContext sFor() {
			return getRuleContext(SForContext.class,0);
		}
		public ForStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends InstruccionesContext {
		public RetornoContext retorno() {
			return getRuleContext(RetornoContext.class,0);
		}
		public ReturnStmtContext(InstruccionesContext ctx) { copyFrom(ctx); }
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
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2161727821141067146L) != 0) || _la==ID_VARIABLE) {
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
				if (((((_la - 29)) & ~0x3f) == 0 && ((1L << (_la - 29)) & 408030265347L) != 0)) {
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
				if (((((_la - 29)) & ~0x3f) == 0 && ((1L << (_la - 29)) & 408030265347L) != 0)) {
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
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2161727821141067146L) != 0) || _la==ID_VARIABLE) {
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
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2161727821141067146L) != 0) || _la==ID_VARIABLE) {
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
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2161727821141067146L) != 0) || _la==ID_VARIABLE) {
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
	}

	public final SForContext sFor() throws RecognitionException {
		SForContext _localctx = new SForContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_sFor);
		try {
			setState(190);
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
				match(ID_VARIABLE);
				setState(189);
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
		public AssignContext assign() {
			return getRuleContext(AssignContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarDclWithTypeAndValueContext(VarDclContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDclWithInferenceContext extends VarDclContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public AssignContext assign() {
			return getRuleContext(AssignContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarDclWithInferenceContext(VarDclContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarDclWithTypeOnlyContext extends VarDclContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public VarDclWithTypeOnlyContext(VarDclContext ctx) { copyFrom(ctx); }
	}

	public final VarDclContext varDcl() throws RecognitionException {
		VarDclContext _localctx = new VarDclContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_varDcl);
		int _la;
		try {
			setState(208);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				_localctx = new VarDclWithTypeAndValueContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(192);
				match(T__15);
				setState(193);
				match(ID_VARIABLE);
				setState(194);
				type();
				setState(195);
				assign();
				setState(196);
				expr(0);
				}
				break;
			case 2:
				_localctx = new VarDclWithTypeOnlyContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(198);
				match(T__15);
				setState(199);
				match(ID_VARIABLE);
				setState(200);
				type();
				}
				break;
			case 3:
				_localctx = new VarDclWithInferenceContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(202);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__15) {
					{
					setState(201);
					match(T__15);
					}
				}

				setState(204);
				match(ID_VARIABLE);
				setState(205);
				assign();
				setState(206);
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
	}

	public final VarDclSliceContext varDclSlice() throws RecognitionException {
		VarDclSliceContext _localctx = new VarDclSliceContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_varDclSlice);
		int _la;
		try {
			setState(231);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID_VARIABLE:
				_localctx = new SliceValoresContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(210);
				match(ID_VARIABLE);
				setState(211);
				assign();
				setState(213); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(212);
					nuevoSlice();
					}
					}
					setState(215); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__18 );
				setState(217);
				type();
				setState(218);
				match(T__0);
				setState(219);
				contenidoSlice();
				setState(220);
				match(T__1);
				}
				break;
			case T__15:
				_localctx = new SliceVacioContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(222);
				match(T__15);
				setState(223);
				match(ID_VARIABLE);
				setState(225); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(224);
					nuevoSlice();
					}
					}
					setState(227); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__18 );
				setState(229);
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
	}

	public final AssignContext assign() throws RecognitionException {
		AssignContext _localctx = new AssignContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_assign);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(233);
			_la = _input.LA(1);
			if ( !(_la==T__16 || _la==T__17) ) {
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
	}

	public final NuevoSliceContext nuevoSlice() throws RecognitionException {
		NuevoSliceContext _localctx = new NuevoSliceContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_nuevoSlice);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(T__18);
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
	}

	public final ContenidoSliceContext contenidoSlice() throws RecognitionException {
		ContenidoSliceContext _localctx = new ContenidoSliceContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_contenidoSlice);
		int _la;
		try {
			setState(260);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__28:
			case T__29:
			case T__42:
			case T__43:
			case T__44:
			case T__45:
			case T__46:
			case T__47:
			case T__48:
			case T__49:
			case T__50:
			case INT:
			case DOUBLE:
			case CHAR:
			case STRING:
			case BOOL:
			case ID_VARIABLE:
				_localctx = new SliceContenidoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(237);
				expr(0);
				setState(242);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(238);
					match(T__3);
					setState(239);
					expr(0);
					}
					}
					setState(244);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case T__0:
				_localctx = new SliceContenidoSliceContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(245);
				match(T__0);
				setState(246);
				contenidoSlice();
				setState(247);
				match(T__1);
				setState(257);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(248);
					match(T__3);
					setState(253);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==T__0) {
						{
						setState(249);
						match(T__0);
						setState(250);
						contenidoSlice();
						setState(251);
						match(T__1);
						}
					}

					}
					}
					setState(259);
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
	}

	public final VarDclStructContext varDclStruct() throws RecognitionException {
		VarDclStructContext _localctx = new VarDclStructContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_varDclStruct);
		int _la;
		try {
			_localctx = new DeclStructDataContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__19) {
				{
				setState(262);
				match(T__19);
				}
			}

			setState(266);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__20) {
				{
				setState(265);
				match(T__20);
				}
			}

			setState(268);
			match(ID_VARIABLE);
			setState(270);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__20) {
				{
				setState(269);
				match(T__20);
				}
			}

			setState(272);
			match(T__0);
			setState(278); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(273);
				match(ID_VARIABLE);
				setState(274);
				type();
				setState(276);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__5) {
					{
					setState(275);
					match(T__5);
					}
				}

				}
				}
				setState(280); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID_VARIABLE );
			setState(282);
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
	}

	public final VarStructDclContext varStructDcl() throws RecognitionException {
		VarStructDclContext _localctx = new VarStructDclContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_varStructDcl);
		int _la;
		try {
			setState(320);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				_localctx = new StructVarTypeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(284);
				match(ID_VARIABLE);
				setState(285);
				match(ID_VARIABLE);
				setState(286);
				match(T__17);
				setState(287);
				match(T__0);
				setState(288);
				match(ID_VARIABLE);
				setState(289);
				match(T__11);
				setState(290);
				expr(0);
				setState(295); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(291);
					match(T__3);
					setState(292);
					match(ID_VARIABLE);
					setState(293);
					match(T__11);
					setState(294);
					expr(0);
					}
					}
					setState(297); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__3 );
				setState(299);
				match(T__1);
				}
				break;
			case 2:
				_localctx = new StructVarTypeInferenceContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(301);
				match(ID_VARIABLE);
				setState(302);
				match(T__16);
				setState(303);
				match(ID_VARIABLE);
				setState(304);
				match(T__0);
				setState(305);
				match(ID_VARIABLE);
				setState(306);
				match(T__11);
				setState(307);
				expr(0);
				setState(314); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(308);
					match(T__3);
					setState(312);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==ID_VARIABLE) {
						{
						setState(309);
						match(ID_VARIABLE);
						setState(310);
						match(T__11);
						setState(311);
						expr(0);
						}
					}

					}
					}
					setState(316); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__3 );
				setState(318);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarExprContext extends VarAsignContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarExprContext(VarAsignContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarAddContext extends VarAsignContext {
		public Token op;
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarAddContext(VarAsignContext ctx) { copyFrom(ctx); }
	}

	public final VarAsignContext varAsign() throws RecognitionException {
		VarAsignContext _localctx = new VarAsignContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_varAsign);
		int _la;
		try {
			setState(351);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				_localctx = new VarExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(322);
				match(ID_VARIABLE);
				setState(323);
				match(T__17);
				setState(324);
				expr(0);
				}
				break;
			case 2:
				_localctx = new VarAddContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(325);
				match(ID_VARIABLE);
				setState(326);
				((VarAddContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__21 || _la==T__22) ) {
					((VarAddContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(327);
				expr(0);
				}
				break;
			case 3:
				_localctx = new VarIncContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(328);
				match(ID_VARIABLE);
				setState(329);
				((VarIncContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__23 || _la==T__24) ) {
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
				setState(330);
				match(ID_VARIABLE);
				setState(335); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(331);
					match(T__25);
					setState(332);
					expr(0);
					setState(333);
					match(T__26);
					}
					}
					setState(337); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__25 );
				setState(339);
				match(T__17);
				setState(340);
				expr(0);
				}
				break;
			case 5:
				_localctx = new StructAccessAsignContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(342);
				match(ID_VARIABLE);
				setState(345); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(343);
					match(T__27);
					setState(344);
					match(ID_VARIABLE);
					}
					}
					setState(347); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__27 );
				setState(349);
				match(T__17);
				setState(350);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionValueContext extends ExprContext {
		public VarCallStatementContext varCallStatement() {
			return getRuleContext(VarCallStatementContext.class,0);
		}
		public CallFunctionValueContext(ExprContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringContext extends ExprContext {
		public TerminalNode STRING() { return getToken(gramaticaParser.STRING, 0); }
		public StringContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAccessContext extends ExprContext {
		public List<TerminalNode> ID_VARIABLE() { return getTokens(gramaticaParser.ID_VARIABLE); }
		public TerminalNode ID_VARIABLE(int i) {
			return getToken(gramaticaParser.ID_VARIABLE, i);
		}
		public StructAccessContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public IdentifierContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CharContext extends ExprContext {
		public TerminalNode CHAR() { return getToken(gramaticaParser.CHAR, 0); }
		public CharContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BooleanContext extends ExprContext {
		public TerminalNode BOOL() { return getToken(gramaticaParser.BOOL, 0); }
		public BooleanContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallFunctionStructValueContext extends ExprContext {
		public VarCallFuncStructContext varCallFuncStruct() {
			return getRuleContext(VarCallFuncStructContext.class,0);
		}
		public CallFunctionStructValueContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayFindIndexContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ArrayFindIndexContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayAppendContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ArrayAppendContext(ExprContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntToStringContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IntToStringContext(ExprContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DoubleContext extends ExprContext {
		public TerminalNode DOUBLE() { return getToken(gramaticaParser.DOUBLE, 0); }
		public DoubleContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntegerContext extends ExprContext {
		public TerminalNode INT() { return getToken(gramaticaParser.INT, 0); }
		public IntegerContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NilContext extends ExprContext {
		public NilContext(ExprContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NotContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NotContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReflectTypeContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReflectTypeContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NegateContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NegateContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayJoinContext extends ExprContext {
		public TerminalNode ID_VARIABLE() { return getToken(gramaticaParser.ID_VARIABLE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ArrayJoinContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatToStringContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public FloatToStringContext(ExprContext ctx) { copyFrom(ctx); }
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
			setState(429);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				{
				_localctx = new NegateContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(354);
				match(T__28);
				setState(355);
				expr(26);
				}
				break;
			case 2:
				{
				_localctx = new NotContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(356);
				match(T__29);
				setState(357);
				expr(25);
				}
				break;
			case 3:
				{
				_localctx = new IntegerContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(358);
				match(INT);
				}
				break;
			case 4:
				{
				_localctx = new DoubleContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(359);
				match(DOUBLE);
				}
				break;
			case 5:
				{
				_localctx = new StringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(360);
				match(STRING);
				}
				break;
			case 6:
				{
				_localctx = new BooleanContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(361);
				match(BOOL);
				}
				break;
			case 7:
				{
				_localctx = new IdentifierContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(362);
				match(ID_VARIABLE);
				}
				break;
			case 8:
				{
				_localctx = new CharContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(363);
				match(CHAR);
				}
				break;
			case 9:
				{
				_localctx = new NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(364);
				match(T__42);
				}
				break;
			case 10:
				{
				_localctx = new ParensContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(365);
				match(T__43);
				setState(366);
				expr(0);
				setState(367);
				match(T__4);
				}
				break;
			case 11:
				{
				_localctx = new ArrayAccessSimpleContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(369);
				match(ID_VARIABLE);
				setState(374); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(370);
						match(T__25);
						setState(371);
						expr(0);
						setState(372);
						match(T__26);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(376); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,39,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 12:
				{
				_localctx = new ArrayFindIndexContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(378);
				match(T__44);
				setState(379);
				match(ID_VARIABLE);
				setState(380);
				match(T__3);
				setState(381);
				expr(0);
				setState(382);
				match(T__4);
				}
				break;
			case 13:
				{
				_localctx = new ArrayJoinContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(384);
				match(T__45);
				setState(385);
				match(ID_VARIABLE);
				setState(386);
				match(T__3);
				setState(387);
				expr(0);
				setState(388);
				match(T__4);
				}
				break;
			case 14:
				{
				_localctx = new ArrayLengthContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(390);
				match(T__46);
				setState(391);
				match(ID_VARIABLE);
				setState(395);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__25) {
					{
					{
					setState(392);
					posicion();
					}
					}
					setState(397);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(398);
				match(T__4);
				}
				break;
			case 15:
				{
				_localctx = new ArrayAppendContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(399);
				match(T__47);
				setState(400);
				match(ID_VARIABLE);
				setState(401);
				match(T__3);
				setState(402);
				expr(0);
				setState(403);
				match(T__4);
				}
				break;
			case 16:
				{
				_localctx = new IntToStringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(405);
				match(T__48);
				setState(406);
				expr(0);
				setState(407);
				match(T__4);
				}
				break;
			case 17:
				{
				_localctx = new FloatToStringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(409);
				match(T__49);
				setState(410);
				expr(0);
				setState(411);
				match(T__4);
				}
				break;
			case 18:
				{
				_localctx = new ReflectTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(413);
				match(T__50);
				setState(414);
				expr(0);
				setState(415);
				match(T__4);
				}
				break;
			case 19:
				{
				_localctx = new StructAccessContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(417);
				match(ID_VARIABLE);
				setState(420); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(418);
						match(T__27);
						setState(419);
						match(ID_VARIABLE);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(422); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(425);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
				case 1:
					{
					setState(424);
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
				setState(427);
				varCallStatement();
				}
				break;
			case 21:
				{
				_localctx = new CallFunctionStructValueContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(428);
				varCallFuncStruct();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(448);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(446);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
					case 1:
						{
						_localctx = new MulDivModuloContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(431);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(432);
						((MulDivModuloContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 15032385536L) != 0)) ) {
							((MulDivModuloContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(433);
						expr(25);
						}
						break;
					case 2:
						{
						_localctx = new AddSubContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(434);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(435);
						((AddSubContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__28 || _la==T__33) ) {
							((AddSubContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(436);
						expr(24);
						}
						break;
					case 3:
						{
						_localctx = new MinorMajorEqualContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(437);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(438);
						((MinorMajorEqualContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 515396075520L) != 0)) ) {
							((MinorMajorEqualContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(439);
						expr(23);
						}
						break;
					case 4:
						{
						_localctx = new EqualsNotEqualsContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(440);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(441);
						((EqualsNotEqualsContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__38 || _la==T__39) ) {
							((EqualsNotEqualsContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(442);
						expr(22);
						}
						break;
					case 5:
						{
						_localctx = new LogicalContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(443);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(444);
						((LogicalContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__40 || _la==T__41) ) {
							((LogicalContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(445);
						expr(21);
						}
						break;
					}
					} 
				}
				setState(450);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
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
	}

	public final PosicionContext posicion() throws RecognitionException {
		PosicionContext _localctx = new PosicionContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_posicion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(451);
			match(T__25);
			setState(452);
			expr(0);
			setState(453);
			match(T__26);
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
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(455);
			_la = _input.LA(1);
			if ( !(((((_la - 21)) & ~0x3f) == 0 && ((1L << (_la - 21)) & 70435316170753L) != 0)) ) {
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
	}

	public final BreakContext break_() throws RecognitionException {
		BreakContext _localctx = new BreakContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_break);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(457);
			match(T__56);
			setState(459);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(458);
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
	}

	public final ContinueContext continue_() throws RecognitionException {
		ContinueContext _localctx = new ContinueContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_continue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(461);
			match(T__57);
			setState(463);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(462);
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
	}

	public final FunctionsContext functions() throws RecognitionException {
		FunctionsContext _localctx = new FunctionsContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_functions);
		int _la;
		try {
			_localctx = new FuncionesContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(465);
			match(T__58);
			setState(466);
			match(ID_VARIABLE);
			setState(467);
			match(T__43);
			setState(478);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID_VARIABLE) {
				{
				setState(468);
				match(ID_VARIABLE);
				setState(469);
				type();
				setState(475);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(470);
					match(T__3);
					setState(471);
					match(ID_VARIABLE);
					setState(472);
					type();
					}
					}
					setState(477);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(480);
			match(T__4);
			setState(482);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 21)) & ~0x3f) == 0 && ((1L << (_la - 21)) & 70435316170753L) != 0)) {
				{
				setState(481);
				valRet();
				}
			}

			setState(484);
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
	}

	public final FunctionStructContext functionStruct() throws RecognitionException {
		FunctionStructContext _localctx = new FunctionStructContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_functionStruct);
		int _la;
		try {
			_localctx = new FuncionesStructsNativasContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(486);
			match(T__58);
			setState(487);
			match(T__43);
			setState(488);
			match(ID_VARIABLE);
			setState(489);
			match(ID_VARIABLE);
			setState(490);
			match(T__4);
			setState(491);
			match(ID_VARIABLE);
			setState(492);
			match(T__43);
			setState(494);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID_VARIABLE) {
				{
				setState(493);
				defParams();
				}
			}

			setState(496);
			match(T__4);
			setState(498);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 21)) & ~0x3f) == 0 && ((1L << (_la - 21)) & 70435316170753L) != 0)) {
				{
				setState(497);
				valRet();
				}
			}

			setState(500);
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
	}

	public final DefParamsContext defParams() throws RecognitionException {
		DefParamsContext _localctx = new DefParamsContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_defParams);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(502);
			match(ID_VARIABLE);
			setState(503);
			type();
			setState(509);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(504);
				match(T__3);
				setState(505);
				match(ID_VARIABLE);
				setState(506);
				type();
				}
				}
				setState(511);
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
	}

	public final VarCallStatementContext varCallStatement() throws RecognitionException {
		VarCallStatementContext _localctx = new VarCallStatementContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_varCallStatement);
		int _la;
		try {
			_localctx = new CallFunctionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(512);
			match(ID_VARIABLE);
			setState(513);
			match(T__43);
			setState(522);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 29)) & ~0x3f) == 0 && ((1L << (_la - 29)) & 408030265347L) != 0)) {
				{
				setState(514);
				expr(0);
				setState(519);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(515);
					match(T__3);
					setState(516);
					expr(0);
					}
					}
					setState(521);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(524);
			match(T__4);
			setState(526);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,56,_ctx) ) {
			case 1:
				{
				setState(525);
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
	}

	public final VarCallFuncStructContext varCallFuncStruct() throws RecognitionException {
		VarCallFuncStructContext _localctx = new VarCallFuncStructContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_varCallFuncStruct);
		int _la;
		try {
			_localctx = new CallFunctionStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(528);
			match(ID_VARIABLE);
			setState(529);
			match(T__27);
			setState(530);
			match(ID_VARIABLE);
			setState(531);
			match(T__43);
			setState(540);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 29)) & ~0x3f) == 0 && ((1L << (_la - 29)) & 408030265347L) != 0)) {
				{
				setState(532);
				expr(0);
				setState(537);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(533);
					match(T__3);
					setState(534);
					expr(0);
					}
					}
					setState(539);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(542);
			match(T__4);
			setState(544);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,59,_ctx) ) {
			case 1:
				{
				setState(543);
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
	}

	public final ValRetContext valRet() throws RecognitionException {
		ValRetContext _localctx = new ValRetContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_valRet);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(546);
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
	}

	public final RetornoContext retorno() throws RecognitionException {
		RetornoContext _localctx = new RetornoContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_retorno);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(548);
			match(T__59);
			setState(550);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,60,_ctx) ) {
			case 1:
				{
				setState(549);
				expr(0);
				}
				break;
			}
			setState(553);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(552);
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
		"\u0004\u0001E\u022c\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"\u0001\u0007\u0001\u0007\u0003\u0007\u00bf\b\u0007\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u00cb"+
		"\b\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u00d1\b\b\u0001\t\u0001\t"+
		"\u0001\t\u0004\t\u00d6\b\t\u000b\t\f\t\u00d7\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0004\t\u00e2\b\t\u000b\t\f\t\u00e3"+
		"\u0001\t\u0001\t\u0003\t\u00e8\b\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0005\f\u00f1\b\f\n\f\f\f\u00f4\t\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u00fe\b\f\u0005"+
		"\f\u0100\b\f\n\f\f\f\u0103\t\f\u0003\f\u0105\b\f\u0001\r\u0003\r\u0108"+
		"\b\r\u0001\r\u0003\r\u010b\b\r\u0001\r\u0001\r\u0003\r\u010f\b\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0003\r\u0115\b\r\u0004\r\u0117\b\r\u000b\r"+
		"\f\r\u0118\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0004\u000e\u0128\b\u000e\u000b\u000e\f\u000e\u0129\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0003\u000e\u0139\b\u000e\u0004\u000e\u013b\b\u000e\u000b\u000e"+
		"\f\u000e\u013c\u0001\u000e\u0001\u000e\u0003\u000e\u0141\b\u000e\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0004\u000f\u0150\b\u000f\u000b\u000f\f\u000f\u0151\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0004\u000f"+
		"\u015a\b\u000f\u000b\u000f\f\u000f\u015b\u0001\u000f\u0001\u000f\u0003"+
		"\u000f\u0160\b\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0004\u0010\u0177"+
		"\b\u0010\u000b\u0010\f\u0010\u0178\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0005\u0010\u018a\b\u0010\n\u0010\f\u0010\u018d\t\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0004\u0010\u01a5\b\u0010\u000b\u0010\f"+
		"\u0010\u01a6\u0001\u0010\u0003\u0010\u01aa\b\u0010\u0001\u0010\u0001\u0010"+
		"\u0003\u0010\u01ae\b\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010"+
		"\u01bf\b\u0010\n\u0010\f\u0010\u01c2\t\u0010\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0003"+
		"\u0013\u01cc\b\u0013\u0001\u0014\u0001\u0014\u0003\u0014\u01d0\b\u0014"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0005\u0015\u01da\b\u0015\n\u0015\f\u0015\u01dd"+
		"\t\u0015\u0003\u0015\u01df\b\u0015\u0001\u0015\u0001\u0015\u0003\u0015"+
		"\u01e3\b\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016"+
		"\u01ef\b\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u01f3\b\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0005\u0017\u01fc\b\u0017\n\u0017\f\u0017\u01ff\t\u0017\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0005\u0018\u0206\b\u0018"+
		"\n\u0018\f\u0018\u0209\t\u0018\u0003\u0018\u020b\b\u0018\u0001\u0018\u0001"+
		"\u0018\u0003\u0018\u020f\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u0218\b\u0019\n"+
		"\u0019\f\u0019\u021b\t\u0019\u0003\u0019\u021d\b\u0019\u0001\u0019\u0001"+
		"\u0019\u0003\u0019\u0221\b\u0019\u0001\u001a\u0001\u001a\u0001\u001b\u0001"+
		"\u001b\u0003\u001b\u0227\b\u001b\u0001\u001b\u0003\u001b\u022a\b\u001b"+
		"\u0001\u001b\u0000\u0001 \u001c\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.0246\u0000\t\u0001"+
		"\u0000\u0011\u0012\u0001\u0000\u0016\u0017\u0001\u0000\u0018\u0019\u0001"+
		"\u0000\u001f!\u0002\u0000\u001d\u001d\"\"\u0001\u0000#&\u0001\u0000\'"+
		"(\u0001\u0000)*\u0003\u0000\u0015\u001548CC\u0277\u0000;\u0001\u0000\u0000"+
		"\u0000\u0002V\u0001\u0000\u0000\u0000\u0004v\u0001\u0000\u0000\u0000\u0006"+
		"\u0085\u0001\u0000\u0000\u0000\b\u0087\u0001\u0000\u0000\u0000\n\u00a0"+
		"\u0001\u0000\u0000\u0000\f\u00a2\u0001\u0000\u0000\u0000\u000e\u00be\u0001"+
		"\u0000\u0000\u0000\u0010\u00d0\u0001\u0000\u0000\u0000\u0012\u00e7\u0001"+
		"\u0000\u0000\u0000\u0014\u00e9\u0001\u0000\u0000\u0000\u0016\u00eb\u0001"+
		"\u0000\u0000\u0000\u0018\u0104\u0001\u0000\u0000\u0000\u001a\u0107\u0001"+
		"\u0000\u0000\u0000\u001c\u0140\u0001\u0000\u0000\u0000\u001e\u015f\u0001"+
		"\u0000\u0000\u0000 \u01ad\u0001\u0000\u0000\u0000\"\u01c3\u0001\u0000"+
		"\u0000\u0000$\u01c7\u0001\u0000\u0000\u0000&\u01c9\u0001\u0000\u0000\u0000"+
		"(\u01cd\u0001\u0000\u0000\u0000*\u01d1\u0001\u0000\u0000\u0000,\u01e6"+
		"\u0001\u0000\u0000\u0000.\u01f6\u0001\u0000\u0000\u00000\u0200\u0001\u0000"+
		"\u0000\u00002\u0210\u0001\u0000\u0000\u00004\u0222\u0001\u0000\u0000\u0000"+
		"6\u0224\u0001\u0000\u0000\u00008:\u0003\u0002\u0001\u000098\u0001\u0000"+
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
		"\f\u0006\u0000\u00ae\u00bf\u0001\u0000\u0000\u0000\u00af\u00b0\u0005\u000e"+
		"\u0000\u0000\u00b0\u00b1\u0003\u0010\b\u0000\u00b1\u00b2\u0005\u0006\u0000"+
		"\u0000\u00b2\u00b3\u0003 \u0010\u0000\u00b3\u00b4\u0005\u0006\u0000\u0000"+
		"\u00b4\u00b5\u0003\u001e\u000f\u0000\u00b5\u00b6\u0003\f\u0006\u0000\u00b6"+
		"\u00bf\u0001\u0000\u0000\u0000\u00b7\u00b8\u0005\u000e\u0000\u0000\u00b8"+
		"\u00b9\u0005C\u0000\u0000\u00b9\u00ba\u0005\u0004\u0000\u0000\u00ba\u00bb"+
		"\u0005C\u0000\u0000\u00bb\u00bc\u0005\u000f\u0000\u0000\u00bc\u00bd\u0005"+
		"C\u0000\u0000\u00bd\u00bf\u0003\f\u0006\u0000\u00be\u00ab\u0001\u0000"+
		"\u0000\u0000\u00be\u00af\u0001\u0000\u0000\u0000\u00be\u00b7\u0001\u0000"+
		"\u0000\u0000\u00bf\u000f\u0001\u0000\u0000\u0000\u00c0\u00c1\u0005\u0010"+
		"\u0000\u0000\u00c1\u00c2\u0005C\u0000\u0000\u00c2\u00c3\u0003$\u0012\u0000"+
		"\u00c3\u00c4\u0003\u0014\n\u0000\u00c4\u00c5\u0003 \u0010\u0000\u00c5"+
		"\u00d1\u0001\u0000\u0000\u0000\u00c6\u00c7\u0005\u0010\u0000\u0000\u00c7"+
		"\u00c8\u0005C\u0000\u0000\u00c8\u00d1\u0003$\u0012\u0000\u00c9\u00cb\u0005"+
		"\u0010\u0000\u0000\u00ca\u00c9\u0001\u0000\u0000\u0000\u00ca\u00cb\u0001"+
		"\u0000\u0000\u0000\u00cb\u00cc\u0001\u0000\u0000\u0000\u00cc\u00cd\u0005"+
		"C\u0000\u0000\u00cd\u00ce\u0003\u0014\n\u0000\u00ce\u00cf\u0003 \u0010"+
		"\u0000\u00cf\u00d1\u0001\u0000\u0000\u0000\u00d0\u00c0\u0001\u0000\u0000"+
		"\u0000\u00d0\u00c6\u0001\u0000\u0000\u0000\u00d0\u00ca\u0001\u0000\u0000"+
		"\u0000\u00d1\u0011\u0001\u0000\u0000\u0000\u00d2\u00d3\u0005C\u0000\u0000"+
		"\u00d3\u00d5\u0003\u0014\n\u0000\u00d4\u00d6\u0003\u0016\u000b\u0000\u00d5"+
		"\u00d4\u0001\u0000\u0000\u0000\u00d6\u00d7\u0001\u0000\u0000\u0000\u00d7"+
		"\u00d5\u0001\u0000\u0000\u0000\u00d7\u00d8\u0001\u0000\u0000\u0000\u00d8"+
		"\u00d9\u0001\u0000\u0000\u0000\u00d9\u00da\u0003$\u0012\u0000\u00da\u00db"+
		"\u0005\u0001\u0000\u0000\u00db\u00dc\u0003\u0018\f\u0000\u00dc\u00dd\u0005"+
		"\u0002\u0000\u0000\u00dd\u00e8\u0001\u0000\u0000\u0000\u00de\u00df\u0005"+
		"\u0010\u0000\u0000\u00df\u00e1\u0005C\u0000\u0000\u00e0\u00e2\u0003\u0016"+
		"\u000b\u0000\u00e1\u00e0\u0001\u0000\u0000\u0000\u00e2\u00e3\u0001\u0000"+
		"\u0000\u0000\u00e3\u00e1\u0001\u0000\u0000\u0000\u00e3\u00e4\u0001\u0000"+
		"\u0000\u0000\u00e4\u00e5\u0001\u0000\u0000\u0000\u00e5\u00e6\u0003$\u0012"+
		"\u0000\u00e6\u00e8\u0001\u0000\u0000\u0000\u00e7\u00d2\u0001\u0000\u0000"+
		"\u0000\u00e7\u00de\u0001\u0000\u0000\u0000\u00e8\u0013\u0001\u0000\u0000"+
		"\u0000\u00e9\u00ea\u0007\u0000\u0000\u0000\u00ea\u0015\u0001\u0000\u0000"+
		"\u0000\u00eb\u00ec\u0005\u0013\u0000\u0000\u00ec\u0017\u0001\u0000\u0000"+
		"\u0000\u00ed\u00f2\u0003 \u0010\u0000\u00ee\u00ef\u0005\u0004\u0000\u0000"+
		"\u00ef\u00f1\u0003 \u0010\u0000\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f1"+
		"\u00f4\u0001\u0000\u0000\u0000\u00f2\u00f0\u0001\u0000\u0000\u0000\u00f2"+
		"\u00f3\u0001\u0000\u0000\u0000\u00f3\u0105\u0001\u0000\u0000\u0000\u00f4"+
		"\u00f2\u0001\u0000\u0000\u0000\u00f5\u00f6\u0005\u0001\u0000\u0000\u00f6"+
		"\u00f7\u0003\u0018\f\u0000\u00f7\u0101\u0005\u0002\u0000\u0000\u00f8\u00fd"+
		"\u0005\u0004\u0000\u0000\u00f9\u00fa\u0005\u0001\u0000\u0000\u00fa\u00fb"+
		"\u0003\u0018\f\u0000\u00fb\u00fc\u0005\u0002\u0000\u0000\u00fc\u00fe\u0001"+
		"\u0000\u0000\u0000\u00fd\u00f9\u0001\u0000\u0000\u0000\u00fd\u00fe\u0001"+
		"\u0000\u0000\u0000\u00fe\u0100\u0001\u0000\u0000\u0000\u00ff\u00f8\u0001"+
		"\u0000\u0000\u0000\u0100\u0103\u0001\u0000\u0000\u0000\u0101\u00ff\u0001"+
		"\u0000\u0000\u0000\u0101\u0102\u0001\u0000\u0000\u0000\u0102\u0105\u0001"+
		"\u0000\u0000\u0000\u0103\u0101\u0001\u0000\u0000\u0000\u0104\u00ed\u0001"+
		"\u0000\u0000\u0000\u0104\u00f5\u0001\u0000\u0000\u0000\u0105\u0019\u0001"+
		"\u0000\u0000\u0000\u0106\u0108\u0005\u0014\u0000\u0000\u0107\u0106\u0001"+
		"\u0000\u0000\u0000\u0107\u0108\u0001\u0000\u0000\u0000\u0108\u010a\u0001"+
		"\u0000\u0000\u0000\u0109\u010b\u0005\u0015\u0000\u0000\u010a\u0109\u0001"+
		"\u0000\u0000\u0000\u010a\u010b\u0001\u0000\u0000\u0000\u010b\u010c\u0001"+
		"\u0000\u0000\u0000\u010c\u010e\u0005C\u0000\u0000\u010d\u010f\u0005\u0015"+
		"\u0000\u0000\u010e\u010d\u0001\u0000\u0000\u0000\u010e\u010f\u0001\u0000"+
		"\u0000\u0000\u010f\u0110\u0001\u0000\u0000\u0000\u0110\u0116\u0005\u0001"+
		"\u0000\u0000\u0111\u0112\u0005C\u0000\u0000\u0112\u0114\u0003$\u0012\u0000"+
		"\u0113\u0115\u0005\u0006\u0000\u0000\u0114\u0113\u0001\u0000\u0000\u0000"+
		"\u0114\u0115\u0001\u0000\u0000\u0000\u0115\u0117\u0001\u0000\u0000\u0000"+
		"\u0116\u0111\u0001\u0000\u0000\u0000\u0117\u0118\u0001\u0000\u0000\u0000"+
		"\u0118\u0116\u0001\u0000\u0000\u0000\u0118\u0119\u0001\u0000\u0000\u0000"+
		"\u0119\u011a\u0001\u0000\u0000\u0000\u011a\u011b\u0005\u0002\u0000\u0000"+
		"\u011b\u001b\u0001\u0000\u0000\u0000\u011c\u011d\u0005C\u0000\u0000\u011d"+
		"\u011e\u0005C\u0000\u0000\u011e\u011f\u0005\u0012\u0000\u0000\u011f\u0120"+
		"\u0005\u0001\u0000\u0000\u0120\u0121\u0005C\u0000\u0000\u0121\u0122\u0005"+
		"\f\u0000\u0000\u0122\u0127\u0003 \u0010\u0000\u0123\u0124\u0005\u0004"+
		"\u0000\u0000\u0124\u0125\u0005C\u0000\u0000\u0125\u0126\u0005\f\u0000"+
		"\u0000\u0126\u0128\u0003 \u0010\u0000\u0127\u0123\u0001\u0000\u0000\u0000"+
		"\u0128\u0129\u0001\u0000\u0000\u0000\u0129\u0127\u0001\u0000\u0000\u0000"+
		"\u0129\u012a\u0001\u0000\u0000\u0000\u012a\u012b\u0001\u0000\u0000\u0000"+
		"\u012b\u012c\u0005\u0002\u0000\u0000\u012c\u0141\u0001\u0000\u0000\u0000"+
		"\u012d\u012e\u0005C\u0000\u0000\u012e\u012f\u0005\u0011\u0000\u0000\u012f"+
		"\u0130\u0005C\u0000\u0000\u0130\u0131\u0005\u0001\u0000\u0000\u0131\u0132"+
		"\u0005C\u0000\u0000\u0132\u0133\u0005\f\u0000\u0000\u0133\u013a\u0003"+
		" \u0010\u0000\u0134\u0138\u0005\u0004\u0000\u0000\u0135\u0136\u0005C\u0000"+
		"\u0000\u0136\u0137\u0005\f\u0000\u0000\u0137\u0139\u0003 \u0010\u0000"+
		"\u0138\u0135\u0001\u0000\u0000\u0000\u0138\u0139\u0001\u0000\u0000\u0000"+
		"\u0139\u013b\u0001\u0000\u0000\u0000\u013a\u0134\u0001\u0000\u0000\u0000"+
		"\u013b\u013c\u0001\u0000\u0000\u0000\u013c\u013a\u0001\u0000\u0000\u0000"+
		"\u013c\u013d\u0001\u0000\u0000\u0000\u013d\u013e\u0001\u0000\u0000\u0000"+
		"\u013e\u013f\u0005\u0002\u0000\u0000\u013f\u0141\u0001\u0000\u0000\u0000"+
		"\u0140\u011c\u0001\u0000\u0000\u0000\u0140\u012d\u0001\u0000\u0000\u0000"+
		"\u0141\u001d\u0001\u0000\u0000\u0000\u0142\u0143\u0005C\u0000\u0000\u0143"+
		"\u0144\u0005\u0012\u0000\u0000\u0144\u0160\u0003 \u0010\u0000\u0145\u0146"+
		"\u0005C\u0000\u0000\u0146\u0147\u0007\u0001\u0000\u0000\u0147\u0160\u0003"+
		" \u0010\u0000\u0148\u0149\u0005C\u0000\u0000\u0149\u0160\u0007\u0002\u0000"+
		"\u0000\u014a\u014f\u0005C\u0000\u0000\u014b\u014c\u0005\u001a\u0000\u0000"+
		"\u014c\u014d\u0003 \u0010\u0000\u014d\u014e\u0005\u001b\u0000\u0000\u014e"+
		"\u0150\u0001\u0000\u0000\u0000\u014f\u014b\u0001\u0000\u0000\u0000\u0150"+
		"\u0151\u0001\u0000\u0000\u0000\u0151\u014f\u0001\u0000\u0000\u0000\u0151"+
		"\u0152\u0001\u0000\u0000\u0000\u0152\u0153\u0001\u0000\u0000\u0000\u0153"+
		"\u0154\u0005\u0012\u0000\u0000\u0154\u0155\u0003 \u0010\u0000\u0155\u0160"+
		"\u0001\u0000\u0000\u0000\u0156\u0159\u0005C\u0000\u0000\u0157\u0158\u0005"+
		"\u001c\u0000\u0000\u0158\u015a\u0005C\u0000\u0000\u0159\u0157\u0001\u0000"+
		"\u0000\u0000\u015a\u015b\u0001\u0000\u0000\u0000\u015b\u0159\u0001\u0000"+
		"\u0000\u0000\u015b\u015c\u0001\u0000\u0000\u0000\u015c\u015d\u0001\u0000"+
		"\u0000\u0000\u015d\u015e\u0005\u0012\u0000\u0000\u015e\u0160\u0003 \u0010"+
		"\u0000\u015f\u0142\u0001\u0000\u0000\u0000\u015f\u0145\u0001\u0000\u0000"+
		"\u0000\u015f\u0148\u0001\u0000\u0000\u0000\u015f\u014a\u0001\u0000\u0000"+
		"\u0000\u015f\u0156\u0001\u0000\u0000\u0000\u0160\u001f\u0001\u0000\u0000"+
		"\u0000\u0161\u0162\u0006\u0010\uffff\uffff\u0000\u0162\u0163\u0005\u001d"+
		"\u0000\u0000\u0163\u01ae\u0003 \u0010\u001a\u0164\u0165\u0005\u001e\u0000"+
		"\u0000\u0165\u01ae\u0003 \u0010\u0019\u0166\u01ae\u0005=\u0000\u0000\u0167"+
		"\u01ae\u0005>\u0000\u0000\u0168\u01ae\u0005@\u0000\u0000\u0169\u01ae\u0005"+
		"A\u0000\u0000\u016a\u01ae\u0005C\u0000\u0000\u016b\u01ae\u0005?\u0000"+
		"\u0000\u016c\u01ae\u0005+\u0000\u0000\u016d\u016e\u0005,\u0000\u0000\u016e"+
		"\u016f\u0003 \u0010\u0000\u016f\u0170\u0005\u0005\u0000\u0000\u0170\u01ae"+
		"\u0001\u0000\u0000\u0000\u0171\u0176\u0005C\u0000\u0000\u0172\u0173\u0005"+
		"\u001a\u0000\u0000\u0173\u0174\u0003 \u0010\u0000\u0174\u0175\u0005\u001b"+
		"\u0000\u0000\u0175\u0177\u0001\u0000\u0000\u0000\u0176\u0172\u0001\u0000"+
		"\u0000\u0000\u0177\u0178\u0001\u0000\u0000\u0000\u0178\u0176\u0001\u0000"+
		"\u0000\u0000\u0178\u0179\u0001\u0000\u0000\u0000\u0179\u01ae\u0001\u0000"+
		"\u0000\u0000\u017a\u017b\u0005-\u0000\u0000\u017b\u017c\u0005C\u0000\u0000"+
		"\u017c\u017d\u0005\u0004\u0000\u0000\u017d\u017e\u0003 \u0010\u0000\u017e"+
		"\u017f\u0005\u0005\u0000\u0000\u017f\u01ae\u0001\u0000\u0000\u0000\u0180"+
		"\u0181\u0005.\u0000\u0000\u0181\u0182\u0005C\u0000\u0000\u0182\u0183\u0005"+
		"\u0004\u0000\u0000\u0183\u0184\u0003 \u0010\u0000\u0184\u0185\u0005\u0005"+
		"\u0000\u0000\u0185\u01ae\u0001\u0000\u0000\u0000\u0186\u0187\u0005/\u0000"+
		"\u0000\u0187\u018b\u0005C\u0000\u0000\u0188\u018a\u0003\"\u0011\u0000"+
		"\u0189\u0188\u0001\u0000\u0000\u0000\u018a\u018d\u0001\u0000\u0000\u0000"+
		"\u018b\u0189\u0001\u0000\u0000\u0000\u018b\u018c\u0001\u0000\u0000\u0000"+
		"\u018c\u018e\u0001\u0000\u0000\u0000\u018d\u018b\u0001\u0000\u0000\u0000"+
		"\u018e\u01ae\u0005\u0005\u0000\u0000\u018f\u0190\u00050\u0000\u0000\u0190"+
		"\u0191\u0005C\u0000\u0000\u0191\u0192\u0005\u0004\u0000\u0000\u0192\u0193"+
		"\u0003 \u0010\u0000\u0193\u0194\u0005\u0005\u0000\u0000\u0194\u01ae\u0001"+
		"\u0000\u0000\u0000\u0195\u0196\u00051\u0000\u0000\u0196\u0197\u0003 \u0010"+
		"\u0000\u0197\u0198\u0005\u0005\u0000\u0000\u0198\u01ae\u0001\u0000\u0000"+
		"\u0000\u0199\u019a\u00052\u0000\u0000\u019a\u019b\u0003 \u0010\u0000\u019b"+
		"\u019c\u0005\u0005\u0000\u0000\u019c\u01ae\u0001\u0000\u0000\u0000\u019d"+
		"\u019e\u00053\u0000\u0000\u019e\u019f\u0003 \u0010\u0000\u019f\u01a0\u0005"+
		"\u0005\u0000\u0000\u01a0\u01ae\u0001\u0000\u0000\u0000\u01a1\u01a4\u0005"+
		"C\u0000\u0000\u01a2\u01a3\u0005\u001c\u0000\u0000\u01a3\u01a5\u0005C\u0000"+
		"\u0000\u01a4\u01a2\u0001\u0000\u0000\u0000\u01a5\u01a6\u0001\u0000\u0000"+
		"\u0000\u01a6\u01a4\u0001\u0000\u0000\u0000\u01a6\u01a7\u0001\u0000\u0000"+
		"\u0000\u01a7\u01a9\u0001\u0000\u0000\u0000\u01a8\u01aa\u0005\u0006\u0000"+
		"\u0000\u01a9\u01a8\u0001\u0000\u0000\u0000\u01a9\u01aa\u0001\u0000\u0000"+
		"\u0000\u01aa\u01ae\u0001\u0000\u0000\u0000\u01ab\u01ae\u00030\u0018\u0000"+
		"\u01ac\u01ae\u00032\u0019\u0000\u01ad\u0161\u0001\u0000\u0000\u0000\u01ad"+
		"\u0164\u0001\u0000\u0000\u0000\u01ad\u0166\u0001\u0000\u0000\u0000\u01ad"+
		"\u0167\u0001\u0000\u0000\u0000\u01ad\u0168\u0001\u0000\u0000\u0000\u01ad"+
		"\u0169\u0001\u0000\u0000\u0000\u01ad\u016a\u0001\u0000\u0000\u0000\u01ad"+
		"\u016b\u0001\u0000\u0000\u0000\u01ad\u016c\u0001\u0000\u0000\u0000\u01ad"+
		"\u016d\u0001\u0000\u0000\u0000\u01ad\u0171\u0001\u0000\u0000\u0000\u01ad"+
		"\u017a\u0001\u0000\u0000\u0000\u01ad\u0180\u0001\u0000\u0000\u0000\u01ad"+
		"\u0186\u0001\u0000\u0000\u0000\u01ad\u018f\u0001\u0000\u0000\u0000\u01ad"+
		"\u0195\u0001\u0000\u0000\u0000\u01ad\u0199\u0001\u0000\u0000\u0000\u01ad"+
		"\u019d\u0001\u0000\u0000\u0000\u01ad\u01a1\u0001\u0000\u0000\u0000\u01ad"+
		"\u01ab\u0001\u0000\u0000\u0000\u01ad\u01ac\u0001\u0000\u0000\u0000\u01ae"+
		"\u01c0\u0001\u0000\u0000\u0000\u01af\u01b0\n\u0018\u0000\u0000\u01b0\u01b1"+
		"\u0007\u0003\u0000\u0000\u01b1\u01bf\u0003 \u0010\u0019\u01b2\u01b3\n"+
		"\u0017\u0000\u0000\u01b3\u01b4\u0007\u0004\u0000\u0000\u01b4\u01bf\u0003"+
		" \u0010\u0018\u01b5\u01b6\n\u0016\u0000\u0000\u01b6\u01b7\u0007\u0005"+
		"\u0000\u0000\u01b7\u01bf\u0003 \u0010\u0017\u01b8\u01b9\n\u0015\u0000"+
		"\u0000\u01b9\u01ba\u0007\u0006\u0000\u0000\u01ba\u01bf\u0003 \u0010\u0016"+
		"\u01bb\u01bc\n\u0014\u0000\u0000\u01bc\u01bd\u0007\u0007\u0000\u0000\u01bd"+
		"\u01bf\u0003 \u0010\u0015\u01be\u01af\u0001\u0000\u0000\u0000\u01be\u01b2"+
		"\u0001\u0000\u0000\u0000\u01be\u01b5\u0001\u0000\u0000\u0000\u01be\u01b8"+
		"\u0001\u0000\u0000\u0000\u01be\u01bb\u0001\u0000\u0000\u0000\u01bf\u01c2"+
		"\u0001\u0000\u0000\u0000\u01c0\u01be\u0001\u0000\u0000\u0000\u01c0\u01c1"+
		"\u0001\u0000\u0000\u0000\u01c1!\u0001\u0000\u0000\u0000\u01c2\u01c0\u0001"+
		"\u0000\u0000\u0000\u01c3\u01c4\u0005\u001a\u0000\u0000\u01c4\u01c5\u0003"+
		" \u0010\u0000\u01c5\u01c6\u0005\u001b\u0000\u0000\u01c6#\u0001\u0000\u0000"+
		"\u0000\u01c7\u01c8\u0007\b\u0000\u0000\u01c8%\u0001\u0000\u0000\u0000"+
		"\u01c9\u01cb\u00059\u0000\u0000\u01ca\u01cc\u0005\u0006\u0000\u0000\u01cb"+
		"\u01ca\u0001\u0000\u0000\u0000\u01cb\u01cc\u0001\u0000\u0000\u0000\u01cc"+
		"\'\u0001\u0000\u0000\u0000\u01cd\u01cf\u0005:\u0000\u0000\u01ce\u01d0"+
		"\u0005\u0006\u0000\u0000\u01cf\u01ce\u0001\u0000\u0000\u0000\u01cf\u01d0"+
		"\u0001\u0000\u0000\u0000\u01d0)\u0001\u0000\u0000\u0000\u01d1\u01d2\u0005"+
		";\u0000\u0000\u01d2\u01d3\u0005C\u0000\u0000\u01d3\u01de\u0005,\u0000"+
		"\u0000\u01d4\u01d5\u0005C\u0000\u0000\u01d5\u01db\u0003$\u0012\u0000\u01d6"+
		"\u01d7\u0005\u0004\u0000\u0000\u01d7\u01d8\u0005C\u0000\u0000\u01d8\u01da"+
		"\u0003$\u0012\u0000\u01d9\u01d6\u0001\u0000\u0000\u0000\u01da\u01dd\u0001"+
		"\u0000\u0000\u0000\u01db\u01d9\u0001\u0000\u0000\u0000\u01db\u01dc\u0001"+
		"\u0000\u0000\u0000\u01dc\u01df\u0001\u0000\u0000\u0000\u01dd\u01db\u0001"+
		"\u0000\u0000\u0000\u01de\u01d4\u0001\u0000\u0000\u0000\u01de\u01df\u0001"+
		"\u0000\u0000\u0000\u01df\u01e0\u0001\u0000\u0000\u0000\u01e0\u01e2\u0005"+
		"\u0005\u0000\u0000\u01e1\u01e3\u00034\u001a\u0000\u01e2\u01e1\u0001\u0000"+
		"\u0000\u0000\u01e2\u01e3\u0001\u0000\u0000\u0000\u01e3\u01e4\u0001\u0000"+
		"\u0000\u0000\u01e4\u01e5\u0003\f\u0006\u0000\u01e5+\u0001\u0000\u0000"+
		"\u0000\u01e6\u01e7\u0005;\u0000\u0000\u01e7\u01e8\u0005,\u0000\u0000\u01e8"+
		"\u01e9\u0005C\u0000\u0000\u01e9\u01ea\u0005C\u0000\u0000\u01ea\u01eb\u0005"+
		"\u0005\u0000\u0000\u01eb\u01ec\u0005C\u0000\u0000\u01ec\u01ee\u0005,\u0000"+
		"\u0000\u01ed\u01ef\u0003.\u0017\u0000\u01ee\u01ed\u0001\u0000\u0000\u0000"+
		"\u01ee\u01ef\u0001\u0000\u0000\u0000\u01ef\u01f0\u0001\u0000\u0000\u0000"+
		"\u01f0\u01f2\u0005\u0005\u0000\u0000\u01f1\u01f3\u00034\u001a\u0000\u01f2"+
		"\u01f1\u0001\u0000\u0000\u0000\u01f2\u01f3\u0001\u0000\u0000\u0000\u01f3"+
		"\u01f4\u0001\u0000\u0000\u0000\u01f4\u01f5\u0003\f\u0006\u0000\u01f5-"+
		"\u0001\u0000\u0000\u0000\u01f6\u01f7\u0005C\u0000\u0000\u01f7\u01fd\u0003"+
		"$\u0012\u0000\u01f8\u01f9\u0005\u0004\u0000\u0000\u01f9\u01fa\u0005C\u0000"+
		"\u0000\u01fa\u01fc\u0003$\u0012\u0000\u01fb\u01f8\u0001\u0000\u0000\u0000"+
		"\u01fc\u01ff\u0001\u0000\u0000\u0000\u01fd\u01fb\u0001\u0000\u0000\u0000"+
		"\u01fd\u01fe\u0001\u0000\u0000\u0000\u01fe/\u0001\u0000\u0000\u0000\u01ff"+
		"\u01fd\u0001\u0000\u0000\u0000\u0200\u0201\u0005C\u0000\u0000\u0201\u020a"+
		"\u0005,\u0000\u0000\u0202\u0207\u0003 \u0010\u0000\u0203\u0204\u0005\u0004"+
		"\u0000\u0000\u0204\u0206\u0003 \u0010\u0000\u0205\u0203\u0001\u0000\u0000"+
		"\u0000\u0206\u0209\u0001\u0000\u0000\u0000\u0207\u0205\u0001\u0000\u0000"+
		"\u0000\u0207\u0208\u0001\u0000\u0000\u0000\u0208\u020b\u0001\u0000\u0000"+
		"\u0000\u0209\u0207\u0001\u0000\u0000\u0000\u020a\u0202\u0001\u0000\u0000"+
		"\u0000\u020a\u020b\u0001\u0000\u0000\u0000\u020b\u020c\u0001\u0000\u0000"+
		"\u0000\u020c\u020e\u0005\u0005\u0000\u0000\u020d\u020f\u0005\u0006\u0000"+
		"\u0000\u020e\u020d\u0001\u0000\u0000\u0000\u020e\u020f\u0001\u0000\u0000"+
		"\u0000\u020f1\u0001\u0000\u0000\u0000\u0210\u0211\u0005C\u0000\u0000\u0211"+
		"\u0212\u0005\u001c\u0000\u0000\u0212\u0213\u0005C\u0000\u0000\u0213\u021c"+
		"\u0005,\u0000\u0000\u0214\u0219\u0003 \u0010\u0000\u0215\u0216\u0005\u0004"+
		"\u0000\u0000\u0216\u0218\u0003 \u0010\u0000\u0217\u0215\u0001\u0000\u0000"+
		"\u0000\u0218\u021b\u0001\u0000\u0000\u0000\u0219\u0217\u0001\u0000\u0000"+
		"\u0000\u0219\u021a\u0001\u0000\u0000\u0000\u021a\u021d\u0001\u0000\u0000"+
		"\u0000\u021b\u0219\u0001\u0000\u0000\u0000\u021c\u0214\u0001\u0000\u0000"+
		"\u0000\u021c\u021d\u0001\u0000\u0000\u0000\u021d\u021e\u0001\u0000\u0000"+
		"\u0000\u021e\u0220\u0005\u0005\u0000\u0000\u021f\u0221\u0005\u0006\u0000"+
		"\u0000\u0220\u021f\u0001\u0000\u0000\u0000\u0220\u0221\u0001\u0000\u0000"+
		"\u0000\u02213\u0001\u0000\u0000\u0000\u0222\u0223\u0003$\u0012\u0000\u0223"+
		"5\u0001\u0000\u0000\u0000\u0224\u0226\u0005<\u0000\u0000\u0225\u0227\u0003"+
		" \u0010\u0000\u0226\u0225\u0001\u0000\u0000\u0000\u0226\u0227\u0001\u0000"+
		"\u0000\u0000\u0227\u0229\u0001\u0000\u0000\u0000\u0228\u022a\u0005\u0006"+
		"\u0000\u0000\u0229\u0228\u0001\u0000\u0000\u0000\u0229\u022a\u0001\u0000"+
		"\u0000\u0000\u022a7\u0001\u0000\u0000\u0000>;EV^aemptv}\u0085\u0093\u0097"+
		"\u009d\u00a0\u00a6\u00be\u00ca\u00d0\u00d7\u00e3\u00e7\u00f2\u00fd\u0101"+
		"\u0104\u0107\u010a\u010e\u0114\u0118\u0129\u0138\u013c\u0140\u0151\u015b"+
		"\u015f\u0178\u018b\u01a6\u01a9\u01ad\u01be\u01c0\u01cb\u01cf\u01db\u01de"+
		"\u01e2\u01ee\u01f2\u01fd\u0207\u020a\u020e\u0219\u021c\u0220\u0226\u0229";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}