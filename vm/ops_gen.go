package vm

// Code generated automatically; DO NOT EDIT

const (
	opret                    bcop = 0
	opjz                     bcop = 1
	oploadk                  bcop = 2
	opsavek                  bcop = 3
	opxchgk                  bcop = 4
	oploadb                  bcop = 5
	opsaveb                  bcop = 6
	oploadv                  bcop = 7
	opsavev                  bcop = 8
	oploadzerov              bcop = 9
	opsavezerov              bcop = 10
	oploadpermzerov          bcop = 11
	opsaveblendv             bcop = 12
	oploads                  bcop = 13
	opsaves                  bcop = 14
	oploadzeros              bcop = 15
	opsavezeros              bcop = 16
	opfalse                  bcop = 17
	opandk                   bcop = 18
	opork                    bcop = 19
	opandnotk                bcop = 20
	opnandk                  bcop = 21
	opxork                   bcop = 22
	opnotk                   bcop = 23
	opxnork                  bcop = 24
	opbroadcastimmf          bcop = 25
	opbroadcastimmi          bcop = 26
	opabsf                   bcop = 27
	opabsi                   bcop = 28
	opnegf                   bcop = 29
	opnegi                   bcop = 30
	opsignf                  bcop = 31
	opsigni                  bcop = 32
	opsquaref                bcop = 33
	opsquarei                bcop = 34
	oproundf                 bcop = 35
	oproundevenf             bcop = 36
	optruncf                 bcop = 37
	opfloorf                 bcop = 38
	opceilf                  bcop = 39
	opaddf                   bcop = 40
	opaddimmf                bcop = 41
	opaddi                   bcop = 42
	opaddimmi                bcop = 43
	opsubf                   bcop = 44
	opsubimmf                bcop = 45
	opsubi                   bcop = 46
	opsubimmi                bcop = 47
	oprsubf                  bcop = 48
	oprsubimmf               bcop = 49
	oprsubi                  bcop = 50
	oprsubimmi               bcop = 51
	opmulf                   bcop = 52
	opmulimmf                bcop = 53
	opmuli                   bcop = 54
	opmulimmi                bcop = 55
	opdivf                   bcop = 56
	opdivimmf                bcop = 57
	oprdivf                  bcop = 58
	oprdivimmf               bcop = 59
	opdivi                   bcop = 60
	opdivimmi                bcop = 61
	oprdivi                  bcop = 62
	oprdivimmi               bcop = 63
	opmodf                   bcop = 64
	opmodimmf                bcop = 65
	oprmodf                  bcop = 66
	oprmodimmf               bcop = 67
	opmodi                   bcop = 68
	opmodimmi                bcop = 69
	oprmodi                  bcop = 70
	oprmodimmi               bcop = 71
	opaddmulimmi             bcop = 72
	opminvaluef              bcop = 73
	opminvalueimmf           bcop = 74
	opmaxvaluef              bcop = 75
	opmaxvalueimmf           bcop = 76
	opminvaluei              bcop = 77
	opminvalueimmi           bcop = 78
	opmaxvaluei              bcop = 79
	opmaxvalueimmi           bcop = 80
	opsqrtf                  bcop = 81
	opcbrtf                  bcop = 82
	opexpf                   bcop = 83
	opexp2f                  bcop = 84
	opexp10f                 bcop = 85
	opexpm1f                 bcop = 86
	oplnf                    bcop = 87
	opln1pf                  bcop = 88
	oplog2f                  bcop = 89
	oplog10f                 bcop = 90
	opsinf                   bcop = 91
	opcosf                   bcop = 92
	optanf                   bcop = 93
	opasinf                  bcop = 94
	opacosf                  bcop = 95
	opatanf                  bcop = 96
	opatan2f                 bcop = 97
	ophypotf                 bcop = 98
	oppowf                   bcop = 99
	opcvtktof64              bcop = 100
	opcvtktoi64              bcop = 101
	opcvti64tof64            bcop = 102
	opcvtf64toi64            bcop = 103
	opfproundu               bcop = 104
	opfproundd               bcop = 105
	opcvti64tostr            bcop = 106
	opcmpeqf                 bcop = 107
	opcmpeqi                 bcop = 108
	opcmpeqimmf              bcop = 109
	opcmpeqimmi              bcop = 110
	opcmpltf                 bcop = 111
	opcmplti                 bcop = 112
	opcmpltimmf              bcop = 113
	opcmpltimmi              bcop = 114
	opcmplef                 bcop = 115
	opcmplei                 bcop = 116
	opcmpleimmf              bcop = 117
	opcmpleimmi              bcop = 118
	opcmpgtf                 bcop = 119
	opcmpgti                 bcop = 120
	opcmpgtimmf              bcop = 121
	opcmpgtimmi              bcop = 122
	opcmpgef                 bcop = 123
	opcmpgei                 bcop = 124
	opcmpgeimmf              bcop = 125
	opcmpgeimmi              bcop = 126
	opisnanf                 bcop = 127
	opchecktag               bcop = 128
	opisnull                 bcop = 129
	opisnotnull              bcop = 130
	opistrue                 bcop = 131
	opisfalse                bcop = 132
	opeqslice                bcop = 133
	opequalv                 bcop = 134
	opeqv4mask               bcop = 135
	opeqv4maskplus           bcop = 136
	opeqv8                   bcop = 137
	opeqv8plus               bcop = 138
	opleneq                  bcop = 139
	opdateaddmonth           bcop = 140
	opdateaddmonthimm        bcop = 141
	opdateaddyear            bcop = 142
	opdatediffparam          bcop = 143
	opdatediffmonthyear      bcop = 144
	opdateextractmicrosecond bcop = 145
	opdateextractmillisecond bcop = 146
	opdateextractsecond      bcop = 147
	opdateextractminute      bcop = 148
	opdateextracthour        bcop = 149
	opdateextractday         bcop = 150
	opdateextractmonth       bcop = 151
	opdateextractyear        bcop = 152
	opdatetounixepoch        bcop = 153
	opdatetruncmillisecond   bcop = 154
	opdatetruncsecond        bcop = 155
	opdatetruncminute        bcop = 156
	opdatetrunchour          bcop = 157
	opdatetruncday           bcop = 158
	opdatetruncmonth         bcop = 159
	opdatetruncyear          bcop = 160
	opunboxts                bcop = 161
	opboxts                  bcop = 162
	optimelt                 bcop = 163
	optimegt                 bcop = 164
	opconsttm                bcop = 165
	optmextract              bcop = 166
	opwidthbucketf           bcop = 167
	opwidthbucketi           bcop = 168
	optimebucketts           bcop = 169
	opgeogridi               bcop = 170
	opgeogridimmi            bcop = 171
	opgeohash                bcop = 172
	opgeohashimm             bcop = 173
	opfindsym                bcop = 174
	opfindsym2               bcop = 175
	opfindsym2rev            bcop = 176
	opfindsym3               bcop = 177
	opblendv                 bcop = 178
	opblendrevv              bcop = 179
	opblendnum               bcop = 180
	opblendnumrev            bcop = 181
	opblendslice             bcop = 182
	opblendslicerev          bcop = 183
	opunpack                 bcop = 184
	optoint                  bcop = 185
	optof64                  bcop = 186
	opboxfloat               bcop = 187
	opboxint                 bcop = 188
	opboxmask                bcop = 189
	opboxmask2               bcop = 190
	opboxmask3               bcop = 191
	opboxstring              bcop = 192
	ophashvalue              bcop = 193
	ophashvalueplus          bcop = 194
	ophashmember             bcop = 195
	ophashlookup             bcop = 196
	opaggsumf                bcop = 197
	opaggsumi                bcop = 198
	opaggminf                bcop = 199
	opaggmini                bcop = 200
	opaggmaxf                bcop = 201
	opaggmaxi                bcop = 202
	opaggcount               bcop = 203
	opaggbucket              bcop = 204
	opaggslotaddf            bcop = 205
	opaggslotaddi            bcop = 206
	opaggslotavgf            bcop = 207
	opaggslotavgi            bcop = 208
	opaggslotminf            bcop = 209
	opaggslotmini            bcop = 210
	opaggslotmaxf            bcop = 211
	opaggslotmaxi            bcop = 212
	opaggslotcount           bcop = 213
	oplitref                 bcop = 214
	opsplit                  bcop = 215
	optuple                  bcop = 216
	opdupv                   bcop = 217
	opzerov                  bcop = 218
	opobjectsize             bcop = 219
	opCmpStrEqCs             bcop = 220
	opCmpStrEqCi             bcop = 221
	opCmpStrEqUTF8Ci         bcop = 222
	opSkip1charLeft          bcop = 223
	opSkip1charRight         bcop = 224
	opSkipNcharLeft          bcop = 225
	opSkipNcharRight         bcop = 226
	opTrimWsLeft             bcop = 227
	opTrimWsRight            bcop = 228
	opTrim4charLeft          bcop = 229
	opTrim4charRight         bcop = 230
	opTrimPrefixCs           bcop = 231
	opTrimPrefixCi           bcop = 232
	opTrimSuffixCs           bcop = 233
	opTrimSuffixCi           bcop = 234
	opContainsSubstrCs       bcop = 235
	opContainsSubstrCi       bcop = 236
	opContainsSuffixCs       bcop = 237
	opContainsSuffixCi       bcop = 238
	opContainsSuffixUTF8Ci   bcop = 239
	opContainsPrefixCs       bcop = 240
	opContainsPrefixCi       bcop = 241
	opContainsPrefixUTF8Ci   bcop = 242
	opLengthStr              bcop = 243
	opSubstr                 bcop = 244
	opSplitPart              bcop = 245
	opMatchpatCs             bcop = 246
	opMatchpatCi             bcop = 247
	opMatchpatUTF8Ci         bcop = 248
	opIsSubnetOfIP4          bcop = 249
	optrap                   bcop = 250
	_maxbcop                      = 251
)
