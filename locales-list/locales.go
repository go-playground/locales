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
			"ar_ER":          ar_ER.New,
			"dje":            dje.New,
			"es_419":         es_419.New,
			"et":             et.New,
			"gv":             gv.New,
			"ii_CN":          ii_CN.New,
			"to_TO":          to_TO.New,
			"fr_SC":          fr_SC.New,
			"rw":             rw.New,
			"ar_JO":          ar_JO.New,
			"az_Cyrl":        az_Cyrl.New,
			"ckb_IR":         ckb_IR.New,
			"dyo":            dyo.New,
			"en_FJ":          en_FJ.New,
			"en_GG":          en_GG.New,
			"fr_GP":          fr_GP.New,
			"ar_MA":          ar_MA.New,
			"fr_CH":          fr_CH.New,
			"fr_GA":          fr_GA.New,
			"fr_MQ":          fr_MQ.New,
			"mgh_MZ":         mgh_MZ.New,
			"ne_NP":          ne_NP.New,
			"dje_NE":         dje_NE.New,
			"en_BE":          en_BE.New,
			"en_VG":          en_VG.New,
			"fo_FO":          fo_FO.New,
			"fr_CA":          fr_CA.New,
			"uz_Cyrl":        uz_Cyrl.New,
			"fil_PH":         fil_PH.New,
			"fr_FR":          fr_FR.New,
			"ja_JP":          ja_JP.New,
			"kde":            kde.New,
			"ksh_DE":         ksh_DE.New,
			"sl_SI":          sl_SI.New,
			"vi":             vi.New,
			"asa":            asa.New,
			"ce":             ce.New,
			"en_GB":          en_GB.New,
			"en_JM":          en_JM.New,
			"en_VI":          en_VI.New,
			"rn":             rn.New,
			"sr_Latn":        sr_Latn.New,
			"qu":             qu.New,
			"am_ET":          am_ET.New,
			"ar_IL":          ar_IL.New,
			"en_SG":          en_SG.New,
			"ewo":            ewo.New,
			"fr_TG":          fr_TG.New,
			"ms_BN":          ms_BN.New,
			"os":             os.New,
			"sk":             sk.New,
			"ar":             ar.New,
			"ar_LB":          ar_LB.New,
			"en_BM":          en_BM.New,
			"es_US":          es_US.New,
			"lt_LT":          lt_LT.New,
			"luo_KE":         luo_KE.New,
			"sl":             sl.New,
			"bez":            bez.New,
			"cs":             cs.New,
			"pt_ST":          pt_ST.New,
			"sv_AX":          sv_AX.New,
			"bn_IN":          bn_IN.New,
			"ca_ES_VALENCIA": ca_ES_VALENCIA.New,
			"en_NG":          en_NG.New,
			"mas_TZ":         mas_TZ.New,
			"pt_BR":          pt_BR.New,
			"ro_MD":          ro_MD.New,
			"mg":             mg.New,
			"bg":             bg.New,
			"en_GD":          en_GD.New,
			"en_IN":          en_IN.New,
			"fa_AF":          fa_AF.New,
			"fr_BI":          fr_BI.New,
			"lb":             lb.New,
			"lg_UG":          lg_UG.New,
			"ru_UA":          ru_UA.New,
			"sw_KE":          sw_KE.New,
			"bem_ZM":         bem_ZM.New,
			"bo_CN":          bo_CN.New,
			"en_FI":          en_FI.New,
			"en_LS":          en_LS.New,
			"ka":             ka.New,
			"ky":             ky.New,
			"lo_LA":          lo_LA.New,
			"ff_CM":          ff_CM.New,
			"ksh":            ksh.New,
			"rm":             rm.New,
			"smn_FI":         smn_FI.New,
			"ta_IN":          ta_IN.New,
			"te":             te.New,
			"yav":            yav.New,
			"es_PE":          es_PE.New,
			"lkt":            lkt.New,
			"twq_NE":         twq_NE.New,
			"de_AT":          de_AT.New,
			"fur_IT":         fur_IT.New,
			"vun_TZ":         vun_TZ.New,
			"ar_DJ":          ar_DJ.New,
			"el_GR":          el_GR.New,
			"ff_MR":          ff_MR.New,
			"ksb_TZ":         ksb_TZ.New,
			"mua_CM":         mua_CM.New,
			"nnh_CM":         nnh_CM.New,
			"or":             or.New,
			"kab_DZ":         kab_DZ.New,
			"en_NZ":          en_NZ.New,
			"en_PG":          en_PG.New,
			"fr_MC":          fr_MC.New,
			"ru_KZ":          ru_KZ.New,
			"kn_IN":          kn_IN.New,
			"en_CX":          en_CX.New,
			"en_TK":          en_TK.New,
			"en_TT":          en_TT.New,
			"es_DO":          es_DO.New,
			"es_GT":          es_GT.New,
			"fr_TD":          fr_TD.New,
			"haw_US":         haw_US.New,
			"pt_MO":          pt_MO.New,
			"yue_HK":         yue_HK.New,
			"sw_CD":          sw_CD.New,
			"agq_CM":         agq_CM.New,
			"ar_BH":          ar_BH.New,
			"es":             es.New,
			"fr_LU":          fr_LU.New,
			"gd":             gd.New,
			"ln_AO":          ln_AO.New,
			"sah":            sah.New,
			"ur":             ur.New,
			"dua_CM":         dua_CM.New,
			"en_CA":          en_CA.New,
			"ga":             ga.New,
			"gl":             gl.New,
			"khq":            khq.New,
			"zh_Hans_HK":     zh_Hans_HK.New,
			"fr_MG":          fr_MG.New,
			"mr_IN":          mr_IN.New,
			"mt_MT":          mt_MT.New,
			"or_IN":          or_IN.New,
			"sr_Cyrl":        sr_Cyrl.New,
			"gsw_LI":         gsw_LI.New,
			"ug":             ug.New,
			"xog":            xog.New,
			"bm":             bm.New,
			"en_DK":          en_DK.New,
			"en_SI":          en_SI.New,
			"it_SM":          it_SM.New,
			"om_ET":          om_ET.New,
			"so":             so.New,
			"ta_LK":          ta_LK.New,
			"nnh":            nnh.New,
			"de_DE":          de_DE.New,
			"en_CH":          en_CH.New,
			"en_DM":          en_DM.New,
			"en_MP":          en_MP.New,
			"en_MU":          en_MU.New,
			"es_CL":          es_CL.New,
			"es_PH":          es_PH.New,
			"cgg_UG":         cgg_UG.New,
			"sv_FI":          sv_FI.New,
			"ti_ET":          ti_ET.New,
			"as_IN":          as_IN.New,
			"brx_IN":         brx_IN.New,
			"es_MX":          es_MX.New,
			"fr_BF":          fr_BF.New,
			"yue":            yue.New,
			"az_Latn":        az_Latn.New,
			"dav_KE":         dav_KE.New,
			"hr":             hr.New,
			"de_LU":          de_LU.New,
			"en_PR":          en_PR.New,
			"en_US":          en_US.New,
			"ff":             ff.New,
			"bs_Cyrl":        bs_Cyrl.New,
			"en_SX":          en_SX.New,
			"gu_IN":          gu_IN.New,
			"kok_IN":         kok_IN.New,
			"ce_RU":          ce_RU.New,
			"dua":            dua.New,
			"dyo_SN":         dyo_SN.New,
			"en_GM":          en_GM.New,
			"fr_BJ":          fr_BJ.New,
			"fr_KM":          fr_KM.New,
			"ko":             ko.New,
			"kw":             kw.New,
			"zh_Hans_MO":     zh_Hans_MO.New,
			"bas":            bas.New,
			"fr_WF":          fr_WF.New,
			"lrc":            lrc.New,
			"en_ZW":          en_ZW.New,
			"fr_GQ":          fr_GQ.New,
			"fr_ML":          fr_ML.New,
			"nyn":            nyn.New,
			"qu_BO":          qu_BO.New,
			"en_150":         en_150.New,
			"en_NA":          en_NA.New,
			"en_SD":          en_SD.New,
			"nl_SX":          nl_SX.New,
			"ses_ML":         ses_ML.New,
			"sw_TZ":          sw_TZ.New,
			"tr_TR":          tr_TR.New,
			"de_BE":          de_BE.New,
			"en_TO":          en_TO.New,
			"nb_NO":          nb_NO.New,
			"rw_RW":          rw_RW.New,
			"ks_IN":          ks_IN.New,
			"yo_BJ":          yo_BJ.New,
			"bs":             bs.New,
			"en_VU":          en_VU.New,
			"fr_DZ":          fr_DZ.New,
			"rn_BI":          rn_BI.New,
			"en_IO":          en_IO.New,
			"en_ZA":          en_ZA.New,
			"es_EC":          es_EC.New,
			"fa_IR":          fa_IR.New,
			"kam":            kam.New,
			"nus":            nus.New,
			"sr_Latn_XK":     sr_Latn_XK.New,
			"cy_GB":          cy_GB.New,
			"fr_YT":          fr_YT.New,
			"nd_ZW":          nd_ZW.New,
			"fi_FI":          fi_FI.New,
			"kn":             kn.New,
			"en_DE":          en_DE.New,
			"en_IL":          en_IL.New,
			"ar_LY":          ar_LY.New,
			"en_UG":          en_UG.New,
			"es_ES":          es_ES.New,
			"fr_GF":          fr_GF.New,
			"lg":             lg.New,
			"ar_TN":          ar_TN.New,
			"bo":             bo.New,
			"en_PK":          en_PK.New,
			"lu":             lu.New,
			"pl":             pl.New,
			"shi_Latn_MA":    shi_Latn_MA.New,
			"root":           root.New,
			"hu_HU":          hu_HU.New,
			"nl_BE":          nl_BE.New,
			"so_KE":          so_KE.New,
			"ln_CD":          ln_CD.New,
			"nn":             nn.New,
			"pt_CH":          pt_CH.New,
			"pt_CV":          pt_CV.New,
			"dsb_DE":         dsb_DE.New,
			"ff_SN":          ff_SN.New,
			"os_GE":          os_GE.New,
			"qu_EC":          qu_EC.New,
			"sr_Latn_BA":     sr_Latn_BA.New,
			"es_CU":          es_CU.New,
			"ro_RO":          ro_RO.New,
			"so_DJ":          so_DJ.New,
			"ar_IQ":          ar_IQ.New,
			"fr_CG":          fr_CG.New,
			"hu":             hu.New,
			"lkt_US":         lkt_US.New,
			"mas_KE":         mas_KE.New,
			"sr":             sr.New,
			"en_KY":          en_KY.New,
			"en_MH":          en_MH.New,
			"en_SZ":          en_SZ.New,
			"tk":             tk.New,
			"tzm_MA":         tzm_MA.New,
			"uz_Latn":        uz_Latn.New,
			"sv":             sv.New,
			"ar_SY":          ar_SY.New,
			"en_BI":          en_BI.New,
			"es_SV":          es_SV.New,
			"mgh":            mgh.New,
			"my_MM":          my_MM.New,
			"naq":            naq.New,
			"seh":            seh.New,
			"saq":            saq.New,
			"dsb":            dsb.New,
			"en_GU":          en_GU.New,
			"es_AR":          es_AR.New,
			"fr_PM":          fr_PM.New,
			"ga_IE":          ga_IE.New,
			"mzn":            mzn.New,
			"pa":             pa.New,
			"te_IN":          te_IN.New,
			"dz":             dz.New,
			"fr_MU":          fr_MU.New,
			"ky_KG":          ky_KG.New,
			"ru_BY":          ru_BY.New,
			"en_SS":          en_SS.New,
			"fr":             fr.New,
			"pt_TL":          pt_TL.New,
			"ru_MD":          ru_MD.New,
			"sk_SK":          sk_SK.New,
			"sr_Cyrl_BA":     sr_Cyrl_BA.New,
			"da_GL":          da_GL.New,
			"fr_CM":          fr_CM.New,
			"hi":             hi.New,
			"seh_MZ":         seh_MZ.New,
			"uz_Arab_AF":     uz_Arab_AF.New,
			"he_IL":          he_IL.New,
			"kln_KE":         kln_KE.New,
			"mn_MN":          mn_MN.New,
			"sq_AL":          sq_AL.New,
			"mgo_CM":         mgo_CM.New,
			"nb_SJ":          nb_SJ.New,
			"agq":            agq.New,
			"asa_TZ":         asa_TZ.New,
			"eu":             eu.New,
			"id":             id.New,
			"zu_ZA":          zu_ZA.New,
			"af_ZA":          af_ZA.New,
			"en_KN":          en_KN.New,
			"en_MG":          en_MG.New,
			"haw":            haw.New,
			"ms_MY":          ms_MY.New,
			"zh_Hans_CN":     zh_Hans_CN.New,
			"chr":            chr.New,
			"de":             de.New,
			"kab":            kab.New,
			"ki_KE":          ki_KE.New,
			"teo_KE":         teo_KE.New,
			"uk":             uk.New,
			"ur_IN":          ur_IN.New,
			"fr_TN":          fr_TN.New,
			"en_AU":          en_AU.New,
			"es_CO":          es_CO.New,
			"eu_ES":          eu_ES.New,
			"fil":            fil.New,
			"fr_MA":          fr_MA.New,
			"fr_MR":          fr_MR.New,
			"fr_SY":          fr_SY.New,
			"it_CH":          it_CH.New,
			"jmc":            jmc.New,
			"yo":             yo.New,
			"ru":             ru.New,
			"en_AI":          en_AI.New,
			"en_TZ":          en_TZ.New,
			"en_VC":          en_VC.New,
			"ko_KR":          ko_KR.New,
			"mt":             mt.New,
			"ses":            ses.New,
			"en_PW":          en_PW.New,
			"fr_CD":          fr_CD.New,
			"kl":             kl.New,
			"ksf":            ksf.New,
			"lt":             lt.New,
			"nd":             nd.New,
			"nmg":            nmg.New,
			"ar_SA":          ar_SA.New,
			"be":             be.New,
			"be_BY":          be_BY.New,
			"ca":             ca.New,
			"is_IS":          is_IS.New,
			"zu":             zu.New,
			"en_CC":          en_CC.New,
			"jmc_TZ":         jmc_TZ.New,
			"mas":            mas.New,
			"nl":             nl.New,
			"saq_KE":         saq_KE.New,
			"smn":            smn.New,
			"en_NR":          en_NR.New,
			"en_TV":          en_TV.New,
			"gd_GB":          gd_GB.New,
			"lag_TZ":         lag_TZ.New,
			"sr_Cyrl_ME":     sr_Cyrl_ME.New,
			"uz_Latn_UZ":     uz_Latn_UZ.New,
			"en_BW":          en_BW.New,
			"tzm":            tzm.New,
			"vo_001":         vo_001.New,
			"ca_FR":          ca_FR.New,
			"guz_KE":         guz_KE.New,
			"mer_KE":         mer_KE.New,
			"shi_Latn":       shi_Latn.New,
			"sw_UG":          sw_UG.New,
			"ta_SG":          ta_SG.New,
			"th_TH":          th_TH.New,
			"zh_Hans":        zh_Hans.New,
			"uz":             uz.New,
			"en_DG":          en_DG.New,
			"fo_DK":          fo_DK.New,
			"nl_AW":          nl_AW.New,
			"nyn_UG":         nyn_UG.New,
			"os_RU":          os_RU.New,
			"pt":             pt.New,
			"so_SO":          so_SO.New,
			"pt_MZ":          pt_MZ.New,
			"en_LC":          en_LC.New,
			"en_US_POSIX":    en_US_POSIX.New,
			"fr_HT":          fr_HT.New,
			"hi_IN":          hi_IN.New,
			"km_KH":          km_KH.New,
			"lo":             lo.New,
			"mk":             mk.New,
			"ff_GN":          ff_GN.New,
			"twq":            twq.New,
			"en_IE":          en_IE.New,
			"fr_NE":          fr_NE.New,
			"kam_KE":         kam_KE.New,
			"ln":             ln.New,
			"th":             th.New,
			"ur_PK":          ur_PK.New,
			"wae_CH":         wae_CH.New,
			"shi_Tfng_MA":    shi_Tfng_MA.New,
			"ar_DZ":          ar_DZ.New,
			"ast_ES":         ast_ES.New,
			"da_DK":          da_DK.New,
			"en_ER":          en_ER.New,
			"gsw":            gsw.New,
			"jgo_CM":         jgo_CM.New,
			"ksf_CM":         ksf_CM.New,
			"si":             si.New,
			"en_RW":          en_RW.New,
			"eo_001":         eo_001.New,
			"fr_CI":          fr_CI.New,
			"kde_TZ":         kde_TZ.New,
			"si_LK":          si_LK.New,
			"ebu_KE":         ebu_KE.New,
			"en_FK":          en_FK.New,
			"es_PR":          es_PR.New,
			"es_VE":          es_VE.New,
			"gsw_CH":         gsw_CH.New,
			"ksb":            ksb.New,
			"ne_IN":          ne_IN.New,
			"bs_Latn":        bs_Latn.New,
			"chr_US":         chr_US.New,
			"ka_GE":          ka_GE.New,
			"mzn_IR":         mzn_IR.New,
			"ps":             ps.New,
			"mfe":            mfe.New,
			"az":             az.New,
			"ca_ES":          ca_ES.New,
			"en_GY":          en_GY.New,
			"en_ZM":          en_ZM.New,
			"fur":            fur.New,
			"gl_ES":          gl_ES.New,
			"lv_LV":          lv_LV.New,
			"uk_UA":          uk_UA.New,
			"ne":             ne.New,
			"tr_CY":          tr_CY.New,
			"vun":            vun.New,
			"en_MW":          en_MW.New,
			"ha":             ha.New,
			"luo":            luo.New,
			"luy_KE":         luy_KE.New,
			"nl_BQ":          nl_BQ.New,
			"ps_AF":          ps_AF.New,
			"sv_SE":          sv_SE.New,
			"ak_GH":          ak_GH.New,
			"en_JE":          en_JE.New,
			"sg_CF":          sg_CF.New,
			"ta":             ta.New,
			"kkj":            kkj.New,
			"bm_ML":          bm_ML.New,
			"cgg":            cgg.New,
			"en_PH":          en_PH.New,
			"en_SB":          en_SB.New,
			"ewo_CM":         ewo_CM.New,
			"ig":             ig.New,
			"kea_CV":         kea_CV.New,
			"naq_NA":         naq_NA.New,
			"ak":             ak.New,
			"kk_KZ":          kk_KZ.New,
			"af_NA":          af_NA.New,
			"ar_KW":          ar_KW.New,
			"en_AT":          en_AT.New,
			"en_BB":          en_BB.New,
			"en_MT":          en_MT.New,
			"fi":             fi.New,
			"zh_Hant_MO":     zh_Hant_MO.New,
			"fy":             fy.New,
			"vo":             vo.New,
			"zgh_MA":         zgh_MA.New,
			"en_MY":          en_MY.New,
			"es_UY":          es_UY.New,
			"km":             km.New,
			"prg_001":        prg_001.New,
			"vai_Latn":       vai_Latn.New,
			"en_PN":          en_PN.New,
			"es_EA":          es_EA.New,
			"es_NI":          es_NI.New,
			"sah_RU":         sah_RU.New,
			"shi_Tfng":       shi_Tfng.New,
			"pa_Guru_IN":     pa_Guru_IN.New,
			"ar_QA":          ar_QA.New,
			"ar_SD":          ar_SD.New,
			"en_NU":          en_NU.New,
			"en_WS":          en_WS.New,
			"fy_NL":          fy_NL.New,
			"hr_BA":          hr_BA.New,
			"om_KE":          om_KE.New,
			"vai_Vaii_LR":    vai_Vaii_LR.New,
			"ar_001":         ar_001.New,
			"ar_SO":          ar_SO.New,
			"ee":             ee.New,
			"el":             el.New,
			"fr_MF":          fr_MF.New,
			"nb":             nb.New,
			"ru_RU":          ru_RU.New,
			"bn_BD":          bn_BD.New,
			"om":             om.New,
			"sn":             sn.New,
			"kea":            kea.New,
			"sq":             sq.New,
			"sq_XK":          sq_XK.New,
			"sr_Cyrl_XK":     sr_Cyrl_XK.New,
			"sr_Latn_RS":     sr_Latn_RS.New,
			"ar_PS":          ar_PS.New,
			"en_BZ":          en_BZ.New,
			"hsb_DE":         hsb_DE.New,
			"hy":             hy.New,
			"it":             it.New,
			"kok":            kok.New,
			"vi_VN":          vi_VN.New,
			"en_001":         en_001.New,
			"en_FM":          en_FM.New,
			"fr_PF":          fr_PF.New,
			"nl_CW":          nl_CW.New,
			"da":             da.New,
			"en_IM":          en_IM.New,
			"es_BR":          es_BR.New,
			"fr_VU":          fr_VU.New,
			"pl_PL":          pl_PL.New,
			"sn_ZW":          sn_ZW.New,
			"fa":             fa.New,
			"nl_SR":          nl_SR.New,
			"pa_Arab":        pa_Arab.New,
			"zh_Hant":        zh_Hant.New,
			"cu_RU":          cu_RU.New,
			"it_IT":          it_IT.New,
			"luy":            luy.New,
			"mfe_MU":         mfe_MU.New,
			"nmg_CM":         nmg_CM.New,
			"pt_PT":          pt_PT.New,
			"yo_NG":          yo_NG.New,
			"ee_TG":          ee_TG.New,
			"en_SL":          en_SL.New,
			"gu":             gu.New,
			"am":             am.New,
			"gsw_FR":         gsw_FR.New,
			"kk":             kk.New,
			"mn":             mn.New,
			"sg":             sg.New,
			"ml_IN":          ml_IN.New,
			"vai_Latn_LR":    vai_Latn_LR.New,
			"zgh":            zgh.New,
			"ar_OM":          ar_OM.New,
			"en_MO":          en_MO.New,
			"kl_GL":          kl_GL.New,
			"shi":            shi.New,
			"vai_Vaii":       vai_Vaii.New,
			"zh_Hant_HK":     zh_Hant_HK.New,
			"is":             is.New,
			"pa_Guru":        pa_Guru.New,
			"pt_AO":          pt_AO.New,
			"teo_UG":         teo_UG.New,
			"ti_ER":          ti_ER.New,
			"ebu":            ebu.New,
			"ig_NG":          ig_NG.New,
			"nl_NL":          nl_NL.New,
			"pt_GW":          pt_GW.New,
			"sr_Latn_ME":     sr_Latn_ME.New,
			"tr":             tr.New,
			"uz_Cyrl_UZ":     uz_Cyrl_UZ.New,
			"ar_YE":          ar_YE.New,
			"en_BS":          en_BS.New,
			"es_CR":          es_CR.New,
			"ml":             ml.New,
			"rwk":            rwk.New,
			"wae":            wae.New,
			"yav_CM":         yav_CM.New,
			"cy":             cy.New,
			"en_HK":          en_HK.New,
			"fr_BL":          fr_BL.New,
			"se":             se.New,
			"tk_TM":          tk_TM.New,
			"ar_EH":          ar_EH.New,
			"az_Latn_AZ":     az_Latn_AZ.New,
			"el_CY":          el_CY.New,
			"ii":             ii.New,
			"mua":            mua.New,
			"ug_CN":          ug_CN.New,
			"fr_BE":          fr_BE.New,
			"ar_SS":          ar_SS.New,
			"bn":             bn.New,
			"en":             en.New,
			"es_BO":          es_BO.New,
			"es_HN":          es_HN.New,
			"es_IC":          es_IC.New,
			"es_PA":          es_PA.New,
			"fr_GN":          fr_GN.New,
			"fr_NC":          fr_NC.New,
			"ko_KP":          ko_KP.New,
			"ro":             ro.New,
			"rwk_TZ":         rwk_TZ.New,
			"sr_Cyrl_RS":     sr_Cyrl_RS.New,
			"uz_Arab":        uz_Arab.New,
			"fr_RE":          fr_RE.New,
			"lrc_IR":         lrc_IR.New,
			"ti":             ti.New,
			"az_Cyrl_AZ":     az_Cyrl_AZ.New,
			"cs_CZ":          cs_CZ.New,
			"sbp":            sbp.New,
			"ta_MY":          ta_MY.New,
			"vai":            vai.New,
			"dz_BT":          dz_BT.New,
			"fr_SN":          fr_SN.New,
			"guz":            guz.New,
			"hr_HR":          hr_HR.New,
			"ja":             ja.New,
			"rof":            rof.New,
			"bas_CM":         bas_CM.New,
			"bez_TZ":         bez_TZ.New,
			"gv_IM":          gv_IM.New,
			"hsb":            hsb.New,
			"kkj_CM":         kkj_CM.New,
			"br_FR":          br_FR.New,
			"en_NL":          en_NL.New,
			"fr_RW":          fr_RW.New,
			"mk_MK":          mk_MK.New,
			"my":             my.New,
			"ar_EG":          ar_EG.New,
			"bem":            bem.New,
			"en_CK":          en_CK.New,
			"en_LR":          en_LR.New,
			"en_TC":          en_TC.New,
			"lrc_IQ":         lrc_IQ.New,
			"mgo":            mgo.New,
			"pa_Arab_PK":     pa_Arab_PK.New,
			"pt_LU":          pt_LU.New,
			"en_KI":          en_KI.New,
			"hy_AM":          hy_AM.New,
			"sw":             sw.New,
			"bo_IN":          bo_IN.New,
			"en_UM":          en_UM.New,
			"ln_CG":          ln_CG.New,
			"xog_UG":         xog_UG.New,
			"zh":             zh.New,
			"bs_Cyrl_BA":     bs_Cyrl_BA.New,
			"en_AG":          en_AG.New,
			"en_AS":          en_AS.New,
			"fr_CF":          fr_CF.New,
			"kln":            kln.New,
			"lv":             lv.New,
			"se_NO":          se_NO.New,
			"ln_CF":          ln_CF.New,
			"ar_MR":          ar_MR.New,
			"ast":            ast.New,
			"bs_Latn_BA":     bs_Latn_BA.New,
			"cu":             cu.New,
			"en_SH":          en_SH.New,
			"ha_NG":          ha_NG.New,
			"he":             he.New,
			"pt_GQ":          pt_GQ.New,
			"rof_TZ":         rof_TZ.New,
			"so_ET":          so_ET.New,
			"yi_001":         yi_001.New,
			"af":             af.New,
			"lu_CD":          lu_CD.New,
			"mr":             mr.New,
			"en_CM":          en_CM.New,
			"rm_CH":          rm_CH.New,
			"se_FI":          se_FI.New,
			"to":             to.New,
			"bg_BG":          bg_BG.New,
			"en_SC":          en_SC.New,
			"en_SE":          en_SE.New,
			"es_PY":          es_PY.New,
			"se_SE":          se_SE.New,
			"mg_MG":          mg_MG.New,
			"ar_AE":          ar_AE.New,
			"br":             br.New,
			"ca_AD":          ca_AD.New,
			"ckb":            ckb.New,
			"fo":             fo.New,
			"ha_NE":          ha_NE.New,
			"lb_LU":          lb_LU.New,
			"sq_MK":          sq_MK.New,
			"de_CH":          de_CH.New,
			"en_GH":          en_GH.New,
			"fr_DJ":          fr_DJ.New,
			"ks":             ks.New,
			"as":             as.New,
			"ckb_IQ":         ckb_IQ.New,
			"de_LI":          de_LI.New,
			"zh_Hant_TW":     zh_Hant_TW.New,
			"teo":            teo.New,
			"brx":            brx.New,
			"ee_GH":          ee_GH.New,
			"en_CY":          en_CY.New,
			"en_KE":          en_KE.New,
			"eo":             eo.New,
			"jgo":            jgo.New,
			"ru_KG":          ru_KG.New,
			"en_NF":          en_NF.New,
			"khq_ML":         khq_ML.New,
			"ki":             ki.New,
			"qu_PE":          qu_PE.New,
			"yi":             yi.New,
			"es_GQ":          es_GQ.New,
			"ms_SG":          ms_SG.New,
			"nn_NO":          nn_NO.New,
			"id_ID":          id_ID.New,
			"ar_KM":          ar_KM.New,
			"ar_TD":          ar_TD.New,
			"ca_IT":          ca_IT.New,
			"dav":            dav.New,
			"en_GI":          en_GI.New,
			"et_EE":          et_EE.New,
			"ha_GH":          ha_GH.New,
			"kw_GB":          kw_GB.New,
			"mer":            mer.New,
			"ms":             ms.New,
			"nus_SS":         nus_SS.New,
			"en_MS":          en_MS.New,
			"lag":            lag.New,
			"prg":            prg.New,
			"sbp_TZ":         sbp_TZ.New,
			"zh_Hans_SG":     zh_Hans_SG.New,
		}
	})
}

// Map returns the map of locales to instance New function
func Map() LocaleMap {
	return localeMap
}
