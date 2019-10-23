package registryv1

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	ibmcloud "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	ibmcloudHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	dockerfileName = "Dockerfile"
	dockerfile     = `FROM golang:1.7-alpine3.6

ARG JQ_VERSION=jq-1.5
RUN apk update
RUN apk add build-base git bash
ADD https://github.com/stedolan/jq/releases/download/${JQ_VERSION}/jq-linux64 /tmp`

	buildResult = `{"stream":"Step 1/5 : FROM golang:1.7-alpine3.6\n"}
{"stream":" ---\u003e 0cf3d3497ae9\n"}
{"stream":"Step 2/5 : ARG JQ_VERSION=jq-1.5\n"}
{"stream":" ---\u003e Running in 1423169fe389\n"}
{"stream":" ---\u003e 28032d919806\n"}
{"stream":"Removing intermediate container 1423169fe389\n"}
{"stream":"Step 3/5 : RUN apk update\n"}
{"stream":" ---\u003e Running in 3be23ccfacb4\n"}
{"stream":"fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/main/x86_64/APKINDEX.tar.gz\n"}
{"stream":"fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/community/x86_64/APKINDEX.tar.gz\n"}
{"stream":"v3.6.3-20-g059e2ef678 [http://dl-cdn.alpinelinux.org/alpine/v3.6/main]\nv3.6.3-21-g0a5db24a90 [http://dl-cdn.alpinelinux.org/alpine/v3.6/community]\nOK: 8445 distinct packages available\n"}
{"stream":" ---\u003e 1eb8d32cae96\n"}
{"stream":"Removing intermediate container 3be23ccfacb4\n"}
{"stream":"Step 4/5 : RUN apk add build-base git bash\n"}
{"stream":" ---\u003e Running in 723a2e9f277a\n"}
{"stream":"(1/30) Upgrading musl (1.1.16-r10 -\u003e 1.1.16-r14)\n"}
{"stream":"(2/30) Installing ncurses-terminfo-base (6.0_p20171125-r1)\n"}
{"stream":"(3/30) Installing ncurses-terminfo (6.0_p20171125-r1)\n"}
{"stream":"(4/30) Installing ncurses-libs (6.0_p20171125-r1)\n"}
{"stream":"(5/30) Installing readline (6.3.008-r5)\n"}
{"stream":"(6/30) Installing bash (4.3.48-r1)\n"}
{"stream":"Executing bash-4.3.48-r1.post-install\n"}
{"stream":"(7/30) Installing binutils-libs (2.30-r1)\n"}
{"stream":"(8/30) Installing binutils (2.30-r1)\n"}
{"stream":"(9/30) Installing gmp (6.1.2-r0)\n"}
{"stream":"(10/30) Installing isl (0.17.1-r0)\n"}
{"stream":"(11/30) Installing libgomp (6.3.0-r4)\n"}
{"stream":"(12/30) Installing libatomic (6.3.0-r4)\n"}
{"stream":"(13/30) Installing pkgconf (1.3.7-r0)\n"}
{"stream":"(14/30) Installing libgcc (6.3.0-r4)\n"}
{"stream":"(15/30) Installing mpfr3 (3.1.5-r0)\n"}
{"stream":"(16/30) Installing mpc1 (1.0.3-r0)\n"}
{"stream":"(17/30) Installing libstdc++ (6.3.0-r4)\n"}
{"stream":"(18/30) Installing gcc (6.3.0-r4)\n"}
{"stream":"(19/30) Installing musl-dev (1.1.16-r14)\n"}
{"stream":"(20/30) Installing libc-dev (0.7.1-r0)\n"}
{"stream":"(21/30) Installing g++ (6.3.0-r4)\n"}
{"stream":"(22/30) Installing make (4.2.1-r0)\n"}
{"stream":"(23/30) Installing fortify-headers (0.8-r0)\n"}
{"stream":"(24/30) Installing build-base (0.5-r0)\n"}
{"stream":"(25/30) Installing libssh2 (1.8.0-r1)\n"}
{"stream":"(26/30) Installing libcurl (7.61.1-r0)\n"}
{"stream":"(27/30) Installing expat (2.2.0-r1)\n"}
{"stream":"(28/30) Installing pcre (8.41-r0)\n"}
{"stream":"(29/30) Installing git (2.13.7-r0)\n"}
{"stream":"(30/30) Upgrading musl-utils (1.1.16-r10 -\u003e 1.1.16-r14)\n"}
{"stream":"Executing busybox-1.26.2-r5.trigger\n"}
{"stream":"OK: 190 MiB in 40 packages\n"}
{"stream":" ---\u003e 7a5a01b469dd\n"}
{"stream":"Removing intermediate container 723a2e9f277a\n"}
{"stream":"Step 5/5 : ADD https://github.com/stedolan/jq/releases/download/${JQ_VERSION}/jq-linux64 /tmp\n"}
{"status":"Downloading","progressDetail":{"current":33339,"total":3027945},"progress":"[\u003e                                                  ]  33.34kB/3.028MB"}
{"status":"Downloading","progressDetail":{"current":451131,"total":3027945},"progress":"[=======\u003e                                           ]  451.1kB/3.028MB"}
{"status":"Downloading","progressDetail":{"current":3027945,"total":3027945},"progress":"[==================================================\u003e]  3.028MB/3.028MB"}
{"stream":"\n"}
{"stream":" ---\u003e c8763987f48f\n"}
{"aux":{"ID":"sha256:c8763987f48f13fb23ec84dded2b241b79149548ab10e051dc96b772b9c212a7"}}
{"stream":"Successfully built c8763987f48f\n"}
{"stream":"Successfully tagged registry.ng.bluemix.net/bkuschel/testimage:latest\n"}
{"status":"The push refers to a repository [registry.ng.bluemix.net/bkuschel/testimage]"}
{"status":"Preparing","progressDetail":{},"id":"54bdabef2225"}
{"status":"Preparing","progressDetail":{},"id":"360ebdf19ecf"}
{"status":"Preparing","progressDetail":{},"id":"8e69d607c320"}
{"status":"Preparing","progressDetail":{},"id":"e5b65e634c73"}
{"status":"Preparing","progressDetail":{},"id":"afc52f56690e"}
{"status":"Preparing","progressDetail":{},"id":"e454a04d0a3e"}
{"status":"Preparing","progressDetail":{},"id":"fc3d79f9e82c"}
{"status":"Preparing","progressDetail":{},"id":"069301b5b9f1"}
{"status":"Preparing","progressDetail":{},"id":"5bef08742407"}
{"status":"Waiting","progressDetail":{},"id":"e454a04d0a3e"}
{"status":"Waiting","progressDetail":{},"id":"fc3d79f9e82c"}
{"status":"Waiting","progressDetail":{},"id":"069301b5b9f1"}
{"status":"Waiting","progressDetail":{},"id":"5bef08742407"}
{"status":"Layer already exists","progressDetail":{},"id":"afc52f56690e"}
{"status":"Layer already exists","progressDetail":{},"id":"e5b65e634c73"}
{"status":"Pushing","progressDetail":{"current":33792,"total":3027945},"progress":"[\u003e                                                  ]  33.79kB/3.028MB","id":"54bdabef2225"}
{"status":"Pushing","progressDetail":{"current":36864,"total":1102904},"progress":"[=\u003e                                                 ]  36.86kB/1.103MB","id":"8e69d607c320"}
{"status":"Layer already exists","progressDetail":{},"id":"fc3d79f9e82c"}
{"status":"Pushing","progressDetail":{"current":525312,"total":173229523},"progress":"[\u003e                                                  ]  525.3kB/173.2MB","id":"360ebdf19ecf"}
{"status":"Layer already exists","progressDetail":{},"id":"e454a04d0a3e"}
{"status":"Pushing","progressDetail":{"current":1108992,"total":1102904},"progress":"[==================================================\u003e]  1.109MB","id":"8e69d607c320"}
{"status":"Pushing","progressDetail":{"current":1246208,"total":3027945},"progress":"[====================\u003e                              ]  1.246MB/3.028MB","id":"54bdabef2225"}
{"status":"Layer already exists","progressDetail":{},"id":"069301b5b9f1"}
{"status":"Layer already exists","progressDetail":{},"id":"5bef08742407"}
{"status":"Pushing","progressDetail":{"current":2589696,"total":3027945},"progress":"[==========================================\u003e        ]   2.59MB/3.028MB","id":"54bdabef2225"}
{"status":"Pushing","progressDetail":{"current":2166272,"total":173229523},"progress":"[\u003e                                                  ]  2.166MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":3030016,"total":3027945},"progress":"[==================================================\u003e]   3.03MB","id":"54bdabef2225"}
{"status":"Pushing","progressDetail":{"current":4338688,"total":173229523},"progress":"[=\u003e                                                 ]  4.339MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":5999616,"total":173229523},"progress":"[=\u003e                                                 ]      6MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":7622656,"total":173229523},"progress":"[==\u003e                                                ]  7.623MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":9250816,"total":173229523},"progress":"[==\u003e                                                ]  9.251MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushed","progressDetail":{},"id":"8e69d607c320"}
{"status":"Pushing","progressDetail":{"current":10364928,"total":173229523},"progress":"[==\u003e                                                ]  10.36MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushed","progressDetail":{},"id":"54bdabef2225"}
{"status":"Pushing","progressDetail":{"current":11465728,"total":173229523},"progress":"[===\u003e                                               ]  11.47MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":13095424,"total":173229523},"progress":"[===\u003e                                               ]   13.1MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":15811072,"total":173229523},"progress":"[====\u003e                                              ]  15.81MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":17427291,"total":173229523},"progress":"[=====\u003e                                             ]  17.43MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":19612160,"total":173229523},"progress":"[=====\u003e                                             ]  19.61MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":21743658,"total":173229523},"progress":"[======\u003e                                            ]  21.74MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":23860185,"total":173229523},"progress":"[======\u003e                                            ]  23.86MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":25434293,"total":173229523},"progress":"[=======\u003e                                           ]  25.43MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":26485290,"total":173229523},"progress":"[=======\u003e                                           ]  26.49MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":28063323,"total":173229523},"progress":"[========\u003e                                          ]  28.06MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":29112723,"total":173229523},"progress":"[========\u003e                                          ]  29.11MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":30162944,"total":173229523},"progress":"[========\u003e                                          ]  30.16MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":31212257,"total":173229523},"progress":"[=========\u003e                                         ]  31.21MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":32786627,"total":173229523},"progress":"[=========\u003e                                         ]  32.79MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":34936132,"total":173229523},"progress":"[==========\u003e                                        ]  34.94MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":36571136,"total":173229523},"progress":"[==========\u003e                                        ]  36.57MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":39271936,"total":173229523},"progress":"[===========\u003e                                       ]  39.27MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":40943104,"total":173229523},"progress":"[===========\u003e                                       ]  40.94MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":43131480,"total":173229523},"progress":"[============\u003e                                      ]  43.13MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":44725320,"total":173229523},"progress":"[============\u003e                                      ]  44.73MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":46927360,"total":173229523},"progress":"[=============\u003e                                     ]  46.93MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":48524496,"total":173229523},"progress":"[==============\u003e                                    ]  48.52MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":50132480,"total":173229523},"progress":"[==============\u003e                                    ]  50.13MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":51774976,"total":173229523},"progress":"[==============\u003e                                    ]  51.77MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":53446144,"total":173229523},"progress":"[===============\u003e                                   ]  53.45MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":55117312,"total":173229523},"progress":"[===============\u003e                                   ]  55.12MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":56788480,"total":173229523},"progress":"[================\u003e                                  ]  56.79MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":58459648,"total":173229523},"progress":"[================\u003e                                  ]  58.46MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":60130816,"total":173229523},"progress":"[=================\u003e                                 ]  60.13MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":61772288,"total":173229523},"progress":"[=================\u003e                                 ]  61.77MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":63387648,"total":173229523},"progress":"[==================\u003e                                ]  63.39MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":64472064,"total":173229523},"progress":"[==================\u003e                                ]  64.47MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":66118656,"total":173229523},"progress":"[===================\u003e                               ]  66.12MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":68294144,"total":173229523},"progress":"[===================\u003e                               ]  68.29MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":69965312,"total":173229523},"progress":"[====================\u003e                              ]  69.97MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":71636480,"total":173229523},"progress":"[====================\u003e                              ]  71.64MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":73307648,"total":173229523},"progress":"[=====================\u003e                             ]  73.31MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":74978816,"total":173229523},"progress":"[=====================\u003e                             ]  74.98MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":76649984,"total":173229523},"progress":"[======================\u003e                            ]  76.65MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":78321152,"total":173229523},"progress":"[======================\u003e                            ]  78.32MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":79992320,"total":173229523},"progress":"[=======================\u003e                           ]  79.99MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":81645056,"total":173229523},"progress":"[=======================\u003e                           ]  81.65MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":83311104,"total":173229523},"progress":"[========================\u003e                          ]  83.31MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":84982272,"total":173229523},"progress":"[========================\u003e                          ]  84.98MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":86639616,"total":173229523},"progress":"[=========================\u003e                         ]  86.64MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":87745536,"total":173229523},"progress":"[=========================\u003e                         ]  87.75MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":89416704,"total":173229523},"progress":"[=========================\u003e                         ]  89.42MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":91087872,"total":173229523},"progress":"[==========================\u003e                        ]  91.09MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":92759040,"total":173229523},"progress":"[==========================\u003e                        ]  92.76MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":94430208,"total":173229523},"progress":"[===========================\u003e                       ]  94.43MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":96101376,"total":173229523},"progress":"[===========================\u003e                       ]   96.1MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":97215488,"total":173229523},"progress":"[============================\u003e                      ]  97.22MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":98886656,"total":173229523},"progress":"[============================\u003e                      ]  98.89MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":100557824,"total":173229523},"progress":"[=============================\u003e                     ]  100.6MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":102228992,"total":173229523},"progress":"[=============================\u003e                     ]  102.2MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":104438272,"total":173229523},"progress":"[==============================\u003e                    ]  104.4MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":105552384,"total":173229523},"progress":"[==============================\u003e                    ]  105.6MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":107223552,"total":173229523},"progress":"[==============================\u003e                    ]  107.2MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":108894720,"total":173229523},"progress":"[===============================\u003e                   ]  108.9MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":110565888,"total":173229523},"progress":"[===============================\u003e                   ]  110.6MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":112237056,"total":173229523},"progress":"[================================\u003e                  ]  112.2MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":113908224,"total":173229523},"progress":"[================================\u003e                  ]  113.9MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":115022336,"total":173229523},"progress":"[=================================\u003e                 ]    115MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":116693504,"total":173229523},"progress":"[=================================\u003e                 ]  116.7MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":118364672,"total":173229523},"progress":"[==================================\u003e                ]  118.4MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":120035840,"total":173229523},"progress":"[==================================\u003e                ]    120MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":121707008,"total":173229523},"progress":"[===================================\u003e               ]  121.7MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":122793984,"total":173229523},"progress":"[===================================\u003e               ]  122.8MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":124465152,"total":173229523},"progress":"[===================================\u003e               ]  124.5MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":125579264,"total":173229523},"progress":"[====================================\u003e              ]  125.6MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":126693376,"total":173229523},"progress":"[====================================\u003e              ]  126.7MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":127807488,"total":173229523},"progress":"[====================================\u003e              ]  127.8MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":128921600,"total":173229523},"progress":"[=====================================\u003e             ]  128.9MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":130592768,"total":173229523},"progress":"[=====================================\u003e             ]  130.6MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":132263936,"total":173229523},"progress":"[======================================\u003e            ]  132.3MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":133378048,"total":173229523},"progress":"[======================================\u003e            ]  133.4MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":135049216,"total":173229523},"progress":"[======================================\u003e            ]    135MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":136720384,"total":173229523},"progress":"[=======================================\u003e           ]  136.7MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":138391552,"total":173229523},"progress":"[=======================================\u003e           ]  138.4MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":140062720,"total":173229523},"progress":"[========================================\u003e          ]  140.1MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":141689344,"total":173229523},"progress":"[========================================\u003e          ]  141.7MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":142794240,"total":173229523},"progress":"[=========================================\u003e         ]  142.8MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":144465408,"total":173229523},"progress":"[=========================================\u003e         ]  144.5MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":146136576,"total":173229523},"progress":"[==========================================\u003e        ]  146.1MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":147807744,"total":173229523},"progress":"[==========================================\u003e        ]  147.8MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":149478912,"total":173229523},"progress":"[===========================================\u003e       ]  149.5MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":151150080,"total":173229523},"progress":"[===========================================\u003e       ]  151.2MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":152264192,"total":173229523},"progress":"[===========================================\u003e       ]  152.3MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":153935360,"total":173229523},"progress":"[============================================\u003e      ]  153.9MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":155606528,"total":173229523},"progress":"[============================================\u003e      ]  155.6MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":157834752,"total":173229523},"progress":"[=============================================\u003e     ]  157.8MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":159503360,"total":173229523},"progress":"[==============================================\u003e    ]  159.5MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":161146880,"total":173229523},"progress":"[==============================================\u003e    ]  161.1MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":162766336,"total":173229523},"progress":"[==============================================\u003e    ]  162.8MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":163867648,"total":173229523},"progress":"[===============================================\u003e   ]  163.9MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":165516800,"total":173229523},"progress":"[===============================================\u003e   ]  165.5MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":167149056,"total":173229523},"progress":"[================================================\u003e  ]  167.1MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":168793600,"total":173229523},"progress":"[================================================\u003e  ]  168.8MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":170412544,"total":173229523},"progress":"[=================================================\u003e ]  170.4MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":172031488,"total":173229523},"progress":"[=================================================\u003e ]    172MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":173113232,"total":173229523},"progress":"[=================================================\u003e ]  173.1MB/173.2MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":174732532,"total":173229523},"progress":"[==================================================\u003e]  174.7MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":175783582,"total":173229523},"progress":"[==================================================\u003e]  175.8MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":176833526,"total":173229523},"progress":"[==================================================\u003e]  176.8MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":177884111,"total":173229523},"progress":"[==================================================\u003e]  177.9MB","id":"360ebdf19ecf"}
{"status":"Pushing","progressDetail":{"current":179670528,"total":173229523},"progress":"[==================================================\u003e]  179.7MB","id":"360ebdf19ecf"}
{"status":"Pushed","progressDetail":{},"id":"360ebdf19ecf"}
{"status":"latest: digest: sha256:a93d0334db5fd9769e1d09adcd13eafb690b170fe294cf0aa806273db5e2bcaf size: 2207"}
{"progressDetail":{},"aux":{"Tag":"latest","Digest":"sha256:a93d0334db5fd9769e1d09adcd13eafb690b170fe294cf0aa806273db5e2bcaf","Size":"2207"}}`

	autoErrorResult = `{
	"code": "CRG0009E",
	"message": "You are not authorized to access the specified account.",
	"request-id": "1044-1540315282.640-825591"
}`

	buildErrorResult = `{"errorDetail":{"message":"Error response from daemon: Cannot locate specified Dockerfile: Dockerfiles"},"error":"Error response from daemon: Cannot locate specified Dockerfile: Dockerfiles"}`
)

var _ = Describe("Builds", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("ImageBuildCallback", func() {
		Context("When build with callback is completed", func() {
			tarBuffer := createTestTar()
			var requestBuffer bytes.Buffer
			buffer := io.TeeReader(tarBuffer, &requestBuffer)
			bodyBytes, _ := ioutil.ReadAll(buffer)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/builds"),
						ghttp.VerifyBody(bodyBytes),
						ghttp.RespondWith(http.StatusOK, buildResult),
					),
				)
			})

			It("should return build with callback results", func() {
				params := ImageBuildRequest{
					T:          "registry.ng.bluemix.net/bkuschel/testimage",
					Dockerfile: "",
					Buildargs:  "",
					Nocache:    false,
					Pull:       false,
					Quiet:      false,
					Squash:     false,
				}
				target := BuildTargetHeader{
					AccountID: "abc",
				}
				i := 0
				respArr := make([]ImageBuildResponse, i, 203)
				err := newBuild(server.URL()).ImageBuildCallback(params, &requestBuffer, target, func(respV ImageBuildResponse) bool {
					respArr = append(respArr, respV)
					i++
					return true
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(respArr).To(HaveLen(203))
				for i, v := range strings.Split(buildResult, "\n") {
					var resp ImageBuildResponse
					err = json.Unmarshal([]byte(v), &resp)
					Expect(err).To(BeNil())
					Expect(respArr[i]).Should(Equal(resp))
				}
			})
		})
		Context("When build with callback auth is unsuccessful", func() {
			tarBuffer := createTestTar()
			var requestBuffer bytes.Buffer
			buffer := io.TeeReader(tarBuffer, &requestBuffer)
			bodyBytes, _ := ioutil.ReadAll(buffer)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/builds"),
						ghttp.VerifyBody(bodyBytes),
						ghttp.RespondWith(http.StatusUnauthorized, autoErrorResult),
					),
				)
			})

			It("should return error during build with callback ", func() {
				params := ImageBuildRequest{
					T:          "registry.ng.bluemix.net/bkuschel/testimage",
					Dockerfile: "",
					Buildargs:  "",
					Nocache:    false,
					Pull:       false,
					Quiet:      false,
					Squash:     false,
				}
				target := BuildTargetHeader{
					AccountID: "abc",
				}
				i := 0
				respArr := make([]ImageBuildResponse, i, 203)
				err := newBuild(server.URL()).ImageBuildCallback(params, &requestBuffer, target, func(respV ImageBuildResponse) bool {
					respArr = append(respArr, respV)
					i++
					return true
				})
				Expect(err).To(HaveOccurred())
				Expect(respArr).To(HaveLen(i))
			})
		})
		Context("When build with callback is unsuccessful", func() {
			tarBuffer := createTestTar()
			var requestBuffer bytes.Buffer
			buffer := io.TeeReader(tarBuffer, &requestBuffer)
			bodyBytes, _ := ioutil.ReadAll(buffer)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/builds"),
						ghttp.VerifyBody(bodyBytes),
						ghttp.RespondWith(http.StatusOK, buildErrorResult),
					),
				)
			})

			It("should return error during build with callback", func() {
				params := ImageBuildRequest{
					T:          "registry.ng.bluemix.net/bkuschel/testimage",
					Dockerfile: "Dockerfiles",
					Buildargs:  "",
					Nocache:    false,
					Pull:       false,
					Quiet:      false,
					Squash:     false,
				}
				target := BuildTargetHeader{
					AccountID: "abc",
				}
				i := 0
				respArr := make([]ImageBuildResponse, i, 203)
				err := newBuild(server.URL()).ImageBuildCallback(params, &requestBuffer, target, func(respV ImageBuildResponse) bool {
					respArr = append(respArr, respV)
					i++
					return true
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(respArr).To(HaveLen(1))
				var resp ImageBuildResponse
				err = json.Unmarshal([]byte(buildErrorResult), &resp)
				Expect(err).To(BeNil())
				Expect(respArr[0]).Should(Equal(resp))
			})
		})
	})
	Describe("ImageBuild", func() {
		Context("When Build is completed", func() {
			tarBuffer := createTestTar()
			var requestBuffer bytes.Buffer
			buffer := io.TeeReader(tarBuffer, &requestBuffer)
			bodyBytes, _ := ioutil.ReadAll(buffer)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/builds"),
						ghttp.VerifyBody(bodyBytes),
						ghttp.RespondWith(http.StatusOK, buildResult),
					),
				)
			})

			It("should return build results", func() {
				params := ImageBuildRequest{
					T:          "registry.ng.bluemix.net/bkuschel/testimage",
					Dockerfile: "",
					Buildargs:  "",
					Nocache:    false,
					Pull:       false,
					Quiet:      false,
					Squash:     false,
				}
				target := BuildTargetHeader{
					AccountID: "abc",
				}
				var b bytes.Buffer
				out := bufio.NewWriter(&b)
				err := newBuild(server.URL()).ImageBuild(params, &requestBuffer, target, out)
				Expect(err).NotTo(HaveOccurred())
				Expect(out).NotTo(BeNil())
				Expect(b.Bytes()).Should(Equal([]byte(buildResult)))
			})
		})
		Context("When build auth is unsuccessful", func() {
			tarBuffer := createTestTar()
			var requestBuffer bytes.Buffer
			buffer := io.TeeReader(tarBuffer, &requestBuffer)
			bodyBytes, _ := ioutil.ReadAll(buffer)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/builds"),
						ghttp.VerifyBody(bodyBytes),
						ghttp.RespondWith(http.StatusUnauthorized, autoErrorResult),
					),
				)
			})

			It("should return error during build", func() {
				params := ImageBuildRequest{
					T:          "registry.ng.bluemix.net/bkuschel/testimage",
					Dockerfile: "",
					Buildargs:  "",
					Nocache:    false,
					Pull:       false,
					Quiet:      false,
					Squash:     false,
				}
				target := BuildTargetHeader{
					AccountID: "abc",
				}
				var b bytes.Buffer
				out := bufio.NewWriter(&b)
				err := newBuild(server.URL()).ImageBuild(params, &requestBuffer, target, out)
				Expect(err).To(HaveOccurred())
				Expect(out).NotTo(BeNil())
			})
		})
		Context("When build is unsuccessful", func() {
			tarBuffer := createTestTar()
			var requestBuffer bytes.Buffer
			buffer := io.TeeReader(tarBuffer, &requestBuffer)
			bodyBytes, _ := ioutil.ReadAll(buffer)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/builds"),
						ghttp.VerifyBody(bodyBytes),
						ghttp.RespondWith(http.StatusOK, buildErrorResult),
					),
				)
			})

			It("should return error during build", func() {
				params := ImageBuildRequest{
					T:          "registry.ng.bluemix.net/bkuschel/testimage",
					Dockerfile: "Dockerfiles",
					Buildargs:  "",
					Nocache:    false,
					Pull:       false,
					Quiet:      false,
					Squash:     false,
				}
				target := BuildTargetHeader{
					AccountID: "abc",
				}
				var b bytes.Buffer
				out := bufio.NewWriter(&b)
				err := newBuild(server.URL()).ImageBuild(params, &requestBuffer, target, out)
				Expect(err).NotTo(HaveOccurred())
				Expect(out).NotTo(BeNil())
				Expect(b.Bytes()).Should(Equal([]byte(buildErrorResult)))
			})
		})
	})
})

func newBuild(url string) Builds {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = ibmcloudHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: ibmcloud.ContainerRegistryService,
	}
	return newBuildAPI(&client)
}

func createTestTar() *bytes.Buffer {
	// Create and add some files to the archive.
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	hdr := &tar.Header{
		Name: dockerfileName,
		Mode: 0600,
		Size: int64(len(dockerfile)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		log.Fatal(err)
	}
	if _, err := tw.Write([]byte(dockerfile)); err != nil {
		log.Fatal(err)
	}
	return &buf
}
