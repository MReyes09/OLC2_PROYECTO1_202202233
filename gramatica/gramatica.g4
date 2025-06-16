grammar gramatica;

// ----------------- LEXER -----------------
INT: [0-9]+;
DOUBLE: [0-9]+ '.' [0-9]+;
CHAR: '\'' . '\'';
STRING: '"' ( '\\' . | ~["\\] )* '"';
BOOL: 'true' | 'false';
BLANCOS: [ \t\r\n]+ -> skip;

ID_VARIABLE: [a-zA-Z_0-9]+;

COMENTARIOLINEA: '//' .*? '\n' -> skip;
COMENTARIOMULTILINEA: '/*' .*? '*/' -> skip;


// ----------------- PARSER -----------------

// ----------------- Lista de instrucciones -----------------
inicio: instrucciones*;

// ----------------- Instrucciones -----------------

instrucciones: imprimir         # PrintStmt
    //| expr                      # ExprStmt 
    | sIf                       # IfStmt
    | sSwitch                   # SwitchInstruccion
    | '{' instrucciones* '}'    # SeccionInstruccion
	| sFor                      # ForStmt
    | varDclSlice               # VarDeclSliceStmt
	| varAsign                  # AsignStmt
	| varDcl                    # VarDeclStmt
    | varDclStruct              # VarDeclStructStmt
    | varStructDcl              # VarStructDclStmt
    | break                     # BreakStmt
    | continue                  # ContinueStmt
    | functions                 # FunctionStmt
    | functionStruct            # FunctionStructStmt
    | varCallStatement          # CallFunctionStmt
    | varCallFuncStruct         # CallFunctionStructStmt
    | retorno                    # ReturnStmt
;

// ----------------- Instruccion imprimir -----------------

imprimir: 'println(' (expr (',' expr)*)?')' ';'? #Println
    | 'print(' (expr(',' expr)*)? ')' ';'?       #Print
;

// ----------------- Sentencias de control -----------------

sIf: 'if' expr block ('else' block)?    # IfOnly
    | 'if' expr block 'else' sIf     # IfAnidado
;

sSwitch: 'switch' expr '{' cases '}' # SwitchStmt
;

cases: 'case' expr ':' instrucciones* cases? # Case
    |'default:' instrucciones*               # Default
;

// ----------------- Bloque de instrucciones -----------------

block: '{' instrucciones* '}'           #blockStmt
;

// ----------------- Sentencia For -----------------
sFor: 'for' expr block                              # ForCondicion
    | 'for' varDcl ';' expr ';' varAsign block      # ForAsignacion
    | 'for' ID_VARIABLE ',' ID_VARIABLE 'in' ID_VARIABLE block # ForRange
;

// ----------------- Declaracion de variables -----------------
varDcl: 'mut' ID_VARIABLE type assign expr # VarDclWithTypeAndValue
      | 'mut' ID_VARIABLE type          # VarDclWithTypeOnly
      | 'mut'? ID_VARIABLE ':=' expr           # VarDclWithInference
;

varDclSlice: ID_VARIABLE assign (nuevoSlice)+ type '{' contenidoSlice '}' # SliceValores
    | 'mut' ID_VARIABLE (nuevoSlice)+ type                                # SliceVacio
    | 'mut' ID_VARIABLE (nuevoSlice)+ type assign expr                    # SliceDcl_Asign
    
;

assign: ':=' 
    | '='
;

nuevoSlice: '[]'
;

contenidoSlice: expr (',' expr)*                                # SliceContenido
    | '{' contenidoSlice '}' (',' ('{' contenidoSlice '}')?)*      # SliceContenidoSlice
;
// ----------------- Declaracion de estructuras -----------------
varDclStruct: 'type'? 'struct'? ID_VARIABLE 'struct'? '{'( type ID_VARIABLE ';'?)+'}'   # DeclStructData
;

varStructDcl: ID_VARIABLE ':=' ID_VARIABLE '{' ID_VARIABLE ':' expr (',' (ID_VARIABLE ':' expr)?)+ '}'  # StructVarTypeInference
;

// ----------------- Asignacion de variables -----------------
varAsign: ID_VARIABLE '=' expr          # varExpr
    | ID_VARIABLE op =('+='|'-=') expr  # varAdd    
    | ID_VARIABLE op = ('++'|'--')      # varInc
    | ID_VARIABLE ('[' expr ']')+ '=' expr # ArrayAccess
    | ID_VARIABLE ('.' ID_VARIABLE)+ '=' expr # StructAccessAsign
;

expr: '-' expr                                                # Negate
    | '!' expr                                              # Not
    | expr op = ('*' | '/' | '%') expr                      # MulDivModulo
    | expr op = ('+' | '-') expr                            # AddSub
    | expr op = ('<' | '>' | '<=' | '>=') expr              # MinorMajorEqual
    | expr op = ('==' | '!=') expr                          # EqualsNotEquals
    | expr op = ('&&' | '||') expr                          # Logical
    | INT                                                   # Integer
    | DOUBLE                                                # Double
    | STRING                                                # String
    | BOOL                                                  # Boolean
    | ID_VARIABLE                                           # Identifier
    | CHAR                                                  # Char
    | 'nil'                                                 # Nil    
    | '(' expr ')'                                          # Parens
    //Acceso a arreglos
    | ID_VARIABLE ('[' expr ']')+                           # ArrayAccessSimple
    //Funciones embebidas
    | 'indexOf('ID_VARIABLE ',' expr ')'               # ArrayFindIndex
    | 'join('ID_VARIABLE ',' expr ')'               # ArrayJoin
    | 'len('ID_VARIABLE (posicion)* ')'                                 # ArrayLength
    | 'append('ID_VARIABLE ',' expr ')'                     # ArrayAppend
    | 'Atoi(' expr ')'                                      # IntToString
    | 'parseFloat(' expr ')'                                  # floatToString
    | 'typeOf(' expr ')'                                      # reflectType
    //Acceso a estructuras
    | ID_VARIABLE ('.' ID_VARIABLE)+ ';'?                   # StructAccess 
    | varCallStatement                                      # CallFunctionValue
    | varCallFuncStruct                                     # CallFunctionStructValue   
;

posicion: '[' expr ']'
;

type: 'int' 
    | 'float64' 
    | 'string' 
    | 'bool' 
    | 'rune'
    | 'struct'
    | ID_VARIABLE
;

break: 'break' ';'?
;

continue: 'continue' ';'?
;

functions: 'func' ID_VARIABLE '(' (ID_VARIABLE type (',' ID_VARIABLE type)*)? ')' valRet? block # Funciones
;

functionStruct: 'func' '(' ID_VARIABLE ID_VARIABLE ')'  ID_VARIABLE '(' defParams? ')' valRet? block   # FuncionesStructsNativas
;

defParams: ID_VARIABLE type (',' ID_VARIABLE type)*
;

varCallStatement: ID_VARIABLE '(' (expr (',' expr)*)? ')' ';'? # CallFunction
;

varCallFuncStruct: ID_VARIABLE '.' ID_VARIABLE '(' (expr (',' expr)*)? ')' ';'? # CallFunctionStruct
;

valRet: type
;

retorno: 'return' expr? ';'?
;
