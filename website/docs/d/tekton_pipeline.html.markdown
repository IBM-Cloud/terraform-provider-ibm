---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline"
description: |-
  Get information about tekton_pipeline
subcategory: "CD Tekton Pipeline"
---

# ibm_cd_tekton_pipeline

Provides a read-only data source for tekton_pipeline. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_tekton_pipeline" "tekton_pipeline" {
	id = "94619026-912b-4d92-8f51-6c74f0692d90"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) ID of current instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the tekton_pipeline.
* `build_number` - (Optional, Integer) The latest pipeline run build number. If this property is absent, the pipeline hasn't had any pipelineRuns.
  * Constraints: The minimum value is `1`.

* `created` - (Required, String) Standard RFC 3339 Date Time String.

* `definitions` - (Required, List) Definition list.
Nested scheme for **definitions**:
	* `id` - (Required, String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `scm_source` - (Required, List) Scm source for tekton pipeline defintion.
	Nested scheme for **scm_source**:
		* `branch` - (Optional, String) A branch of the repo, branch field doesn't coexist with tag field.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `path` - (Required, String) The path to the definitions yaml files.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `tag` - (Optional, String) A tag of the repo.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]{1,235}$/`.
		* `url` - (Required, String) General href URL.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `service_instance_id` - (Required, String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

* `enabled` - (Required, Boolean) Flag whether this pipeline enabled.

* `html_url` - (Required, String) Dashboard URL of this pipeline.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `name` - (Required, String) String.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,235}[a-zA-Z0-9]$/`.

* `pipeline_definition` - (Optional, List) Tekton pipeline definition document detail object. If this property is absent, the pipeline has no definitions added.
Nested scheme for **pipeline_definition**:
	* `id` - (Optional, String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `status` - (Optional, String) The state of pipeline definition status.
	  * Constraints: Allowable values are: `updated`, `outdated`, `updating`, `failed`.

* `properties` - (Required, List) Tekton pipeline's environment properties.
Nested scheme for **properties**:
	* `default` - (Optional, String) Default option for SINGLE_SELECT property type.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `enum` - (Optional, List) Options for SINGLE_SELECT property type.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `name` - (Required, String) Property name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
	* `path` - (Optional, String) property path for INTEGRATION type properties.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
	* `type` - (Required, String) Property type.
	  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`, `APPCONFIG`.
	* `value` - (Optional, String) String format property value.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.

* `resource_group_id` - (Required, String) ID.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]+$/`.

* `status` - (Required, String) Pipeline status.
  * Constraints: Allowable values are: `configured`, `configuring`.

* `toolchain` - (Required, List) Toolchain object.
Nested scheme for **toolchain**:
	* `crn` - (Required, String) The CRN for the toolchain that containing the tekton pipeline.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `id` - (Required, String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

* `triggers` - (Required, List) Tekton pipeline triggers list.
Nested scheme for **triggers**:
	* `concurrency` - (Optional, List) Concurrency object.
	Nested scheme for **concurrency**:
		* `max_concurrent_runs` - (Optional, Integer) Defines the maximum number of concurrent runs for this trigger.
		  * Constraints: The minimum value is `0`.
	* `cron` - (Optional, String) Needed only for timer trigger type. Cron expression for timer trigger. Maximum frequency is every 5 minutes.
	  * Constraints: The maximum length is `253` characters. The minimum length is `5` characters. The value must match regular expression `/^(\\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\\*\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\\*|([0-9]|1[0-9]|2[0-3])|\\*\/([0-9]|1[0-9]|2[0-3])) (\\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\\*\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\\*|([1-9]|1[0-2])|\\*\/([1-9]|1[0-2])) (\\*|([0-6])|\\*\/([0-6]))$/`.
	* `disabled` - (Optional, Boolean) flag whether the trigger is disabled.
	* `event_listener` - (Optional, String) Event listener name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `events` - (Optional, List) Needed only for git trigger type. Events object defines the events this git trigger listening to.
	Nested scheme for **events**:
		* `pull_request` - (Optional, Boolean) If true, the trigger starts when tekton pipeline service receive a repo pull reqeust's 'open' or 'update' git webhook event.
		* `pull_request_closed` - (Optional, Boolean) If true, the trigger starts when tekton pipeline service receive a repo pull reqeust's 'close' git webhook event.
		* `push` - (Optional, Boolean) If true, the trigger starts when tekton pipeline service receive a repo's 'push' git webhook event.
	* `id` - (Optional, String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `name` - (Optional, String) name of the duplicated trigger.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,235}[a-zA-Z0-9]$/`.
	* `properties` - (Optional, List) Trigger properties.
	Nested scheme for **properties**:
		* `default` - (Optional, String) Default option for SINGLE_SELECT property type.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `enum` - (Optional, List) Options for SINGLE_SELECT property type.
		  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `href` - (Optional, String) General href URL.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `name` - (Required, String) Property name.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
		* `path` - (Optional, String) property path for INTEGRATION type properties.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
		* `type` - (Required, String) Property type.
		  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`, `APPCONFIG`.
		* `value` - (Optional, String) String format property value.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
	* `scm_source` - (Optional, List) Scm source for git type tekton pipeline trigger.
	Nested scheme for **scm_source**:
		* `blind_connection` - (Optional, Boolean) Needed only for git trigger type. Branch name of the repo.
		* `branch` - (Optional, String) Needed only for git trigger type. Branch name of the repo. Branch field doesn't coexist with pattern field.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `hook_id` - (Optional, String) Webhook ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `pattern` - (Optional, String) Needed only for git trigger type. Git branch or tag pattern to listen to. Please refer to https://github.com/micromatch/micromatch for pattern syntax.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^.{1,235}$/`.
		* `url` - (Optional, String) Needed only for git trigger type. Repo URL that listening to.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `secret` - (Optional, List) Needed only for generic trigger type. Secret used to start generic trigger.
	Nested scheme for **secret**:
		* `algorithm` - (Optional, String) Algorithm used for "digestMatches" secret type.
		  * Constraints: Allowable values are: `md4`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`, `sha512_224`, `sha512_256`, `ripemd160`.
		* `key_name` - (Optional, String) Secret name, not needed if type is "internalValidation".
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `source` - (Optional, String) Secret location, not needed if secret type is "internalValidation".
		  * Constraints: Allowable values are: `header`, `payload`, `query`.
		* `type` - (Optional, String) Secret type.
		  * Constraints: Allowable values are: `tokenMatches`, `digestMatches`, `internalValidation`.
		* `value` - (Optional, String) Secret value, not needed if secret type is "internalValidation".
		  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `service_instance_id` - (Optional, String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `source_trigger_id` - (Optional, String) source trigger ID to clone from.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `tags` - (Optional, List) Trigger tags array.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `timezone` - (Optional, String) Needed only for timer trigger type. Timezones for timer trigger.
	  * Constraints: Allowable values are: `Africa/Abidjan`, `Africa/Accra`, `Africa/Addis_Ababa`, `Africa/Algiers`, `Africa/Asmara`, `Africa/Asmera`, `Africa/Bamako`, `Africa/Bangui`, `Africa/Banjul`, `Africa/Bissau`, `Africa/Blantyre`, `Africa/Brazzaville`, `Africa/Bujumbura`, `Africa/Cairo`, `Africa/Casablanca`, `Africa/Ceuta`, `Africa/Conakry`, `Africa/Dakar`, `Africa/Dar_es_Salaam`, `Africa/Djibouti`, `Africa/Douala`, `Africa/El_Aaiun`, `Africa/Freetown`, `Africa/Gaborone`, `Africa/Harare`, `Africa/Johannesburg`, `Africa/Juba`, `Africa/Kampala`, `Africa/Khartoum`, `Africa/Kigali`, `Africa/Kinshasa`, `Africa/Lagos`, `Africa/Libreville`, `Africa/Lome`, `Africa/Luanda`, `Africa/Lubumbashi`, `Africa/Lusaka`, `Africa/Malabo`, `Africa/Maputo`, `Africa/Maseru`, `Africa/Mbabane`, `Africa/Mogadishu`, `Africa/Monrovia`, `Africa/Nairobi`, `Africa/Ndjamena`, `Africa/Niamey`, `Africa/Nouakchott`, `Africa/Ouagadougou`, `Africa/Porto-Novo`, `Africa/Sao_Tome`, `Africa/Timbuktu`, `Africa/Tripoli`, `Africa/Tunis`, `Africa/Windhoek`, `America/Adak`, `America/Anchorage`, `America/Anguilla`, `America/Antigua`, `America/Araguaina`, `America/Argentina/Buenos_Aires`, `America/Argentina/Catamarca`, `America/Argentina/ComodRivadavia`, `America/Argentina/Cordoba`, `America/Argentina/Jujuy`, `America/Argentina/La_Rioja`, `America/Argentina/Mendoza`, `America/Argentina/Rio_Gallegos`, `America/Argentina/Salta`, `America/Argentina/San_Juan`, `America/Argentina/San_Luis`, `America/Argentina/Tucuman`, `America/Argentina/Ushuaia`, `America/Aruba`, `America/Asuncion`, `America/Atikokan`, `America/Atka`, `America/Bahia`, `America/Bahia_Banderas`, `America/Barbados`, `America/Belem`, `America/Belize`, `America/Blanc-Sablon`, `America/Boa_Vista`, `America/Bogota`, `America/Boise`, `America/Buenos_Aires`, `America/Cambridge_Bay`, `America/Campo_Grande`, `America/Cancun`, `America/Caracas`, `America/Catamarca`, `America/Cayenne`, `America/Cayman`, `America/Chicago`, `America/Chihuahua`, `America/Coral_Harbour`, `America/Cordoba`, `America/Costa_Rica`, `America/Creston`, `America/Cuiaba`, `America/Curacao`, `America/Danmarkshavn`, `America/Dawson`, `America/Dawson_Creek`, `America/Denver`, `America/Detroit`, `America/Dominica`, `America/Edmonton`, `America/Eirunepe`, `America/El_Salvador`, `America/Ensenada`, `America/Fort_Nelson`, `America/Fort_Wayne`, `America/Fortaleza`, `America/Glace_Bay`, `America/Godthab`, `America/Goose_Bay`, `America/Grand_Turk`, `America/Grenada`, `America/Guadeloupe`, `America/Guatemala`, `America/Guayaquil`, `America/Guyana`, `America/Halifax`, `America/Havana`, `America/Hermosillo`, `America/Indiana/Indianapolis`, `America/Indiana/Knox`, `America/Indiana/Marengo`, `America/Indiana/Petersburg`, `America/Indiana/Tell_City`, `America/Indiana/Vevay`, `America/Indiana/Vincennes`, `America/Indiana/Winamac`, `America/Indianapolis`, `America/Inuvik`, `America/Iqaluit`, `America/Jamaica`, `America/Jujuy`, `America/Juneau`, `America/Kentucky/Louisville`, `America/Kentucky/Monticello`, `America/Knox_IN`, `America/Kralendijk`, `America/La_Paz`, `America/Lima`, `America/Los_Angeles`, `America/Louisville`, `America/Lower_Princes`, `America/Maceio`, `America/Managua`, `America/Manaus`, `America/Marigot`, `America/Martinique`, `America/Matamoros`, `America/Mazatlan`, `America/Mendoza`, `America/Menominee`, `America/Merida`, `America/Metlakatla`, `America/Mexico_City`, `America/Miquelon`, `America/Moncton`, `America/Monterrey`, `America/Montevideo`, `America/Montreal`, `America/Montserrat`, `America/Nassau`, `America/New_York`, `America/Nipigon`, `America/Nome`, `America/Noronha`, `America/North_Dakota/Beulah`, `America/North_Dakota/Center`, `America/North_Dakota/New_Salem`, `America/Ojinaga`, `America/Panama`, `America/Pangnirtung`, `America/Paramaribo`, `America/Phoenix`, `America/Port-au-Prince`, `America/Port_of_Spain`, `America/Porto_Acre`, `America/Porto_Velho`, `America/Puerto_Rico`, `America/Punta_Arenas`, `America/Rainy_River`, `America/Rankin_Inlet`, `America/Recife`, `America/Regina`, `America/Resolute`, `America/Rio_Branco`, `America/Rosario`, `America/Santa_Isabel`, `America/Santarem`, `America/Santiago`, `America/Santo_Domingo`, `America/Sao_Paulo`, `America/Scoresbysund`, `America/Shiprock`, `America/Sitka`, `America/St_Barthelemy`, `America/St_Johns`, `America/St_Kitts`, `America/St_Lucia`, `America/St_Thomas`, `America/St_Vincent`, `America/Swift_Current`, `America/Tegucigalpa`, `America/Thule`, `America/Thunder_Bay`, `America/Tijuana`, `America/Toronto`, `America/Tortola`, `America/Vancouver`, `America/Virgin`, `America/Whitehorse`, `America/Winnipeg`, `America/Yakutat`, `America/Yellowknife`, `Antarctica/Casey`, `Antarctica/Davis`, `Antarctica/DumontDUrville`, `Antarctica/Macquarie`, `Antarctica/Mawson`, `Antarctica/McMurdo`, `Antarctica/Palmer`, `Antarctica/Rothera`, `Antarctica/South_Pole`, `Antarctica/Syowa`, `Antarctica/Troll`, `Antarctica/Vostok`, `Arctic/Longyearbyen`, `Asia/Aden`, `Asia/Almaty`, `Asia/Amman`, `Asia/Anadyr`, `Asia/Aqtau`, `Asia/Aqtobe`, `Asia/Ashgabat`, `Asia/Ashkhabad`, `Asia/Atyrau`, `Asia/Baghdad`, `Asia/Bahrain`, `Asia/Baku`, `Asia/Bangkok`, `Asia/Barnaul`, `Asia/Beirut`, `Asia/Bishkek`, `Asia/Brunei`, `Asia/Calcutta`, `Asia/Chita`, `Asia/Choibalsan`, `Asia/Chongqing`, `Asia/Chungking`, `Asia/Colombo`, `Asia/Dacca`, `Asia/Damascus`, `Asia/Dhaka`, `Asia/Dili`, `Asia/Dubai`, `Asia/Dushanbe`, `Asia/Famagusta`, `Asia/Gaza`, `Asia/Harbin`, `Asia/Hebron`, `Asia/Ho_Chi_Minh`, `Asia/Hong_Kong`, `Asia/Hovd`, `Asia/Irkutsk`, `Asia/Istanbul`, `Asia/Jakarta`, `Asia/Jayapura`, `Asia/Jerusalem`, `Asia/Kabul`, `Asia/Kamchatka`, `Asia/Karachi`, `Asia/Kashgar`, `Asia/Kathmandu`, `Asia/Katmandu`, `Asia/Khandyga`, `Asia/Kolkata`, `Asia/Krasnoyarsk`, `Asia/Kuala_Lumpur`, `Asia/Kuching`, `Asia/Kuwait`, `Asia/Macao`, `Asia/Macau`, `Asia/Magadan`, `Asia/Makassar`, `Asia/Manila`, `Asia/Muscat`, `Asia/Nicosia`, `Asia/Novokuznetsk`, `Asia/Novosibirsk`, `Asia/Omsk`, `Asia/Oral`, `Asia/Phnom_Penh`, `Asia/Pontianak`, `Asia/Pyongyang`, `Asia/Qatar`, `Asia/Qostanay`, `Asia/Qyzylorda`, `Asia/Rangoon`, `Asia/Riyadh`, `Asia/Saigon`, `Asia/Sakhalin`, `Asia/Samarkand`, `Asia/Seoul`, `Asia/Shanghai`, `Asia/Singapore`, `Asia/Srednekolymsk`, `Asia/Taipei`, `Asia/Tashkent`, `Asia/Tbilisi`, `Asia/Tehran`, `Asia/Tel_Aviv`, `Asia/Thimbu`, `Asia/Thimphu`, `Asia/Tokyo`, `Asia/Tomsk`, `Asia/Ujung_Pandang`, `Asia/Ulaanbaatar`, `Asia/Ulan_Bator`, `Asia/Urumqi`, `Asia/Ust-Nera`, `Asia/Vientiane`, `Asia/Vladivostok`, `Asia/Yakutsk`, `Asia/Yangon`, `Asia/Yekaterinburg`, `Asia/Yerevan`, `Atlantic/Azores`, `Atlantic/Bermuda`, `Atlantic/Canary`, `Atlantic/Cape_Verde`, `Atlantic/Faeroe`, `Atlantic/Faroe`, `Atlantic/Jan_Mayen`, `Atlantic/Madeira`, `Atlantic/Reykjavik`, `Atlantic/South_Georgia`, `Atlantic/St_Helena`, `Atlantic/Stanley`, `Australia/ACT`, `Australia/Adelaide`, `Australia/Brisbane`, `Australia/Broken_Hill`, `Australia/Canberra`, `Australia/Currie`, `Australia/Darwin`, `Australia/Eucla`, `Australia/Hobart`, `Australia/LHI`, `Australia/Lindeman`, `Australia/Lord_Howe`, `Australia/Melbourne`, `Australia/NSW`, `Australia/North`, `Australia/Perth`, `Australia/Queensland`, `Australia/South`, `Australia/Sydney`, `Australia/Tasmania`, `Australia/Victoria`, `Australia/West`, `Australia/Yancowinna`, `Brazil/Acre`, `Brazil/DeNoronha`, `Brazil/East`, `Brazil/West`, `CET`, `CST6CDT`, `Canada/Atlantic`, `Canada/Central`, `Canada/Eastern`, `Canada/Mountain`, `Canada/Newfoundland`, `Canada/Pacific`, `Canada/Saskatchewan`, `Canada/Yukon`, `Chile/Continental`, `Chile/EasterIsland`, `Cuba`, `EET`, `EST`, `EST5EDT`, `Egypt`, `Eire`, `Etc/GMT`, `Etc/GMT+0`, `Etc/GMT+1`, `Etc/GMT+10`, `Etc/GMT+11`, `Etc/GMT+12`, `Etc/GMT+2`, `Etc/GMT+3`, `Etc/GMT+4`, `Etc/GMT+5`, `Etc/GMT+6`, `Etc/GMT+7`, `Etc/GMT+8`, `Etc/GMT+9`, `Etc/GMT-0`, `Etc/GMT-1`, `Etc/GMT-10`, `Etc/GMT-11`, `Etc/GMT-12`, `Etc/GMT-13`, `Etc/GMT-14`, `Etc/GMT-2`, `Etc/GMT-3`, `Etc/GMT-4`, `Etc/GMT-5`, `Etc/GMT-6`, `Etc/GMT-7`, `Etc/GMT-8`, `Etc/GMT-9`, `Etc/GMT0`, `Etc/Greenwich`, `Etc/UCT`, `Etc/UTC`, `Etc/Universal`, `Etc/Zulu`, `Europe/Amsterdam`, `Europe/Andorra`, `Europe/Astrakhan`, `Europe/Athens`, `Europe/Belfast`, `Europe/Belgrade`, `Europe/Berlin`, `Europe/Bratislava`, `Europe/Brussels`, `Europe/Bucharest`, `Europe/Budapest`, `Europe/Busingen`, `Europe/Chisinau`, `Europe/Copenhagen`, `Europe/Dublin`, `Europe/Gibraltar`, `Europe/Guernsey`, `Europe/Helsinki`, `Europe/Isle_of_Man`, `Europe/Istanbul`, `Europe/Jersey`, `Europe/Kaliningrad`, `Europe/Kiev`, `Europe/Kirov`, `Europe/Lisbon`, `Europe/Ljubljana`, `Europe/London`, `Europe/Luxembourg`, `Europe/Madrid`, `Europe/Malta`, `Europe/Mariehamn`, `Europe/Minsk`, `Europe/Monaco`, `Europe/Moscow`, `Europe/Nicosia`, `Europe/Oslo`, `Europe/Paris`, `Europe/Podgorica`, `Europe/Prague`, `Europe/Riga`, `Europe/Rome`, `Europe/Samara`, `Europe/San_Marino`, `Europe/Sarajevo`, `Europe/Saratov`, `Europe/Simferopol`, `Europe/Skopje`, `Europe/Sofia`, `Europe/Stockholm`, `Europe/Tallinn`, `Europe/Tirane`, `Europe/Tiraspol`, `Europe/Ulyanovsk`, `Europe/Uzhgorod`, `Europe/Vaduz`, `Europe/Vatican`, `Europe/Vienna`, `Europe/Vilnius`, `Europe/Volgograd`, `Europe/Warsaw`, `Europe/Zagreb`, `Europe/Zaporozhye`, `Europe/Zurich`, `GB`, `GB-Eire`, `GMT`, `GMT+0`, `GMT-0`, `GMT0`, `Greenwich`, `HST`, `Hongkong`, `Iceland`, `Indian/Antananarivo`, `Indian/Chagos`, `Indian/Christmas`, `Indian/Cocos`, `Indian/Comoro`, `Indian/Kerguelen`, `Indian/Mahe`, `Indian/Maldives`, `Indian/Mauritius`, `Indian/Mayotte`, `Indian/Reunion`, `Iran`, `Israel`, `Jamaica`, `Japan`, `Kwajalein`, `Libya`, `MET`, `MST`, `MST7MDT`, `Mexico/BajaNorte`, `Mexico/BajaSur`, `Mexico/General`, `NZ`, `NZ-CHAT`, `Navajo`, `PRC`, `PST8PDT`, `Pacific/Apia`, `Pacific/Auckland`, `Pacific/Bougainville`, `Pacific/Chatham`, `Pacific/Chuuk`, `Pacific/Easter`, `Pacific/Efate`, `Pacific/Enderbury`, `Pacific/Fakaofo`, `Pacific/Fiji`, `Pacific/Funafuti`, `Pacific/Galapagos`, `Pacific/Gambier`,
	`Pacific/Guadalcanal`, `Pacific/Guam`, `Pacific/Honolulu`, `Pacific/Johnston`, `Pacific/Kiritimati`, `Pacific/Kosrae`, `Pacific/Kwajalein`, `Pacific/Majuro`, `Pacific/Marquesas`, `Pacific/Midway`, `Pacific/Nauru`, `Pacific/Niue`, `Pacific/Norfolk`, `Pacific/Noumea`, `Pacific/Pago_Pago`, `Pacific/Palau`, `Pacific/Pitcairn`, `Pacific/Pohnpei`, `Pacific/Ponape`, `Pacific/Port_Moresby`, `Pacific/Rarotonga`, `Pacific/Saipan`, `Pacific/Samoa`, `Pacific/Tahiti`, `Pacific/Tarawa`, `Pacific/Tongatapu`, `Pacific/Truk`, `Pacific/Wake`, `Pacific/Wallis`, `Pacific/Yap`, `Poland`, `Portugal`, `ROC`, `ROK`, `Singapore`, `Turkey`, `UCT`, `US/Alaska`, `US/Aleutian`, `US/Arizona`, `US/Central`, `US/East-Indiana`, `US/Eastern`, `US/Hawaii`, `US/Indiana-Starke`, `US/Michigan`, `US/Mountain`, `US/Pacific`, `US/Pacific-New`, `US/Samoa`, `UTC`, `Universal`, `W-SU`, `WET`, `Zulu`.
	* `type` - (Optional, String) Trigger type.
	  * Constraints: Allowable values are: .
	* `worker` - (Optional, List) Trigger worker used to run the trigger, the trigger worker overrides the default pipeline worker.If not exist, this trigger uses default pipeline worker.
	Nested scheme for **worker**:
		* `id` - (Required, String)
		* `name` - (Optional, String) worker name.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_. \\(\\)\\[\\]]{1,235}$/`.
		* `type` - (Optional, String) worker type.
		  * Constraints: Allowable values are: `private`, `public`.

* `updated_at` - (Required, String) Standard RFC 3339 Date Time String.

* `worker` - (Required, List) Default pipeline worker used to run the pipeline.
Nested scheme for **worker**:
	* `id` - (Required, String)
	* `name` - (Optional, String) worker name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_. \\(\\)\\[\\]]{1,235}$/`.
	* `type` - (Optional, String) worker type.
	  * Constraints: Allowable values are: `private`, `public`.

