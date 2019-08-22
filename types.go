package lame

/*
#cgo LDFLAGS: -lmp3lame
#include <lame/lame.h>
*/
import "C"

type lameglobal *C.lame_global_flags

// General constants
const (
	bitDepth = 16
)

// MpegMode is a MPEG mode constants type
type MpegMode int

// MPEG modes
const (
	MpegStereo       MpegMode = C.STEREO
	MpegJointStereo  MpegMode = C.JOINT_STEREO
	MpegDualChannel  MpegMode = C.DUAL_CHANNEL /* LAME doesn't supports this! */
	MpegMono         MpegMode = C.MONO
	MpegNotSet       MpegMode = C.NOT_SET
	MpegMaxIndicator MpegMode = C.MAX_INDICATOR /* Don't use this! It's used for sanity checks. */
)

// VBRMode is a VBR mode constants type
type VBRMode int

// VBR modes
const (
	VBROff          VBRMode = C.vbr_off
	VBRMT           VBRMode = C.vbr_mt /* obsolete, same as vbr_mtrh */
	VBRRH           VBRMode = C.vbr_rh
	VBRABR          VBRMode = C.vbr_abr
	VBRMTRH         VBRMode = C.vbr_mtrh
	VBRMaxIndicator VBRMode = C.vbr_max_indicator /* Don't use this! It's used for sanity checks.       */
	VBRDefault      VBRMode = C.vbr_default
)

// PaddingType is a padding type constants type
type PaddingType int

// Padding types
const (
	PadNo           PaddingType = C.PAD_NO
	PadAll          PaddingType = C.PAD_ALL
	PadAdjust       PaddingType = C.PAD_ADJUST
	PadMaxIndicator PaddingType = C.PAD_MAX_INDICATOR /* Don't use this! It's used for sanity checks. */
)

// PresetMode is a preset mode constants type
type PresetMode int

// Preset modes
const (
	PresetABR8   PresetMode = C.ABR_8
	PresetABR320 PresetMode = C.ABR_320
	PresetV9     PresetMode = C.V9
	PresetVBR10  PresetMode = C.VBR_10
	PresetV8     PresetMode = C.V8
	PresetVBR20  PresetMode = C.VBR_20
	PresetV7     PresetMode = C.V7
	PresetVBR30  PresetMode = C.VBR_30
	PresetV6     PresetMode = C.V6
	PresetVBR40  PresetMode = C.VBR_40
	PresetV5     PresetMode = C.V5
	PresetVBR50  PresetMode = C.VBR_50
	PresetV4     PresetMode = C.V4
	PresetVBR60  PresetMode = C.VBR_60
	PresetV3     PresetMode = C.V3
	PresetVBR70  PresetMode = C.VBR_70
	PresetV2     PresetMode = C.V2
	PresetVBR80  PresetMode = C.VBR_80
	PresetV1     PresetMode = C.V1
	PresetVBR90  PresetMode = C.VBR_90
	PresetV0     PresetMode = C.V0
	PresetVBR100 PresetMode = C.VBR_100

	// Compatibility presets
	PresetR3Mix        PresetMode = C.R3MIX
	PresetStandard     PresetMode = C.STANDARD
	PresetExtreme      PresetMode = C.EXTREME
	PresetInsane       PresetMode = C.INSANE
	PresetStandardFast PresetMode = C.STANDARD_FAST
	PresetExtremeFast  PresetMode = C.EXTREME_FAST
	PresetMedium       PresetMode = C.MEDIUM
	PresetMediumFast   PresetMode = C.MEDIUM_FAST
)

// Error lame error type
type Error int

// Lame errors
const (
	ErrorBufferTooSmall         Error = -1
	ErrorMalloc                 Error = -2
	ErrorParamsNotInitialized   Error = -3
	ErrorPsychoAcousticProblems Error = -4
)

func (e Error) Error() string {
	switch e {
	case ErrorBufferTooSmall:
		return "buffer too small"
	case ErrorMalloc:
		return "error allocating memory"
	case ErrorParamsNotInitialized:
		return "lame_init_params not called"
	case ErrorPsychoAcousticProblems:
		return "psycho acoustic problems"
	default:
		return "unknown error"
	}
}

func convError(errCode int) error {
	if errCode >= 0 {
		return nil
	}
	return Error(errCode)
}
