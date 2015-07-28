// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package sspi

import "unsafe"
import "syscall"

var _ unsafe.Pointer

var (
	modsecur32 = syscall.NewLazyDLL("secur32.dll")

	procQuerySecurityPackageInfoW  = modsecur32.NewProc("QuerySecurityPackageInfoW")
	procFreeContextBuffer          = modsecur32.NewProc("FreeContextBuffer")
	procAcquireCredentialsHandleW  = modsecur32.NewProc("AcquireCredentialsHandleW")
	procFreeCredentialsHandle      = modsecur32.NewProc("FreeCredentialsHandle")
	procInitializeSecurityContextW = modsecur32.NewProc("InitializeSecurityContextW")
	procAcceptSecurityContext      = modsecur32.NewProc("AcceptSecurityContext")
	procCompleteAuthToken          = modsecur32.NewProc("CompleteAuthToken")
	procDeleteSecurityContext      = modsecur32.NewProc("DeleteSecurityContext")
	procImpersonateSecurityContext = modsecur32.NewProc("ImpersonateSecurityContext")
	procRevertSecurityContext      = modsecur32.NewProc("RevertSecurityContext")
	procQueryContextAttributesW    = modsecur32.NewProc("QueryContextAttributesW")
	procEncryptMessage             = modsecur32.NewProc("EncryptMessage")
	procDecryptMessage             = modsecur32.NewProc("DecryptMessage")
	procApplyControlToken          = modsecur32.NewProc("ApplyControlToken")
)

func QuerySecurityPackageInfo(pkgname *uint16, pkginfo **SecPkgInfo) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procQuerySecurityPackageInfoW.Addr(), 2, uintptr(unsafe.Pointer(pkgname)), uintptr(unsafe.Pointer(pkginfo)), 0)
	ret = syscall.Errno(r0)
	return
}

func FreeContextBuffer(buf *byte) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procFreeContextBuffer.Addr(), 1, uintptr(unsafe.Pointer(buf)), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func AcquireCredentialsHandle(principal *uint16, pkgname *uint16, creduse uint32, logonid *LUID, authdata *byte, getkeyfn uintptr, getkeyarg uintptr, handle *CredHandle, expiry *syscall.Filetime) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall9(procAcquireCredentialsHandleW.Addr(), 9, uintptr(unsafe.Pointer(principal)), uintptr(unsafe.Pointer(pkgname)), uintptr(creduse), uintptr(unsafe.Pointer(logonid)), uintptr(unsafe.Pointer(authdata)), uintptr(getkeyfn), uintptr(getkeyarg), uintptr(unsafe.Pointer(handle)), uintptr(unsafe.Pointer(expiry)))
	ret = syscall.Errno(r0)
	return
}

func FreeCredentialsHandle(handle *CredHandle) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procFreeCredentialsHandle.Addr(), 1, uintptr(unsafe.Pointer(handle)), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func InitializeSecurityContext(credential *CredHandle, context *CtxtHandle, targname *uint16, contextreq uint32, reserved1 uint32, targdatarep uint32, input *SecBufferDesc, reserved2 uint32, newcontext *CtxtHandle, output *SecBufferDesc, contextattr *uint32, expiry *syscall.Filetime) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall12(procInitializeSecurityContextW.Addr(), 12, uintptr(unsafe.Pointer(credential)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(targname)), uintptr(contextreq), uintptr(reserved1), uintptr(targdatarep), uintptr(unsafe.Pointer(input)), uintptr(reserved2), uintptr(unsafe.Pointer(newcontext)), uintptr(unsafe.Pointer(output)), uintptr(unsafe.Pointer(contextattr)), uintptr(unsafe.Pointer(expiry)))
	ret = syscall.Errno(r0)
	return
}

func AcceptSecurityContext(credential *CredHandle, context *CtxtHandle, input *SecBufferDesc, contextreq uint32, targdatarep uint32, newcontext *CtxtHandle, output *SecBufferDesc, contextattr *uint32, expiry *syscall.Filetime) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall9(procAcceptSecurityContext.Addr(), 9, uintptr(unsafe.Pointer(credential)), uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(input)), uintptr(contextreq), uintptr(targdatarep), uintptr(unsafe.Pointer(newcontext)), uintptr(unsafe.Pointer(output)), uintptr(unsafe.Pointer(contextattr)), uintptr(unsafe.Pointer(expiry)))
	ret = syscall.Errno(r0)
	return
}

func CompleteAuthToken(context *CtxtHandle, token *SecBufferDesc) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procCompleteAuthToken.Addr(), 2, uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(token)), 0)
	ret = syscall.Errno(r0)
	return
}

func DeleteSecurityContext(context *CtxtHandle) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procDeleteSecurityContext.Addr(), 1, uintptr(unsafe.Pointer(context)), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func ImpersonateSecurityContext(context *CtxtHandle) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procImpersonateSecurityContext.Addr(), 1, uintptr(unsafe.Pointer(context)), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func RevertSecurityContext(context *CtxtHandle) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procRevertSecurityContext.Addr(), 1, uintptr(unsafe.Pointer(context)), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func QueryContextAttributes(context *CtxtHandle, attribute uint32, buf *byte) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procQueryContextAttributesW.Addr(), 3, uintptr(unsafe.Pointer(context)), uintptr(attribute), uintptr(unsafe.Pointer(buf)))
	ret = syscall.Errno(r0)
	return
}

func EncryptMessage(context *CtxtHandle, qop uint32, message *SecBufferDesc, messageseqno uint32) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall6(procEncryptMessage.Addr(), 4, uintptr(unsafe.Pointer(context)), uintptr(qop), uintptr(unsafe.Pointer(message)), uintptr(messageseqno), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func DecryptMessage(context *CtxtHandle, message *SecBufferDesc, messageseqno uint32, qop *uint32) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall6(procDecryptMessage.Addr(), 4, uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(message)), uintptr(messageseqno), uintptr(unsafe.Pointer(qop)), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func ApplyControlToken(context *CtxtHandle, input *SecBufferDesc) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procApplyControlToken.Addr(), 2, uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(input)), 0)
	ret = syscall.Errno(r0)
	return
}
