package main

import (
	test "example/test_tool"
	"fmt"
	"strings"
)

func listswitch(index int) {
	switch index {
	case 1:
		test.Halloworld()
	case 2:
		test.Value()
	case 3:
		test.Vartest()
	case 4:
		test.ConstTest()
	case 5:
		test.Forlooptest()
	case 6:
		var n int
		var b string
		for {
			fmt.Println("Enter one number:")
			fmt.Scanln(&n)
			test.IfElse(n)
			fmt.Println("Continue ? Y/N")
			fmt.Scanln(&b)
			if strings.ToUpper(b) == "Y" {
				continue
			} else {
				break
			}
		}
	case 7:
		var n int
		var b string
		for {
			fmt.Println("Enter one number:")
			fmt.Scanln(&n)
			test.Switchtest(n)
			fmt.Println("Continue ? Y/N")
			fmt.Scanln(&b)
			if strings.ToUpper(b) == "Y" {
				continue
			} else {
				break
			}
		}
	case 8:
		test.SingleArr()
		test.TwoDimensionArr()
	case 9:
		test.Slicetest()
	case 10:
		test.Maptest()
	case 11:
		test.Rangetest()
	case 12:
		test.Functest()
	case 13:
		test.Multireturntest()
	case 14:
		test.Varifunctest()
	case 15:
		test.Closuretest()
	case 16:
		test.Recursiontest()
	case 17:
		test.PointerTest()
	case 18:
		test.StringRunestest()
	case 19:
		test.Structuretest()
	case 20:
		test.Methodtest()
	case 21:
		test.Interfacetest()
	case 22:
		test.Embaddingtest()
	case 23:
		test.Genericstest()
	case 24:
		test.Errortest()
	case 25:
		test.Goroutinetest()
	case 26:
		test.Channeltest()
	case 27:
		test.Channelbuffertest()
	case 28:
		test.ChannelSynctest()
	case 29:
		test.Channeldirectiontest()
	case 30:
		test.Selecttest()
	case 31:
		test.Timeouttest()
	case 32:
		test.Nonblockchantest()
	case 33:
		test.Closingchanneltest()
	case 34:
		test.Rangeoverchanneltest()
	case 35:
		test.Timertest()
	case 36:
		test.Tickertest()
	case 37:
		test.Workerpooltest()
	case 38:
		test.Waitgrouptest()
	case 39:
		test.Ratelimittest()
	case 40:
		test.AtomicCountertest()
	case 41:
		test.Mutextest()
	case 42:
		test.Statefulgoroutine()
	case 43:
		test.Sortingtest()
	case 44:
		test.Sortfunctiontest()
	case 45:
		test.Panictest()
	case 46:
		test.Defertest()
	case 47:
		test.Recovertest()
	case 48:
		test.Stringfunctest()
	case 49:
		test.Stringformat()
	case 50:
		test.Texttemplate()
	case 51:
		test.Regularexpression()
	case 52:
		test.JSONtest()
	case 53:
		test.Xmltest()
	case 54:
		test.Timetest()
	case 55:
		test.Epochtest()
	case 56:
		test.Timeformatparsing()
	case 57:
		test.Randomnumtest()
	case 58:
		test.Numparse()
	case 59:
		test.Urlparse()
	case 60:
		test.Hashing()
	case 61:
		test.Encodingb64()
	case 62:
		test.Readfile()
	case 63:
		test.Writefile()
	case 64:
		test.Linefilter()
	case 65:
		test.Filepath()
	case 66:
		test.Director()
	case 67:
		test.TempfileDirector()
	case 68:
		test.Embaddir()
		//	case 69:
		//		test.Testandbenchmarking()
		//	case 70:
		//		Commamdlineargu()
		//	case 71:
		//		Command-Line Flags
		//	case 72:
		//      Command-Line Subcommands
		//	case 73:
		//Environment Variables
	case 74:
		test.Httpclient()
		//	case 75:
		//http server
		//	case 76:
		//server context
	case 77:
		test.Spawningproc()
	case 78:
		test.Execproc()
	case 79:
		test.Signaltest()
		//	case 80:
		//error test
	default:
		fmt.Println("index cannot found")
	}
}

func main() {
	var index int
	for {
		fmt.Println("index from 1~80, index<=0 to exit")
		fmt.Println("choose the index you want:")
		fmt.Scanln(&index)
		listswitch(index)
		if index <= 0 {
			break
		}
	}
	fmt.Println("END THE EXAMPLE")
}
