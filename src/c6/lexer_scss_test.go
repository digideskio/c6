package c6

import "c6/ast"
import "testing"

/**
Variable assignment test cases
*/
func TestLexerVariableAssignmentWithCommentBlock(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: /* comment */ #fff /* comment 2 */;`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON, ast.T_HEX_COLOR, ast.T_SEMICOLON})
}

func TestLexerVariableAssignmentHexColor(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: #fff;`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON, ast.T_HEX_COLOR, ast.T_SEMICOLON})
}

func TestLexerVariableAssignmentString(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: 'string';`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON, ast.T_Q_STRING, ast.T_SEMICOLON})
}

func TestLexerVariableAssignmentQQString(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: "string";`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON, ast.T_QQ_STRING, ast.T_SEMICOLON})
}

func TestLexerVariableAssignmentQQStringEscape(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: "str\ning";`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON, ast.T_QQ_STRING, ast.T_SEMICOLON})
}

func TestLexerVariableAssignmentWithDefaultFlag(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: #fff !default;`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON, ast.T_HEX_COLOR, ast.T_DEFAULT, ast.T_SEMICOLON})
}

func TestLexerVariableAssignmentWithImportantFlag(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: #fff !important;`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON, ast.T_HEX_COLOR, ast.T_IMPORTANT, ast.T_SEMICOLON})
}

func TestLexerVariableAssignmentWithOptionalFlag(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: #fff !optional;`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON, ast.T_HEX_COLOR, ast.T_OPTIONAL, ast.T_SEMICOLON})
}

func TestLexerVariableAssignmentWithInterp(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: #{ 10 + 2 }px;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_INTERPOLATION_START, ast.T_INTEGER, ast.T_PLUS, ast.T_INTEGER, ast.T_INTERPOLATION_END, ast.T_LITERAL_CONCAT, ast.T_IDENT, ast.T_SEMICOLON,
	})
}

func TestLexerVariableAssignmentWithIdent(t *testing.T) {
	AssertLexerTokenSequence(t, `$media: screen;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_IDENT, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$feature: -webkit-min-device-pixel-ratio;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_IDENT, ast.T_SEMICOLON,
	})
}

func TestLexerVariableAssignmentWithFloatingNumber(t *testing.T) {
	AssertLexerTokenSequence(t, `$value: 1.5;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_SEMICOLON,
	})
}

func TestLexerVariableAssignmentWithPtValue(t *testing.T) {
	AssertLexerTokenSequence(t, `$foo: 10pt;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_INTEGER, ast.T_UNIT_PT, ast.T_SEMICOLON,
	})
}

func TestLexerVariableAssignmentWithLengthValueStartWithDot(t *testing.T) {
	AssertLexerTokenSequence(t, `$foo: .3em;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_EM, ast.T_SEMICOLON,
	})
}

func TestLexerVariableAssignmentWithLengthValue(t *testing.T) {
	AssertLexerTokenSequence(t, `$foo: 10px;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_INTEGER, ast.T_UNIT_PX, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$foo: 0.3em;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_EM, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$foo: 0.3rem;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_REM, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$foo: 0.3pt;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_PT, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$foo: 0.3in;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_IN, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$foo: 0.3cm;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_CM, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$foo: 0.3ch;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_CH, ast.T_SEMICOLON,
	})
}

func TestLexerVariableAssignmentWithDpiValue(t *testing.T) {
	AssertLexerTokenSequence(t, `$foo: 0.3dpi;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_DPI, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$foo: 0.3dpcm;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_DPCM, ast.T_SEMICOLON,
	})
	AssertLexerTokenSequence(t, `$foo: 0.3dppx;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_FLOAT, ast.T_UNIT_DPPX, ast.T_SEMICOLON,
	})
}

func TestLexerVariableAssignmentWithPercent(t *testing.T) {
	AssertLexerTokenSequence(t, `width: 20%;`, []ast.TokenType{
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON, ast.T_INTEGER, ast.T_UNIT_PERCENT, ast.T_SEMICOLON,
	})
}

func TestLexerMultipleVariableAssignment(t *testing.T) {
	AssertLexerTokenSequence(t, `$favorite: #fff; $foo: 10em;`, []ast.TokenType{
		ast.T_VARIABLE, ast.T_COLON, ast.T_HEX_COLOR, ast.T_SEMICOLON,
		ast.T_VARIABLE, ast.T_COLON, ast.T_INTEGER, ast.T_UNIT_EM, ast.T_SEMICOLON,
	})
}

/**********************************
Map Value test cases
************************************/

func TestLexerMap(t *testing.T) {
	AssertLexerTokenSequence(t, `$var: (foo: 1, bar: 2);`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON,
		ast.T_PAREN_START,
		ast.T_IDENT, ast.T_COLON, ast.T_INTEGER, ast.T_COMMA,
		ast.T_IDENT, ast.T_COLON, ast.T_INTEGER,
		ast.T_PAREN_END,
	})
}

func TestLexerMapWithExtraComma(t *testing.T) {
	AssertLexerTokenSequence(t, `$var: (foo: 1, bar: 2, );`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON,
		ast.T_PAREN_START,
		ast.T_IDENT, ast.T_COLON, ast.T_INTEGER, ast.T_COMMA,
		ast.T_IDENT, ast.T_COLON, ast.T_INTEGER, ast.T_COMMA,
		ast.T_PAREN_END,
	})
}

func TestLexerMapWithExpression(t *testing.T) {
	AssertLexerTokenSequence(t, `$var: (foo: 2px + 3px, bar: $var2);`, []ast.TokenType{ast.T_VARIABLE, ast.T_COLON,
		ast.T_PAREN_START,
		ast.T_IDENT, ast.T_COLON, ast.T_INTEGER, ast.T_UNIT_PX, ast.T_PLUS, ast.T_INTEGER, ast.T_UNIT_PX, ast.T_COMMA,
		ast.T_IDENT, ast.T_COLON, ast.T_VARIABLE,
		ast.T_PAREN_END,
	})
}

func TestLexerMapValue(t *testing.T) {
	AssertLexerTokenSequence(t, `$foo: (foo: 1, bar: 2);`, []ast.TokenType{
		ast.T_VARIABLE,
		ast.T_COLON,
		ast.T_PAREN_START,
		ast.T_IDENT,
		ast.T_COLON,
		ast.T_INTEGER,
		ast.T_COMMA,

		ast.T_IDENT,
		ast.T_COLON,
		ast.T_INTEGER,
		ast.T_PAREN_END,

		ast.T_SEMICOLON,
	})
}

/*************************************************************
Interpolation test cases
******************************************************/

func TestLexerInterpolationPropertyValue(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		-webkit-transition: #{ 1 + 2 }px
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON,
		ast.T_INTERPOLATION_START,
		ast.T_INTEGER,
		ast.T_PLUS,
		ast.T_INTEGER,
		ast.T_INTERPOLATION_END,
		ast.T_LITERAL_CONCAT,
		ast.T_IDENT,
		ast.T_BRACE_END})
}

func TestLexerInterpolationPropertyValueList(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		padding: #{ 1 + 2 }px 10px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON,
		ast.T_INTERPOLATION_START,
		ast.T_INTEGER, // 1
		ast.T_PLUS,
		ast.T_INTEGER, // 2
		ast.T_INTERPOLATION_END,
		ast.T_LITERAL_CONCAT, // px
		ast.T_IDENT,
		ast.T_INTEGER,
		ast.T_UNIT_PX,
		ast.T_SEMICOLON,
		ast.T_BRACE_END})
}

func TestLexerInterpolationComplex1(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		padding: (10+10)#{ 2 * 2 }px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON,
		ast.T_PAREN_START,
		ast.T_INTEGER,
		ast.T_PLUS,
		ast.T_INTEGER,
		ast.T_PAREN_END,
		ast.T_LITERAL_CONCAT,
		ast.T_INTERPOLATION_START,
		ast.T_INTEGER,
		ast.T_MUL,
		ast.T_INTEGER,
		ast.T_INTERPOLATION_END,
		ast.T_LITERAL_CONCAT,
		ast.T_IDENT,
	})
}

func TestLexerInterpolationComplex2(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		padding: (10+10) #{ 2 * 2 } 10px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON,
		ast.T_PAREN_START,
		ast.T_INTEGER,
		ast.T_PLUS,
		ast.T_INTEGER,
		ast.T_PAREN_END,

		ast.T_INTERPOLATION_START,
		ast.T_INTEGER,
		ast.T_MUL,
		ast.T_INTEGER,
		ast.T_INTERPOLATION_END,

		ast.T_INTEGER,
		ast.T_UNIT_PX,
	})
}

func TestLexerInterpolationLeadingAndTrailing(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		padding: 10#{ 1 + 1 }#{ 2 + 2 }px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON,

		ast.T_INTEGER,
		ast.T_LITERAL_CONCAT,

		ast.T_INTERPOLATION_START,
		ast.T_INTEGER, ast.T_PLUS, ast.T_INTEGER,
		ast.T_INTERPOLATION_END,

		ast.T_LITERAL_CONCAT,

		ast.T_INTERPOLATION_START,
		ast.T_INTEGER, ast.T_PLUS, ast.T_INTEGER,
		ast.T_INTERPOLATION_END,

		ast.T_LITERAL_CONCAT,

		ast.T_IDENT,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerInterpolationConcatInterpolation(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		padding: #{ 1 + 2 }#{ 3 + 4 }px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON,

		ast.T_INTERPOLATION_START,
		ast.T_INTEGER, ast.T_PLUS, ast.T_INTEGER,
		ast.T_INTERPOLATION_END,

		ast.T_LITERAL_CONCAT,

		ast.T_INTERPOLATION_START,
		ast.T_INTEGER, ast.T_PLUS, ast.T_INTEGER,
		ast.T_INTERPOLATION_END,

		ast.T_LITERAL_CONCAT,

		ast.T_IDENT,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerInterpolationPropertyValueListWithoutSemiColon(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		padding: #{ 1 + 2 }px 10px
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON,
		ast.T_INTERPOLATION_START,
		ast.T_INTEGER,
		ast.T_PLUS,
		ast.T_INTEGER,
		ast.T_INTERPOLATION_END,
		ast.T_LITERAL_CONCAT,
		ast.T_IDENT,
		ast.T_INTEGER,
		ast.T_UNIT_PX,
		ast.T_BRACE_END,
	})
}

func TestLexerInterpolationPropertyNameSuffix(t *testing.T) {

	AssertLexerTokenSequence(t, `.test {
		padding-#{ 'left' }: 10px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_LITERAL_CONCAT,
		ast.T_INTERPOLATION_START,
		ast.T_Q_STRING,
		ast.T_INTERPOLATION_END,
		ast.T_COLON,
	})

}

func TestLexerInterpolationPropertyName(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		#{ "foo" }: #{ 1 + 2 }px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_INTERPOLATION_START,
		ast.T_QQ_STRING,
		ast.T_INTERPOLATION_END,
		ast.T_COLON,
		ast.T_INTERPOLATION_START,
		ast.T_INTEGER,
		ast.T_PLUS,
		ast.T_INTEGER,
		ast.T_INTERPOLATION_END,
		ast.T_LITERAL_CONCAT,
		ast.T_IDENT,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerInterpolationPropertyName2(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		-#{ "moz" }-border-radius: #{ 1 + 2 }px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_LITERAL_CONCAT,
		ast.T_INTERPOLATION_START,
		ast.T_QQ_STRING,
		ast.T_INTERPOLATION_END,
		ast.T_LITERAL_CONCAT,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_COLON,
		ast.T_INTERPOLATION_START,
		ast.T_INTEGER,
		ast.T_PLUS,
		ast.T_INTEGER,
		ast.T_INTERPOLATION_END,
		ast.T_LITERAL_CONCAT,
		ast.T_IDENT,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerInterpolationPropertyName3(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		padding-#{ $direction }: 10px;
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_LITERAL_CONCAT,
		ast.T_INTERPOLATION_START,
		ast.T_VARIABLE,
		ast.T_INTERPOLATION_END,
		ast.T_COLON,
		ast.T_INTEGER,
		ast.T_UNIT_PX,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerRuleWithSubRule(t *testing.T) {
	AssertLexerTokenSequence(t, `.test {
		-webkit-transition: none;
		.foo { color: #fff; }
	}`, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON, ast.T_IDENT, ast.T_SEMICOLON,
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON, ast.T_HEX_COLOR, ast.T_SEMICOLON,
		ast.T_BRACE_END,
		ast.T_BRACE_END,
	})
}

/**********************************************************************
If Statement Test Case
***********************************************************************/
func TestLexerIfStatementTrueCondition(t *testing.T) {
	AssertLexerTokenSequence(t, `
	@if true {
		color: red;
	}
	`, []ast.TokenType{ast.T_IF, ast.T_TRUE, ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_COLON,
		ast.T_IDENT,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerIfStatementFalseCondition(t *testing.T) {
	AssertLexerTokenSequence(t, `
	@if false {
		color: red;
	}
	`, []ast.TokenType{ast.T_IF, ast.T_FALSE, ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_COLON,
		ast.T_IDENT,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerIfStatementTrueOrFalseCondition(t *testing.T) {
	AssertLexerTokenSequence(t, `
	@if true or false {
		color: red;
	}
	`, []ast.TokenType{
		ast.T_IF, ast.T_TRUE, ast.T_LOGICAL_OR, ast.T_FALSE, ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_COLON,
		ast.T_IDENT,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerIfStatementFunctionCallEqualToNumber(t *testing.T) {
	AssertLexerTokenSequence(t, `
	@if type-of(nth($x, 3)) == 10 {
	}
	`, []ast.TokenType{
		ast.T_IF, ast.T_FUNCTION_NAME, ast.T_PAREN_START, ast.T_FUNCTION_NAME, ast.T_PAREN_START, ast.T_VARIABLE, ast.T_COMMA, ast.T_INTEGER, ast.T_PAREN_END, ast.T_PAREN_END, ast.T_EQUAL, ast.T_INTEGER, ast.T_BRACE_START,
	})
}

/**********************************************************************
For statement
***********************************************************************/
func TestLexerForStatementSimpleFromThrough(t *testing.T) {
	AssertLexerTokenSequence(t, `@for $var from 1 through 20 {  }`, []ast.TokenType{
		ast.T_FOR, ast.T_VARIABLE, ast.T_FOR_FROM, ast.T_INTEGER, ast.T_FOR_THROUGH, ast.T_INTEGER, ast.T_BRACE_START, ast.T_BRACE_END,
	})
}

func TestLexerNestedProperty(t *testing.T) {
	code := `
	.foo {
		border: {
			color: #fff;
			width: 100px;
			style: dashed;
		}
	}
	`
	AssertLexerTokenSequence(t, code, []ast.TokenType{
		ast.T_CLASS_SELECTOR,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_COLON,
		ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_COLON,
		ast.T_HEX_COLOR,
		ast.T_SEMICOLON,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_COLON,
		ast.T_INTEGER,
		ast.T_UNIT_PX,
		ast.T_SEMICOLON,
		ast.T_PROPERTY_NAME_TOKEN,
		ast.T_COLON,
		ast.T_IDENT,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
		ast.T_BRACE_END,
	})
}

/************************************************************
Function test case
*************************************************************/
func TestLexerFunctionGridWidth(t *testing.T) {
	code := `
	@function grid-width($n) {
		@return $n * $grid-width + ($n - 1) * $gutter-width;
	}
	`
	AssertLexerTokenSequence(t, code, []ast.TokenType{
		ast.T_FUNCTION, ast.T_FUNCTION_NAME, ast.T_PAREN_START, ast.T_VARIABLE, ast.T_PAREN_END, ast.T_BRACE_START,
		ast.T_RETURN, ast.T_VARIABLE, ast.T_MUL, ast.T_VARIABLE, ast.T_PLUS, ast.T_PAREN_START, ast.T_VARIABLE, ast.T_MINUS, ast.T_INTEGER, ast.T_PAREN_END, ast.T_MUL, ast.T_VARIABLE,
		ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

/************************************************************
Mixin test cases
*************************************************************/
func TestLexerMixinSimple(t *testing.T) {
	code := `
	@mixin large-text {
	  font-size: 32px;
	}
	`
	AssertLexerTokenSequence(t, code, []ast.TokenType{
		ast.T_MIXIN, ast.T_IDENT, ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON, ast.T_INTEGER, ast.T_UNIT_PX, ast.T_SEMICOLON,
		ast.T_BRACE_END,
	})
}

func TestLexerMixinArguments(t *testing.T) {
	code := `
@mixin sexy-border($color, $width) {
  border: {
    color: $color;
    width: $width;
    style: dashed;
  }
}
`
	AssertLexerTokenSequence(t, code, []ast.TokenType{
		ast.T_MIXIN, ast.T_FUNCTION_NAME, ast.T_PAREN_START, ast.T_VARIABLE, ast.T_COMMA, ast.T_VARIABLE, ast.T_PAREN_END, ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON, ast.T_BRACE_START,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON, ast.T_VARIABLE, ast.T_SEMICOLON,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON, ast.T_VARIABLE, ast.T_SEMICOLON,
		ast.T_PROPERTY_NAME_TOKEN, ast.T_COLON, ast.T_IDENT, ast.T_SEMICOLON,
		ast.T_BRACE_END,
		ast.T_BRACE_END,
	})
}
