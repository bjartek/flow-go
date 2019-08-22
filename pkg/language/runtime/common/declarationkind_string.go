// Code generated by "stringer -type=DeclarationKind"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DeclarationKindUnknown-0]
	_ = x[DeclarationKindValue-1]
	_ = x[DeclarationKindFunction-2]
	_ = x[DeclarationKindVariable-3]
	_ = x[DeclarationKindConstant-4]
	_ = x[DeclarationKindType-5]
	_ = x[DeclarationKindParameter-6]
	_ = x[DeclarationKindArgumentLabel-7]
	_ = x[DeclarationKindStructure-8]
	_ = x[DeclarationKindField-9]
	_ = x[DeclarationKindInitializer-10]
	_ = x[DeclarationKindInterface-11]
	_ = x[DeclarationKindImport-12]
}

const _DeclarationKind_name = "DeclarationKindUnknownDeclarationKindValueDeclarationKindFunctionDeclarationKindVariableDeclarationKindConstantDeclarationKindTypeDeclarationKindParameterDeclarationKindArgumentLabelDeclarationKindStructureDeclarationKindFieldDeclarationKindInitializerDeclarationKindInterfaceDeclarationKindImport"

var _DeclarationKind_index = [...]uint16{0, 22, 42, 65, 88, 111, 130, 154, 182, 206, 226, 252, 276, 297}

func (i DeclarationKind) String() string {
	if i < 0 || i >= DeclarationKind(len(_DeclarationKind_index)-1) {
		return "DeclarationKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeclarationKind_name[_DeclarationKind_index[i]:_DeclarationKind_index[i+1]]
}
