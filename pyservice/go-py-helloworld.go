package main

import (
    "github.com/sbinet/go-python"
    "fmt"
)

func init() {
    err := python.Initialize()
    if err != nil {
        panic(err.Error())
    }
}

var PyStr = python.PyString_FromString
var GoStr = python.PyString_AS_STRING


// InsertBeforeSysPath will add given dir to python import path
func InsertBeforeSysPath(p string) string {
    sysModule := python.PyImport_ImportModule("sys")
    path := sysModule.GetAttrString("path")
    python.PyList_Insert(path, 0, PyStr(p))
    fmt.Println(path)
    return GoStr(path.Repr())
}

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python.PyObject {
    sysModule := python.PyImport_ImportModule("sys") // import sys
    path := sysModule.GetAttrString("path")                    // path = sys.path
    python.PyList_Insert(path, 0, PyStr(dir))                     // path.insert(0, dir)   
    fmt.Println(path)
    return python.PyImport_ImportModule(name)            // return __import__(name)
}

func main() {
    // import hello
    ret :=InsertBeforeSysPath("/opt/py37/lib/python3.7/site-packages")
    fmt.Println(ret)
    hello := ImportModule("/opt/py_hyg/helloworld", "hello1")
    if hello ==nil{
	python.PyErr_Print()
	fmt.Println("importmodule error!")
    }
    fmt.Println(hello)
    fmt.Printf("[MODULE] repr(hello) = %s\n", GoStr(hello.Repr()))

    // print(hello.a)
    a := hello.GetAttrString("a")
    fmt.Printf("[VARS] a = %#v\n", python.PyInt_AsLong(a))

    // print(hello.b)
    b := hello.GetAttrString("b")
    fmt.Printf("[FUNC] b = %#v\n", b)

    // args = tuple("xixi",)
    bArgs := python.PyTuple_New(1)
    python.PyTuple_SetItem(bArgs, 0, PyStr("xixi"))

    // b(*args)
    res := b.Call(bArgs, python.Py_None)
    fmt.Printf("[CALL] b('xixi') = %s\n", GoStr(res))

    // sklearn
    //sklearn := hello.GetAttrString("sklearn")
    //skVersion := sklearn.GetAttrString("__version__")
    //fmt.Printf("[IMPORT] sklearn = %s\n", GoStr(sklearn.Repr()))
    //fmt.Printf("[IMPORT] sklearn version =  %s\n", GoStr(skVersion.Repr()))
}

