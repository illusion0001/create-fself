package oelf

var _extraLibraryToModule = map[string]string{
	"libSceUserServiceForNpToolkit": "libSceUserService",
	"libSceUserServiceForShellCore": "libSceUserService",
	"libSceUserServiceRegisteredUserIdList": "libSceUserService",
	"libkernel_avlfmem": "libkernel",
	"libkernel_cpumode": "libkernel",
	"libkernel_cpumode_platform": "libkernel",
	"libkernel_dmem_aliasing": "libkernel",
	"libkernel_dmem_aliasing2": "libkernel",
	"libkernel_dmem_aliasing2_for_dev": "libkernel",
	"libkernel_exception": "libkernel",
	"libkernel_jvm": "libkernel",
	"libkernel_module_extension": "libkernel",
	"libkernel_module_info": "libkernel",
	"libkernel_module_load_check": "libkernel",
	"libkernel_pre250mmap": "libkernel",
	"libkernel_ps2emu": "libkernel",
	"libkernel_psmkit": "libkernel",
	"libkernel_qadisc": "libkernel",
	"libkernel_sysc_se": "libkernel",
	"libkernel_unity": "libkernel",
	"libSceCoredump": "libkernel",
	"libSceCoredump_debug": "libkernel",
	"libSceOpenPsId": "libkernel",
	"libScePosix": "libkernel",
	"libSceNpManagerCompat": "libSceNpManager",
	"libSceNpManagerForSys": "libSceNpManager",
	"libSceNpManagerForToolkit": "libSceNpManager",
	"libSceNpManagerIsPlusMember": "libSceNpManager",
	"libkernel_devtoollog": "libkernel",
	"libSceIpmiDbg": "libSceIpmi",
	"libSceGameCustomDataDialog": "libGameCustomDataDialog",
	"libSceGameCustomDataDialogCompat": "libGameCustomDataDialog",
	"libSceWebBrowserDialogLimited": "libSceWebBrowserDialog",
	"libSceNpMatching2Compat": "libSceNpMatching2",
	"libSceGnmDebugModuleReset": "libSceGnmDriver",
	"libSceGnmDebugReset": "libSceGnmDriver",
	"libSceGnmDriverCompat": "libSceGnmDriver",
	"libSceGnmDriverResourceRegistration": "libSceGnmDriver",
	"libSceGnmGetGpuCoreClockFrequency": "libSceGnmDriver",
	"libSceGnmWaitFreeSubmit": "libSceGnmDriver",
	"libSceLibcInternalExt": "libSceLibcInternal",
	"libSceJson2": "libSceJson",
	"libSceDbgKeyboard": "libSceKeyboard",
	"libSceNpAppLauncher": "libSceNpGameIntent",
	"libSceNetDebug": "libSceNet",
	"libSceAppMessaging": "libSceSystemService",
	"libSceLncUtil": "libSceSystemService",
	"libSceShellCoreUtil": "libSceSystemService",
	"libSceSystemService_jvm": "libSceSystemService",
	"libSceSystemServiceActivateHevc": "libSceSystemService",
	"libSceSystemServiceActivateHevcSoft": "libSceSystemService",
	"libSceSystemServiceActivateMpeg2": "libSceSystemService",
	"libSceSystemServiceAppLaunchLink": "libSceSystemService",
	"libSceSystemServiceClosedCaption": "libSceSystemService",
	"libSceSystemServiceDbg": "libSceSystemService",
	"libSceSystemServiceEyeToEyeDistance": "libSceSystemService",
	"libSceSystemServiceForShellCoreOnly": "libSceSystemService",
	"libSceSystemServicePadspkRouting": "libSceSystemService",
	"libSceSystemServicePartyVoice": "libSceSystemService",
	"libSceSystemServicePlatformPrivacy": "libSceSystemService",
	"libSceSystemServicePowerControl": "libSceSystemService",
	"libSceSystemServicePowerSaveLevel": "libSceSystemService",
	"libSceSystemServicePs2Emu": "libSceSystemService",
	"libSceSystemServicePsmKit": "libSceSystemService",
	"libSceSystemServiceStore": "libSceSystemService",
	"libSceSystemServiceSuspend": "libSceSystemService",
	"libSceSystemServiceTelemetry": "libSceSystemService",
	"libSceSystemServiceTournamentMlg": "libSceSystemService",
	"libSceSystemServiceUdsApp": "libSceSystemService",
	"libSceSystemServiceVideoServiceWebApp": "libSceSystemService",
	"libSceSystemServiceVideoToken": "libSceSystemService",
	"libSceSystemServiceVoiceRecognition": "libSceSystemService",
	"libSceSystemServiceWebApp": "libSceSystemService",
	"libSceSystemServiceWebBrowser": "libSceSystemService",
	"libSceSystemServiceYouTubeAccountLinkStatus": "libSceSystemService",
	"libSceSystemStateMgr": "libSceSystemService",
	"libSceNpFriendListDialogCompat": "libSceNpFriendListDialog",
	"libSceNpScoreCompat": "libSceNpScore",
	"libSceHmdDistortion": "libSceHmd",
	"libsceHmdReprojectionMultilayer": "libSceHmd",
	"libSceNpAuthCompat": "libSceNpAuth",
	"libSceNpTusCompat": "libSceNpTus",
	"libSceJpegDec_jvm": "libSceJpegDec",
	"libSceNpUtilityCompat": "libSceNpUtility",
	"libSceNetBwe": "libSceNetCtl",
	"libSceNetCtlAp": "libSceNetCtl",
	"libSceNetCtlApIpcInt": "libSceNetCtl",
	"libSceNetCtlForNpToolkit": "libSceNetCtl",
	"libSceNetCtlV6": "libSceNetCtl",
	"libSceDeviceService": "libSceMbus",
	"libSceMbusDebug": "libSceMbus",
	"libSceOrbisCompatForVideoService": "libSceOrbisCompat",
	"libSceWebSecurity": "libSceOrbisCompat",
	"libSceNpWebApi2AsyncRestricted": "libSceNpWebApi2",
	"libSceAudioDeviceControl": "libSceAudioOut",
	"libSceAudioOutDeviceService": "libSceAudioOut",
	"libSceAudioOutSparkControl": "libSceAudioOut",
	"libSceDbgAudioOut": "libSceAudioOut",
	"libSceNpProfileDialogCompat": "libSceNpProfileDialog",
	"libSceDbgPlayGo": "libScePlayGo",
	"libSceWebKit2ForVideoService": "libSceWebKit2",
	"libSceGLSlimServerVSH": "libSceSlimGLVSH",
	"libSceScreenShotDrc": "libSceScreenShot",
	"libSceNpSns": "libSceNpManager",
	"libSceNpSnsTwitch": "libSceNpManager",
	"libSceNpSnsYouTube": "libSceNpManager",
	"libSceNpSnsTwitchDialog": "libSceNpSnsDialog",
	"libSceNpWebApiCompat": "libSceNpWebApi",
	"libSceUsbStorageAux": "libSceUsbStorage",
	"libSceContentBinder": "libSceContentSearch",
	"libSceInvitationDialogCompat": "libSceInvitationDialog",
	"libSceNpPartyCompat": "libSceNpParty",
	"libSceNpSignalingCompat": "libSceNpSignaling",
	"libc_setjmp": "libc",
	"libSceAsyncStorageInternalAux": "libSceAsyncStorageInternal",
	"libSceAppContent": "libSceAppContentUtil",
	"libSceAppContentBundle": "libSceAppContentUtil",
	"libSceAppContentIro": "libSceAppContentUtil",
	"libSceAppContentPft": "libSceAppContentUtil",
	"libSceAppContentSc": "libSceAppContentUtil",
	"libSceDbgVideoOut": "libSceVideoOut",
	"libSceDbgVideoOutSub4k": "libSceVideoOut",
	"libSceVideoOutAniso": "libSceVideoOut",
	"libSceVideoOutBaseMode4k": "libSceVideoOut",
	"libSceVideoOutExtra": "libSceVideoOut",
	"libSceVideoOutHdr": "libSceVideoOut",
	"libSceVideoOutRawEdid": "libSceVideoOut",
	"libSceVideoOutWindow2Pre400": "libSceVideoOut",
	"libSceSharePlayCompat": "libSceSharePlay",
	"libScePngDec_jvm": "libScePngDec",
	"libSceAvSettingDebug": "libSceAvSetting",
	"libSceGameLiveStreaming_debug": "libSceGameLiveStreaming",
	"libSceGameLiveStreaming_direct_streaming": "libSceGameLiveStreaming",
	"libSceGameLiveStreamingCompat": "libSceGameLiveStreaming",
	"libSceVrTrackerDeviceRejection": "libSceVrTracker",
	"libSceVrTrackerFourDeviceAllowed": "libSceVrTracker",
	"libSceVrTrackerGpuTest": "libSceVrTracker",
	"libSceVrTrackerLiveCapture": "libSceVrTracker",
	"libSceNpCommonCompat": "libSceNpCommon",
}
