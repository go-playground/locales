package localeslist

import (
	"sync"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/af"
	"github.com/go-playground/locales/af_NA"
	"github.com/go-playground/locales/af_ZA"
	"github.com/go-playground/locales/agq"
	"github.com/go-playground/locales/agq_CM"
	"github.com/go-playground/locales/ak"
	"github.com/go-playground/locales/ak_GH"
	"github.com/go-playground/locales/am"
	"github.com/go-playground/locales/am_ET"
	"github.com/go-playground/locales/ar"
	"github.com/go-playground/locales/ar_001"
	"github.com/go-playground/locales/ar_AE"
	"github.com/go-playground/locales/ar_BH"
	"github.com/go-playground/locales/ar_DJ"
	"github.com/go-playground/locales/ar_DZ"
	"github.com/go-playground/locales/ar_EG"
	"github.com/go-playground/locales/ar_EH"
	"github.com/go-playground/locales/ar_ER"
	"github.com/go-playground/locales/ar_IL"
	"github.com/go-playground/locales/ar_IQ"
	"github.com/go-playground/locales/ar_JO"
	"github.com/go-playground/locales/ar_KM"
	"github.com/go-playground/locales/ar_KW"
	"github.com/go-playground/locales/ar_LB"
	"github.com/go-playground/locales/ar_LY"
	"github.com/go-playground/locales/ar_MA"
	"github.com/go-playground/locales/ar_MR"
	"github.com/go-playground/locales/ar_OM"
	"github.com/go-playground/locales/ar_PS"
	"github.com/go-playground/locales/ar_QA"
	"github.com/go-playground/locales/ar_SA"
	"github.com/go-playground/locales/ar_SD"
	"github.com/go-playground/locales/ar_SO"
	"github.com/go-playground/locales/ar_SS"
	"github.com/go-playground/locales/ar_SY"
	"github.com/go-playground/locales/ar_TD"
	"github.com/go-playground/locales/ar_TN"
	"github.com/go-playground/locales/ar_YE"
	"github.com/go-playground/locales/as"
	"github.com/go-playground/locales/as_IN"
	"github.com/go-playground/locales/asa"
	"github.com/go-playground/locales/asa_TZ"
	"github.com/go-playground/locales/ast"
	"github.com/go-playground/locales/ast_ES"
	"github.com/go-playground/locales/az"
	"github.com/go-playground/locales/az_Cyrl"
	"github.com/go-playground/locales/az_Cyrl_AZ"
	"github.com/go-playground/locales/az_Latn"
	"github.com/go-playground/locales/az_Latn_AZ"
	"github.com/go-playground/locales/bas"
	"github.com/go-playground/locales/bas_CM"
	"github.com/go-playground/locales/be"
	"github.com/go-playground/locales/be_BY"
	"github.com/go-playground/locales/bem"
	"github.com/go-playground/locales/bem_ZM"
	"github.com/go-playground/locales/bez"
	"github.com/go-playground/locales/bez_TZ"
	"github.com/go-playground/locales/bg"
	"github.com/go-playground/locales/bg_BG"
	"github.com/go-playground/locales/bm"
	"github.com/go-playground/locales/bm_ML"
	"github.com/go-playground/locales/bn"
	"github.com/go-playground/locales/bn_BD"
	"github.com/go-playground/locales/bn_IN"
	"github.com/go-playground/locales/bo"
	"github.com/go-playground/locales/bo_CN"
	"github.com/go-playground/locales/bo_IN"
	"github.com/go-playground/locales/br"
	"github.com/go-playground/locales/br_FR"
	"github.com/go-playground/locales/brx"
	"github.com/go-playground/locales/brx_IN"
	"github.com/go-playground/locales/bs"
	"github.com/go-playground/locales/bs_Cyrl"
	"github.com/go-playground/locales/bs_Cyrl_BA"
	"github.com/go-playground/locales/bs_Latn"
	"github.com/go-playground/locales/bs_Latn_BA"
	"github.com/go-playground/locales/ca"
	"github.com/go-playground/locales/ca_AD"
	"github.com/go-playground/locales/ca_ES"
	"github.com/go-playground/locales/ca_ES_VALENCIA"
	"github.com/go-playground/locales/ca_FR"
	"github.com/go-playground/locales/ca_IT"
	"github.com/go-playground/locales/ce"
	"github.com/go-playground/locales/ce_RU"
	"github.com/go-playground/locales/cgg"
	"github.com/go-playground/locales/cgg_UG"
	"github.com/go-playground/locales/chr"
	"github.com/go-playground/locales/chr_US"
	"github.com/go-playground/locales/ckb"
	"github.com/go-playground/locales/ckb_IQ"
	"github.com/go-playground/locales/ckb_IR"
	"github.com/go-playground/locales/cs"
	"github.com/go-playground/locales/cs_CZ"
	"github.com/go-playground/locales/cu"
	"github.com/go-playground/locales/cu_RU"
	"github.com/go-playground/locales/cy"
	"github.com/go-playground/locales/cy_GB"
	"github.com/go-playground/locales/da"
	"github.com/go-playground/locales/da_DK"
	"github.com/go-playground/locales/da_GL"
	"github.com/go-playground/locales/dav"
	"github.com/go-playground/locales/dav_KE"
	"github.com/go-playground/locales/de"
	"github.com/go-playground/locales/de_AT"
	"github.com/go-playground/locales/de_BE"
	"github.com/go-playground/locales/de_CH"
	"github.com/go-playground/locales/de_DE"
	"github.com/go-playground/locales/de_LI"
	"github.com/go-playground/locales/de_LU"
	"github.com/go-playground/locales/dje"
	"github.com/go-playground/locales/dje_NE"
	"github.com/go-playground/locales/dsb"
	"github.com/go-playground/locales/dsb_DE"
	"github.com/go-playground/locales/dua"
	"github.com/go-playground/locales/dua_CM"
	"github.com/go-playground/locales/dyo"
	"github.com/go-playground/locales/dyo_SN"
	"github.com/go-playground/locales/dz"
	"github.com/go-playground/locales/dz_BT"
	"github.com/go-playground/locales/ebu"
	"github.com/go-playground/locales/ebu_KE"
	"github.com/go-playground/locales/ee"
	"github.com/go-playground/locales/ee_GH"
	"github.com/go-playground/locales/ee_TG"
	"github.com/go-playground/locales/el"
	"github.com/go-playground/locales/el_CY"
	"github.com/go-playground/locales/el_GR"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/en_001"
	"github.com/go-playground/locales/en_150"
	"github.com/go-playground/locales/en_AG"
	"github.com/go-playground/locales/en_AI"
	"github.com/go-playground/locales/en_AS"
	"github.com/go-playground/locales/en_AT"
	"github.com/go-playground/locales/en_AU"
	"github.com/go-playground/locales/en_BB"
	"github.com/go-playground/locales/en_BE"
	"github.com/go-playground/locales/en_BI"
	"github.com/go-playground/locales/en_BM"
	"github.com/go-playground/locales/en_BS"
	"github.com/go-playground/locales/en_BW"
	"github.com/go-playground/locales/en_BZ"
	"github.com/go-playground/locales/en_CA"
	"github.com/go-playground/locales/en_CC"
	"github.com/go-playground/locales/en_CH"
	"github.com/go-playground/locales/en_CK"
	"github.com/go-playground/locales/en_CM"
	"github.com/go-playground/locales/en_CX"
	"github.com/go-playground/locales/en_CY"
	"github.com/go-playground/locales/en_DE"
	"github.com/go-playground/locales/en_DG"
	"github.com/go-playground/locales/en_DK"
	"github.com/go-playground/locales/en_DM"
	"github.com/go-playground/locales/en_ER"
	"github.com/go-playground/locales/en_FI"
	"github.com/go-playground/locales/en_FJ"
	"github.com/go-playground/locales/en_FK"
	"github.com/go-playground/locales/en_FM"
	"github.com/go-playground/locales/en_GB"
	"github.com/go-playground/locales/en_GD"
	"github.com/go-playground/locales/en_GG"
	"github.com/go-playground/locales/en_GH"
	"github.com/go-playground/locales/en_GI"
	"github.com/go-playground/locales/en_GM"
	"github.com/go-playground/locales/en_GU"
	"github.com/go-playground/locales/en_GY"
	"github.com/go-playground/locales/en_HK"
	"github.com/go-playground/locales/en_IE"
	"github.com/go-playground/locales/en_IL"
	"github.com/go-playground/locales/en_IM"
	"github.com/go-playground/locales/en_IN"
	"github.com/go-playground/locales/en_IO"
	"github.com/go-playground/locales/en_JE"
	"github.com/go-playground/locales/en_JM"
	"github.com/go-playground/locales/en_KE"
	"github.com/go-playground/locales/en_KI"
	"github.com/go-playground/locales/en_KN"
	"github.com/go-playground/locales/en_KY"
	"github.com/go-playground/locales/en_LC"
	"github.com/go-playground/locales/en_LR"
	"github.com/go-playground/locales/en_LS"
	"github.com/go-playground/locales/en_MG"
	"github.com/go-playground/locales/en_MH"
	"github.com/go-playground/locales/en_MO"
	"github.com/go-playground/locales/en_MP"
	"github.com/go-playground/locales/en_MS"
	"github.com/go-playground/locales/en_MT"
	"github.com/go-playground/locales/en_MU"
	"github.com/go-playground/locales/en_MW"
	"github.com/go-playground/locales/en_MY"
	"github.com/go-playground/locales/en_NA"
	"github.com/go-playground/locales/en_NF"
	"github.com/go-playground/locales/en_NG"
	"github.com/go-playground/locales/en_NL"
	"github.com/go-playground/locales/en_NR"
	"github.com/go-playground/locales/en_NU"
	"github.com/go-playground/locales/en_NZ"
	"github.com/go-playground/locales/en_PG"
	"github.com/go-playground/locales/en_PH"
	"github.com/go-playground/locales/en_PK"
	"github.com/go-playground/locales/en_PN"
	"github.com/go-playground/locales/en_PR"
	"github.com/go-playground/locales/en_PW"
	"github.com/go-playground/locales/en_RW"
	"github.com/go-playground/locales/en_SB"
	"github.com/go-playground/locales/en_SC"
	"github.com/go-playground/locales/en_SD"
	"github.com/go-playground/locales/en_SE"
	"github.com/go-playground/locales/en_SG"
	"github.com/go-playground/locales/en_SH"
	"github.com/go-playground/locales/en_SI"
	"github.com/go-playground/locales/en_SL"
	"github.com/go-playground/locales/en_SS"
	"github.com/go-playground/locales/en_SX"
	"github.com/go-playground/locales/en_SZ"
	"github.com/go-playground/locales/en_TC"
	"github.com/go-playground/locales/en_TK"
	"github.com/go-playground/locales/en_TO"
	"github.com/go-playground/locales/en_TT"
	"github.com/go-playground/locales/en_TV"
	"github.com/go-playground/locales/en_TZ"
	"github.com/go-playground/locales/en_UG"
	"github.com/go-playground/locales/en_UM"
	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/locales/en_US_POSIX"
	"github.com/go-playground/locales/en_VC"
	"github.com/go-playground/locales/en_VG"
	"github.com/go-playground/locales/en_VI"
	"github.com/go-playground/locales/en_VU"
	"github.com/go-playground/locales/en_WS"
	"github.com/go-playground/locales/en_ZA"
	"github.com/go-playground/locales/en_ZM"
	"github.com/go-playground/locales/en_ZW"
	"github.com/go-playground/locales/eo"
	"github.com/go-playground/locales/eo_001"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/es_419"
	"github.com/go-playground/locales/es_AR"
	"github.com/go-playground/locales/es_BO"
	"github.com/go-playground/locales/es_BR"
	"github.com/go-playground/locales/es_CL"
	"github.com/go-playground/locales/es_CO"
	"github.com/go-playground/locales/es_CR"
	"github.com/go-playground/locales/es_CU"
	"github.com/go-playground/locales/es_DO"
	"github.com/go-playground/locales/es_EA"
	"github.com/go-playground/locales/es_EC"
	"github.com/go-playground/locales/es_ES"
	"github.com/go-playground/locales/es_GQ"
	"github.com/go-playground/locales/es_GT"
	"github.com/go-playground/locales/es_HN"
	"github.com/go-playground/locales/es_IC"
	"github.com/go-playground/locales/es_MX"
	"github.com/go-playground/locales/es_NI"
	"github.com/go-playground/locales/es_PA"
	"github.com/go-playground/locales/es_PE"
	"github.com/go-playground/locales/es_PH"
	"github.com/go-playground/locales/es_PR"
	"github.com/go-playground/locales/es_PY"
	"github.com/go-playground/locales/es_SV"
	"github.com/go-playground/locales/es_US"
	"github.com/go-playground/locales/es_UY"
	"github.com/go-playground/locales/es_VE"
	"github.com/go-playground/locales/et"
	"github.com/go-playground/locales/et_EE"
	"github.com/go-playground/locales/eu"
	"github.com/go-playground/locales/eu_ES"
	"github.com/go-playground/locales/ewo"
	"github.com/go-playground/locales/ewo_CM"
	"github.com/go-playground/locales/fa"
	"github.com/go-playground/locales/fa_AF"
	"github.com/go-playground/locales/fa_IR"
	"github.com/go-playground/locales/ff"
	"github.com/go-playground/locales/ff_CM"
	"github.com/go-playground/locales/ff_GN"
	"github.com/go-playground/locales/ff_MR"
	"github.com/go-playground/locales/ff_SN"
	"github.com/go-playground/locales/fi"
	"github.com/go-playground/locales/fi_FI"
	"github.com/go-playground/locales/fil"
	"github.com/go-playground/locales/fil_PH"
	"github.com/go-playground/locales/fo"
	"github.com/go-playground/locales/fo_DK"
	"github.com/go-playground/locales/fo_FO"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/fr_BE"
	"github.com/go-playground/locales/fr_BF"
	"github.com/go-playground/locales/fr_BI"
	"github.com/go-playground/locales/fr_BJ"
	"github.com/go-playground/locales/fr_BL"
	"github.com/go-playground/locales/fr_CA"
	"github.com/go-playground/locales/fr_CD"
	"github.com/go-playground/locales/fr_CF"
	"github.com/go-playground/locales/fr_CG"
	"github.com/go-playground/locales/fr_CH"
	"github.com/go-playground/locales/fr_CI"
	"github.com/go-playground/locales/fr_CM"
	"github.com/go-playground/locales/fr_DJ"
	"github.com/go-playground/locales/fr_DZ"
	"github.com/go-playground/locales/fr_FR"
	"github.com/go-playground/locales/fr_GA"
	"github.com/go-playground/locales/fr_GF"
	"github.com/go-playground/locales/fr_GN"
	"github.com/go-playground/locales/fr_GP"
	"github.com/go-playground/locales/fr_GQ"
	"github.com/go-playground/locales/fr_HT"
	"github.com/go-playground/locales/fr_KM"
	"github.com/go-playground/locales/fr_LU"
	"github.com/go-playground/locales/fr_MA"
	"github.com/go-playground/locales/fr_MC"
	"github.com/go-playground/locales/fr_MF"
	"github.com/go-playground/locales/fr_MG"
	"github.com/go-playground/locales/fr_ML"
	"github.com/go-playground/locales/fr_MQ"
	"github.com/go-playground/locales/fr_MR"
	"github.com/go-playground/locales/fr_MU"
	"github.com/go-playground/locales/fr_NC"
	"github.com/go-playground/locales/fr_NE"
	"github.com/go-playground/locales/fr_PF"
	"github.com/go-playground/locales/fr_PM"
	"github.com/go-playground/locales/fr_RE"
	"github.com/go-playground/locales/fr_RW"
	"github.com/go-playground/locales/fr_SC"
	"github.com/go-playground/locales/fr_SN"
	"github.com/go-playground/locales/fr_SY"
	"github.com/go-playground/locales/fr_TD"
	"github.com/go-playground/locales/fr_TG"
	"github.com/go-playground/locales/fr_TN"
	"github.com/go-playground/locales/fr_VU"
	"github.com/go-playground/locales/fr_WF"
	"github.com/go-playground/locales/fr_YT"
	"github.com/go-playground/locales/fur"
	"github.com/go-playground/locales/fur_IT"
	"github.com/go-playground/locales/fy"
	"github.com/go-playground/locales/fy_NL"
	"github.com/go-playground/locales/ga"
	"github.com/go-playground/locales/ga_IE"
	"github.com/go-playground/locales/gd"
	"github.com/go-playground/locales/gd_GB"
	"github.com/go-playground/locales/gl"
	"github.com/go-playground/locales/gl_ES"
	"github.com/go-playground/locales/gsw"
	"github.com/go-playground/locales/gsw_CH"
	"github.com/go-playground/locales/gsw_FR"
	"github.com/go-playground/locales/gsw_LI"
	"github.com/go-playground/locales/gu"
	"github.com/go-playground/locales/gu_IN"
	"github.com/go-playground/locales/guz"
	"github.com/go-playground/locales/guz_KE"
	"github.com/go-playground/locales/gv"
	"github.com/go-playground/locales/gv_IM"
	"github.com/go-playground/locales/ha"
	"github.com/go-playground/locales/ha_GH"
	"github.com/go-playground/locales/ha_NE"
	"github.com/go-playground/locales/ha_NG"
	"github.com/go-playground/locales/haw"
	"github.com/go-playground/locales/haw_US"
	"github.com/go-playground/locales/he"
	"github.com/go-playground/locales/he_IL"
	"github.com/go-playground/locales/hi"
	"github.com/go-playground/locales/hi_IN"
	"github.com/go-playground/locales/hr"
	"github.com/go-playground/locales/hr_BA"
	"github.com/go-playground/locales/hr_HR"
	"github.com/go-playground/locales/hsb"
	"github.com/go-playground/locales/hsb_DE"
	"github.com/go-playground/locales/hu"
	"github.com/go-playground/locales/hu_HU"
	"github.com/go-playground/locales/hy"
	"github.com/go-playground/locales/hy_AM"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/locales/id_ID"
	"github.com/go-playground/locales/ig"
	"github.com/go-playground/locales/ig_NG"
	"github.com/go-playground/locales/ii"
	"github.com/go-playground/locales/ii_CN"
	"github.com/go-playground/locales/is"
	"github.com/go-playground/locales/is_IS"
	"github.com/go-playground/locales/it"
	"github.com/go-playground/locales/it_CH"
	"github.com/go-playground/locales/it_IT"
	"github.com/go-playground/locales/it_SM"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/ja_JP"
	"github.com/go-playground/locales/jgo"
	"github.com/go-playground/locales/jgo_CM"
	"github.com/go-playground/locales/jmc"
	"github.com/go-playground/locales/jmc_TZ"
	"github.com/go-playground/locales/ka"
	"github.com/go-playground/locales/ka_GE"
	"github.com/go-playground/locales/kab"
	"github.com/go-playground/locales/kab_DZ"
	"github.com/go-playground/locales/kam"
	"github.com/go-playground/locales/kam_KE"
	"github.com/go-playground/locales/kde"
	"github.com/go-playground/locales/kde_TZ"
	"github.com/go-playground/locales/kea"
	"github.com/go-playground/locales/kea_CV"
	"github.com/go-playground/locales/khq"
	"github.com/go-playground/locales/khq_ML"
	"github.com/go-playground/locales/ki"
	"github.com/go-playground/locales/ki_KE"
	"github.com/go-playground/locales/kk"
	"github.com/go-playground/locales/kk_KZ"
	"github.com/go-playground/locales/kkj"
	"github.com/go-playground/locales/kkj_CM"
	"github.com/go-playground/locales/kl"
	"github.com/go-playground/locales/kl_GL"
	"github.com/go-playground/locales/kln"
	"github.com/go-playground/locales/kln_KE"
	"github.com/go-playground/locales/km"
	"github.com/go-playground/locales/km_KH"
	"github.com/go-playground/locales/kn"
	"github.com/go-playground/locales/kn_IN"
	"github.com/go-playground/locales/ko"
	"github.com/go-playground/locales/ko_KP"
	"github.com/go-playground/locales/ko_KR"
	"github.com/go-playground/locales/kok"
	"github.com/go-playground/locales/kok_IN"
	"github.com/go-playground/locales/ks"
	"github.com/go-playground/locales/ks_IN"
	"github.com/go-playground/locales/ksb"
	"github.com/go-playground/locales/ksb_TZ"
	"github.com/go-playground/locales/ksf"
	"github.com/go-playground/locales/ksf_CM"
	"github.com/go-playground/locales/ksh"
	"github.com/go-playground/locales/ksh_DE"
	"github.com/go-playground/locales/kw"
	"github.com/go-playground/locales/kw_GB"
	"github.com/go-playground/locales/ky"
	"github.com/go-playground/locales/ky_KG"
	"github.com/go-playground/locales/lag"
	"github.com/go-playground/locales/lag_TZ"
	"github.com/go-playground/locales/lb"
	"github.com/go-playground/locales/lb_LU"
	"github.com/go-playground/locales/lg"
	"github.com/go-playground/locales/lg_UG"
	"github.com/go-playground/locales/lkt"
	"github.com/go-playground/locales/lkt_US"
	"github.com/go-playground/locales/ln"
	"github.com/go-playground/locales/ln_AO"
	"github.com/go-playground/locales/ln_CD"
	"github.com/go-playground/locales/ln_CF"
	"github.com/go-playground/locales/ln_CG"
	"github.com/go-playground/locales/lo"
	"github.com/go-playground/locales/lo_LA"
	"github.com/go-playground/locales/lrc"
	"github.com/go-playground/locales/lrc_IQ"
	"github.com/go-playground/locales/lrc_IR"
	"github.com/go-playground/locales/lt"
	"github.com/go-playground/locales/lt_LT"
	"github.com/go-playground/locales/lu"
	"github.com/go-playground/locales/lu_CD"
	"github.com/go-playground/locales/luo"
	"github.com/go-playground/locales/luo_KE"
	"github.com/go-playground/locales/luy"
	"github.com/go-playground/locales/luy_KE"
	"github.com/go-playground/locales/lv"
	"github.com/go-playground/locales/lv_LV"
	"github.com/go-playground/locales/mas"
	"github.com/go-playground/locales/mas_KE"
	"github.com/go-playground/locales/mas_TZ"
	"github.com/go-playground/locales/mer"
	"github.com/go-playground/locales/mer_KE"
	"github.com/go-playground/locales/mfe"
	"github.com/go-playground/locales/mfe_MU"
	"github.com/go-playground/locales/mg"
	"github.com/go-playground/locales/mg_MG"
	"github.com/go-playground/locales/mgh"
	"github.com/go-playground/locales/mgh_MZ"
	"github.com/go-playground/locales/mgo"
	"github.com/go-playground/locales/mgo_CM"
	"github.com/go-playground/locales/mk"
	"github.com/go-playground/locales/mk_MK"
	"github.com/go-playground/locales/ml"
	"github.com/go-playground/locales/ml_IN"
	"github.com/go-playground/locales/mn"
	"github.com/go-playground/locales/mn_MN"
	"github.com/go-playground/locales/mr"
	"github.com/go-playground/locales/mr_IN"
	"github.com/go-playground/locales/ms"
	"github.com/go-playground/locales/ms_BN"
	"github.com/go-playground/locales/ms_MY"
	"github.com/go-playground/locales/ms_SG"
	"github.com/go-playground/locales/mt"
	"github.com/go-playground/locales/mt_MT"
	"github.com/go-playground/locales/mua"
	"github.com/go-playground/locales/mua_CM"
	"github.com/go-playground/locales/my"
	"github.com/go-playground/locales/my_MM"
	"github.com/go-playground/locales/mzn"
	"github.com/go-playground/locales/mzn_IR"
	"github.com/go-playground/locales/naq"
	"github.com/go-playground/locales/naq_NA"
	"github.com/go-playground/locales/nb"
	"github.com/go-playground/locales/nb_NO"
	"github.com/go-playground/locales/nb_SJ"
	"github.com/go-playground/locales/nd"
	"github.com/go-playground/locales/nd_ZW"
	"github.com/go-playground/locales/ne"
	"github.com/go-playground/locales/ne_IN"
	"github.com/go-playground/locales/ne_NP"
	"github.com/go-playground/locales/nl"
	"github.com/go-playground/locales/nl_AW"
	"github.com/go-playground/locales/nl_BE"
	"github.com/go-playground/locales/nl_BQ"
	"github.com/go-playground/locales/nl_CW"
	"github.com/go-playground/locales/nl_NL"
	"github.com/go-playground/locales/nl_SR"
	"github.com/go-playground/locales/nl_SX"
	"github.com/go-playground/locales/nmg"
	"github.com/go-playground/locales/nmg_CM"
	"github.com/go-playground/locales/nn"
	"github.com/go-playground/locales/nn_NO"
	"github.com/go-playground/locales/nnh"
	"github.com/go-playground/locales/nnh_CM"
	"github.com/go-playground/locales/nus"
	"github.com/go-playground/locales/nus_SS"
	"github.com/go-playground/locales/nyn"
	"github.com/go-playground/locales/nyn_UG"
	"github.com/go-playground/locales/om"
	"github.com/go-playground/locales/om_ET"
	"github.com/go-playground/locales/om_KE"
	"github.com/go-playground/locales/or"
	"github.com/go-playground/locales/or_IN"
	"github.com/go-playground/locales/os"
	"github.com/go-playground/locales/os_GE"
	"github.com/go-playground/locales/os_RU"
	"github.com/go-playground/locales/pa"
	"github.com/go-playground/locales/pa_Arab"
	"github.com/go-playground/locales/pa_Arab_PK"
	"github.com/go-playground/locales/pa_Guru"
	"github.com/go-playground/locales/pa_Guru_IN"
	"github.com/go-playground/locales/pl"
	"github.com/go-playground/locales/pl_PL"
	"github.com/go-playground/locales/prg"
	"github.com/go-playground/locales/prg_001"
	"github.com/go-playground/locales/ps"
	"github.com/go-playground/locales/ps_AF"
	"github.com/go-playground/locales/pt"
	"github.com/go-playground/locales/pt_AO"
	"github.com/go-playground/locales/pt_BR"
	"github.com/go-playground/locales/pt_CH"
	"github.com/go-playground/locales/pt_CV"
	"github.com/go-playground/locales/pt_GQ"
	"github.com/go-playground/locales/pt_GW"
	"github.com/go-playground/locales/pt_LU"
	"github.com/go-playground/locales/pt_MO"
	"github.com/go-playground/locales/pt_MZ"
	"github.com/go-playground/locales/pt_PT"
	"github.com/go-playground/locales/pt_ST"
	"github.com/go-playground/locales/pt_TL"
	"github.com/go-playground/locales/qu"
	"github.com/go-playground/locales/qu_BO"
	"github.com/go-playground/locales/qu_EC"
	"github.com/go-playground/locales/qu_PE"
	"github.com/go-playground/locales/rm"
	"github.com/go-playground/locales/rm_CH"
	"github.com/go-playground/locales/rn"
	"github.com/go-playground/locales/rn_BI"
	"github.com/go-playground/locales/ro"
	"github.com/go-playground/locales/ro_MD"
	"github.com/go-playground/locales/ro_RO"
	"github.com/go-playground/locales/rof"
	"github.com/go-playground/locales/rof_TZ"
	"github.com/go-playground/locales/root"
	"github.com/go-playground/locales/ru"
	"github.com/go-playground/locales/ru_BY"
	"github.com/go-playground/locales/ru_KG"
	"github.com/go-playground/locales/ru_KZ"
	"github.com/go-playground/locales/ru_MD"
	"github.com/go-playground/locales/ru_RU"
	"github.com/go-playground/locales/ru_UA"
	"github.com/go-playground/locales/rw"
	"github.com/go-playground/locales/rw_RW"
	"github.com/go-playground/locales/rwk"
	"github.com/go-playground/locales/rwk_TZ"
	"github.com/go-playground/locales/sah"
	"github.com/go-playground/locales/sah_RU"
	"github.com/go-playground/locales/saq"
	"github.com/go-playground/locales/saq_KE"
	"github.com/go-playground/locales/sbp"
	"github.com/go-playground/locales/sbp_TZ"
	"github.com/go-playground/locales/se"
	"github.com/go-playground/locales/se_FI"
	"github.com/go-playground/locales/se_NO"
	"github.com/go-playground/locales/se_SE"
	"github.com/go-playground/locales/seh"
	"github.com/go-playground/locales/seh_MZ"
	"github.com/go-playground/locales/ses"
	"github.com/go-playground/locales/ses_ML"
	"github.com/go-playground/locales/sg"
	"github.com/go-playground/locales/sg_CF"
	"github.com/go-playground/locales/shi"
	"github.com/go-playground/locales/shi_Latn"
	"github.com/go-playground/locales/shi_Latn_MA"
	"github.com/go-playground/locales/shi_Tfng"
	"github.com/go-playground/locales/shi_Tfng_MA"
	"github.com/go-playground/locales/si"
	"github.com/go-playground/locales/si_LK"
	"github.com/go-playground/locales/sk"
	"github.com/go-playground/locales/sk_SK"
	"github.com/go-playground/locales/sl"
	"github.com/go-playground/locales/sl_SI"
	"github.com/go-playground/locales/smn"
	"github.com/go-playground/locales/smn_FI"
	"github.com/go-playground/locales/sn"
	"github.com/go-playground/locales/sn_ZW"
	"github.com/go-playground/locales/so"
	"github.com/go-playground/locales/so_DJ"
	"github.com/go-playground/locales/so_ET"
	"github.com/go-playground/locales/so_KE"
	"github.com/go-playground/locales/so_SO"
	"github.com/go-playground/locales/sq"
	"github.com/go-playground/locales/sq_AL"
	"github.com/go-playground/locales/sq_MK"
	"github.com/go-playground/locales/sq_XK"
	"github.com/go-playground/locales/sr"
	"github.com/go-playground/locales/sr_Cyrl"
	"github.com/go-playground/locales/sr_Cyrl_BA"
	"github.com/go-playground/locales/sr_Cyrl_ME"
	"github.com/go-playground/locales/sr_Cyrl_RS"
	"github.com/go-playground/locales/sr_Cyrl_XK"
	"github.com/go-playground/locales/sr_Latn"
	"github.com/go-playground/locales/sr_Latn_BA"
	"github.com/go-playground/locales/sr_Latn_ME"
	"github.com/go-playground/locales/sr_Latn_RS"
	"github.com/go-playground/locales/sr_Latn_XK"
	"github.com/go-playground/locales/sv"
	"github.com/go-playground/locales/sv_AX"
	"github.com/go-playground/locales/sv_FI"
	"github.com/go-playground/locales/sv_SE"
	"github.com/go-playground/locales/sw"
	"github.com/go-playground/locales/sw_CD"
	"github.com/go-playground/locales/sw_KE"
	"github.com/go-playground/locales/sw_TZ"
	"github.com/go-playground/locales/sw_UG"
	"github.com/go-playground/locales/ta"
	"github.com/go-playground/locales/ta_IN"
	"github.com/go-playground/locales/ta_LK"
	"github.com/go-playground/locales/ta_MY"
	"github.com/go-playground/locales/ta_SG"
	"github.com/go-playground/locales/te"
	"github.com/go-playground/locales/te_IN"
	"github.com/go-playground/locales/teo"
	"github.com/go-playground/locales/teo_KE"
	"github.com/go-playground/locales/teo_UG"
	"github.com/go-playground/locales/th"
	"github.com/go-playground/locales/th_TH"
	"github.com/go-playground/locales/ti"
	"github.com/go-playground/locales/ti_ER"
	"github.com/go-playground/locales/ti_ET"
	"github.com/go-playground/locales/tk"
	"github.com/go-playground/locales/tk_TM"
	"github.com/go-playground/locales/to"
	"github.com/go-playground/locales/to_TO"
	"github.com/go-playground/locales/tr"
	"github.com/go-playground/locales/tr_CY"
	"github.com/go-playground/locales/tr_TR"
	"github.com/go-playground/locales/twq"
	"github.com/go-playground/locales/twq_NE"
	"github.com/go-playground/locales/tzm"
	"github.com/go-playground/locales/tzm_MA"
	"github.com/go-playground/locales/ug"
	"github.com/go-playground/locales/ug_CN"
	"github.com/go-playground/locales/uk"
	"github.com/go-playground/locales/uk_UA"
	"github.com/go-playground/locales/ur"
	"github.com/go-playground/locales/ur_IN"
	"github.com/go-playground/locales/ur_PK"
	"github.com/go-playground/locales/uz"
	"github.com/go-playground/locales/uz_Arab"
	"github.com/go-playground/locales/uz_Arab_AF"
	"github.com/go-playground/locales/uz_Cyrl"
	"github.com/go-playground/locales/uz_Cyrl_UZ"
	"github.com/go-playground/locales/uz_Latn"
	"github.com/go-playground/locales/uz_Latn_UZ"
	"github.com/go-playground/locales/vai"
	"github.com/go-playground/locales/vai_Latn"
	"github.com/go-playground/locales/vai_Latn_LR"
	"github.com/go-playground/locales/vai_Vaii"
	"github.com/go-playground/locales/vai_Vaii_LR"
	"github.com/go-playground/locales/vi"
	"github.com/go-playground/locales/vi_VN"
	"github.com/go-playground/locales/vo"
	"github.com/go-playground/locales/vo_001"
	"github.com/go-playground/locales/vun"
	"github.com/go-playground/locales/vun_TZ"
	"github.com/go-playground/locales/wae"
	"github.com/go-playground/locales/wae_CH"
	"github.com/go-playground/locales/xog"
	"github.com/go-playground/locales/xog_UG"
	"github.com/go-playground/locales/yav"
	"github.com/go-playground/locales/yav_CM"
	"github.com/go-playground/locales/yi"
	"github.com/go-playground/locales/yi_001"
	"github.com/go-playground/locales/yo"
	"github.com/go-playground/locales/yo_BJ"
	"github.com/go-playground/locales/yo_NG"
	"github.com/go-playground/locales/yue"
	"github.com/go-playground/locales/yue_HK"
	"github.com/go-playground/locales/zgh"
	"github.com/go-playground/locales/zgh_MA"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hans"
	"github.com/go-playground/locales/zh_Hans_CN"
	"github.com/go-playground/locales/zh_Hans_HK"
	"github.com/go-playground/locales/zh_Hans_MO"
	"github.com/go-playground/locales/zh_Hans_SG"
	"github.com/go-playground/locales/zh_Hant"
	"github.com/go-playground/locales/zh_Hant_HK"
	"github.com/go-playground/locales/zh_Hant_MO"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/locales/zu"
	"github.com/go-playground/locales/zu_ZA"
)

// LocaleFunc is the function to run in order to create
// a new instance of a given locale
type LocaleFunc func() locales.Translator

// LocaleMap is map of locale string to instance function
type LocaleMap map[string]LocaleFunc

var (
	once      sync.Once
	localeMap LocaleMap
)

func init() {
	once.Do(func() {
		localeMap = map[string]LocaleFunc{
			"rof":            rof.New,
			"ru_BY":          ru_BY.New,
			"vo_001":         vo_001.New,
			"root":           root.New,
			"cy_GB":          cy_GB.New,
			"es_CR":          es_CR.New,
			"fr_FR":          fr_FR.New,
			"gsw":            gsw.New,
			"nl":             nl.New,
			"nnh":            nnh.New,
			"sw_UG":          sw_UG.New,
			"bm":             bm.New,
			"en_PN":          en_PN.New,
			"vi":             vi.New,
			"et":             et.New,
			"gu_IN":          gu_IN.New,
			"ti_ET":          ti_ET.New,
			"en_CM":          en_CM.New,
			"es_PE":          es_PE.New,
			"fr_VU":          fr_VU.New,
			"ga_IE":          ga_IE.New,
			"guz_KE":         guz_KE.New,
			"kde_TZ":         kde_TZ.New,
			"be":             be.New,
			"en_BI":          en_BI.New,
			"dua":            dua.New,
			"en_PW":          en_PW.New,
			"en_SI":          en_SI.New,
			"hsb":            hsb.New,
			"mt":             mt.New,
			"zu_ZA":          zu_ZA.New,
			"cu":             cu.New,
			"en_MT":          en_MT.New,
			"kk":             kk.New,
			"qu_PE":          qu_PE.New,
			"en_FK":          en_FK.New,
			"jmc":            jmc.New,
			"de_BE":          de_BE.New,
			"fur":            fur.New,
			"mr_IN":          mr_IN.New,
			"pt_GW":          pt_GW.New,
			"ta_SG":          ta_SG.New,
			"ca_ES_VALENCIA": ca_ES_VALENCIA.New,
			"chr_US":         chr_US.New,
			"pt_ST":          pt_ST.New,
			"sl_SI":          sl_SI.New,
			"ta_LK":          ta_LK.New,
			"vai_Latn_LR":    vai_Latn_LR.New,
			"cgg":            cgg.New,
			"or_IN":          or_IN.New,
			"en_CX":          en_CX.New,
			"fr_KM":          fr_KM.New,
			"fr_RW":          fr_RW.New,
			"asa":            asa.New,
			"de_AT":          de_AT.New,
			"zgh":            zgh.New,
			"ar_SD":          ar_SD.New,
			"fr_DZ":          fr_DZ.New,
			"es_EA":          es_EA.New,
			"ko_KP":          ko_KP.New,
			"uz_Latn_UZ":     uz_Latn_UZ.New,
			"asa_TZ":         asa_TZ.New,
			"es_419":         es_419.New,
			"ja_JP":          ja_JP.New,
			"en_SB":          en_SB.New,
			"it":             it.New,
			"lt_LT":          lt_LT.New,
			"sah_RU":         sah_RU.New,
			"zh_Hant_HK":     zh_Hant_HK.New,
			"es_HN":          es_HN.New,
			"es_US":          es_US.New,
			"en_SD":          en_SD.New,
			"fr_GQ":          fr_GQ.New,
			"en_CA":          en_CA.New,
			"en_JE":          en_JE.New,
			"tr_CY":          tr_CY.New,
			"tzm":            tzm.New,
			"ca":             ca.New,
			"dyo":            dyo.New,
			"wae":            wae.New,
			"xog_UG":         xog_UG.New,
			"bs_Latn_BA":     bs_Latn_BA.New,
			"da":             da.New,
			"fr_SC":          fr_SC.New,
			"ro_RO":          ro_RO.New,
			"uz_Cyrl_UZ":     uz_Cyrl_UZ.New,
			"en_NL":          en_NL.New,
			"en_TT":          en_TT.New,
			"nl_SX":          nl_SX.New,
			"ar_DZ":          ar_DZ.New,
			"mgo_CM":         mgo_CM.New,
			"es_CO":          es_CO.New,
			"fa":             fa.New,
			"mzn":            mzn.New,
			"sg":             sg.New,
			"ug":             ug.New,
			"yav":            yav.New,
			"ar_DJ":          ar_DJ.New,
			"be_BY":          be_BY.New,
			"en_PG":          en_PG.New,
			"is":             is.New,
			"tr":             tr.New,
			"en_150":         en_150.New,
			"en_BB":          en_BB.New,
			"en_AU":          en_AU.New,
			"en_SE":          en_SE.New,
			"to_TO":          to_TO.New,
			"de_DE":          de_DE.New,
			"dje_NE":         dje_NE.New,
			"fo_FO":          fo_FO.New,
			"ksb_TZ":         ksb_TZ.New,
			"es_ES":          es_ES.New,
			"saq_KE":         saq_KE.New,
			"ak":             ak.New,
			"en_NF":          en_NF.New,
			"bg_BG":          bg_BG.New,
			"en_MO":          en_MO.New,
			"ar_YE":          ar_YE.New,
			"ast_ES":         ast_ES.New,
			"fil_PH":         fil_PH.New,
			"ln_CG":          ln_CG.New,
			"se_FI":          se_FI.New,
			"smn":            smn.New,
			"vai_Latn":       vai_Latn.New,
			"ar_TD":          ar_TD.New,
			"en_GU":          en_GU.New,
			"smn_FI":         smn_FI.New,
			"zh_Hans_HK":     zh_Hans_HK.New,
			"en_IL":          en_IL.New,
			"id":             id.New,
			"ki":             ki.New,
			"ky_KG":          ky_KG.New,
			"so_ET":          so_ET.New,
			"sq_MK":          sq_MK.New,
			"uz_Arab_AF":     uz_Arab_AF.New,
			"uz_Cyrl":        uz_Cyrl.New,
			"en_MS":          en_MS.New,
			"en_VI":          en_VI.New,
			"mn_MN":          mn_MN.New,
			"os_RU":          os_RU.New,
			"ar_IL":          ar_IL.New,
			"ee_TG":          ee_TG.New,
			"is_IS":          is_IS.New,
			"nb_SJ":          nb_SJ.New,
			"twq_NE":         twq_NE.New,
			"chr":            chr.New,
			"ii":             ii.New,
			"cu_RU":          cu_RU.New,
			"en_PK":          en_PK.New,
			"ewo_CM":         ewo_CM.New,
			"rm":             rm.New,
			"ar_SS":          ar_SS.New,
			"bo":             bo.New,
			"fr_MR":          fr_MR.New,
			"he":             he.New,
			"nnh_CM":         nnh_CM.New,
			"sr_Cyrl_ME":     sr_Cyrl_ME.New,
			"yi_001":         yi_001.New,
			"ast":            ast.New,
			"de":             de.New,
			"fur_IT":         fur_IT.New,
			"uk_UA":          uk_UA.New,
			"ar_EH":          ar_EH.New,
			"en_US_POSIX":    en_US_POSIX.New,
			"lg":             lg.New,
			"sr_Latn_ME":     sr_Latn_ME.New,
			"os_GE":          os_GE.New,
			"pl_PL":          pl_PL.New,
			"sr_Cyrl_BA":     sr_Cyrl_BA.New,
			"fr_BL":          fr_BL.New,
			"kl":             kl.New,
			"ff_MR":          ff_MR.New,
			"pt_CH":          pt_CH.New,
			"sl":             sl.New,
			"vai":            vai.New,
			"en_TV":          en_TV.New,
			"es_PA":          es_PA.New,
			"pa_Guru":        pa_Guru.New,
			"rof_TZ":         rof_TZ.New,
			"se":             se.New,
			"seh_MZ":         seh_MZ.New,
			"sw_TZ":          sw_TZ.New,
			"zh_Hant_MO":     zh_Hant_MO.New,
			"af_NA":          af_NA.New,
			"naq":            naq.New,
			"fr_CG":          fr_CG.New,
			"ka":             ka.New,
			"kab_DZ":         kab_DZ.New,
			"kea_CV":         kea_CV.New,
			"lt":             lt.New,
			"lv_LV":          lv_LV.New,
			"bem":            bem.New,
			"dav_KE":         dav_KE.New,
			"mfe_MU":         mfe_MU.New,
			"vun_TZ":         vun_TZ.New,
			"yo_BJ":          yo_BJ.New,
			"zgh_MA":         zgh_MA.New,
			"cs":             cs.New,
			"fr_SY":          fr_SY.New,
			"ms_SG":          ms_SG.New,
			"sw":             sw.New,
			"teo_KE":         teo_KE.New,
			"en_WS":          en_WS.New,
			"ml":             ml.New,
			"ksf":            ksf.New,
			"shi_Latn_MA":    shi_Latn_MA.New,
			"ar_SY":          ar_SY.New,
			"en_KE":          en_KE.New,
			"fa_AF":          fa_AF.New,
			"fr_GP":          fr_GP.New,
			"khq_ML":         khq_ML.New,
			"kok_IN":         kok_IN.New,
			"ta":             ta.New,
			"vo":             vo.New,
			"en_MY":          en_MY.New,
			"en_NZ":          en_NZ.New,
			"fr_CI":          fr_CI.New,
			"ha_NE":          ha_NE.New,
			"ii_CN":          ii_CN.New,
			"sr_Latn_BA":     sr_Latn_BA.New,
			"yue":            yue.New,
			"en_RW":          en_RW.New,
			"fi":             fi.New,
			"lag":            lag.New,
			"bs_Latn":        bs_Latn.New,
			"en_TZ":          en_TZ.New,
			"fr_RE":          fr_RE.New,
			"fr_TN":          fr_TN.New,
			"kln_KE":         kln_KE.New,
			"ta_IN":          ta_IN.New,
			"zh":             zh.New,
			"agq_CM":         agq_CM.New,
			"es_GQ":          es_GQ.New,
			"id_ID":          id_ID.New,
			"kok":            kok.New,
			"ps":             ps.New,
			"vun":            vun.New,
			"dua_CM":         dua_CM.New,
			"en_FI":          en_FI.New,
			"luy":            luy.New,
			"te":             te.New,
			"xog":            xog.New,
			"ce":             ce.New,
			"kw":             kw.New,
			"en_ER":          en_ER.New,
			"en_SG":          en_SG.New,
			"en_SS":          en_SS.New,
			"eo_001":         eo_001.New,
			"fr_GN":          fr_GN.New,
			"fr_MQ":          fr_MQ.New,
			"ar_PS":          ar_PS.New,
			"bn_IN":          bn_IN.New,
			"mk":             mk.New,
			"pl":             pl.New,
			"nd":             nd.New,
			"ses_ML":         ses_ML.New,
			"ckb_IR":         ckb_IR.New,
			"fr":             fr.New,
			"fr_NC":          fr_NC.New,
			"hu":             hu.New,
			"om_KE":          om_KE.New,
			"sr_Latn_RS":     sr_Latn_RS.New,
			"ar_KM":          ar_KM.New,
			"ar_TN":          ar_TN.New,
			"km":             km.New,
			"kw_GB":          kw_GB.New,
			"luy_KE":         luy_KE.New,
			"ug_CN":          ug_CN.New,
			"he_IL":          he_IL.New,
			"kl_GL":          kl_GL.New,
			"en_GB":          en_GB.New,
			"pa_Guru_IN":     pa_Guru_IN.New,
			"gv_IM":          gv_IM.New,
			"ne_NP":          ne_NP.New,
			"ti":             ti.New,
			"cs_CZ":          cs_CZ.New,
			"en_UG":          en_UG.New,
			"hu_HU":          hu_HU.New,
			"mgo":            mgo.New,
			"ru_UA":          ru_UA.New,
			"de_LU":          de_LU.New,
			"en_ZM":          en_ZM.New,
			"fr_MC":          fr_MC.New,
			"lu":             lu.New,
			"sr_Cyrl_XK":     sr_Cyrl_XK.New,
			"zh_Hans_MO":     zh_Hans_MO.New,
			"en_BE":          en_BE.New,
			"fo_DK":          fo_DK.New,
			"mua_CM":         mua_CM.New,
			"ru_RU":          ru_RU.New,
			"saq":            saq.New,
			"es_GT":          es_GT.New,
			"eu":             eu.New,
			"fr_NE":          fr_NE.New,
			"gsw_LI":         gsw_LI.New,
			"kam_KE":         kam_KE.New,
			"sv_SE":          sv_SE.New,
			"da_GL":          da_GL.New,
			"en_BZ":          en_BZ.New,
			"en_NA":          en_NA.New,
			"es_EC":          es_EC.New,
			"mas":            mas.New,
			"bg":             bg.New,
			"ckb":            ckb.New,
			"en_NR":          en_NR.New,
			"lo":             lo.New,
			"mgh_MZ":         mgh_MZ.New,
			"mt_MT":          mt_MT.New,
			"nmg_CM":         nmg_CM.New,
			"ta_MY":          ta_MY.New,
			"ar_SO":          ar_SO.New,
			"en_CC":          en_CC.New,
			"yi":             yi.New,
			"en_MW":          en_MW.New,
			"lg_UG":          lg_UG.New,
			"mg_MG":          mg_MG.New,
			"nn_NO":          nn_NO.New,
			"ru_KZ":          ru_KZ.New,
			"az":             az.New,
			"bs_Cyrl_BA":     bs_Cyrl_BA.New,
			"fr_ML":          fr_ML.New,
			"hr":             hr.New,
			"mer":            mer.New,
			"prg_001":        prg_001.New,
			"bo_IN":          bo_IN.New,
			"en_SX":          en_SX.New,
			"zh_Hans":        zh_Hans.New,
			"ar_QA":          ar_QA.New,
			"mr":             mr.New,
			"pt_MZ":          pt_MZ.New,
			"guz":            guz.New,
			"ha_NG":          ha_NG.New,
			"ks_IN":          ks_IN.New,
			"zh_Hans_SG":     zh_Hans_SG.New,
			"en_GG":          en_GG.New,
			"en_UM":          en_UM.New,
			"hsb_DE":         hsb_DE.New,
			"ckb_IQ":         ckb_IQ.New,
			"gd":             gd.New,
			"en_SZ":          en_SZ.New,
			"fy_NL":          fy_NL.New,
			"mgh":            mgh.New,
			"pt_AO":          pt_AO.New,
			"si_LK":          si_LK.New,
			"en_LC":          en_LC.New,
			"en_MU":          en_MU.New,
			"mas_TZ":         mas_TZ.New,
			"mg":             mg.New,
			"teo":            teo.New,
			"fr_DJ":          fr_DJ.New,
			"lrc":            lrc.New,
			"or":             or.New,
			"fr_MF":          fr_MF.New,
			"nl_SR":          nl_SR.New,
			"fr_CD":          fr_CD.New,
			"gl":             gl.New,
			"ks":             ks.New,
			"nl_NL":          nl_NL.New,
			"ru":             ru.New,
			"sw_KE":          sw_KE.New,
			"en_PH":          en_PH.New,
			"fr_BF":          fr_BF.New,
			"te_IN":          te_IN.New,
			"tr_TR":          tr_TR.New,
			"br":             br.New,
			"en_GH":          en_GH.New,
			"ml_IN":          ml_IN.New,
			"nb_NO":          nb_NO.New,
			"nus_SS":         nus_SS.New,
			"so_SO":          so_SO.New,
			"am_ET":          am_ET.New,
			"az_Cyrl":        az_Cyrl.New,
			"sr_Latn":        sr_Latn.New,
			"en_CY":          en_CY.New,
			"lrc_IQ":         lrc_IQ.New,
			"mua":            mua.New,
			"ro_MD":          ro_MD.New,
			"so_KE":          so_KE.New,
			"cy":             cy.New,
			"en_BW":          en_BW.New,
			"en_KN":          en_KN.New,
			"en_US":          en_US.New,
			"es_UY":          es_UY.New,
			"gsw_CH":         gsw_CH.New,
			"nl_BQ":          nl_BQ.New,
			"rw_RW":          rw_RW.New,
			"en_GD":          en_GD.New,
			"en_IM":          en_IM.New,
			"hy":             hy.New,
			"pt_BR":          pt_BR.New,
			"ar_EG":          ar_EG.New,
			"en_DE":          en_DE.New,
			"ar":             ar.New,
			"en_SL":          en_SL.New,
			"es_DO":          es_DO.New,
			"eu_ES":          eu_ES.New,
			"fr_GA":          fr_GA.New,
			"hi_IN":          hi_IN.New,
			"it_IT":          it_IT.New,
			"nd_ZW":          nd_ZW.New,
			"dje":            dje.New,
			"en_NG":          en_NG.New,
			"rn":             rn.New,
			"uz_Latn":        uz_Latn.New,
			"fil":            fil.New,
			"sr_Latn_XK":     sr_Latn_XK.New,
			"ur_IN":          ur_IN.New,
			"zh_Hant_TW":     zh_Hant_TW.New,
			"en_AT":          en_AT.New,
			"es_PY":          es_PY.New,
			"cgg_UG":         cgg_UG.New,
			"el_CY":          el_CY.New,
			"en_BS":          en_BS.New,
			"en_CK":          en_CK.New,
			"en_KI":          en_KI.New,
			"mfe":            mfe.New,
			"ar_LB":          ar_LB.New,
			"bas_CM":         bas_CM.New,
			"ne_IN":          ne_IN.New,
			"wae_CH":         wae_CH.New,
			"ur":             ur.New,
			"haw_US":         haw_US.New,
			"ko_KR":          ko_KR.New,
			"ms_BN":          ms_BN.New,
			"pa":             pa.New,
			"en_VC":          en_VC.New,
			"es_NI":          es_NI.New,
			"en_DG":          en_DG.New,
			"en_TO":          en_TO.New,
			"ja":             ja.New,
			"jgo_CM":         jgo_CM.New,
			"ka_GE":          ka_GE.New,
			"kkj":            kkj.New,
			"bn":             bn.New,
			"en_AS":          en_AS.New,
			"ln_CD":          ln_CD.New,
			"vi_VN":          vi_VN.New,
			"en_ZA":          en_ZA.New,
			"fr_YT":          fr_YT.New,
			"it_CH":          it_CH.New,
			"sq_XK":          sq_XK.New,
			"yav_CM":         yav_CM.New,
			"ar_OM":          ar_OM.New,
			"dav":            dav.New,
			"en_HK":          en_HK.New,
			"fr_CM":          fr_CM.New,
			"ksh_DE":         ksh_DE.New,
			"pt":             pt.New,
			"as":             as.New,
			"dz_BT":          dz_BT.New,
			"fa_IR":          fa_IR.New,
			"ff_SN":          ff_SN.New,
			"fr_BJ":          fr_BJ.New,
			"kea":            kea.New,
			"vai_Vaii_LR":    vai_Vaii_LR.New,
			"ar_JO":          ar_JO.New,
			"bm_ML":          bm_ML.New,
			"en_AG":          en_AG.New,
			"en_IE":          en_IE.New,
			"ff_CM":          ff_CM.New,
			"kln":            kln.New,
			"ky":             ky.New,
			"ru_KG":          ru_KG.New,
			"as_IN":          as_IN.New,
			"bez":            bez.New,
			"sk":             sk.New,
			"es_CL":          es_CL.New,
			"pa_Arab_PK":     pa_Arab_PK.New,
			"el_GR":          el_GR.New,
			"en_CH":          en_CH.New,
			"ko":             ko.New,
			"ksh":            ksh.New,
			"shi":            shi.New,
			"sq":             sq.New,
			"ca_AD":          ca_AD.New,
			"fr_CF":          fr_CF.New,
			"hr_BA":          hr_BA.New,
			"lo_LA":          lo_LA.New,
			"mzn_IR":         mzn_IR.New,
			"os":             os.New,
			"si":             si.New,
			"tzm_MA":         tzm_MA.New,
			"ca_FR":          ca_FR.New,
			"el":             el.New,
			"en_TK":          en_TK.New,
			"en_VG":          en_VG.New,
			"fr_WF":          fr_WF.New,
			"ga":             ga.New,
			"rwk_TZ":         rwk_TZ.New,
			"uz":             uz.New,
			"ar_ER":          ar_ER.New,
			"ar_LY":          ar_LY.New,
			"ln_AO":          ln_AO.New,
			"shi_Tfng":       shi_Tfng.New,
			"yo":             yo.New,
			"en_DK":          en_DK.New,
			"ki_KE":          ki_KE.New,
			"en_MP":          en_MP.New,
			"yue_HK":         yue_HK.New,
			"ar_AE":          ar_AE.New,
			"ee_GH":          ee_GH.New,
			"fr_LU":          fr_LU.New,
			"ar_IQ":          ar_IQ.New,
			"ar_MR":          ar_MR.New,
			"lb_LU":          lb_LU.New,
			"mk_MK":          mk_MK.New,
			"th_TH":          th_TH.New,
			"tk_TM":          tk_TM.New,
			"yo_NG":          yo_NG.New,
			"gsw_FR":         gsw_FR.New,
			"khq":            khq.New,
			"de_CH":          de_CH.New,
			"en_VU":          en_VU.New,
			"en_ZW":          en_ZW.New,
			"es":             es.New,
			"fr_BE":          fr_BE.New,
			"gd_GB":          gd_GB.New,
			"ar_001":         ar_001.New,
			"ca_IT":          ca_IT.New,
			"sv_FI":          sv_FI.New,
			"kde":            kde.New,
			"to":             to.New,
			"twq":            twq.New,
			"bo_CN":          bo_CN.New,
			"da_DK":          da_DK.New,
			"es_CU":          es_CU.New,
			"sr_Cyrl_RS":     sr_Cyrl_RS.New,
			"az_Latn_AZ":     az_Latn_AZ.New,
			"en_IN":          en_IN.New,
			"nl_AW":          nl_AW.New,
			"sg_CF":          sg_CF.New,
			"es_PR":          es_PR.New,
			"fr_MG":          fr_MG.New,
			"kk_KZ":          kk_KZ.New,
			"my":             my.New,
			"my_MM":          my_MM.New,
			"nus":            nus.New,
			"ebu_KE":         ebu_KE.New,
			"es_BR":          es_BR.New,
			"en_BM":          en_BM.New,
			"kab":            kab.New,
			"en":             en.New,
			"en_LS":          en_LS.New,
			"es_IC":          es_IC.New,
			"fr_TD":          fr_TD.New,
			"haw":            haw.New,
			"lag_TZ":         lag_TZ.New,
			"bs_Cyrl":        bs_Cyrl.New,
			"de_LI":          de_LI.New,
			"shi_Tfng_MA":    shi_Tfng_MA.New,
			"lkt_US":         lkt_US.New,
			"pt_CV":          pt_CV.New,
			"gv":             gv.New,
			"om":             om.New,
			"rn_BI":          rn_BI.New,
			"ru_MD":          ru_MD.New,
			"rwk":            rwk.New,
			"en_GM":          en_GM.New,
			"fr_GF":          fr_GF.New,
			"es_PH":          es_PH.New,
			"fo":             fo.New,
			"fr_MU":          fr_MU.New,
			"ha_GH":          ha_GH.New,
			"dsb":            dsb.New,
			"en_PR":          en_PR.New,
			"mn":             mn.New,
			"nb":             nb.New,
			"pt_PT":          pt_PT.New,
			"qu_BO":          qu_BO.New,
			"uk":             uk.New,
			"ak_GH":          ak_GH.New,
			"az_Cyrl_AZ":     az_Cyrl_AZ.New,
			"jmc_TZ":         jmc_TZ.New,
			"ksb":            ksb.New,
			"sq_AL":          sq_AL.New,
			"es_MX":          es_MX.New,
			"fr_TG":          fr_TG.New,
			"fr_MA":          fr_MA.New,
			"nmg":            nmg.New,
			"nyn":            nyn.New,
			"sw_CD":          sw_CD.New,
			"uz_Arab":        uz_Arab.New,
			"zu":             zu.New,
			"en_GY":          en_GY.New,
			"es_BO":          es_BO.New,
			"es_VE":          es_VE.New,
			"mas_KE":         mas_KE.New,
			"mer_KE":         mer_KE.New,
			"naq_NA":         naq_NA.New,
			"pt_TL":          pt_TL.New,
			"teo_UG":         teo_UG.New,
			"bem_ZM":         bem_ZM.New,
			"en_MH":          en_MH.New,
			"sk_SK":          sk_SK.New,
			"th":             th.New,
			"bs":             bs.New,
			"seh":            seh.New,
			"bas":            bas.New,
			"nl_CW":          nl_CW.New,
			"sbp_TZ":         sbp_TZ.New,
			"ln_CF":          ln_CF.New,
			"ms_MY":          ms_MY.New,
			"se_NO":          se_NO.New,
			"fr_SN":          fr_SN.New,
			"hy_AM":          hy_AM.New,
			"fr_BI":          fr_BI.New,
			"fr_CH":          fr_CH.New,
			"en_SH":          en_SH.New,
			"et_EE":          et_EE.New,
			"lkt":            lkt.New,
			"nyn_UG":         nyn_UG.New,
			"sah":            sah.New,
			"vai_Vaii":       vai_Vaii.New,
			"af_ZA":          af_ZA.New,
			"ca_ES":          ca_ES.New,
			"fr_CA":          fr_CA.New,
			"fy":             fy.New,
			"ln":             ln.New,
			"nn":             nn.New,
			"sn":             sn.New,
			"agq":            agq.New,
			"ff":             ff.New,
			"kn":             kn.New,
			"rm_CH":          rm_CH.New,
			"se_SE":          se_SE.New,
			"sr_Cyrl":        sr_Cyrl.New,
			"ti_ER":          ti_ER.New,
			"zh_Hant":        zh_Hant.New,
			"dyo_SN":         dyo_SN.New,
			"hi":             hi.New,
			"en_JM":          en_JM.New,
			"fr_PF":          fr_PF.New,
			"ig_NG":          ig_NG.New,
			"ses":            ses.New,
			"ar_KW":          ar_KW.New,
			"ce_RU":          ce_RU.New,
			"ewo":            ewo.New,
			"ff_GN":          ff_GN.New,
			"ha":             ha.New,
			"ps_AF":          ps_AF.New,
			"sbp":            sbp.New,
			"sn_ZW":          sn_ZW.New,
			"en_GI":          en_GI.New,
			"en_TC":          en_TC.New,
			"prg":            prg.New,
			"so_DJ":          so_DJ.New,
			"sr":             sr.New,
			"ur_PK":          ur_PK.New,
			"brx":            brx.New,
			"en_SC":          en_SC.New,
			"en_AI":          en_AI.New,
			"en_KY":          en_KY.New,
			"es_AR":          es_AR.New,
			"fr_HT":          fr_HT.New,
			"jgo":            jgo.New,
			"pt_GQ":          pt_GQ.New,
			"am":             am.New,
			"br_FR":          br_FR.New,
			"sv_AX":          sv_AX.New,
			"ee":             ee.New,
			"en_DM":          en_DM.New,
			"en_LR":          en_LR.New,
			"ksf_CM":         ksf_CM.New,
			"lv":             lv.New,
			"af":             af.New,
			"brx_IN":         brx_IN.New,
			"om_ET":          om_ET.New,
			"so":             so.New,
			"gu":             gu.New,
			"kam":            kam.New,
			"en_FJ":          en_FJ.New,
			"en_NU":          en_NU.New,
			"luo_KE":         luo_KE.New,
			"ms":             ms.New,
			"shi_Latn":       shi_Latn.New,
			"ar_SA":          ar_SA.New,
			"dsb_DE":         dsb_DE.New,
			"fi_FI":          fi_FI.New,
			"ig":             ig.New,
			"ne":             ne.New,
			"ar_MA":          ar_MA.New,
			"bn_BD":          bn_BD.New,
			"dz":             dz.New,
			"gl_ES":          gl_ES.New,
			"kn_IN":          kn_IN.New,
			"pa_Arab":        pa_Arab.New,
			"qu":             qu.New,
			"rw":             rw.New,
			"az_Latn":        az_Latn.New,
			"bez_TZ":         bez_TZ.New,
			"tk":             tk.New,
			"zh_Hans_CN":     zh_Hans_CN.New,
			"pt_MO":          pt_MO.New,
			"qu_EC":          qu_EC.New,
			"ar_BH":          ar_BH.New,
			"en_FM":          en_FM.New,
			"hr_HR":          hr_HR.New,
			"kkj_CM":         kkj_CM.New,
			"pt_LU":          pt_LU.New,
			"es_SV":          es_SV.New,
			"fr_PM":          fr_PM.New,
			"luo":            luo.New,
			"nl_BE":          nl_BE.New,
			"ro":             ro.New,
			"sv":             sv.New,
			"lrc_IR":         lrc_IR.New,
			"lu_CD":          lu_CD.New,
			"en_MG":          en_MG.New,
			"it_SM":          it_SM.New,
			"km_KH":          km_KH.New,
			"ebu":            ebu.New,
			"en_IO":          en_IO.New,
			"lb":             lb.New,
			"en_001":         en_001.New,
			"eo":             eo.New,
		}
	})
}

// Map returns the map of locales to instance New function
func Map() LocaleMap {
	return localeMap
}
