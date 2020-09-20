package modlib

import (
	"sync"

	. "github.com/NjinN/RML/go/core"
)

func InitMod(ctx *BindMap) {

	// init Robot start
	var robotObj = BindMap{make(map[string]*Token, 8), ctx, USR_CTX, sync.RWMutex{}}
	var robotToken = Token{Tp: OBJECT, Val: &robotObj}

	var scrollMouseToken = Token{
		NATIVE,
		Native{
			"scroll_mouse",
			3,
			ScrollMouse,
			nil,
		},
	}
	robotObj.PutNow("scroll_mouse", &scrollMouseToken)

	var mouseClickToken = Token{
		NATIVE,
		Native{
			"mouse_click",
			3,
			MouseClick,
			nil,
		},
	}
	robotObj.PutNow("mouse_click", &mouseClickToken)

	var moveMouseToken = Token{
		NATIVE,
		Native{
			"move_mouse",
			3,
			MoveMouse,
			nil,
		},
	}
	robotObj.PutNow("move_mouse", &moveMouseToken)

	var moveMouseSmoothToken = Token{
		NATIVE,
		Native{
			"move_mouse_smooth",
			5,
			MoveMouseSmooth,
			nil,
		},
	}
	robotObj.PutNow("move_mouse_smooth", &moveMouseSmoothToken)

	var typeStrToken = Token{
		NATIVE,
		Native{
			"type_str",
			2,
			TypeStr,
			nil,
		},
	}
	robotObj.PutNow("type_str", &typeStrToken)

	var keyTapToken = Token{
		NATIVE,
		Native{
			"key_tap",
			2,
			KeyTap,
			nil,
		},
	}
	robotObj.PutNow("key_tap", &keyTapToken)

	var ccopyToken = Token{
		NATIVE,
		Native{
			"copy_str",
			2,
			Ccopy,
			nil,
		},
	}
	robotObj.PutNow("copy_str", &ccopyToken)

	var readCopyToken = Token{
		NATIVE,
		Native{
			"read_copy",
			1,
			ReadCopy,
			nil,
		},
	}
	robotObj.PutNow("read_copy", &readCopyToken)

	var getMousePosToken = Token{
		NATIVE,
		Native{
			"mouse_pos",
			1,
			GetMousePos,
			nil,
		},
	}
	robotObj.PutNow("mouse_pos", &getMousePosToken)

	var getPixelColorToken = Token{
		NATIVE,
		Native{
			"pixel_color",
			3,
			GetPixelColor,
			nil,
		},
	}
	robotObj.PutNow("pixel_color", &getPixelColorToken)

	var captureScreenAllToken = Token{
		NATIVE,
		Native{
			"capture_all",
			2,
			CaptureScreenAll,
			nil,
		},
	}
	robotObj.PutNow("capture_all", &captureScreenAllToken)

	var captureScreenToken = Token{
		NATIVE,
		Native{
			"capture",
			6,
			CaptureScreen,
			nil,
		},
	}
	robotObj.PutNow("capture", &captureScreenToken)

	var findPicToken = Token{
		NATIVE,
		Native{
			"find_pic",
			4,
			FindPic,
			nil,
		},
	}
	robotObj.PutNow("find_pic", &findPicToken)

	var matchScreenToken = Token{
		NATIVE,
		Native{
			"match_screen",
			3,
			MatchScreen,
			nil,
		},
	}
	robotObj.PutNow("match_screen", &matchScreenToken)

	var eventHookToken = Token{
		NATIVE,
		Native{
			"event_hook",
			4,
			EventHook,
			nil,
		},
	}
	robotObj.PutNow("event_hook", &eventHookToken)

	var eventHooksToken = Token{
		NATIVE,
		Native{
			"event_hooks",
			2,
			EventHooks,
			nil,
		},
	}
	robotObj.PutNow("event_hooks", &eventHooksToken)

	var startHookToken = Token{
		NATIVE,
		Native{
			"start_hook",
			1,
			StartHook,
			nil,
		},
	}
	robotObj.PutNow("start_hook", &startHookToken)

	var endHookToken = Token{
		NATIVE,
		Native{
			"end_hook",
			1,
			EndHook,
			nil,
		},
	}
	robotObj.PutNow("end_hook", &endHookToken)

	var findIdsToken = Token{
		NATIVE,
		Native{
			"find_pids",
			2,
			FindPids,
			nil,
		},
	}
	robotObj.PutNow("find_pids", &findIdsToken)

	var activePidToken = Token{
		NATIVE,
		Native{
			"active_pid",
			2,
			ActivePid,
			nil,
		},
	}
	robotObj.PutNow("active_pid", &activePidToken)

	var killPidToken = Token{
		NATIVE,
		Native{
			"kill_pid",
			2,
			KillPid,
			nil,
		},
	}
	robotObj.PutNow("kill_pid", &killPidToken)

	var activeProcToken = Token{
		NATIVE,
		Native{
			"active_proc",
			2,
			ActiveProc,
			nil,
		},
	}
	robotObj.PutNow("active_proc", &activeProcToken)

	var pidExistsToken = Token{
		NATIVE,
		Native{
			"pid_exists",
			2,
			PidExists,
			nil,
		},
	}
	robotObj.PutNow("pid_exists", &pidExistsToken)

	var showAlertToken = Token{
		NATIVE,
		Native{
			"show_alert",
			3,
			ShowAlert,
			nil,
		},
	}
	robotObj.PutNow("show_alert", &showAlertToken)

	var getTitleToken = Token{
		NATIVE,
		Native{
			"get_title",
			1,
			GetTitle,
			nil,
		},
	}
	robotObj.PutNow("get_title", &getTitleToken)

	var getPidToken = Token{
		NATIVE,
		Native{
			"get_pid",
			1,
			GetPid,
			nil,
		},
	}
	robotObj.PutNow("get_pid", &getPidToken)

	var is64BitToken = Token{
		NATIVE,
		Native{
			"is64bit",
			1,
			Is64Bit,
			nil,
		},
	}
	robotObj.PutNow("is64bit", &is64BitToken)



	ctx.PutNow("robot", &robotToken)
	// init Robot end
	


	// init wlog start
	var wlogObj = BindMap{make(map[string]*Token, 8), ctx, USR_CTX, sync.RWMutex{}}
	var wlogToken = Token{OBJECT, &wlogObj}

	var logToken = Token{
		NATIVE,
		Native{
			"log",
			2,
			WlogLog,
			nil,
		},
	}
	wlogObj.PutNow("log", &logToken)

	var outputToken = Token{
		NATIVE,
		Native{
			"output",
			2,
			WlogOutput,
			nil,
		},
	}
	wlogObj.PutNow("output", &outputToken)

	var successToken = Token{
		NATIVE,
		Native{
			"success",
			2,
			WlogSuccess,
			nil,
		},
	}
	wlogObj.PutNow("success", &successToken)

	var infoToken = Token{
		NATIVE,
		Native{
			"info",
			2,
			WlogInfo,
			nil,
		},
	}
	wlogObj.PutNow("info", &infoToken)

	var errorToken = Token{
		NATIVE,
		Native{
			"error",
			2,
			WlogError,
			nil,
		},
	}
	wlogObj.PutNow("error", &errorToken)

	var warnToken = Token{
		NATIVE,
		Native{
			"warn",
			2,
			WlogWarn,
			nil,
		},
	}
	wlogObj.PutNow("warn", &warnToken)

	var runningToken = Token{
		NATIVE,
		Native{
			"running",
			2,
			WlogRunning,
			nil,
		},
	}
	wlogObj.PutNow("running", &runningToken)

	var askToken = Token{
		NATIVE,
		Native{
			"ask",
			2,
			WlogAsk,
			nil,
		},
	}
	wlogObj.PutNow("ask", &askToken)


	ctx.PutNow("log", &wlogToken)
	// init wlog end


	// init audio start
	var audioObj = BindMap{make(map[string]*Token, 8), ctx, USR_CTX, sync.RWMutex{}}
	var audioToken = Token{OBJECT, &audioObj}

	var playAudioToken = Token{
		NATIVE,
		Native{
			"play",
			2,
			AudioPlay,
			nil,
		},
	}
	audioObj.PutNow("play", &playAudioToken)

	var playAudioLoopToken = Token{
		NATIVE,
		Native{
			"play-loop",
			2,
			AudioPlayLoop,
			nil,
		},
	}
	audioObj.PutNow("play-loop", &playAudioLoopToken)

	var stopAudioToken = Token{
		NATIVE,
		Native{
			"stop",
			1,
			AudioStop,
			nil,
		},
	}
	audioObj.PutNow("stop", &stopAudioToken)

	ctx.PutNow("audio", &audioToken)
	// init audio end
}
