# Sort URL List

URLの一覧が記載されたXMLファイルがあります。

それを読み取って、URLを逆順にして、ソートして表示します。

<br><br>

## 依存パッケージ

radix treeを使います。

https://github.com/takamitsu-iida/radix

<br><br>

## XMLファイルの構造

このような構造をしています。

```xml
<?xml version="1.0" encoding="utf-8"?>
<domainlist>
  <service name="MS-Apps">
    <dst>
      <url>outlook.office.com</url>
      <url>outlook.office365.com</url>
      <url>*.protection.outlook.com</url>
      <url>*.lync.com</url>
      <url>*.teams.microsoft.com</url>
      <url>teams.microsoft.com</url>
    </dst>
  </service>
</domainlist>
```

<br><br>

## 仕組み

XMLファイルから読み取ったURLを逆順にします。
たとえば、`outlook.office.com` であれば `com.office.outlook` というようにします。

これをradix treeに格納します。
radix treeでは辞書順にソートされていますので、ツリーをルートから辿っていけばソートされて表示されます。

radix treeに格納する際には同じキーがすでに存在するかわかりますので、重複したエントリの存在もわかります。

<br><br><br><br>

## 以下、実行結果

数字はXMLファイルで読み取った順番です。

ワイルドカード（*）の解釈次第では、不要なエントリも散見します。
たとえば、この二つ。

```text
com.microsoft.* 99
com.microsoft.admin 245
```

`com.microsoft.admin` は `com.microsoft.*` にも一致します。
ロンゲストマッチ方式で挙動を変えたいのであれば意味を持ちますが、単純にフィルタリングするだけなら `com.microsoft.admin` は不要です。

```bash
iida@s400win:~/go/src/sort-url$ go run main.go
ai.cortana.* 244
com.aadrm.* 144
com.acompli.www 179
com.aspnetcdn.ajax 164
com.assets-yammer.* 193
com.azure.informationprotection.* 146
com.azure.management 115
com.azurerms.* 145
com.bing.appex.* 180
com.bing.c 182
com.bing.platform.officespeech 227
com.bing.www 173
com.cloudappsecurity.portal.* 117
com.digicert.cacerts 210
com.digicert.crl3 215
com.digicert.crl4 216
com.digicert.ocsp 219
com.digicert.ocspx 223
com.digicert.www 225
com.geotrust.* 202
com.getmicrosoftkey.partnerservices 187
com.globalsign.crl 212
com.globalsign.ocsp 220
com.globalsign.ocsp2 222
com.globalsign.secure 224
com.identrust.apps 209
com.identrust.crl 214
com.identrust.ocsp.trustid.isrg 217
com.linkedin.platform 135
com.live.account 177
com.live.c 183
com.live.g 83
com.live.hip.* 106
com.live.login 190
com.live.office 12
com.live.officeapps 104
com.live.officeapps.* 10
com.live.officeapps.odc.cdn 246
com.live.officeapps.uci.cdn 247
com.live.services.directory 185
com.live.signup 188
com.live.ssw 75
com.live.storage 76
com.live.storage.users.* 68
com.lync.* 3
com.microsoft.* 99
com.microsoft.admin 245
com.microsoft.appsforoffice 128
com.microsoft.aria.* 124
com.microsoft.auth.* 13
com.microsoft.c1 131
com.microsoft.compass-ssl 71
com.microsoft.compliance 41
com.microsoft.compliance.* 38
com.microsoft.crl 159
com.microsoft.data.events.* 125
com.microsoft.defender 42
com.microsoft.docs 133
com.microsoft.flow.* 238
com.microsoft.go 163
com.microsoft.graph 26
com.microsoft.login 28
com.microsoft.manage.* 230
com.microsoft.mscrl 218
com.microsoft.msdn 134
com.microsoft.myanalytics 233
com.microsoft.myanalytics-gcc 234
com.microsoft.myanalytics.cdnprod 232
com.microsoft.office.r 157
com.microsoft.office15client 160
com.microsoft.officecdn 48
com.microsoft.officeclient 161
com.microsoft.officepreviewredir 155
com.microsoft.officeredir 156
com.microsoft.officeredir.o15 154
com.microsoft.security 44
com.microsoft.security.* 40
com.microsoft.sls.activation 158
com.microsoft.support 139
com.microsoft.support.dgps 132
com.microsoft.teams 5
com.microsoft.teams.* 4
com.microsoft.technet 140
com.microsoft.www 226
com.microsoftazuread-sso.autologon 20
com.microsoftonline-p.* 107
com.microsoftonline-p.aadcdn.secure.* 54
com.microsoftonline-p.login 30
com.microsoftonline-p.nexus 35
com.microsoftonline.* 50
com.microsoftonline.adminwebservice 18
com.microsoftonline.becws 21
com.microsoftonline.companymanager 24
com.microsoftonline.login 29
com.microsoftonline.login-us 34
com.microsoftonline.login.* 53
com.microsoftonline.login.ccs 22
com.microsoftonline.login.device 25
com.microsoftonline.logincert 32
com.microsoftonline.loginex 33
com.microsoftonline.passwordreset 36
com.microsoftonline.passwordreset.api 19
com.microsoftonline.provisioningapi 37
com.microsoftstream.* 92
com.microsoftusercontent.* 236
com.msftidentity.* 14
com.msidentity.* 15
com.msn.appex-rf.* 181
com.msocdn.* 152
com.msocdn.prod 136
com.msocdn.shellprod 137
com.msocsp.ocsp 221
com.o365weve.* 126
com.oaspapps.firstpartyapps 120
com.oaspapps.firstpartyapps.excelbingmap 168
com.oaspapps.firstpartyapps.peoplegraph 170
com.oaspapps.firstpartyapps.telemetryservice 122
com.oaspapps.firstpartyapps.wikipedia 172
com.oaspapps.wus-firstpartyapps 123
com.office.* 231
com.office.online.* 11
com.office.outlook 0
com.office.outlook.* 59
com.office.protection 43
com.office.protection.* 39
com.office365.* 49
com.office365.outlook 1
com.office365.res.r1 56
com.office365.res.r3 57
com.office365.res.r4 58
com.office365.smtp 55
com.omniroot.* 203
com.onedrive.admin 81
com.onedrive.www 105
com.onenote.* 98
com.onmicrosoft.* 153
com.onmicrosoft.*.autodiscover 240
com.outlook.* 60
com.outlook.protection.* 2
com.outlook.protection.mail.* 62
com.outlook.www 194
com.outlookmobile.* 175
com.powerapps.* 239
com.public-trust.* 204
com.sfbassets.* 63
com.sharepoint.* 9
com.sharepoint.*-admin 89
com.sharepoint.*-files 90
com.sharepoint.*-myfiles 91
com.sharepoint.aa-rt.global.gr.* 87
com.sharepointonline.* 85
com.skype.* 228
com.skype.broadcast 7
com.skype.broadcast.* 6
com.skypeassets.secure.* 73
com.skypeforbusiness.* 8
com.sway 199
com.sway-cdn.eus-www 195
com.sway-cdn.wus-www 197
com.sway-extensions.eus-www 196
com.sway-extensions.wus-www 198
com.sway.www 200
com.symcb.* 205
com.symcd.* 206
com.verisign.* 207
com.visualstudio.services.dc 149
com.windows.activity 242
com.windows.activity.* 241
com.windows.adl.* 69
com.windows.wns.* 80
com.windowsazure.activedirectory.account 16
com.yammer.* 191
com.yammerusercontent.* 192
ms.aka 67
ms.gfx.auth 189
ms.gfx.mem 150
ms.mstea.* 72
ms.onestore.assets 129
ms.sfx.oneclient 84
ms.staffhub 151
ms.svc.* 88
net.acompli.* 174
net.akadns.com.oaspapps.firstpartyapps.prod 121
net.akamaihd.spoprod-a 86
net.azure-apim.* 237
net.azure.amp 127
net.azure.media.* 95
net.azure.onyx.nps 93
net.azure.portal.hosting.informationprotection 148
net.azureedge.* 94
net.azureedge.mlccdnprod 74
net.bing.c 167
net.bing.mm.tse1 171
net.cloudapp.* 143
net.edgesuite.com.microsoft.officecdn 47
net.entrust.* 201
net.globalsign.crl 213
net.live.apis 103
net.live.docs 186
net.live.docs.d 184
net.microsoftonline-p.clientconfig 23
net.msauth.* 108
net.msauthimages.* 109
net.msecnd.* 110
net.msedge.* 70
net.msedge.ocos-office365-s2s 169
net.msft.ad.dc.policykeyservice 116
net.msftauth.* 111
net.msftauthimages.* 112
net.office.* 100
net.office.account 45
net.office.attachments 61
net.office.bing.insertmedia 162
net.office.cdn.* 96
net.office.cdn.teams.statics 46
net.office.cdn.workplaceanalytics 235
net.office.config.* 229
net.office.content.support 138
net.office.osi.contentstorage 97
net.office.osi.videocontent 141
net.office.osi.videoplayercdn 142
net.office.suite 118
net.onenote.*cdn 101
net.phonefactor.* 113
net.trafficmanager.apac.production.search.* 77
net.trafficmanager.emea.production.search.* 78
net.trafficmanager.us.production.search.* 79
net.verisign.* 208
net.virtualearth.* 166
net.virtualearth.dev.ecn 147
net.windows-ppe.login 176
net.windows.* 51
net.windows.accesscontrol.accounts 17
net.windows.core.blob.* 119
net.windows.core.blob.mlccdn 66
net.windows.enterpriseregistration 114
net.windows.graph 27
net.windows.login 31
net.windows.login.* 52
net.windows.mediaservices.keydelivery.* 64
net.windows.mediaservices.streaming.* 65
org.letsencrypt.int-x3.cert 211
org.letsencrypt.int-x3.ocsp 243
===duplicate===
officeclient.microsoft.com  is duplicated.
ajax.aspnetcdn.com  is duplicated.
auth.gfx.ms  is duplicated.
login.live.com  is duplicated.
cdn.odc.officeapps.live.com  is duplicated.
iida@s400win:~/go/src/sort-url$
```
