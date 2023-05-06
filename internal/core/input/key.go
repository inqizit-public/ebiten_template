package input

import "fmt"

type KeyInput interface {
	IsKeyPressed(key Key) bool
	IsKeyJustPressed(key Key) bool
	IsKeyJustReleased(key Key) bool
	KeyPressedDuration(key Key) int // return how long the key has been pressed in ticks
	PressedKeys() []Key
	AppendPressedKeys(keys []Key) []Key
	AppendJustPressedKeys(keys []Key) []Key
	AppendJustReleaseKeys(keys []Key) []Key
}

type Key int

const (
	KeyA Key = iota
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyAltLeft
	KeyAltRight
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyArrowUp
	KeyBackquote
	KeyBackslash
	KeyBackspace
	KeyBracketLeft
	KeyBracketRight
	KeyCapsLock
	KeyComma
	KeyContextMenu
	KeyControlLeft
	KeyControlRight
	KeyDelete
	KeyDigit0
	KeyDigit1
	KeyDigit2
	KeyDigit3
	KeyDigit4
	KeyDigit5
	KeyDigit6
	KeyDigit7
	KeyDigit8
	KeyDigit9
	KeyEnd
	KeyEnter
	KeyEqual
	KeyEscape
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyHome
	KeyInsert
	KeyMetaLeft
	KeyMetaRight
	KeyMinus
	KeyNumLock
	KeyNumpad0
	KeyNumpad1
	KeyNumpad2
	KeyNumpad3
	KeyNumpad4
	KeyNumpad5
	KeyNumpad6
	KeyNumpad7
	KeyNumpad8
	KeyNumpad9
	KeyNumpadAdd
	KeyNumpadDecimal
	KeyNumpadDivide
	KeyNumpadEnter
	KeyNumpadEqual
	KeyNumpadMultiply
	KeyNumpadSubtract
	KeyPageDown
	KeyPageUp
	KeyPause
	KeyPeriod
	KeyPrintScreen
	KeyQuote
	KeyScrollLock
	KeySemicolon
	KeyShiftLeft
	KeyShiftRight
	KeySlash
	KeySpace
	KeyTab
	KeyReserved0
	KeyReserved1
	KeyReserved2
	KeyReserved3
)

func (k Key) String() string {
	switch k {
	case KeyA:
		return "KeyA"
	case KeyB:
		return "KeyB"
	case KeyC:
		return "KeyC"
	case KeyD:
		return "KeyD"
	case KeyE:
		return "KeyE"
	case KeyF:
		return "KeyF"
	case KeyG:
		return "KeyG"
	case KeyH:
		return "KeyH"
	case KeyI:
		return "KeyI"
	case KeyJ:
		return "KeyJ"
	case KeyK:
		return "KeyK"
	case KeyL:
		return "KeyL"
	case KeyM:
		return "KeyM"
	case KeyN:
		return "KeyN"
	case KeyO:
		return "KeyO"
	case KeyP:
		return "KeyP"
	case KeyQ:
		return "KeyQ"
	case KeyR:
		return "KeyR"
	case KeyS:
		return "KeyS"
	case KeyT:
		return "KeyT"
	case KeyU:
		return "KeyU"
	case KeyV:
		return "KeyV"
	case KeyW:
		return "KeyW"
	case KeyX:
		return "KeyX"
	case KeyY:
		return "KeyY"
	case KeyZ:
		return "KeyZ"
	case KeyAltLeft:
		return "KeyAltLeft"
	case KeyAltRight:
		return "KeyAltRight"
	case KeyArrowDown:
		return "KeyArrowDown"
	case KeyArrowLeft:
		return "KeyArrowLeft"
	case KeyArrowRight:
		return "KeyArrowRight"
	case KeyArrowUp:
		return "KeyArrowUp"
	case KeyBackquote:
		return "KeyBackquote"
	case KeyBackslash:
		return "KeyBackslash"
	case KeyBackspace:
		return "KeyBackspace"
	case KeyBracketLeft:
		return "KeyBracketLeft"
	case KeyBracketRight:
		return "KeyBracketRight"
	case KeyCapsLock:
		return "KeyCapsLock"
	case KeyComma:
		return "KeyComma"
	case KeyContextMenu:
		return "KeyContextMenu"
	case KeyControlLeft:
		return "KeyControlLeft"
	case KeyControlRight:
		return "KeyControlRight"
	case KeyDelete:
		return "KeyDelete"
	case KeyDigit0:
		return "KeyDigit0"
	case KeyDigit1:
		return "KeyDigit1"
	case KeyDigit2:
		return "KeyDigit2"
	case KeyDigit3:
		return "KeyDigit3"
	case KeyDigit4:
		return "KeyDigit4"
	case KeyDigit5:
		return "KeyDigit5"
	case KeyDigit6:
		return "KeyDigit6"
	case KeyDigit7:
		return "KeyDigit7"
	case KeyDigit8:
		return "KeyDigit8"
	case KeyDigit9:
		return "KeyDigit9"
	case KeyEnd:
		return "KeyEnd"
	case KeyEnter:
		return "KeyEnter"
	case KeyEqual:
		return "KeyEqual"
	case KeyEscape:
		return "KeyEscape"
	case KeyF1:
		return "KeyF1"
	case KeyF2:
		return "KeyF2"
	case KeyF3:
		return "KeyF3"
	case KeyF4:
		return "KeyF4"
	case KeyF5:
		return "KeyF5"
	case KeyF6:
		return "KeyF6"
	case KeyF7:
		return "KeyF7"
	case KeyF8:
		return "KeyF8"
	case KeyF9:
		return "KeyF9"
	case KeyF10:
		return "KeyF10"
	case KeyF11:
		return "KeyF11"
	case KeyF12:
		return "KeyF12"
	case KeyHome:
		return "KeyHome"
	case KeyInsert:
		return "KeyInsert"
	case KeyMetaLeft:
		return "KeyMetaLeft"
	case KeyMetaRight:
		return "KeyMetaRight"
	case KeyMinus:
		return "KeyMinus"
	case KeyNumLock:
		return "KeyNumLock"
	case KeyNumpad0:
		return "KeyNumpad0"
	case KeyNumpad1:
		return "KeyNumpad1"
	case KeyNumpad2:
		return "KeyNumpad2"
	case KeyNumpad3:
		return "KeyNumpad3"
	case KeyNumpad4:
		return "KeyNumpad4"
	case KeyNumpad5:
		return "KeyNumpad5"
	case KeyNumpad6:
		return "KeyNumpad6"
	case KeyNumpad7:
		return "KeyNumpad7"
	case KeyNumpad8:
		return "KeyNumpad8"
	case KeyNumpad9:
		return "KeyNumpad9"
	case KeyNumpadAdd:
		return "KeyNumpadAdd"
	case KeyNumpadDecimal:
		return "KeyNumpadDecimal"
	case KeyNumpadDivide:
		return "KeyNumpadDivide"
	case KeyNumpadEnter:
		return "KeyNumpadEnter"
	case KeyNumpadEqual:
		return "KeyNumpadEqual"
	case KeyNumpadMultiply:
		return "KeyNumpadMultiply"
	case KeyNumpadSubtract:
		return "KeyNumpadSubtract"
	case KeyPageDown:
		return "KeyPageDown"
	case KeyPageUp:
		return "KeyPageUp"
	case KeyPause:
		return "KeyPause"
	case KeyPeriod:
		return "KeyPeriod"
	case KeyPrintScreen:
		return "KeyPrintScreen"
	case KeyQuote:
		return "KeyQuote"
	case KeyScrollLock:
		return "KeyScrollLock"
	case KeySemicolon:
		return "KeySemicolon"
	case KeyShiftLeft:
		return "KeyShiftLeft"
	case KeyShiftRight:
		return "KeyShiftRight"
	case KeySlash:
		return "KeySlash"
	case KeySpace:
		return "KeySpace"
	case KeyTab:
		return "KeyTab"
	}
	panic(fmt.Sprintf("ui: invalid key: %d", k))
}
