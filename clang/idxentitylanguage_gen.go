package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type IdxEntityLanguage uint32

const (
	IdxEntityLang_None  IdxEntityLanguage = C.CXIdxEntityLang_None
	IdxEntityLang_C                       = C.CXIdxEntityLang_C
	IdxEntityLang_ObjC                    = C.CXIdxEntityLang_ObjC
	IdxEntityLang_CXX                     = C.CXIdxEntityLang_CXX
	IdxEntityLang_Swift                   = C.CXIdxEntityLang_Swift
)

func (iel IdxEntityLanguage) Spelling() string {
	switch iel {
	case IdxEntityLang_None:
		return "IdxEntityLang=None"
	case IdxEntityLang_C:
		return "IdxEntityLang=C"
	case IdxEntityLang_ObjC:
		return "IdxEntityLang=ObjC"
	case IdxEntityLang_CXX:
		return "IdxEntityLang=CXX"
	case IdxEntityLang_Swift:
		return "IdxEntityLang=Swift"
	}

	return fmt.Sprintf("IdxEntityLanguage unknown %d", int(iel))
}

func (iel IdxEntityLanguage) String() string {
	return iel.Spelling()
}
