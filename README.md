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

## 実行

```bash
iida@s400win:~/go/src/dlcheck$ go run main.go
ai.cortana.* 244
com.aadrm.* 144
com.acompli.www 179
com.aspnetcdn.ajax 164
com.assets-yammer.* 193
com.azure.informationprotection.* 146
com.azure.management 115
...

===duplicate===
officeclient.microsoft.com  is duplicated.
ajax.aspnetcdn.com  is duplicated.
auth.gfx.ms  is duplicated.
login.live.com  is duplicated.
cdn.odc.officeapps.live.com  is duplicated.
```

<br><br>

## 仕組み

XMLファイルから読み取ったURLを逆順にします。
たとえば、`outlook.office.com` であれば `com.office.outlook` というようにします。

これをradix treeに格納します。
radix treeでは辞書順にソートされていますので、ツリーをルートから辿っていけばソートされて表示されます。

radix treeに格納する際には同じキーがすでに存在するかわかりますので、重複したエントリの存在もわかります。
