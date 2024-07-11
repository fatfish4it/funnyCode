; #NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; ; #Warn  ; Enable warnings to assist with detecting common errors.
; SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
; SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.
!+r::
Reload
Return

; AppsKey::RWin
; MsgBox "call AppsKey" ;不生效
; Return

!+s::
;test-run
; workPath := "O:\test_go\_to_log\"
; run, %ComSpec% /c cd /d %workPath% & go run %workPath%test.go & pause, ,
;v1.1
; run, %ComSpec% /c echo 1234 & pause, ,
;v1.2
; run, %ComSpec% /c echo 123 & ping 127.1 -n 6 >nul, ,
;v1.3
; run, %ComSpec% /c echo 123 & timeout /nobreak /t 5 >nul, ,
;v1.4
; run, %ComSpec% /c echo 123 & timeout /t 5, ,
;v2.0
run, timeout /t 3, ,
Return

;测试专用快捷键
; !+d::
; ; SetKeyDelay, -1
; SendInput {Raw}asdkjfhaksdjfhaksasdkjfhaksdj
; Return


#+x::
; workPath := "O:\test_go\_here\"
; run, %ComSpec% /c cd /d %workPath% & test.exe, ,Hide
; run, %ComSpec% /c log.exe & pause, ,
run, %ComSpec% /c log.exe , ,Hide
; MsgBox % workPath
; MsgBox %path%
Return


; #+f::
; cmdReturn(command=5566)
; {
;     ; WshShell 对象: http://msdn.microsoft.com/en-us/library/aew9yb99
;     shell := ComObjCreate("WScript.Shell")
;     ; 通过 cmd.exe 执行单条命令
;     exec := shell.Exec(ComSpec " /C " command)
;     ; 读取并返回命令的输出
;     return exec.StdOut.ReadAll()
; }
; ; com := cmdReturn("echo 8899")
; ; MsgBox %com%
; MsgBox cmdReturn("O:\\test_go\\_just_print\\print_some.exe")
; Return

; #+p::
!+p::
shell := ComObjCreate("WScript.Shell")
; 通过 cmd.exe 执行单条命令
exec := shell.Exec("cmd.exe /c O:\\test_go\\_just_print\\pp.exe 123 456")
; 读取并返回命令的输出
MsgBox % exec.StdOut.ReadAll()
; MsgBox nnnnn
Return 

; #+q::
; result := GetSelectedText()
; msgbox, %result%
; return

; !+1::
;;清空剪贴板
; Clipboard = 
; Send ^C
; MsgBox % Clipboard
; Return 

;;F7::^ins
;;将F7，设置为复制

