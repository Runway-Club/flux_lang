// Generated from /Users/ambrose/Repos/flux_lang/parser/Math.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Math extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TEXT=1, TEXT_TYPE=2, BOOLEAN=3, BOOLEAN_TYPE=4, NUMBER=5, NUMBER_TYPE=6, 
		NULL=7, DIGIT=8, OCTET=9, IPV4=10, IPV4_TYPE=11, OP_ADD=12, OP_SUB=13, 
		OP_MUL=14, OP_DIV=15, OP_MOD=16, OP_POW=17, OP_EQ=18, OP_NEQ=19, OP_GT=20, 
		OP_LT=21, OP_GTE=22, OP_LTE=23, OP_AND=24, OP_OR=25, OP_NOT=26, OP_XOR=27, 
		L_BLOCK=28, R_BLOCK=29, L_PAREN=30, R_PAREN=31, L_SQUARE=32, R_SQUARE=33, 
		COMMA=34, DOT=35, VAR_IDENTIFIER=36, COMMON_IDENTIFIER=37, AT=38, NEWLINE=39, 
		WS=40;
	public static final int
		RULE_op_level1 = 0, RULE_op_level2 = 1, RULE_op_level3 = 2, RULE_op_level4 = 3, 
		RULE_op_level5 = 4, RULE_numeric_expression = 5, RULE_logical_expression = 6, 
		RULE_comparative_expression = 7, RULE_math_expression = 8, RULE_get_array_element = 9, 
		RULE_get_child = 10, RULE_get_var = 11, RULE_function_call = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"op_level1", "op_level2", "op_level3", "op_level4", "op_level5", "numeric_expression", 
			"logical_expression", "comparative_expression", "math_expression", "get_array_element", 
			"get_child", "get_var", "function_call"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'text'", null, "'bool'", null, "'num'", "'na'", null, null, 
			null, "'ipv4'", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'='", "'!='", 
			"'>'", "'<'", "'>='", "'<='", "'and'", "'or'", "'not'", "'xor'", "'{'", 
			"'}'", "'('", "')'", "'['", "']'", "','", "'.'", null, null, "'@'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE", 
			"NULL", "DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "OP_ADD", "OP_SUB", "OP_MUL", 
			"OP_DIV", "OP_MOD", "OP_POW", "OP_EQ", "OP_NEQ", "OP_GT", "OP_LT", "OP_GTE", 
			"OP_LTE", "OP_AND", "OP_OR", "OP_NOT", "OP_XOR", "L_BLOCK", "R_BLOCK", 
			"L_PAREN", "R_PAREN", "L_SQUARE", "R_SQUARE", "COMMA", "DOT", "VAR_IDENTIFIER", 
			"COMMON_IDENTIFIER", "AT", "NEWLINE", "WS"
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
	public String getGrammarFileName() { return "Math.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Math(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op_level1Context extends ParserRuleContext {
		public TerminalNode OP_MUL() { return getToken(Math.OP_MUL, 0); }
		public TerminalNode OP_DIV() { return getToken(Math.OP_DIV, 0); }
		public TerminalNode OP_MOD() { return getToken(Math.OP_MOD, 0); }
		public TerminalNode OP_POW() { return getToken(Math.OP_POW, 0); }
		public Op_level1Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level1; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level1(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level1(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level1(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level1Context op_level1() throws RecognitionException {
		Op_level1Context _localctx = new Op_level1Context(_ctx, getState());
		enterRule(_localctx, 0, RULE_op_level1);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(26);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 245760L) != 0)) ) {
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
	public static class Op_level2Context extends ParserRuleContext {
		public TerminalNode OP_ADD() { return getToken(Math.OP_ADD, 0); }
		public TerminalNode OP_SUB() { return getToken(Math.OP_SUB, 0); }
		public Op_level2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level2; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level2(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level2(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level2(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level2Context op_level2() throws RecognitionException {
		Op_level2Context _localctx = new Op_level2Context(_ctx, getState());
		enterRule(_localctx, 2, RULE_op_level2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(28);
			_la = _input.LA(1);
			if ( !(_la==OP_ADD || _la==OP_SUB) ) {
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
	public static class Op_level3Context extends ParserRuleContext {
		public TerminalNode OP_AND() { return getToken(Math.OP_AND, 0); }
		public TerminalNode OP_OR() { return getToken(Math.OP_OR, 0); }
		public TerminalNode OP_XOR() { return getToken(Math.OP_XOR, 0); }
		public Op_level3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level3; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level3(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level3(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level3(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level3Context op_level3() throws RecognitionException {
		Op_level3Context _localctx = new Op_level3Context(_ctx, getState());
		enterRule(_localctx, 4, RULE_op_level3);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 184549376L) != 0)) ) {
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
	public static class Op_level4Context extends ParserRuleContext {
		public TerminalNode OP_EQ() { return getToken(Math.OP_EQ, 0); }
		public TerminalNode OP_NEQ() { return getToken(Math.OP_NEQ, 0); }
		public TerminalNode OP_LT() { return getToken(Math.OP_LT, 0); }
		public TerminalNode OP_GT() { return getToken(Math.OP_GT, 0); }
		public TerminalNode OP_GTE() { return getToken(Math.OP_GTE, 0); }
		public TerminalNode OP_LTE() { return getToken(Math.OP_LTE, 0); }
		public Op_level4Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level4; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level4(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level4(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level4(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level4Context op_level4() throws RecognitionException {
		Op_level4Context _localctx = new Op_level4Context(_ctx, getState());
		enterRule(_localctx, 6, RULE_op_level4);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 16515072L) != 0)) ) {
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
	public static class Op_level5Context extends ParserRuleContext {
		public TerminalNode OP_NOT() { return getToken(Math.OP_NOT, 0); }
		public Op_level5Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_level5; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterOp_level5(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitOp_level5(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitOp_level5(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Op_level5Context op_level5() throws RecognitionException {
		Op_level5Context _localctx = new Op_level5Context(_ctx, getState());
		enterRule(_localctx, 8, RULE_op_level5);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(34);
			match(OP_NOT);
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
	public static class Numeric_expressionContext extends ParserRuleContext {
		public TerminalNode L_PAREN() { return getToken(Math.L_PAREN, 0); }
		public List<Numeric_expressionContext> numeric_expression() {
			return getRuleContexts(Numeric_expressionContext.class);
		}
		public Numeric_expressionContext numeric_expression(int i) {
			return getRuleContext(Numeric_expressionContext.class,i);
		}
		public TerminalNode R_PAREN() { return getToken(Math.R_PAREN, 0); }
		public TerminalNode NUMBER() { return getToken(Math.NUMBER, 0); }
		public Get_varContext get_var() {
			return getRuleContext(Get_varContext.class,0);
		}
		public Function_callContext function_call() {
			return getRuleContext(Function_callContext.class,0);
		}
		public TerminalNode OP_SUB() { return getToken(Math.OP_SUB, 0); }
		public Op_level1Context op_level1() {
			return getRuleContext(Op_level1Context.class,0);
		}
		public Op_level2Context op_level2() {
			return getRuleContext(Op_level2Context.class,0);
		}
		public Numeric_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numeric_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterNumeric_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitNumeric_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitNumeric_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Numeric_expressionContext numeric_expression() throws RecognitionException {
		return numeric_expression(0);
	}

	private Numeric_expressionContext numeric_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Numeric_expressionContext _localctx = new Numeric_expressionContext(_ctx, _parentState);
		Numeric_expressionContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_numeric_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				setState(37);
				match(L_PAREN);
				setState(38);
				numeric_expression(0);
				setState(39);
				match(R_PAREN);
				}
				break;
			case 2:
				{
				setState(41);
				match(NUMBER);
				}
				break;
			case 3:
				{
				setState(42);
				get_var();
				}
				break;
			case 4:
				{
				setState(43);
				function_call();
				}
				break;
			case 5:
				{
				setState(44);
				match(OP_SUB);
				setState(45);
				numeric_expression(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(58);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(56);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new Numeric_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_numeric_expression);
						setState(48);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(49);
						op_level1();
						setState(50);
						numeric_expression(7);
						}
						break;
					case 2:
						{
						_localctx = new Numeric_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_numeric_expression);
						setState(52);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(53);
						op_level2();
						setState(54);
						numeric_expression(6);
						}
						break;
					}
					} 
				}
				setState(60);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
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
	public static class Logical_expressionContext extends ParserRuleContext {
		public TerminalNode L_PAREN() { return getToken(Math.L_PAREN, 0); }
		public List<Logical_expressionContext> logical_expression() {
			return getRuleContexts(Logical_expressionContext.class);
		}
		public Logical_expressionContext logical_expression(int i) {
			return getRuleContext(Logical_expressionContext.class,i);
		}
		public TerminalNode R_PAREN() { return getToken(Math.R_PAREN, 0); }
		public TerminalNode OP_NOT() { return getToken(Math.OP_NOT, 0); }
		public TerminalNode BOOLEAN() { return getToken(Math.BOOLEAN, 0); }
		public Get_varContext get_var() {
			return getRuleContext(Get_varContext.class,0);
		}
		public Function_callContext function_call() {
			return getRuleContext(Function_callContext.class,0);
		}
		public Op_level3Context op_level3() {
			return getRuleContext(Op_level3Context.class,0);
		}
		public Op_level4Context op_level4() {
			return getRuleContext(Op_level4Context.class,0);
		}
		public Logical_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logical_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterLogical_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitLogical_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitLogical_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Logical_expressionContext logical_expression() throws RecognitionException {
		return logical_expression(0);
	}

	private Logical_expressionContext logical_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Logical_expressionContext _localctx = new Logical_expressionContext(_ctx, _parentState);
		Logical_expressionContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_logical_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(62);
				match(L_PAREN);
				setState(63);
				logical_expression(0);
				setState(64);
				match(R_PAREN);
				}
				break;
			case 2:
				{
				setState(66);
				match(OP_NOT);
				setState(67);
				logical_expression(4);
				}
				break;
			case 3:
				{
				setState(68);
				match(BOOLEAN);
				}
				break;
			case 4:
				{
				setState(69);
				get_var();
				}
				break;
			case 5:
				{
				setState(70);
				function_call();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(83);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(81);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new Logical_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_logical_expression);
						setState(73);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(74);
						op_level3();
						setState(75);
						logical_expression(7);
						}
						break;
					case 2:
						{
						_localctx = new Logical_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_logical_expression);
						setState(77);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(78);
						op_level4();
						setState(79);
						logical_expression(6);
						}
						break;
					}
					} 
				}
				setState(85);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
	public static class Comparative_expressionContext extends ParserRuleContext {
		public List<Numeric_expressionContext> numeric_expression() {
			return getRuleContexts(Numeric_expressionContext.class);
		}
		public Numeric_expressionContext numeric_expression(int i) {
			return getRuleContext(Numeric_expressionContext.class,i);
		}
		public Op_level4Context op_level4() {
			return getRuleContext(Op_level4Context.class,0);
		}
		public List<Logical_expressionContext> logical_expression() {
			return getRuleContexts(Logical_expressionContext.class);
		}
		public Logical_expressionContext logical_expression(int i) {
			return getRuleContext(Logical_expressionContext.class,i);
		}
		public Op_level3Context op_level3() {
			return getRuleContext(Op_level3Context.class,0);
		}
		public Op_level5Context op_level5() {
			return getRuleContext(Op_level5Context.class,0);
		}
		public TerminalNode L_PAREN() { return getToken(Math.L_PAREN, 0); }
		public Comparative_expressionContext comparative_expression() {
			return getRuleContext(Comparative_expressionContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(Math.R_PAREN, 0); }
		public Comparative_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparative_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterComparative_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitComparative_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitComparative_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Comparative_expressionContext comparative_expression() throws RecognitionException {
		return comparative_expression(0);
	}

	private Comparative_expressionContext comparative_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Comparative_expressionContext _localctx = new Comparative_expressionContext(_ctx, _parentState);
		Comparative_expressionContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_comparative_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(87);
				numeric_expression(0);
				setState(88);
				op_level4();
				setState(89);
				numeric_expression(0);
				}
				break;
			case 2:
				{
				setState(91);
				logical_expression(0);
				setState(92);
				op_level3();
				setState(93);
				logical_expression(0);
				}
				break;
			case 3:
				{
				setState(95);
				op_level5();
				setState(96);
				match(L_PAREN);
				setState(97);
				comparative_expression(0);
				setState(98);
				match(R_PAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(108);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Comparative_expressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_comparative_expression);
					setState(102);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(103);
					op_level3();
					setState(104);
					logical_expression(0);
					}
					} 
				}
				setState(110);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
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
	public static class Math_expressionContext extends ParserRuleContext {
		public Get_varContext get_var() {
			return getRuleContext(Get_varContext.class,0);
		}
		public Numeric_expressionContext numeric_expression() {
			return getRuleContext(Numeric_expressionContext.class,0);
		}
		public Logical_expressionContext logical_expression() {
			return getRuleContext(Logical_expressionContext.class,0);
		}
		public Math_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_math_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterMath_expression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitMath_expression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitMath_expression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Math_expressionContext math_expression() throws RecognitionException {
		Math_expressionContext _localctx = new Math_expressionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_math_expression);
		try {
			setState(114);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				get_var();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				numeric_expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(113);
				logical_expression(0);
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
	public static class Get_array_elementContext extends ParserRuleContext {
		public TerminalNode VAR_IDENTIFIER() { return getToken(Math.VAR_IDENTIFIER, 0); }
		public TerminalNode L_SQUARE() { return getToken(Math.L_SQUARE, 0); }
		public Numeric_expressionContext numeric_expression() {
			return getRuleContext(Numeric_expressionContext.class,0);
		}
		public TerminalNode R_SQUARE() { return getToken(Math.R_SQUARE, 0); }
		public Get_array_elementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_get_array_element; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterGet_array_element(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitGet_array_element(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitGet_array_element(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Get_array_elementContext get_array_element() throws RecognitionException {
		Get_array_elementContext _localctx = new Get_array_elementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_get_array_element);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			match(VAR_IDENTIFIER);
			setState(117);
			match(L_SQUARE);
			setState(118);
			numeric_expression(0);
			setState(119);
			match(R_SQUARE);
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
	public static class Get_childContext extends ParserRuleContext {
		public List<TerminalNode> VAR_IDENTIFIER() { return getTokens(Math.VAR_IDENTIFIER); }
		public TerminalNode VAR_IDENTIFIER(int i) {
			return getToken(Math.VAR_IDENTIFIER, i);
		}
		public TerminalNode DOT() { return getToken(Math.DOT, 0); }
		public Get_array_elementContext get_array_element() {
			return getRuleContext(Get_array_elementContext.class,0);
		}
		public Get_childContext get_child() {
			return getRuleContext(Get_childContext.class,0);
		}
		public Get_childContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_get_child; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterGet_child(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitGet_child(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitGet_child(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Get_childContext get_child() throws RecognitionException {
		Get_childContext _localctx = new Get_childContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_get_child);
		try {
			setState(136);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(121);
				match(VAR_IDENTIFIER);
				setState(122);
				match(DOT);
				setState(123);
				match(VAR_IDENTIFIER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				match(VAR_IDENTIFIER);
				setState(125);
				match(DOT);
				setState(126);
				get_array_element();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(127);
				match(VAR_IDENTIFIER);
				setState(128);
				match(DOT);
				setState(129);
				get_child();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(130);
				get_array_element();
				setState(131);
				match(DOT);
				setState(132);
				get_child();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(134);
				get_array_element();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(135);
				match(VAR_IDENTIFIER);
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
	public static class Get_varContext extends ParserRuleContext {
		public TerminalNode VAR_IDENTIFIER() { return getToken(Math.VAR_IDENTIFIER, 0); }
		public Get_array_elementContext get_array_element() {
			return getRuleContext(Get_array_elementContext.class,0);
		}
		public Get_childContext get_child() {
			return getRuleContext(Get_childContext.class,0);
		}
		public Get_varContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_get_var; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterGet_var(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitGet_var(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitGet_var(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Get_varContext get_var() throws RecognitionException {
		Get_varContext _localctx = new Get_varContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_get_var);
		try {
			setState(141);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(138);
				match(VAR_IDENTIFIER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(139);
				get_array_element();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(140);
				get_child();
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
	public static class Function_callContext extends ParserRuleContext {
		public TerminalNode VAR_IDENTIFIER() { return getToken(Math.VAR_IDENTIFIER, 0); }
		public TerminalNode L_PAREN() { return getToken(Math.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(Math.R_PAREN, 0); }
		public List<Math_expressionContext> math_expression() {
			return getRuleContexts(Math_expressionContext.class);
		}
		public Math_expressionContext math_expression(int i) {
			return getRuleContext(Math_expressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(Math.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(Math.COMMA, i);
		}
		public Function_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterFunction_call(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitFunction_call(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitFunction_call(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Function_callContext function_call() throws RecognitionException {
		Function_callContext _localctx = new Function_callContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_function_call);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(VAR_IDENTIFIER);
			setState(144);
			match(L_PAREN);
			setState(153);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 69860335656L) != 0)) {
				{
				setState(145);
				math_expression();
				setState(150);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(146);
					match(COMMA);
					setState(147);
					math_expression();
					}
					}
					setState(152);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(155);
			match(R_PAREN);
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
		case 5:
			return numeric_expression_sempred((Numeric_expressionContext)_localctx, predIndex);
		case 6:
			return logical_expression_sempred((Logical_expressionContext)_localctx, predIndex);
		case 7:
			return comparative_expression_sempred((Comparative_expressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean numeric_expression_sempred(Numeric_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		}
		return true;
	}
	private boolean logical_expression_sempred(Logical_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 6);
		case 3:
			return precpred(_ctx, 5);
		}
		return true;
	}
	private boolean comparative_expression_sempred(Comparative_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001(\u009e\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005/\b\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0005\u00059\b\u0005\n\u0005\f\u0005<\t\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006H\b\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0005\u0006R\b\u0006\n\u0006\f\u0006U\t\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0003\u0007e\b\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0005\u0007k\b\u0007\n\u0007\f\u0007n\t\u0007\u0001\b\u0001"+
		"\b\u0001\b\u0003\bs\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u0089\b\n\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0003\u000b\u008e\b\u000b\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0005\f\u0095\b\f\n\f\f\f\u0098\t\f\u0003\f\u009a\b"+
		"\f\u0001\f\u0001\f\u0001\f\u0000\u0003\n\f\u000e\r\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u0000\u0004\u0001\u0000\u000e"+
		"\u0011\u0001\u0000\f\r\u0002\u0000\u0018\u0019\u001b\u001b\u0001\u0000"+
		"\u0012\u0017\u00aa\u0000\u001a\u0001\u0000\u0000\u0000\u0002\u001c\u0001"+
		"\u0000\u0000\u0000\u0004\u001e\u0001\u0000\u0000\u0000\u0006 \u0001\u0000"+
		"\u0000\u0000\b\"\u0001\u0000\u0000\u0000\n.\u0001\u0000\u0000\u0000\f"+
		"G\u0001\u0000\u0000\u0000\u000ed\u0001\u0000\u0000\u0000\u0010r\u0001"+
		"\u0000\u0000\u0000\u0012t\u0001\u0000\u0000\u0000\u0014\u0088\u0001\u0000"+
		"\u0000\u0000\u0016\u008d\u0001\u0000\u0000\u0000\u0018\u008f\u0001\u0000"+
		"\u0000\u0000\u001a\u001b\u0007\u0000\u0000\u0000\u001b\u0001\u0001\u0000"+
		"\u0000\u0000\u001c\u001d\u0007\u0001\u0000\u0000\u001d\u0003\u0001\u0000"+
		"\u0000\u0000\u001e\u001f\u0007\u0002\u0000\u0000\u001f\u0005\u0001\u0000"+
		"\u0000\u0000 !\u0007\u0003\u0000\u0000!\u0007\u0001\u0000\u0000\u0000"+
		"\"#\u0005\u001a\u0000\u0000#\t\u0001\u0000\u0000\u0000$%\u0006\u0005\uffff"+
		"\uffff\u0000%&\u0005\u001e\u0000\u0000&\'\u0003\n\u0005\u0000\'(\u0005"+
		"\u001f\u0000\u0000(/\u0001\u0000\u0000\u0000)/\u0005\u0005\u0000\u0000"+
		"*/\u0003\u0016\u000b\u0000+/\u0003\u0018\f\u0000,-\u0005\r\u0000\u0000"+
		"-/\u0003\n\u0005\u0001.$\u0001\u0000\u0000\u0000.)\u0001\u0000\u0000\u0000"+
		".*\u0001\u0000\u0000\u0000.+\u0001\u0000\u0000\u0000.,\u0001\u0000\u0000"+
		"\u0000/:\u0001\u0000\u0000\u000001\n\u0006\u0000\u000012\u0003\u0000\u0000"+
		"\u000023\u0003\n\u0005\u000739\u0001\u0000\u0000\u000045\n\u0005\u0000"+
		"\u000056\u0003\u0002\u0001\u000067\u0003\n\u0005\u000679\u0001\u0000\u0000"+
		"\u000080\u0001\u0000\u0000\u000084\u0001\u0000\u0000\u00009<\u0001\u0000"+
		"\u0000\u0000:8\u0001\u0000\u0000\u0000:;\u0001\u0000\u0000\u0000;\u000b"+
		"\u0001\u0000\u0000\u0000<:\u0001\u0000\u0000\u0000=>\u0006\u0006\uffff"+
		"\uffff\u0000>?\u0005\u001e\u0000\u0000?@\u0003\f\u0006\u0000@A\u0005\u001f"+
		"\u0000\u0000AH\u0001\u0000\u0000\u0000BC\u0005\u001a\u0000\u0000CH\u0003"+
		"\f\u0006\u0004DH\u0005\u0003\u0000\u0000EH\u0003\u0016\u000b\u0000FH\u0003"+
		"\u0018\f\u0000G=\u0001\u0000\u0000\u0000GB\u0001\u0000\u0000\u0000GD\u0001"+
		"\u0000\u0000\u0000GE\u0001\u0000\u0000\u0000GF\u0001\u0000\u0000\u0000"+
		"HS\u0001\u0000\u0000\u0000IJ\n\u0006\u0000\u0000JK\u0003\u0004\u0002\u0000"+
		"KL\u0003\f\u0006\u0007LR\u0001\u0000\u0000\u0000MN\n\u0005\u0000\u0000"+
		"NO\u0003\u0006\u0003\u0000OP\u0003\f\u0006\u0006PR\u0001\u0000\u0000\u0000"+
		"QI\u0001\u0000\u0000\u0000QM\u0001\u0000\u0000\u0000RU\u0001\u0000\u0000"+
		"\u0000SQ\u0001\u0000\u0000\u0000ST\u0001\u0000\u0000\u0000T\r\u0001\u0000"+
		"\u0000\u0000US\u0001\u0000\u0000\u0000VW\u0006\u0007\uffff\uffff\u0000"+
		"WX\u0003\n\u0005\u0000XY\u0003\u0006\u0003\u0000YZ\u0003\n\u0005\u0000"+
		"Ze\u0001\u0000\u0000\u0000[\\\u0003\f\u0006\u0000\\]\u0003\u0004\u0002"+
		"\u0000]^\u0003\f\u0006\u0000^e\u0001\u0000\u0000\u0000_`\u0003\b\u0004"+
		"\u0000`a\u0005\u001e\u0000\u0000ab\u0003\u000e\u0007\u0000bc\u0005\u001f"+
		"\u0000\u0000ce\u0001\u0000\u0000\u0000dV\u0001\u0000\u0000\u0000d[\u0001"+
		"\u0000\u0000\u0000d_\u0001\u0000\u0000\u0000el\u0001\u0000\u0000\u0000"+
		"fg\n\u0002\u0000\u0000gh\u0003\u0004\u0002\u0000hi\u0003\f\u0006\u0000"+
		"ik\u0001\u0000\u0000\u0000jf\u0001\u0000\u0000\u0000kn\u0001\u0000\u0000"+
		"\u0000lj\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000\u0000m\u000f\u0001"+
		"\u0000\u0000\u0000nl\u0001\u0000\u0000\u0000os\u0003\u0016\u000b\u0000"+
		"ps\u0003\n\u0005\u0000qs\u0003\f\u0006\u0000ro\u0001\u0000\u0000\u0000"+
		"rp\u0001\u0000\u0000\u0000rq\u0001\u0000\u0000\u0000s\u0011\u0001\u0000"+
		"\u0000\u0000tu\u0005$\u0000\u0000uv\u0005 \u0000\u0000vw\u0003\n\u0005"+
		"\u0000wx\u0005!\u0000\u0000x\u0013\u0001\u0000\u0000\u0000yz\u0005$\u0000"+
		"\u0000z{\u0005#\u0000\u0000{\u0089\u0005$\u0000\u0000|}\u0005$\u0000\u0000"+
		"}~\u0005#\u0000\u0000~\u0089\u0003\u0012\t\u0000\u007f\u0080\u0005$\u0000"+
		"\u0000\u0080\u0081\u0005#\u0000\u0000\u0081\u0089\u0003\u0014\n\u0000"+
		"\u0082\u0083\u0003\u0012\t\u0000\u0083\u0084\u0005#\u0000\u0000\u0084"+
		"\u0085\u0003\u0014\n\u0000\u0085\u0089\u0001\u0000\u0000\u0000\u0086\u0089"+
		"\u0003\u0012\t\u0000\u0087\u0089\u0005$\u0000\u0000\u0088y\u0001\u0000"+
		"\u0000\u0000\u0088|\u0001\u0000\u0000\u0000\u0088\u007f\u0001\u0000\u0000"+
		"\u0000\u0088\u0082\u0001\u0000\u0000\u0000\u0088\u0086\u0001\u0000\u0000"+
		"\u0000\u0088\u0087\u0001\u0000\u0000\u0000\u0089\u0015\u0001\u0000\u0000"+
		"\u0000\u008a\u008e\u0005$\u0000\u0000\u008b\u008e\u0003\u0012\t\u0000"+
		"\u008c\u008e\u0003\u0014\n\u0000\u008d\u008a\u0001\u0000\u0000\u0000\u008d"+
		"\u008b\u0001\u0000\u0000\u0000\u008d\u008c\u0001\u0000\u0000\u0000\u008e"+
		"\u0017\u0001\u0000\u0000\u0000\u008f\u0090\u0005$\u0000\u0000\u0090\u0099"+
		"\u0005\u001e\u0000\u0000\u0091\u0096\u0003\u0010\b\u0000\u0092\u0093\u0005"+
		"\"\u0000\u0000\u0093\u0095\u0003\u0010\b\u0000\u0094\u0092\u0001\u0000"+
		"\u0000\u0000\u0095\u0098\u0001\u0000\u0000\u0000\u0096\u0094\u0001\u0000"+
		"\u0000\u0000\u0096\u0097\u0001\u0000\u0000\u0000\u0097\u009a\u0001\u0000"+
		"\u0000\u0000\u0098\u0096\u0001\u0000\u0000\u0000\u0099\u0091\u0001\u0000"+
		"\u0000\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a\u009b\u0001\u0000"+
		"\u0000\u0000\u009b\u009c\u0005\u001f\u0000\u0000\u009c\u0019\u0001\u0000"+
		"\u0000\u0000\r.8:GQSdlr\u0088\u008d\u0096\u0099";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}