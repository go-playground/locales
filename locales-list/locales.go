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
			"ar_LY":          ar_LY.New,
			"fr_TD":          fr_TD.New,
			"ar_JO":          ar_JO.New,
			"en_FK":          en_FK.New,
			"es_419":         es_419.New,
			"hr":             hr.New,
			"si_LK":          si_LK.New,
			"sl_SI":          sl_SI.New,
			"fr_DZ":          fr_DZ.New,
			"ky":             ky.New,
			"saq":            saq.New,
			"dz_BT":          dz_BT.New,
			"fr_VU":          fr_VU.New,
			"sv_FI":          sv_FI.New,
			"brx":            brx.New,
			"el_CY":          el_CY.New,
			"en_MU":          en_MU.New,
			"es_BR":          es_BR.New,
			"pa_Arab_PK":     pa_Arab_PK.New,
			"pl_PL":          pl_PL.New,
			"az_Cyrl_AZ":     az_Cyrl_AZ.New,
			"be_BY":          be_BY.New,
			"en":             en.New,
			"en_SH":          en_SH.New,
			"fr_GN":          fr_GN.New,
			"ko_KP":          ko_KP.New,
			"mfe":            mfe.New,
			"nmg_CM":         nmg_CM.New,
			"dav":            dav.New,
			"it_SM":          it_SM.New,
			"kab":            kab.New,
			"sk_SK":          sk_SK.New,
			"sr_Cyrl_RS":     sr_Cyrl_RS.New,
			"bem":            bem.New,
			"dsb":            dsb.New,
			"en_BB":          en_BB.New,
			"ha_GH":          ha_GH.New,
			"ms_SG":          ms_SG.New,
			"pt_PT":          pt_PT.New,
			"uz_Cyrl_UZ":     uz_Cyrl_UZ.New,
			"agq":            agq.New,
			"ar_SY":          ar_SY.New,
			"dua_CM":         dua_CM.New,
			"en_IN":          en_IN.New,
			"et_EE":          et_EE.New,
			"kw_GB":          kw_GB.New,
			"mk_MK":          mk_MK.New,
			"mn":             mn.New,
			"ne_NP":          ne_NP.New,
			"tr_TR":          tr_TR.New,
			"wae":            wae.New,
			"ewo":            ewo.New,
			"nl_AW":          nl_AW.New,
			"qu_EC":          qu_EC.New,
			"br_FR":          br_FR.New,
			"fr_GF":          fr_GF.New,
			"gd_GB":          gd_GB.New,
			"hi":             hi.New,
			"ar_LB":          ar_LB.New,
			"en_GD":          en_GD.New,
			"nyn_UG":         nyn_UG.New,
			"sr_Latn_BA":     sr_Latn_BA.New,
			"twq":            twq.New,
			"yav_CM":         yav_CM.New,
			"af":             af.New,
			"mua":            mua.New,
			"nd":             nd.New,
			"ro_MD":          ro_MD.New,
			"en_SZ":          en_SZ.New,
			"fr_SN":          fr_SN.New,
			"seh_MZ":         seh_MZ.New,
			"en_VI":          en_VI.New,
			"fr_BI":          fr_BI.New,
			"fr_CG":          fr_CG.New,
			"gsw_CH":         gsw_CH.New,
			"is_IS":          is_IS.New,
			"pt_ST":          pt_ST.New,
			"sl":             sl.New,
			"bn_IN":          bn_IN.New,
			"el":             el.New,
			"en_ZA":          en_ZA.New,
			"fo_FO":          fo_FO.New,
			"fr_CA":          fr_CA.New,
			"kok":            kok.New,
			"lt":             lt.New,
			"ur":             ur.New,
			"xog_UG":         xog_UG.New,
			"da_GL":          da_GL.New,
			"hsb_DE":         hsb_DE.New,
			"ml_IN":          ml_IN.New,
			"ar_PS":          ar_PS.New,
			"en_SB":          en_SB.New,
			"es_EC":          es_EC.New,
			"es_VE":          es_VE.New,
			"kab_DZ":         kab_DZ.New,
			"ksb_TZ":         ksb_TZ.New,
			"my_MM":          my_MM.New,
			"ru_RU":          ru_RU.New,
			"so_ET":          so_ET.New,
			"zh_Hans_CN":     zh_Hans_CN.New,
			"en_NZ":          en_NZ.New,
			"mt":             mt.New,
			"pa_Guru":        pa_Guru.New,
			"shi_Latn":       shi_Latn.New,
			"rof_TZ":         rof_TZ.New,
			"sv_AX":          sv_AX.New,
			"tk":             tk.New,
			"eu":             eu.New,
			"gsw_FR":         gsw_FR.New,
			"pt_AO":          pt_AO.New,
			"sah_RU":         sah_RU.New,
			"sbp_TZ":         sbp_TZ.New,
			"uk":             uk.New,
			"wae_CH":         wae_CH.New,
			"ast_ES":         ast_ES.New,
			"fr_MA":          fr_MA.New,
			"gd":             gd.New,
			"my":             my.New,
			"ar_TD":          ar_TD.New,
			"en_MP":          en_MP.New,
			"fr_ML":          fr_ML.New,
			"gu":             gu.New,
			"ka_GE":          ka_GE.New,
			"vai_Vaii_LR":    vai_Vaii_LR.New,
			"yo_BJ":          yo_BJ.New,
			"lrc":            lrc.New,
			"ses_ML":         ses_ML.New,
			"sw_TZ":          sw_TZ.New,
			"de_AT":          de_AT.New,
			"en_AT":          en_AT.New,
			"fr_WF":          fr_WF.New,
			"teo":            teo.New,
			"ug":             ug.New,
			"ar_MA":          ar_MA.New,
			"es_PE":          es_PE.New,
			"ff":             ff.New,
			"kok_IN":         kok_IN.New,
			"sbp":            sbp.New,
			"se_FI":          se_FI.New,
			"ses":            ses.New,
			"da_DK":          da_DK.New,
			"en_KE":          en_KE.New,
			"en_RW":          en_RW.New,
			"ca_ES":          ca_ES.New,
			"et":             et.New,
			"fo_DK":          fo_DK.New,
			"vi":             vi.New,
			"ar_SS":          ar_SS.New,
			"en_FJ":          en_FJ.New,
			"en_WS":          en_WS.New,
			"eu_ES":          eu_ES.New,
			"mas_KE":         mas_KE.New,
			"th_TH":          th_TH.New,
			"zh_Hans":        zh_Hans.New,
			"en_GH":          en_GH.New,
			"en_NG":          en_NG.New,
			"es_BO":          es_BO.New,
			"gl":             gl.New,
			"ig_NG":          ig_NG.New,
			"ta_MY":          ta_MY.New,
			"ckb_IQ":         ckb_IQ.New,
			"dsb_DE":         dsb_DE.New,
			"en_BZ":          en_BZ.New,
			"en_ER":          en_ER.New,
			"gv":             gv.New,
			"ar_AE":          ar_AE.New,
			"bs_Cyrl":        bs_Cyrl.New,
			"zh_Hans_MO":     zh_Hans_MO.New,
			"as":             as.New,
			"en_DK":          en_DK.New,
			"en_VC":          en_VC.New,
			"es_PA":          es_PA.New,
			"es_UY":          es_UY.New,
			"ro":             ro.New,
			"ti":             ti.New,
			"fr_MU":          fr_MU.New,
			"uz_Cyrl":        uz_Cyrl.New,
			"en_AG":          en_AG.New,
			"en_BS":          en_BS.New,
			"en_SG":          en_SG.New,
			"fr_BJ":          fr_BJ.New,
			"ga":             ga.New,
			"hy":             hy.New,
			"sw_CD":          sw_CD.New,
			"tr":             tr.New,
			"fr_PF":          fr_PF.New,
			"ki_KE":          ki_KE.New,
			"lu_CD":          lu_CD.New,
			"ml":             ml.New,
			"teo_UG":         teo_UG.New,
			"fr_CM":          fr_CM.New,
			"khq":            khq.New,
			"mzn":            mzn.New,
			"pt":             pt.New,
			"shi_Tfng_MA":    shi_Tfng_MA.New,
			"smn_FI":         smn_FI.New,
			"teo_KE":         teo_KE.New,
			"en_AI":          en_AI.New,
			"en_ZM":          en_ZM.New,
			"es_DO":          es_DO.New,
			"es_US":          es_US.New,
			"sr_Latn_ME":     sr_Latn_ME.New,
			"tr_CY":          tr_CY.New,
			"yue":            yue.New,
			"az_Cyrl":        az_Cyrl.New,
			"bo_IN":          bo_IN.New,
			"en_TT":          en_TT.New,
			"mk":             mk.New,
			"sq_XK":          sq_XK.New,
			"ak_GH":          ak_GH.New,
			"fr_FR":          fr_FR.New,
			"fr_MG":          fr_MG.New,
			"is":             is.New,
			"it_CH":          it_CH.New,
			"kde_TZ":         kde_TZ.New,
			"cy":             cy.New,
			"so_KE":          so_KE.New,
			"cgg_UG":         cgg_UG.New,
			"fur":            fur.New,
			"kk":             kk.New,
			"pt_MZ":          pt_MZ.New,
			"sr_Latn_XK":     sr_Latn_XK.New,
			"en_TK":          en_TK.New,
			"es":             es.New,
			"fr_CF":          fr_CF.New,
			"fr_YT":          fr_YT.New,
			"te_IN":          te_IN.New,
			"twq_NE":         twq_NE.New,
			"ar_SD":          ar_SD.New,
			"es_GQ":          es_GQ.New,
			"ff_MR":          ff_MR.New,
			"fur_IT":         fur_IT.New,
			"hu_HU":          hu_HU.New,
			"ii":             ii.New,
			"lag":            lag.New,
			"shi":            shi.New,
			"vo":             vo.New,
			"en_MO":          en_MO.New,
			"es_EA":          es_EA.New,
			"luy_KE":         luy_KE.New,
			"sq_MK":          sq_MK.New,
			"ckb_IR":         ckb_IR.New,
			"en_LR":          en_LR.New,
			"gv_IM":          gv_IM.New,
			"mr":             mr.New,
			"tzm":            tzm.New,
			"ar":             ar.New,
			"chr_US":         chr_US.New,
			"ee":             ee.New,
			"en_BM":          en_BM.New,
			"en_MG":          en_MG.New,
			"mg_MG":          mg_MG.New,
			"mgh_MZ":         mgh_MZ.New,
			"prg_001":        prg_001.New,
			"ak":             ak.New,
			"ar_KW":          ar_KW.New,
			"ebu":            ebu.New,
			"naq":            naq.New,
			"az_Latn":        az_Latn.New,
			"en_BI":          en_BI.New,
			"es_CL":          es_CL.New,
			"fr_BE":          fr_BE.New,
			"fr_HT":          fr_HT.New,
			"fr_SC":          fr_SC.New,
			"lg_UG":          lg_UG.New,
			"shi_Tfng":       shi_Tfng.New,
			"te":             te.New,
			"ar_EH":          ar_EH.New,
			"bs_Latn_BA":     bs_Latn_BA.New,
			"ka":             ka.New,
			"kkj":            kkj.New,
			"ln_CD":          ln_CD.New,
			"nl_SR":          nl_SR.New,
			"os_RU":          os_RU.New,
			"pt_TL":          pt_TL.New,
			"ca_AD":          ca_AD.New,
			"dje":            dje.New,
			"en_MH":          en_MH.New,
			"fi_FI":          fi_FI.New,
			"fr":             fr.New,
			"lo":             lo.New,
			"mas_TZ":         mas_TZ.New,
			"ne_IN":          ne_IN.New,
			"yo":             yo.New,
			"cgg":            cgg.New,
			"dz":             dz.New,
			"en_FM":          en_FM.New,
			"fr_RE":          fr_RE.New,
			"ne":             ne.New,
			"rof":            rof.New,
			"ru_KZ":          ru_KZ.New,
			"uz_Arab_AF":     uz_Arab_AF.New,
			"cs":             cs.New,
			"ebu_KE":         ebu_KE.New,
			"ee_TG":          ee_TG.New,
			"gsw_LI":         gsw_LI.New,
			"lb_LU":          lb_LU.New,
			"ru_BY":          ru_BY.New,
			"vai":            vai.New,
			"ar_OM":          ar_OM.New,
			"nb_SJ":          nb_SJ.New,
			"nnh":            nnh.New,
			"vai_Latn":       vai_Latn.New,
			"ar_DZ":          ar_DZ.New,
			"en_US_POSIX":    en_US_POSIX.New,
			"rm":             rm.New,
			"ta_LK":          ta_LK.New,
			"xog":            xog.New,
			"en_GI":          en_GI.New,
			"en_SS":          en_SS.New,
			"fr_GA":          fr_GA.New,
			"fr_MC":          fr_MC.New,
			"ar_BH":          ar_BH.New,
			"dyo_SN":         dyo_SN.New,
			"en_AU":          en_AU.New,
			"fr_CD":          fr_CD.New,
			"ko_KR":          ko_KR.New,
			"sw_KE":          sw_KE.New,
			"zh_Hans_SG":     zh_Hans_SG.New,
			"en_CY":          en_CY.New,
			"fr_CI":          fr_CI.New,
			"kn":             kn.New,
			"ar_QA":          ar_QA.New,
			"az_Latn_AZ":     az_Latn_AZ.New,
			"bg_BG":          bg_BG.New,
			"ce_RU":          ce_RU.New,
			"es_CR":          es_CR.New,
			"fil":            fil.New,
			"sv_SE":          sv_SE.New,
			"ta":             ta.New,
			"ca_ES_VALENCIA": ca_ES_VALENCIA.New,
			"el_GR":          el_GR.New,
			"en_CH":          en_CH.New,
			"en_SL":          en_SL.New,
			"ewo_CM":         ewo_CM.New,
			"fr_BF":          fr_BF.New,
			"lg":             lg.New,
			"ca_FR":          ca_FR.New,
			"en_US":          en_US.New,
			"nl_CW":          nl_CW.New,
			"sah":            sah.New,
			"zu":             zu.New,
			"en_CM":          en_CM.New,
			"fr_TG":          fr_TG.New,
			"rwk":            rwk.New,
			"vun_TZ":         vun_TZ.New,
			"fa":             fa.New,
			"nus_SS":         nus_SS.New,
			"da":             da.New,
			"en_BW":          en_BW.New,
			"fr_MR":          fr_MR.New,
			"lb":             lb.New,
			"pt_MO":          pt_MO.New,
			"ro_RO":          ro_RO.New,
			"ar_KM":          ar_KM.New,
			"es_CU":          es_CU.New,
			"fr_RW":          fr_RW.New,
			"ki":             ki.New,
			"pa_Guru_IN":     pa_Guru_IN.New,
			"sn":             sn.New,
			"sr_Cyrl_ME":     sr_Cyrl_ME.New,
			"ug_CN":          ug_CN.New,
			"en_CA":          en_CA.New,
			"kea":            kea.New,
			"zh_Hant":        zh_Hant.New,
			"cs_CZ":          cs_CZ.New,
			"de_CH":          de_CH.New,
			"en_DE":          en_DE.New,
			"en_NF":          en_NF.New,
			"es_IC":          es_IC.New,
			"hsb":            hsb.New,
			"ig":             ig.New,
			"nb":             nb.New,
			"pt_LU":          pt_LU.New,
			"rm_CH":          rm_CH.New,
			"rn":             rn.New,
			"sr_Cyrl_XK":     sr_Cyrl_XK.New,
			"ta_IN":          ta_IN.New,
			"ar_001":         ar_001.New,
			"ar_IL":          ar_IL.New,
			"de_LU":          de_LU.New,
			"ff_SN":          ff_SN.New,
			"ha":             ha.New,
			"luy":            luy.New,
			"pt_GQ":          pt_GQ.New,
			"en_CC":          en_CC.New,
			"en_SC":          en_SC.New,
			"haw":            haw.New,
			"rwk_TZ":         rwk_TZ.New,
			"sk":             sk.New,
			"zh_Hant_MO":     zh_Hant_MO.New,
			"en_GU":          en_GU.New,
			"km_KH":          km_KH.New,
			"ln_AO":          ln_AO.New,
			"mfe_MU":         mfe_MU.New,
			"se_SE":          se_SE.New,
			"seh":            seh.New,
			"ca":             ca.New,
			"en_TC":          en_TC.New,
			"fr_DJ":          fr_DJ.New,
			"ga_IE":          ga_IE.New,
			"kkj_CM":         kkj_CM.New,
			"smn":            smn.New,
			"sr_Cyrl_BA":     sr_Cyrl_BA.New,
			"tk_TM":          tk_TM.New,
			"bn":             bn.New,
			"en_NA":          en_NA.New,
			"ja_JP":          ja_JP.New,
			"jgo_CM":         jgo_CM.New,
			"nnh_CM":         nnh_CM.New,
			"pt_BR":          pt_BR.New,
			"sn_ZW":          sn_ZW.New,
			"yav":            yav.New,
			"ar_MR":          ar_MR.New,
			"brx_IN":         brx_IN.New,
			"fi":             fi.New,
			"ms_MY":          ms_MY.New,
			"mt_MT":          mt_MT.New,
			"ti_ER":          ti_ER.New,
			"vi_VN":          vi_VN.New,
			"en_LS":          en_LS.New,
			"en_TO":          en_TO.New,
			"fy_NL":          fy_NL.New,
			"so_DJ":          so_DJ.New,
			"sv":             sv.New,
			"kl":             kl.New,
			"luo":            luo.New,
			"nl":             nl.New,
			"to_TO":          to_TO.New,
			"en_VG":          en_VG.New,
			"mn_MN":          mn_MN.New,
			"asa_TZ":         asa_TZ.New,
			"bas_CM":         bas_CM.New,
			"en_PW":          en_PW.New,
			"se_NO":          se_NO.New,
			"af_ZA":          af_ZA.New,
			"rw":             rw.New,
			"vun":            vun.New,
			"zh_Hans_HK":     zh_Hans_HK.New,
			"ar_IQ":          ar_IQ.New,
			"ee_GH":          ee_GH.New,
			"ff_GN":          ff_GN.New,
			"fr_NC":          fr_NC.New,
			"mg":             mg.New,
			"pa_Arab":        pa_Arab.New,
			"yue_HK":         yue_HK.New,
			"en_MY":          en_MY.New,
			"es_ES":          es_ES.New,
			"hu":             hu.New,
			"eo_001":         eo_001.New,
			"es_PY":          es_PY.New,
			"fr_CH":          fr_CH.New,
			"gl_ES":          gl_ES.New,
			"mgh":            mgh.New,
			"yi_001":         yi_001.New,
			"zgh":            zgh.New,
			"en_DM":          en_DM.New,
			"pt_CV":          pt_CV.New,
			"ru_KG":          ru_KG.New,
			"ti_ET":          ti_ET.New,
			"en_CK":          en_CK.New,
			"haw_US":         haw_US.New,
			"jgo":            jgo.New,
			"lt_LT":          lt_LT.New,
			"mzn_IR":         mzn_IR.New,
			"naq_NA":         naq_NA.New,
			"om":             om.New,
			"zh_Hant_HK":     zh_Hant_HK.New,
			"es_PR":          es_PR.New,
			"kam_KE":         kam_KE.New,
			"ksb":            ksb.New,
			"pa":             pa.New,
			"agq_CM":         agq_CM.New,
			"bs":             bs.New,
			"kam":            kam.New,
			"nl_BQ":          nl_BQ.New,
			"saq_KE":         saq_KE.New,
			"uz":             uz.New,
			"ha_NG":          ha_NG.New,
			"id_ID":          id_ID.New,
			"ks":             ks.New,
			"nl_BE":          nl_BE.New,
			"or":             or.New,
			"asa":            asa.New,
			"fa_IR":          fa_IR.New,
			"mer_KE":         mer_KE.New,
			"vai_Latn_LR":    vai_Latn_LR.New,
			"vo_001":         vo_001.New,
			"yi":             yi.New,
			"cu":             cu.New,
			"en_NR":          en_NR.New,
			"en_PR":          en_PR.New,
			"nmg":            nmg.New,
			"pt_CH":          pt_CH.New,
			"ast":            ast.New,
			"en_PG":          en_PG.New,
			"es_NI":          es_NI.New,
			"ff_CM":          ff_CM.New,
			"lkt":            lkt.New,
			"ln_CG":          ln_CG.New,
			"pl":             pl.New,
			"ps":             ps.New,
			"to":             to.New,
			"fr_NE":          fr_NE.New,
			"ru_UA":          ru_UA.New,
			"uz_Arab":        uz_Arab.New,
			"en_GY":          en_GY.New,
			"en_LC":          en_LC.New,
			"en_SI":          en_SI.New,
			"kn_IN":          kn_IN.New,
			"mas":            mas.New,
			"qu_BO":          qu_BO.New,
			"dav_KE":         dav_KE.New,
			"de_DE":          de_DE.New,
			"dua":            dua.New,
			"en_AS":          en_AS.New,
			"fy":             fy.New,
			"he":             he.New,
			"guz":            guz.New,
			"kl_GL":          kl_GL.New,
			"ky_KG":          ky_KG.New,
			"ms":             ms.New,
			"ur_PK":          ur_PK.New,
			"zh_Hant_TW":     zh_Hant_TW.New,
			"ar_ER":          ar_ER.New,
			"en_150":         en_150.New,
			"en_GB":          en_GB.New,
			"en_NL":          en_NL.New,
			"lkt_US":         lkt_US.New,
			"luo_KE":         luo_KE.New,
			"mua_CM":         mua_CM.New,
			"en_IE":          en_IE.New,
			"fr_LU":          fr_LU.New,
			"guz_KE":         guz_KE.New,
			"lag_TZ":         lag_TZ.New,
			"ps_AF":          ps_AF.New,
			"sw_UG":          sw_UG.New,
			"zgh_MA":         zgh_MA.New,
			"bm_ML":          bm_ML.New,
			"en_GM":          en_GM.New,
			"es_GT":          es_GT.New,
			"ja":             ja.New,
			"ksh_DE":         ksh_DE.New,
			"os":             os.New,
			"ar_EG":          ar_EG.New,
			"en_TZ":          en_TZ.New,
			"he_IL":          he_IL.New,
			"nyn":            nyn.New,
			"yo_NG":          yo_NG.New,
			"am":             am.New,
			"bas":            bas.New,
			"en_IO":          en_IO.New,
			"en_KN":          en_KN.New,
			"rw_RW":          rw_RW.New,
			"az":             az.New,
			"bem_ZM":         bem_ZM.New,
			"en_MW":          en_MW.New,
			"es_AR":          es_AR.New,
			"hr_BA":          hr_BA.New,
			"se":             se.New,
			"tzm_MA":         tzm_MA.New,
			"ce":             ce.New,
			"en_CX":          en_CX.New,
			"en_IM":          en_IM.New,
			"ksf_CM":         ksf_CM.New,
			"sq":             sq.New,
			"bo_CN":          bo_CN.New,
			"en_001":         en_001.New,
			"en_IL":          en_IL.New,
			"en_SE":          en_SE.New,
			"fil_PH":         fil_PH.New,
			"fr_KM":          fr_KM.New,
			"mgo":            mgo.New,
			"chr":            chr.New,
			"ckb":            ckb.New,
			"en_PN":          en_PN.New,
			"en_SX":          en_SX.New,
			"en_VU":          en_VU.New,
			"gu_IN":          gu_IN.New,
			"lrc_IR":         lrc_IR.New,
			"be":             be.New,
			"en_MT":          en_MT.New,
			"om_ET":          om_ET.New,
			"ar_DJ":          ar_DJ.New,
			"ar_YE":          ar_YE.New,
			"de_BE":          de_BE.New,
			"en_HK":          en_HK.New,
			"en_KY":          en_KY.New,
			"en_NU":          en_NU.New,
			"ha_NE":          ha_NE.New,
			"nn_NO":          nn_NO.New,
			"nus":            nus.New,
			"sg_CF":          sg_CF.New,
			"vai_Vaii":       vai_Vaii.New,
			"kea_CV":         kea_CV.New,
			"ks_IN":          ks_IN.New,
			"ksf":            ksf.New,
			"mgo_CM":         mgo_CM.New,
			"os_GE":          os_GE.New,
			"qu":             qu.New,
			"ca_IT":          ca_IT.New,
			"fr_MQ":          fr_MQ.New,
			"jmc":            jmc.New,
			"uz_Latn_UZ":     uz_Latn_UZ.New,
			"bg":             bg.New,
			"en_KI":          en_KI.New,
			"it_IT":          it_IT.New,
			"lrc_IQ":         lrc_IQ.New,
			"nn":             nn.New,
			"sr_Cyrl":        sr_Cyrl.New,
			"ar_SA":          ar_SA.New,
			"de_LI":          de_LI.New,
			"en_DG":          en_DG.New,
			"en_PK":          en_PK.New,
			"nd_ZW":          nd_ZW.New,
			"zh":             zh.New,
			"ar_SO":          ar_SO.New,
			"bez":            bez.New,
			"bo":             bo.New,
			"en_UG":          en_UG.New,
			"fr_GP":          fr_GP.New,
			"fr_SY":          fr_SY.New,
			"kde":            kde.New,
			"ko":             ko.New,
			"bm":             bm.New,
			"es_CO":          es_CO.New,
			"es_HN":          es_HN.New,
			"fr_BL":          fr_BL.New,
			"fr_PM":          fr_PM.New,
			"jmc_TZ":         jmc_TZ.New,
			"ln_CF":          ln_CF.New,
			"gsw":            gsw.New,
			"kln_KE":         kln_KE.New,
			"nb_NO":          nb_NO.New,
			"om_KE":          om_KE.New,
			"ta_SG":          ta_SG.New,
			"cy_GB":          cy_GB.New,
			"en_FI":          en_FI.New,
			"pt_GW":          pt_GW.New,
			"en_JE":          en_JE.New,
			"it":             it.New,
			"mr_IN":          mr_IN.New,
			"qu_PE":          qu_PE.New,
			"sg":             sg.New,
			"sr_Latn_RS":     sr_Latn_RS.New,
			"cu_RU":          cu_RU.New,
			"en_ZW":          en_ZW.New,
			"es_PH":          es_PH.New,
			"prg":            prg.New,
			"uk_UA":          uk_UA.New,
			"zu_ZA":          zu_ZA.New,
			"mer":            mer.New,
			"nl_NL":          nl_NL.New,
			"ru":             ru.New,
			"so":             so.New,
			"th":             th.New,
			"af_NA":          af_NA.New,
			"en_BE":          en_BE.New,
			"en_PH":          en_PH.New,
			"fo":             fo.New,
			"ksh":            ksh.New,
			"si":             si.New,
			"eo":             eo.New,
			"km":             km.New,
			"fr_TN":          fr_TN.New,
			"ms_BN":          ms_BN.New,
			"so_SO":          so_SO.New,
			"sr_Latn":        sr_Latn.New,
			"en_GG":          en_GG.New,
			"en_JM":          en_JM.New,
			"en_UM":          en_UM.New,
			"kln":            kln.New,
			"shi_Latn_MA":    shi_Latn_MA.New,
			"bn_BD":          bn_BD.New,
			"bs_Latn":        bs_Latn.New,
			"dje_NE":         dje_NE.New,
			"en_MS":          en_MS.New,
			"en_TV":          en_TV.New,
			"fa_AF":          fa_AF.New,
			"hy_AM":          hy_AM.New,
			"khq_ML":         khq_ML.New,
			"ln":             ln.New,
			"ru_MD":          ru_MD.New,
			"root":           root.New,
			"br":             br.New,
			"hr_HR":          hr_HR.New,
			"nl_SX":          nl_SX.New,
			"as_IN":          as_IN.New,
			"hi_IN":          hi_IN.New,
			"kk_KZ":          kk_KZ.New,
			"kw":             kw.New,
			"lv":             lv.New,
			"en_SD":          en_SD.New,
			"rn_BI":          rn_BI.New,
			"sw":             sw.New,
			"am_ET":          am_ET.New,
			"ar_TN":          ar_TN.New,
			"es_MX":          es_MX.New,
			"es_SV":          es_SV.New,
			"fr_GQ":          fr_GQ.New,
			"fr_MF":          fr_MF.New,
			"lu":             lu.New,
			"lv_LV":          lv_LV.New,
			"uz_Latn":        uz_Latn.New,
			"ii_CN":          ii_CN.New,
			"or_IN":          or_IN.New,
			"ur_IN":          ur_IN.New,
			"bez_TZ":         bez_TZ.New,
			"bs_Cyrl_BA":     bs_Cyrl_BA.New,
			"de":             de.New,
			"dyo":            dyo.New,
			"id":             id.New,
			"lo_LA":          lo_LA.New,
			"sq_AL":          sq_AL.New,
			"sr":             sr.New,
		}
	})
}

// Map returns the map of locales to instance New function
func Map() LocaleMap {
	return localeMap
}
