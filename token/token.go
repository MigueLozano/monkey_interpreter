package token

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifiers + literals
    IDENT   = "IDENT"   // add, foobar, x, y, ...
    INT     = "INT"     // 1387263

    // Operators
    ASSIGN  = "="
    PLUS    = "+"

    // Delimiters
    COMMA   = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET      = "LET"
)
