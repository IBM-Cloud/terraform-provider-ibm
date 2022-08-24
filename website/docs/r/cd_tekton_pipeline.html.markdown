---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline"
description: |-
  Manages cd_tekton_pipeline.
subcategory: "CD Tekton Pipeline"
---

# ibm_cd_tekton_pipeline

Provides a resource for cd_tekton_pipeline. This allows cd_tekton_pipeline to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
  worker {
		id = "id"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `worker` - (Optional, List) Worker object containing worker ID only. If omitted the IBM Managed shared workers are used by default.
Nested scheme for **worker**:
	* `id` - (Required, String)

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_tekton_pipeline.
* `build_number` - (Integer) The latest pipeline run build number. If this property is absent, the pipeline hasn't had any pipeline runs.
  * Constraints: The minimum value is `1`.
* `created_at` - (String) Standard RFC 3339 Date Time String.
* `definitions` - (List) Definition list.
Nested scheme for **definitions**:
	* `id` - (String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `scm_source` - (List) SCM source for Tekton pipeline definition.
	Nested scheme for **scm_source**:
		* `branch` - (String) A branch from the repo. One of branch or tag must be specified, but only one or the other.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `path` - (String) The path to the definition's yaml files.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `service_instance_id` - (String) ID of the SCM repository service instance.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
		* `tag` - (String) A tag from the repo. One of branch or tag must be specified, but only one or the other.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]{1,235}$/`.
		* `url` - (Forces new resource, String) URL of the definition repository.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `enabled` - (Boolean) Flag whether this pipeline is enabled.
* `html_url` - (String) Dashboard URL of this pipeline.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `name` - (String) String.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,235}[a-zA-Z0-9]$/`.
* `properties` - (List) Tekton pipeline's environment properties.
Nested scheme for **properties**:
	* `default` - (String) Default option for SINGLE_SELECT property type. Only needed when using SINGLE_SELECT property type.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `enum` - (List) Options for SINGLE_SELECT property type. Only needed when using SINGLE_SELECT property type.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `name` - (Forces new resource, String) Property name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
	* `path` - (String) A dot notation path for INTEGRATION type properties to select a value from the tool integration.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
	* `type` - (String) Property type.
	  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`, `APPCONFIG`.
	* `value` - (String) Property value.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
* `resource_group_id` - (String) ID.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_]+$/`.
* `status` - (String) Pipeline status.
  * Constraints: Allowable values are: `configured`, `configuring`.
* `toolchain` - (List) Toolchain object.
Nested scheme for **toolchain**:
	* `crn` - (String) The CRN for the toolchain that contains the Tekton pipeline.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `id` - (String) UUID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `triggers` - (List) Tekton pipeline triggers list.
Nested scheme for **triggers**:
	* `cron` - (String) Only needed for timer triggers. Cron expression for timer trigger. Maximum frequency is every 5 minutes.
	  * Constraints: The maximum length is `253` characters. The minimum length is `5` characters. The value must match regular expression `/^(\\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\\*\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\\*|([0-9]|1[0-9]|2[0-3])|\\*\/([0-9]|1[0-9]|2[0-3])) (\\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\\*\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\\*|([1-9]|1[0-2])|\\*\/([1-9]|1[0-2])) (\\*|([0-6])|\\*\/([0-6]))$/`.
	* `disabled` - (Boolean) Flag whether the trigger is disabled. If omitted the trigger is enabled by default.
	* `event_listener` - (String) Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `events` - (List) Only needed for Git triggers. Events object defines the events to which this Git trigger listens.
	Nested scheme for **events**:
		* `pull_request` - (Boolean) If true, the trigger listens for 'open pull request' or 'update pull request' Git webhook events.
		* `pull_request_closed` - (Boolean) If true, the trigger listens for 'close pull request' Git webhook events.
		* `push` - (Boolean) If true, the trigger listens for 'push' Git webhook events.
	* `href` - (String) API URL for interacting with the trigger.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) ID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `max_concurrent_runs` - (Integer) Defines the maximum number of concurrent runs for this trigger. Omit this property to disable the concurrency limit.
	* `name` - (String) Trigger name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,235}[a-zA-Z0-9]$/`.
	* `properties` - (List) Trigger properties.
	Nested scheme for **properties**:
		* `default` - (String) Default option for SINGLE_SELECT property type. Only needed for SINGLE_SELECT property type.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `enum` - (List) Options for SINGLE_SELECT property type. Only needed for SINGLE_SELECT property type.
		  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `href` - (String) API URL for interacting with the trigger property.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `name` - (Forces new resource, String) Property name.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
		* `path` - (String) A dot notation path for INTEGRATION type properties to select a value from the tool integration. If left blank the full tool integration JSON will be selected.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
		* `type` - (String) Property type.
		  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`, `APPCONFIG`.
		* `value` - (String) Property value. Can be empty and should be omitted for SINGLE_SELECT property type.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
	* `scm_source` - (List) SCM source repository for a Git trigger. Only needed for Git triggers.
	Nested scheme for **scm_source**:
		* `blind_connection` - (Boolean) Set this boolean to true if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide. False by default.
		* `branch` - (String) Name of a branch from the repo. One of branch or tag must be specified, but only one or the other.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `hook_id` - (String) ID of the webhook from the repo. Computed upon creation of the trigger.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `pattern` - (String) Git branch or tag pattern to listen to. Please refer to https://github.com/micromatch/micromatch for pattern syntax.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^.{1,235}$/`.
		* `service_instance_id` - (String) ID of the repository service instance.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
		* `url` - (Forces new resource, String) URL of the repository to which the trigger is listening.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `secret` - (List) Only needed for generic webhook trigger type. Secret used to start generic webhook trigger.
	Nested scheme for **secret**:
		* `algorithm` - (String) Algorithm used for "digestMatches" secret type. Only needed for "digestMatches" secret type.
		  * Constraints: Allowable values are: `md4`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`, `sha512_224`, `sha512_256`, `ripemd160`.
		* `key_name` - (String) Secret name, not needed if type is "internalValidation".
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
		* `source` - (String) Secret location, not needed if secret type is "internalValidation".
		  * Constraints: Allowable values are: `header`, `payload`, `query`.
		* `type` - (String) Secret type.
		  * Constraints: Allowable values are: `tokenMatches`, `digestMatches`, `internalValidation`.
		* `value` - (String) Secret value, not needed if secret type is "internalValidation".
		  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `source_trigger_id` - (String) ID of the trigger to duplicate. Only needed when duplicating a trigger.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `tags` - (List) Trigger tags array.
	  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,235}$/`.
	* `timezone` - (String) Only needed for timer triggers. Timezone for timer trigger.
	  * Constraints: Allowable values are: `Africa/Abidjan`, `Africa/Accra`, `Africa/Addis_Ababa`, `Africa/Algiers`, `Africa/Asmara`, `Africa/Asmera`, `Africa/Bamako`, `Africa/Bangui`, `Africa/Banjul`, `Africa/Bissau`, `Africa/Blantyre`, `Africa/Brazzaville`, `Africa/Bujumbura`, `Africa/Cairo`, `Africa/Casablanca`, `Africa/Ceuta`, `Africa/Conakry`, `Africa/Dakar`, `Africa/Dar_es_Salaam`, `Africa/Djibouti`, `Africa/Douala`, `Africa/El_Aaiun`, `Africa/Freetown`, `Africa/Gaborone`, `Africa/Harare`, `Africa/Johannesburg`, `Africa/Juba`, `Africa/Kampala`, `Africa/Khartoum`, `Africa/Kigali`, `Africa/Kinshasa`, `Africa/Lagos`, `Africa/Libreville`, `Africa/Lome`, `Africa/Luanda`, `Africa/Lubumbashi`, `Africa/Lusaka`, `Africa/Malabo`, `Africa/Maputo`, `Africa/Maseru`, `Africa/Mbabane`, `Africa/Mogadishu`, `Africa/Monrovia`, `Africa/Nairobi`, `Africa/Ndjamena`, `Africa/Niamey`, `Africa/Nouakchott`, `Africa/Ouagadougou`, `Africa/Porto-Novo`, `Africa/Sao_Tome`, `Africa/Timbuktu`, `Africa/Tripoli`, `Africa/Tunis`, `Africa/Windhoek`, `America/Adak`, `America/Anchorage`, `America/Anguilla`, `America/Antigua`, `America/Araguaina`, `America/Argentina/Buenos_Aires`, `America/Argentina/Catamarca`, `America/Argentina/ComodRivadavia`, `America/Argentina/Cordoba`, `America/Argentina/Jujuy`, `America/Argentina/La_Rioja`, `America/Argentina/Mendoza`, `America/Argentina/Rio_Gallegos`, `America/Argentina/Salta`, `America/Argentina/San_Juan`, `America/Argentina/San_Luis`, `America/Argentina/Tucuman`, `America/Argentina/Ushuaia`, `America/Aruba`, `America/Asuncion`, `America/Atikokan`, `America/Atka`, `America/Bahia`, `America/Bahia_Banderas`, `America/Barbados`, `America/Belem`, `America/Belize`, `America/Blanc-Sablon`, `America/Boa_Vista`, `America/Bogota`, `America/Boise`, `America/Buenos_Aires`, `America/Cambridge_Bay`, `America/Campo_Grande`, `America/Cancun`, `America/Caracas`, `America/Catamarca`, `America/Cayenne`, `America/Cayman`, `America/Chicago`, `America/Chihuahua`, `America/Coral_Harbour`, `America/Cordoba`, `America/Costa_Rica`, `America/Creston`, `America/Cuiaba`, `America/Curacao`, `America/Danmarkshavn`, `America/Dawson`, `America/Dawson_Creek`, `America/Denver`, `America/Detroit`, `America/Dominica`, `America/Edmonton`, `America/Eirunepe`, `America/El_Salvador`, `America/Ensenada`, `America/Fort_Nelson`, `America/Fort_Wayne`, `America/Fortaleza`, `America/Glace_Bay`, `America/Godthab`, `America/Goose_Bay`, `America/Grand_Turk`, `America/Grenada`, `America/Guadeloupe`, `America/Guatemala`, `America/Guayaquil`, `America/Guyana`, `America/Halifax`, `America/Havana`, `America/Hermosillo`, `America/Indiana/Indianapolis`, `America/Indiana/Knox`, `America/Indiana/Marengo`, `America/Indiana/Petersburg`, `America/Indiana/Tell_City`, `America/Indiana/Vevay`, `America/Indiana/Vincennes`, `America/Indiana/Winamac`, `America/Indianapolis`, `America/Inuvik`, `America/Iqaluit`, `America/Jamaica`, `America/Jujuy`, `America/Juneau`, `America/Kentucky/Louisville`, `America/Kentucky/Monticello`, `America/Knox_IN`, `America/Kralendijk`, `America/La_Paz`, `America/Lima`, `America/Los_Angeles`, `America/Louisville`, `America/Lower_Princes`, `America/Maceio`, `America/Managua`, `America/Manaus`, `America/Marigot`, `America/Martinique`, `America/Matamoros`, `America/Mazatlan`, `America/Mendoza`, `America/Menominee`, `America/Merida`, `America/Metlakatla`, `America/Mexico_City`, `America/Miquelon`, `America/Moncton`, `America/Monterrey`, `America/Montevideo`, `America/Montreal`, `America/Montserrat`, `America/Nassau`, `America/New_York`, `America/Nipigon`, `America/Nome`, `America/Noronha`, `America/North_Dakota/Beulah`, `America/North_Dakota/Center`, `America/North_Dakota/New_Salem`, `America/Ojinaga`, `America/Panama`, `America/Pangnirtung`, `America/Paramaribo`, `America/Phoenix`, `America/Port-au-Prince`, `America/Port_of_Spain`, `America/Porto_Acre`, `America/Porto_Velho`, `America/Puerto_Rico`, `America/Punta_Arenas`, `America/Rainy_River`, `America/Rankin_Inlet`, `America/Recife`, `America/Regina`, `America/Resolute`, `America/Rio_Branco`, `America/Rosario`, `America/Santa_Isabel`, `America/Santarem`, `America/Santiago`, `America/Santo_Domingo`, `America/Sao_Paulo`, `America/Scoresbysund`, `America/Shiprock`, `America/Sitka`, `America/St_Barthelemy`, `America/St_Johns`, `America/St_Kitts`, `America/St_Lucia`, `America/St_Thomas`, `America/St_Vincent`, `America/Swift_Current`, `America/Tegucigalpa`, `America/Thule`, `America/Thunder_Bay`, `America/Tijuana`, `America/Toronto`, `America/Tortola`, `America/Vancouver`, `America/Virgin`, `America/Whitehorse`, `America/Winnipeg`, `America/Yakutat`, `America/Yellowknife`, `Antarctica/Casey`, `Antarctica/Davis`, `Antarctica/DumontDUrville`, `Antarctica/Macquarie`, `Antarctica/Mawson`, `Antarctica/McMurdo`, `Antarctica/Palmer`, `Antarctica/Rothera`, `Antarctica/South_Pole`, `Antarctica/Syowa`, `Antarctica/Troll`, `Antarctica/Vostok`, `Arctic/Longyearbyen`, `Asia/Aden`, `Asia/Almaty`, `Asia/Amman`, `Asia/Anadyr`, `Asia/Aqtau`, `Asia/Aqtobe`, `Asia/Ashgabat`, `Asia/Ashkhabad`, `Asia/Atyrau`, `Asia/Baghdad`, `Asia/Bahrain`, `Asia/Baku`, `Asia/Bangkok`, `Asia/Barnaul`, `Asia/Beirut`, `Asia/Bishkek`, `Asia/Brunei`, `Asia/Calcutta`, `Asia/Chita`, `Asia/Choibalsan`, `Asia/Chongqing`, `Asia/Chungking`, `Asia/Colombo`, `Asia/Dacca`, `Asia/Damascus`, `Asia/Dhaka`, `Asia/Dili`, `Asia/Dubai`, `Asia/Dushanbe`, `Asia/Famagusta`, `Asia/Gaza`, `Asia/Harbin`, `Asia/Hebron`, `Asia/Ho_Chi_Minh`, `Asia/Hong_Kong`, `Asia/Hovd`, `Asia/Irkutsk`, `Asia/Istanbul`, `Asia/Jakarta`, `Asia/Jayapura`, `Asia/Jerusalem`, `Asia/Kabul`, `Asia/Kamchatka`, `Asia/Karachi`, `Asia/Kashgar`, `Asia/Kathmandu`, `Asia/Katmandu`, `Asia/Khandyga`, `Asia/Kolkata`, `Asia/Krasnoyarsk`, `Asia/Kuala_Lumpur`, `Asia/Kuching`, `Asia/Kuwait`, `Asia/Macao`, `Asia/Macau`, `Asia/Magadan`, `Asia/Makassar`, `Asia/Manila`, `Asia/Muscat`, `Asia/Nicosia`, `Asia/Novokuznetsk`, `Asia/Novosibirsk`, `Asia/Omsk`, `Asia/Oral`, `Asia/Phnom_Penh`, `Asia/Pontianak`, `Asia/Pyongyang`, `Asia/Qatar`, `Asia/Qostanay`, `Asia/Qyzylorda`, `Asia/Rangoon`, `Asia/Riyadh`, `Asia/Saigon`, `Asia/Sakhalin`, `Asia/Samarkand`, `Asia/Seoul`, `Asia/Shanghai`, `Asia/Singapore`, `Asia/Srednekolymsk`, `Asia/Taipei`, `Asia/Tashkent`, `Asia/Tbilisi`, `Asia/Tehran`, `Asia/Tel_Aviv`, `Asia/Thimbu`, `Asia/Thimphu`, `Asia/Tokyo`, `Asia/Tomsk`, `Asia/Ujung_Pandang`, `Asia/Ulaanbaatar`, `Asia/Ulan_Bator`, `Asia/Urumqi`, `Asia/Ust-Nera`, `Asia/Vientiane`, `Asia/Vladivostok`, `Asia/Yakutsk`, `Asia/Yangon`, `Asia/Yekaterinburg`, `Asia/Yerevan`, `Atlantic/Azores`, `Atlantic/Bermuda`, `Atlantic/Canary`, `Atlantic/Cape_Verde`, `Atlantic/Faeroe`, `Atlantic/Faroe`, `Atlantic/Jan_Mayen`, `Atlantic/Madeira`, `Atlantic/Reykjavik`, `Atlantic/South_Georgia`, `Atlantic/St_Helena`, `Atlantic/Stanley`, `Australia/ACT`, `Australia/Adelaide`, `Australia/Brisbane`, `Australia/Broken_Hill`, `Australia/Canberra`, `Australia/Currie`, `Australia/Darwin`, `Australia/Eucla`, `Australia/Hobart`, `Australia/LHI`, `Australia/Lindeman`, `Australia/Lord_Howe`, `Australia/Melbourne`, `Australia/NSW`, `Australia/North`, `Australia/Perth`, `Australia/Queensland`, `Australia/South`, `Australia/Sydney`, `Australia/Tasmania`, `Australia/Victoria`, `Australia/West`, `Australia/Yancowinna`, `Brazil/Acre`, `Brazil/DeNoronha`, `Brazil/East`, `Brazil/West`, `CET`, `CST6CDT`, `Canada/Atlantic`, `Canada/Central`, `Canada/Eastern`, `Canada/Mountain`, `Canada/Newfoundland`, `Canada/Pacific`, `Canada/Saskatchewan`, `Canada/Yukon`, `Chile/Continental`, `Chile/EasterIsland`, `Cuba`, `EET`, `EST`, `EST5EDT`, `Egypt`, `Eire`, `Etc/GMT`, `Etc/GMT+0`, `Etc/GMT+1`, `Etc/GMT+10`, `Etc/GMT+11`, `Etc/GMT+12`, `Etc/GMT+2`, `Etc/GMT+3`, `Etc/GMT+4`, `Etc/GMT+5`, `Etc/GMT+6`, `Etc/GMT+7`, `Etc/GMT+8`, `Etc/GMT+9`, `Etc/GMT-0`, `Etc/GMT-1`, `Etc/GMT-10`, `Etc/GMT-11`, `Etc/GMT-12`, `Etc/GMT-13`, `Etc/GMT-14`, `Etc/GMT-2`, `Etc/GMT-3`, `Etc/GMT-4`, `Etc/GMT-5`, `Etc/GMT-6`, `Etc/GMT-7`, `Etc/GMT-8`, `Etc/GMT-9`, `Etc/GMT0`, `Etc/Greenwich`, `Etc/UCT`, `Etc/UTC`, `Etc/Universal`, `Etc/Zulu`, `Europe/Amsterdam`, `Europe/Andorra`, `Europe/Astrakhan`, `Europe/Athens`, `Europe/Belfast`, `Europe/Belgrade`, `Europe/Berlin`, `Europe/Bratislava`, `Europe/Brussels`, `Europe/Bucharest`, `Europe/Budapest`, `Europe/Busingen`, `Europe/Chisinau`, `Europe/Copenhagen`, `Europe/Dublin`, `Europe/Gibraltar`, `Europe/Guernsey`, `Europe/Helsinki`, `Europe/Isle_of_Man`, `Europe/Istanbul`, `Europe/Jersey`, `Europe/Kaliningrad`, `Europe/Kiev`, `Europe/Kirov`, `Europe/Lisbon`, `Europe/Ljubljana`, `Europe/London`, `Europe/Luxembourg`, `Europe/Madrid`, `Europe/Malta`, `Europe/Mariehamn`, `Europe/Minsk`, `Europe/Monaco`, `Europe/Moscow`, `Europe/Nicosia`, `Europe/Oslo`, `Europe/Paris`, `Europe/Podgorica`, `Europe/Prague`, `Europe/Riga`, `Europe/Rome`, `Europe/Samara`, `Europe/San_Marino`, `Europe/Sarajevo`, `Europe/Saratov`, `Europe/Simferopol`, `Europe/Skopje`, `Europe/Sofia`, `Europe/Stockholm`, `Europe/Tallinn`, `Europe/Tirane`, `Europe/Tiraspol`, `Europe/Ulyanovsk`, `Europe/Uzhgorod`, `Europe/Vaduz`, `Europe/Vatican`, `Europe/Vienna`, `Europe/Vilnius`, `Europe/Volgograd`, `Europe/Warsaw`, `Europe/Zagreb`, `Europe/Zaporozhye`, `Europe/Zurich`, `GB`, `GB-Eire`, `GMT`, `GMT+0`, `GMT-0`, `GMT0`, `Greenwich`, `HST`, `Hongkong`, `Iceland`, `Indian/Antananarivo`, `Indian/Chagos`, `Indian/Christmas`, `Indian/Cocos`, `Indian/Comoro`, `Indian/Kerguelen`, `Indian/Mahe`, `Indian/Maldives`, `Indian/Mauritius`, `Indian/Mayotte`, `Indian/Reunion`, `Iran`, `Israel`, `Jamaica`, `Japan`, `Kwajalein`, `Libya`, `MET`, `MST`, `MST7MDT`, `Mexico/BajaNorte`, `Mexico/BajaSur`, `Mexico/General`, `NZ`, `NZ-CHAT`, `Navajo`, `PRC`, `PST8PDT`, `Pacific/Apia`, `Pacific/Auckland`, `Pacific/Bougainville`, `Pacific/Chatham`, `Pacific/Chuuk`, `Pacific/Easter`, `Pacific/Efate`, `Pacific/Enderbury`, `Pacific/Fakaofo`, `Pacific/Fiji`, `Pacific/Funafuti`, `Pacific/Galapagos`, `Pacific/Gambier`,
	`Pacific/Guadalcanal`, `Pacific/Guam`, `Pacific/Honolulu`, `Pacific/Johnston`, `Pacific/Kiritimati`, `Pacific/Kosrae`, `Pacific/Kwajalein`, `Pacific/Majuro`, `Pacific/Marquesas`, `Pacific/Midway`, `Pacific/Nauru`, `Pacific/Niue`, `Pacific/Norfolk`, `Pacific/Noumea`, `Pacific/Pago_Pago`, `Pacific/Palau`, `Pacific/Pitcairn`, `Pacific/Pohnpei`, `Pacific/Ponape`, `Pacific/Port_Moresby`, `Pacific/Rarotonga`, `Pacific/Saipan`, `Pacific/Samoa`, `Pacific/Tahiti`, `Pacific/Tarawa`, `Pacific/Tongatapu`, `Pacific/Truk`, `Pacific/Wake`, `Pacific/Wallis`, `Pacific/Yap`, `Poland`, `Portugal`, `ROC`, `ROK`, `Singapore`, `Turkey`, `UCT`, `US/Alaska`, `US/Aleutian`, `US/Arizona`, `US/Central`, `US/East-Indiana`, `US/Eastern`, `US/Hawaii`, `US/Indiana-Starke`, `US/Michigan`, `US/Mountain`, `US/Pacific`, `US/Pacific-New`, `US/Samoa`, `UTC`, `Universal`, `W-SU`, `WET`, `Zulu`.
	* `type` - (String) Trigger type.
	  * Constraints: Allowable values are: .
	* `worker` - (List) Worker used to run the trigger. If not specified the trigger will use the default pipeline worker.
	Nested scheme for **worker**:
		* `id` - (Forces new resource, String) ID of the worker.
		* `name` - (String) Name of the worker. Computed based on the worker ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_. \\(\\)\\[\\]]{1,235}$/`.
		* `type` - (String) Type of the worker. Computed based on the worker ID.
		  * Constraints: Allowable values are: `private`, `public`.
* `updated_at` - (String) Standard RFC 3339 Date Time String.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_cd_tekton_pipeline` resource by using `id`. UUID.

# Syntax
```
$ terraform import ibm_cd_tekton_pipeline.cd_tekton_pipeline <id>
```
