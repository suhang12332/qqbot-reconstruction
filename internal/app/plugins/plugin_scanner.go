package plugins

import (
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "reflect"
)

// 自动扫描目录并注册实现了 Plugin 接口的类型
func ScanAndRegisterPlugins(dir string, registry *PluginRegistry) error {
    fset := token.NewFileSet()
    pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
    if err != nil {
        return err
    }

    for _, pkg := range pkgs {
        for _, file := range pkg.Files {
            ast.Inspect(file, func(n ast.Node) bool {
                // 检查所有的类型声明
                if genDecl, ok := n.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
                    for _, spec := range genDecl.Specs {
                        if typeSpec, ok := spec.(*ast.TypeSpec); ok {
                            // 判断类型是否为接口类型
                            if _, ok := typeSpec.Type.(*ast.InterfaceType); ok {
                                continue // 跳过接口类型的定义
                            }

                            // 判断类型是否实现了 Plugin 接口
                            if implementsPlugin(file, typeSpec) {
                                typeName := typeSpec.Name.Name
                                registry.Register(typeName, reflect.TypeOf((*Plugin)(nil)).Elem())
                                fmt.Printf("Registered plugin: %s\n", typeName)
                            }
                        }
                    }
                }
                return true
            })
        }
    }

    return nil
}

// 判断类型是否实现了 Plugin 接口
func implementsPlugin(file *ast.File, typeSpec *ast.TypeSpec) bool {
    methods := map[string]bool{
        "Execute":      false,
        "GetWhiteList": false,
    }

    for _, decl := range file.Decls {
        if funcDecl, ok := decl.(*ast.FuncDecl); ok {
            // 只检查方法为接收者的函数声明
            if funcDecl.Recv != nil {
                recvTypeName := funcDecl.Recv.List[0].Type.(*ast.Ident).Name
                if recvTypeName == typeSpec.Name.Name {
                    if _, ok := methods[funcDecl.Name.Name]; ok {
                        methods[funcDecl.Name.Name] = true
                    }
                }
            }
        }
    }

    for _, implemented := range methods {
        if !implemented {
            return false
        }
    }
    return true
}
