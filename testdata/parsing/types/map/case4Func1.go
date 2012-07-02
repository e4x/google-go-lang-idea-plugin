package main
type myHandler struct {
handlers map[string]func(w http.ResponseWriter, r *http.Request, queues *yqs.Queues)
templates map[string]*template.Template
}

/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiElement(WS_NEW_LINES)('\n')
  TypeDeclarationsImpl
    PsiElement(KEYWORD_TYPE)('type')
    PsiWhiteSpace(' ')
    TypeSpecImpl
      TypeNameDeclaration(myHandler)
        PsiElement(IDENTIFIER)('myHandler')
      PsiWhiteSpace(' ')
      TypeStructImpl
        PsiElement(KEYWORD_STRUCT)('struct')
        PsiWhiteSpace(' ')
        PsiElement({)('{')
        PsiElement(WS_NEW_LINES)('\n')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('handlers')
          PsiWhiteSpace(' ')
          TypeMapImpl
            PsiElement(KEYWORD_MAP)('map')
            PsiElement([)('[')
            TypeNameImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('string')
            PsiElement(])(']')
            TypeFunctionImpl
              PsiElement(KEYWORD_FUNC)('func')
              PsiElement(()('(')
              FunctionParameterListImpl
                FunctionParameterImpl
                  LiteralIdentifierImpl
                    PsiElement(IDENTIFIER)('w')
                  PsiWhiteSpace(' ')
                  TypeNameImpl
                    LiteralIdentifierImpl
                      PsiElement(IDENTIFIER)('http')
                      PsiElement(.)('.')
                      PsiElement(IDENTIFIER)('ResponseWriter')
                PsiElement(,)(',')
                PsiWhiteSpace(' ')
                FunctionParameterImpl
                  LiteralIdentifierImpl
                    PsiElement(IDENTIFIER)('r')
                  PsiWhiteSpace(' ')
                  TypePointerImpl
                    PsiElement(*)('*')
                    TypeNameImpl
                      LiteralIdentifierImpl
                        PsiElement(IDENTIFIER)('http')
                        PsiElement(.)('.')
                        PsiElement(IDENTIFIER)('Request')
                PsiElement(,)(',')
                PsiWhiteSpace(' ')
                FunctionParameterImpl
                  LiteralIdentifierImpl
                    PsiElement(IDENTIFIER)('queues')
                  PsiWhiteSpace(' ')
                  TypePointerImpl
                    PsiElement(*)('*')
                    TypeNameImpl
                      LiteralIdentifierImpl
                        PsiElement(IDENTIFIER)('yqs')
                        PsiElement(.)('.')
                        PsiElement(IDENTIFIER)('Queues')
              PsiElement())(')')
        PsiElement(WS_NEW_LINES)('\n')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('templates')
          PsiWhiteSpace(' ')
          TypeMapImpl
            PsiElement(KEYWORD_MAP)('map')
            PsiElement([)('[')
            TypeNameImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('string')
            PsiElement(])(']')
            TypePointerImpl
              PsiElement(*)('*')
              TypeNameImpl
                LiteralIdentifierImpl
                  PsiElement(IDENTIFIER)('template')
                  PsiElement(.)('.')
                  PsiElement(IDENTIFIER)('Template')
        PsiElement(WS_NEW_LINES)('\n')
        PsiElement(})('}')
  PsiElement(WS_NEW_LINES)('\n')