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
		NULL=7, DIGIT=8, OCTET=9, IPV4=10, IPV4_TYPE=11, LOOP=12, IF=13, ELSE=14, 
		FUNC=15, RETURN=16, RETURN_TYPE=17, OP_ADD=18, OP_SUB=19, OP_MUL=20, OP_DIV=21, 
		OP_MOD=22, OP_POW=23, OP_EQ=24, OP_NEQ=25, OP_GT=26, OP_LT=27, OP_GTE=28, 
		OP_LTE=29, OP_AND=30, OP_OR=31, OP_NOT=32, OP_XOR=33, L_BLOCK=34, R_BLOCK=35, 
		L_PAREN=36, R_PAREN=37, L_SQUARE=38, R_SQUARE=39, COMMA=40, DOT=41, VAR_IDENTIFIER=42, 
		COMMON_IDENTIFIER=43, AT=44, NEWLINE=45, WS=46;
	public static final int
		RULE_op_level1 = 0, RULE_op_level2 = 1, RULE_op_level3 = 2, RULE_op_level4 = 3, 
		RULE_op_level5 = 4, RULE_numeric_expression = 5, RULE_logical_expression = 6, 
		RULE_comparative_expression = 7, RULE_math_expression = 8, RULE_get_array_element = 9, 
		RULE_get_child = 10, RULE_get_var = 11, RULE_args = 12, RULE_function_call = 13;
	private static String[] makeRuleNames() {
		return new String[] {
			"op_level1", "op_level2", "op_level3", "op_level4", "op_level5", "numeric_expression", 
			"logical_expression", "comparative_expression", "math_expression", "get_array_element", 
			"get_child", "get_var", "args", "function_call"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'text'", null, "'bool'", null, "'num'", "'na'", null, null, 
			null, "'ipv4'", "'loop'", "'if'", "'else'", "'fun'", "'return'", "'->'", 
			"'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'='", "'!='", "'>'", "'<'", 
			"'>='", "'<='", "'and'", "'or'", "'not'", "'xor'", "'{'", "'}'", "'('", 
			"')'", "'['", "']'", "','", "'.'", null, null, "'@'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "NUMBER", "NUMBER_TYPE", 
			"NULL", "DIGIT", "OCTET", "IPV4", "IPV4_TYPE", "LOOP", "IF", "ELSE", 
			"FUNC", "RETURN", "RETURN_TYPE", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", 
			"OP_MOD", "OP_POW", "OP_EQ", "OP_NEQ", "OP_GT", "OP_LT", "OP_GTE", "OP_LTE", 
			"OP_AND", "OP_OR", "OP_NOT", "OP_XOR", "L_BLOCK", "R_BLOCK", "L_PAREN", 
			"R_PAREN", "L_SQUARE", "R_SQUARE", "COMMA", "DOT", "VAR_IDENTIFIER", 
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
			setState(28);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 15728640L) != 0)) ) {
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
			setState(30);
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
			setState(32);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 11811160064L) != 0)) ) {
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
			setState(34);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1056964608L) != 0)) ) {
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
			setState(36);
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
			setState(48);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				setState(39);
				match(L_PAREN);
				setState(40);
				numeric_expression(0);
				setState(41);
				match(R_PAREN);
				}
				break;
			case 2:
				{
				setState(43);
				match(NUMBER);
				}
				break;
			case 3:
				{
				setState(44);
				get_var();
				}
				break;
			case 4:
				{
				setState(45);
				function_call();
				}
				break;
			case 5:
				{
				setState(46);
				match(OP_SUB);
				setState(47);
				numeric_expression(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(60);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(58);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new Numeric_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_numeric_expression);
						setState(50);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(51);
						op_level1();
						setState(52);
						numeric_expression(7);
						}
						break;
					case 2:
						{
						_localctx = new Numeric_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_numeric_expression);
						setState(54);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(55);
						op_level2();
						setState(56);
						numeric_expression(6);
						}
						break;
					}
					} 
				}
				setState(62);
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
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(64);
				match(L_PAREN);
				setState(65);
				logical_expression(0);
				setState(66);
				match(R_PAREN);
				}
				break;
			case 2:
				{
				setState(68);
				match(OP_NOT);
				setState(69);
				logical_expression(4);
				}
				break;
			case 3:
				{
				setState(70);
				match(BOOLEAN);
				}
				break;
			case 4:
				{
				setState(71);
				get_var();
				}
				break;
			case 5:
				{
				setState(72);
				function_call();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(85);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(83);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new Logical_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_logical_expression);
						setState(75);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(76);
						op_level3();
						setState(77);
						logical_expression(7);
						}
						break;
					case 2:
						{
						_localctx = new Logical_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_logical_expression);
						setState(79);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(80);
						op_level4();
						setState(81);
						logical_expression(6);
						}
						break;
					}
					} 
				}
				setState(87);
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
			setState(102);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(89);
				numeric_expression(0);
				setState(90);
				op_level4();
				setState(91);
				numeric_expression(0);
				}
				break;
			case 2:
				{
				setState(93);
				logical_expression(0);
				setState(94);
				op_level3();
				setState(95);
				logical_expression(0);
				}
				break;
			case 3:
				{
				setState(97);
				op_level5();
				setState(98);
				match(L_PAREN);
				setState(99);
				comparative_expression(0);
				setState(100);
				match(R_PAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(110);
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
					setState(104);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(105);
					op_level3();
					setState(106);
					logical_expression(0);
					}
					} 
				}
				setState(112);
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
			setState(116);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(113);
				get_var();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(114);
				numeric_expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(115);
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
			setState(118);
			match(VAR_IDENTIFIER);
			setState(119);
			match(L_SQUARE);
			setState(120);
			numeric_expression(0);
			setState(121);
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
			setState(138);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(123);
				match(VAR_IDENTIFIER);
				setState(124);
				match(DOT);
				setState(125);
				match(VAR_IDENTIFIER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
				match(VAR_IDENTIFIER);
				setState(127);
				match(DOT);
				setState(128);
				get_array_element();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(129);
				match(VAR_IDENTIFIER);
				setState(130);
				match(DOT);
				setState(131);
				get_child();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(132);
				get_array_element();
				setState(133);
				match(DOT);
				setState(134);
				get_child();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(136);
				get_array_element();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(137);
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
			setState(143);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(140);
				match(VAR_IDENTIFIER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(141);
				get_array_element();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(142);
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
	public static class ArgsContext extends ParserRuleContext {
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
		public ArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_args; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).enterArgs(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof MathListener ) ((MathListener)listener).exitArgs(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof MathVisitor ) return ((MathVisitor<? extends T>)visitor).visitArgs(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArgsContext args() throws RecognitionException {
		ArgsContext _localctx = new ArgsContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_args);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
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
		public ArgsContext args() {
			return getRuleContext(ArgsContext.class,0);
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
		enterRule(_localctx, 26, RULE_function_call);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			match(VAR_IDENTIFIER);
			setState(154);
			match(L_PAREN);
			setState(156);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4471061479464L) != 0)) {
				{
				setState(155);
				args();
				}
			}

			setState(158);
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
		"\u0004\u0001.\u00a1\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u00051\b\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0005\u0005;\b\u0005\n\u0005\f\u0005>\t\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006J\b\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0005\u0006T\b\u0006\n\u0006\f\u0006W\t\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0003\u0007g\b\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0005\u0007m\b\u0007\n\u0007\f\u0007p\t\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0003\bu\b\b\u0001\t\u0001\t\u0001\t\u0001\t"+
		"\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u008b"+
		"\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u0090\b\u000b\u0001"+
		"\f\u0001\f\u0001\f\u0005\f\u0095\b\f\n\f\f\f\u0098\t\f\u0001\r\u0001\r"+
		"\u0001\r\u0003\r\u009d\b\r\u0001\r\u0001\r\u0001\r\u0000\u0003\n\f\u000e"+
		"\u000e\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018"+
		"\u001a\u0000\u0004\u0001\u0000\u0014\u0017\u0001\u0000\u0012\u0013\u0002"+
		"\u0000\u001e\u001f!!\u0001\u0000\u0018\u001d\u00ac\u0000\u001c\u0001\u0000"+
		"\u0000\u0000\u0002\u001e\u0001\u0000\u0000\u0000\u0004 \u0001\u0000\u0000"+
		"\u0000\u0006\"\u0001\u0000\u0000\u0000\b$\u0001\u0000\u0000\u0000\n0\u0001"+
		"\u0000\u0000\u0000\fI\u0001\u0000\u0000\u0000\u000ef\u0001\u0000\u0000"+
		"\u0000\u0010t\u0001\u0000\u0000\u0000\u0012v\u0001\u0000\u0000\u0000\u0014"+
		"\u008a\u0001\u0000\u0000\u0000\u0016\u008f\u0001\u0000\u0000\u0000\u0018"+
		"\u0091\u0001\u0000\u0000\u0000\u001a\u0099\u0001\u0000\u0000\u0000\u001c"+
		"\u001d\u0007\u0000\u0000\u0000\u001d\u0001\u0001\u0000\u0000\u0000\u001e"+
		"\u001f\u0007\u0001\u0000\u0000\u001f\u0003\u0001\u0000\u0000\u0000 !\u0007"+
		"\u0002\u0000\u0000!\u0005\u0001\u0000\u0000\u0000\"#\u0007\u0003\u0000"+
		"\u0000#\u0007\u0001\u0000\u0000\u0000$%\u0005 \u0000\u0000%\t\u0001\u0000"+
		"\u0000\u0000&\'\u0006\u0005\uffff\uffff\u0000\'(\u0005$\u0000\u0000()"+
		"\u0003\n\u0005\u0000)*\u0005%\u0000\u0000*1\u0001\u0000\u0000\u0000+1"+
		"\u0005\u0005\u0000\u0000,1\u0003\u0016\u000b\u0000-1\u0003\u001a\r\u0000"+
		"./\u0005\u0013\u0000\u0000/1\u0003\n\u0005\u00010&\u0001\u0000\u0000\u0000"+
		"0+\u0001\u0000\u0000\u00000,\u0001\u0000\u0000\u00000-\u0001\u0000\u0000"+
		"\u00000.\u0001\u0000\u0000\u00001<\u0001\u0000\u0000\u000023\n\u0006\u0000"+
		"\u000034\u0003\u0000\u0000\u000045\u0003\n\u0005\u00075;\u0001\u0000\u0000"+
		"\u000067\n\u0005\u0000\u000078\u0003\u0002\u0001\u000089\u0003\n\u0005"+
		"\u00069;\u0001\u0000\u0000\u0000:2\u0001\u0000\u0000\u0000:6\u0001\u0000"+
		"\u0000\u0000;>\u0001\u0000\u0000\u0000<:\u0001\u0000\u0000\u0000<=\u0001"+
		"\u0000\u0000\u0000=\u000b\u0001\u0000\u0000\u0000><\u0001\u0000\u0000"+
		"\u0000?@\u0006\u0006\uffff\uffff\u0000@A\u0005$\u0000\u0000AB\u0003\f"+
		"\u0006\u0000BC\u0005%\u0000\u0000CJ\u0001\u0000\u0000\u0000DE\u0005 \u0000"+
		"\u0000EJ\u0003\f\u0006\u0004FJ\u0005\u0003\u0000\u0000GJ\u0003\u0016\u000b"+
		"\u0000HJ\u0003\u001a\r\u0000I?\u0001\u0000\u0000\u0000ID\u0001\u0000\u0000"+
		"\u0000IF\u0001\u0000\u0000\u0000IG\u0001\u0000\u0000\u0000IH\u0001\u0000"+
		"\u0000\u0000JU\u0001\u0000\u0000\u0000KL\n\u0006\u0000\u0000LM\u0003\u0004"+
		"\u0002\u0000MN\u0003\f\u0006\u0007NT\u0001\u0000\u0000\u0000OP\n\u0005"+
		"\u0000\u0000PQ\u0003\u0006\u0003\u0000QR\u0003\f\u0006\u0006RT\u0001\u0000"+
		"\u0000\u0000SK\u0001\u0000\u0000\u0000SO\u0001\u0000\u0000\u0000TW\u0001"+
		"\u0000\u0000\u0000US\u0001\u0000\u0000\u0000UV\u0001\u0000\u0000\u0000"+
		"V\r\u0001\u0000\u0000\u0000WU\u0001\u0000\u0000\u0000XY\u0006\u0007\uffff"+
		"\uffff\u0000YZ\u0003\n\u0005\u0000Z[\u0003\u0006\u0003\u0000[\\\u0003"+
		"\n\u0005\u0000\\g\u0001\u0000\u0000\u0000]^\u0003\f\u0006\u0000^_\u0003"+
		"\u0004\u0002\u0000_`\u0003\f\u0006\u0000`g\u0001\u0000\u0000\u0000ab\u0003"+
		"\b\u0004\u0000bc\u0005$\u0000\u0000cd\u0003\u000e\u0007\u0000de\u0005"+
		"%\u0000\u0000eg\u0001\u0000\u0000\u0000fX\u0001\u0000\u0000\u0000f]\u0001"+
		"\u0000\u0000\u0000fa\u0001\u0000\u0000\u0000gn\u0001\u0000\u0000\u0000"+
		"hi\n\u0002\u0000\u0000ij\u0003\u0004\u0002\u0000jk\u0003\f\u0006\u0000"+
		"km\u0001\u0000\u0000\u0000lh\u0001\u0000\u0000\u0000mp\u0001\u0000\u0000"+
		"\u0000nl\u0001\u0000\u0000\u0000no\u0001\u0000\u0000\u0000o\u000f\u0001"+
		"\u0000\u0000\u0000pn\u0001\u0000\u0000\u0000qu\u0003\u0016\u000b\u0000"+
		"ru\u0003\n\u0005\u0000su\u0003\f\u0006\u0000tq\u0001\u0000\u0000\u0000"+
		"tr\u0001\u0000\u0000\u0000ts\u0001\u0000\u0000\u0000u\u0011\u0001\u0000"+
		"\u0000\u0000vw\u0005*\u0000\u0000wx\u0005&\u0000\u0000xy\u0003\n\u0005"+
		"\u0000yz\u0005\'\u0000\u0000z\u0013\u0001\u0000\u0000\u0000{|\u0005*\u0000"+
		"\u0000|}\u0005)\u0000\u0000}\u008b\u0005*\u0000\u0000~\u007f\u0005*\u0000"+
		"\u0000\u007f\u0080\u0005)\u0000\u0000\u0080\u008b\u0003\u0012\t\u0000"+
		"\u0081\u0082\u0005*\u0000\u0000\u0082\u0083\u0005)\u0000\u0000\u0083\u008b"+
		"\u0003\u0014\n\u0000\u0084\u0085\u0003\u0012\t\u0000\u0085\u0086\u0005"+
		")\u0000\u0000\u0086\u0087\u0003\u0014\n\u0000\u0087\u008b\u0001\u0000"+
		"\u0000\u0000\u0088\u008b\u0003\u0012\t\u0000\u0089\u008b\u0005*\u0000"+
		"\u0000\u008a{\u0001\u0000\u0000\u0000\u008a~\u0001\u0000\u0000\u0000\u008a"+
		"\u0081\u0001\u0000\u0000\u0000\u008a\u0084\u0001\u0000\u0000\u0000\u008a"+
		"\u0088\u0001\u0000\u0000\u0000\u008a\u0089\u0001\u0000\u0000\u0000\u008b"+
		"\u0015\u0001\u0000\u0000\u0000\u008c\u0090\u0005*\u0000\u0000\u008d\u0090"+
		"\u0003\u0012\t\u0000\u008e\u0090\u0003\u0014\n\u0000\u008f\u008c\u0001"+
		"\u0000\u0000\u0000\u008f\u008d\u0001\u0000\u0000\u0000\u008f\u008e\u0001"+
		"\u0000\u0000\u0000\u0090\u0017\u0001\u0000\u0000\u0000\u0091\u0096\u0003"+
		"\u0010\b\u0000\u0092\u0093\u0005(\u0000\u0000\u0093\u0095\u0003\u0010"+
		"\b\u0000\u0094\u0092\u0001\u0000\u0000\u0000\u0095\u0098\u0001\u0000\u0000"+
		"\u0000\u0096\u0094\u0001\u0000\u0000\u0000\u0096\u0097\u0001\u0000\u0000"+
		"\u0000\u0097\u0019\u0001\u0000\u0000\u0000\u0098\u0096\u0001\u0000\u0000"+
		"\u0000\u0099\u009a\u0005*\u0000\u0000\u009a\u009c\u0005$\u0000\u0000\u009b"+
		"\u009d\u0003\u0018\f\u0000\u009c\u009b\u0001\u0000\u0000\u0000\u009c\u009d"+
		"\u0001\u0000\u0000\u0000\u009d\u009e\u0001\u0000\u0000\u0000\u009e\u009f"+
		"\u0005%\u0000\u0000\u009f\u001b\u0001\u0000\u0000\u0000\r0:<ISUfnt\u008a"+
		"\u008f\u0096\u009c";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}